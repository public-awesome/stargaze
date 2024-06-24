package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v14/x/globalfee/types"
)

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(types.ParamsKey, bz)
	return nil
}

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return params
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params
}

func (k Keeper) IsPrivilegedAddress(ctx sdk.Context, address string) bool {
	privilegedAddresses := k.GetParams(ctx).PrivilegedAddresses
	for _, paddr := range privilegedAddresses {
		if address == paddr {
			return true
		}
	}
	return false
}
