package keeper

// TODO: Define if your module needs Parameters, if not this can be deleted

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/funding/types"
)

// GetParams returns the total set of funding parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the funding parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}
