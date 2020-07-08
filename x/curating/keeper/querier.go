package keeper

import (
	"fmt"
	"strconv"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// [TODO]
// https://github.com/public-awesome/stakebird/issues/57
// https://github.com/public-awesome/stakebird/issues/58

// NewQuerier creates a new querier for curating clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryParams:
			return queryParams(ctx, k)
		case types.QueryPost:
			return queryPost(ctx, path[1:], req, k)
		// case types.QueryPosts:
		// 	return queryPosts(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown curating query endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryPost(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	i64, _ := strconv.ParseUint(path[0], 10, 32)
	vendorID := uint32(i64)
	postID := path[1]

	post, _, err := k.GetPost(ctx, vendorID, postID)
	if err != nil {
		fmt.Println(err)
	}

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, post)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// func queryPosts(ctx sdk.Context, k Keeper) ([]byte, error) {
// 	posts := k.GetPosts(ctx)

// 	res, err := codec.MarshalJSONIndent(types.ModuleCdc, posts)
// 	if err != nil {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
// 	}

// 	return res, nil
// }
