//go:build integration

package repository

import (
	"context"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestListPendingPushTasksReturnsDueRows(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 110_000}, "local-userupdates")
	now := time.Now().UTC()
	dueID := base + 1
	futureID := base + 2

	insertPushTaskOutboxRow(t, repo, model.PushTaskOutbox{
		TaskId:            dueID,
		UserId:            base + 1001,
		Pts:               1,
		PushType:          PushTypeUserUpdate,
		PeerType:          payload.PeerTypeUser,
		PeerId:            base + 2001,
		OperationId:       "v1:test:push:due",
		PushPartitionId:   7,
		TaskSchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskCodec:         PayloadCodecJSON,
		TaskPayload:       []byte(`{"schema_version":1}`),
		Status:            PushTaskStatusPending,
		AvailableAt:       mysqlTestTime(now.Add(-time.Minute)),
		NextRetryAt:       mysqlTestTime(now.Add(-time.Minute)),
		PublishedAt:       mysqlTestTime(time.Unix(0, 0).UTC()),
	})
	insertPushTaskOutboxRow(t, repo, model.PushTaskOutbox{
		TaskId:            futureID,
		UserId:            base + 1002,
		Pts:               1,
		PushType:          PushTypeUserUpdate,
		PeerType:          payload.PeerTypeUser,
		PeerId:            base + 2002,
		OperationId:       "v1:test:push:future",
		PushPartitionId:   8,
		TaskSchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskCodec:         PayloadCodecJSON,
		TaskPayload:       []byte(`{"schema_version":1}`),
		Status:            PushTaskStatusPending,
		AvailableAt:       mysqlTestTime(now.Add(time.Hour)),
		NextRetryAt:       mysqlTestTime(now.Add(time.Hour)),
		PublishedAt:       mysqlTestTime(time.Unix(0, 0).UTC()),
	})

	tasks, err := repo.ListPendingPushTasks(ctx, now, 10_000)
	if err != nil {
		t.Fatalf("ListPendingPushTasks() error = %v", err)
	}
	if !containsPushTask(tasks, dueID) {
		t.Fatalf("due task not returned: %+v", tasks)
	}
	if containsPushTask(tasks, futureID) {
		t.Fatalf("future task should not be returned: %+v", tasks)
	}
}

func TestTryMarkPushTaskPublishingClaimsOnce(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 120_000}, "local-userupdates")
	taskID := base + 11
	insertPushTaskOutboxRow(t, repo, model.PushTaskOutbox{
		TaskId:            taskID,
		UserId:            base + 1101,
		Pts:               1,
		PushType:          PushTypeUserUpdate,
		PeerType:          payload.PeerTypeUser,
		PeerId:            base + 2101,
		OperationId:       "v1:test:push:claim",
		PushPartitionId:   9,
		TaskSchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskCodec:         PayloadCodecJSON,
		TaskPayload:       []byte(`{"schema_version":1}`),
		Status:            PushTaskStatusPending,
		AvailableAt:       mysqlTestTime(time.Now().Add(-time.Minute)),
		NextRetryAt:       mysqlTestTime(time.Now().Add(-time.Minute)),
		PublishedAt:       mysqlTestTime(time.Unix(0, 0).UTC()),
	})

	now := time.Now().UTC()
	leaseExpiresAt := now.Add(time.Minute)
	claimed, err := repo.TryMarkPushTaskPublishing(ctx, taskID, now, leaseExpiresAt)
	if err != nil {
		t.Fatalf("TryMarkPushTaskPublishing() error = %v", err)
	}
	if !claimed {
		t.Fatal("first claim = false, want true")
	}
	claimed, err = repo.TryMarkPushTaskPublishing(ctx, taskID, now.Add(time.Second), leaseExpiresAt.Add(time.Second))
	if err != nil {
		t.Fatalf("TryMarkPushTaskPublishing() second error = %v", err)
	}
	if claimed {
		t.Fatal("second claim = true, want false")
	}
	row, err := repo.models.PushTaskOutboxModel.FindOne(ctx, taskID)
	if err != nil {
		t.Fatalf("FindOne() error = %v", err)
	}
	if row.Status != PushTaskStatusPublishing {
		t.Fatalf("status = %d, want publishing", row.Status)
	}
	if normalizeDBTestTime(t, row.AvailableAt) != mysqlTestTime(leaseExpiresAt) {
		t.Fatalf("available_at lease = %q, want %q", row.AvailableAt, mysqlTestTime(leaseExpiresAt))
	}
}

func TestResetExpiredPublishingTasksMovesRowsBackToPending(t *testing.T) {
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 130_000}, "local-userupdates")
	taskID := base + 21
	insertPushTaskOutboxRow(t, repo, model.PushTaskOutbox{
		TaskId:            taskID,
		UserId:            base + 1201,
		Pts:               1,
		PushType:          PushTypeUserUpdate,
		PeerType:          payload.PeerTypeUser,
		PeerId:            base + 2201,
		OperationId:       "v1:test:push:stale",
		PushPartitionId:   10,
		TaskSchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskCodec:         PayloadCodecJSON,
		TaskPayload:       []byte(`{"schema_version":1}`),
		Status:            PushTaskStatusPublishing,
		AvailableAt:       mysqlTestTime(time.Now().Add(-2 * time.Hour)),
		NextRetryAt:       mysqlTestTime(time.Now().Add(time.Hour)),
		PublishedAt:       mysqlTestTime(time.Unix(0, 0).UTC()),
	})

	now := time.Now().UTC()
	count, err := repo.ResetExpiredPublishingTasks(ctx, now, 10)
	if err != nil {
		t.Fatalf("ResetExpiredPublishingTasks() error = %v", err)
	}
	if count == 0 {
		t.Fatal("reset count = 0, want > 0")
	}
	row, err := repo.models.PushTaskOutboxModel.FindOne(ctx, taskID)
	if err != nil {
		t.Fatalf("FindOne() error = %v", err)
	}
	if row.Status != PushTaskStatusPending {
		t.Fatalf("status = %d, want pending", row.Status)
	}
}

func insertPushTaskOutboxRow(t *testing.T, repo *Repository, row model.PushTaskOutbox) {
	t.Helper()
	_, err := repo.db.Exec(context.Background(), `
INSERT INTO push_task_outbox
	(task_id, user_id, pts, push_type, peer_type, peer_id, operation_id,
	 push_partition_id, task_schema_version, task_codec, task_payload, status,
	 publish_attempts, available_at, next_retry_at, published_topic,
	 published_partition, published_offset, last_error_code, published_at)
VALUES
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		row.TaskId,
		row.UserId,
		row.Pts,
		row.PushType,
		row.PeerType,
		row.PeerId,
		row.OperationId,
		row.PushPartitionId,
		row.TaskSchemaVersion,
		row.TaskCodec,
		row.TaskPayload,
		row.Status,
		row.PublishAttempts,
		row.AvailableAt,
		row.NextRetryAt,
		row.PublishedTopic,
		row.PublishedPartition,
		row.PublishedOffset,
		row.LastErrorCode,
		row.PublishedAt,
	)
	if err != nil {
		t.Fatalf("insert push_task_outbox row: %v", err)
	}
}

func containsPushTask(tasks []PushTask, taskID int64) bool {
	for _, task := range tasks {
		if task.TaskID == taskID {
			return true
		}
	}
	return false
}
