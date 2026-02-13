package types

import errorsmod "cosmossdk.io/errors"

var (
	DefaultCodespace    = ModuleName
	ErrContractPaused   = errorsmod.Register(DefaultCodespace, 2, "contract is paused")
	ErrCodeIDPaused     = errorsmod.Register(DefaultCodespace, 3, "code ID is paused")
	ErrContractNotExist = errorsmod.Register(DefaultCodespace, 4, "contract with given address does not exist")
	ErrCodeIDNotExist   = errorsmod.Register(DefaultCodespace, 5, "code id does not exist")
	ErrUnauthorized     = errorsmod.Register(DefaultCodespace, 6, "sender is unauthorized to perform the operation")
	ErrAlreadyPaused    = errorsmod.Register(DefaultCodespace, 7, "already paused")
	ErrNotPaused        = errorsmod.Register(DefaultCodespace, 8, "not paused")
	ErrNestedMsgTooDeep = errorsmod.Register(DefaultCodespace, 9, "nested message depth exceeds maximum")
)
