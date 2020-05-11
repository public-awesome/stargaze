package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// ParamSubspace defines the expected Subspace interfacace
// type ParamSubspace interface {
// 	WithKeyTable(table params.KeyTable) params.Subspace
// 	Get(ctx sdk.Context, key []byte, ptr interface{})
// 	GetParamSet(ctx sdk.Context, ps params.ParamSet)
// 	SetParamSet(ctx sdk.Context, ps params.ParamSet)
// }

/*
When a module wishes to interact with another module, it is good practice to define what it will use
as an interface so the module cannot use things that are not permitted.
type BankKeeper interface {
	SubtractCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, error)
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}
*/

type StakingKeeper interface {
	// Delegate(ctx sdk.Context, delAddr sdk.AccAddress, bondAmt sdk.Int, tokenSrc sdk.BondStatus,
	// 	validator stakingexported.ValidatorI, subtractAccount bool) (newShares sdk.Dec, err error)

	Delegate(ctx sdk.Context, delAddr sdk.AccAddress, bondAmt sdk.Int, tokenSrc sdk.BondStatus,
		validator stakingtypes.Validator, subtractAccount bool) (newShares sdk.Dec, err error)

	// GetValidator(ctx sdk.Context, valAddress sdk.ValAddress) (validator stakingexported.ValidatorI, found bool)

	GetValidator(ctx sdk.Context, valAddress sdk.ValAddress) (validator stakingtypes.Validator, found bool)

	Unbond(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, shares sdk.Dec) (amount sdk.Int, err error)
}
