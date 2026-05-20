//go:build integration

package repository

import (
	"context"
	"database/sql"
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
	body := []byte{1, byte(base)}
	hash := payload.HashBytes(body)
	cleanupAuthSeqDifferenceRows(t, context.Background(), db, userID, AuthSeqPayloadID(hash))
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
	body := []byte{2, byte(base)}
	hash := payload.HashBytes(body)
	cleanupAuthSeqDifferenceRows(t, context.Background(), db, userID, AuthSeqPayloadID(hash))
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

func TestGetDifferenceLimitedAuthSeqRowsReturnsDeliveredCursor(t *testing.T) {
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 304_000}, "local-userupdates")
	userID := base + 12_000
	body1 := []byte{3, byte(base)}
	body2 := []byte{4, byte(base)}
	hash1 := payload.HashBytes(body1)
	hash2 := payload.HashBytes(body2)
	cleanupAuthSeqDifferenceRows(t, context.Background(), db, userID, AuthSeqPayloadID(hash1), AuthSeqPayloadID(hash2))

	if _, err := repo.AppendAuthSeqUpdate(context.Background(), AuthSeqUpdateAppendInput{
		UserID:               userID,
		TargetPermAuthKeyIDs: []int64{100},
		OperationID:          fmt.Sprintf("op-auth-diff-limit-1-%d", base),
		UpdateType:           "updatePeerSettings",
		ReplayPolicy:         AuthSeqReplayPolicyDurableReplay,
		VisibilityPolicy:     AuthSeqVisibilityAllUserAuthKeys,
		Layer:                AuthSeqLayer,
		TLBytes:              body1,
		PayloadHash:          hash1,
		Now:                  1779234430,
	}); err != nil {
		t.Fatalf("append first auth seq error = %v", err)
	}
	if _, err := repo.AppendAuthSeqUpdate(context.Background(), AuthSeqUpdateAppendInput{
		UserID:               userID,
		TargetPermAuthKeyIDs: []int64{100},
		OperationID:          fmt.Sprintf("op-auth-diff-limit-2-%d", base),
		UpdateType:           "updatePeerSettings",
		ReplayPolicy:         AuthSeqReplayPolicyDurableReplay,
		VisibilityPolicy:     AuthSeqVisibilityAllUserAuthKeys,
		Layer:                AuthSeqLayer,
		TLBytes:              body2,
		PayloadHash:          hash2,
		Now:                  1779234431,
	}); err != nil {
		t.Fatalf("append second auth seq error = %v", err)
	}
	date := int64(0)
	diff, err := repo.GetDifference(context.Background(), GetDifferenceInput{UserID: userID, PermAuthKeyID: 100, Date: &date, Limit: 1})
	if err != nil {
		t.Fatalf("GetDifference() error = %v", err)
	}
	if len(diff.AuthSeqEvents) != 1 {
		t.Fatalf("auth seq events = %d, want 1", len(diff.AuthSeqEvents))
	}
	event := diff.AuthSeqEvents[0]
	if event.Seq != 1 || event.Date != 1779234430 {
		t.Fatalf("delivered event seq/date = %d/%d, want 1/1779234430", event.Seq, event.Date)
	}
	if diff.State.Seq != event.Seq || diff.State.Date != event.Date {
		t.Fatalf("state seq/date = %d/%d, want delivered %d/%d", diff.State.Seq, diff.State.Date, event.Seq, event.Date)
	}
}

func cleanupAuthSeqDifferenceRows(t *testing.T, ctx context.Context, db interface {
	Exec(context.Context, string, ...interface{}) (sql.Result, error)
}, userID int64, payloadIDs ...string) {
	t.Helper()
	t.Cleanup(func() {
		for _, statement := range []string{
			"DELETE FROM auth_seq_deliveries WHERE user_id = ?",
			"DELETE FROM auth_seq_state WHERE user_id = ?",
		} {
			if _, err := db.Exec(ctx, statement, userID); err != nil {
				t.Fatalf("cleanup %q: %v", statement, err)
			}
		}
		for _, payloadID := range payloadIDs {
			if payloadID == "" {
				continue
			}
			if _, err := db.Exec(ctx, "DELETE FROM auth_update_payloads WHERE payload_id = ?", payloadID); err != nil {
				t.Fatalf("cleanup auth_update_payloads: %v", err)
			}
		}
	})
}
