package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/public-awesome/stargaze/v14/x/cron/types"
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

func (m msgServer) PromoteToPrivilegedContract(goCtx context.Context, msg *types.MsgPromoteToPrivilegedContract) (*types.MsgPromoteToPrivilegedContractResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	authorityAddr, err := sdk.AccAddressFromBech32(msg.GetAuthority())
	if err != nil {
		return nil, err
	}

	if !m.isAuthorized(ctx, authorityAddr.String()) {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender address is not authorized address to promote contracts")
	}

	contractAddr, err := sdk.AccAddressFromBech32(msg.GetContract())
	if err != nil {
		return nil, err
	}
	err = m.SetPrivileged(ctx, contractAddr)
	if err != nil {
		return nil, err
	}

	return &types.MsgPromoteToPrivilegedContractResponse{}, nil
}

func (m msgServer) DemoteFromPrivilegedContract(goCtx context.Context, msg *types.MsgDemoteFromPrivilegedContract) (*types.MsgDemoteFromPrivilegedContractResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	authorityAddr, err := sdk.AccAddressFromBech32(msg.GetAuthority())
	if err != nil {
		return nil, err
	}

	if !m.isAuthorized(ctx, authorityAddr.String()) {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender address is not authorized address to demote contracts")
	}

	contractAddr, err := sdk.AccAddressFromBech32(msg.GetContract())
	if err != nil {
		return nil, err
	}
	err = m.UnsetPrivileged(ctx, contractAddr)
	if err != nil {
		return nil, err
	}
	return &types.MsgDemoteFromPrivilegedContractResponse{}, nil
}

func (m msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.GetAuthority() != m.Keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "sender address is not authorized address to update module params")
	}

	// need to explicitly validate as x/gov invokes this msg and it does not validate
	if err := msg.GetParams().Validate(); err != nil {
		return nil, err
	}

	err := m.SetParams(ctx, msg.GetParams())
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

func (m msgServer) isAuthorized(ctx sdk.Context, actor string) bool {
	if actor == m.Keeper.GetAuthority() {
		return true
	}
	if m.Keeper.IsAdminAddress(ctx, actor) {
		return true
	}
	return false
}
