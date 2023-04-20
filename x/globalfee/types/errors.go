package types

import sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	DefaultCodespace    = ModuleName
	ErrInvalidMethods   = sdkErrors.Register(DefaultCodespace, 2, "invalid method in code/contract authorization") // Code or Contract Authorizations have invalid methods configured
	ErrContractNotExist = sdkErrors.Register(DefaultCodespace, 3, "contract with given address does not exist")
	ErrCodeIdNotExist   = sdkErrors.Register(DefaultCodespace, 4, "code id does not exist")
)
