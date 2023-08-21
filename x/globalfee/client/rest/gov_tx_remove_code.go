package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/public-awesome/stargaze/v12/x/globalfee/types"
)

func ProposalRemoveCodeAuthorizationHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "globalfee_remove_code_authorization",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			var req RemoveCodeAuthorizationProposalJSONReq
			if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
				return
			}
			toStdTxResponse(cliCtx, w, req)
		},
	}
}

type RemoveCodeAuthorizationProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	CodeID uint64 `json:"code_id" yaml:"code_id"`
}

func (s RemoveCodeAuthorizationProposalJSONReq) Content() govtypes.Content {
	return &types.RemoveCodeAuthorizationProposal{
		Title:       s.Title,
		Description: s.Description,
		CodeID:      s.CodeID,
	}
}

func (s RemoveCodeAuthorizationProposalJSONReq) GetProposer() string {
	return s.Proposer
}

func (s RemoveCodeAuthorizationProposalJSONReq) GetDeposit() sdk.Coins {
	return s.Deposit
}

func (s RemoveCodeAuthorizationProposalJSONReq) GetBaseReq() rest.BaseReq {
	return s.BaseReq
}
