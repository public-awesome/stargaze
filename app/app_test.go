package app_test

import (
	"testing"

	"github.com/public-awesome/stargaze/v12/testutil/simapp"
)

func TestAnteHandler(t *testing.T) {
	simapp.New(t)
	// suite.app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "stargaze-1", Time: time.Now().UTC()})
}
