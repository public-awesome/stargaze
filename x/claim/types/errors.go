package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/claim module sentinel errors
var (
	ErrIncorrectModuleAccountBalance = sdkerrors.Register(ModuleName, 2, "claim module account balance != sum of all claim record InitialClaimableAmounts")
)
