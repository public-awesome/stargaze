package keeper

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/public-awesome/stargaze/v15/x/cron/types"
)

// SetPrivileged checks if the given contract exists and adds it to the list of privilege contracts
func (k Keeper) SetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) error {
	if k.wasmKeeper.HasContractInfo(ctx, contractAddr) {
		if !k.IsPrivileged(ctx, contractAddr) {
			if err := k.PrivilegedContracts.Set(ctx, contractAddr.Bytes(), []byte{1}); err != nil {
				return err
			}
		}
		event := sdk.NewEvent(
			types.EventTypeSetContractPriviledge,
			sdk.NewAttribute(wasmtypes.AttributeKeyContractAddr, contractAddr.String()),
		)
		ctx.EventManager().EmitEvent(event)
	} else {
		return types.ErrContractDoesNotExist
	}
	return nil
}

// UnsetPrivileged checks if the given contract exists and if it has privilege and remove it from the list of privileg contracts
func (k Keeper) UnsetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) error {
	if k.wasmKeeper.HasContractInfo(ctx, contractAddr) {
		if k.IsPrivileged(ctx, contractAddr) {
			if err := k.PrivilegedContracts.Remove(ctx, contractAddr.Bytes()); err != nil {
				return err
			}

			event := sdk.NewEvent(
				types.EventTypeUnsetContractPriviledge,
				sdk.NewAttribute(wasmtypes.AttributeKeyContractAddr, contractAddr.String()),
			)
			ctx.EventManager().EmitEvent(event)
		} else {
			return types.ErrContractPrivilegeNotSet
		}
	} else {
		return types.ErrContractDoesNotExist
	}
	return nil
}

// IsPrivileged returns if the given contract is part of the privilege contract list
func (k Keeper) IsPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) bool {
	has, err := k.PrivilegedContracts.Has(ctx, contractAddr.Bytes())
	if err != nil {
		return false
	}
	return has
}

// IteratePrivileged executes the given func on all the privilege contracts
func (k Keeper) IteratePrivileged(ctx sdk.Context, doSomething func(sdk.AccAddress) bool) {
	iterator, err := k.PrivilegedContracts.Iterate(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		key, err := iterator.Key()
		if err != nil {
			panic(err)
		}
		contractAddr := sdk.AccAddress(key)
		if doSomething(contractAddr) {
			return
		}
	}
}
