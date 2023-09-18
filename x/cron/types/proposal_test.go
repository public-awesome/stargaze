package types

import (
	"strings"
	"testing"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestValidatePromoteToPrivilegedContractProposal(t *testing.T) {
	specs := map[string]struct {
		src    *PromoteToPrivilegedContractProposal
		expErr bool
	}{
		"all good": {
			src: PromoteProposalFixture(),
		},
		"with empty contract address": {
			src: PromoteProposalFixture(func(p *PromoteToPrivilegedContractProposal) {
				p.Contract = ""
			}),
			expErr: true,
		},
		"with invalid contract address": {
			src: PromoteProposalFixture(func(p *PromoteToPrivilegedContractProposal) {
				p.Contract = "invalid address"
			}),
			expErr: true,
		},

		"base data missing": {
			src: PromoteProposalFixture(func(p *PromoteToPrivilegedContractProposal) {
				p.Title = ""
			}),
			expErr: true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := spec.src.ValidateBasic()
			if spec.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateDemotePrivilegedContractProposal(t *testing.T) {
	specs := map[string]struct {
		src    *DemotePrivilegedContractProposal
		expErr bool
	}{
		"all good": {
			src: DemoteProposalFixture(),
		},
		"with empty contract address": {
			src: DemoteProposalFixture(func(p *DemotePrivilegedContractProposal) {
				p.Contract = ""
			}),
			expErr: true,
		},
		"with invalid contract address": {
			src: DemoteProposalFixture(func(p *DemotePrivilegedContractProposal) {
				p.Contract = "invalid address"
			}),
			expErr: true,
		},

		"base data missing": {
			src: DemoteProposalFixture(func(p *DemotePrivilegedContractProposal) {
				p.Title = ""
			}),
			expErr: true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := spec.src.ValidateBasic()
			if spec.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestProposalYaml(t *testing.T) {
	specs := map[string]struct {
		src govtypes.Content
		exp string
	}{
		"promote proposal": {
			src: PromoteProposalFixture(),
			exp: `title: Foo
description: Bar
contract: cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du
`,
		},
		"demote proposal": {
			src: DemoteProposalFixture(),
			exp: `title: Foo
description: Bar
contract: cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du
`,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			v, err := yaml.Marshal(&spec.src)
			require.NoError(t, err)
			assert.Equal(t, spec.exp, string(v))
		})
	}
}

func TestValidateProposalCommons(t *testing.T) {
	type commonProposal struct {
		Title, Description string
	}

	specs := map[string]struct {
		src    commonProposal
		expErr bool
	}{
		"all good": {src: commonProposal{
			Title:       "Foo",
			Description: "Bar",
		}},
		"prevent empty title": {
			src: commonProposal{
				Description: "Bar",
			},
			expErr: true,
		},
		"prevent white space only title": {
			src: commonProposal{
				Title:       " ",
				Description: "Bar",
			},
			expErr: true,
		},
		"prevent leading white spaces in title": {
			src: commonProposal{
				Title:       " Foo",
				Description: "Bar",
			},
			expErr: true,
		},
		"prevent title exceeds max length ": {
			src: commonProposal{
				Title:       strings.Repeat("a", govtypes.MaxTitleLength+1),
				Description: "Bar",
			},
			expErr: true,
		},
		"prevent empty description": {
			src: commonProposal{
				Title: "Foo",
			},
			expErr: true,
		},
		"prevent leading white spaces in description": {
			src: commonProposal{
				Title:       "Foo",
				Description: " Bar",
			},
			expErr: true,
		},
		"prevent white space only description": {
			src: commonProposal{
				Title:       "Foo",
				Description: " ",
			},
			expErr: true,
		},
		"prevent descr exceeds max length ": {
			src: commonProposal{
				Title:       "Foo",
				Description: strings.Repeat("a", govtypes.MaxDescriptionLength+1),
			},
			expErr: true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := validateProposalCommons(spec.src.Title, spec.src.Description)
			if spec.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
