package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/public-awesome/stargaze/v11/x/cron/types"
)

var _ types.QueryServer = Keeper{}

// ListPrivileged lists the addresses of all the contracts which have been promoted to privilege status
func (k Keeper) ListPrivileged(c context.Context, req *types.QueryListPrivilegedRequest) (*types.QueryListPrivilegedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	var result types.QueryListPrivilegedResponse

	k.IteratePrivileged(ctx, func(addr sdk.AccAddress) bool {
		result.ContractAddresses = append(result.ContractAddresses, addr.String())
		return false
	})

	return &result, nil
}
