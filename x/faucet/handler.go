package faucet

import (
	"fmt"
	"strings"

	"github.com/public-awesome/stakebird/x/faucet/internal/keeper"
	"github.com/public-awesome/stakebird/x/faucet/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "faucet" type messages.
func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case *types.MsgMint:
			return handleMsgMint(ctx, keeper, msg)
		case *types.MsgFaucetKey:
			return handleMsgFaucetKey(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized faucet Msg type: %v", msg.Type()))
		}
	}
}

// Handle a message to Mint
func handleMsgMint(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgMint) (*sdk.Result, error) {
	minter, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return nil, err
	}
	keeper.Logger(ctx).Info("received mint message: %s", msg)

	if strings.TrimSpace(msg.Denom) == "" {
		return nil, sdkerrors.Wrap(err, "invalid denomination")
	}
	err = keeper.MintAndSend(ctx, minter, msg.Time, msg.Denom)
	if err != nil {
		return nil, sdkerrors.Wrap(err, fmt.Sprintf(",in [%v] hours", keeper.Limit().Hours()))
	}

	return &sdk.Result{}, nil
}

// Handle a message to Mint
func handleMsgFaucetKey(ctx sdk.Context, keeper keeper.Keeper, msg *types.MsgFaucetKey) (*sdk.Result, error) {
	keeper.Logger(ctx).Info("received faucet message: %s", msg)
	if keeper.HasFaucetKey(ctx) {
		return nil, types.ErrFaucetKeyExisted
	}
	keeper.SetFaucetKey(ctx, msg.Armor)

	return &sdk.Result{}, nil
}
