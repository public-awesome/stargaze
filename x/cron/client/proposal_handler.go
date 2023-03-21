package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/public-awesome/stargaze/v9/x/cron/client/cli"
	"github.com/public-awesome/stargaze/v9/x/cron/client/rest"
)

var (
	SetPrivilegeProposalHandler   = govclient.NewProposalHandler(cli.ProposalSetPrivilegeContractCmd, rest.ProposalSetPrivilegeContractHandler)
	UnsetPrivilegeProposalHandler = govclient.NewProposalHandler(cli.ProposalUnsetPrivilegeContractCmd, rest.ProposalDemotePrivilegeContractHandler)
)
