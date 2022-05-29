package keeper

import (
	"context"

	"github.com/public-awesome/stargaze/v5/x/alloc/types"
)

func (k msgServer) FundFeeCollector(goCtx context.Context, msg *types.MsgFundFeeCollector) (*types.MsgFundFeeCollectorResponse, error) {
	return &types.MsgFundFeeCollectorResponse{}, nil
}
