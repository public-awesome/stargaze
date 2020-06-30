package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/gorilla/mux"
	"github.com/public-awesome/stakebird/x/curating/types"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/curating/{creatorAddr}/posts", postPostsHandlerFn(cliCtx),
	).Methods("POST")

	r.HandleFunc(
		"/curating/{curatorAddr}/upvotes", postUpvoteHandlerFn(cliCtx),
	).Methods("POST")
}

// PostRequest is the REST API request to register a post
type PostRequest struct {
	BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
	RewardAccount string       `json:"reward_account,omitempty" yaml:"reward_account"`
	Deposit       sdk.Coin     `json:"deposit" yaml:"deposit"`
	VendorID      uint32       `json:"vendor_id" yaml:"vendor_id"`
	PostID        string       `json:"post_id" yaml:"post_id"`
	Body          string       `json:"body" yaml:"body"`
}

func postPostsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
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

		var rewardAddr sdk.AccAddress
		if req.RewardAccount != "" {
			rewardAddr, err = sdk.AccAddressFromBech32(req.RewardAccount)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		msg := types.NewMsgPost(req.VendorID, req.PostID, fromAddr, rewardAddr, req.Body, req.Deposit)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

// UpvoteRequest is the REST API request to perform an upvote
type UpvoteRequest struct {
	BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
	RewardAccount string       `json:"reward_account,omitempty" yaml:"reward_account"`
	Deposit       sdk.Coin     `json:"deposit" yaml:"deposit"`
	VendorID      uint32       `json:"vendor_id" yaml:"vendor_id"`
	PostID        string       `json:"post_id" yaml:"post_id"`
	VoteNum       int32        `json:"vote_num" yaml:"vote_num"`
}

func postUpvoteHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpvoteRequest

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

		var rewardAddr sdk.AccAddress
		if req.RewardAccount != "" {
			rewardAddr, err = sdk.AccAddressFromBech32(req.RewardAccount)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		} else {
			rewardAddr = fromAddr
		}

		msg := types.NewMsgUpvote(
			req.VendorID, req.PostID, fromAddr, rewardAddr, req.VoteNum, req.Deposit)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
