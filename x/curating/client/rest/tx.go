package rest

import (
	"bytes"
	"net/http"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/gorilla/mux"
	"github.com/public-awesome/stakebird/x/stake/types"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/stake/{delegatorAddr}/delegations",
		postDelegationsHandlerFn(cliCtx),
	).Methods("POST")

	r.HandleFunc(
		"/stake/{delegatorAddr}/posts",
		postPostsHandlerFn(cliCtx),
	).Methods("POST")
}

// DelegateRequest defines the properties of a delegation request's body.
type DelegateRequest struct {
	BaseReq          rest.BaseReq   `json:"base_req" yaml:"base_req"`
	DelegatorAddress sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"` // in bech32
	ValidatorAddress sdk.ValAddress `json:"validator_address" yaml:"validator_address"` // in bech32
	Amount           sdk.Coin       `json:"amount" yaml:"amount"`
	VendorID         uint64         `json:"vendor_id" yaml:"vendor_id"`
	PostID           uint64         `json:"post_id" yaml:"post_id"`
}

type PostRequest struct {
	*DelegateRequest

	Body         string        `json:"body" yaml:"body"`
	VotingPeriod time.Duration `json:"voting_period" yaml:"voting_period"`
}

func postDelegationsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if !bytes.Equal(fromAddr, req.DelegatorAddress) {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "must use own delegator address")
			return
		}

		msgDel := types.NewMsgDelegate(req.VendorID, req.PostID, req.DelegatorAddress, req.ValidatorAddress, req.Amount)
		msg := types.NewMsgPost(req.Body, msgDel, req.VotingPeriod)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

func postPostsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DelegateRequest

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if !bytes.Equal(fromAddr, req.DelegatorAddress) {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "must use own delegator address")
			return
		}

		msg := types.NewMsgDelegate(req.VendorID, req.PostID, req.DelegatorAddress, req.ValidatorAddress, req.Amount)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
