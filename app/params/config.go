package params

import (
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	oracleconfig "github.com/skip-mev/slinky/oracle/config"
)

// CustomAppConfig defines the configuration for the Nois app.
type CustomAppConfig struct {
	serverconfig.Config
	Oracle oracleconfig.AppConfig `mapstructure:"oracle" json:"oracle"`
	Wasm   wasmtypes.WasmConfig   `mapstructure:"wasm"`
}

func CustomconfigTemplate(config wasmtypes.WasmConfig) string {
	return serverconfig.DefaultConfigTemplate + wasmtypes.ConfigTemplate(config) + oracleconfig.DefaultConfigTemplate
}

func DefaultConfig() (string, interface{}) {
	serverConfig := serverconfig.DefaultConfig()
	serverConfig.MinGasPrices = "0ustars"

	wasmConfig := wasmtypes.DefaultWasmConfig()
	simulationLimit := uint64(50_000_000)

	wasmConfig.SimulationGasLimit = &simulationLimit
	wasmConfig.SmartQueryGasLimit = 25_000_000
	wasmConfig.MemoryCacheSize = 512
	wasmConfig.ContractDebugMode = false

	oracleConfig := oracleconfig.AppConfig{
		Enabled:        true,
		OracleAddress:  "localhost:8080",
		ClientTimeout:  time.Second * 2,
		MetricsEnabled: true,
	}

	customConfig := CustomAppConfig{
		Config: *serverConfig,
		Oracle: oracleConfig,
		Wasm:   wasmConfig,
	}

	return CustomconfigTemplate(wasmConfig), customConfig
}
