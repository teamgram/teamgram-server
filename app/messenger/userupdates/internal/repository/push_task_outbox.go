package repository

import (
	"context"
	"errors"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

func (r *Repository) ListPendingPushTasks(ctx context.Context, now time.Time, limit int32) ([]PushTask, error) {
	rows, err := r.models.PushTaskOutboxModel.SelectDueForPublish(
		ctx,
		PushTaskStatusPending,
		PushTaskStatusFailedRetryable,
		mysqlTimestamp(now),
		limit,
	)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return []PushTask{}, nil
		}
		return nil, storageError("list pending push tasks", err)
	}
	out := make([]PushTask, 0, len(rows))
	for _, row := range rows {
		out = append(out, PushTask{
			TaskID:          row.TaskId,
			UserID:          row.UserId,
			Pts:             row.Pts,
			PushType:        row.PushType,
			PeerType:        row.PeerType,
			PeerID:          row.PeerId,
			OperationID:     row.OperationId,
			PushPartitionID: row.PushPartitionId,
			TaskPayload:     row.TaskPayload,
		})
	}
	return out, nil
}

func (r *Repository) TryMarkPushTaskPublishing(ctx context.Context, taskID int64, now time.Time, leaseExpiresAt time.Time) (bool, error) {
	rows, err := r.models.PushTaskOutboxModel.TryMarkPublishingDue(
		ctx,
		PushTaskStatusPublishing,
		mysqlTimestamp(leaseExpiresAt),
		taskID,
		PushTaskStatusPending,
		PushTaskStatusFailedRetryable,
		mysqlTimestamp(now),
	)
	if err != nil {
		return false, storageError("mark push task publishing", err)
	}
	return rows == 1, nil
}

func (r *Repository) MarkPushTaskPublished(ctx context.Context, taskID int64, ack KafkaAck) error {
	rows, err := r.models.PushTaskOutboxModel.MarkPublished(ctx, PushTaskStatusPublished, ack.Topic, ack.Partition, ack.Offset, mysqlNow(), taskID)
	if err != nil {
		return storageError("mark push task published", err)
	}
	if rows == 0 {
		return storageError("mark push task published", sqlx.ErrNotFound)
	}
	return nil
}

func (r *Repository) MarkPushTaskPublishFailed(ctx context.Context, taskID int64, code string, nextRetryAt time.Time) error {
	next := mysqlTimestamp(nextRetryAt)
	rows, err := r.models.PushTaskOutboxModel.MarkPublishFailed(ctx, PushTaskStatusFailedRetryable, next, next, code, taskID)
	if err != nil {
		return storageError("mark push task publish failed", err)
	}
	if rows == 0 {
		return storageError("mark push task publish failed", sqlx.ErrNotFound)
	}
	return nil
}

func (r *Repository) ResetExpiredPublishingTasks(ctx context.Context, now time.Time, limit int32) (int64, error) {
	nowString := mysqlTimestamp(now)
	rows, err := r.models.PushTaskOutboxModel.ResetExpiredPublishing(
		ctx,
		PushTaskStatusPending,
		nowString,
		PushTaskStatusPublishing,
		nowString,
		limit,
	)
	if err != nil {
		return 0, storageError("reset stale publishing push tasks", err)
	}
	return rows, nil
}

func mysqlTimestamp(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05.000000")
}
