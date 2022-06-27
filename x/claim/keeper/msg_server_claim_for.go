package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v6/x/claim/types"
)

func (k msgServer) ClaimFor(goCtx context.Context, msg *types.MsgClaimFor) (*types.MsgClaimForResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}
	params := k.GetParams(ctx)
	if !params.IsAirdropEnabled(ctx.BlockTime()) {
		return nil, types.ErrAirdropNotEnabled
	}
	allowed := false
	for _, authorization := range params.AllowedClaimers {
		if authorization.ContractAddress == msg.Sender && authorization.Action == msg.Action {
			allowed = true
			break
		}
	}
	if !allowed {
		return nil, types.ErrUnauthorizedClaimer
	}
	coins, err := k.Keeper.ClaimCoinsForAction(ctx, address, msg.GetAction())
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
	})
	return &types.MsgClaimForResponse{
		Address:       msg.Address,
		ClaimedAmount: coins,
	}, nil
}
