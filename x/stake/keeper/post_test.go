package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPost(t *testing.T) {
	_, app, ctx := createTestInput()

	postID := uint64(500)
	vendorID := uint64(100)
	votingPeriod := time.Hour * 24 * 7
	app.StakeKeeper.CreatePost(ctx, postID, vendorID, "body string", votingPeriod)

	_, found := app.StakeKeeper.GetPost(ctx, vendorID, postID)
	require.True(t, found, "Post should be found")
}
