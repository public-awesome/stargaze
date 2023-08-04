package e2e

import (
	"fmt"

	"github.com/strangelove-ventures/interchaintest/v4/ibc"
)

var (
	coinType = "118"
	denom    = "ustars"

	dockerImage = ibc.DockerImage{
		Repository: "publicawesome/stargaze",
		Version:    "local-dev",
		UidGid:     "1025:1025",
	}

	stargazeCfg = ibc.ChainConfig{
		Type:                   "cosmos",
		Name:                   "local-stargaze",
		ChainID:                "testing",
		Images:                 []ibc.DockerImage{dockerImage},
		Bin:                    "starsd",
		Bech32Prefix:           "stars",
		Denom:                  denom,
		CoinType:               coinType,
		GasPrices:              fmt.Sprintf("0%s", denom),
		GasAdjustment:          2.0,
		TrustingPeriod:         "112h",
		NoHostMount:            false,
		SkipGenTx:              false,
		PreGenesis:             nil,
		ModifyGenesis:          nil,
		ConfigFileOverrides:    nil,
		UsingNewGenesisCommand: false,
	}
)
