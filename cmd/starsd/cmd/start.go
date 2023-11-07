package cmd

import (
	"fmt"

	wasmvmapi "github.com/CosmWasm/wasmvm"
	"github.com/spf13/cobra"
)

var LibwasmVersion = "1.5.0"

func CheckLibwasmVersion(_ *cobra.Command, _ []string) error {
	version, err := wasmvmapi.LibwasmvmVersion()
	if err != nil {
		return fmt.Errorf("unable to retrieve libwasmversion %w", err)
	}
	if version != LibwasmVersion {
		return fmt.Errorf("libwasmversion mismatch. got: %s; expected: %s", version, LibwasmVersion)
	}
	return nil
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
