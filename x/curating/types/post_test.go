package types_test

import (
	"crypto/sha256"
	"encoding/binary"
	"strconv"
	"testing"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/public-awesome/stakebird/simapp"
	"github.com/public-awesome/stakebird/x/curating/types"
)

func TestPostString(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	vendorID := uint32(1)
	postIDBz, err := postIDBytes("100")
	require.NoError(t, err)
	addresses := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
	bodyHash, err := hash("lorem ipsum")
	require.NoError(t, err)
	curatingEndTime := time.Now()

	post := types.NewPost(vendorID, postIDBz, bodyHash, addresses[0], addresses[1], curatingEndTime)
	require.Equal(t, post.String(), "hello")
}

// postIDBytes returns the byte representation of a postID
func postIDBytes(postID string) ([]byte, error) {
	postIDInt64, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		return nil, err
	}

	postIDBz := make([]byte, 8)
	binary.BigEndian.PutUint64(postIDBz, uint64(postIDInt64))

	return postIDBz, nil
}

func hash(body string) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write([]byte(body))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
