package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v18/x/pauser/types"
)

// SetPausedContract stores a paused contract.
func (k Keeper) SetPausedContract(ctx sdk.Context, pc types.PausedContract) error {
	contractAddr, err := sdk.AccAddressFromBech32(pc.ContractAddress)
	if err != nil {
		return err
	}
	return k.PausedContracts.Set(ctx, contractAddr.Bytes(), pc)
}

// GetPausedContract returns a paused contract by address.
func (k Keeper) GetPausedContract(ctx sdk.Context, contractAddr sdk.AccAddress) (types.PausedContract, error) {
	return k.PausedContracts.Get(ctx, contractAddr.Bytes())
}

// DeletePausedContract removes a paused contract.
func (k Keeper) DeletePausedContract(ctx sdk.Context, contractAddr sdk.AccAddress) error {
	return k.PausedContracts.Remove(ctx, contractAddr.Bytes())
}

// IsContractPaused checks if a contract is paused.
func (k Keeper) IsContractPaused(ctx sdk.Context, contractAddr sdk.AccAddress) bool {
	has, err := k.PausedContracts.Has(ctx, contractAddr.Bytes())
	if err != nil {
		return false
	}
	return has
}

// IteratePausedContracts iterates over all paused contracts.
func (k Keeper) IteratePausedContracts(ctx sdk.Context, cb func(types.PausedContract) bool) {
	iterator, err := k.PausedContracts.Iterate(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		pc, err := iterator.Value()
		if err != nil {
			panic(err)
		}
		if cb(pc) {
			return
		}
	}
}
