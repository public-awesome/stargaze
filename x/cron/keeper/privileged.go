package keeper

import (
	"cosmossdk.io/store/prefix"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/public-awesome/stargaze/v14/x/cron/types"
)

// SetPrivileged checks if the given contract exists and adds it to the list of privilege contracts
func (k Keeper) SetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) error {
	if k.wasmKeeper.HasContractInfo(ctx, contractAddr) {
		if !k.IsPrivileged(ctx, contractAddr) {
			if err := k.PrivilegedContracts.Set(ctx, contractAddr.Bytes(), []byte{1}); err != nil {
				return err
			}
		}
		event := sdk.NewEvent(
			types.EventTypeSetContractPriviledge,
			sdk.NewAttribute(wasmtypes.AttributeKeyContractAddr, contractAddr.String()),
		)
		ctx.EventManager().EmitEvent(event)
	} else {
		return types.ErrContractDoesNotExist
	}
	return nil
}

// UnsetPrivileged checks if the given contract exists and if it has privilege and remove it from the list of privileg contracts
func (k Keeper) UnsetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) error {
	if k.wasmKeeper.HasContractInfo(ctx, contractAddr) {
		if k.IsPrivileged(ctx, contractAddr) {
			if err := k.PrivilegedContracts.Remove(ctx, contractAddr.Bytes()); err != nil {
				return err
			}

			event := sdk.NewEvent(
				types.EventTypeUnsetContractPriviledge,
				sdk.NewAttribute(wasmtypes.AttributeKeyContractAddr, contractAddr.String()),
			)
			ctx.EventManager().EmitEvent(event)
		} else {
			return types.ErrContractPrivilegeNotSet
		}
	} else {
		return types.ErrContractDoesNotExist
	}
	return nil
}

// IsPrivileged returns if the given contract is part of the privilege contract list
func (k Keeper) IsPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) bool {
	has, err := k.PrivilegedContracts.Has(ctx, contractAddr.Bytes())
	if err != nil {
		return false
	}
	return has
}

// IteratePrivileged executes the given func on all the privilege contracts
func (k Keeper) IteratePrivileged(ctx sdk.Context, cb func(sdk.AccAddress) bool) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrivilegedContractsPrefix)
	iter := prefixStore.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		// cb returns true to stop early
		if cb(iter.Key()) {
			return
		}
	}
}
