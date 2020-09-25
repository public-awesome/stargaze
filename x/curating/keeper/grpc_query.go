package keeper

import (
	"context"

	"github.com/public-awesome/stakebird/x/curating/types"
)

var _ types.QueryServer = Keeper{}

// Post returns a post based on vendor and post id
func (k Keeper) Post(context.Context, *types.QueryPostRequest) (*types.QueryPostResponse, error) {
	// TODO(jhernandezb): complete query
	return nil, nil
}
