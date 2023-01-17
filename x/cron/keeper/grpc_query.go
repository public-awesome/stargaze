package keeper

import (
	"github.com/public-awesome/stargaze/v8/x/cron/types"
)

var _ types.QueryServer = Keeper{}
