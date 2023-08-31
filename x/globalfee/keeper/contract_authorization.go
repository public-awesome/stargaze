package keeper

import (
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v12/x/globalfee/types"
)

// IterateContractAuthorizations executes the given func on all the contract authorizations
func (k Keeper) IterateContractAuthorizations(ctx sdk.Context, cb func(types.ContractAuthorization) bool) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.ContractAuthorizationPrefix)
	iter := prefixStore.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var ca types.ContractAuthorization
		k.cdc.MustUnmarshal(iter.Value(), &ca)
		// cb returns true to stop early
		if cb(ca) {
			return
		}
	}
}

// GetContractAuthorization gets any authorizations set up for the given contract address
func (k Keeper) GetContractAuthorization(ctx sdk.Context, contractAddr sdk.AccAddress) (types.ContractAuthorization, bool) {
	store := ctx.KVStore(k.storeKey)

	var ca types.ContractAuthorization
	bz := store.Get(types.GetContractAuthorizationPrefix(contractAddr))
	if bz == nil {
		return ca, false
	}

	k.cdc.MustUnmarshal(bz, &ca)
	return ca, true
}

// SetContractAuthorization creates of updates provided authorizations for given contract address
func (k Keeper) SetContractAuthorization(ctx sdk.Context, ca types.ContractAuthorization) error {
	if err := ca.Validate(); err != nil {
		return err
	}

	if !k.wasmKeeper.HasContractInfo(ctx, sdk.MustAccAddressFromBech32(ca.GetContractAddress())) {
		return types.ErrContractNotExist
	}

	store := ctx.KVStore(k.storeKey)
	value := k.cdc.MustMarshal(&ca)

	store.Set(types.GetContractAuthorizationPrefix(sdk.MustAccAddressFromBech32(ca.ContractAddress)), value)
	return nil
}

// DeleteContractAuthorization deletes any existing authorizations for given contract address
func (k Keeper) DeleteContractAuthorization(ctx sdk.Context, contractAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)

	store.Delete(types.GetContractAuthorizationPrefix(contractAddr))
}
