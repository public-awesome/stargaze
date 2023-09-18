package e2e

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/client"
	"github.com/strangelove-ventures/interchaintest/v4"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
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
		Name:                   "stargaze-local",
		ChainID:                "stargaze-local-1",
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
		ModifyGenesis:          cosmos.ModifyGenesis(getTestGenesis()), // Modifying genesis to have test-friendly gov params
		ConfigFileOverrides:    nil,
		UsingNewGenesisCommand: false,
	}
)

func startChain(t *testing.T, version string) (*cosmos.CosmosChain, *client.Client, context.Context) {
	// Configuring the chain factory. We are building Stargaze chain with the version that matches the `initialVersion` value
	numOfVals := 5
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          "stargaze",
			ChainName:     "stargaze-1",
			Version:       version,
			ChainConfig:   stargazeCfg,
			NumValidators: &numOfVals,
			NumFullNodes:  &numOfVals,
		},
	})
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	stargazeChain := chains[0].(*cosmos.CosmosChain)

	ic := interchaintest.NewInterchain().AddChain(stargazeChain)
	client, network := interchaintest.DockerSetup(t)
	ctx := context.Background()
	require.NoError(t, ic.Build(ctx, nil, interchaintest.InterchainBuildOptions{
		TestName:         t.Name(),
		Client:           client,
		NetworkID:        network,
		SkipPathCreation: true,
	}))
	t.Cleanup(func() {
		_ = ic.Close()
	})
	return stargazeChain, client, ctx
}

const (
	votingPeriod     = "10s"    // Reducing voting period for testing
	maxDepositPeriod = "10s"    // Reducing max deposit period for testing
	depositDenom     = "ustars" // The bond denom to be used to deposit for propsals
)

func getTestGenesis() []cosmos.GenesisKV {
	return []cosmos.GenesisKV{
		{
			Key:   "app_state.gov.voting_params.voting_period",
			Value: votingPeriod,
		},
		{
			Key:   "app_state.gov.deposit_params.max_deposit_period",
			Value: maxDepositPeriod,
		},
		{
			Key:   "app_state.gov.deposit_params.min_deposit.0.denom",
			Value: depositDenom,
		},
	}
}
