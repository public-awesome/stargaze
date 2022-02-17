package wasm

import (
	"encoding/json"
	"fmt"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

const (
	DistributionRoute = "distribution"
)

var _ Encoder = CustomDistributionEncoder

type FundCommunityPool struct {
	Amount wasmvmtypes.Coins `json:"amount"`
}

func (fcp FundCommunityPool) Encode(contract sdk.AccAddress) ([]sdk.Msg, error) {
	amount, err := convertWasmCoinsToSdkCoins(fcp.Amount)
	if err != nil {
		return nil, err
	}
	msg := distributiontypes.NewMsgFundCommunityPool(amount, contract)
	return []sdk.Msg{msg}, nil
}

type DistributionMsg struct {
	FundCommunityPool *FundCommunityPool `json:"fund_community_pool,omitempty"`
}

func CustomDistributionEncoder(contract sdk.AccAddress, data json.RawMessage, version string) ([]sdk.Msg, error) {
	msg := &DistributionMsg{}
	err := json.Unmarshal(data, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	if msg.FundCommunityPool != nil {
		return msg.FundCommunityPool.Encode(contract)
	}
	return nil, fmt.Errorf("wasm: invalid custom distribution message")
}

func convertWasmCoinsToSdkCoins(coins []wasmvmtypes.Coin) (sdk.Coins, error) {
	var toSend sdk.Coins
	for _, coin := range coins {
		c, err := convertWasmCoinToSdkCoin(coin)
		if err != nil {
			return nil, err
		}
		toSend = append(toSend, c)
	}
	return toSend, nil
}

func convertWasmCoinToSdkCoin(coin wasmvmtypes.Coin) (sdk.Coin, error) {
	amount, ok := sdk.NewIntFromString(coin.Amount)
	if !ok {
		return sdk.Coin{}, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, coin.Amount+coin.Denom)
	}
	r := sdk.Coin{
		Denom:  coin.Denom,
		Amount: amount,
	}
	return r, r.Validate()
}
