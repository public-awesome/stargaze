package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v3/x/claim/types"
)

func (k msgServer) ClaimFor(goCtx context.Context, msg *types.MsgClaimFor) (*types.MsgClaimForResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgClaimForResponse{}, nil
}
