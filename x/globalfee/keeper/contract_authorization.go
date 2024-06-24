package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v15/x/globalfee/types"
)

// IterateContractAuthorizations executes the given func on all the contract authorizations
func (k Keeper) IterateContractAuthorizations(ctx sdk.Context, cb func(types.ContractAuthorization) bool) {
	iterator, err := k.ContractAuthorizations.Iterate(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		contractAuth, err := iterator.Value()
		if err != nil {
			panic(err)
		}
		if cb(contractAuth) {
			return
		}
	}
}

// GetContractAuthorization gets any authorizations set up for the given contract address
func (k Keeper) GetContractAuthorization(ctx sdk.Context, contractAddr sdk.AccAddress) (types.ContractAuthorization, error) {
	return k.ContractAuthorizations.Get(ctx, contractAddr.Bytes())
}

// SetContractAuthorization creates of updates provided authorizations for given contract address
func (k Keeper) SetContractAuthorization(ctx sdk.Context, ca types.ContractAuthorization) error {
	if err := ca.Validate(); err != nil {
		return err
	}

	if !k.wasmKeeper.HasContractInfo(ctx, sdk.MustAccAddressFromBech32(ca.GetContractAddress())) {
		return types.ErrContractNotExist
	}

	return k.ContractAuthorizations.Set(ctx, sdk.MustAccAddressFromBech32(ca.ContractAddress), ca)
}

// DeleteContractAuthorization deletes any existing authorizations for given contract address
func (k Keeper) DeleteContractAuthorization(ctx sdk.Context, contractAddr sdk.AccAddress) error {
	return k.ContractAuthorizations.Remove(ctx, contractAddr.Bytes())
}

// HasContractAuthorization checks if the given contract address has any authorizations
func (k Keeper) HasContractAuthorization(ctx sdk.Context, contractAddr sdk.AccAddress) bool {
	has, err := k.ContractAuthorizations.Has(ctx, contractAddr.Bytes())
	if err != nil {
		panic(err)
	}
	return has
}
