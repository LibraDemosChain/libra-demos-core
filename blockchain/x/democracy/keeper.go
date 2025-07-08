package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/librademos/librademos-chain/x/democracy/types"
)

// Keeper of the democracy store
type Keeper struct {
	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	authKeeper types.AccountKeeper
	bankKeeper types.BankKeeper

	// the address capable of executing a MsgUpdateParams message
	authority string
}

// NewKeeper creates a democracy keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	authKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	authority string,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		paramstore: ps,
		authKeeper: authKeeper,
		bankKeeper: bankKeeper,
		authority:  authority,
	}
}

// Logger returns a module-specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// GetAuthority returns the x/democracy module's authority
func (k Keeper) GetAuthority() string {
	return k.authority
}

// SetProposal sets a proposal to store
func (k Keeper) SetProposal(ctx sdk.Context, proposal types.Proposal) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&proposal)
	store.Set(types.ProposalKey(proposal.ProposalId), bz)
}

// GetProposal gets a proposal from store by ProposalID
func (k Keeper) GetProposal(ctx sdk.Context, proposalID uint64) (proposal types.Proposal, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ProposalKey(proposalID))
	if bz == nil {
		return proposal, false
	}

	k.cdc.MustUnmarshal(bz, &proposal)
	return proposal, true
}

// DeleteProposal deletes a proposal from store
func (k Keeper) DeleteProposal(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(k.storeKey)
	proposal, found := k.GetProposal(ctx, proposalID)
	if !found {
		return
	}

	k.RemoveFromActiveProposalQueue(ctx, proposalID, proposal.VotingEndTime)
	store.Delete(types.ProposalKey(proposalID))
}

// GetProposalID gets the highest proposal ID
func (k Keeper) GetProposalID(ctx sdk.Context) (proposalID uint64, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ProposalIDKey)
	if bz == nil {
		return 0, fmt.Errorf("initial proposal ID hasn't been set")
	}

	proposalID = types.GetProposalIDFromBytes(bz)
	return proposalID, nil
}

// SetProposalID sets the new proposal ID to the store
func (k Keeper) SetProposalID(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ProposalIDKey, types.GetProposalIDBytes(proposalID))
}

// ActivateVotingPeriod activates the voting period of a proposal
func (k Keeper) ActivateVotingPeriod(ctx sdk.Context, proposal types.Proposal) {
	proposal.VotingStartTime = ctx.BlockHeader().Time
	proposal.Status = types.StatusVotingPeriod
	depositParams := k.GetDepositParams(ctx)
	proposal.VotingEndTime = proposal.VotingStartTime.Add(depositParams.VotingPeriod)

	k.SetProposal(ctx, proposal)

	k.RemoveFromInactiveProposalQueue(ctx, proposal.ProposalId, proposal.DepositEndTime)
	k.InsertActiveProposalQueue(ctx, proposal.ProposalId, proposal.VotingEndTime)
}

// IterateProposals iterates over all proposals and performs a callback function
func (k Keeper) IterateProposals(ctx sdk.Context, cb func(proposal types.Proposal) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ProposalsKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var proposal types.Proposal
		k.cdc.MustUnmarshal(iterator.Value(), &proposal)

		if cb(proposal) {
			break
		}
	}
}

// IterateActiveProposalsQueue iterates over proposals in the active proposal queue
func (k Keeper) IterateActiveProposalsQueue(ctx sdk.Context, endTime time.Time, cb func(proposal types.Proposal) (stop bool)) {
	iterator := k.ActiveProposalQueueIterator(ctx, endTime)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		proposalID, _ := types.SplitActiveProposalQueueKey(iterator.Key())
		proposal, found := k.GetProposal(ctx, proposalID)
		if !found {
			panic(fmt.Sprintf("proposal %d does not exist", proposalID))
		}

		if cb(proposal) {
			break
		}
	}
}

// ActiveProposalQueueIterator returns an sdk.Iterator for all the proposals in the active queue
func (k Keeper) ActiveProposalQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(types.ActiveProposalQueuePrefix, sdk.PrefixEndBytes(types.ActiveProposalQueueKey(endTime, 0)))
}

// InactiveProposalQueueIterator returns an sdk.Iterator for all the proposals in the inactive queue
func (k Keeper) InactiveProposalQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(types.InactiveProposalQueuePrefix, sdk.PrefixEndBytes(types.InactiveProposalQueueKey(endTime, 0)))
}

// InsertActiveProposalQueue inserts a ProposalID into the active proposal queue
func (k Keeper) InsertActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	bz := types.GetProposalIDBytes(proposalID)
	store.Set(types.ActiveProposalQueueKey(endTime, proposalID), bz)
}

// RemoveFromActiveProposalQueue removes a proposalID from the Active Proposal Queue
func (k Keeper) RemoveFromActiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ActiveProposalQueueKey(endTime, proposalID))
}

// InsertInactiveProposalQueue inserts a ProposalID into the inactive proposal queue
func (k Keeper) InsertInactiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	bz := types.GetProposalIDBytes(proposalID)
	store.Set(types.InactiveProposalQueueKey(endTime, proposalID), bz)
}

// RemoveFromInactiveProposalQueue removes a proposalID from the Inactive Proposal Queue
func (k Keeper) RemoveFromInactiveProposalQueue(ctx sdk.Context, proposalID uint64, endTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.InactiveProposalQueueKey(endTime, proposalID))
}
