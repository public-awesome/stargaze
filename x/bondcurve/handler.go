package bondcurve

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rocket-protocol/stakebird/x/bondcurve/types"
)

// NewHandler creates an sdk.Handler for all the bondcurve type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// case MsgBuy:
		// 	return handleMsgBuy(ctx, k, msg)
		case types.MsgLockCollateral:
			return handleMsgLockCollateral(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleBuy does x
func handleMsgBuy(ctx sdk.Context, k Keeper, msg types.MsgBuy) (*sdk.Result, error) {
	// err := k.Buy(ctx, msg.ValidatorAddr)
	// if err != nil {
	// 	return nil, err
	// }

	// TODO: Define your msg events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgLockCollateral(ctx sdk.Context, k Keeper, msg types.MsgLockCollateral) (*sdk.Result, error) {
	// transfer/ibczeroxfer/stake
	denom := fmt.Sprintf("transfer/%s/%s", types.Counterparty, types.CounterpartyDenom)
	lockCoin := sdk.NewCoin(denom, msg.Amount.Amount)

	err := k.SupplyKeeper.SendCoinsFromAccountToModule(ctx, msg.Sender, ModuleName, sdk.NewCoins(lockCoin))
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds, "can't transfer %s coins from sender to module account", denom)
	}

	newCoin := sdk.NewCoin(types.Denom, msg.Amount.Amount)
	err = k.SupplyKeeper.MintCoins(ctx, ModuleName, sdk.NewCoins(newCoin))
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "can't mint %s", newCoin.Denom)
	}
	err = k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, ModuleName, msg.Sender, sdk.NewCoins(newCoin))
	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds, "can't transfer %s coins from module account to sender", newCoin.Denom)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
