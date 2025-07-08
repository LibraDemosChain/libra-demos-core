package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/librademos/librademos-chain/x/democracy/types"
)

// SubmitProposal creates a new proposal given a content
func (keeper Keeper) SubmitProposal(ctx sdk.Context, content types.Content, proposer sdk.AccAddress, initialDeposit sdk.Coins) (types.Proposal, error) {
	if err := keeper.validateInitialDeposit(ctx, initialDeposit); err != nil {
		return types.Proposal{}, err
	}

	proposalID, err := keeper.GetProposalID(ctx)
	if err != nil {
		return types.Proposal{}, err
	}

	submitTime := ctx.BlockHeader().Time
	depositPeriod := keeper.GetDepositParams(ctx).MaxDepositPeriod

	proposal := types.Proposal{
		ProposalId:       proposalID,
		Content:          content,
		Status:           types.StatusDepositPeriod,
		FinalTallyResult: types.EmptyTallyResult(),
		SubmitTime:       submitTime,
		DepositEndTime:   submitTime.Add(depositPeriod),
		TotalDeposit:     sdk.NewCoins(),
		VotingStartTime:  time.Time{},
		VotingEndTime:    time.Time{},
	}

	keeper.SetProposal(ctx, proposal)
	keeper.InsertInactiveProposalQueue(ctx, proposalID, proposal.DepositEndTime)
	keeper.SetProposalID(ctx, proposalID+1)

	// called after a proposal is submitted
	keeper.AfterProposalSubmission(ctx, proposalID)

	err = keeper.AddDeposit(ctx, proposalID, proposer, initialDeposit)
	if err != nil {
		return proposal, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSubmitProposal,
			sdk.NewAttribute(types.AttributeKeyProposalID, fmt.Sprintf("%d", proposalID)),
			sdk.NewAttribute(types.AttributeKeyProposalType, content.ProposalType()),
		),
	)

	return proposal, nil
}

// GetProposal get proposal from store by ProposalID
func (keeper Keeper) GetProposal(ctx sdk.Context, proposalID uint64) (proposal types.Proposal, found bool) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(types.ProposalKey(proposalID))
	if bz == nil {
		return
	}

	keeper.cdc.MustUnmarshal(bz, &proposal)
	return proposal, true
}

// SetProposal set a proposal to store
func (keeper Keeper) SetProposal(ctx sdk.Context, proposal types.Proposal) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshal(&proposal)
	store.Set(types.ProposalKey(proposal.ProposalId), bz)
}

// DeleteProposal deletes a proposal from store
func (keeper Keeper) DeleteProposal(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	proposal, found := keeper.GetProposal(ctx, proposalID)
	if !found {
		return
	}

	keeper.RemoveFromActiveProposalQueue(ctx, proposalID, proposal.VotingEndTime)
	keeper.RemoveFromInactiveProposalQueue(ctx, proposalID, proposal.DepositEndTime)
	store.Delete(types.ProposalKey(proposalID))
}

// ActivateVotingPeriod activates the voting period of a proposal
func (keeper Keeper) ActivateVotingPeriod(ctx sdk.Context, proposal types.Proposal) {
	proposal.VotingStartTime = ctx.BlockHeader().Time
	proposal.Status = types.StatusVotingPeriod
	votingPeriod := keeper.GetVotingParams(ctx).VotingPeriod
	proposal.VotingEndTime = proposal.VotingStartTime.Add(votingPeriod)

	keeper.SetProposal(ctx, proposal)

	keeper.RemoveFromInactiveProposalQueue(ctx, proposal.ProposalId, proposal.DepositEndTime)
	keeper.InsertActiveProposalQueue(ctx, proposal.ProposalId, proposal.VotingEndTime)
}

// ValidateInitialDeposit validates if initial deposit is greater than or equal to the minimum
// required at the time of proposal submission. This threshold amount is determined by
// the deposit parameters. Returns nil on success, error otherwise.
func (keeper Keeper) validateInitialDeposit(ctx sdk.Context, initialDeposit sdk.Coins) error {
	minInitialDepositRatio := keeper.GetDepositParams(ctx).MinInitialDepositRatio
	if minInitialDepositRatio.IsZero() {
		return nil
	}

	minDepositCoins := keeper.GetDepositParams(ctx).MinDeposit
	for i := range minDepositCoins {
		minDepositCoins[i].Amount = sdk.NewDecFromInt(minDepositCoins[i].Amount).Mul(minInitialDepositRatio).TruncateInt()
	}

	if !initialDeposit.IsAllGTE(minDepositCoins) {
		return sdkerrors.Wrapf(types.ErrInsufficientDeposit,
			"insufficient initial deposit: got %s, required %s", initialDeposit, minDepositCoins)
	}

	return nil
}

