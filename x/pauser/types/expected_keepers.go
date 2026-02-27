package types

import (
	"context"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// WasmKeeper defines the expected interface needed to check contract and code info.
type WasmKeeper interface {
	GetCodeInfo(ctx context.Context, codeID uint64) *wasmtypes.CodeInfo
	HasContractInfo(ctx context.Context, contractAddr sdk.AccAddress) bool
	GetContractInfo(ctx context.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
}
