package streaming

import (
	"context"
	"sort"
	"strings"
	"sync"

	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/store/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/nats-io/nats.go"
)

var _ baseapp.StreamingService = &StreamingService{}

// Nats streaming option keys
const (
	OptNatsUrl = "store.streamers.nats.nats_url"
)

// StreamingService is a concrete implementation of StreamingService that accumulate the state changes in current block,
// writes the ordered changeset out to a Nats Jet Stream service.
type StreamingService struct {
	listeners          []*types.MemoryListener // the listeners that will be initialized with BaseApp
	codec              codec.BinaryCodec       // marshaller used for re-marshalling the ABCI messages to write them out to the destination files
	conn               *nats.Conn
	currentBlockNumber int64
}

// NewNatsStreamingService is the streaming.ServiceConstructor function for
// creating a NatsStreamingService.
func NewNatsStreamingService(
	opts servertypes.AppOptions,
	keys []storetypes.StoreKey,
	cdc codec.BinaryCodec,
) (baseapp.StreamingService, error) {
	natsUrl := cast.ToString(opts.Get(OptNatsUrl))
	if natsUrl == "" {
		natsUrl = nats.DefaultURL
	}

	service, err := NewStreamingService(natsUrl, keys, cdc)
	if err != nil {
		return nil, err
	}

	return service, nil
}

// NewStreamingService creates a new StreamingService with a nats connection for the provided url.
func NewStreamingService(natsUrl string, storeKeys []types.StoreKey, cdc codec.BinaryCodec) (*StreamingService, error) {
	// sort by the storeKeys first
	sort.SliceStable(storeKeys, func(i, j int) bool {
		return strings.Compare(storeKeys[i].Name(), storeKeys[j].Name()) < 0
	})

	listeners := make([]*types.MemoryListener, len(storeKeys))
	for i, key := range storeKeys {
		listeners[i] = types.NewMemoryListener(key)
	}

	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	return &StreamingService{listeners, cdc, conn, 0}, nil
}

// Listeners satisfies the baseapp.StreamingService interface
func (nss *StreamingService) Listeners() map[types.StoreKey][]types.WriteListener {
	listeners := make(map[types.StoreKey][]types.WriteListener, len(nss.listeners))
	for _, listener := range nss.listeners {
		listeners[listener.StoreKey()] = []types.WriteListener{listener}
	}
	return listeners
}

// ListenBeginBlock satisfies the baseapp.ABCIListener interface
// It sets the currentBlockNumber.
func (nss *StreamingService) ListenBeginBlock(ctx context.Context, req abci.RequestBeginBlock, res abci.ResponseBeginBlock) error {
	nss.currentBlockNumber = req.GetHeader().Height
	return nil
}

// ListenDeliverTx satisfies the baseapp.ABCIListener interface
func (nss *StreamingService) ListenDeliverTx(ctx context.Context, req abci.RequestDeliverTx, res abci.ResponseDeliverTx) error {
	return nil
}

// ListenEndBlock satisfies the baseapp.ABCIListener interface
// It merge the state caches of all the listeners together, and write out to the versionStore.
func (nss *StreamingService) ListenEndBlock(ctx context.Context, req abci.RequestEndBlock, res abci.ResponseEndBlock) error {
	return nil
}

func (nss *StreamingService) ListenCommit(ctx context.Context, res abci.ResponseCommit) error {
	// concat the state caches
	for _, listener := range nss.listeners {
		for _, changeSet := range listener.PopStateCache() {
			message, err := nss.codec.Marshal(&changeSet)
			if err != nil {
				return err
			}

			nss.conn.Publish(listener.StoreKey().Name(), message)
		}
	}

	return nil
}

// Stream satisfies the baseapp.StreamingService interface
func (nss *StreamingService) Stream(wg *sync.WaitGroup) error {
	return nil
}

// Close satisfies the io.Closer interface, which satisfies the baseapp.StreamingService interface
func (nss *StreamingService) Close() error {
	nss.conn.Close()
	return nil
}
