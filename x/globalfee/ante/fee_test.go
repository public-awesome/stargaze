package ante_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func TestFeeDecoratorAntehandler(t *testing.T) {
	// Test cases to check
	// where min_gas_price = min gas price configured by the validator
	// where globalfee = minimum gas price configured by x/globalfee module param
	// where feeSent = the amount of fee sent by the user in the tx
	// where msgType = whether this msg has been configured to be free by the x/globalfee module
	//
	// min_gas_price:  empty, globalfee: 5stake, feeSent: 1stake msgType: notAuthd - should fail
	// min_gas_price:  empty, globalfee: 5stake, feeSent: 7stake msgType: notAuthd - should pass
	// min_gas_price: 0stake, globalfee: 5stake, feeSent: 0stake msgType: notAuthd - should pass
	// min_gas_price: 2stake, globalfee: 5stake, feeSent: 1stake msgType: notAuthd - should fail
	// min_gas_price: 2stake, globalfee: 5stake, feeSent: 3stake msgType: notAuthd - should pass
	// min_gas_price: 2stake, globalfee: 5stake, feeSent: 0stake msgType: authd    - should pass

	// todo add the actual tests
}

func getTestAccount() (privateKey secp256k1.PrivKey, publicKey crypto.PubKey, accountAddress sdk.AccAddress) {
	privateKey = secp256k1.GenPrivKey()
	publicKey = privateKey.PubKey()
	accountAddress = sdk.AccAddress(publicKey.Address())
	return
}
