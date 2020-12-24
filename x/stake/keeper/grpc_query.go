package keeper

import (
	"context"
	"fmt"

	"github.com/bwmarrin/snowflake"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/stake/types"
)

var _ types.QueryServer = Keeper{}

// Stakes returns stakes for a post
func (k Keeper) Stakes(c context.Context, req *types.QueryStakesRequest) (*types.QueryStakesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	k.Logger(ctx).Info("in here")
	fmt.Println("in here yo")

	postID, err := postIDBytes(req.PostId)
	if err != nil {
		return nil, err
	}
	stakes := k.GetStakes(ctx, req.VendorId, postID)
	// fmt.Println(req.VendorId)
	// fmt.Println(postID)
	fmt.Println(stakes)

	return &types.QueryStakesResponse{Stakes: stakes}, nil
}

// Stake returns a Stake
func (k Keeper) Stake(c context.Context, req *types.QueryStakeRequest) (*types.QueryStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	postID, err := postIDBytes(req.PostId)
	if err != nil {
		return nil, err
	}

	delegator, err := sdk.AccAddressFromBech32(req.Delegator)
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

// postIDBytes returns the byte representation of a postID int64
func postIDBytes(postID string) ([]byte, error) {
	pID, err := snowflake.ParseString(postID)
	if err != nil {
		return nil, err
	}

	temp := pID.IntBytes()

	return temp[:], nil
}
