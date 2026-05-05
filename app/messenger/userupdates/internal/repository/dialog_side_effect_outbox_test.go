//go:build integration

package repository

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestApplySendMessageClearDraftWritesSideEffectInFinalTransaction(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 10_001
	peerID := base + 10_002
	sourcePermAuthKeyID := base + 99_001
	repo := NewForTest(db, &testIDGenerator{next: base + 100_000}, "local-userupdates")
	in := buildApplyInput(t, userID, userID, peerID, true, "clear draft")
	op := decodeApplyOperation(t, in)
	op.ClearDraft = true
	op.SourcePermAuthKeyID = sourcePermAuthKeyID
	op.ClearDraftBeforeDate = op.Date - 10
	in = encodeApplyOperation(t, in, op)

	if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, in); err != nil {
		t.Fatalf("ApplyUserOperation() error = %v", err)
	}

	rows, err := repo.models.DialogSideEffectOutboxModel.SelectBySourceOperationKind(
		ctx,
		in.OperationID,
		DialogSideEffectKindClearDraftAfterSend,
	)
	if err != nil {
		t.Fatalf("SelectBySourceOperationKind() error = %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("side effect rows = %d, want 1: %+v", len(rows), rows)
	}
	got := rows[0]
	if got.UserId != userID || got.PeerType != payload.PeerTypeUser || got.PeerId != peerID {
		t.Fatalf("side effect peer mismatch: %+v", got)
	}
	if got.SourcePermAuthKeyId != sourcePermAuthKeyID || got.SourceOperationId != in.OperationID {
		t.Fatalf("side effect source mismatch: %+v", got)
	}
	if normalizeDBTestTime(t, got.ClearBeforeDate) != mysqlTestTime(time.Unix(int64(op.ClearDraftBeforeDate), 0).UTC()) {
		t.Fatalf("clear_before_date = %q, want operation clear boundary", got.ClearBeforeDate)
	}
	if got.Status != DialogSideEffectStatusPending || got.AttemptCount != 0 {
		t.Fatalf("side effect status mismatch: %+v", got)
	}
}

func TestApplySavedMessageWritesSavedDialogTopSideEffect(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	userID := base + 20_001
	repo := NewForTest(db, &testIDGenerator{next: base + 200_000}, "local-userupdates")
	in := buildApplyInput(t, userID, userID, userID, true, "saved message")
	op := decodeApplyOperation(t, in)
	op.PeerSeq = 8
	op.CanonicalMessageID = userID*10 + 88
	in.OperationID = payload.SenderOperationID(op.CanonicalMessageID, userID)
	in = encodeApplyOperation(t, in, op)

	if _, err := repo.ClaimPartitionOwner(ctx, in.PartitionID); err != nil {
		t.Fatalf("ClaimPartitionOwner() error = %v", err)
	}
	if _, err := repo.ApplyUserOperation(ctx, in); err != nil {
		t.Fatalf("ApplyUserOperation() error = %v", err)
	}

	rows, err := repo.models.DialogSideEffectOutboxModel.SelectBySourceOperationKind(
		ctx,
		in.OperationID,
		DialogSideEffectKindUpsertSavedDialogFromMessage,
	)
	if err != nil {
		t.Fatalf("SelectBySourceOperationKind() error = %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("side effect rows = %d, want 1: %+v", len(rows), rows)
	}
	got := rows[0]
	if got.SourcePeerSeq != op.PeerSeq || got.SourceCanonicalMessageId != op.CanonicalMessageID {
		t.Fatalf("saved top source mismatch: %+v", got)
	}
	if normalizeDBTestTime(t, got.SourceMessageDate) != mysqlTestTime(time.Unix(int64(op.Date), 0).UTC()) {
		t.Fatalf("source_message_date = %q, want %q", got.SourceMessageDate, mysqlTestTime(time.Unix(int64(op.Date), 0).UTC()))
	}
	var saved savedDialogSideEffectPayloadV1
	if err := json.Unmarshal(got.Payload, &saved); err != nil {
		t.Fatalf("unmarshal saved payload: %v", err)
	}
	if saved.TopPeerSeq != op.PeerSeq || saved.TopCanonicalMessageID != op.CanonicalMessageID || !saved.Top {
		t.Fatalf("saved payload mismatch: %+v", saved)
	}
}

func TestDialogSideEffectOutboxRetryMovesToBlocked(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 300_000}, "local-userupdates")
	now := time.Now().UTC().Truncate(time.Second)
	row := DialogSideEffect{
		SideEffectID:             base + 300_001,
		Kind:                     DialogSideEffectKindClearDraftAfterSend,
		UserID:                   base + 30_001,
		PeerType:                 payload.PeerTypeUser,
		PeerID:                   base + 30_002,
		SourcePermAuthKeyID:      base + 30_003,
		SourceOperationID:        "test-dialog-side-effect-blocked-" + time.Now().UTC().Format("150405.000000000"),
		SourceMessageDate:        now,
		SourcePeerSeq:            7,
		SourceCanonicalMessageID: base + 30_004,
		ClearBeforeDate:          now,
		PayloadSchemaVersion:     1,
		Payload:                  []byte(`{"schema_version":1}`),
		PayloadHash:              payload.HashBytes([]byte(`{"schema_version":1}`)),
		Status:                   DialogSideEffectStatusPending,
		AttemptCount:             BlockedAfterAttempts - 1,
		NextRetryAt:              now.Add(-time.Minute),
	}
	if err := db.Transact(ctx, func(tx *sqlx.Tx) error {
		return repo.InsertDialogSideEffectTx(repo.models.WithTx(tx), row)
	}); err != nil {
		t.Fatalf("InsertDialogSideEffectTx() error = %v", err)
	}

	if err := repo.MarkDialogSideEffectRetryableFailure(ctx, row.SideEffectID, "dialog_unavailable", now); err != nil {
		t.Fatalf("MarkDialogSideEffectRetryableFailure() error = %v", err)
	}

	got, err := repo.models.DialogSideEffectOutboxModel.SelectOne(ctx, row.SideEffectID)
	if err != nil {
		t.Fatalf("SelectOne() error = %v", err)
	}
	if got.Status != DialogSideEffectStatusBlocked {
		t.Fatalf("status = %d, want blocked: %+v", got.Status, got)
	}
	if got.LastErrorCode != "dialog_unavailable" {
		t.Fatalf("last_error_code = %q, want dialog_unavailable", got.LastErrorCode)
	}
}

func decodeApplyOperation(t *testing.T, in ApplyUserOperationInput) payload.MessageOperationV1 {
	t.Helper()
	var op payload.MessageOperationV1
	if err := json.Unmarshal(in.Payload, &op); err != nil {
		t.Fatalf("unmarshal operation: %v", err)
	}
	return op
}

func encodeApplyOperation(t *testing.T, in ApplyUserOperationInput, op payload.MessageOperationV1) ApplyUserOperationInput {
	t.Helper()
	body, err := json.Marshal(op)
	if err != nil {
		t.Fatalf("marshal operation: %v", err)
	}
	in.PeerType = op.PeerType
	in.PeerID = op.PeerID
	in.Payload = body
	in.PayloadHash = payload.HashBytes(body)
	return in
}

var _ = model.DialogSideEffectOutbox{}
