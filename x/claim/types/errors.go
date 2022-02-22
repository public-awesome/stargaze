package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/claim module sentinel errors
var (
	ErrAirdropNotEnabled             = sdkerrors.Register(ModuleName, 2, "airdrop not enabled")
	ErrIncorrectModuleAccountBalance = sdkerrors.Register(ModuleName, 3, "claim module account balance != sum of all claim record InitialClaimableAmounts")
	ErrUnauthorizedClaimer           = sdkerrors.Register(ModuleName, 4, "address is not allowed to claim")
)
