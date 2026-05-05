package repository

import (
	"errors"
	"math"
	"testing"
)

func TestMakePeerDialogIDRoundTrip(t *testing.T) {
	for _, tt := range []struct {
		peerType int32
		peerID   int64
	}{
		{PeerTypeUser, 10},
		{PeerTypeChat, 20},
		{PeerTypeChannel, 30},
	} {
		t.Run("", func(t *testing.T) {
			id, err := MakePeerDialogID(tt.peerType, tt.peerID)
			if err != nil {
				t.Fatalf("MakePeerDialogID(%d,%d) error = %v", tt.peerType, tt.peerID, err)
			}
			got, err := SplitPeerDialogID(id)
			if err != nil {
				t.Fatalf("SplitPeerDialogID(%d) error = %v", id, err)
			}
			if got.PeerType != tt.peerType || got.PeerID != tt.peerID {
				t.Fatalf("round trip = (%d,%d), want (%d,%d)", got.PeerType, got.PeerID, tt.peerType, tt.peerID)
			}
		})
	}
}

func TestMakePeerDialogIDRejectsInvalidPeer(t *testing.T) {
	if _, err := MakePeerDialogID(99, 1); !errors.Is(err, ErrInvalidPeerType) {
		t.Fatalf("MakePeerDialogID invalid type error = %v, want ErrInvalidPeerType", err)
	}
	if _, err := MakePeerDialogID(PeerTypeUser, 0); !errors.Is(err, ErrInvalidPeerID) {
		t.Fatalf("MakePeerDialogID non-positive id error = %v, want ErrInvalidPeerID", err)
	}
	if _, err := MakePeerDialogID(PeerTypeUser, (math.MaxInt64-15)/16+1); !errors.Is(err, ErrPeerDialogIDOverflow) {
		t.Fatalf("MakePeerDialogID overflow error = %v, want ErrPeerDialogIDOverflow", err)
	}
	if _, err := SplitPeerDialogID(16); !errors.Is(err, ErrCorruptPeerDialogID) {
		t.Fatalf("SplitPeerDialogID corrupt error = %v, want ErrCorruptPeerDialogID", err)
	}
}

func TestPrivatePairScopeCanonicalizesUsers(t *testing.T) {
	first, err := MakePrivatePairScope(5, 9)
	if err != nil {
		t.Fatalf("MakePrivatePairScope(5,9) error = %v", err)
	}
	second, err := MakePrivatePairScope(9, 5)
	if err != nil {
		t.Fatalf("MakePrivatePairScope(9,5) error = %v", err)
	}
	if first != second {
		t.Fatalf("scope mismatch: first=%+v second=%+v", first, second)
	}
	if first.ScopeType != PeerPolicyScopePrivatePair || first.ScopeID != "5:9" {
		t.Fatalf("scope = %+v, want private_pair 5:9", first)
	}
	if first.PeerType != 0 || first.PeerID != 0 {
		t.Fatalf("private pair diagnostic peer = (%d,%d), want (0,0)", first.PeerType, first.PeerID)
	}
}

func TestPrivatePairScopeRejectsNonPositiveUserID(t *testing.T) {
	if _, err := MakePrivatePairScope(0, 9); !errors.Is(err, ErrInvalidPeerID) {
		t.Fatalf("MakePrivatePairScope(0,9) error = %v, want ErrInvalidPeerID", err)
	}
	if _, err := MakePrivatePairScope(5, -1); !errors.Is(err, ErrInvalidPeerID) {
		t.Fatalf("MakePrivatePairScope(5,-1) error = %v, want ErrInvalidPeerID", err)
	}
}
