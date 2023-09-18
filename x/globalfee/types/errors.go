package types

import errorsmod "cosmossdk.io/errors"

var (
	DefaultCodespace    = ModuleName
	ErrInvalidMethods   = errorsmod.Register(DefaultCodespace, 2, "invalid method in code/contract authorization") // Code or Contract Authorizations have invalid methods configured
	ErrContractNotExist = errorsmod.Register(DefaultCodespace, 3, "contract with given address does not exist")
	ErrCodeIDNotExist   = errorsmod.Register(DefaultCodespace, 4, "code id does not exist")
	ErrUnauthorized     = errorsmod.Register(DefaultCodespace, 5, "sender is unauthorized to perform the operation")
)
