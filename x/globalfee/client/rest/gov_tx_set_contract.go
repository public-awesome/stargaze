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

func ProposalSetContractAuthorizationHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "globalfee_set_contract_authorization",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			var req SetContractAuthorizationProposalJSONReq
			if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
				return
			}
			toStdTxResponse(cliCtx, w, req)
		},
	}
}

type SetContractAuthorizationProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	ContractAddress string   `json:"contract_address" yaml:"contract_address"`
	Methods         []string `json:"methods" yaml:"methods"`
}

func (s SetContractAuthorizationProposalJSONReq) Content() govtypes.Content {
	return &types.SetContractAuthorizationProposal{
		Title:       s.Title,
		Description: s.Description,
		ContractAuthorization: &types.ContractAuthorization{
			ContractAddress: s.ContractAddress,
			Methods:         s.Methods,
		},
	}
}

func (s SetContractAuthorizationProposalJSONReq) GetProposer() string {
	return s.Proposer
}

func (s SetContractAuthorizationProposalJSONReq) GetDeposit() sdk.Coins {
	return s.Deposit
}

func (s SetContractAuthorizationProposalJSONReq) GetBaseReq() rest.BaseReq {
	return s.BaseReq
}
