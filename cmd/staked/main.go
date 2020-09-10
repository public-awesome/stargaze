package main

import (
	"os"

	"github.com/public-awesome/stakebird/cmd/staked/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
