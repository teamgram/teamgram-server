//go:build integration

package repository

import (
	"context"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestGetMessageViewsByPeerSeqsReturnsRequestedTopViews(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 9001
	repo := NewForTest(db, &testIDGenerator{next: base + 90_000}, "local-userupdates")

	in := buildApplyInput(t, userID, userID, userID+1, true, "dialog top")
	if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, in); err != nil {
		t.Fatalf("ApplyUserOperation() error = %v", err)
	}

	rows, err := repo.GetMessageViewsByPeerSeqs(ctx, userID, []MessageViewPeerSeq{
		{PeerType: payload.PeerTypeUser, PeerID: userID + 1, PeerSeq: 1},
		{PeerType: payload.PeerTypeUser, PeerID: userID + 2, PeerSeq: 1},
	})
	if err != nil {
		t.Fatalf("GetMessageViewsByPeerSeqs() error = %v", err)
	}
	row, ok := rows[MessageViewPeerSeq{PeerType: payload.PeerTypeUser, PeerID: userID + 1, PeerSeq: 1}]
	if !ok {
		t.Fatalf("existing message view missing: %+v", rows)
	}
	if row.UserID != userID || row.PeerSeq != 1 || string(row.ViewPayload) == "" {
		t.Fatalf("unexpected message view: %+v", row)
	}
	if _, ok := rows[MessageViewPeerSeq{PeerType: payload.PeerTypeUser, PeerID: userID + 2, PeerSeq: 1}]; ok {
		t.Fatalf("missing peer should be omitted: %+v", rows)
	}
}
