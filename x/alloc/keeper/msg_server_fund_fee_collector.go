package keeper

import (
	"context"

	"github.com/public-awesome/stargaze/v5/x/alloc/types"
)

func (k msgServer) FundFairburnPool(goCtx context.Context, msg *types.MsgFundFairburnPool) (*types.MsgFundFairburnPoolResponse, error) {
	return &types.MsgFundFairburnPoolResponse{}, nil
}
