package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/version"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/public-awesome/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

// AddAirdropAccounts Add balances of accounts to genesis, based on cosmos hub snapshot file
func AddAirdropAccounts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-airdrop-accounts [airdrop-snapshot-file] [denom] [community pool]",
		Short: "Add balances of accounts to genesis, based on cosmos hub snapshot file. The snapshot creation will be used to calc balances total, which will be added to the community pool you supply",
		Args:  cobra.ExactArgs(3),
		Long: fmt.Sprintf(`Add balances of accounts to genesis, based on cosmos hub snapshot file
Example:
$ %s add-airdrop-accounts /path/to/snapshot.json ujuno 2000
`, version.AppName),

		RunE: func(cmd *cobra.Command, args []string) error {

			var ctx = client.GetClientContextFromCmd(cmd)
			aminoCodec := ctx.LegacyAmino.Amino
			depCdc := ctx.JSONMarshaler
			cdc := depCdc.(codec.Marshaler)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(ctx.HomeDir)

			blob, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			snapshot := make(map[string]SnapshotFields)
			err = aminoCodec.UnmarshalJSON(blob, &snapshot)
			if err != nil {
				return err
			}

			denom := args[1]

			genFile := config.GenesisFile()

			// err, possibly redundant
			// genFileBlob, err := ioutil.ReadFile(genFile)
			// if err != nil {
			//	return err
			// }

			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			authGenState := authtypes.GetGenesisStateFromAppState(cdc, appState)

			accs, err := authtypes.UnpackAccounts(authGenState.Accounts)
			if err != nil {
				return fmt.Errorf("failed to get accounts from any: %w", err)
			}

			bankGenState := banktypes.GetGenesisStateFromAppState(depCdc, appState)

			fmt.Printf("Accounts %d\n", len(snapshot))

			// while this loop runs, collect balance to
			// populate supply later
			totalAtomSupplyFromBalances := sdk.NewInt(0)
			count := 0
			for address, acc := range snapshot {
				count++

				// Skip empty accounts
				if acc.JunoBalance.LTE(sdk.NewInt(0)) {
					continue
				}

				addr, err := ConvertCosmosAddressToJuno(address)
				if err != nil {
					return err
				}

				// Skip if account already exists
				if accs.Contains(addr) {
					continue
				}

				coin := sdk.NewCoin(denom, acc.JunoBalance)
				coins := sdk.NewCoins(coin)

				// add to total supply
				totalAtomSupplyFromBalances = totalAtomSupplyFromBalances.Add(acc.JunoBalance)

				// create concrete account type based on input parameters
				balances := banktypes.Balance{Address: addr.String(), Coins: coins.Sort()}
				genAccount := authtypes.NewBaseAccount(addr, nil, 0, 0)

				accs = append(accs, genAccount)

				bankGenState.Balances = append(bankGenState.Balances, balances)

				if count%1000 == 0 {
					fmt.Printf("Progress (%d of %d)\n", count, len(snapshot))
				}
			}

			fmt.Println("Done! Sorting...")

			accs = authtypes.SanitizeGenesisAccounts(accs)
			bankGenState.Balances = banktypes.SanitizeGenesisBalances(bankGenState.Balances)

			genAccs, err := authtypes.PackAccounts(accs)
			if err != nil {
				return fmt.Errorf("failed to convert accounts into any's: %w", err)
			}
			authGenState.Accounts = genAccs

			authGenStateBz, err := cdc.MarshalJSON(&authGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal auth genesis state: %w", err)
			}
			appState[authtypes.ModuleName] = authGenStateBz

			// let's set community pool in the genesis
			// it is nested in distribution
			var distGenesisState v038distribution.GenesisState
			aminoCodec.UnmarshalJSON(appState[v038distribution.ModuleName], &distGenesisState)
			distributionGenState := &distGenesisState

			cpAmountStr := args[2]
			cpAmount, ok := sdk.NewIntFromString(cpAmountStr)
			if !ok {
				return fmt.Errorf("failed to parse total supply arg: %s", cpAmountStr)
			}

			// needs to be v034distr.FeePool
			// so we need to Coin -> DecCoin
			cpCoin := sdk.NewCoin(denom, cpAmount)
			// cpCoins := sdk.NewCoins(cpCoin)
			communityPoolConfig := sdk.NewDecCoinsFromCoins(cpCoin)

			distributionGenState.FeePool.CommunityPool = communityPoolConfig
			distroGenStateJSON, err := aminoCodec.MarshalJSON(distributionGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal community pool genesis state: %w", err)
			}

			appState[v038distribution.ModuleName] = distroGenStateJSON

			// let's set supply before bank gen state is written
			totalAtomSupply := cpAmount.Add(totalAtomSupplyFromBalances)
			supplyCoin := sdk.NewCoin(denom, totalAtomSupply)
			supplyConfig := sdk.NewCoins(supplyCoin)

			bankGenState.Supply = supplyConfig

			bankGenStateBz, err := cdc.MarshalJSON(bankGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal bank genesis state: %w", err)
			}

			appState[banktypes.ModuleName] = bankGenStateBz

			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			genDoc.AppState = appStateJSON

			fmt.Printf("Saving genesis...")
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	return cmd
}
