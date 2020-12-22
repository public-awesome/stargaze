package stake_test

// func TestHandleMsgStake(t *testing.T) {
// 	app := simapp.Setup(false)
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

// 	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.NewInt(10_000_000))

// 	handler := stake.NewHandler(app.StakeKeeper)
// 	msg := types.NewMsgStake(1, "123", addrs[0], sdk.NewInt(1_000_000))
// 	_, err := handler(ctx, msg)
// 	assert.NoError(t, err)
// }

// func TestHandleMsgUnstake(t *testing.T) {
// 	app := simapp.Setup(false)
// 	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

// 	addrs = simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.NewInt(10_000_000))

// 	handler := curating.NewHandler(app.CuratingKeeper)
// 	msgUpvote := types.NewMsgUpvote(1, "123", addrs[0], nil, 1)
// 	_, err := handler(ctx, msgUpvote)
// 	assert.NoError(t, err)
// }
