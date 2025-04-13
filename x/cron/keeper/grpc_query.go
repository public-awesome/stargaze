package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/public-awesome/stargaze/v16/x/cron/types"
)

var _ types.QueryServer = &QueryServer{}

// QueryServer implements the module gRPC query service.
type QueryServer struct {
	keeper Keeper
}

// NewQueryServer creates a new gRPC query server.
func NewQueryServer(keeper Keeper) *QueryServer {
	return &QueryServer{
		keeper: keeper,
	}
}

// ListPrivileged lists the addresses of all the contracts which have been promoted to privilege status
func (q QueryServer) ListPrivileged(c context.Context, req *types.QueryListPrivilegedRequest) (*types.QueryListPrivilegedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	var result types.QueryListPrivilegedResponse

	q.keeper.IteratePrivileged(ctx, func(addr sdk.AccAddress) bool {
		result.ContractAddresses = append(result.ContractAddresses, addr.String())
		return false
	})

	return &result, nil
}

// Params fetches all the params of x/cron module
func (q QueryServer) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	params, err := q.keeper.GetParams(ctx)

	return &types.QueryParamsResponse{
		Params: params,
	}, err
}
