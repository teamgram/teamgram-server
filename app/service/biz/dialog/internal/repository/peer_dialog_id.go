package repository

import (
	"errors"
	"fmt"
	"math"
)

const (
	PeerTypeUser    int32 = 1
	PeerTypeChat    int32 = 2
	PeerTypeChannel int32 = 3

	PeerPolicyScopePrivatePair = "private_pair"
)

var (
	ErrInvalidPeerType      = errors.New("dialog: invalid peer type")
	ErrInvalidPeerID        = errors.New("dialog: invalid peer id")
	ErrPeerDialogIDOverflow = errors.New("dialog: peer dialog id overflow")
	ErrCorruptPeerDialogID  = errors.New("dialog: corrupt peer dialog id")
)

type PeerRef struct {
	PeerType int32
	PeerID   int64
}

type PolicyScope struct {
	ScopeType string
	ScopeID   string
	PeerType  int32
	PeerID    int64
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

func SplitPeerDialogID(peerDialogID int64) (PeerRef, error) {
	if peerDialogID <= 0 {
		return PeerRef{}, ErrCorruptPeerDialogID
	}
	normalized := int32(peerDialogID % 16)
	peerID := peerDialogID / 16
	peerType, err := denormalizePeerType(normalized)
	if err != nil {
		return PeerRef{}, fmt.Errorf("%w: normalized_type=%d", ErrCorruptPeerDialogID, normalized)
	}
	if peerID <= 0 {
		return PeerRef{}, ErrCorruptPeerDialogID
	}
	return PeerRef{PeerType: peerType, PeerID: peerID}, nil
}

func MakePrivatePairScope(userA, userB int64) (PolicyScope, error) {
	if userA <= 0 || userB <= 0 {
		return PolicyScope{}, ErrInvalidPeerID
	}
	if userB < userA {
		userA, userB = userB, userA
	}
	return PolicyScope{
		ScopeType: PeerPolicyScopePrivatePair,
		ScopeID:   fmt.Sprintf("%d:%d", userA, userB),
		PeerType:  0,
		PeerID:    0,
	}, nil
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
