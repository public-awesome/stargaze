package params

import (
	"fmt"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	oracleconfig "github.com/skip-mev/slinky/oracle/config"
)

// CustomAppConfig defines the configuration for the Nois app.
type CustomAppConfig struct {
	serverconfig.Config
	Wasm   wasmtypes.WasmConfig   `mapstructure:"wasm"`
	Oracle oracleconfig.AppConfig `mapstructure:"oracle"`
}

// ConfigTemplate toml snippet for app.toml
func OracleConfigTemplate(c oracleconfig.AppConfig) string {

	return fmt.Sprintf(`
###############################################################################
###                                  Oracle                                 ###
###############################################################################
[oracle]
# Enabled indicates whether the oracle is enabled.
enabled = %t

# Oracle Address is the URL of the out of process oracle sidecar. This is used to
# connect to the oracle sidecar when the application boots up. Note that the address
# can be modified at any point, but will only take effect after the application is
# restarted. This can be the address of an oracle container running on the same
# machine or a remote machine.
oracle_address = "%s"

# Client Timeout is the time that the client is willing to wait for responses from 
# the oracle before timing out.
client_timeout = "%s"

# MetricsEnabled determines whether oracle metrics are enabled. Specifically
# this enables instrumentation of the oracle client and the interaction between
# the oracle and the app.
metrics_enabled = %t
`, c.Enabled, c.OracleAddress, c.ClientTimeout.String(), c.MetricsEnabled)
}
func CustomconfigTemplate(config wasmtypes.WasmConfig, oracleConfig oracleconfig.AppConfig) string {
	return serverconfig.DefaultConfigTemplate + wasmtypes.ConfigTemplate(config) + OracleConfigTemplate(oracleConfig)
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
	customConfig := CustomAppConfig{
		Config: *serverConfig,
		Wasm:   wasmConfig,
		Oracle: oracleconfig.AppConfig{
			Enabled:        false,
			OracleAddress:  "localhost:8080",
			ClientTimeout:  time.Second * 1,
			MetricsEnabled: false,
		},
	}

	return CustomconfigTemplate(wasmConfig, customConfig.Oracle), customConfig
}
