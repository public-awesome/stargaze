package types_test

import (
	"crypto/sha256"
	"testing"

	"github.com/public-awesome/stakebird/x/curating/types"

	"github.com/stretchr/testify/require"
)

// func TestPostString(t *testing.T) {
// 	app := simapp.Setup(false)
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

// 	vendorID := uint32(1)
// 	postID, err := types.PostIDFromString("1000")
// 	require.NoError(t, err)
// 	addresses := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(1000000))
// 	bodyHash, err := hash("lorem ipsum")
// 	require.NoError(t, err)
// 	curatingEndTime := time.Now()

// 	post := types.NewPost(vendorID, postID, bodyHash, addresses[0], addresses[1], curatingEndTime)
// 	_, err = post.MarshalJSON()
// 	require.NoError(t, err)
// }

func TestPostID(t *testing.T) {
	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)

	require.Equal(t, "500", postID.String())
	bz1 := postID.Bytes()
	require.Len(t, bz1, 8)
	bz2, err := postID.Marshal()
	require.NoError(t, err)
	require.NotEqual(t, bz1, bz2)

	err = postID.Unmarshal(bz2)
	require.NoError(t, err)
}

func hash(body string) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write([]byte(body))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
