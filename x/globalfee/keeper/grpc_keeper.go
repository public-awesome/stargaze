package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v14/x/globalfee/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (q QueryServer) CodeAuthorization(c context.Context, req *types.QueryCodeAuthorizationRequest) (*types.QueryCodeAuthorizationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	ca, found := q.keeper.GetCodeAuthorization(ctx, req.GetCodeId())
	if !found {
		return nil, status.Errorf(codes.NotFound, "code authorization: not found")
	}

	return &types.QueryCodeAuthorizationResponse{
		Methods: ca.GetMethods(),
	}, nil
}

func (q QueryServer) ContractAuthorization(c context.Context, req *types.QueryContractAuthorizationRequest) (*types.QueryContractAuthorizationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	contractAddr, err := sdk.AccAddressFromBech32(req.GetContractAddress())
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	ca, found := q.keeper.GetContractAuthorization(ctx, contractAddr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "contract authorization: not found")
	}

	return &types.QueryContractAuthorizationResponse{
		Methods: ca.GetMethods(),
	}, nil
}

func (q QueryServer) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := q.keeper.GetParams(ctx)
	return &types.QueryParamsResponse{Params: &params}, nil
}

func (q QueryServer) Authorizations(c context.Context, _ *types.QueryAuthorizationsRequest) (*types.QueryAuthorizationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	res := types.QueryAuthorizationsResponse{}
	q.keeper.IterateCodeAuthorizations(ctx, func(ca types.CodeAuthorization) bool {
		res.CodeAuthorizations = append(res.CodeAuthorizations, &ca)
		return false
	})
	q.keeper.IterateContractAuthorizations(ctx, func(ca types.ContractAuthorization) bool {
		res.ContractAuthorizations = append(res.ContractAuthorizations, &ca)
		return false
	})
	return &res, nil
}
