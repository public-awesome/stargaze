package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/alloc module sentinel errors
var (
	ErrUnauthorized = errorsmod.Register(ModuleName, 2, "sender is unauthorized to perform the operation")
)
