package e2e

import (
	"fmt"

	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
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
		ModifyGenesis:          cosmos.ModifyGenesis(getTestGenesis()),
		ConfigFileOverrides:    nil,
		UsingNewGenesisCommand: false,
	}
)

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
		{
			Key:   "app_state.interchainaccounts.host_genesis_state.params.allow_messages",
			Value: []string{"/cosmos.bank.v1beta1.MsgSend", "/cosmos.staking.v1beta1.MsgDelegate"},
		},
	}
}
