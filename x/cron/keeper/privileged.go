package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v8/x/cron/types"
)

func (k Keeper) SetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PrivilegedContractsKey(contractAddr), []byte{1})
}

func (k Keeper) UnsetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PrivilegedContractsKey(contractAddr))
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
