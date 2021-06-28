package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	// ProposalTypeCommunityPoolIBCSpend defines the type for a CommunityPoolSpendProposal
	ProposalTypeCommunityPoolIBCSpend = "CommunityPoolIBCSpend"
)

// Assert CommunityPoolSpendProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &CommunityPoolIBCSpendProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeCommunityPoolIBCSpend)
	govtypes.RegisterProposalTypeCodec(&CommunityPoolIBCSpendProposal{}, "stargaze/CommunityPoolIBCSpendProposal")
}

// NewCommunityPoolIBCSpendProposal creates a new community pool spned proposal.
//nolint:interfacer
func NewCommunityPoolIBCSpendProposal(
	title, description string,
	recipient string,
	amount sdk.Coins,
	sourceChannel string,
	timeout uint64,
) *CommunityPoolIBCSpendProposal {
	return &CommunityPoolIBCSpendProposal{title, description, recipient, amount, sourceChannel, timeout}
}

// GetTitle returns the title of a community pool spend proposal.
func (csp *CommunityPoolIBCSpendProposal) GetTitle() string { return csp.Title }

// GetDescription returns the description of a community pool spend proposal.
func (csp *CommunityPoolIBCSpendProposal) GetDescription() string { return csp.Description }

// GetDescription returns the routing key of a community pool spend proposal.
func (csp *CommunityPoolIBCSpendProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a community pool spend proposal.
func (csp *CommunityPoolIBCSpendProposal) ProposalType() string {
	return ProposalTypeCommunityPoolIBCSpend
}

// ValidateBasic runs basic stateless validity checks
func (csp *CommunityPoolIBCSpendProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(csp)
	if err != nil {
		return err
	}
	if !csp.Amount.IsValid() {
		return ErrInvalidProposalAmount
	}
	if csp.Recipient == "" {
		return ErrEmptyProposalRecipient
	}

	return nil
}

// String implements the Stringer interface.
func (csp CommunityPoolIBCSpendProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Community Pool IBC Spend Proposal:
  Title:       %s
  Description: %s
  Recipient:   %s
  Amount:      %s
  Channel:     %s
  Timeout:     %d
`, csp.Title, csp.Description, csp.Recipient, csp.Amount, csp.SourceChannel, csp.Timeout))
	return b.String()
}
