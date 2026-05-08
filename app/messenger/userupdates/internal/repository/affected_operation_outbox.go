package repository

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

type AffectedOutboxRow struct {
	OutboxID           int64
	UserID             int64
	RequesterUserID    int64
	OperationID        string
	OpType             int32
	OperationKind      string
	PeerType           int32
	PeerID             int64
	PayloadCodec       int32
	PayloadHash        []byte
	Payload            []byte
	DeliveryPolicy     int32
	Status             int32
	RetryCount         int32
	AvailableAt        int64
	ProcessingDeadline int64
	LastErrorCode      string
	LastErrorMessage   string
	BucketID           int32
	PartitionID        int32
	OwnerTokenPayload  []byte
}

func (r *Repository) insertAffectedOutboxesTx(ctx context.Context, txModels *model.TxModels, rows []AffectedOutbox) error {
	if len(rows) == 0 {
		return nil
	}
	for i := range rows {
		row := rows[i]
		if err := validateAffectedOutbox(row); err != nil {
			return err
		}
		outboxID, err := r.idgen.NextID(ctx)
		if err != nil {
			return storageError("next affected operation outbox id", err)
		}
		_, affected, err := txModels.AffectedOperationOutboxModel.InsertIgnore(&model.AffectedOperationOutbox{
			OutboxId:          outboxID,
			UserId:            row.TargetUserID,
			RequesterUserId:   row.RequesterUserID,
			OperationId:       row.OperationID,
			OpType:            row.OpType,
			OperationKind:     row.OperationKind,
			PeerType:          row.PeerType,
			PeerId:            row.PeerID,
			PayloadCodec:      row.PayloadCodec,
			PayloadHash:       row.PayloadHash,
			Payload:           row.Payload,
			DeliveryPolicy:    DeliveryPolicyDurableAsync,
			Status:            AffectedOutboxStatusPending,
			RetryCount:        0,
			AvailableAt:       unixNow(),
			BucketId:          row.TargetBucketID,
			PartitionId:       row.TargetPartitionID,
			OwnerTokenPayload: row.OwnerTokenPayload,
		})
		if err != nil {
			return storageError("insert affected operation outbox", err)
		}
		if affected != 0 {
			continue
		}
		existing, err := txModels.AffectedOperationOutboxModel.SelectByUserOperation(row.TargetUserID, row.OperationID)
		if err != nil {
			return storageError("select affected operation outbox duplicate", err)
		}
		if !bytes.Equal(existing.PayloadHash, row.PayloadHash) {
			return userupdates.ErrOperationPayloadConflict
		}
	}
	return nil
}

func (r *Repository) ListPendingAffectedOutboxes(ctx context.Context, now time.Time, limit int32) ([]AffectedOutboxRow, error) {
	if _, err := r.requireDB(); err != nil {
		return nil, err
	}
	rows, err := r.models.AffectedOperationOutboxModel.SelectPending(ctx, AffectedOutboxStatusPending, now.UTC().Unix(), limit)
	if err != nil {
		return nil, storageError("list pending affected operation outboxes", err)
	}
	out := make([]AffectedOutboxRow, 0, len(rows))
	for _, row := range rows {
		out = append(out, affectedOutboxRowFromModel(row))
	}
	return out, nil
}

func (r *Repository) TryMarkAffectedOutboxProcessing(ctx context.Context, outboxID int64, now time.Time, deadline time.Time) (bool, error) {
	if _, err := r.requireDB(); err != nil {
		return false, err
	}
	rows, err := r.models.AffectedOperationOutboxModel.TryMarkProcessing(
		ctx,
		AffectedOutboxStatusProcessing,
		deadline.UTC().Unix(),
		outboxID,
		AffectedOutboxStatusPending,
		now.UTC().Unix(),
	)
	if err != nil {
		return false, storageError("mark affected operation outbox processing", err)
	}
	return rows == 1, nil
}

func (r *Repository) MarkAffectedOutboxCompleted(ctx context.Context, outboxID int64, processingDeadline time.Time) error {
	if _, err := r.requireDB(); err != nil {
		return err
	}
	rows, err := r.models.AffectedOperationOutboxModel.MarkCompleted(
		ctx,
		AffectedOutboxStatusCompleted,
		outboxID,
		AffectedOutboxStatusProcessing,
		processingDeadline.UTC().Unix(),
	)
	if err != nil {
		return storageError("mark affected operation outbox completed", err)
	}
	if rows == 0 {
		return storageError("mark affected operation outbox completed", model.ErrNotFound)
	}
	return nil
}

