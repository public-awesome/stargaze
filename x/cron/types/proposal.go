package types

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
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
	v1beta1.RegisterProposalType(string(ProposalTypePromoteContract))
	v1beta1.RegisterProposalType(string(ProposalTypeDemoteContract))
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
		return errorsmod.Wrap(err, "contract")
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
		return errorsmod.Wrap(err, "contract")
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
		return errorsmod.Wrap(govtypes.ErrInvalidProposalContent, "proposal title must not start/end with white spaces")
	}
	if len(title) == 0 {
		return errorsmod.Wrap(govtypes.ErrInvalidProposalContent, "proposal title cannot be blank")
	}
	if len(title) > v1beta1.MaxTitleLength {
		return errorsmod.Wrapf(govtypes.ErrInvalidProposalContent, "proposal title is longer than max length of %d", v1beta1.MaxTitleLength)
	}
	if strings.TrimSpace(description) != description {
		return errorsmod.Wrap(govtypes.ErrInvalidProposalContent, "proposal description must not start/end with white spaces")
	}
	if len(description) == 0 {
		return errorsmod.Wrap(govtypes.ErrInvalidProposalContent, "proposal description cannot be blank")
	}
	if len(description) > v1beta1.MaxDescriptionLength {
		return errorsmod.Wrapf(govtypes.ErrInvalidProposalContent, "proposal description is longer than max length of %d", v1beta1.MaxDescriptionLength)
	}
	return nil
}
