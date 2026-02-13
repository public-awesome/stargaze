package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v17/x/pauser/types"
)

// SetParams sets the module params.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	return k.Params.Set(ctx, params)
}

// GetParams returns the module params.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params, err error) {
	return k.Params.Get(ctx)
}

// IsPrivilegedAddress checks if the given address is a privileged address.
func (k Keeper) IsPrivilegedAddress(ctx sdk.Context, address string) bool {
	params, err := k.GetParams(ctx)
	if err != nil {
		return false
	}
	for _, paddr := range params.PrivilegedAddresses {
		if address == paddr {
			return true
		}
	}
	return false
}
