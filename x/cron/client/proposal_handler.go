package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/public-awesome/stargaze/v11/x/cron/client/cli"
)

var (
	SetPrivilegeProposalHandler   = govclient.NewProposalHandler(cli.ProposalSetPrivilegeContractCmd)
	UnsetPrivilegeProposalHandler = govclient.NewProposalHandler(cli.ProposalUnsetPrivilegeContractCmd)
)
