package types_test

import (
	fmt "fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stargaze/v17/x/tokenfactory/types"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	"github.com/cometbft/cometbft/crypto/ed25519"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// TestMsgCreateDenom tests if valid/invalid create denom messages are properly validated/invalidated
func TestMsgCreateDenom(t *testing.T) {
	// generate a private/public key pair and get the respective address
	pk1 := ed25519.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pk1.Address())

	// make a proper createDenom message
	createMsg := func(after func(msg types.MsgCreateDenom) types.MsgCreateDenom) types.MsgCreateDenom {
		properMsg := *types.NewMsgCreateDenom(
			addr1.String(),
			"bitcoin",
		)

		return after(properMsg)
	}

	// validate createDenom message was created as intended
	msg := createMsg(func(msg types.MsgCreateDenom) types.MsgCreateDenom {
		return msg
	})
	require.Equal(t, msg.Route(), types.RouterKey)
	require.Equal(t, msg.Type(), "create_denom")
	signers := msg.GetSigners()
	require.Equal(t, len(signers), 1)
	require.Equal(t, signers[0].String(), addr1.String())

	tests := []struct {
		name       string
		msg        types.MsgCreateDenom
		expectPass bool
	}{
		{
			name: "proper msg",
			msg: createMsg(func(msg types.MsgCreateDenom) types.MsgCreateDenom {
				return msg
			}),
			expectPass: true,
		},
		{
			name: "empty sender",
			msg: createMsg(func(msg types.MsgCreateDenom) types.MsgCreateDenom {
				msg.Sender = ""
				return msg
			}),
			expectPass: false,
		},
		{
			name: "invalid subdenom",
			msg: createMsg(func(msg types.MsgCreateDenom) types.MsgCreateDenom {
				msg.Subdenom = "thissubdenomismuchtoolongasdkfjaasdfdsafsdlkfnmlksadmflksmdlfmlsakmfdsafasdfasdf"
				return msg
			}),
			expectPass: false,
		},
	}

	for _, test := range tests {
		if test.expectPass {
			require.NoError(t, test.msg.ValidateBasic(), "test: %v", test.name)
		} else {
			require.Error(t, test.msg.ValidateBasic(), "test: %v", test.name)
		}
	}
}

// TestMsgMint tests if valid/invalid create denom messages are properly validated/invalidated
func TestMsgMint(t *testing.T) {
	// generate a private/public key pair and get the respective address
	pk1 := ed25519.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pk1.Address())

	// make a proper mint message
	createMsg := func(after func(msg types.MsgMint) types.MsgMint) types.MsgMint {
		properMsg := *types.NewMsgMint(
			addr1.String(),
			sdk.NewCoin("bitcoin", math.NewInt(500000000)),
		)

		return after(properMsg)
	}

	// validate mint message was created as intended
	msg := createMsg(func(msg types.MsgMint) types.MsgMint {
		return msg
	})
	require.Equal(t, msg.Route(), types.RouterKey)
	require.Equal(t, msg.Type(), "tf_mint")
	signers := msg.GetSigners()
	require.Equal(t, len(signers), 1)
	require.Equal(t, signers[0].String(), addr1.String())

	tests := []struct {
		name       string
		msg        types.MsgMint
		expectPass bool
	}{
		{
			name: "proper msg",
			msg: createMsg(func(msg types.MsgMint) types.MsgMint {
				return msg
			}),
			expectPass: true,
		},
		{
			name: "empty sender",
			msg: createMsg(func(msg types.MsgMint) types.MsgMint {
				msg.Sender = ""
				return msg
			}),
			expectPass: false,
		},
		{
			name: "zero amount",
			msg: createMsg(func(msg types.MsgMint) types.MsgMint {
				msg.Amount = sdk.NewCoin("bitcoin", math.ZeroInt())
				return msg
			}),
			expectPass: false,
		},
		{
			name: "negative amount",
			msg: createMsg(func(msg types.MsgMint) types.MsgMint {
				msg.Amount.Amount = math.NewInt(-10000000)
				return msg
			}),
			expectPass: false,
		},
		{
			name: "proper msg with mint_to",
			msg: createMsg(func(msg types.MsgMint) types.MsgMint {
				msg.MintToAddress = msg.Sender
				return msg
			}),
			expectPass: true,
		},
		{
			name: "invalid mint to address",
			msg: createMsg(func(msg types.MsgMint) types.MsgMint {
				msg.MintToAddress = "someaddress"
				return msg
			}),
			expectPass: false,
		},
	}

	for _, test := range tests {
		if test.expectPass {
			require.NoError(t, test.msg.ValidateBasic(), "test: %v", test.name)
		} else {
			require.Error(t, test.msg.ValidateBasic(), "test: %v", test.name)
		}
	}
}

// TestMsgBurn tests if valid/invalid create denom messages are properly validated/invalidated
func TestMsgBurn(t *testing.T) {
	// generate a private/public key pair and get the respective address
	pk1 := ed25519.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pk1.Address())

	// make a proper burn message
	baseMsg := types.NewMsgBurn(
		addr1.String(),
		sdk.NewCoin("bitcoin", math.NewInt(500000000)),
	)

	// validate burn message was created as intended
	require.Equal(t, baseMsg.Route(), types.RouterKey)
	require.Equal(t, baseMsg.Type(), "tf_burn")
	signers := baseMsg.GetSigners()
	require.Equal(t, len(signers), 1)
	require.Equal(t, signers[0].String(), addr1.String())

	tests := []struct {
		name       string
		msg        func() *types.MsgBurn
		expectPass bool
	}{
		{
			name: "proper msg",
			msg: func() *types.MsgBurn {
				msg := baseMsg
				return msg
			},
			expectPass: true,
		},
		{
			name: "empty sender",
			msg: func() *types.MsgBurn {
				msg := baseMsg
				msg.Sender = ""
				return msg
			},
			expectPass: false,
		},
		{
			name: "zero amount",
			msg: func() *types.MsgBurn {
				msg := baseMsg
				msg.Amount.Amount = math.ZeroInt()
				return msg
			},
			expectPass: false,
		},
		{
			name: "negative amount",
			msg: func() *types.MsgBurn {
				msg := baseMsg
				msg.Amount.Amount = math.NewInt(-10000000)
				return msg
			},
			expectPass: false,
		},
	}

	for _, test := range tests {
		if test.expectPass {
			require.NoError(t, test.msg().ValidateBasic(), "test: %v", test.name)
		} else {
			require.Error(t, test.msg().ValidateBasic(), "test: %v", test.name)
		}
	}
}

