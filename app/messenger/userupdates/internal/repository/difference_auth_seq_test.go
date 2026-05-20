//go:build integration

package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestGetStateUsesPerAuthKeySeqState(t *testing.T) {
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 302_000}, "local-userupdates")
	userID := base + 10_000
	operationID := fmt.Sprintf("op-auth-state-%d", base)
	body := []byte{1}
	hash := payload.HashBytes(body)
	if _, err := repo.AppendAuthSeqUpdate(context.Background(), AuthSeqUpdateAppendInput{
		UserID:               userID,
		TargetPermAuthKeyIDs: []int64{100},
		OperationID:          operationID,
		UpdateType:           "updatePeerSettings",
		ReplayPolicy:         AuthSeqReplayPolicyDurableReplay,
		VisibilityPolicy:     AuthSeqVisibilityAllUserAuthKeys,
		Layer:                AuthSeqLayer,
		TLBytes:              body,
		PayloadHash:          hash,
		Now:                  1779234419,
	}); err != nil {
		t.Fatalf("append auth seq error = %v", err)
	}
	state, err := repo.GetState(context.Background(), userID, 100)
	if err != nil {
		t.Fatalf("GetState() error = %v", err)
	}
	if state.Seq != 1 || state.Date != 1779234419 {
		t.Fatalf("state seq/date = %d/%d, want 1/1779234419", state.Seq, state.Date)
	}
	other, err := repo.GetState(context.Background(), userID, 200)
	if err != nil {
		t.Fatalf("GetState(other) error = %v", err)
	}
	if other.Seq != 0 || other.Date != 0 {
		t.Fatalf("other state seq/date = %d/%d, want zero", other.Seq, other.Date)
	}
}

func TestGetDifferenceReadsOnlyCurrentAuthKeyDeliveries(t *testing.T) {
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 303_000}, "local-userupdates")
	userID := base + 11_000
	operationID := fmt.Sprintf("op-auth-diff-%d", base)
	body := []byte{2}
	hash := payload.HashBytes(body)
	if _, err := repo.AppendAuthSeqUpdate(context.Background(), AuthSeqUpdateAppendInput{
		UserID:               userID,
		TargetPermAuthKeyIDs: []int64{100, 200},
		OperationID:          operationID,
		UpdateType:           "updatePeerSettings",
		ReplayPolicy:         AuthSeqReplayPolicyDurableReplay,
		VisibilityPolicy:     AuthSeqVisibilityAllUserAuthKeys,
		Layer:                AuthSeqLayer,
		TLBytes:              body,
		PayloadHash:          hash,
		Now:                  1779234420,
	}); err != nil {
		t.Fatalf("append auth seq error = %v", err)
	}
	date := int64(0)
	diff, err := repo.GetDifference(context.Background(), GetDifferenceInput{UserID: userID, PermAuthKeyID: 100, Date: &date})
	if err != nil {
		t.Fatalf("GetDifference() error = %v", err)
	}
	if len(diff.AuthSeqEvents) != 1 {
		t.Fatalf("auth seq events = %d, want 1", len(diff.AuthSeqEvents))
	}
	if diff.AuthSeqEvents[0].PermAuthKeyID != 100 {
		t.Fatalf("auth seq event auth key = %d, want 100", diff.AuthSeqEvents[0].PermAuthKeyID)
	}
	if diff.AuthSeqEvents[0].Date != 1779234420 {
		t.Fatalf("auth seq event date = %d, want 1779234420", diff.AuthSeqEvents[0].Date)
	}
}
