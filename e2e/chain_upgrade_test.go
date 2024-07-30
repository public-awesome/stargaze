package e2e

import (
	"context"
	"strconv"
	"testing"
	"time"

	"cosmossdk.io/math"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/docker/docker/client"
	interchaintest "github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

const (
	initialVersion = "v14.0.0-rc2" // The last released git tag version of the stargaze binary. This version of the image is fetched from Heigliner backage repository.
	upgradeName    = "v15"         // The upcoming version name - Should match with upgrade handler name. This version needs to be built locally for tests. Using `make build-docker`
)

const (
	haltHeightDelta    = int64(20) // The number of blocks after which to apply upgrade after creation of proposal.
	blocksAfterUpgrade = int64(10) // The number of blocks to wait for after the upgrade has been applied.
	votingPeriod       = "30s"     // Reducing voting period for testing
	maxDepositPeriod   = "10s"     // Reducing max deposit period for testing
	depositDenom       = "ustars"  // The bond denom to be used to deposit for propsals
)

func TestChainUpgrade(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	stargazeChain, client, ctx := startChain(t)
	chainUser := fundChainUser(t, ctx, t.Name(), stargazeChain)

	// Creating a contract before upgrade and ensuring expected state
	codeId, err := stargazeChain.StoreContract(ctx, chainUser.KeyName(), "artifacts/cron_counter.wasm")
	require.NoError(t, err)
	initMsg := `{}`
	contractAddress, err := InstantiateContract(stargazeChain, chainUser, ctx, codeId, initMsg)
	require.NoError(t, err)
	var queryRes QueryContractResponse
	err = stargazeChain.QueryContract(ctx, contractAddress, QueryMsg{GetCount: &struct{}{}}, &queryRes)
	require.NoError(t, err)
	require.Equal(t, int64(0), queryRes.Data.UpCount)

	haltHeight := submitUpgradeProposalAndVote(t, ctx, stargazeChain, chainUser)

	height, err := stargazeChain.Height(ctx)
	require.NoError(t, err, "error fetching height before upgrade")

	timeoutCtx, timeoutCtxCancel := context.WithTimeout(ctx, time.Second*45)
	defer timeoutCtxCancel()

	// This should timeout due to chain halt at upgrade height.
	_ = testutil.WaitForBlocks(timeoutCtx, int(haltHeight)-int(height)+1, stargazeChain)

	height, err = stargazeChain.Height(ctx)
	require.NoError(t, err, "error fetching height after chain should have halted")

	// Make sure that chain is halted
	require.Equal(t, haltHeight, height, "height is not equal to halt height")

	// Bring down nodes to prepare for upgrade
	err = stargazeChain.StopAllNodes(ctx)
	require.NoError(t, err, "error stopping node(s)")

	// Upgrade version on all nodes - We are passing in the local image for the upgrade build using `make build-docker`
	stargazeChain.UpgradeVersion(ctx, client, "publicawesome/stargaze", "local-dev")

	// Start all nodes back up.
	// Validators reach consensus on first block after upgrade height
	// And chain block production resumes 🎉
	err = stargazeChain.StartAllNodes(ctx)
	require.NoError(t, err, "error starting upgraded node(s)")

	timeoutCtx, timeoutCtxCancel = context.WithTimeout(ctx, time.Second*45)
	defer timeoutCtxCancel()

	err = testutil.WaitForBlocks(timeoutCtx, int(blocksAfterUpgrade), stargazeChain)
	require.NoError(t, err, "chain did not produce blocks after upgrade")

	// Ensure contract behavior is as expected after upgrade
	execMsg := `{"increment":{}}`
	err = ExecuteContract(stargazeChain, chainUser, ctx, contractAddress, execMsg)
	require.NoError(t, err)
	err = stargazeChain.QueryContract(ctx, contractAddress, QueryMsg{GetCount: &struct{}{}}, &queryRes)
	require.NoError(t, err)
	require.Equal(t, int64(1), queryRes.Data.UpCount)
}

func submitUpgradeProposalAndVote(t *testing.T, ctx context.Context, stargazeChain *cosmos.CosmosChain, chainUser ibc.Wallet) int64 {
	height, err := stargazeChain.Height(ctx) // The current chain height
	require.NoError(t, err, "error fetching height before submit upgrade proposal")

	haltHeight := height + haltHeightDelta // The height at which upgrade should be applied

	govAuthorityAddr, err := stargazeChain.GetGovernanceAddress(ctx)
	require.NoError(t, err, "error fetching governance address")
	proposalMsg := upgradetypes.MsgSoftwareUpgrade{
		Authority: govAuthorityAddr,
		Plan: upgradetypes.Plan{
			Name:   upgradeName,
			Height: int64(haltHeight),
		},
	}
	proposal, err := stargazeChain.BuildProposal([]cosmos.ProtoMessage{&proposalMsg},
		"Test Upgrade",
		"Every PR we perform an upgrade check to ensure nothing breaks",
		"metadata",
		"10000000000"+stargazeChain.Config().Denom,
		chainUser.KeyName(),
		false,
	)
	require.NoError(t, err, "error building proposal tx")

	upgradeTx, err := stargazeChain.SubmitProposal(ctx, chainUser.KeyName(), proposal) // Submitting the software upgrade proposal
	require.NoError(t, err, "error submitting software upgrade proposal tx")

	proposalID, err := strconv.ParseUint(upgradeTx.ProposalID, 10, 64)
	require.NoError(t, err, "error parsing proposal ID")

	err = testutil.WaitForBlocks(ctx, 2, stargazeChain)
	require.NoError(t, err, "error waiting for blocks after proposal submission")

	// Vote on the proposal
	for _, n := range stargazeChain.Nodes() {
		if n.Validator {
			n := n
			_, err = n.ExecTx(ctx, "validator",
				"gov", "vote", upgradeTx.ProposalID, cosmos.ProposalVoteYes,
				"--gas", "auto", "--gas-adjustment", "2.0",
			)
			require.NoError(t, err, "failed to submit votes")
		}
	}

	_, err = cosmos.PollForProposalStatusV1(ctx, stargazeChain, height, height+haltHeightDelta, proposalID, govv1.ProposalStatus_PROPOSAL_STATUS_PASSED)
	require.NoError(t, err, "proposal status did not change to passed in expected number of blocks")
	return haltHeight
}

func fundChainUser(t *testing.T, ctx context.Context, userName string, stargazeChain *cosmos.CosmosChain) ibc.Wallet {
	userFunds := math.NewInt(10_000_000_000_000)
	users := interchaintest.GetAndFundTestUsers(t, ctx, userName, userFunds, stargazeChain)
	return users[0]
}

func startChain(t *testing.T) (*cosmos.CosmosChain, *client.Client, context.Context) {
	// Configuring the chain factory. We are building Stargaze chain with the version that matches the `initialVersion` value
	numOfVals := 5
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:      "stargaze",
			ChainName: "stargaze-1",
			Version:   initialVersion,
			ChainConfig: ibc.ChainConfig{
				ModifyGenesis: cosmos.ModifyGenesis(getTestGenesis()), // Modifying genesis to have test-friendly gov params
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
	return stargazeChain, client, ctx
}

func getTestGenesis() []cosmos.GenesisKV {
	return []cosmos.GenesisKV{
		{
			Key:   "app_state.gov.params.voting_period",
			Value: votingPeriod,
		},
	}
}
