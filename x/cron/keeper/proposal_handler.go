package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	errorsmod "cosmossdk.io/errors"
	"github.com/public-awesome/stargaze/v12/x/cron/types"
)

// govKeeper is a subset of Keeper that is needed for the gov proposal handling
type govKeeper interface {
	SetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) error
	UnsetPrivileged(ctx sdk.Context, contractAddr sdk.AccAddress) error
}

// NewProposalHandler creates a new governance Handler for wasm proposals
func NewProposalHandler(k Keeper) govtypes.Handler {
	return NewProposalHandlerX(k)
}

func NewProposalHandlerX(k govKeeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.PromoteToPrivilegedContractProposal:
			return handlePromoteContractProposal(ctx, k, *c)
		case *types.DemotePrivilegedContractProposal:
			return handleDemoteContractProposal(ctx, k, *c)
		default:
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized cron srcProposal content type: %T", c)
		}
	}
}

func handlePromoteContractProposal(ctx sdk.Context, k govKeeper, p types.PromoteToPrivilegedContractProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	contractAddr, err := sdk.AccAddressFromBech32(p.Contract)
	if err != nil {
		return errorsmod.Wrap(err, "contract address")
	}

	err = k.SetPrivileged(ctx, contractAddr)
	if err != nil {
		return err
	}
	return nil
}

func handleDemoteContractProposal(ctx sdk.Context, k govKeeper, p types.DemotePrivilegedContractProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}
	contractAddr, err := sdk.AccAddressFromBech32(p.Contract)
	if err != nil {
		return errorsmod.Wrap(err, "contract address")
	}

	err = k.UnsetPrivileged(ctx, contractAddr)
	if err != nil {
		return err
	}
	return nil
}
