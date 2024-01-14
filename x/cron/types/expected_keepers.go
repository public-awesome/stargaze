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
