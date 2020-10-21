package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stakebird/x/faucet/internal/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the curating MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	minter, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return nil, err
	}
	k.Logger(ctx).Info("received mint message: %s", msg)

	if strings.TrimSpace(msg.Denom) == "" {
		return nil, sdkerrors.Wrap(err, "invalid denomination")
	}
	err = k.MintAndSend(ctx, minter, msg.Time, msg.Denom)
	if err != nil {
		return nil, sdkerrors.Wrap(err, fmt.Sprintf(",in [%v] hours", k.Limit().Hours()))
	}

	return &types.MsgMintResponse{}, nil
}

func (k msgServer) FaucetKey(goCtx context.Context, msg *types.MsgFaucetKey) (*types.MsgFaucetKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("received faucet message: %s", msg)
	if k.HasFaucetKey(ctx) {
		return nil, types.ErrFaucetKeyExisted
	}
	k.SetFaucetKey(ctx, msg.Armor)

	return &types.MsgFaucetKeyResponse{}, nil
}
