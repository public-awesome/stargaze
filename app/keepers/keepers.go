package keepers

import (
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
)

type StargazeKeepers struct {
	IBCKeeper *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
}
