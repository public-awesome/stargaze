package app

import (
	"strings"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	tokenfactorytypes "github.com/public-awesome/stargaze/v12/x/tokenfactory/types"
)

var wasmCapabilities = []string{
	"iterator",
	"staking",
	"stargate",
	"stargaze",
	"cosmwasm_1_1",
	"cosmwasm_1_2",
	"cosmwasm_1_3",
	"token_factory",
}

func AcceptedStargateQueries() wasmkeeper.AcceptedStargateQueries {
	return wasmkeeper.AcceptedStargateQueries{
		"/osmosis.tokenfactory.v1beta1.Query/Params":                 &tokenfactorytypes.QueryParamsResponse{},
		"/osmosis.tokenfactory.v1beta1.Query/DenomAuthorityMetadata": &tokenfactorytypes.QueryDenomAuthorityMetadataResponse{},
		"/osmosis.tokenfactory.v1beta1.Query/DenomsFromCreator":      &tokenfactorytypes.QueryDenomsFromCreatorResponse{},
	}
}

func GetWasmCapabilities() string {
	return strings.Join(wasmCapabilities, ",")
}
