package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/public-awesome/stargaze/v10/x/globalfee/types"
)

type govKeeper interface {
	SetCodeAuthorization(ctx sdk.Context, ca types.CodeAuthorization) error
	DeleteCodeAuthorization(ctx sdk.Context, codeID uint64)
	SetContractAuthorization(ctx sdk.Context, ca types.ContractAuthorization) error
	DeleteContractAuthorization(ctx sdk.Context, contractAddr sdk.AccAddress)
}

// NewProposalHandler creates a new governance Handler for wasm proposals
func NewProposalHandler(k Keeper) govtypes.Handler {
	return NewProposalHandlerX(k)
}

func NewProposalHandlerX(k govKeeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.SetCodeAuthorizationProposal:
			return handleSetCodeAuthorizationProposal(ctx, k, *c)
		case *types.RemoveCodeAuthorizationProposal:
			return handleDeleteCodeAuthorizationProposal(ctx, k, *c)
		case *types.SetContractAuthorizationProposal:
			return handleSetContractAuthorizationProposal(ctx, k, *c)
		case *types.RemoveContractAuthorizationProposal:
			return handleDeleteContractAuthorizationProposal(ctx, k, *c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized globalfee srcProposal content type: %T", c)
		}
	}
}

func handleSetCodeAuthorizationProposal(ctx sdk.Context, k govKeeper, p types.SetCodeAuthorizationProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	return k.SetCodeAuthorization(ctx, *p.GetCodeAuthorization())
}

func handleDeleteCodeAuthorizationProposal(ctx sdk.Context, k govKeeper, p types.RemoveCodeAuthorizationProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	k.DeleteCodeAuthorization(ctx, p.GetCodeId())
	return nil
}

func handleSetContractAuthorizationProposal(ctx sdk.Context, k govKeeper, p types.SetContractAuthorizationProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	return k.SetContractAuthorization(ctx, *p.GetContractAuthorization())
}

func handleDeleteContractAuthorizationProposal(ctx sdk.Context, k govKeeper, p types.RemoveContractAuthorizationProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	contractAddr, err := sdk.AccAddressFromBech32(p.GetContractAddress())
	if err != nil {
		return sdkerrors.Wrap(err, "invalid contract address")
	}

	k.DeleteContractAuthorization(ctx, contractAddr)
	return nil
}
