package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v17/x/pauser/types"
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

func (k msgServer) PauseContract(goCtx context.Context, msg *types.MsgPauseContract) (*types.MsgPauseContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender is not a privileged address")
	}

	contractAddr, err := sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return nil, err
	}

	if !k.wasmKeeper.HasContractInfo(ctx, contractAddr) {
		return nil, types.ErrContractNotExist
	}

	if k.Keeper.IsContractPaused(ctx, contractAddr) {
		return nil, errorsmod.Wrap(types.ErrAlreadyPaused, "contract is already paused")
	}

	pc := types.PausedContract{
		ContractAddress: msg.ContractAddress,
		PausedBy:        msg.Sender,
		PausedAt:        ctx.BlockTime(),
	}
	if err := k.SetPausedContract(ctx, pc); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeContractPaused,
			sdk.NewAttribute(types.AttributeKeyContractAddress, msg.ContractAddress),
			sdk.NewAttribute(types.AttributeKeyPausedBy, msg.Sender),
		),
	)

	return &types.MsgPauseContractResponse{}, nil
}

func (k msgServer) UnpauseContract(goCtx context.Context, msg *types.MsgUnpauseContract) (*types.MsgUnpauseContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender is not a privileged address")
	}

	contractAddr, err := sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return nil, err
	}

	if !k.Keeper.IsContractPaused(ctx, contractAddr) {
		return nil, errorsmod.Wrap(types.ErrNotPaused, "contract is not paused")
	}

	if err := k.DeletePausedContract(ctx, contractAddr); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeContractUnpaused,
			sdk.NewAttribute(types.AttributeKeyContractAddress, msg.ContractAddress),
			sdk.NewAttribute(types.AttributeKeyPausedBy, msg.Sender),
		),
	)

	return &types.MsgUnpauseContractResponse{}, nil
}

func (k msgServer) PauseCodeID(goCtx context.Context, msg *types.MsgPauseCodeID) (*types.MsgPauseCodeIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender is not a privileged address")
	}

	if k.wasmKeeper.GetCodeInfo(ctx, msg.CodeID) == nil {
		return nil, types.ErrCodeIDNotExist
	}

	if k.Keeper.IsCodeIDPaused(ctx, msg.CodeID) {
		return nil, errorsmod.Wrap(types.ErrAlreadyPaused, "code ID is already paused")
	}

	pc := types.PausedCodeID{
		CodeID:   msg.CodeID,
		PausedBy: msg.Sender,
		PausedAt: ctx.BlockTime(),
	}
	if err := k.SetPausedCodeID(ctx, pc); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCodeIDPaused,
			sdk.NewAttribute(types.AttributeKeyCodeID, strconv.FormatUint(msg.CodeID, 10)),
			sdk.NewAttribute(types.AttributeKeyPausedBy, msg.Sender),
		),
	)

	return &types.MsgPauseCodeIDResponse{}, nil
}

func (k msgServer) UnpauseCodeID(goCtx context.Context, msg *types.MsgUnpauseCodeID) (*types.MsgUnpauseCodeIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !k.IsPrivilegedAddress(ctx, msg.Sender) {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender is not a privileged address")
	}

	if !k.Keeper.IsCodeIDPaused(ctx, msg.CodeID) {
		return nil, errorsmod.Wrap(types.ErrNotPaused, "code ID is not paused")
	}

	if err := k.DeletePausedCodeID(ctx, msg.CodeID); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCodeIDUnpaused,
			sdk.NewAttribute(types.AttributeKeyCodeID, strconv.FormatUint(msg.CodeID, 10)),
			sdk.NewAttribute(types.AttributeKeyPausedBy, msg.Sender),
		),
	)

	return &types.MsgUnpauseCodeIDResponse{}, nil
}

func (k msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	if msg.Sender != k.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender address is not the authority for updating module params")
	}

	if err := msg.GetParams().Validate(); err != nil {
		return nil, err
	}

	err = k.SetParams(ctx, msg.GetParams())
	return &types.MsgUpdateParamsResponse{}, err
}
