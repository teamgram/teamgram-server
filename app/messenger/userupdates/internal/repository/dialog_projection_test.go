//go:build integration

package repository

import (
	"context"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestListDialogsOrdersByTopDatePeerSeqAndSkipsHidden(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 1001
	repo := NewForTest(db, &testIDGenerator{next: base + 10_000}, "local-userupdates")

	insertUserDialogProjection(t, repo, userID, payload.PeerTypeUser, 10, 1, 1_778_201_600, false)
	insertUserDialogProjection(t, repo, userID, payload.PeerTypeUser, 11, 2, 1_778_201_600, false)
	insertUserDialogProjection(t, repo, userID, payload.PeerTypeUser, 12, 1, 1_778_201_601, true)

	rows, err := repo.ListDialogProjections(ctx, userID, DialogProjectionCursor{}, 10)
	if err != nil {
		t.Fatalf("ListDialogProjections() error = %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("row count = %d, want 2: %+v", len(rows), rows)
	}
	if rows[0].PeerID != 11 || rows[1].PeerID != 10 {
		t.Fatalf("order = [%d %d], want [11 10]", rows[0].PeerID, rows[1].PeerID)
	}
}

func TestListDialogsHandlesZeroDeletedAt(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 1501
	repo := NewForTest(db, &testIDGenerator{next: base + 15_000}, "local-userupdates")

	insertUserDialogProjectionWithDeletedAt(t, repo, userID, payload.PeerTypeUser, 10, 1, 1_778_201_600, false, int64(0))

	rows, err := repo.ListDialogProjections(ctx, userID, DialogProjectionCursor{}, 10)
	if err != nil {
		t.Fatalf("ListDialogProjections() error = %v", err)
	}
	if len(rows) != 1 || rows[0].PeerID != 10 {
		t.Fatalf("rows = %+v, want peer 10", rows)
	}
}

func TestGetDialogsByPeersPreservesRequestedPeerCoverage(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 2001
	repo := NewForTest(db, &testIDGenerator{next: base + 20_000}, "local-userupdates")

	insertUserDialogProjection(t, repo, userID, payload.PeerTypeUser, 10, 1, 1_778_201_600, false)

	rows, err := repo.GetDialogProjectionsByPeers(ctx, userID, []DialogProjectionPeer{
		{PeerType: payload.PeerTypeUser, PeerID: 10},
		{PeerType: payload.PeerTypeUser, PeerID: 11},
	})
	if err != nil {
		t.Fatalf("GetDialogProjectionsByPeers() error = %v", err)
	}
	if _, ok := rows[DialogProjectionPeer{PeerType: payload.PeerTypeUser, PeerID: 10}]; !ok {
		t.Fatalf("existing peer missing: %+v", rows)
	}
	if _, ok := rows[DialogProjectionPeer{PeerType: payload.PeerTypeUser, PeerID: 11}]; ok {
		t.Fatalf("missing peer should be omitted: %+v", rows)
	}
}

func TestCountVisibleDialogsSkipsHidden(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 3001
	repo := NewForTest(db, &testIDGenerator{next: base + 30_000}, "local-userupdates")

	insertUserDialogProjection(t, repo, userID, payload.PeerTypeUser, 10, 1, 1_778_201_600, false)
	insertUserDialogProjection(t, repo, userID, payload.PeerTypeUser, 11, 2, 1_778_201_600, true)

	count, err := repo.CountVisibleDialogs(ctx, userID)
	if err != nil {
		t.Fatalf("CountVisibleDialogs() error = %v", err)
	}
	if count != 1 {
		t.Fatalf("CountVisibleDialogs() = %d, want 1", count)
	}
}

func insertUserDialogProjection(t *testing.T, repo *Repository, userID int64, peerType int32, peerID int64, topPeerSeq int64, topMessageDate int64, hidden bool) {
	t.Helper()
	deletedAt := any(int64(0))
	insertUserDialogProjectionWithDeletedAt(t, repo, userID, peerType, peerID, topPeerSeq, topMessageDate, hidden, deletedAt)
}

func insertUserDialogProjectionWithDeletedAt(t *testing.T, repo *Repository, userID int64, peerType int32, peerID int64, topPeerSeq int64, topMessageDate int64, hidden bool, deletedAt any) {
	t.Helper()
	_, err := repo.db.Exec(context.Background(), `
INSERT INTO user_dialogs
	(user_id, peer_type, peer_id, top_peer_seq, top_canonical_message_id,
	 top_message_date, top_message_status, read_inbox_max_peer_seq,
	 read_outbox_max_peer_seq, unread_count, unread_mentions_count,
	 unread_reactions_count, unread_mark, pinned_peer_seq,
	 pinned_canonical_message_id, has_scheduled, available_min_peer_seq,
	 hidden, deleted_at, last_pts, last_pts_at, dialog_schema_version,
	 dialog_payload)
VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		userID,
		peerType,
		peerID,
		topPeerSeq,
		peerID*10+topPeerSeq,
		topMessageDate,
		MessageStatusLive,
		int64(0),
		int64(0),
		int32(0),
		int32(0),
		int32(0),
		false,
		int64(0),
		int64(0),
		false,
		int64(0),
		hidden,
		deletedAt,
		topPeerSeq,
		topMessageDate,
		int32(1),
		[]byte(`{"test":true}`),
	)
	if err != nil {
		t.Fatalf("insert user_dialogs projection: %v", err)
	}
}
