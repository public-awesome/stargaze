package types

import sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	DefaultCodespace = ModuleName
	ErrInternal      = sdkErrors.Register(DefaultCodespace, 0, "internal error") // smth went wrong
)
