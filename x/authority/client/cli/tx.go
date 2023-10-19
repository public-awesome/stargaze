package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v13/x/authority/types"
	"github.com/spf13/cobra"
)

// GetTxCmd builds tx command group for the module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdSubmitProposal())
	return cmd
}

func CmdSubmitProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proposal [path/to/proposal.json]",
		Short: "Submit a proposal along with some messages",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msgs, err := parseSubmitProposal(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}

			msg, err := types.NewMsgExecuteProposal(msgs, clientCtx.GetFromAddress().String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// proposal defines the new Msg-based proposal.
type proposal struct {
	// Msgs defines an array of sdk.Msgs proto-JSON-encoded as Anys.
	Messages []json.RawMessage `json:"messages,omitempty"`
}

// parseSubmitProposal reads and parses the proposal.
func parseSubmitProposal(cdc codec.Codec, path string) ([]sdk.Msg, error) {
	var proposal proposal

	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &proposal)
	if err != nil {
		return nil, err
	}

	msgs := make([]sdk.Msg, len(proposal.Messages))
	for i, anyJSON := range proposal.Messages {
		var msg sdk.Msg
		err := cdc.UnmarshalInterfaceJSON(anyJSON, &msg)
		if err != nil {
			return nil, err
		}

		msgs[i] = msg
	}

	return msgs, nil
}
