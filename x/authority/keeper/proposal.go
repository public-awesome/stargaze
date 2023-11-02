package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/public-awesome/stargaze/v13/x/authority/types"
)

// ExecuteProposalMessages executes the array of messages in the proposal if their authorization matches
func (keeper Keeper) ExecuteProposalMessages(ctx sdk.Context, messages []sdk.Msg, proposer sdk.AccAddress) (uint64, error) {
	var events sdk.Events

	// Loop through all messages and confirm that each has a handler and the authorizations for the msg are valid
	for _, msg := range messages {
		// perform a basic validation of the message
		if err := msg.ValidateBasic(); err != nil {
			return 0, errorsmod.Wrap(govtypes.ErrInvalidProposalMsg, err.Error())
		}

		signers := msg.GetSigners()
		if len(signers) != 1 {
			return 0, govtypes.ErrInvalidSigner
		}

		valid, err := keeper.IsAuthorized(ctx, msg, proposer.String())
		if !valid {
			return 0, err
		}

		// use the msg service router to see that there is a valid route for that message.
		handler := keeper.router.Handler(msg)
		if handler == nil {
			return 0, errorsmod.Wrap(govtypes.ErrUnroutableProposalMsg, sdk.MsgTypeURL(msg))
		}

		var res *sdk.Result
		res, err = handler(ctx, msg)
		if err != nil {
			return 0, err
		}

		events = append(events, res.GetEvents()...)
	}

	ctx.EventManager().EmitEvents(events)
	return 0, nil
}

func (keeper Keeper) IsAuthorized(ctx sdk.Context, msg sdk.Msg, proposer string) (bool, error) {
	authorizations := keeper.GetParams(ctx).Authorizations
	msgType := sdk.MsgTypeURL(msg)
	auth, found := types.GetMsgAuthorization(msgType, authorizations)
	if !found {
		return false, errorsmod.Wrap(types.ErrAuthorizationNotFound, "authorization not found for given msg type: "+msgType)
	}

	if !auth.IsAuthorized(proposer) {
		return false, errorsmod.Wrap(govtypes.ErrInvalidSigner, "sender address "+proposer+" is not authorized to execute this proposal"+msgType)
	}
	return true, nil
}
