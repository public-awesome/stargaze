package app_test

import (
	"testing"

	"github.com/public-awesome/stargaze/v6/testutil/simapp"
)

func TestAnteHandler(t *testing.T) {
	simapp.New(t.TempDir())
	// suite.app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "stargaze-1", Time: time.Now().UTC()})
}
