package keeper

import (
	"github.com/public-awesome/stargaze/x/ibc-spend/types"
)

var _ types.QueryServer = Keeper{}
