package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/public-awesome/stargaze/v8/x/cron/types"
)

func ProposalSetPrivilegeContractHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "wasm_set_privilege",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			var req SetPrivilegeProposalJSONReq
			if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
				return
			}
			toStdTxResponse(cliCtx, w, req)
		},
	}
}

func ProposalDemotePrivilegeContractHandler(cliCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "wasm_demote_privilege",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			var req SetPrivilegeProposalJSONReq
			if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
				return
			}
			toStdTxResponse(cliCtx, w, req)
		},
	}
}

type SetPrivilegeProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	Contract string `json:"contract" yaml:"contract"`
}

func (s SetPrivilegeProposalJSONReq) Content() govtypes.Content {
	return &types.PromoteToPrivilegedContractProposal{
		Title:       s.Title,
		Description: s.Description,
		Contract:    s.Contract,
	}
}

func (s SetPrivilegeProposalJSONReq) GetProposer() string {
	return s.Proposer
}

func (s SetPrivilegeProposalJSONReq) GetDeposit() sdk.Coins {
	return s.Deposit
}

func (s SetPrivilegeProposalJSONReq) GetBaseReq() rest.BaseReq {
	return s.BaseReq
}

type UnsetPrivilegeProposalJSONReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Proposer string    `json:"proposer" yaml:"proposer"`
	Deposit  sdk.Coins `json:"deposit" yaml:"deposit"`

	Contract string `json:"contract" yaml:"contract"`
}

func (s UnsetPrivilegeProposalJSONReq) Content() govtypes.Content {
	return &types.DemotePrivilegedContractProposal{
		Title:       s.Title,
		Description: s.Description,
		Contract:    s.Contract,
	}
}

func (s UnsetPrivilegeProposalJSONReq) GetProposer() string {
	return s.Proposer
}

func (s UnsetPrivilegeProposalJSONReq) GetDeposit() sdk.Coins {
	return s.Deposit
}

func (s UnsetPrivilegeProposalJSONReq) GetBaseReq() rest.BaseReq {
	return s.BaseReq
}

type wasmProposalData interface {
	Content() govtypes.Content
	GetProposer() string
	GetDeposit() sdk.Coins
	GetBaseReq() rest.BaseReq
}

func toStdTxResponse(cliCtx client.Context, w http.ResponseWriter, data wasmProposalData) {
	proposerAddr, err := sdk.AccAddressFromBech32(data.GetProposer())
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	msg, err := govtypes.NewMsgSubmitProposal(data.Content(), data.GetDeposit(), proposerAddr)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := msg.ValidateBasic(); err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	baseReq := data.GetBaseReq().Sanitize()
	if !baseReq.ValidateBasic(w) {
		return
	}
	tx.WriteGeneratedTxResponse(cliCtx, w, baseReq, msg)
}
