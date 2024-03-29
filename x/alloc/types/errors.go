package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/alloc module sentinel errors
var (
	ErrUnauthorized = errorsmod.Register(ModuleName, 3, "sender is unauthorized to perform the operation")
)
