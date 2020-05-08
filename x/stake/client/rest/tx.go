package rest

// func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
// 	r.HandleFunc(
// 		"/stake/{delegatorAddr}/delegations",
// 		postDelegationsHandlerFn(cliCtx),
// 	).Methods("POST")
// }

// // DelegateRequest defines the properties of a delegation request's body.
// type DelegateRequest struct {
// 	BaseReq          rest.BaseReq   `json:"base_req" yaml:"base_req"`
// 	DelegatorAddress sdk.AccAddress `json:"delegator_address" yaml:"delegator_address"` // in bech32
// 	ValidatorAddress sdk.ValAddress `json:"validator_address" yaml:"validator_address"` // in bech32
// 	Amount           sdk.Coin       `json:"amount" yaml:"amount"`
// }

// func postDelegationsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var req DelegateRequest

// 		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
// 			return
// 		}

// 		req.BaseReq = req.BaseReq.Sanitize()
// 		if !req.BaseReq.ValidateBasic(w) {
// 			return
// 		}

// 		// msg := types.NewMsgDelegate(req.DelegatorAddress, req.ValidatorAddress, req.Amount)
// 		// if err := msg.ValidateBasic(); err != nil {
// 		// 	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
// 		// 	return
// 		// }

// 		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
// 		if err != nil {
// 			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
// 			return
// 		}

// 		if !bytes.Equal(fromAddr, req.DelegatorAddress) {
// 			rest.WriteErrorResponse(w, http.StatusUnauthorized, "must use own delegator address")
// 			return
// 		}

// 		// utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
// 		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{})
// 	}
// }
