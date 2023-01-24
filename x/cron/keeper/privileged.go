package keeper

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/public-awesome/stargaze/v8/x/cron/types"
)

func (k Keeper) SetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) error {
	if k.HasContractInfo(ctx, contractAddr) {
		store := ctx.KVStore(k.storeKey)
		store.Set(types.PrivilegedContractsKey(contractAddr), []byte{1})

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

func (k Keeper) UnsetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) error {
	if k.HasContractInfo(ctx, contractAddr) {
		if k.IsPrivileged(ctx, contractAddr) {
			store := ctx.KVStore(k.storeKey)
			store.Delete(types.PrivilegedContractsKey(contractAddr))

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

func (k Keeper) IsPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.PrivilegedContractsKey(contractAddr))
}

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
