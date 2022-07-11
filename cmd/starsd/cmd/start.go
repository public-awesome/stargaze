package cmd

import (
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmapi "github.com/CosmWasm/wasmvm/api"
	"github.com/spf13/cobra"
)

var LibwasmVersion = "1.0.0"

func CheckLibwasmVersion(cmd *cobra.Command, args []string) error {
	version, err := wasmvmapi.LibwasmvmVersion()
	if err != nil {
		return fmt.Errorf("unable to retrieve libwasmversion %w", err)
	}
	if version != LibwasmVersion {
		return fmt.Errorf("libwasmversion mismatch. got: %s; expected: %s", version, LibwasmVersion)
	}
	return nil
}

func CustomStart(startCmd *cobra.Command) {
	wasm.AddModuleInitFlags(startCmd)
	startCmd.PreRunE = chainPreRuns(CheckLibwasmVersion, startCmd.PreRunE)
}

type preRunFn func(cmd *cobra.Command, args []string) error

func chainPreRuns(pfns ...preRunFn) preRunFn {
	return func(cmd *cobra.Command, args []string) error {
		for _, pfn := range pfns {
			if pfn != nil {
				if err := pfn(cmd, args); err != nil {
					return err
				}
			}
		}
		return nil
	}
}
