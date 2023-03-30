package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (ca ContractAuthorization) Validate() error {
	_, err := sdk.AccAddressFromBech32(ca.GetContractAddress())
	if err != nil {
		return err
	}

	return validateMethods(ca.Methods)
}
