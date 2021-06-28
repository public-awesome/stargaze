package client

import (
	"github.com/cosmos/cosmos-sdk/x/distribution/client/rest"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/public-awesome/stargaze/x/ibc-spend/client/cli"
)

// ProposalHandler is the IBC community spend proposal handler.
var (
	// TODO: add rest
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
