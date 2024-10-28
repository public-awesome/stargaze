package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	tokenfactorytypes "github.com/public-awesome/stargaze/v15/x/tokenfactory/types"
)

var wasmCapabilities = []string{
	"stargaze",
	"token_factory",
}

func AcceptedStargateQueries() wasmkeeper.AcceptedQueries {
	return wasmkeeper.AcceptedQueries{
		// tokenfactory
		"/osmosis.tokenfactory.v1beta1.Query/Params":                 &tokenfactorytypes.QueryParamsResponse{},
		"/osmosis.tokenfactory.v1beta1.Query/DenomAuthorityMetadata": &tokenfactorytypes.QueryDenomAuthorityMetadataResponse{},
		"/osmosis.tokenfactory.v1beta1.Query/DenomsFromCreator":      &tokenfactorytypes.QueryDenomsFromCreatorResponse{},
	}
}

func GetWasmCapabilities() []string {
	return append(wasmkeeper.BuiltInCapabilities(), wasmCapabilities...)
}
