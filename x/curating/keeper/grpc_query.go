package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

var _ types.QueryServer = Keeper{}

// Params returns module params
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{
		Params: k.GetParams(ctx),
	}, nil
}

// Posts returns all posts based on vendor
func (k Keeper) Posts(c context.Context, req *types.QueryPostsRequest) (*types.QueryPostsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	posts := k.GetPosts(ctx, req.VendorId)

	return &types.QueryPostsResponse{Posts: posts}, nil
}

// Post returns a post based on vendor and post id
func (k Keeper) Post(c context.Context, req *types.QueryPostRequest) (*types.QueryPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	postID, err := types.PostIDFromString(req.PostId)
	if err != nil {
		return nil, err
	}

	post, found, err := k.GetPost(ctx, req.VendorId, postID)
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
	var upvotes []types.Upvote
	fmt.Println("HELLO................")

	postID, err := types.PostIDFromString(req.PostId)
	if err != nil {
		return nil, err
	}

	post, found, err := k.GetPost(ctx, req.VendorId, postID)
	if err != nil || !found {
		return nil, types.ErrPostNotFound
	}
	fmt.Println("HELLO................")
	k.IterateUpvotes(ctx, req.VendorId, post.PostID, func(upvote types.Upvote) (stop bool) {
		fmt.Println("in here...")
		upvotes = append(upvotes, upvote)
		return false
	})
	return &types.QueryUpvotesResponse{
		Upvotes: upvotes,
	}, nil
}
