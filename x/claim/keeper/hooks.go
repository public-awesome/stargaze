package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/public-awesome/stargaze/v12/x/claim/types"
)

func (k Keeper) AfterProposalVote(ctx sdk.Context, _ uint64, voterAddr sdk.AccAddress) {
	_, err := k.ClaimCoinsForAction(ctx, voterAddr, types.ActionVote)
	if err != nil {
		panic(err.Error())
	}
}

func (k Keeper) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, _ sdk.ValAddress) {
	// must not run on genesis
	if ctx.BlockHeight() <= 1 {
		return
	}
	_, err := k.ClaimCoinsForAction(ctx, delAddr, types.ActionDelegateStake)
	if err != nil {
		panic(err.Error())
	}
}

// ________________________________________________________________________________________

// Hooks wrapper struct for slashing keeper
type Hooks struct {
	k Keeper
}

var (
	_ govtypes.GovHooks         = Hooks{}
	_ stakingtypes.StakingHooks = Hooks{}
)

// Return the wrapper struct
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// governance hooks
func (h Hooks) AfterProposalSubmission(_ sdk.Context, _ uint64) {}

func (h Hooks) AfterProposalDeposit(_ sdk.Context, _ uint64, _ sdk.AccAddress) {
}

func (h Hooks) AfterProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	h.k.AfterProposalVote(ctx, proposalID, voterAddr)
}

func (h Hooks) AfterProposalInactive(_ sdk.Context, _ uint64) {}
func (h Hooks) AfterProposalActive(_ sdk.Context, _ uint64)   {}

func (h Hooks) AfterProposalFailedMinDeposit(_ sdk.Context, _ uint64)  {}
func (h Hooks) AfterProposalVotingPeriodEnded(_ sdk.Context, _ uint64) {}

// staking hooks
func (h Hooks) AfterValidatorCreated(_ sdk.Context, _ sdk.ValAddress)   {}
func (h Hooks) BeforeValidatorModified(_ sdk.Context, _ sdk.ValAddress) {}
func (h Hooks) AfterValidatorRemoved(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
}

func (h Hooks) AfterValidatorBonded(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
}

func (h Hooks) AfterValidatorBeginUnbonding(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) {
}

func (h Hooks) BeforeDelegationCreated(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
}

func (h Hooks) BeforeDelegationSharesModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
}

func (h Hooks) BeforeDelegationRemoved(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) {
}

func (h Hooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.k.AfterDelegationModified(ctx, delAddr, valAddr)
}
func (h Hooks) BeforeValidatorSlashed(_ sdk.Context, _ sdk.ValAddress, _ sdk.Dec) {}
