package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v15/x/globalfee/types"
)

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	return k.Params.Set(ctx, params)
}

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params, err error) {
	return k.Params.Get(ctx)
}

func (k Keeper) IsPrivilegedAddress(ctx sdk.Context, address string) bool {
	params, err := k.GetParams(ctx)
	if err != nil {
		return false
	}
	privilegedAddresses := params.PrivilegedAddresses
	for _, paddr := range privilegedAddresses {
		if address == paddr {
			return true
		}
	}
	return false
}
