package user

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stakebird/x/user/keeper"
	"github.com/public-awesome/stakebird/x/user/types"
)

// NewHandler creates an sdk.Handler for all the x/user type messages
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case *types.MsgVouch:
			return handleMsgVouch(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgVouch(ctx sdk.Context, k keeper.Keeper, msg *types.MsgVouch) (*sdk.Result, error) {
	voucher, err := sdk.AccAddressFromBech32(msg.Voucher)
	if err != nil {
		return nil, err
	}
	vouched, err := sdk.AccAddressFromBech32(msg.Vouched)
	if err != nil {
		return nil, err
	}

	err = k.CreateVouch(
		ctx, voucher, vouched, msg.Comment)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Voucher),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
