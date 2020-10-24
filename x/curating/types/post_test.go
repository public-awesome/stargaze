package types_test

import (
	"crypto/sha256"
	"testing"
	time "time"

	"github.com/bwmarrin/snowflake"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/simapp"
	"github.com/public-awesome/stakebird/x/curating/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestPostString(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	postIDBz, err := postIDBytes("1000")
	require.NoError(t, err)
	addresses := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	bodyHash, err := hash("lorem ipsum")
	require.NoError(t, err)
	curatingEndTime := time.Now()

	post := types.NewPost(vendorID, postIDBz, bodyHash, addresses[0], addresses[1], curatingEndTime)
	_, err = post.MarshalJSON()
	require.NoError(t, err)
}

// postIDBytes returns the byte representation of a postID int64
func postIDBytes(postID string) ([]byte, error) {
	pID, err := snowflake.ParseString(postID)
	if err != nil {
		return nil, err
	}

	temp := pID.IntBytes()

	return temp[:], nil
}

func hash(body string) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write([]byte(body))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
