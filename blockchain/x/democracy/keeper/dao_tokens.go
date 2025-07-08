
// Gestion des tokens DAO pour la gouvernance décentralisée

package keeper

import (
        "fmt"
        "strings"
        "time"

        sdk "github.com/cosmos/cosmos-sdk/types"
        sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
        banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

        "github.com/librademos/librademos-chain/x/democracy/types"
)

// CreateDAOToken creates a new DAO-specific token
func (k Keeper) CreateDAOToken(ctx sdk.Context, location string, level types.DAOLevel, parentDenom string) (types.DAOToken, error) {
        denom := types.CreateDAOTokenDenom(location, level)
        
        // Validate denomination
        if err := types.ValidateDAOTokenDenom(denom); err != nil {
                return types.DAOToken{}, err
        }

        // Check if token already exists
        if k.DAOTokenExists(ctx, denom) {
                return types.DAOToken{}, sdkerrors.Wrapf(types.ErrDAOTokenExists, "token %s already exists", denom)
        }

        // Create token metadata for bank module
        metadata := banktypes.Metadata{
                Description: fmt.Sprintf("LibraDemos Chain token for %s", location),
                DenomUnits: []banktypes.DenomUnit{
                        {
                                Denom:    denom,
                                Exponent: 0,
                                Aliases:  []string{fmt.Sprintf("micro%s", denom)},
                        },
                        {
                                Denom:    strings.TrimPrefix(denom, "u"),
                                Exponent: 6,
                                Aliases:  []string{strings.ToUpper(strings.TrimPrefix(denom, "u"))},
                        },
                },
                Base:    denom,
                Display: strings.TrimPrefix(denom, "u"),
                Name:    fmt.Sprintf("LibraDemos %s", location),
                Symbol:  strings.ToUpper(strings.TrimPrefix(denom, "u")),
        }

        k.bankKeeper.SetDenomMetaData(ctx, metadata)

        // Calculate initial supply based on expected population
        distribution := types.GetStandardDistribution(level)
        initialSupply := k.calculateInitialSupply(ctx, location, level, distribution)

        // Create DAO token
        daoToken := types.DAOToken{
                Denom:          denom,
                Name:           metadata.Name,
                DAOLocation:    location,
                DAOLevel:       level,
                ParentDAODenom: parentDenom,
                TotalSupply:    initialSupply,
                CreatedAt:      ctx.BlockTime().Format(time.RFC3339),
                Active:         true,
                ExchangeRate:   sdk.NewDecWithPrec(1, 0), // 1:1 initially
                VotingWeight:   k.calculateVotingWeight(ctx, location, level),
        }

        // Store DAO token
        k.SetDAOToken(ctx, daoToken)

        // Mint initial supply to module account
        if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, initialSupply))); err != nil {
                return types.DAOToken{}, err
        }

        // Emit event
        ctx.EventManager().EmitEvent(
                sdk.NewEvent(
                        types.EventTypeCreateDAOToken,
                        sdk.NewAttribute(types.AttributeKeyDAOTokenDenom, denom),
                        sdk.NewAttribute(types.AttributeKeyDAOLocation, location),
                        sdk.NewAttribute(types.AttributeKeyDAOLevel, level.String()),
                        sdk.NewAttribute(types.AttributeKeyTotalSupply, initialSupply.String()),
                ),
        )

        return daoToken, nil
}

// GetDAOToken retrieves a DAO token by denomination
func (k Keeper) GetDAOToken(ctx sdk.Context, denom string) (types.DAOToken, bool) {
        store := ctx.KVStore(k.storeKey)
        key := types.DAOTokenKey(denom)
        bz := store.Get(key)
        if bz == nil {
                return types.DAOToken{}, false
        }

        var token types.DAOToken
        k.cdc.MustUnmarshal(bz, &token)
        return token, true
}

// SetDAOToken stores a DAO token
func (k Keeper) SetDAOToken(ctx sdk.Context, token types.DAOToken) {
        store := ctx.KVStore(k.storeKey)
        key := types.DAOTokenKey(token.Denom)
        bz := k.cdc.MustMarshal(&token)
        store.Set(key, bz)
}

