package curating

import (
	"github.com/public-awesome/stakebird/x/curating/keeper"
	"github.com/public-awesome/stakebird/x/curating/types"

	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	// TODO: define constants that you would like exposed from your module

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	QueryParams       = types.QueryParams
	QuerierRoute      = types.QuerierRoute
	RewardPoolName    = types.RewardPoolName
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewDelegation       = stakingTypes.NewDelegation
	// StakeIndexFromKey   = types.StakeIndexFromKey

	// variable aliases
	ModuleCdc                = types.ModuleCdc
	EventTypeCuratingEndTime = types.EventTypeCuratingEndTime
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params
	Post         = types.Post
	Delegation   = stakingTypes.Delegation
)
