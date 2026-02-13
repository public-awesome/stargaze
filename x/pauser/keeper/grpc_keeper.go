package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v17/x/pauser/types"
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

func (q QueryServer) PausedContracts(c context.Context, _ *types.QueryPausedContractsRequest) (*types.QueryPausedContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var pausedContracts []types.PausedContract
	q.keeper.IteratePausedContracts(ctx, func(pc types.PausedContract) bool {
		pausedContracts = append(pausedContracts, pc)
		return false
	})

	return &types.QueryPausedContractsResponse{
		PausedContracts: pausedContracts,
	}, nil
}

func (q QueryServer) PausedCodeIDs(c context.Context, _ *types.QueryPausedCodeIDsRequest) (*types.QueryPausedCodeIDsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var pausedCodeIDs []types.PausedCodeID
	q.keeper.IteratePausedCodeIDs(ctx, func(pc types.PausedCodeID) bool {
		pausedCodeIDs = append(pausedCodeIDs, pc)
		return false
	})

	return &types.QueryPausedCodeIDsResponse{
		PausedCodeIds: pausedCodeIDs,
	}, nil
}

func (q QueryServer) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params, err := q.keeper.GetParams(ctx)
	return &types.QueryParamsResponse{Params: &params}, err
}