// DAOTokenExists checks if a DAO token exists
func (k Keeper) DAOTokenExists(ctx sdk.Context, denom string) bool {
        _, found := k.GetDAOToken(ctx, denom)
        return found
}

// GetAllDAOTokens returns all DAO tokens
func (k Keeper) GetAllDAOTokens(ctx sdk.Context) []types.DAOToken {
        var tokens []types.DAOToken
        k.IterateDAOTokens(ctx, func(token types.DAOToken) bool {
                tokens = append(tokens, token)
                return false
        })
        return tokens
}

// IterateDAOTokens iterates over all DAO tokens
func (k Keeper) IterateDAOTokens(ctx sdk.Context, cb func(token types.DAOToken) bool) {
        store := ctx.KVStore(k.storeKey)
        iterator := sdk.KVStorePrefixIterator(store, types.DAOTokensKeyPrefix)

        defer iterator.Close()
        for ; iterator.Valid(); iterator.Next() {
                var token types.DAOToken
                k.cdc.MustUnmarshal(iterator.Value(), &token)

                if cb(token) {
                        break
                }
        }
}

// AllocateDAOTokensToCitizen distributes DAO tokens to a newly verified citizen
func (k Keeper) AllocateDAOTokensToCitizen(ctx sdk.Context, citizenAddr sdk.AccAddress, location string) error {
        // Get citizen to verify location
        citizen, found := k.GetCitizen(ctx, citizenAddr)
        if !found {
                return types.ErrCitizenNotFound
        }

        // Find all applicable DAO tokens for this location
        daoTokens := k.getApplicableDAOTokens(ctx, citizen.Location)

        for _, token := range daoTokens {
                distribution := types.GetStandardDistribution(token.DAOLevel)
                
                // Allocate initial tokens
                if distribution.InitialAllocation.GT(sdk.ZeroInt()) {
                        coins := sdk.NewCoins(sdk.NewCoin(token.Denom, distribution.InitialAllocation))
                        
                        if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, citizenAddr, coins); err != nil {
                                k.Logger(ctx).Error("Failed to allocate DAO tokens", "citizen", citizenAddr, "token", token.Denom, "error", err)
                                continue
                        }

                        // Record allocation
                        k.recordTokenAllocation(ctx, citizenAddr, token.Denom, distribution.InitialAllocation, "initial_allocation")
                }
        }

        return nil
}

// RewardCitizenForVoting rewards a citizen with DAO tokens for voting
func (k Keeper) RewardCitizenForVoting(ctx sdk.Context, voterAddr sdk.AccAddress, proposalLocation string) error {
        // Get applicable DAO tokens based on proposal location
        daoTokens := k.getApplicableDAOTokens(ctx, proposalLocation)

        for _, token := range daoTokens {
                distribution := types.GetStandardDistribution(token.DAOLevel)
                
                if distribution.VotingReward.GT(sdk.ZeroInt()) {
                        coins := sdk.NewCoins(sdk.NewCoin(token.Denom, distribution.VotingReward))
                        
                        if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, voterAddr, coins); err != nil {
                                k.Logger(ctx).Error("Failed to reward voting", "voter", voterAddr, "token", token.Denom, "error", err)
                                continue
                        }

                        k.recordTokenAllocation(ctx, voterAddr, token.Denom, distribution.VotingReward, "voting_reward")
                }
        }

        return nil
}

// RewardCitizenForProposal rewards a citizen with DAO tokens for creating a quality proposal
func (k Keeper) RewardCitizenForProposal(ctx sdk.Context, proposerAddr sdk.AccAddress, proposalLocation string) error {
        daoTokens := k.getApplicableDAOTokens(ctx, proposalLocation)

        for _, token := range daoTokens {
                distribution := types.GetStandardDistribution(token.DAOLevel)
                
                if distribution.ProposalReward.GT(sdk.ZeroInt()) {
                        coins := sdk.NewCoins(sdk.NewCoin(token.Denom, distribution.ProposalReward))
                        
                        if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, proposerAddr, coins); err != nil {
                                k.Logger(ctx).Error("Failed to reward proposal", "proposer", proposerAddr, "token", token.Denom, "error", err)
                                continue
                        }

                        k.recordTokenAllocation(ctx, proposerAddr, token.Denom, distribution.ProposalReward, "proposal_reward")
                }
        }

        return nil
}

