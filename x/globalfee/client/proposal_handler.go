package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/public-awesome/stargaze/v10/x/globalfee/client/cli"
	"github.com/public-awesome/stargaze/v10/x/globalfee/client/rest"
)

var (
	SetCodeAuthorizationProposalHandler        = govclient.NewProposalHandler(cli.CmdProposalSetCodeAuthorization, rest.ProposalSetCodeAuthorizationHandler)
	RemoveCodeAuthorizationProposalHandler     = govclient.NewProposalHandler(cli.CmdProposalRemoveCodeAuthorization, rest.ProposalRemoveCodeAuthorizationHandler)
	SetContractAuthorizationProposalHandler    = govclient.NewProposalHandler(cli.CmdProposalSetContractAuthorization, rest.ProposalSetContractAuthorizationHandler)
	RemoveContractAuthorizationProposalHandler = govclient.NewProposalHandler(cli.CmdProposalRemoveContractAuthorization, rest.ProposalRemoveContractAuthorizationHandler)
)
