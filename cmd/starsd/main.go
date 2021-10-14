package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/public-awesome/stargaze/app"
	airdrop "github.com/public-awesome/stargaze/cmd/starsd/cmd"
	"github.com/tendermint/spm/cosmoscmd"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)
	rootCmd.AddCommand(airdrop.ExportHubSnapshotCmd())
	rootCmd.AddCommand(airdrop.ExportOsmosisSnapshotCmd())
	rootCmd.AddCommand(airdrop.ExportRegenSnapshotCmd())
	rootCmd.AddCommand(airdrop.ExportSnapshotCmd())
	rootCmd.AddCommand(airdrop.AddAirdropCmd())
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