// DistributeMonthlyUBI distributes monthly Universal Basic Income in DAO tokens
func (k Keeper) DistributeMonthlyUBI(ctx sdk.Context) error {
        // This would typically be called by a monthly cron job or governance proposal
        
        // Get all active citizens
        citizens := k.GetAllCitizens(ctx)
        
        for _, citizen := range citizens {
                if !citizen.Active {
                        continue
                }

                citizenAddr := citizen.Address
                daoTokens := k.getApplicableDAOTokens(ctx, citizen.Location)

                for _, token := range daoTokens {
                        distribution := types.GetStandardDistribution(token.DAOLevel)
                        
                        if distribution.MonthlyUBI.GT(sdk.ZeroInt()) {
                                coins := sdk.NewCoins(sdk.NewCoin(token.Denom, distribution.MonthlyUBI))
                                
                                if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, citizenAddr, coins); err != nil {
                                        k.Logger(ctx).Error("Failed to distribute UBI", "citizen", citizenAddr, "token", token.Denom, "error", err)
                                        continue
                                }

                                k.recordTokenAllocation(ctx, citizenAddr, token.Denom, distribution.MonthlyUBI, "monthly_ubi")
                        }
                }
        }

        ctx.EventManager().EmitEvent(
                sdk.NewEvent(
                        types.EventTypeDistributeUBI,
                        sdk.NewAttribute(types.AttributeKeyBlockHeight, fmt.Sprintf("%d", ctx.BlockHeight())),
                        sdk.NewAttribute(types.AttributeKeyCitizenCount, fmt.Sprintf("%d", len(citizens))),
                ),
        )

        return nil
}

// ExchangeDAOTokens allows exchanging between DAO tokens in the hierarchy
func (k Keeper) ExchangeDAOTokens(ctx sdk.Context, fromAddr sdk.AccAddress, fromDenom, toDenom string, amount sdk.Int) error {
        // Get both tokens
        fromToken, foundFrom := k.GetDAOToken(ctx, fromDenom)
        if !foundFrom {
                return sdkerrors.Wrapf(types.ErrDAOTokenNotFound, "from token %s not found", fromDenom)
        }

        toToken, foundTo := k.GetDAOToken(ctx, toDenom)
        if !foundTo {
                return sdkerrors.Wrapf(types.ErrDAOTokenNotFound, "to token %s not found", toDenom)
        }

        // Validate exchange is allowed (must be in same hierarchy or global)
        if !k.isExchangeAllowed(fromToken, toToken) {
                return sdkerrors.Wrap(types.ErrInvalidExchange, "exchange not allowed between these tokens")
        }

        // Calculate exchange rate
        exchangeRate := k.calculateExchangeRate(fromToken, toToken)
        toAmount := sdk.NewDecFromInt(amount).Mul(exchangeRate).TruncateInt()

        // Execute exchange
        fromCoins := sdk.NewCoins(sdk.NewCoin(fromDenom, amount))
        toCoins := sdk.NewCoins(sdk.NewCoin(toDenom, toAmount))

        // Send from tokens to module
        if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, fromAddr, types.ModuleName, fromCoins); err != nil {
                return err
        }

        // Send to tokens to user
        if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, fromAddr, toCoins); err != nil {
                return err
        }

        // Emit exchange event
        ctx.EventManager().EmitEvent(
                sdk.NewEvent(
                        types.EventTypeExchangeDAOTokens,
                        sdk.NewAttribute(types.AttributeKeyExchanger, fromAddr.String()),
                        sdk.NewAttribute(types.AttributeKeyFromDenom, fromDenom),
                        sdk.NewAttribute(types.AttributeKeyToDenom, toDenom),
                        sdk.NewAttribute(types.AttributeKeyFromAmount, amount.String()),
                        sdk.NewAttribute(types.AttributeKeyToAmount, toAmount.String()),
                        sdk.NewAttribute(types.AttributeKeyExchangeRate, exchangeRate.String()),
                ),
        )

        return nil
}

// Helper functions

