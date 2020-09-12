package rest

import (
	"net/http"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/public-awesome/stakebird/x/faucet/internal/types"
)

// PostProposalReq defines the properties of a proposal request's body.
type PostMintReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Minter  string       `json:"minter" yaml:"minter"` // Address of the minter
	Denom   string       `json:"denom" yaml:"denom"`   // Denom of the token
}

func mintHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostMintReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		sender, err := sdk.AccAddressFromBech32(baseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		minter, err := sdk.AccAddressFromBech32(req.Minter)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		denom := req.Denom

		// create the message
		msg := types.NewMsgMint(sender, minter, time.Now().Unix(), denom)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		authclient.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
