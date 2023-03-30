package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v9/x/globalfee/types"
)

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

func (k Keeper) SetContractAuthorization(ctx sdk.Context, ca types.ContractAuthorization) error {
	if err := ca.Validate(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	value := k.cdc.MustMarshal(&ca)

	store.Set(types.GetContractAuthorizationPrefix(sdk.MustAccAddressFromBech32(ca.ContractAddress)), value)
	return nil
}

func (k Keeper) DeleteContractAuthorization(ctx sdk.Context, contractAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)

	store.Delete(types.GetContractAuthorizationPrefix(contractAddr))
}