func (k Keeper) calculateInitialSupply(ctx sdk.Context, location string, level types.DAOLevel, distribution types.DAOTokenDistribution) sdk.Int {
        // Estimate population based on level and location
        var estimatedPopulation int64

        switch level {
        case types.DAOLevelGlobal:
                estimatedPopulation = 67000000 // France population
        case types.DAOLevelRegional:
                estimatedPopulation = 5000000 // Average region
        case types.DAOLevelDepartmental:
                estimatedPopulation = 650000 // Average department
        case types.DAOLevelCommunal:
                estimatedPopulation = 50000 // Average commune
        default:
                estimatedPopulation = 10000
        }

        // Supply = (initial allocation + 12 months UBI + governance reserves) * population
        yearlySupply := distribution.InitialAllocation.Add(distribution.MonthlyUBI.MulRaw(12))
        reserves := yearlySupply.MulRaw(2) // 2x for reserves
        totalPerCitizen := yearlySupply.Add(reserves)

        return totalPerCitizen.MulRaw(estimatedPopulation)
}

func (k Keeper) calculateVotingWeight(ctx sdk.Context, location string, level types.DAOLevel) sdk.Dec {
        // Weight based on population ratios - simplified for now
        switch level {
        case types.DAOLevelRegional:
                return sdk.NewDecWithPrec(10, 0) // 10% average
        case types.DAOLevelDepartmental:
                return sdk.NewDecWithPrec(5, 0) // 5% average
        case types.DAOLevelCommunal:
                return sdk.NewDecWithPrec(1, 0) // 1% average
        default:
                return sdk.NewDecWithPrec(1, 0)
        }
}

func (k Keeper) getApplicableDAOTokens(ctx sdk.Context, location string) []types.DAOToken {
        var applicableTokens []types.DAOToken

        // Always include global LDC
        if globalToken, found := k.GetDAOToken(ctx, "uldc"); found {
                applicableTokens = append(applicableTokens, globalToken)
        }

        // Get location-specific tokens based on hierarchy
        // This is simplified - in practice you'd parse the location to find all applicable levels
        
        allTokens := k.GetAllDAOTokens(ctx)
        for _, token := range allTokens {
                if k.isTokenApplicableForLocation(token, location) {
                        applicableTokens = append(applicableTokens, token)
                }
        }

        return applicableTokens
}

func (k Keeper) isTokenApplicableForLocation(token types.DAOToken, citizenLocation string) bool {
        // Simplified location matching - in practice this would be more sophisticated
        // matching the hierarchical structure of French administrative divisions
        
        if token.DAOLevel == types.DAOLevelGlobal {
                return true
        }

        // For now, simple string contains check
        return strings.Contains(strings.ToLower(citizenLocation), strings.ToLower(token.DAOLocation))
}

func (k Keeper) isExchangeAllowed(fromToken, toToken types.DAOToken) bool {
        // Allow exchange to global LDC from any token
        if toToken.DAOLevel == types.DAOLevelGlobal {
                return true
        }

        // Allow exchange from global LDC to any token
        if fromToken.DAOLevel == types.DAOLevelGlobal {
                return true
        }

        // Allow exchange within the same hierarchy
        return fromToken.ParentDAODenom == toToken.ParentDAODenom ||
                fromToken.Denom == toToken.ParentDAODenom ||
                toToken.Denom == fromToken.ParentDAODenom
}

func (k Keeper) calculateExchangeRate(fromToken, toToken types.DAOToken) sdk.Dec {
        // For now, all tokens have 1:1 exchange rate
        // In practice, this could be based on:
        // - Economic activity in the region
        // - Population density
        // - Governance activity
        // - Supply and demand
        
        return sdk.NewDecWithPrec(1, 0)
}

func (k Keeper) recordTokenAllocation(ctx sdk.Context, recipient sdk.AccAddress, denom string, amount sdk.Int, reason string) {
        // Record allocation for audit trail
        ctx.EventManager().EmitEvent(
                sdk.NewEvent(
                        types.EventTypeTokenAllocation,
                        sdk.NewAttribute(types.AttributeKeyRecipient, recipient.String()),
                        sdk.NewAttribute(types.AttributeKeyTokenDenom, denom),
                        sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
                        sdk.NewAttribute(types.AttributeKeyReason, reason),
                        sdk.NewAttribute(types.AttributeKeyBlockHeight, fmt.Sprintf("%d", ctx.BlockHeight())),
                ),
        )
}
