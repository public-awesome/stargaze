package curating

import (
	"github.com/public-awesome/stakebird/x/curating/keeper"
	"github.com/public-awesome/stakebird/x/curating/types"
)

// constants exposed from module
const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	QueryParams       = types.QueryParams
	QuerierRoute      = types.QuerierRoute
	RewardPoolName    = types.RewardPoolName
)

// functions aliases
var (
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc                = types.ModuleCdc
	EventTypePost            = types.EventTypePost
	EventTypeUpvote          = types.EventTypeUpvote
	EventTypeCuratingEndTime = types.EventTypeCuratingEndTime
)

// type aliases
type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params
	Post         = types.Post
	Upvote       = types.Upvote
)
