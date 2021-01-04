package types

import (
	"encoding/json"

	"github.com/bwmarrin/snowflake"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewStake allocates and returns a new `Stake` struct
func NewStake(
	vendorID uint32, postID []byte,
	delegator sdk.AccAddress, validator sdk.ValAddress, amount sdk.Int) Stake {

	return Stake{
		VendorID:  vendorID,
		PostID:    postID,
		Delegator: delegator.String(),
		Validator: validator.String(),
		Amount:    amount,
	}
}

// Stakes is a collection of Stake 🥩
type Stakes []Stake

type stakePretty struct {
	VendorID  uint32  `json:"vendor_id" yaml:"vendor_id"`
	PostID    string  `json:"post_id" yaml:"post_id"`
	Delegator string  `json:"delegator" yaml:"delegator"`
	Validator string  `json:"validator" yaml:"validator"`
	Amount    sdk.Int `json:"amount" yaml:"amount"`
}

// MarshalJSON defines custom encoding scheme
func (s Stake) MarshalJSON() ([]byte, error) {
	var temp [8]byte
	copy(temp[:], s.PostID) // convert a postID byte slice into a fixed 8 byte array
	postID := snowflake.ParseIntBytes(temp)

	out, err := json.Marshal(stakePretty{
		VendorID:  s.VendorID,
		PostID:    postID.String(),
		Delegator: s.Delegator,
		Validator: s.Validator,
		Amount:    s.Amount,
	})
	if err != nil {
		return out, err
	}

	return out, nil
}
