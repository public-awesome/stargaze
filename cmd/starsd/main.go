package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/public-awesome/stargaze/v6/app"
	"github.com/public-awesome/stargaze/v6/cmd/starsd/cmd"
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
		cosmoscmd.AddCustomInitCmd(cmd.InitCmd(app.ModuleBasics, app.DefaultNodeHome)),
		cosmoscmd.AddSubCmd(cmd.PrepareGenesisCmd(app.DefaultNodeHome, app.ModuleBasics)),
		cosmoscmd.CustomizeStartCmd(cmd.CustomStart),
		cosmoscmd.AddSubCmd(cmd.AddGenesisWasmMsgCmd(app.DefaultNodeHome)),
		// this line is used by starport scaffolding # root/arguments
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
