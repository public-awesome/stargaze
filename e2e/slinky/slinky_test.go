package slinky_test

import (
	"encoding/json"
	"fmt"
	"github.com/icza/dyno"
	"strconv"
	"strings"
	"testing"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/skip-mev/slinky/tests/integration"
	marketmapmodule "github.com/skip-mev/slinky/x/marketmap"
	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"
	"github.com/skip-mev/slinky/x/oracle"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
	interchaintest "github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/stretchr/testify/suite"

	"github.com/public-awesome/stargaze/v14/app"
)

func init() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(app.Bech32PrefixAccAddr, app.Bech32PrefixAccPub)
	cfg.SetBech32PrefixForValidator(app.Bech32PrefixValAddr, app.Bech32PrefixValPub)
	cfg.SetBech32PrefixForConsensusNode(app.Bech32PrefixConsAddr, app.Bech32PrefixConsPub)
	cfg.SetAddressVerifier(wasmtypes.VerifyAddressLen())
	cfg.Seal()
}

var (
	dockerImage = ibc.DockerImage{
		Repository: "publicawesome/stargaze",
		Version:    "local-dev",
		UidGid:     "1025:1025",
	}

	numValidators = 4
	numFullNodes  = 0
	noHostMount   = false

	oracleImage = ibc.DockerImage{
		Repository: "ghcr.io/skip-mev/slinky-sidecar",
		Version:    "latest",
		UidGid:     "1000:1000",
	}
	encodingConfig = testutil.MakeTestEncodingConfig(
		bank.AppModuleBasic{},
		oracle.AppModuleBasic{},
		gov.AppModuleBasic{},
		auth.AppModuleBasic{},
		marketmapmodule.AppModuleBasic{},
	)

	defaultGenesis = marketmaptypes.DefaultGenesisState()
	params         = marketmaptypes.Params{
		MarketAuthorities: []string{"stars1salrqasc4de7qzdf0zkr9cpxg7gld4q28qrj7f"},
		Admin:             "stars1salrqasc4de7qzdf0zkr9cpxg7gld4q28qrj7f",
	}

	defaultGenesisKV = []cosmos.GenesisKV{
		{
			Key:   "consensus.params.abci.vote_extensions_enable_height",
			Value: "2",
		},
		{
			Key:   "consensus.params.block.max_gas",
			Value: "1000000000",
		},
		{
			Key:   "app_state.oracle",
			Value: oracletypes.DefaultGenesisState(),
		},
		{
			Key: "app_state.marketmap",
			Value: marketmaptypes.GenesisState{
				MarketMap:   defaultGenesis.MarketMap,
				LastUpdated: 0,
				Params:      params,
			},
		},
	}

	denom = "ustars"
	spec  = &interchaintest.ChainSpec{
		ChainName:     "slinky",
		Name:          "slinky",
		NumValidators: &numValidators,
		NumFullNodes:  &numFullNodes,
		NoHostMount:   &noHostMount,
		ChainConfig: ibc.ChainConfig{
			EncodingConfig: &encodingConfig,
			Images: []ibc.DockerImage{
				dockerImage,
			},
			Type:                "cosmos",
			Name:                "slinky",
			Denom:               denom,
			ChainID:             "chain-id-0",
			Bin:                 "starsd",
			Bech32Prefix:        "stars",
			CoinType:            "118",
			GasPrices:           fmt.Sprintf("1%s", denom),
			GasAdjustment:       2.0,
			TrustingPeriod:      "112h",
			NoHostMount:         false,
			SkipGenTx:           false,
			PreGenesis:          nil,
			ModifyGenesis:       ModifyGenesis(defaultGenesisKV),
			ConfigFileOverrides: nil,
		},
	}
)

func TestSlinkyOracleIntegration(t *testing.T) {
	baseSuite := integration.NewSlinkyIntegrationSuite(
		spec,
		oracleImage,
		integration.WithDenom(denom),
	)

	suite.Run(t, integration.NewSlinkyOracleIntegrationSuite(baseSuite))
}

func ModifyGenesis(genesisKV []cosmos.GenesisKV) func(ibc.ChainConfig, []byte) ([]byte, error) {
	return func(chainConfig ibc.ChainConfig, genbz []byte) ([]byte, error) {
		g := make(map[string]interface{})
		if err := json.Unmarshal(genbz, &g); err != nil {
			return nil, fmt.Errorf("failed to unmarshal genesis file: %w", err)
		}

		for idx, values := range genesisKV {
			splitPath := strings.Split(values.Key, ".")

			path := make([]interface{}, len(splitPath))
			for i, component := range splitPath {
				if v, err := strconv.Atoi(component); err == nil {
					path[i] = v
				} else {
					path[i] = component
				}
			}

			if err := dyno.Set(g, values.Value, path...); err != nil {
				return nil, fmt.Errorf("failed to set key '%s' as '%+v' (index:%d) in genesis json: %w", values.Key, values.Value, idx, err)
			}
		}

		out, err := json.Marshal(g)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal genesis bytes to json: %w", err)
		}

		// panic(string(out))

		return out, nil
	}
}
