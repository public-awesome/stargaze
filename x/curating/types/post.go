package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewPost allocates and returns a new `Post` struct
func NewPost(
	vendorID uint32, postIDBz []byte, bodyHash []byte, creator,
	rewardAccount sdk.AccAddress, curatingEndTime time.Time) Post {

	return Post{
		VendorID:        vendorID,
		PostID:          postIDBz,
		Creator:         creator.String(),
		RewardAccount:   rewardAccount.String(),
		BodyHash:        bodyHash,
		CuratingEndTime: curatingEndTime,
	}
}

// func (cp CuratedPost) String() string {
// 	postID, err := strconv.ParseInt(string(cp.PostID), 10, 64)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return strconv.FormatInt(postID, 10)
// }

// String implements the stringer interface for Post
// func (p *Post) String() string {
// 	// out, err := yaml.Marshal(p)
// 	// if err != nil {
// 	// 	return ""
// 	// }
// 	// return string(out)
// 	return "hello"
// }

// MarshalJSON defines custom encoding scheme
// func (p Post) MarshalJSON() ([]byte, error) {
// 	// if i.i == nil { // Necessary since default Uint initialization has i.i as nil
// 	// 	i.i = new(big.Int)
// 	// }
// 	// return marshalJSON(i.i)
// 	// return []byte("hello"), nil

// 	out, err := json.Marshal(p)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return out, nil
// }

// MarshalJSON returns the JSON representation of a ModuleAccount.
// func (ma ModuleAccount) MarshalJSON() ([]byte, error) {
// 	accAddr, err := sdk.AccAddressFromBech32(ma.Address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return json.Marshal(moduleAccountPretty{
// 		Address:       accAddr,
// 		PubKey:        "",
// 		AccountNumber: ma.AccountNumber,
// 		Sequence:      ma.Sequence,
// 		Name:          ma.Name,
// 		Permissions:   ma.Permissions,
// 	})
// }

// UnmarshalJSON defines custom decoding scheme
// func (p *Post) UnmarshalJSON(bz []byte) error {
// 	// if i.i == nil { // Necessary since default Int initialization has i.i as nil
// 	// 	i.i = new(big.Int)
// 	// }
// 	// return unmarshalJSON(i.i, bz)
// 	return nil
// }

// MarshalYAML returns the YAML representation.
// func (p *Post) MarshalYAML() (interface{}, error) {
// 	// return i.String(), nil
// 	return "hello", nil
// }
