package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/stake/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the curating MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) Stake(goCtx context.Context, msg *types.MsgStake) (*types.MsgStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// creator, err := sdk.AccAddressFromBech32(msg.Delegator)
	// if err != nil {
	// 	return nil, err
	// }

	// err = k.CreatePost(
	// 	ctx, msg.VendorID, msg.PostID, msg.Body, creator, rewardAccount)
	// if err != nil {
	// 	return nil, err
	// }

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Delegator),
		),
	})
	return &types.MsgStakeResponse{}, nil
}
