package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v17/x/globalfee/types"
)

// IterateCodeAuthorizations executes the given func on all the code authorizations
func (k Keeper) IterateCodeAuthorizations(ctx sdk.Context, cb func(types.CodeAuthorization) bool) {
	iterator, err := k.CodeAuthorizations.Iterate(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		codeAuth, err := iterator.Value()
		if err != nil {
			panic(err)
		}
		if cb(codeAuth) {
			return
		}
	}
}

// GetCodeAuthorization gets any authorizations set up for the given code id
func (k Keeper) GetCodeAuthorization(ctx sdk.Context, codeID uint64) (types.CodeAuthorization, error) {
	return k.CodeAuthorizations.Get(ctx, codeID)
}

// SetCodeAuthorization creates of updates provided authorizations for given code id
func (k Keeper) SetCodeAuthorization(ctx sdk.Context, ca types.CodeAuthorization) error {
	if err := ca.Validate(); err != nil {
		return err
	}

	if k.wasmKeeper.GetCodeInfo(ctx, ca.GetCodeID()) == nil {
		return types.ErrCodeIDNotExist
	}

	return k.CodeAuthorizations.Set(ctx, ca.GetCodeID(), ca)
}

// DeleteCodeAuthorization deletes any existing authorizations for given code id
func (k Keeper) DeleteCodeAuthorization(ctx sdk.Context, codeID uint64) error {
	return k.CodeAuthorizations.Remove(ctx, codeID)
}

// HasCodeAuthorization checks if the given code id has any authorizations
func (k Keeper) HasCodeAuthorization(ctx sdk.Context, codeID uint64) bool {
	has, err := k.CodeAuthorizations.Has(ctx, codeID)
	if err != nil {
		panic(err)
	}
	return has
}
