package types

import (
	"context"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// WasmKeeper defines the expected interface needed to setup and execute privilege contracts.
type WasmKeeper interface {
	GetCodeInfo(ctx context.Context, codeID uint64) *wasmtypes.CodeInfo
	// HasContractInfo checks if a contract with given address exists
	HasContractInfo(ctx context.Context, contractAddr sdk.AccAddress) bool
	GetContractInfo(ctx context.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
}
