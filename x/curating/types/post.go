package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// // Marshal implements the gogo proto custom type interface.
// func (i Int) Marshal() ([]byte, error) {
// 	if i.i == nil {
// 		i.i = new(big.Int)
// 	}
// 	return i.i.MarshalText()
// }

// // MarshalTo implements the gogo proto custom type interface.
// func (i *Int) MarshalTo(data []byte) (n int, err error) {
// 	if i.i == nil {
// 		i.i = new(big.Int)
// 	}
// 	if len(i.i.Bytes()) == 0 {
// 		copy(data, []byte{0x30})
// 		return 1, nil
// 	}

// 	bz, err := i.Marshal()
// 	if err != nil {
// 		return 0, err
// 	}

// 	copy(data, bz)
// 	return len(bz), nil
// }

// // Unmarshal implements the gogo proto custom type interface.
// func (i *Int) Unmarshal(data []byte) error {
// 	if len(data) == 0 {
// 		i = nil
// 		return nil
// 	}

// 	if i.i == nil {
// 		i.i = new(big.Int)
// 	}

// 	if err := i.i.UnmarshalText(data); err != nil {
// 		return err
// 	}

// 	if i.i.BitLen() > maxBitLen {
// 		return fmt.Errorf("integer out of range; got: %d, max: %d", i.i.BitLen(), maxBitLen)
// 	}

// 	return nil
// }

// // Size implements the gogo proto custom type interface.
// func (i *Int) Size() int {
// 	bz, _ := i.Marshal()
// 	return len(bz)
// }

// type PostID []byte

// PostID is a byte array
type PostID struct {
	id []byte
}

// Marshal implements the gogo proto custom type interface.
func (p PostID) Marshal() ([]byte, error) {
	if p.id == nil {
		p.id = []byte{}
	}

	return p.id, nil
}

// Size implements the gogo proto custom type interface.
func (p *PostID) Size() int {
	bz, _ := p.Marshal()
	return len(bz)
}

// MarshalJSON encodes a post id snowflake as JSON bytes
// func (p PostID) MarshalJSON() ([]byte, error) {
// 	var temp [8]byte
// 	copy(temp[:], p)
// 	id := snowflake.ParseIntBytes(temp)

// 	out, err := json.Marshal(id.String())
// 	if err != nil {
// 		return out, err
// 	}

// 	return out, nil
// }

// Posts is a collection of Post objects
type Posts []Post

// CuratingQueue is a collection of VPPairs objects
type CuratingQueue []VPPair

// NewPost allocates and returns a new `Post` struct
func NewPost(
	vendorID uint32, postIDBz PostID, bodyHash []byte, creator,
	rewardAccount sdk.AccAddress, curatingEndTime time.Time) Post {

	return Post{
		VendorID:        vendorID,
		PostID:          &postIDBz,
		Creator:         creator.String(),
		RewardAccount:   rewardAccount.String(),
		BodyHash:        bodyHash,
		CuratingEndTime: curatingEndTime,
	}
}

// postPretty is a representation of `Post` suitable for silly hoomans
// type postPretty struct {
// 	VendorID        uint32 `json:"vendor_id" yaml:"vendor_id"`
// 	PostID          string `json:"post_id" yaml:"post_id"`
// 	Creator         string `json:"creator" yaml:"creator"`
// 	RewardAccount   string `json:"reward_account" yaml:"reward_account"`
// 	BodyHash        string `json:"body_hash" yaml:"body_hash"`
// 	CuratingEndTime string `json:"curating_end_time" yaml:"curating_end_time"`
// }

// // MarshalJSON defines custom encoding scheme
// func (p Post) MarshalJSON() ([]byte, error) {
// 	var temp [8]byte
// 	copy(temp[:], p.PostID) // convert a postID byte slice into a fixed 8 byte array
// 	postID := snowflake.ParseIntBytes(temp)

// 	out, err := json.Marshal(postPretty{
// 		VendorID:        p.VendorID,
// 		PostID:          postID.String(),
// 		Creator:         p.Creator,
// 		RewardAccount:   p.RewardAccount,
// 		BodyHash:        hex.EncodeToString(p.BodyHash),
// 		CuratingEndTime: p.CuratingEndTime.Format(time.RFC3339),
// 	})
// 	if err != nil {
// 		return out, err
// 	}

// 	return out, nil
// }

// func (vo *VoteOption) UnmarshalJSON(data []byte) error {
// 	var s string
// 	err := json.Unmarshal(data, &s)
// 	if err != nil {
// 		return err
// 	}

// 	bz2, err := VoteOptionFromString(s)
// 	if err != nil {
// 		return err
// 	}

// 	*vo = bz2
// 	return nil
// }

// UnmarshalJSON decodes JSON bytes into a Post
// func (p *Post) UnmarshalJSON(data []byte) error {
// 	// TODO: do a custom unmarshaller for PostID only?

// 	return nil
// }

// PostIDStr returns a string representation of the underlying bytes that conforms an id.
// func (p Post) PostIDStr() string {
// 	var temp [8]byte
// 	copy(temp[:], p.PostID) // convert a postID byte slice into a fixed 8 byte array
// 	postID := snowflake.ParseIntBytes(temp)
// 	return postID.String()
// }

// type vpPairPretty struct {
// 	VendorID uint32 `json:"vendor_id" yaml:"vendor_id"`
// 	PostID   string `json:"post_id" yaml:"post_id"`
// }

// // MarshalJSON defines custom encoding scheme
// func (vp VPPair) MarshalJSON() ([]byte, error) {
// 	var temp [8]byte
// 	copy(temp[:], vp.PostID) // convert a postID byte slice into a fixed 8 byte array
// 	postID := snowflake.ParseIntBytes(temp)

// 	out, err := json.Marshal(vpPairPretty{
// 		VendorID: vp.VendorID,
// 		PostID:   postID.String(),
// 	})
// 	if err != nil {
// 		return out, err
// 	}

// 	return out, nil
// }
