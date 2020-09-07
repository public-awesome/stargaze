package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// You can see how they are constructed below:
var (
	ErrPostNotFound = sdkerrors.Register(ModuleName, 1, "Post not found")
	ErrAlreadyVoted = sdkerrors.Register(ModuleName, 2, "Already voted")
)
