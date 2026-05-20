//go:build integration

package repository

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func TestAppendAuthSeqUpdateExpandsPerAuthKey(t *testing.T) {
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 300_000}, "local-userupdates")
	userID := base + 1001
	operationID := fmt.Sprintf("op-auth-seq-1-%d", base)
	body := []byte{0x01, 0x02, 0x03}
	hash := payload.HashBytes(body)

	got, err := repo.AppendAuthSeqUpdate(context.Background(), AuthSeqUpdateAppendInput{
		UserID:               userID,
		SourcePermAuthKeyID:  11,
		TargetPermAuthKeyIDs: []int64{22, 33},
		OperationID:          operationID,
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
	userID := base + 1002
	operationID := fmt.Sprintf("op-auth-seq-idempotent-%d", base)
	body := []byte{0x04, 0x05}
	hash := payload.HashBytes(body)
	in := AuthSeqUpdateAppendInput{
		UserID:               userID,
		TargetPermAuthKeyIDs: []int64{44},
		OperationID:          operationID,
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

func TestAppendAuthSeqUpdateIdempotentPayloadConflictDoesNotAdvanceSeq(t *testing.T) {
	db := openIntegrationDB(t)
	ctx := context.Background()
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 302_000}, "local-userupdates")
	userID := base + 1003
	authKeyID := int64(55)
	operationID := fmt.Sprintf("op-auth-seq-conflict-%d", base)
	body := []byte{0x06, 0x07}
	hash := payload.HashBytes(body)
	in := AuthSeqUpdateAppendInput{
		UserID:               userID,
		TargetPermAuthKeyIDs: []int64{authKeyID},
		OperationID:          operationID,
		UpdateType:           "updatePeerSettings",
		ReplayPolicy:         AuthSeqReplayPolicyDurableReplay,
		VisibilityPolicy:     AuthSeqVisibilityAllUserAuthKeys,
		Layer:                AuthSeqLayer,
		TLBytes:              body,
		PayloadHash:          hash,
		Now:                  1779234421,
	}
	if _, err := repo.AppendAuthSeqUpdate(ctx, in); err != nil {
		t.Fatalf("first append error = %v", err)
	}
	stateBefore, err := repo.models.AuthSeqStateModel.SelectByUserAuthKey(ctx, userID, authKeyID)
	if err != nil {
		t.Fatalf("SelectByUserAuthKey(before) error = %v", err)
	}

	conflictBody := []byte{0x08, 0x09}
	in.TLBytes = conflictBody
	in.PayloadHash = payload.HashBytes(conflictBody)
	in.Now++
	_, err = repo.AppendAuthSeqUpdate(ctx, in)
	if !errors.Is(err, userupdates.ErrOperationPayloadConflict) {
		t.Fatalf("conflict append error = %v, want ErrOperationPayloadConflict", err)
	}
	stateAfter, err := repo.models.AuthSeqStateModel.SelectByUserAuthKey(ctx, userID, authKeyID)
	if err != nil {
		t.Fatalf("SelectByUserAuthKey(after) error = %v", err)
	}
	if stateAfter.Seq != stateBefore.Seq {
		t.Fatalf("seq changed from %d to %d", stateBefore.Seq, stateAfter.Seq)
	}
}
