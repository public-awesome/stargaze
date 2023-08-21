package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/public-awesome/stargaze/v12/x/globalfee/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) SetCodeAuthorization(goCtx context.Context, msg *types.MsgSetCodeAuthorization) (*types.MsgSetCodeAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "sender address is not privileged address")
	}
	err = k.Keeper.SetCodeAuthorization(ctx, *msg.CodeAuthorization)
	if err != nil {
		return nil, err
	}
	return &types.MsgSetCodeAuthorizationResponse{}, nil
}

func (k msgServer) RemoveCodeAuthorization(goCtx context.Context, msg *types.MsgRemoveCodeAuthorization) (*types.MsgRemoveCodeAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "sender address is not privileged address")
	}
	k.Keeper.DeleteCodeAuthorization(ctx, msg.GetCodeID())
	return &types.MsgRemoveCodeAuthorizationResponse{}, nil
}

func (k msgServer) SetContractAuthorization(goCtx context.Context, msg *types.MsgSetContractAuthorization) (*types.MsgSetContractAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "sender address is not privileged address")
	}
	err = k.Keeper.SetContractAuthorization(ctx, *msg.ContractAuthorization)
	if err != nil {
		return nil, err
	}
	return &types.MsgSetContractAuthorizationResponse{}, nil
}

func (k msgServer) RemoveContractAuthorization(goCtx context.Context, msg *types.MsgRemoveContractAuthorization) (*types.MsgRemoveContractAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "sender address is not privileged address")
	}
	contractAddr, err := sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return nil, err
	}
	k.Keeper.DeleteContractAuthorization(ctx, contractAddr)
	return &types.MsgRemoveContractAuthorizationResponse{}, nil
}
