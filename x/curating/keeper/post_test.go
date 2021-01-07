package keeper_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/simapp"
	"github.com/public-awesome/stakebird/x/curating/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestCreatePost(t *testing.T) {
	fakedenom := "fakedenom"
	app := simapp.SetupWithStakeDenom(false, "fakedenom")
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	postID, err := types.PostIDFromString("500")
	require.NoError(t, err)
	vendorID := uint32(1)
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(27_000_000), fakedenom)

	ctx = ctx.WithBlockTime(time.Now())

	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string", addrs[0], addrs[0])
	require.NoError(t, err)

	post, found, err := app.CuratingKeeper.GetPost(ctx, vendorID, postID)
	require.NoError(t, err)
	require.True(t, found, "post should be found")
	require.Equal(t, "500", post.PostID.String())
	obj, err := json.MarshalIndent(post.PostID, "", " ")
	if err != nil {
		panic(err)
	}
	require.Equal(t, "\"500\"", string(obj))

	// obj, err = json.MarshalIndent(post, "", " ")
	// if err != nil {
	// 	panic(err)
	// }
	// require.Equal(t, "\"500\"", string(obj))
	// fmt.Println(string(obj))

	curatingQueue := app.CuratingKeeper.GetCurationQueueTimeSlice(ctx, ctx.BlockTime().Add(10*time.Minute))
	fmt.Println(curatingQueue)
	require.Equal(t, curatingQueue[0].PostID.String(), "500")
	require.Equal(t, curatingQueue[0].VendorID, uint32(1))

	// add another post
	postID, err = types.PostIDFromString("501")
	err = app.CuratingKeeper.CreatePost(ctx, vendorID, postID, "body string1", addrs[0], addrs[0])
	require.NoError(t, err)

	// fast forward block time
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Hour*24*3 + 1))

	// test expired post iterator
	app.CuratingKeeper.IterateExpiredPosts(ctx, func(post types.Post) (stop bool) {
		fmt.Println("post ", post)
		// require.Equal(t, post.PostID.String(), "500")
		return false
	})
}
