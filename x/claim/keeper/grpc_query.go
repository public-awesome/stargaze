package keeper

import (
	"github.com/public-awesome/stargaze/x/claim/types"
)

var _ types.QueryServer = Keeper{}
