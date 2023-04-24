package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/public-awesome/stargaze/v9/x/globalfee/client/cli"
)

var (
	SetCodeAuthorizationProposalHandler        = govclient.NewProposalHandler(cli.CmdProposalSetCodeAuthorization, nil)
	RemoveCodeAuthorizationProposalHandler     = govclient.NewProposalHandler(cli.CmdProposalRemoveCodeAuthorization, nil)
	SetContractAuthorizationProposalHandler    = govclient.NewProposalHandler(cli.CmdProposalSetContractAuthorization, nil)
	RemoveContractAuthorizationProposalHandler = govclient.NewProposalHandler(cli.CmdProposalRemoveContractAuthorization, nil)
)
