package paging

import "testing"

func TestNormalizeDialogLimit(t *testing.T) {
	tests := []struct {
		name string
		in   int32
		want int32
	}{
		{name: "zero", in: 0, want: 0},
		{name: "negative", in: -10, want: 0},
		{name: "inside cap", in: 100, want: 100},
		{name: "over cap", in: 600, want: DialogListLimitHardCap},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeDialogLimit(tt.in); got != tt.want {
				t.Fatalf("NormalizeDialogLimit(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestPeerDialogIDRoundTrip(t *testing.T) {
	for _, tt := range []struct {
		peerType int32
		peerID   int64
	}{
		{PeerTypeUser, 10},
		{PeerTypeChat, 20},
		{PeerTypeChannel, 30},
	} {
		id, err := MakePeerDialogID(tt.peerType, tt.peerID)
		if err != nil {
			t.Fatalf("MakePeerDialogID(%d,%d): %v", tt.peerType, tt.peerID, err)
		}
		gotType, gotID, err := SplitPeerDialogID(id)
		if err != nil {
			t.Fatalf("SplitPeerDialogID(%d): %v", id, err)
		}
		if gotType != tt.peerType || gotID != tt.peerID {
			t.Fatalf("round trip = (%d,%d), want (%d,%d)", gotType, gotID, tt.peerType, tt.peerID)
		}
	}
}

func TestPeerDialogIDRejectsInvalidInput(t *testing.T) {
	if _, err := MakePeerDialogID(99, 1); err == nil {
		t.Fatalf("expected invalid peer type error")
	}
	if _, err := MakePeerDialogID(PeerTypeUser, 0); err == nil {
		t.Fatalf("expected non-positive peer id error")
	}
	if _, _, err := SplitPeerDialogID(16); err == nil {
		t.Fatalf("expected unknown normalized type error")
	}
}

func TestPinnedSnapshotMismatchRestartsPage(t *testing.T) {
	cursor := DialogCursor{Section: DialogSectionPinned, PinnedSnapshotVersion: 10}
	if !cursor.ShouldRestartForPinnedVersion(11) {
		t.Fatalf("expected stale pinned cursor to restart")
	}
	if cursor.ShouldRestartForPinnedVersion(10) {
		t.Fatalf("did not expect matching pinned cursor to restart")
	}
}

func TestRegularCursorIgnoresPinOrder(t *testing.T) {
	a := DialogCursor{Section: DialogSectionRegular, TopMessageDate: 100, TopPeerSeq: 9, PeerType: PeerTypeUser, PeerID: 1, PinOrder: 999}
	b := DialogCursor{Section: DialogSectionRegular, TopMessageDate: 100, TopPeerSeq: 9, PeerType: PeerTypeUser, PeerID: 1, PinOrder: 1}
	if a.RegularKey() != b.RegularKey() {
		t.Fatalf("regular cursor key must ignore pin order")
	}
}
