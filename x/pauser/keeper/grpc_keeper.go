package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/public-awesome/stargaze/v18/x/pauser/types"
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

func (q QueryServer) IsContractPaused(c context.Context, req *types.QueryIsContractPausedRequest) (*types.QueryIsContractPausedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	contractAddr, err := sdk.AccAddressFromBech32(req.GetContractAddress())
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	pc, err := q.keeper.GetPausedContract(ctx, contractAddr)
	if err != nil {
		return &types.QueryIsContractPausedResponse{Paused: false}, nil
	}

	return &types.QueryIsContractPausedResponse{
		Paused:         true,
		PausedContract: &pc,
	}, nil
}

func (q QueryServer) IsCodeIDPaused(c context.Context, req *types.QueryIsCodeIDPausedRequest) (*types.QueryIsCodeIDPausedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	pc, err := q.keeper.GetPausedCodeID(ctx, req.GetCodeId())
	if err != nil {
		return &types.QueryIsCodeIDPausedResponse{Paused: false}, nil
	}

	return &types.QueryIsCodeIDPausedResponse{
		Paused:       true,
		PausedCodeId: &pc,
	}, nil
}

func (q QueryServer) PausedContracts(c context.Context, req *types.QueryPausedContractsRequest) (*types.QueryPausedContractsResponse, error) {
	results, pageRes, err := query.CollectionPaginate(
		c,
		q.keeper.PausedContracts,
		req.GetPagination(),
		func(_ []byte, value types.PausedContract) (types.PausedContract, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.QueryPausedContractsResponse{
		PausedContracts: results,
		Pagination:      pageRes,
	}, nil
}

func (q QueryServer) PausedCodeIDs(c context.Context, req *types.QueryPausedCodeIDsRequest) (*types.QueryPausedCodeIDsResponse, error) {
	results, pageRes, err := query.CollectionPaginate(
		c,
		q.keeper.PausedCodeIDs,
		req.GetPagination(),
		func(_ uint64, value types.PausedCodeID) (types.PausedCodeID, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.QueryPausedCodeIDsResponse{
		PausedCodeIds: results,
		Pagination:    pageRes,
	}, nil
}

func (q QueryServer) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params, err := q.keeper.GetParams(ctx)
	return &types.QueryParamsResponse{Params: &params}, err
}
