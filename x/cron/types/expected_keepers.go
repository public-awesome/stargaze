package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// WasmKeeper defines the expected interface needed to setup and execute privilege contracts.
type WasmKeeper interface {
	// HasContractInfo checks if a contract with given address exists
	HasContractInfo(ctx context.Context, contractAddr sdk.AccAddress) bool
	// Sudo allows priviledged access to a contract
	Sudo(ctx context.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
}

// PauserKeeper defines the interface to check if a contract is paused.
type PauserKeeper interface {
	IsExecutionPaused(ctx sdk.Context, contractAddr sdk.AccAddress) bool
}
