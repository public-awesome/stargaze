package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

var _ types.QueryServer = Keeper{}

// Post returns a post based on vendor and post id
func (k Keeper) Post(c context.Context, req *types.QueryPostRequest) (*types.QueryPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	post, found, err := k.GetPost(ctx, req.VendorId, req.PostId)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("post does not exist")
	}
	return &types.QueryPostResponse{
		Post: &post,
	}, nil
}

// Upvotes returns all upvotes for a given vendor and post id
func (k Keeper) Upvotes(c context.Context, req *types.QueryUpvotesRequest) (*types.QueryUpvotesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var upvotes []*types.Upvote
	post, found, err := k.GetPost(ctx, req.VendorId, req.PostId)
	if err != nil || !found {
		return nil, types.ErrPostNotFound
	}
	k.IterateUpvotes(ctx, req.VendorId, post.PostIDHash, func(upvote types.Upvote) (stop bool) {
		upvotes = append(upvotes, &upvote)
		return false
	})
	return &types.QueryUpvotesResponse{
		Upvotes: upvotes,
	}, nil
}
