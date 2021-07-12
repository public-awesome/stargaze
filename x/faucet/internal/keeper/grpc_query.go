package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/x/faucet/internal/types"
)

var _ types.QueryServer = Keeper{}

// FaucetKey returns the stored faucet key
func (k Keeper) FaucetKey(c context.Context, req *types.QueryFaucetKeyRequest) (*types.QueryFaucetKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if !k.HasFaucetKey(ctx) {
		return nil, types.ErrKeyNotFound
	}
	return &types.QueryFaucetKeyResponse{FaucetKey: k.GetFaucetKey(ctx)}, nil
}
