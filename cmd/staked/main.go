package main

import (
	"os"

	"github.com/public-awesome/stakebird/app"
	"github.com/public-awesome/stakebird/cmd/staked/cmd"
)

func main() {
	app.ConfigureAccountPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
