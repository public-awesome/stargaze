package keeper

import (
	"context"

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

	_, err = m.Keeper.ExecuteProposalMessages(ctx, proposalMsgs, proposer)
	if err != nil {
		return nil, err
	}

	return &types.MsgExecuteProposalResponse{}, nil
}
