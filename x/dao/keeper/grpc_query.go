package keeper

import (
	"github.com/public-awesome/stargaze/x/dao/types"
)

var _ types.QueryServer = Keeper{}
