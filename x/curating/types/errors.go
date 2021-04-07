package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/curating module sentinel errors
var (
	ErrPostNotFound  = sdkerrors.Register(ModuleName, 2, "Post not found")
	ErrAlreadyVoted  = sdkerrors.Register(ModuleName, 3, "Already voted")
	ErrDuplicatePost = sdkerrors.Register(ModuleName, 4, "Post already exists")
	ErrPostExpired   = sdkerrors.Register(ModuleName, 5, "Post already expired")
	ErrInvalidPostID = sdkerrors.Register(ModuleName, 6, "PostID cannot be nil")
)
