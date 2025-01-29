package cmd

import (
	"time"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
)

const flagSkipPreferredSettings = "skip-preferred-settings"

type PreferredSetting struct {
	ViperKey string
	Value    string
	Set      func(cmd *cobra.Command, serverCtx *server.Context, key, value string) error
}

var preferredSettings = []PreferredSetting{
	{
		ViperKey: "consensus.timeout_commit",
		Value:    "3s",
		Set: func(cmd *cobra.Command, serverCtx *server.Context, key, value string) error {
			serverCtx.Viper.Set(key, value)
			serverCtx.Config.Consensus.TimeoutCommit = time.Duration(3 * time.Second)
			return nil
		},
	},
	{
		ViperKey: "consensus.timeout_propose",
		Value:    "2s",
		Set: func(cmd *cobra.Command, serverCtx *server.Context, key, value string) error {
			serverCtx.Viper.Set(key, value)
			serverCtx.Config.Consensus.TimeoutPropose = time.Duration(2 * time.Second)
			return nil
		}},
	{
		ViperKey: "wasm.memory_cache_size",
		Value:    "1024",
		Set: func(cmd *cobra.Command, serverCtx *server.Context, key, value string) error {
			serverCtx.Viper.Set(key, value)
			return nil
		},
	},
}

func SetPreferredSettings(cmd *cobra.Command, _ []string) error {

	skip, err := cmd.Flags().GetBool(flagSkipPreferredSettings)
	if err != nil {
		return err
	}
	if skip {
		return nil
	}

	serverCtx := server.GetServerContextFromCmd(cmd)

	for _, setting := range preferredSettings {
		err := setting.Set(cmd, serverCtx, setting.ViperKey, setting.Value)
		if err != nil {
			return err
		}
	}

	return server.SetCmdServerContext(cmd, serverCtx)
}

func LogPreferredSettings(cmd *cobra.Command, _ []string) error {
	serverCtx := server.GetServerContextFromCmd(cmd)

	skip, err := cmd.Flags().GetBool(flagSkipPreferredSettings)
	if err != nil {
		return err
	}

	if !skip {
		serverCtx.Logger.Info("using preferred settings use --skip-preferred-settings to disable")
	}

	serverCtx.Logger.Info("using timeout_commit", "value", serverCtx.Config.Consensus.TimeoutCommit.String())
	serverCtx.Logger.Info("using timeout_propose", "value", serverCtx.Config.Consensus.TimeoutPropose.String())
	serverCtx.Logger.Info("using wasm.memory_cache_size", "value", serverCtx.Viper.Get("wasm.memory_cache_size"))

	return nil
}
