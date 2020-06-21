package keeper

// TODO: Define if your module needs Parameters, if not this can be deleted

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// ParamTable for staking module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&types.Params{})
}

func (k Keeper) CurationWindow(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyCurationWindow, &res)
	return
}

// GetParams returns the total set of stake parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the stake parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