func (r *Repository) MarkAffectedOutboxRetryable(ctx context.Context, outboxID int64, processingDeadline time.Time, code string, message string, nextAvailableAt time.Time) error {
	if _, err := r.requireDB(); err != nil {
		return err
	}
	rows, err := r.models.AffectedOperationOutboxModel.MarkRetryable(
		ctx,
		AffectedOutboxStatusPending,
		nextAvailableAt.UTC().Unix(),
		code,
		message,
		outboxID,
		AffectedOutboxStatusProcessing,
		processingDeadline.UTC().Unix(),
	)
	if err != nil {
		return storageError("mark affected operation outbox retryable", err)
	}
	if rows == 0 {
		return storageError("mark affected operation outbox retryable", model.ErrNotFound)
	}
	return nil
}

func (r *Repository) MarkAffectedOutboxFailedTerminal(ctx context.Context, outboxID int64, processingDeadline time.Time, code string, message string) error {
	if _, err := r.requireDB(); err != nil {
		return err
	}
	rows, err := r.models.AffectedOperationOutboxModel.MarkFailedTerminal(
		ctx,
		AffectedOutboxStatusFailedTerminal,
		code,
		message,
		outboxID,
		AffectedOutboxStatusProcessing,
		processingDeadline.UTC().Unix(),
	)
	if err != nil {
		return storageError("mark affected operation outbox failed terminal", err)
	}
	if rows == 0 {
		return storageError("mark affected operation outbox failed terminal", model.ErrNotFound)
	}
	return nil
}

func (r *Repository) ResetExpiredAffectedOutboxes(ctx context.Context, now time.Time, limit int32) (int64, error) {
	if _, err := r.requireDB(); err != nil {
		return 0, err
	}
	rows, err := r.models.AffectedOperationOutboxModel.ResetExpiredProcessing(
		ctx,
		AffectedOutboxStatusPending,
		now.UTC().Unix(),
		AffectedOutboxStatusProcessing,
		limit,
	)
	if err != nil {
		return 0, storageError("reset expired affected operation outboxes", err)
	}
	return rows, nil
}

func affectedOutboxRowFromModel(row model.AffectedOperationOutbox) AffectedOutboxRow {
	return AffectedOutboxRow{
		OutboxID:           row.OutboxId,
		UserID:             row.UserId,
		RequesterUserID:    row.RequesterUserId,
		OperationID:        row.OperationId,
		OpType:             row.OpType,
		OperationKind:      row.OperationKind,
		PeerType:           row.PeerType,
		PeerID:             row.PeerId,
		PayloadCodec:       row.PayloadCodec,
		PayloadHash:        row.PayloadHash,
		Payload:            row.Payload,
		DeliveryPolicy:     row.DeliveryPolicy,
		Status:             row.Status,
		RetryCount:         row.RetryCount,
		AvailableAt:        row.AvailableAt,
		ProcessingDeadline: row.ProcessingDeadline,
		LastErrorCode:      row.LastErrorCode,
		LastErrorMessage:   row.LastErrorMessage,
		BucketID:           row.BucketId,
		PartitionID:        row.PartitionId,
		OwnerTokenPayload:  row.OwnerTokenPayload,
	}
}

func validateAffectedOutbox(row AffectedOutbox) error {
	switch {
	case row.TargetUserID <= 0:
		return fmt.Errorf("%w: affected outbox target user id is required", userupdates.ErrOperationTerminal)
	case row.OperationID == "":
		return fmt.Errorf("%w: affected outbox operation id is required", userupdates.ErrOperationTerminal)
	case row.OperationKind == "":
		return fmt.Errorf("%w: affected outbox operation kind is required", userupdates.ErrOperationTerminal)
	case row.DeliveryPolicy != DeliveryPolicyDurableAsync:
		return fmt.Errorf("%w: affected outbox delivery policy=%d", userupdates.ErrOperationTerminal, row.DeliveryPolicy)
	case row.PayloadCodec != PayloadCodecJSON:
		return fmt.Errorf("%w: affected outbox payload codec=%d", userupdates.ErrOperationTerminal, row.PayloadCodec)
	case !bytes.Equal(row.PayloadHash, payload.HashBytes(row.Payload)):
		return userupdates.ErrOperationPayloadConflict
	default:
		return nil
	}
}
