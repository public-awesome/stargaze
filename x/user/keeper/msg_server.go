package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/user/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the user MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) Vouch(goCtx context.Context, msg *types.MsgVouch) (*types.MsgVouchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
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

	return &types.MsgVouchResponse{}, nil
}