// TestMsgChangeAdmin tests if valid/invalid create denom messages are properly validated/invalidated
func TestMsgChangeAdmin(t *testing.T) {
	// generate a private/public key pair and get the respective address
	pk1 := ed25519.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pk1.Address())
	pk2 := ed25519.GenPrivKey().PubKey()
	addr2 := sdk.AccAddress(pk2.Address())
	tokenFactoryDenom := fmt.Sprintf("factory/%s/bitcoin", addr1.String())

	// make a proper changeAdmin message
	baseMsg := types.NewMsgChangeAdmin(
		addr1.String(),
		tokenFactoryDenom,
		addr2.String(),
	)

	// validate changeAdmin message was created as intended
	require.Equal(t, baseMsg.Route(), types.RouterKey)
	require.Equal(t, baseMsg.Type(), "change_admin")
	signers := baseMsg.GetSigners()
	require.Equal(t, len(signers), 1)
	require.Equal(t, signers[0].String(), addr1.String())

	tests := []struct {
		name       string
		msg        func() *types.MsgChangeAdmin
		expectPass bool
	}{
		{
			name: "proper msg",
			msg: func() *types.MsgChangeAdmin {
				msg := baseMsg
				return msg
			},
			expectPass: true,
		},
		{
			name: "empty sender",
			msg: func() *types.MsgChangeAdmin {
				msg := baseMsg
				msg.Sender = ""
				return msg
			},
			expectPass: false,
		},
		{
			name: "empty newAdmin",
			msg: func() *types.MsgChangeAdmin {
				msg := baseMsg
				msg.NewAdmin = ""
				return msg
			},
			expectPass: false,
		},
		{
			name: "invalid denom",
			msg: func() *types.MsgChangeAdmin {
				msg := baseMsg
				msg.Denom = "bitcoin"
				return msg
			},
			expectPass: false,
		},
	}

	for _, test := range tests {
		if test.expectPass {
			require.NoError(t, test.msg().ValidateBasic(), "test: %v", test.name)
		} else {
			require.Error(t, test.msg().ValidateBasic(), "test: %v", test.name)
		}
	}
}

// TestMsgSetDenomMetadata tests if valid/invalid create denom messages are properly validated/invalidated
func TestMsgSetDenomMetadata(t *testing.T) {
	// generate a private/public key pair and get the respective address
	pk1 := ed25519.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pk1.Address())
	tokenFactoryDenom := fmt.Sprintf("factory/%s/bitcoin", addr1.String())
	denomMetadata := banktypes.Metadata{
		Description: "nakamoto",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    tokenFactoryDenom,
				Exponent: 0,
			},
			{
				Denom:    "sats",
				Exponent: 6,
			},
		},
		Display: "sats",
		Base:    tokenFactoryDenom,
		Name:    "bitcoin",
		Symbol:  "BTC",
	}
	invalidDenomMetadata := banktypes.Metadata{
		Description: "nakamoto",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "bitcoin",
				Exponent: 0,
			},
			{
				Denom:    "sats",
				Exponent: 6,
			},
		},
		Display: "sats",
		Base:    "bitcoin",
		Name:    "bitcoin",
		Symbol:  "BTC",
	}

	// make a proper setDenomMetadata message
	baseMsg := types.NewMsgSetDenomMetadata(
		addr1.String(),
		denomMetadata,
	)

	// validate setDenomMetadata message was created as intended
	require.Equal(t, baseMsg.Route(), types.RouterKey)
	require.Equal(t, baseMsg.Type(), "set_denom_metadata")
	signers := baseMsg.GetSigners()
	require.Equal(t, len(signers), 1)
	require.Equal(t, signers[0].String(), addr1.String())

	tests := []struct {
		name       string
		msg        func() *types.MsgSetDenomMetadata
		expectPass bool
	}{
		{
			name: "proper msg",
			msg: func() *types.MsgSetDenomMetadata {
				msg := baseMsg
				return msg
			},
			expectPass: true,
		},
		{
			name: "empty sender",
			msg: func() *types.MsgSetDenomMetadata {
				msg := baseMsg
				msg.Sender = ""
				return msg
			},
			expectPass: false,
		},
		{
			name: "invalid metadata",
			msg: func() *types.MsgSetDenomMetadata {
				msg := baseMsg
				msg.Metadata.Name = ""
				return msg
			},

			expectPass: false,
		},
		{
			name: "invalid base",
			msg: func() *types.MsgSetDenomMetadata {
				msg := baseMsg
				msg.Metadata = invalidDenomMetadata
				return msg
			},
			expectPass: false,
		},
	}

	for _, test := range tests {
		if test.expectPass {
			require.NoError(t, test.msg().ValidateBasic(), "test: %v", test.name)
		} else {
			require.Error(t, test.msg().ValidateBasic(), "test: %v", test.name)
		}
	}
}
