package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/librademos/librademos-chain/x/democracy/types"
)

// GetDepositParams returns the current DepositParams from the global param store
func (k Keeper) GetDepositParams(ctx sdk.Context) types.DepositParams {
	var depositParams types.DepositParams
	k.paramstore.Get(ctx, types.KeyDepositParams, &depositParams)
	return depositParams
}

// SetDepositParams sets DepositParams to the global param store
func (k Keeper) SetDepositParams(ctx sdk.Context, depositParams types.DepositParams) {
	k.paramstore.Set(ctx, types.KeyDepositParams, &depositParams)
}

// GetVotingParams returns the current VotingParams from the global param store
func (k Keeper) GetVotingParams(ctx sdk.Context) types.VotingParams {
	var votingParams types.VotingParams
	k.paramstore.Get(ctx, types.KeyVotingParams, &votingParams)
	return votingParams
}

// SetVotingParams sets VotingParams to the global param store
func (k Keeper) SetVotingParams(ctx sdk.Context, votingParams types.VotingParams) {
	k.paramstore.Set(ctx, types.KeyVotingParams, &votingParams)
}

// GetTallyParams returns the current TallyParams from the global param store
func (k Keeper) GetTallyParams(ctx sdk.Context) types.TallyParams {
	var tallyParams types.TallyParams
	k.paramstore.Get(ctx, types.KeyTallyParams, &tallyParams)
	return tallyParams
}

// SetTallyParams sets TallyParams to the global param store
func (k Keeper) SetTallyParams(ctx sdk.Context, tallyParams types.TallyParams) {
	k.paramstore.Set(ctx, types.KeyTallyParams, &tallyParams)
}

// GetParams returns all the democracy module parameters
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	params.DepositParams = k.GetDepositParams(ctx)
	params.VotingParams = k.GetVotingParams(ctx)
	params.TallyParams = k.GetTallyParams(ctx)
	return params
}

// SetParams sets all the democracy module parameters
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.SetDepositParams(ctx, params.DepositParams)
	k.SetVotingParams(ctx, params.VotingParams)
	k.SetTallyParams(ctx, params.TallyParams)
}

// IsCitizen checks if the given address is a verified citizen
func (k Keeper) IsCitizen(ctx sdk.Context, address sdk.AccAddress) bool {
	_, found := k.GetCitizen(ctx, address)
	return found
}

// GetCitizen gets a citizen from the store
func (k Keeper) GetCitizen(ctx sdk.Context, address sdk.AccAddress) (citizen types.Citizen, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.CitizenKey(address))
	if bz == nil {
		return citizen, false
	}

	k.cdc.MustUnmarshal(bz, &citizen)
	return citizen, true
}

// SetCitizen sets a citizen to the store
func (k Keeper) SetCitizen(ctx sdk.Context, citizen types.Citizen) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&citizen)
	store.Set(types.CitizenKey(citizen.Address), bz)
}

// DeleteCitizen deletes a citizen from the store
func (k Keeper) DeleteCitizen(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.CitizenKey(address))
}

// IterateCitizens iterates over all citizens and performs a callback function
func (k Keeper) IterateCitizens(ctx sdk.Context, cb func(citizen types.Citizen) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.CitizensKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var citizen types.Citizen
		k.cdc.MustUnmarshal(iterator.Value(), &citizen)

		if cb(citizen) {
			break
		}
	}
}

// GetAllCitizens returns all citizens from the store
func (k Keeper) GetAllCitizens(ctx sdk.Context) (citizens []types.Citizen) {
	k.IterateCitizens(ctx, func(citizen types.Citizen) bool {
		citizens = append(citizens, citizen)
		return false
	})
	return citizens
}

// AddCitizen adds a new verified citizen
func (k Keeper) AddCitizen(ctx sdk.Context, address sdk.AccAddress, kycHash, location string) error {
	// Check if citizen already exists
	if k.IsCitizen(ctx, address) {
		return types.ErrCitizenAlreadyExists
	}

	citizen := types.Citizen{
		Address:    address,
		KYCHash:    kycHash,
		Location:   location,
		VerifiedAt: ctx.BlockHeader().Time,
		Active:     true,
	}

	k.SetCitizen(ctx, citizen)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeVerifyCitizen,
			sdk.NewAttribute(types.AttributeKeyCitizenAddress, address.String()),
			sdk.NewAttribute(types.AttributeKeyKYCHash, kycHash),
			sdk.NewAttribute(types.AttributeKeyLocation, location),
		),
	)

	return nil
}

// DeactivateCitizen deactivates a citizen (soft delete)
func (k Keeper) DeactivateCitizen(ctx sdk.Context, address sdk.AccAddress) error {
	citizen, found := k.GetCitizen(ctx, address)
	if !found {
		return types.ErrCitizenNotFound
	}

	citizen.Active = false
	k.SetCitizen(ctx, citizen)

	return nil
}

// UpdateCitizenLocation updates a citizen's location
func (k Keeper) UpdateCitizenLocation(ctx sdk.Context, address sdk.AccAddress, newLocation string) error {
	citizen, found := k.GetCitizen(ctx, address)
	if !found {
		return types.ErrCitizenNotFound
	}

	citizen.Location = newLocation
	k.SetCitizen(ctx, citizen)

	return nil
}
