package curating

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// NewHandler creates an sdk.Handler for all the x/curating type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgPost:
			return handleMsgPost(ctx, k, msg)
		case types.MsgUpvote:
			return handleMsgUpvote(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgUpvote calls the keeper to perform the upvote operation
func handleMsgUpvote(ctx sdk.Context, k Keeper, msg types.MsgUpvote) (*sdk.Result, error) {
	// err := k.Delegate(ctx, msg.VendorID, msg.PostID, msg.DelegatorAddr, msg.ValidatorAddr, msg.Amount)
	// if err != nil {
	// 	return nil, err
	// }

	// ctx.EventManager().EmitEvents(sdk.Events{
	// 	sdk.NewEvent(
	// 		types.EventTypeDelegate,
	// 		sdk.NewAttribute(types.AttributeKeyVendorID, strconv.FormatUint(msg.VendorID, 10)),
	// 		sdk.NewAttribute(types.AttributeKeyPostID, strconv.FormatUint(msg.PostID, 10)),
	// 		sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
	// 	),
	// 	sdk.NewEvent(
	// 		sdk.EventTypeMessage,
	// 		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
	// 		sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddr.String()),
	// 	),
	// })

	// k.CreateUpvote(ctx, msg.VendorID, msg.PostID, msg.Curator, msg.RewardAccount, msg.VoteNum, msg.Deposit)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgPost(ctx sdk.Context, k Keeper, msg types.MsgPost) (*sdk.Result, error) {
	k.CreatePost(ctx, msg.VendorID, msg.PostID, msg.Body, msg.Deposit, msg.Creator, msg.RewardAccount)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypePost,
			sdk.NewAttribute(types.AttributeKeyVendorID, strconv.FormatUint(uint64(msg.VendorID), 10)),
			sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeKeyRewardAccount, msg.RewardAccount.String()),
			sdk.NewAttribute(types.AttributeKeyBody, msg.Body),
			sdk.NewAttribute(types.AttributeKeyDeposit, msg.Deposit.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
