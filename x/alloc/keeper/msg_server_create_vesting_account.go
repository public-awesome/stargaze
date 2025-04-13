package keeper

import (
	"context"

	"github.com/public-awesome/stargaze/v16/x/alloc/types"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateVestingAccount(_ context.Context, _ *types.MsgCreateVestingAccount) (*types.MsgCreateVestingAccountResponse, error) { //nolint:staticcheck
	return nil, errorsmod.Wrapf(sdkerrors.ErrNotSupported, "support for creating vesting account has been removed in favor of sdk's version")
}
