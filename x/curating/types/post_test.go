package types_test

import (
	"testing"

	"github.com/public-awesome/stakebird/x/curating/types"

	"github.com/stretchr/testify/require"
)

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
