package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/public-awesome/stargaze/x/ibc-spend/types"
)

// HandleCommunityPoolIBCSpendProposal is a handler for executing a passed community spend proposal
func HandleCommunityPoolIBCSpendProposal(ctx sdk.Context, k Keeper, p *types.CommunityPoolIBCSpendProposal) error {
	logger := k.Logger(ctx)

	moduleAddr := k.ak.GetModuleAddress(types.ModuleName)

	// distribute from community pool to module account
	err := k.distrKeeper.DistributeFromFeePool(ctx, p.Amount, moduleAddr)
	if err != nil {
		return err
	}

	fmt.Println(p.String())

	sourcePort := k.transferKeeper.GetPort(ctx)
	sourceChannel := p.SourceChannel
	coinToSend := p.Amount[0]
	sender := moduleAddr
	receiver := p.Recipient
	// height := clienttypes.GetSelfHeight(ctx)
	// timeoutHeight := clienttypes.NewHeight(height.RevisionNumber, height.RevisionHeight+p.Timeout)
	timeoutHeight := clienttypes.NewHeight(0, 110)

	fmt.Printf("sourcePort %v\n", sourcePort)
	fmt.Printf("sourceChannel %v\n", sourceChannel)
	fmt.Printf("coinToSend %v\n", coinToSend.String())
	fmt.Printf("sender %v\n", sender.String())
	fmt.Printf("receiver %v\n", receiver)
	fmt.Printf("timeoutHeight %v\n", timeoutHeight)

	// ibc xfer from module account
	err = k.transferKeeper.SendTransfer(ctx, sourcePort, sourceChannel, coinToSend, sender, receiver, timeoutHeight, 0)
	if err != nil {
		return err
	}

	fmt.Println(
		"transferred from the community pool to IBC recipient", "amount", p.Amount.String(), "recipient", p.Recipient)

	logger.Info(
		"transferred from the community pool to IBC recipient", "amount", p.Amount.String(), "recipient", p.Recipient)

	return nil
}
