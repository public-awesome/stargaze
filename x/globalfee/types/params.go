package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gopkg.in/yaml.v2"
)

func (p Params) Validate() error {
	for _, addr := range p.GetPrivilegedAddress() {
		_, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
