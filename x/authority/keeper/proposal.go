package keeper

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	govV1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

// IterateActiveProposalsQueue iterates over the proposals in the active proposal queue
// and performs a callback function
func (k Keeper) IterateActiveProposalsQueue(ctx sdk.Context, cb func(proposal govV1.Proposal) (stop bool)) {
	iterator := k.ActiveProposalQueueIterator(ctx)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		proposalID := types.GetProposalIDFromBytes(iterator.Value())
		proposal, found := k.GetProposal(ctx, proposalID)
		if !found {
			panic(fmt.Sprintf("proposal %d does not exist", proposalID))
		}

		if cb(proposal) {
			break
		}
	}
}

// ActiveProposalQueueIterator returns an sdk.Iterator for all the proposals in the Active Queue
func (k Keeper) ActiveProposalQueueIterator(ctx sdk.Context) sdk.Iterator {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.ActiveProposalQueuePrefix)
	return prefixStore.Iterator(nil, nil)
}

// GetProposal get proposal from store by ProposalID
func (k Keeper) GetProposal(ctx sdk.Context, proposalID uint64) (govV1.Proposal, bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.ProposalKey(proposalID))
	if bz == nil {
		return govV1.Proposal{}, false
	}

	var proposal govV1.Proposal
	k.MustUnmarshalProposal(bz, &proposal)

	return proposal, true
}

func (k Keeper) MustUnmarshalProposal(bz []byte, proposal *govV1.Proposal) {
	err := k.UnmarshalProposal(bz, proposal)
	if err != nil {
		panic(err)
	}
}

func (k Keeper) UnmarshalProposal(bz []byte, proposal *govV1.Proposal) error {
	err := k.cdc.Unmarshal(bz, proposal)
	if err != nil {
		return err
	}
	return nil
}

// SubmitProposal creates a new proposal given an array of messages
func (keeper Keeper) SubmitProposal(ctx sdk.Context, messages []sdk.Msg, proposer sdk.AccAddress) (uint64, error) {
	// Will hold a comma-separated string of all Msg type URLs.
	msgsStr := ""

	var (
		events sdk.Events
	)

	// Loop through all messages and confirm that each has a handler and the gov module account
	// as the only signer
	for _, msg := range messages {
		msgsStr += fmt.Sprintf(",%s", sdk.MsgTypeURL(msg))

		// perform a basic validation of the message
		if err := msg.ValidateBasic(); err != nil {
			return 0, errorsmod.Wrap(types.ErrInvalidProposalMsg, err.Error())
		}

		signers := msg.GetSigners()
		if len(signers) != 1 {
			return 0, types.ErrInvalidSigner
		}

		// // assert that the governance module account is the only signer of the messages
		// if !signers[0].Equals(keeper.GetGovernanceAccount(ctx).GetAddress()) {
		// 	return v1.Proposal{}, sdkerrors.Wrapf(types.ErrInvalidSigner, signers[0].String())
		// }

		// use the msg service router to see that there is a valid route for that message.
		handler := keeper.router.Handler(msg)
		if handler == nil {
			return 0, errorsmod.Wrap(types.ErrUnroutableProposalMsg, sdk.MsgTypeURL(msg))
		}

		var res *sdk.Result
		res, err := handler(ctx, msg)
		if err != nil {
			return 0, err
		}

		events = append(events, res.GetEvents()...)

	}

	return 0, nil
}

// // SetProposal set a proposal to store
// func (k Keeper) SetProposal(ctx sdk.Context, proposal govV1.Proposal) {
// 	store := ctx.KVStore(k.storeKey)

// 	bz := k.MustMarshalProposal(proposal)

// 	store.Set(types.ProposalKey(proposal.Id), bz)
// }

// // GetProposalID gets the highest proposal ID
// func (k Keeper) GetProposalID(ctx sdk.Context) (proposalID uint64, err error) {
// 	store := ctx.KVStore(k.storeKey)
// 	bz := store.Get(types.ProposalIDKey)
// 	if bz == nil {
// 		return 0, errorsmod.Wrap(types.ErrInvalidGenesis, "initial proposal ID hasn't been set")
// 	}

// 	proposalID = types.GetProposalIDFromBytes(bz)
// 	return proposalID, nil
// }

// // SetProposalID sets the new proposal ID to the store
// func (k Keeper) SetProposalID(ctx sdk.Context, proposalID uint64) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Set(types.ProposalIDKey, types.GetProposalIDBytes(proposalID))
// }

// // InsertActiveProposalQueue inserts a ProposalID into the active proposal queue
// func (k Keeper) InsertActiveProposalQueue(ctx sdk.Context, proposalID uint64) {
// 	store := ctx.KVStore(k.storeKey)
// 	store.Set(types.ActiveProposalQueueKey(proposalID), types.GetProposalIDBytes(proposalID))
// }
