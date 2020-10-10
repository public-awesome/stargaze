package curating

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stakebird/x/curating/keeper"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// NewHandler creates an sdk.Handler for all the x/curating type messages
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgPost:
			return handleMsgPost(ctx, k, msg)
		case *types.MsgUpvote:
			return handleMsgUpvote(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgPost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgPost) (*sdk.Result, error) {
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
	err = k.CreatePost(
		ctx, msg.VendorID, msg.PostID, msg.Body, creator, rewardAccount)
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

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

// handleMsgUpvote calls the keeper to perform the upvote operation
func handleMsgUpvote(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpvote) (*sdk.Result, error) {
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

	err = k.CreateUpvote(
		ctx, msg.VendorID, msg.PostID, curator, rewardAccount, msg.VoteNum)
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

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
