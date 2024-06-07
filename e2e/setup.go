package e2e

import (
	"fmt"

	"github.com/strangelove-ventures/interchaintest/v8/ibc"
)

var (
	coinType = "118"
	denom    = "ustars"

	dockerImage = ibc.DockerImage{
		Repository: "publicawesome/stargaze",
		Version:    "local",
		UidGid:     "1025:1025",
	}

	stargazeCfg = ibc.ChainConfig{
		Type:                "cosmos",
		Name:                "stargaze-local",
		ChainID:             "stargaze-local-1",
		Images:              []ibc.DockerImage{dockerImage},
		Bin:                 "starsd",
		Bech32Prefix:        "stars",
		Denom:               denom,
		CoinType:            coinType,
		GasPrices:           fmt.Sprintf("0%s", denom),
		GasAdjustment:       2.0,
		TrustingPeriod:      "112h",
		NoHostMount:         false,
		SkipGenTx:           false,
		PreGenesis:          nil,
		ModifyGenesis:       nil,
		ConfigFileOverrides: nil,
	}
)
