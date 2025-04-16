package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/gogoproto/proto"
	tokenfactorytypes "github.com/public-awesome/stargaze/v17/x/tokenfactory/types"
)

var wasmCapabilities = []string{
	"stargaze",
	"token_factory",
}

func AcceptedStargateQueries() wasmkeeper.AcceptedQueries {
	return wasmkeeper.AcceptedQueries{
		// tokenfactory
		"/osmosis.tokenfactory.v1beta1.Query/Params":                 func() proto.Message { return &tokenfactorytypes.QueryParamsResponse{} },
		"/osmosis.tokenfactory.v1beta1.Query/DenomAuthorityMetadata": func() proto.Message { return &tokenfactorytypes.QueryDenomAuthorityMetadataResponse{} },
		"/osmosis.tokenfactory.v1beta1.Query/DenomsFromCreator":      func() proto.Message { return &tokenfactorytypes.QueryDenomsFromCreatorResponse{} },
	}
}

func GetWasmCapabilities() []string {
	return append(wasmkeeper.BuiltInCapabilities(), wasmCapabilities...)
}

// initialize wasm overrides default 800kb max size for contract uploads
func initializeWasm() {
	wasmtypes.MaxWasmSize = 2_621_440 // 2.5 * 1024 * 1024
	wasmtypes.MaxProposalWasmSize = wasmtypes.MaxWasmSize
}
