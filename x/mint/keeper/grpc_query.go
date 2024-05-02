package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v14/x/mint/types"
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

// Params returns params of the mint module.
func (q QueryServer) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params, err := q.keeper.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, err
}

// AnnualProvisions returns minter.AnnualProvisions of the mint module.
func (q QueryServer) AnnualProvisions(c context.Context, _ *types.QueryAnnualProvisionsRequest) (*types.QueryAnnualProvisionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	minter, err := q.keeper.GetMinter(ctx)

	return &types.QueryAnnualProvisionsResponse{AnnualProvisions: minter.AnnualProvisions}, err
}
