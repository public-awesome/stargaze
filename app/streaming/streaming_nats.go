package streaming

import (
	"context"
	"fmt"
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

// Nats streaming option keys
const (
	OptNatsUrl             = "store.streamers.nats.url"
	OptNatsPrefix          = "store.streamers.nats.prefix"
	OptNatsOutputMetadata  = "store.streamers.nats.output_metadata"
	OptNatsStopNodeOnError = "store.streamers.nats.stop_node_on_error"

	NatsBlocksSegment    = "blocks"
	NatsChangeSetSegment = "changeset"
)

var _ baseapp.StreamingService = &StreamingService{}

// StreamingService is a concrete implementation of StreamingService that accumulate the state changes in current block,
// writes the ordered changeset out to a Nats Jet Stream service.
type StreamingService struct {
	natsConn       *nats.Conn              // nats connection
	storeListeners []*types.MemoryListener // a series of KVStore listeners for each KVStore
	subjectPrefix  string                  // optional prefix for each of the generated subjects
	codec          codec.BinaryCodec       // marshaller used for re-marshalling the ABCI messages to write them out to the destination files

	blockMetadata types.BlockMetadata

	// outputMetadata, if true, writes additional metadata out to nats
	outputMetadata bool

	// stopNodeOnErr, if true, will panic and stop the node during ABCI Commit
	// to ensure eventual consistency of the output, otherwise, any errors are
	// logged and ignored which could yield data loss in streamed output.
	stopNodeOnErr bool
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

	subjectPrefix := cast.ToString(opts.Get(OptNatsPrefix))
	outputMetadata := cast.ToBool(opts.Get(OptNatsOutputMetadata))
	stopNodeOnError := cast.ToBool(opts.Get(OptNatsStopNodeOnError))

	return NewStreamingService(natsUrl, keys, subjectPrefix, cdc, outputMetadata, stopNodeOnError)
}

// NewStreamingService creates a new StreamingService with a nats connection for the provided url.
func NewStreamingService(
	natsUrl string,
	storeKeys []types.StoreKey,
	subjectPrefix string,
	cdc codec.BinaryCodec,
	outputMetadata bool,
	stopNodeOnErr bool,
) (*StreamingService, error) {
	// sort by the storeKeys first
	sort.SliceStable(storeKeys, func(i, j int) bool {
		return strings.Compare(storeKeys[i].Name(), storeKeys[j].Name()) < 0
	})

	storeListeners := make([]*types.MemoryListener, len(storeKeys))
	for i, key := range storeKeys {
		storeListeners[i] = types.NewMemoryListener(key)
	}

	natsConn, err := nats.Connect(natsUrl)
	if err != nil {
		return nil, err
	}

	return &StreamingService{
		natsConn:       natsConn,
		storeListeners: storeListeners,
		subjectPrefix:  subjectPrefix,
		codec:          cdc,
		outputMetadata: outputMetadata,
		stopNodeOnErr:  stopNodeOnErr,
	}, nil
}

// Listeners satisfies the StreamingService interface. It returns the
// StreamingService's underlying WriteListeners. Use for registering the
// underlying WriteListeners with the BaseApp.
func (nss *StreamingService) Listeners() map[types.StoreKey][]types.WriteListener {
	listeners := make(map[types.StoreKey][]types.WriteListener, len(nss.storeListeners))
	for _, listener := range nss.storeListeners {
		listeners[listener.StoreKey()] = []types.WriteListener{listener}
	}

	return listeners
}

// ListenBeginBlock satisfies the ABCIListener interface. It sets the received
// BeginBlock request, response and the current block number. Note, these are
// not published to nats until ListenCommit is executed and outputMetadata is set,
// after which it will be reset again on the next block.
func (nss *StreamingService) ListenBeginBlock(ctx context.Context, req abci.RequestBeginBlock, res abci.ResponseBeginBlock) error {
	nss.blockMetadata.RequestBeginBlock = &req
	nss.blockMetadata.ResponseBeginBlock = &res
	return nil
}

// ListenDeliverTx satisfies the ABCIListener interface. It appends the received
// DeliverTx request and response to a list of DeliverTxs objects. Note, these
// are not published to nats until ListenCommit is executed and outputMetadata is
// set, after which it will be reset again on the next block.
func (nss *StreamingService) ListenDeliverTx(ctx context.Context, req abci.RequestDeliverTx, res abci.ResponseDeliverTx) error {
	nss.blockMetadata.DeliverTxs = append(nss.blockMetadata.DeliverTxs, &types.BlockMetadata_DeliverTx{
		Request:  &req,
		Response: &res,
	})

	return nil
}

// ListenEndBlock satisfies the ABCIListener interface. It sets the received
// EndBlock request, response and the current block number. Note, these are
// not published to nats until ListenCommit is executed and outputMetadata is set,
// after which it will be reset again on the next block.
func (nss *StreamingService) ListenEndBlock(ctx context.Context, req abci.RequestEndBlock, res abci.ResponseEndBlock) error {
	nss.blockMetadata.RequestEndBlock = &req
	nss.blockMetadata.ResponseEndBlock = &res
	return nil
}

// ListenCommit satisfies the ABCIListener interface. It is executed during the
// ABCI Commit request and is responsible for writing all staged data to files.
// It will only return a non-nil error when stopNodeOnErr is set.
func (nss *StreamingService) ListenCommit(ctx context.Context, res abci.ResponseCommit) error {
	nss.blockMetadata.ResponseCommit = &res

	if err := nss.doListenCommit(); err != nil {
		if nss.stopNodeOnErr {
			return err
		}
	}

	return nil
}

func (nss *StreamingService) doListenCommit() (err error) {
	// Write to target files, the file size is written at the beginning, which can
	// be used to detect completeness.
	metadataSubject := NatsBlocksSegment
	if nss.subjectPrefix != "" {
		metadataSubject = fmt.Sprintf("%s.%s", nss.subjectPrefix, metadataSubject)
	}

	if nss.outputMetadata {
		message, err := nss.codec.Marshal(&nss.blockMetadata)
		if err != nil {
			return err
		}
		err = nss.natsConn.Publish(metadataSubject, message)
		if err != nil {
			return err
		}
	}

	changeSetSubjet := NatsChangeSetSegment
	if nss.subjectPrefix != "" {
		changeSetSubjet = fmt.Sprintf("%s.%s", nss.subjectPrefix, changeSetSubjet)
	}

	for _, listener := range nss.storeListeners {
		for _, changeSet := range listener.PopStateCache() {
			message, err := nss.codec.Marshal(&changeSet)
			if err != nil {
				return err
			}

			subject := fmt.Sprintf("%s.%s", changeSetSubjet, listener.StoreKey().Name())
			err = nss.natsConn.Publish(subject, message)
			if err != nil {
				return err
			}
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
	nss.natsConn.Close()
	return nil
}
