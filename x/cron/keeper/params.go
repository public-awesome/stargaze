package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v12/x/cron/types"
)

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	if err := params.Validate(); err != nil {
		panic(err)
	}
	k.paramstore.SetParamSet(ctx, &params)
}

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) IsAdminAddress(ctx sdk.Context, address string) bool {
	privilegedAddresses := k.GetParams(ctx).AdminAddresses
	for _, paddr := range privilegedAddresses {
		if address == paddr {
			return true
		}
	}
	return false
}
