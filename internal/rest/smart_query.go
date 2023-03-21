package rest

import (
	"encoding/json"
	"net/http"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/client"
)

// BatchedQuerierHandler returns a handler that performas batch queries to smart contracts.
func BatchedQuerierHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		batchRequest := &BatchQueryRequest{}
		err := json.NewDecoder(r.Body).Decode(batchRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if len(batchRequest.QueryRequests) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if len(batchRequest.QueryRequests) > 10 {
			w.WriteHeader(http.StatusBadRequest)
		}
		querier := wasmtypes.NewQueryClient(clientCtx)
		responses := make([]*BatchQueryResponse, len(batchRequest.QueryRequests))
		for i := range batchRequest.QueryRequests {
			resp, err := querier.SmartContractState(r.Context(), &batchRequest.QueryRequests[i])
			batchQueryResponse := &BatchQueryResponse{
				Index: i,
			}
			if err != nil {
				batchQueryResponse.Error = err.Error()
			}
			if resp != nil && resp.Data != nil {
				batchQueryResponse.Data = resp.Data
			}
			responses[i] = batchQueryResponse
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responses)
	}
}

type BatchQueryRequest struct {
	QueryRequests []wasmtypes.QuerySmartContractStateRequest `json:"query_requests"`
}

type BatchQueryResponse struct {
	Index int                          `json:"index"`
	Error string                       `json:"error"`
	Data  wasmtypes.RawContractMessage `json:"result"`
}
