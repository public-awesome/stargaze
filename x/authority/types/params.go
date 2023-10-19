package types

import (
	"gopkg.in/yaml.v2"
)

// default module parameters
func DefaultParams() Params {
	return Params{
		Authorizations: []*Authorization{},
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
