package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/x/curating/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the curating MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) Post(goCtx context.Context, msg *types.MsgPost) (*types.MsgPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	rewardAccount := sdk.AccAddress{}
	if strings.TrimSpace(msg.RewardAccount) != "" {
		rewardAccount, err = sdk.AccAddressFromBech32(msg.RewardAccount)
	}
	if err != nil {
		return nil, err
	}

	postID, err := types.PostIDFromString(msg.PostID)
	if err != nil {
		return nil, err
	}

	bodyHash, err := types.BodyHashFromString(msg.Body)
	if err != nil {
		return nil, err
	}

	post, err := k.CreatePost(
		ctx, msg.VendorID, &postID, bodyHash, msg.Body, creator, rewardAccount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})
	return &types.MsgPostResponse{
		PostID: post.PostID.String(),
	}, nil
}

func (k msgServer) Upvote(goCtx context.Context, msg *types.MsgUpvote) (*types.MsgUpvoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	curator, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}

	rewardAccount := sdk.AccAddress{}
	if strings.TrimSpace(msg.RewardAccount) != "" {
		rewardAccount, err = sdk.AccAddressFromBech32(msg.RewardAccount)
	}
	if err != nil {
		return nil, err
	}

	postID, err := types.PostIDFromString(msg.PostID)
	if err != nil {
		return nil, err
	}

	err = k.CreateUpvote(
		ctx, msg.VendorID, postID, curator, rewardAccount, msg.VoteNum)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Curator),
		),
	})
	return &types.MsgUpvoteResponse{}, nil
}
