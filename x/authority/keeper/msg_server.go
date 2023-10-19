package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v13/x/authority/types"
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

// ExecuteProposal implements types.MsgServer.
func (m msgServer) ExecuteProposal(goCtx context.Context, msg *types.MsgExecuteProposal) (*types.MsgExecuteProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposalMsgs, err := msg.GetMsgs()
	if err != nil {
		return nil, err
	}

	proposer, err := sdk.AccAddressFromBech32(msg.GetAuthority())
	if err != nil {
		return nil, err
	}

	valid, err := m.Keeper.IsAuthorized(ctx, proposalMsgs, msg.GetAuthority())
	if !valid {
		return nil, err
	}

	_, err = m.Keeper.SubmitProposal(ctx, proposalMsgs, proposer)
	if err != nil {
		return nil, err
	}

	return &types.MsgExecuteProposalResponse{}, nil
}

func (keeper Keeper) IsAuthorized(ctx sdk.Context, msgs []sdk.Msg, proposer string) (bool, error) {
	authorizations := keeper.GetParams(ctx).Authorizations
	for _, msg := range msgs { // Checking authorizations for all the msgs in the proposal
		msgType := sdk.MsgTypeURL(msg)
		auth, found := types.GetMsgAuthorization(msgType, authorizations)
		if !found {
			return false, errorsmod.Wrap(types.ErrAuthorizationNotFound, "authorization not found for given msg type: "+msgType)
		}

		if !auth.IsAuthorized(proposer) {
			return false, errorsmod.Wrap(types.ErrUnauthorized, "authority address is not authorized address to ")
		}

	}
	return true, nil
}
