package e2e

import (
	"context"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"gopkg.in/yaml.v2"
)

type QueryMsg struct {
	GetCount *struct{} `json:"get_count"`
}

type QueryContractResponse struct {
	Data QueryContractResponseObj `json:"data"`
}

type QueryContractResponseObj struct {
	UpCount   int64 `json:"up_count"`
	DownCount int64 `json:"down_count"`
}

func InstantiateContract(chain *cosmos.CosmosChain, user ibc.Wallet, ctx context.Context, codeId string, initMsg string) (string, error) {
	// Instantiate the contract
	cmd := []string{
		chain.Config().Bin, "tx", "wasm", "instantiate", codeId, initMsg,
		"--label", "cwica-contract", "--admin", user.FormattedAddress(),
		"--from", user.KeyName(), "--keyring-backend", keyring.BackendTest,
		"--fees", "2000ustars",
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
		"--output", "json",
		"-y",
	}
	if _, _, err := chain.Exec(ctx, cmd, nil); err != nil {
		return "", err
	}
	if err := testutil.WaitForBlocks(ctx, 1, chain); err != nil {
		return "", err
	}

	// Getting the contract address
	cmd = []string{
		chain.Config().Bin, "q", "wasm", "list-contract-by-code", codeId,
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
	}
	stdout, _, err := chain.Exec(ctx, cmd, nil)
	if err != nil {
		return "", err
	}
	contactsRes := cosmos.QueryContractResponse{}
	if err = yaml.Unmarshal(stdout, &contactsRes); err != nil {
		return "", err
	}
	return contactsRes.Contracts[0], nil
}

func ExecuteContract(chain *cosmos.CosmosChain, user ibc.Wallet, ctx context.Context, contractAddress string, execMsg string) error {
	cmd := []string{
		chain.Config().Bin, "tx", "wasm", "execute", contractAddress, execMsg,
		"--from", user.KeyName(), "--keyring-backend", keyring.BackendTest,
		"--fees", "2000ustars",
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
		"--output", "json",
		"-y",
	}
	if _, _, err := chain.Exec(ctx, cmd, nil); err != nil {
		return err
	}
	return testutil.WaitForBlocks(ctx, 1, chain)
}
