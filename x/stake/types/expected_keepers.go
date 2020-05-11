package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// stakingexported "github.com/cosmos/cosmos-sdk/x/staking/exported"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

/*
When a module wishes to interact with another module, it is good practice to define what it will use
as an interface so the module cannot use things that are not permitted.
*/

type StakingKeeper interface {
	// TODO: figure out why the exported validator interface doesn't compile [shanev]
	// Delegate(ctx sdk.Context, delAddr sdk.AccAddress, bondAmt sdk.Int, tokenSrc sdk.BondStatus,
	// 	validator stakingexported.ValidatorI, subtractAccount bool) (newShares sdk.Dec, err error)

	Delegate(ctx sdk.Context, delAddr sdk.AccAddress, bondAmt sdk.Int, tokenSrc sdk.BondStatus,
		validator stakingtypes.Validator, subtractAccount bool) (newShares sdk.Dec, err error)

	// GetValidator(ctx sdk.Context, valAddress sdk.ValAddress) (validator stakingexported.ValidatorI, found bool)

	GetValidator(ctx sdk.Context, valAddress sdk.ValAddress) (validator stakingtypes.Validator, found bool)

	Unbond(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, shares sdk.Dec) (amount sdk.Int, err error)
}
