package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// WasmKeeper defines the expected interface needed to setup and execute privilege contracts.
type WasmKeeper interface {
	// HasContractInfo checks if a contract with given address exists
	HasContractInfo(ctx sdk.Context, contractAddr sdk.AccAddress) bool
	// Sudo allows priviledged access to a contract
	Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
}
