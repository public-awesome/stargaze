package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/public-awesome/stargaze/x/ibc-spend/types"
)

// HandleCommunityPoolIBCSpendProposal is a handler for executing a passed community spend proposal
func HandleCommunityPoolIBCSpendProposal(ctx sdk.Context, k Keeper, p *types.CommunityPoolIBCSpendProposal) error {
	moduleAccount := k.ak.GetModuleAddress(types.ModuleName)

	// distribute from community pool to module account
	err := k.distrKeeper.DistributeFromFeePool(ctx, p.Amount, moduleAccount)
	if err != nil {
		return err
	}

	sourcePort := k.transferKeeper.GetPort(ctx)
	sourceChannel := p.SourceChannel
	coinToSend := p.Amount[0]
	sender := moduleAccount
	receiver := p.Recipient
	height := clienttypes.GetSelfHeight(ctx)
	timeoutHeight := clienttypes.NewHeight(height.RevisionNumber, height.RevisionHeight+p.Timeout)

	// ibc xfer from module account
	err = k.transferKeeper.SendTransfer(ctx, sourcePort, sourceChannel, coinToSend, sender, receiver, timeoutHeight, 0)
	if err != nil {
		return err
	}

	logger := k.Logger(ctx)
	logger.Info(
		"transferred from the community pool to IBC recipient", "amount", p.Amount.String(), "recipient", p.Recipient)

	return nil
}
