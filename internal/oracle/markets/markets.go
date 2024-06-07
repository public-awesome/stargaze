package markets

import (
	_ "embed"
	"encoding/json"
	"slices"
	"strings"

	marketmaptypes "github.com/skip-mev/slinky/x/marketmap/types"
)

//go:embed markets.json
var marketsRawBz []byte

func Map() (markets marketmaptypes.MarketMap, err error) {
	err = json.Unmarshal(marketsRawBz, &markets)
	return
}

func Slice() ([]marketmaptypes.Market, error) {
	m, err := Map()
	markets := make([]marketmaptypes.Market, 0, len(m.Markets))
	if err != nil {
		return nil, err
	}
	for _, market := range m.Markets {
		markets = append(markets, market)
	}

	// sort them for deterministic purposes
	slices.SortFunc(markets, func(a, b marketmaptypes.Market) int {
		return strings.Compare(a.Ticker.String(), b.Ticker.String())
	})
	return markets, nil
}
