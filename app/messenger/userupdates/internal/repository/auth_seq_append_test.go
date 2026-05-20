//go:build integration

package repository

import (
	"context"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestAppendAuthSeqUpdateExpandsPerAuthKey(t *testing.T) {
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 300_000}, "local-userupdates")
	body := []byte{0x01, 0x02, 0x03}
	hash := payload.HashBytes(body)

	got, err := repo.AppendAuthSeqUpdate(context.Background(), AuthSeqUpdateAppendInput{
		UserID:               1001,
		SourcePermAuthKeyID:  11,
		TargetPermAuthKeyIDs: []int64{22, 33},
		OperationID:          "op-auth-seq-1",
		UpdateType:           "updatePeerSettings",
		ReplayPolicy:         AuthSeqReplayPolicyDurableReplay,
		VisibilityPolicy:     AuthSeqVisibilityNotSourcePermAuthKey,
		Layer:                AuthSeqLayer,
		TLBytes:              body,
		PayloadHash:          hash,
		Now:                  1779234419,
	})
	if err != nil {
		t.Fatalf("AppendAuthSeqUpdate() error = %v", err)
	}
	if len(got.Deliveries) != 2 {
		t.Fatalf("deliveries = %d, want 2", len(got.Deliveries))
	}
	if got.Deliveries[0].PermAuthKeyID != 22 || got.Deliveries[1].PermAuthKeyID != 33 {
		t.Fatalf("delivery auth keys = %+v", got.Deliveries)
	}
	if got.Deliveries[0].Seq != 1 || got.Deliveries[1].Seq != 1 {
		t.Fatalf("delivery seqs = %+v, want per-key seq 1", got.Deliveries)
	}
}

func TestAppendAuthSeqUpdateIdempotentPerAuthKey(t *testing.T) {
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 301_000}, "local-userupdates")
	body := []byte{0x04, 0x05}
	hash := payload.HashBytes(body)
	in := AuthSeqUpdateAppendInput{
		UserID:               1002,
		TargetPermAuthKeyIDs: []int64{44},
		OperationID:          "op-auth-seq-idempotent",
		UpdateType:           "updatePeerSettings",
		ReplayPolicy:         AuthSeqReplayPolicyDurableReplay,
		VisibilityPolicy:     AuthSeqVisibilityAllUserAuthKeys,
		Layer:                AuthSeqLayer,
		TLBytes:              body,
		PayloadHash:          hash,
		Now:                  1779234420,
	}
	first, err := repo.AppendAuthSeqUpdate(context.Background(), in)
	if err != nil {
		t.Fatalf("first append error = %v", err)
	}
	second, err := repo.AppendAuthSeqUpdate(context.Background(), in)
	if err != nil {
		t.Fatalf("second append error = %v", err)
	}
	if !second.AlreadyApplied {
		t.Fatalf("AlreadyApplied = false, want true")
	}
	if first.Deliveries[0].Seq != second.Deliveries[0].Seq {
		t.Fatalf("seq changed from %d to %d", first.Deliveries[0].Seq, second.Deliveries[0].Seq)
	}
}
