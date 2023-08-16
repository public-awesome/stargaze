package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	legacygovtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

type ProposalType string

const (
	ProposalTypeSetCodeAuthorization        ProposalType = "ProposalTypeSetCodeAuthorization"
	ProposalTypeRemoveCodeAuthorization     ProposalType = "ProposalTypeRemoveCodeAuthorization"
	ProposalTypeSetContractAuthorization    ProposalType = "ProposalTypeSetContractAuthorization"
	ProposalTypeRemoveContractAuthorization ProposalType = "ProposalTypeRemoveContractAuthorization"
)

// EnableAllProposals contains all twasm gov types as keys.
var EnableAllProposals = []ProposalType{
	ProposalTypeSetCodeAuthorization,
	ProposalTypeRemoveCodeAuthorization,
	ProposalTypeSetContractAuthorization,
	ProposalTypeRemoveContractAuthorization,
}

func init() { // register new content types with the sdk
	legacygovtypes.RegisterProposalType(string(ProposalTypeSetCodeAuthorization))
	legacygovtypes.RegisterProposalType(string(ProposalTypeRemoveCodeAuthorization))
	legacygovtypes.RegisterProposalType(string(ProposalTypeSetContractAuthorization))
	legacygovtypes.RegisterProposalType(string(ProposalTypeRemoveContractAuthorization))

	// legacygovtypes.RegisterProposalTypeCodec(&SetCodeAuthorizationProposal{}, "globalfee/SetCodeAuthorizationProposal")
	// legacygovtypes.RegisterProposalTypeCodec(&RemoveCodeAuthorizationProposal{}, "globalfee/RemoveCodeAuthorizationProposal")
	// legacygovtypes.RegisterProposalTypeCodec(&SetContractAuthorizationProposal{}, "globalfee/SetContractAuthorizationProposal")
	// legacygovtypes.RegisterProposalTypeCodec(&RemoveContractAuthorizationProposal{}, "globalfee/RemoveContractAuthorizationProposal")
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p SetCodeAuthorizationProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type
func (p SetCodeAuthorizationProposal) ProposalType() string {
	return string(ProposalTypeSetCodeAuthorization)
}

// ValidateBasic validates the proposal
func (p SetCodeAuthorizationProposal) ValidateBasic() error {
	if err := validateProposalCommons(p.Title, p.Description); err != nil {
		return err
	}
	return p.CodeAuthorization.Validate()
}

// MarshalYAML pretty prints the wasm byte code
func (p SetCodeAuthorizationProposal) MarshalYAML() (interface{}, error) {
	return p, nil
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p RemoveCodeAuthorizationProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type
func (p RemoveCodeAuthorizationProposal) ProposalType() string {
	return string(ProposalTypeRemoveCodeAuthorization)
}

// ValidateBasic validates the proposal
func (p RemoveCodeAuthorizationProposal) ValidateBasic() error {
	return validateProposalCommons(p.Title, p.Description)
}

// MarshalYAML pretty prints the wasm byte code
func (p RemoveCodeAuthorizationProposal) MarshalYAML() (interface{}, error) {
	return p, nil
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p SetContractAuthorizationProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type
func (p SetContractAuthorizationProposal) ProposalType() string {
	return string(ProposalTypeSetContractAuthorization)
}

// ValidateBasic validates the proposal
func (p SetContractAuthorizationProposal) ValidateBasic() error {
	if err := validateProposalCommons(p.Title, p.Description); err != nil {
		return err
	}
	return p.ContractAuthorization.Validate()
}

// MarshalYAML pretty prints the wasm byte code
func (p SetContractAuthorizationProposal) MarshalYAML() (interface{}, error) {
	return p, nil
}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p RemoveContractAuthorizationProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type
func (p RemoveContractAuthorizationProposal) ProposalType() string {
	return string(ProposalTypeRemoveContractAuthorization)
}

// ValidateBasic validates the proposal
func (p RemoveContractAuthorizationProposal) ValidateBasic() error {
	if err := validateProposalCommons(p.Title, p.Description); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(p.ContractAddress); err != nil {
		return sdkerrors.Wrap(err, "contract")
	}
	return nil
}

// MarshalYAML pretty prints the wasm byte code
func (p RemoveContractAuthorizationProposal) MarshalYAML() (interface{}, error) {
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
	if len(title) > legacygovtypes.MaxTitleLength {
		return sdkerrors.Wrapf(govtypes.ErrInvalidProposalContent, "proposal title is longer than max length of %d", legacygovtypes.MaxTitleLength)
	}
	if strings.TrimSpace(description) != description {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "proposal description must not start/end with white spaces")
	}
	if len(description) == 0 {
		return sdkerrors.Wrap(govtypes.ErrInvalidProposalContent, "proposal description cannot be blank")
	}
	if len(description) > legacygovtypes.MaxDescriptionLength {
		return sdkerrors.Wrapf(govtypes.ErrInvalidProposalContent, "proposal description is longer than max length of %d", legacygovtypes.MaxDescriptionLength)
	}
	return nil
}
