package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	curatingtypes "github.com/public-awesome/stakebird/x/curating/types"
	"github.com/public-awesome/stakebird/x/stake/types"
)

var _ types.QueryServer = Keeper{}

// Stakes returns stakes for a post
func (k Keeper) Stakes(c context.Context, req *types.QueryStakesRequest) (*types.QueryStakesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	postID, err := curatingtypes.PostIDFromString(req.PostId)
	if err != nil {
		return nil, err
	}

	stakes := k.GetStakes(ctx, req.VendorId, postID)

	return &types.QueryStakesResponse{Stakes: stakes}, nil
}

// Stake returns a Stake
func (k Keeper) Stake(c context.Context, req *types.QueryStakeRequest) (*types.QueryStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	delegator, err := sdk.AccAddressFromBech32(req.Delegator)
	if err != nil {
		return nil, err
	}

	postID, err := curatingtypes.PostIDFromString(req.PostId)
	if err != nil {
		return nil, err
	}

	stake, found, err := k.GetStake(ctx, req.VendorId, postID, delegator)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("stake does not exist")
	}
	return &types.QueryStakeResponse{
		Stake: &stake,
	}, nil
}
