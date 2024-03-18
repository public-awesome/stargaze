package statesync

import (
	"io"

	errorsmod "cosmossdk.io/errors"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	snapshot "github.com/cosmos/cosmos-sdk/snapshots/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	SnapshotName   = "sg_version_snapshotter"
	SnapshotFormat = 0xFF
)

var _ snapshot.ExtensionSnapshotter = &VersionSnapshotter{}

type ConsensusParamsGetter interface {
	GetConsensusParams(ctx sdk.Context) *tmproto.ConsensusParams
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
func (vs *VersionSnapshotter) SnapshotExtension(height uint64, payloadWriter snapshot.ExtensionPayloadWriter) error {
	cms, err := vs.ms.CacheMultiStoreWithVersion(int64(height))
	if err != nil {
		return err
	}
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	params := vs.consensusParamGetter.GetConsensusParams(ctx)
	// default to 1 for testnet
	appVersion := uint64(0)
	if params != nil && params.Version != nil {
		appVersion = params.Version.GetApp()
	}
	bz := sdk.Uint64ToBigEndian(appVersion)
	return payloadWriter(bz)
}

// Restore restores a state snapshot from the protobuf items read from the reader.
func (vs *VersionSnapshotter) RestoreExtension(_ uint64, format uint32, payloadReader snapshot.ExtensionPayloadReader) error {
	if format == SnapshotFormat {
		for {
			payload, err := payloadReader()
			if err == io.EOF {
				break
			} else if err != nil {
				return errorsmod.Wrap(err, "invalid protobuf message")
			}
			if payload == nil {
				break
			}
			appVersion := sdk.BigEndianToUint64(payload)
			vs.versionSetter.SetProtocolVersion(appVersion)
		}
		return nil
	}
	return snapshot.ErrUnknownFormat
}

func (vs *VersionSnapshotter) SnapshotFormat() uint32 {
	return SnapshotFormat
}

func (vs *VersionSnapshotter) SupportedFormats() []uint32 {
	return []uint32{SnapshotFormat}
}
