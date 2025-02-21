package params

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

// CustomAppConfig defines the configuration for the Nois app.
type CustomAppConfig struct {
	serverconfig.Config
	Wasm wasmtypes.NodeConfig `mapstructure:"wasm" json:"wasm"`
}

func CustomconfigTemplate(config wasmtypes.NodeConfig) string {
	return serverconfig.DefaultConfigTemplate + wasmtypes.ConfigTemplate(config)
}

func DefaultConfig() (string, interface{}) {
	serverConfig := serverconfig.DefaultConfig()
	serverConfig.MinGasPrices = "0ustars"

	wasmConfig := wasmtypes.DefaultNodeConfig()
	simulationLimit := uint64(50_000_000)

	wasmConfig.SimulationGasLimit = &simulationLimit
	wasmConfig.SmartQueryGasLimit = 25_000_000
	wasmConfig.MemoryCacheSize = 1024
	wasmConfig.ContractDebugMode = false

	customConfig := CustomAppConfig{
		Config: *serverConfig,
		Wasm:   wasmConfig,
	}

	return CustomconfigTemplate(wasmConfig), customConfig
}
