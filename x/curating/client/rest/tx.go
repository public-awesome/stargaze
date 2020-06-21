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
		"/curating/{creatorAddr}/posts",
		postPostsHandlerFn(cliCtx),
	).Methods("POST")
}

type PostRequest struct {
	BaseReq  rest.BaseReq `json:"base_req" yaml:"base_req"`
	Stake    sdk.Coin     `json:"stake" yaml:"stake"`
	VendorID uint64       `json:"vendor_id" yaml:"vendor_id"`
	PostID   uint64       `json:"post_id" yaml:"post_id"`
	BodyHash string       `json:"body_hash" yaml:"body_hash"`
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

		msg := types.NewMsgPost(req.VendorID, req.PostID, fromAddr, req.BodyHash, req.Stake)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}
