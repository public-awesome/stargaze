package e2e

import (
	"context"
	"testing"
	"time"

	interchaintest "github.com/strangelove-ventures/interchaintest/v4"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
	"github.com/strangelove-ventures/interchaintest/v4/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

const (
	haltHeightDelta    = uint64(10) // will propose upgrade this many blocks in the future
	blocksAfterUpgrade = uint64(10)
	votingPeriod       = "10s"
	maxDepositPeriod   = "10s"
)

const (
	initialVersion = "v11.0.0"
	upgradeVersion = "v12"
)

func TestChainUpgrade(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	numOfVals := 5
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:      "stargaze",
			ChainName: "stargaze-1",
			Version:   initialVersion,
			ChainConfig: ibc.ChainConfig{
				ModifyGenesis: cosmos.ModifyGenesis(getTestGenesis()),
			},
			NumValidators: &numOfVals,
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

	const userFunds = int64(10_000_000_000_000)
	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), userFunds, stargazeChain)
	chainUser := users[0]

	height, err := stargazeChain.Height(ctx)
	require.NoError(t, err, "error fetching height before submit upgrade proposal")

	haltHeight := height + haltHeightDelta

	proposal := cosmos.SoftwareUpgradeProposal{
		Deposit:     "10000000000" + stargazeChain.Config().Denom, // greater than min deposit
		Title:       "Chain Upgrade 1",
		Name:        upgradeVersion,
		Description: "First chain software upgrade",
		Height:      haltHeight,
	}

	upgradeTx, err := stargazeChain.UpgradeProposal(ctx, chainUser.KeyName(), proposal)
	require.NoError(t, err, "error submitting software upgrade proposal tx")

	err = stargazeChain.VoteOnProposalAllValidators(ctx, upgradeTx.ProposalID, cosmos.ProposalVoteYes)
	require.NoError(t, err, "failed to submit votes")

	_, err = cosmos.PollForProposalStatus(ctx, stargazeChain, height, height+haltHeightDelta, upgradeTx.ProposalID, cosmos.ProposalStatusPassed)
	require.NoError(t, err, "proposal status did not change to passed in expected number of blocks")

	height, err = stargazeChain.Height(ctx)
	require.NoError(t, err, "error fetching height before upgrade")

	timeoutCtx, timeoutCtxCancel := context.WithTimeout(ctx, time.Second*45)
	defer timeoutCtxCancel()

	// this should timeout due to chain halt at upgrade height.
	_ = testutil.WaitForBlocks(timeoutCtx, int(haltHeight-height)+1, stargazeChain)

	height, err = stargazeChain.Height(ctx)
	require.NoError(t, err, "error fetching height after chain should have halted")

	// make sure that chain is halted
	require.Equal(t, haltHeight, height, "height is not equal to halt height")

	// bring down nodes to prepare for upgrade
	err = stargazeChain.StopAllNodes(ctx)
	require.NoError(t, err, "error stopping node(s)")

	// upgrade version on all nodes
	stargazeChain.UpgradeVersion(ctx, client, "publicawesome/stargaze", "local-dev")

	// start all nodes back up.
	// validators reach consensus on first block after upgrade height
	// and chain block production resumes.
	err = stargazeChain.StartAllNodes(ctx)
	require.NoError(t, err, "error starting upgraded node(s)")

	timeoutCtx, timeoutCtxCancel = context.WithTimeout(ctx, time.Second*45)
	defer timeoutCtxCancel()

	err = testutil.WaitForBlocks(timeoutCtx, int(blocksAfterUpgrade), stargazeChain)
	require.NoError(t, err, "chain did not produce blocks after upgrade")
}

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
			Value: "ustars",
		},
	}
}
