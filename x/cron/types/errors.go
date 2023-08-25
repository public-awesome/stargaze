package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/cron module sentinel errors
var (
	ErrContractDoesNotExist    = errorsmod.Register(ModuleName, 2, "contract does not exist to modify its privilege")
	ErrContractPrivilegeNotSet = errorsmod.Register(ModuleName, 3, "contract does not have privilege set and therefore cannot unset its privilege")
)
