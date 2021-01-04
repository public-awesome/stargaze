package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// InitGenesis initializes the curating module state
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	k.SetParams(ctx, state.Params)

	// NOTE: since the reward pool is a module account, the auth module should
	// take care of importing the amount into the account except for the
	// genesis block
	if k.GetRewardPoolBalance(ctx).IsZero() {
		err := k.InitializeRewardPool(ctx, k.GetParams(ctx).InitialRewardPool)
		if err != nil {
			panic(err)
		}
	}

	for _, post := range state.Posts {
		k.SetPost(ctx, post)
		if ctx.BlockTime().Before(post.CuratingEndTime) {
			k.InsertCurationQueue(ctx, post.VendorID, post.PostID, post.CuratingEndTime)
		}
	}

	for _, upvote := range state.Upvotes {
		curator, err := sdk.AccAddressFromBech32(upvote.Curator)
		if err != nil {
			panic(err)
		}
		k.SetUpvote(ctx, upvote, curator)
	}
}

// ExportGenesis exports the curating module state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	var posts types.Posts
	var upvotes types.Upvotes
	var curationQueue types.CuratingQueue

	// since 1 = twitter, for now
	vendorID := uint32(1)

	k.IteratePosts(ctx, vendorID, func(post types.Post) bool {
		posts = append(posts, post)

		k.IterateUpvotes(ctx, vendorID, post.PostID, func(upvote types.Upvote) bool {
			upvotes = append(upvotes, upvote)
			return false
		})

		if ctx.BlockTime().Before(post.CuratingEndTime) {
			vpPair := types.VPPair{VendorID: post.VendorID, PostID: post.PostID}
			curationQueue = append(curationQueue, vpPair)
		}
		return false
	})

	return &types.GenesisState{
		Params:        k.GetParams(ctx),
		Posts:         posts,
		Upvotes:       upvotes,
		CuratingQueue: curationQueue,
	}
}
