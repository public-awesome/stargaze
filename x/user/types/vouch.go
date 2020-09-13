package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewVouch returns a new instance of the Vouch object
func NewVouch(voucher, vouched sdk.AccAddress, comment string) Vouch {
	return Vouch{
		Voucher: voucher,
		Vouched: vouched,
		Comment: comment,
	}
}
