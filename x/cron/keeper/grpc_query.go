package keeper

import (
	"context"

	"github.com/public-awesome/stargaze/v8/x/cron/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ListPrivileged(c context.Context, req *types.QueryListPrivilegedRequest) (*types.QueryListPrivilegedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// todo

	return &types.QueryListPrivilegedResponse{}, nil
}
