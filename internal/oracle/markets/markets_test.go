package markets_test

import (
	"testing"

	"github.com/public-awesome/stargaze/v15/internal/oracle/markets"
	"github.com/stretchr/testify/require"
)

func TestMarkets(t *testing.T) {
	m, err := markets.Map()
	require.NoError(t, err)
	require.Len(t, m.Markets, 4)

	s, err := markets.Slice()
	require.NoError(t, err)
	require.Len(t, s, 4)

	require.Equal(t, "ATOM/USD", s[0].Ticker.String())
	require.Equal(t, "BTC/USD", s[1].Ticker.String())
	require.Equal(t, "ETH/USD", s[2].Ticker.String())
	require.Equal(t, "USDT/USD", s[3].Ticker.String())
}
