package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/curating module sentinel errors
var (
	ErrStaketNotFound = sdkerrors.Register(ModuleName, 1, "Stake not found")
	ErrAmountTooLarge = sdkerrors.Register(ModuleName, 1, "Unstake amount too large")
)
