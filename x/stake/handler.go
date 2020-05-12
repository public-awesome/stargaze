package stake

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rocket-protocol/stakebird/x/stake/types"
)

// NewHandler creates an sdk.Handler for all the x/stake type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgPost:
			return handleMsgPost(ctx, k, msg)
		case types.MsgDelegate:
			return handleMsgDelegate(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgDelegate calls the keeper to perform the Delegate operation
func handleMsgDelegate(ctx sdk.Context, k Keeper, msg types.MsgDelegate) (*sdk.Result, error) {
	err := k.Delegate(ctx, msg.VendorID, msg.PostID, msg.Downvote, msg.DelegatorAddr, msg.ValidatorAddr, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDelegate,
			sdk.NewAttribute(types.AttributeKeyVendorID, strconv.FormatUint(msg.VendorID, 10)),
			sdk.NewAttribute(types.AttributeKeyPostID, strconv.FormatUint(msg.PostID, 10)),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddr.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgPost(ctx sdk.Context, k Keeper, msg types.MsgPost) (*sdk.Result, error) {
	k.CreatePost(ctx, msg.PostID, msg.VendorID, msg.Body, msg.VotingPeriod)
	err := k.Delegate(ctx, msg.VendorID, msg.PostID, msg.DelegatorAddr, msg.ValidatorAddr, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypePost,
			sdk.NewAttribute(types.AttributeKeyVendorID, strconv.FormatUint(msg.VendorID, 10)),
			sdk.NewAttribute(types.AttributeKeyPostID, strconv.FormatUint(msg.PostID, 10)),
			sdk.NewAttribute(types.AttributeKeyBody, msg.Body),
			sdk.NewAttribute(types.AttributeKeyVotingPeriod, msg.VotingPeriod.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddr.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
