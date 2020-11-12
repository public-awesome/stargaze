package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/curating module sentinel errors
var (
	ErrPostNotFound  = sdkerrors.Register(ModuleName, 1, "Post not found")
	ErrAlreadyVoted  = sdkerrors.Register(ModuleName, 2, "Already voted")
	ErrDuplicatePost = sdkerrors.Register(ModuleName, 3, "Post already exists")
	ErrPostExpired   = sdkerrors.Register(ModuleName, 4, "Post already expired")
)
