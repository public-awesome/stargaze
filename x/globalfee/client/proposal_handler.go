package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/public-awesome/stargaze/v12/x/globalfee/client/cli"
)

var (
	SetCodeAuthorizationProposalHandler        = govclient.NewProposalHandler(cli.CmdProposalSetCodeAuthorization)
	RemoveCodeAuthorizationProposalHandler     = govclient.NewProposalHandler(cli.CmdProposalRemoveCodeAuthorization)
	SetContractAuthorizationProposalHandler    = govclient.NewProposalHandler(cli.CmdProposalSetContractAuthorization)
	RemoveContractAuthorizationProposalHandler = govclient.NewProposalHandler(cli.CmdProposalRemoveContractAuthorization)
)
