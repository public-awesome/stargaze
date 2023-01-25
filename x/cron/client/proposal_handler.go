package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/public-awesome/stargaze/v8/x/cron/client/cli"
)

var (
	SetPrivilegeProposalHandler   = govclient.NewProposalHandler(cli.ProposalSetPrivilegeContractCmd, nil)   // todo rest
	UnsetPrivilegeProposalHandler = govclient.NewProposalHandler(cli.ProposalUnsetPrivilegeContractCmd, nil) // todo rest
)
