package paging

import (
	"errors"
	"fmt"
	"math"
)

const (
	PeerTypeUser    int32 = 1
	PeerTypeChat    int32 = 2
	PeerTypeChannel int32 = 3

	DialogListLimitHardCap          int32 = 500
	DialogOverfetchFactor           int32 = 3
	DialogOverfetchBatchMin         int32 = 50
	DialogOverfetchHardCap          int32 = 2000
	DialogHydrateBatchSize          int32 = 200
	DialogMaxHydratePeersPerRequest int32 = 2000
)

type DialogSection string

const (
	DialogSectionPinned  DialogSection = "pinned"
	DialogSectionRegular DialogSection = "regular"
)

var (
	ErrInvalidPeerType      = errors.New("paging: invalid peer type")
	ErrInvalidPeerID        = errors.New("paging: invalid peer id")
	ErrPeerDialogIDOverflow = errors.New("paging: peer dialog id overflow")
	ErrCorruptPeerDialogID  = errors.New("paging: corrupt peer dialog id")
)

type DialogCursor struct {
	FolderID              int32
	Section               DialogSection
	PinnedSnapshotVersion int64
	PinOrder              int64
	TopMessageDate        int64
	TopPeerSeq            int64
	PeerType              int32
	PeerID                int64
}

type RegularDialogKey struct {
	TopMessageDate int64
	TopPeerSeq     int64
	PeerType       int32
	PeerID         int64
}

func NormalizeDialogLimit(limit int32) int32 {
	if limit <= 0 {
		return 0
	}
	if limit > DialogListLimitHardCap {
		return DialogListLimitHardCap
	}
	return limit
}

func DialogOverfetchLimit(limit int32) int32 {
	limit = NormalizeDialogLimit(limit)
	if limit == 0 {
		return 0
	}
	scan := limit * DialogOverfetchFactor
	if scan < DialogOverfetchBatchMin {
		scan = DialogOverfetchBatchMin
	}
	if scan > DialogOverfetchHardCap {
		scan = DialogOverfetchHardCap
	}
	return scan
}

func (c DialogCursor) ShouldRestartForPinnedVersion(current int64) bool {
	return c.Section == DialogSectionPinned && c.PinnedSnapshotVersion != current
}

func (c DialogCursor) RegularKey() RegularDialogKey {
	return RegularDialogKey{
		TopMessageDate: c.TopMessageDate,
		TopPeerSeq:     c.TopPeerSeq,
		PeerType:       c.PeerType,
		PeerID:         c.PeerID,
	}
}

func MakePeerDialogID(peerType int32, peerID int64) (int64, error) {
	normalized, err := normalizePeerType(peerType)
	if err != nil {
		return 0, err
	}
	if peerID <= 0 {
		return 0, ErrInvalidPeerID
	}
	if peerID > (math.MaxInt64-15)/16 {
		return 0, ErrPeerDialogIDOverflow
	}
	return peerID*16 + int64(normalized), nil
}

func SplitPeerDialogID(peerDialogID int64) (int32, int64, error) {
	if peerDialogID <= 0 {
		return 0, 0, ErrCorruptPeerDialogID
	}
	normalized := int32(peerDialogID % 16)
	peerID := peerDialogID / 16
	peerType, err := denormalizePeerType(normalized)
	if err != nil {
		return 0, 0, fmt.Errorf("%w: normalized_type=%d", ErrCorruptPeerDialogID, normalized)
	}
	if peerID <= 0 {
		return 0, 0, ErrCorruptPeerDialogID
	}
	return peerType, peerID, nil
}

func normalizePeerType(peerType int32) (int32, error) {
	switch peerType {
	case PeerTypeUser, PeerTypeChat, PeerTypeChannel:
		return peerType, nil
	default:
		return 0, ErrInvalidPeerType
	}
}

func denormalizePeerType(normalized int32) (int32, error) {
	switch normalized {
	case PeerTypeUser, PeerTypeChat, PeerTypeChannel:
		return normalized, nil
	default:
		return 0, ErrInvalidPeerType
	}
}
