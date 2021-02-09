package types

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/bwmarrin/snowflake"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CustomProtobufType defines the interface custom gogo proto types must implement
// in order to be used as a "customtype" extension.
//
// ref: https://github.com/gogo/protobuf/blob/master/custom_types.md
type CustomProtobufType interface {
	Marshal() ([]byte, error)
	MarshalTo(data []byte) (n int, err error)
	Unmarshal(data []byte) error
	Size() int

	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

var _ CustomProtobufType = (*PostID)(nil)
var _ CustomProtobufType = (*BodyHash)(nil)

// PostID represents a Twitter post ID (for now)
type PostID struct {
	id snowflake.ID
}

// PostIDFromString does exactly whats on the label
func PostIDFromString(id string) (PostID, error) {
	postID, err := snowflake.ParseString(id)
	if err != nil {
		return PostID{}, err
	}
	return PostID{id: postID}, nil
}

// String like the array of chars, not the theory
func (p PostID) String() string {
	return p.id.String()
}

// Bytes returns bytes in big endian
func (p PostID) Bytes() []byte {
	temp := p.id.IntBytes()
	return temp[:]
}

// Marshal implements the gogo proto custom type interface
func (p PostID) Marshal() ([]byte, error) {
	return p.id.Bytes(), nil
}

// MarshalTo implements the gogo proto custom type interface
func (p PostID) MarshalTo(data []byte) (n int, err error) {
	bz, err := p.Marshal()
	if err != nil {
		return 0, err
	}

	copy(data, bz)
	return len(bz), nil
}

// Unmarshal implements the gogo proto custom type interface
func (p *PostID) Unmarshal(data []byte) error {
	id, err := snowflake.ParseBytes(data)
	if err != nil {
		return err
	}
	p.id = id

	return nil
}

// Size implements the gogo proto custom type interface
func (p PostID) Size() int {
	bz, err := p.Marshal()
	if err != nil {
		return 0
	}
	return len(bz)
}

// MarshalJSON implements the gogo proto custom type interface
func (p PostID) MarshalJSON() ([]byte, error) {
	return p.id.MarshalJSON()
}

// UnmarshalJSON implements the gogo proto custom type interface
func (p *PostID) UnmarshalJSON(data []byte) error {
	err := p.id.UnmarshalJSON(data)
	return err
}

// Equal compares post id is the same
func (p PostID) Equal(p2 PostID) bool {
	return p.id == p2.id
}

// Posts is a collection of Post objects
type Posts []Post

// CuratingQueue is a collection of VPPairs objects
type CuratingQueue []VPPair

// NewPost allocates and returns a new `Post` struct
func NewPost(
	vendorID uint32, postID PostID, bodyHash BodyHash, creator,
	rewardAccount sdk.AccAddress, curatingEndTime time.Time) Post {

	return Post{
		VendorID:        vendorID,
		PostID:          postID,
		Creator:         creator.String(),
		RewardAccount:   rewardAccount.String(),
		BodyHash:        bodyHash,
		CuratingEndTime: curatingEndTime,
		TotalVotes:      0,
		TotalAmount:     sdk.Coin{},
	}
}

type BodyHash struct {
	data []byte
}

// BodyHashFromString does exactly whats on the label
func BodyHashFromString(body string) (BodyHash, error) {
	h := sha256.New()
	_, err := h.Write([]byte(body))
	if err != nil {
		return BodyHash{}, err
	}
	digest := h.Sum(nil)
	return BodyHash{digest[:20]}, nil
}

// String returns the hex string of the body hash
func (b *BodyHash) String() string {
	return hex.EncodeToString(b.data)
}

// Marshal implements the gogo proto custom type interface
func (b BodyHash) Marshal() ([]byte, error) {
	return b.data, nil
}

// MarshalJSON implements the gogo proto custom type interface
func (b BodyHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.data)
}

// MarshalTo implements the gogo proto custom type interface
func (b BodyHash) MarshalTo(data []byte) (n int, err error) {
	bz, err := b.Marshal()
	if err != nil {
		return 0, err
	}

	copy(data, bz)
	return len(bz), nil
}

// Size implements the gogo proto custom type interface
func (b BodyHash) Size() int {
	bz, err := b.Marshal()
	if err != nil {
		return 0
	}
	return len(bz)
}

// Unmarshal implements the gogo proto custom type interface
func (b *BodyHash) Unmarshal(data []byte) error {
	b.data = data
	return nil
}

// UnmarshalJSON implements the gogo proto custom type interface
func (b *BodyHash) UnmarshalJSON(data []byte) error {
	var d []byte
	err := json.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	b.data = d

	return nil
}