// Tally iterates over the votes and updates the tally of a proposal based on the voting power of the voters
func (keeper Keeper) Tally(ctx sdk.Context, proposal types.Proposal) (passes bool, burnDeposits bool, tallyResults types.TallyResult) {
	results := make(map[types.VoteOption]sdk.Dec)
	results[types.OptionYes] = sdk.ZeroDec()
	results[types.OptionAbstain] = sdk.ZeroDec()
	results[types.OptionNo] = sdk.ZeroDec()
	results[types.OptionNoWithVeto] = sdk.ZeroDec()

	totalVotingPower := sdk.ZeroDec()
	currValidators := make(map[string]types.ValidatorGovInfo)

	// fetch all the bonded validators, insert them into currValidators
	keeper.sk.IterateBondedValidatorsByPower(ctx, func(index int64, validator types.StakingValidator) (stop bool) {
		currValidators[validator.GetOperator().String()] = types.NewValidatorGovInfo(
			validator.GetOperator(),
			validator.GetBondedTokens(),
			validator.GetDelegatorShares(),
			sdk.ZeroDec(),
			types.WeightedVoteOptions{},
		)

		return false
	})

	keeper.IterateVotes(ctx, proposal.ProposalId, func(vote types.Vote) bool {
		// if validator, just record it in the map
		voter, err := sdk.AccAddressFromBech32(vote.Voter)
		if err != nil {
			panic(err)
		}

		valAddrStr := sdk.ValAddress(voter.Bytes()).String()
		if val, ok := currValidators[valAddrStr]; ok {
			val.Vote = vote.Options
			currValidators[valAddrStr] = val
		}

		// iterate over all delegations from voter, deduct from any delegated-to validators
		keeper.sk.IterateDelegations(ctx, voter, func(index int64, delegation types.StakingDelegation) (stop bool) {
			valAddrStr := delegation.GetValidatorAddr().String()

			if val, ok := currValidators[valAddrStr]; ok {
				// There is no need to handle the special case that validator address equal to voter address.
				// Because voter's voting power will tally again even if there will deduct voter's voting power from validator.
				val.DelegatorDeductions = val.DelegatorDeductions.Add(delegation.GetShares())
				currValidators[valAddrStr] = val

				delegatorShare, err := keeper.sk.ValidatorByConsAddr(ctx, val.Address.Bytes()).SharesFromTokens(delegation.GetShares().TruncateInt())
				if err != nil {
					panic(err)
				}

				for _, option := range vote.Options {
					subPower := option.Weight.Mul(delegatorShare)
					results[option.Option] = results[option.Option].Add(subPower)
				}
				totalVotingPower = totalVotingPower.Add(delegatorShare)
			}

			return false
		})

		keeper.deleteVote(ctx, vote.ProposalId, voter)
		return false
	})

	// iterate over the validators again to tally their voting power
	for _, val := range currValidators {
		if len(val.Vote) == 0 {
			continue
		}

		sharesAfterDeductions := val.DelegatorShares.Sub(val.DelegatorDeductions)
		votingPower := sharesAfterDeductions

		for _, option := range val.Vote {
			subPower := option.Weight.Mul(votingPower)
			results[option.Option] = results[option.Option].Add(subPower)
		}
		totalVotingPower = totalVotingPower.Add(votingPower)
	}

	tallyParams := keeper.GetTallyParams(ctx)
	tallyResults = types.NewTallyResultFromMap(results)

	// TODO: Upgrade the spec to cover all of these cases & remove pseudocode.
	// If there is no staked coins, the proposal fails
	if keeper.sk.TotalBondedTokens(ctx).IsZero() {
		return false, false, tallyResults
	}

	// If there is not enough quorum of votes, the proposal fails
	percentVoting := totalVotingPower.Quo(sdk.NewDecFromInt(keeper.sk.TotalBondedTokens(ctx)))
	quorum, _ := sdk.NewDecFromStr(tallyParams.Quorum)
	if percentVoting.LT(quorum) {
		return false, true, tallyResults
	}

	// If no one votes (everyone abstains), proposal fails
	if totalVotingPower.Sub(results[types.OptionAbstain]).Equal(sdk.ZeroDec()) {
		return false, false, tallyResults
	}

	// If more than 1/3 of voters veto, proposal fails
	vetoThreshold, _ := sdk.NewDecFromStr(tallyParams.VetoThreshold)
	if results[types.OptionNoWithVeto].Quo(totalVotingPower).GT(vetoThreshold) {
		return false, true, tallyResults
	}

	// If more than 1/2 of non-abstaining voters vote Yes, proposal passes
	threshold, _ := sdk.NewDecFromStr(tallyParams.Threshold)
	if results[types.OptionYes].Quo(totalVotingPower.Sub(results[types.OptionAbstain])).GT(threshold) {
		return true, false, tallyResults
	}

	// If more than 1/2 of non-abstaining voters vote No, proposal fails
	return false, false, tallyResults
}
