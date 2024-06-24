package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	tokenfactorytypes "github.com/public-awesome/stargaze/v15/x/tokenfactory/types"
	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
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

		// oracle
		"/slinky.oracle.v1.Query/GetAllCurrencyPairs": &oracletypes.GetAllCurrencyPairsResponse{},
		"/slinky.oracle.v1.Query/GetPrice":            &oracletypes.GetPriceResponse{},
		"/slinky.oracle.v1.Query/GetPrices":           &oracletypes.GetPricesResponse{},

		// marketmap
		"/slinky.marketmap.v1.Query/MarketMap":   &marketmaptypes.MarketMapResponse{},
		"/slinky.marketmap.v1.Query/LastUpdated": &marketmaptypes.LastUpdatedResponse{},
		"/slinky.marketmap.v1.Query/Params":      &marketmaptypes.ParamsResponse{},
	}
}

func GetWasmCapabilities() []string {
	return append(wasmkeeper.BuiltInCapabilities(), wasmCapabilities...)
}
