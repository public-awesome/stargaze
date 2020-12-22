package keeper

import (
	"context"

	"github.com/public-awesome/stakebird/x/stake/types"
)

var _ types.QueryServer = Keeper{}

// Stake returns a Stake based on vendor and Stake id
func (k Keeper) Stake(c context.Context, req *types.QueryStakeRequest) (*types.QueryStakeResponse, error) {
	// ctx := sdk.UnwrapSDKContext(c)
	// Stake, found, err := k.GetStake(ctx, req.VendorId, req.StakeId)
	// if err != nil {
	// 	return nil, err
	// }
	// if !found {
	// 	return nil, fmt.Errorf("Stake does not exist")
	// }
	// return &types.QueryStakeResponse{
	// 	Stake: &Stake,
	// }, nil
	return nil, nil
}
