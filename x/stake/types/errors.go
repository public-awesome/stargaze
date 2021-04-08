package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/curating module sentinel errors
var (
	// error code 1 is reserved for internal ABCI error code
	ErrStaketNotFound     = sdkerrors.Register(ModuleName, 2, "stake not found")
	ErrAmountTooLarge     = sdkerrors.Register(ModuleName, 3, "unstake amount too large")
	ErrCurationNotExpired = sdkerrors.Register(ModuleName, 4, "post is still being curated")
)
