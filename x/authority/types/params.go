package types

import (
	"gopkg.in/yaml.v2"
)

// default module parameters
func DefaultParams() Params {
	return Params{
		Authorizations: []*Authorization{
			{
				MsgTypeUrl: "/cosmos.distribution.v1beta1.MsgCommunityPoolSpend",
				Addresses:  []string{"stars1x394ype3x8nt9wz0j78m8c8kcezpslrcjmmrc0"},
			},
			{
				MsgTypeUrl: "/cosmwasm.wasm.v1.MsgUpdateParams",
				Addresses:  []string{"stars1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cyt4fdd"},
			},
		},
	}
}

func NewParams(authorizations []*Authorization) Params {
	return Params{
		Authorizations: authorizations,
	}
}

func (p Params) Validate() error {
	for _, auth := range p.GetAuthorizations() {
		if err := auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
