package types

import (
	"strings"
	"testing"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestValidateSetCodeAuthorizationProposal(t *testing.T) {
	specs := map[string]struct {
		src    *SetCodeAuthorizationProposal
		expErr bool
	}{
		"all good": {
			src: SetCodeAuthorizationProposalFixture(),
		},
		"base data missing": {
			src: SetCodeAuthorizationProposalFixture(func(p *SetCodeAuthorizationProposal) {
				p.Title = ""
			}),
			expErr: true,
		},
		"code auth missing": {
			src: SetCodeAuthorizationProposalFixture(func(p *SetCodeAuthorizationProposal) {
				p.CodeAuthorization = &CodeAuthorization{}
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

func TestValidateRemoveCodeAuthorizationProposal(t *testing.T) {
	specs := map[string]struct {
		src    *RemoveCodeAuthorizationProposal
		expErr bool
	}{
		"all good": {
			src: RemoveCodeAuthorizationProposalFixture(),
		},
		"base data missing": {
			src: RemoveCodeAuthorizationProposalFixture(func(p *RemoveCodeAuthorizationProposal) {
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

func TestValidateSetContractAuthorizationProposal(t *testing.T) {
	specs := map[string]struct {
		src    *SetContractAuthorizationProposal
		expErr bool
	}{
		"all good": {
			src: SetContractAuthorizationProposalFixture(),
		},
		"base data missing": {
			src: SetContractAuthorizationProposalFixture(func(p *SetContractAuthorizationProposal) {
				p.Title = ""
			}),
			expErr: true,
		},
		"contract auth missing": {
			src: SetContractAuthorizationProposalFixture(func(p *SetContractAuthorizationProposal) {
				p.ContractAuthorization = &ContractAuthorization{}
			}),
			expErr: true,
		},
		"contarct addr invalid": {
			src: SetContractAuthorizationProposalFixture(func(p *SetContractAuthorizationProposal) {
				p.ContractAuthorization = &ContractAuthorization{
					ContractAddress: "👻",
				}
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

func TestValidateRemoveContractAuthorizationProposal(t *testing.T) {
	specs := map[string]struct {
		src    *RemoveContractAuthorizationProposal
		expErr bool
	}{
		"all good": {
			src: RemoveContractAuthorizationProposalFixture(),
		},
		"base data missing": {
			src: RemoveContractAuthorizationProposalFixture(func(p *RemoveContractAuthorizationProposal) {
				p.Title = ""
			}),
			expErr: true,
		},
		"contarct addr invalid": {
			src: RemoveContractAuthorizationProposalFixture(func(p *RemoveContractAuthorizationProposal) {
				p.ContractAddress = "👻"
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
		"set code authorization proposal": {
			src: SetCodeAuthorizationProposalFixture(),
			exp: `title: Foo
description: Bar
code_authorization:
  code_id: 1
  methods:
  - mint
  - unlist
`,
		},
		"remove code authorization proposal": {
			src: RemoveCodeAuthorizationProposalFixture(),
			exp: `title: Foo
description: Bar
code_id: 1
`,
		},
		"set contract authorization proposal": {
			src: SetContractAuthorizationProposalFixture(),
			exp: `title: Foo
description: Bar
contract_authorization:
  contract_address: cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du
  methods:
  - mint
  - unlist
`,
		},
		"remove contract authorization proposal": {
			src: RemoveContractAuthorizationProposalFixture(),
			exp: `title: Foo
description: Bar
contract_address: cosmos1qyqszqgpqyqszqgpqyqszqgpqyqszqgpjnp7du
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
