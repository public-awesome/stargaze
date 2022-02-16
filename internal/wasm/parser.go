package wasm

import (
	"encoding/json"
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// Encoder describes behavior for provenance smart contract message encoding.
// The contract address must ALWAYS be set as the Msg signer.
type Encoder func(contract sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error)

// MessageEncoders provides stargaze custom encoder for contracts
func MessageEncoders() *wasm.MessageEncoders {
	return &wasm.MessageEncoders{
		Custom: customEncoders(),
	}
}

type MessageWrapper struct {
	Route   string          `json:"route"`
	MsgData json.RawMessage `json:"msg_data"`
	Version string          `json:"version"`
}

type FundCommunityPool struct {
	Amount wasmvmtypes.Coins `json:"amount"`
}

type DistributionMsg struct {
	FundCommunityPool *FundCommunityPool `json:"fund_community_pool,omitempty"`
}

func customEncoders() wasm.CustomEncoder {
	return func(sender sdk.AccAddress, m json.RawMessage) ([]sdk.Msg, error) {
		fmt.Println("sender", sender)
		fmt.Println("msg", string(m))
		msgWrapper := &MessageWrapper{}
		err := json.Unmarshal(m, msgWrapper)
		if err != nil {
			return nil, err
		}
		if msgWrapper.Route == "distribution" {
			fmt.Println("route", "here")
			msg := &DistributionMsg{}
			err = json.Unmarshal(msgWrapper.MsgData, msg)
			if err != nil {
				fmt.Println("route", "err")
				return nil, err
			}
			fmt.Println("unmarshaled", msg, string(msgWrapper.MsgData))

			if msg.FundCommunityPool != nil {
				amount, err := convertWasmCoinsToSdkCoins(msg.FundCommunityPool.Amount)
				if err != nil {
					return nil, err
				}
				m := distributiontypes.MsgFundCommunityPool{
					Amount:    amount,
					Depositor: sender.String(),
				}
				err = m.ValidateBasic()
				if err != nil {
					return nil, err
				}
				msgs := []sdk.Msg{&m}
				fmt.Println("here---\n\n\n\n\n---")
				return msgs, nil
			}
		}
		return nil, nil
	}
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
