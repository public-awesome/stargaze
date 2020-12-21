package keeper

import (
	"context"
	"strings"

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
	return &types.MsgPostResponse{}, nil
}
