package statesync

import (
	"io"

	snapshot "github.com/cosmos/cosmos-sdk/snapshots/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	protoio "github.com/gogo/protobuf/io"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const (
	SnapshotName   = "sg_version_snapshotter"
	SnapshotFormat = 0xFF
)

var _ snapshot.ExtensionSnapshotter = &VersionSnapshotter{}

type ConsensusParamsGetter interface {
	GetConsensusParams(ctx sdk.Context) *abci.ConsensusParams
}

type ProtocolVersionSetter interface {
	SetProtocolVersion(uint64)
}

type VersionSnapshotter struct {
	consensusParamGetter ConsensusParamsGetter
	versionSetter        ProtocolVersionSetter
	ms                   sdk.MultiStore
}

func NewVersionSnapshotter(ms sdk.MultiStore, cpg ConsensusParamsGetter, vs ProtocolVersionSetter) *VersionSnapshotter {
	return &VersionSnapshotter{
		consensusParamGetter: cpg,
		versionSetter:        vs,
		ms:                   ms,
	}
}

func (vs *VersionSnapshotter) SnapshotName() string {
	return SnapshotName
}

// Snapshot writes snapshot items into the protobuf writer.
func (vs *VersionSnapshotter) Snapshot(height uint64, protoWriter protoio.Writer) error {
	cms, err := vs.ms.CacheMultiStoreWithVersion(int64(height))
	if err != nil {
		return err
	}
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	params := vs.consensusParamGetter.GetConsensusParams(ctx)
	// default to 1 for stargaze
	appVersion := uint64(1)
	if params != nil && params.Version != nil && params.Version.GetAppVersion() > 0 {
		appVersion = params.Version.GetAppVersion()
	}
	bz := sdk.Uint64ToBigEndian(appVersion)
	return snapshot.WriteExtensionItem(protoWriter, bz)
}

// Restore restores a state snapshot from the protobuf items read from the reader.
func (vs *VersionSnapshotter) Restore(height uint64, format uint32, protoReader protoio.Reader) (snapshot.SnapshotItem, error) {
	if format == SnapshotFormat {
		var item snapshot.SnapshotItem
		for {
			item = snapshot.SnapshotItem{}
			err := protoReader.ReadMsg(&item)
			if err == io.EOF {
				break
			} else if err != nil {
				return snapshot.SnapshotItem{}, sdkerrors.Wrap(err, "invalid protobuf message")
			}
			payload := item.GetExtensionPayload()
			if payload == nil {
				break
			}
			appVersion := sdk.BigEndianToUint64(payload.Payload)
			vs.versionSetter.SetProtocolVersion(appVersion)
		}
		return item, nil
	}
	return snapshot.SnapshotItem{}, snapshot.ErrUnknownFormat
}

func (vs *VersionSnapshotter) SnapshotFormat() uint32 {
	return SnapshotFormat
}

func (vs *VersionSnapshotter) SupportedFormats() []uint32 {
	return []uint32{SnapshotFormat}
}
