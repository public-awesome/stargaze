package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cron module sentinel errors
var (
	ErrContractDoesNotExist    = sdkerrors.Register(ModuleName, 2, "contract does not exist to modify its privilege")
	ErrContractPrivilegeNotSet = sdkerrors.Register(ModuleName, 3, "contract does not have privilege set and therefore cannot unset its privilege")
)
