package types

import (
	fmt "fmt"
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

// PostID represents a Twitter post ID (for now)
type PostID struct {
	id snowflake.ID
}

func PostIDFromString(id string) (PostID, error) {
	postID, err := snowflake.ParseString(id)
	if err != nil {
		return PostID{}, err
	}
	return PostID{id: postID}, nil
}

func (p PostID) String() string {
	return p.id.String()
}

func (p PostID) Bytes() []byte {
	temp := p.id.IntBytes()
	return temp[:]
}

// Marshal implements the gogo proto custom type interface.
func (p PostID) Marshal() ([]byte, error) {
	return p.id.Bytes(), nil
}

func (p PostID) MarshalTo(data []byte) (n int, err error) {
	return 8, nil
}

func (p PostID) Unmarshal(data []byte) error {
	return nil
}

// Size implements the gogo proto custom type interface.
func (p *PostID) Size() int {
	bz, err := p.Marshal()
	if err != nil {
		fmt.Println(err)
	}
	return len(bz)
}

func (p PostID) MarshalJSON() ([]byte, error) {
	return p.id.MarshalJSON()
}

func (p PostID) UnmarshalJSON(data []byte) error {
	return p.id.UnmarshalJSON(data)
}

func (p PostID) Equal(p2 PostID) bool {
	return p.id == p2.id
}

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
		PostID:          postIDBz,
		Creator:         creator.String(),
		RewardAccount:   rewardAccount.String(),
		BodyHash:        bodyHash,
		CuratingEndTime: curatingEndTime,
	}
}
