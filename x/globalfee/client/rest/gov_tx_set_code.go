package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/public-awesome/stargaze/v11/x/globalfee/types"
)

func ProposalSetCodeAuthorizationHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "globalfee_set_code_authorization",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			var req SetCodeAuthorizationProposalJSONReq
			if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
				return
			}
			toStdTxResponse(cliCtx, w, req)
		},
	}
}

type SetCodeAuthorizationProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	CodeID  uint64   `json:"code_id" yaml:"code_id"`
	Methods []string `json:"methods" yaml:"methods"`
}

func (s SetCodeAuthorizationProposalJSONReq) Content() govtypes.Content {
	return &types.SetCodeAuthorizationProposal{
		Title:       s.Title,
		Description: s.Description,
		CodeAuthorization: &types.CodeAuthorization{
			CodeID:  s.CodeID,
			Methods: s.Methods,
		},
	}
}

func (s SetCodeAuthorizationProposalJSONReq) GetProposer() string {
	return s.Proposer
}

func (s SetCodeAuthorizationProposalJSONReq) GetDeposit() sdk.Coins {
	return s.Deposit
}

func (s SetCodeAuthorizationProposalJSONReq) GetBaseReq() rest.BaseReq {
	return s.BaseReq
}
