package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/public-awesome/stargaze/v11/app"
	"github.com/public-awesome/stargaze/v11/cmd/starsd/cmd"
)

// func main() {
// 	rootCmd, _ := cosmoscmd.NewRootCmd(
// 		"stars",
// 		app.AccountAddressPrefix,
// 		app.DefaultNodeHome,
// 		app.Name,
// 		app.ModuleBasics,
// 		app.NewStargazeApp,
// 		cosmoscmd.AddSubCmd(cmd.TestnetCmd(app.ModuleBasics)),
// 		cosmoscmd.AddCustomInitCmd(cmd.InitCmd(app.ModuleBasics, app.DefaultNodeHome)),
// 		cosmoscmd.AddSubCmd(cmd.PrepareGenesisCmd(app.DefaultNodeHome, app.ModuleBasics)),
// 		cosmoscmd.CustomizeStartCmd(cmd.CustomStart),
// 		// this line is used by starport scaffolding # root/arguments
// 	)
// 	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
// 		os.Exit(1)
// 	}
// }

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
