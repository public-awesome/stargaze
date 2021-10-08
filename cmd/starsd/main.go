package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/public-awesome/stargaze/app"
	"github.com/public-awesome/stargaze/cmd/starsd/cmd"
	airdrop "github.com/public-awesome/stargaze/cmd/starsd/cmd"
	"github.com/tendermint/spm/cosmoscmd"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		"stars",
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.NewStargazeApp,
		cosmoscmd.AddSubCmd(cmd.TestnetCmd(app.ModuleBasics)),
		// this line is used by starport scaffolding # root/arguments
	)
	rootCmd.AddCommand(airdrop.ExportHubSnapshotCmd())
	rootCmd.AddCommand(airdrop.ExportOsmosisSnapshotCmd())
	rootCmd.AddCommand(airdrop.ExportRegenSnapshotCmd())
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
