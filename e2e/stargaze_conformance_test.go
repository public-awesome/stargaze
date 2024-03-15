package e2e

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
)

func TestStargazeConformance(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	t.Parallel()
	stargazeChain, _, ctx := startChain(t)
	chainUser1 := fundChainUser(t, ctx, "user1", stargazeChain)

	// TESTING MODULE : Wasmd
	wasmdConformance(ctx, stargazeChain, chainUser1, t)

	// TESTING MODULE : Tokenfactory - from user
	tokenFactoryConformance(t, ctx, stargazeChain, chainUser1)
}

func wasmdConformance(ctx context.Context, stargazeChain *cosmos.CosmosChain, chainUser1 ibc.Wallet, t *testing.T) {
	codeId, err := stargazeChain.StoreContract(ctx, chainUser1.KeyName(), "artifacts/cron_counter.wasm")
	require.NoError(t, err)

	initMsg := `{}`
	contractAddress, err := InstantiateContract(stargazeChain, chainUser1, ctx, codeId, initMsg)
	require.NoError(t, err)

	var queryRes QueryContractResponse
	err = stargazeChain.QueryContract(ctx, contractAddress, QueryMsg{GetCount: &struct{}{}}, &queryRes)
	require.NoError(t, err)
	require.Equal(t, int64(0), queryRes.Data.UpCount)

	execMsg := `{"increment":{}}`
	err = ExecuteContract(stargazeChain, chainUser1, ctx, contractAddress, execMsg)
	require.NoError(t, err)

	err = stargazeChain.QueryContract(ctx, contractAddress, QueryMsg{GetCount: &struct{}{}}, &queryRes)
	require.NoError(t, err)
	require.Equal(t, int64(1), queryRes.Data.UpCount)
}

func tokenFactoryConformance(t *testing.T, ctx context.Context, stargazeChain *cosmos.CosmosChain, chainUser1 ibc.Wallet) {
	subDenomName := "ictest"

	registerDenom := []string{
		stargazeChain.Config().Bin, "tx", "tokenfactory", "create-denom", subDenomName,
		"--from", chainUser1.KeyName(),
		"--chain-id", stargazeChain.Config().ChainID,
		"--home", stargazeChain.HomeDir(),
		"--node", stargazeChain.GetRPCAddress(),
		"--fees", "2000" + stargazeChain.Config().Denom,
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}
	_, _, err := stargazeChain.Exec(ctx, registerDenom, nil)
	require.NoError(t, err)
	err = testutil.WaitForBlocks(ctx, 1, stargazeChain)
	require.NoError(t, err)

	queryDenomsByUser := []string{
		stargazeChain.Config().Bin, "query", "tokenfactory", "denoms-from-creator", chainUser1.FormattedAddress(),
		"--chain-id", stargazeChain.Config().ChainID,
		"--home", stargazeChain.HomeDir(),
		"--node", stargazeChain.GetRPCAddress(),
		"--output", "json",
	}
	stdout, _, err := stargazeChain.Exec(ctx, queryDenomsByUser, nil)
	require.NoError(t, err)
	var tfdenoms QueryTokenFactoryDenomsFromCreatorResponse
	err = json.Unmarshal(stdout, &tfdenoms)
	require.NoError(t, err)
	require.Len(t, tfdenoms.Denoms, 1)
	tfDenom := tfdenoms.Denoms[0]

	mintToken := []string{
		stargazeChain.Config().Bin, "tx", "tokenfactory", "mint", "1234" + tfDenom,
		"--from", chainUser1.KeyName(),
		"--chain-id", stargazeChain.Config().ChainID,
		"--home", stargazeChain.HomeDir(),
		"--node", stargazeChain.GetRPCAddress(),
		"--fees", "2000" + stargazeChain.Config().Denom,
		"--keyring-backend", keyring.BackendTest,
		"--output", "json",
		"-y",
	}
	_, _, err = stargazeChain.Exec(ctx, mintToken, nil)
	require.NoError(t, err)
	err = testutil.WaitForBlocks(ctx, 2, stargazeChain)
	require.NoError(t, err)

	balance, err := stargazeChain.GetBalance(ctx, chainUser1.FormattedAddress(), tfDenom)
	require.NoError(t, err)
	require.Equal(t, "1234", balance.String())

	burnToken := []string{
		stargazeChain.Config().Bin, "tx", "tokenfactory", "burn", "234" + tfDenom,
		"--from", chainUser1.KeyName(),
		"--chain-id", stargazeChain.Config().ChainID,
		"--home", stargazeChain.HomeDir(),
		"--node", stargazeChain.GetRPCAddress(),
		"--fees", "2000" + stargazeChain.Config().Denom,
		"--keyring-backend", keyring.BackendTest,
		"--output", "json",
		"-y",
	}
	_, _, err = stargazeChain.Exec(ctx, burnToken, nil)
	require.NoError(t, err)
	err = testutil.WaitForBlocks(ctx, 2, stargazeChain)
	require.NoError(t, err)

	balance, err = stargazeChain.GetBalance(ctx, chainUser1.FormattedAddress(), tfDenom)
	require.NoError(t, err)
	require.Equal(t, "1000", balance.String())
}

type QueryTokenFactoryDenomsFromCreatorResponse struct {
	Denoms []string `json:"denoms"`
}
