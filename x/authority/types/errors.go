package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/authority module sentinel errors
var (
	ErrUnauthorized          = errorsmod.Register(ModuleName, 2, "sender is unauthorized to perform the operation")
	ErrAuthorizationNotFound = errorsmod.Register(ModuleName, 3, "")
)
