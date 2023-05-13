package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

type ProposalType string

const (
	ProposalTypePromoteContract ProposalType = "PromoteToPrivilegedContract"
	ProposalTypeDemoteContract  ProposalType = "DemotePrivilegedContract"
)

// EnableAllProposals contains all twasm gov types as keys.
var EnableAllProposals = []ProposalType{
	ProposalTypePromoteContract,
	ProposalTypeDemoteContract,
}

func init() { // register new content types with the sdk
	govv1beta1.RegisterProposalType(string(ProposalTypePromoteContract))
	govv1beta1.RegisterProposalType(string(ProposalTypeDemoteContract))

	govv1.RegisterProposalTypeCodec(&PromoteToPrivilegedContractProposal{}, "cron/PromoteToPrivilegedContractProposal")
	govtypes.RegisterProposalTypeCodec(&DemotePrivilegedContractProposal{}, "cron/DemotePrivilegedContractProposal")
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p PromoteToPrivilegedContractProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type
func (p PromoteToPrivilegedContractProposal) ProposalType() string {
	return string(ProposalTypePromoteContract)
}

// ValidateBasic validates the proposal
func (p PromoteToPrivilegedContractProposal) ValidateBasic() error {
	if err := validateProposalCommons(p.Title, p.Description); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(p.Contract); err != nil {
		return sdkerrors.Wrap(err, "contract")
	}
	return nil
}

// MarshalYAML pretty prints the wasm byte code
func (p PromoteToPrivilegedContractProposal) MarshalYAML() (interface{}, error) {
	return p, nil
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p DemotePrivilegedContractProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type
func (p DemotePrivilegedContractProposal) ProposalType() string {
	return string(ProposalTypeDemoteContract)
}

// ValidateBasic validates the proposal
func (p DemotePrivilegedContractProposal) ValidateBasic() error {
	if err := validateProposalCommons(p.Title, p.Description); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(p.Contract); err != nil {
		return sdkerrors.Wrap(err, "contract")
	}
	return nil
}

// MarshalYAML pretty prints the wasm byte code
func (p DemotePrivilegedContractProposal) MarshalYAML() (interface{}, error) {
	return p, nil
}

// common validations
func validateProposalCommons(title, description string) error {
	if strings.TrimSpace(title) != title {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "proposal title must not start/end with white spaces")
	}
	if len(title) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "proposal title cannot be blank")
	}
	if strings.TrimSpace(description) != description {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "proposal description must not start/end with white spaces")
	}
	if len(description) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "proposal description cannot be blank")
	}
	return nil
}
