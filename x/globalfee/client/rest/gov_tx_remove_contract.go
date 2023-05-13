package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/public-awesome/stargaze/v10/x/globalfee/types"
)

func ProposalRemoveContractAuthorizationHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "globalfee_remove_contract_authorization",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			var req RemoveContractAuthorizationProposalJSONReq
			if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
				return
			}
			toStdTxResponse(cliCtx, w, req)
		},
	}
}

type RemoveContractAuthorizationProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	ContractAddress string `json:"contract_address" yaml:"contract_address"`
}

func (s RemoveContractAuthorizationProposalJSONReq) Content() govtypes.Content {
	return &types.RemoveContractAuthorizationProposal{
		Title:           s.Title,
		Description:     s.Description,
		ContractAddress: s.ContractAddress,
	}
}

func (s RemoveContractAuthorizationProposalJSONReq) GetProposer() string {
	return s.Proposer
}

func (s RemoveContractAuthorizationProposalJSONReq) GetDeposit() sdk.Coins {
	return s.Deposit
}

func (s RemoveContractAuthorizationProposalJSONReq) GetBaseReq() rest.BaseReq {
	return s.BaseReq
}
