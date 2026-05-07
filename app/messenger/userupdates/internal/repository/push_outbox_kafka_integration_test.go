//go:build integration && kafka

package repository

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func TestPushOutboxWorkerMarksRowPublished(t *testing.T) {
	if os.Getenv("TEAMGRAM_TEST_KAFKA_BROKERS") == "" {
		t.Skip("TEAMGRAM_TEST_KAFKA_BROKERS is empty")
	}
	ctx := context.Background()
	db := openIntegrationDB(t)
	base := time.Now().UnixNano() % 1_000_000_000
	repo := NewForTest(db, &testIDGenerator{next: base + 150_000}, "local-userupdates")
	taskID := base + 31
	insertPushTaskOutboxRow(t, repo, model.PushTaskOutbox{
		TaskId:            taskID,
		UserId:            base + 7001,
		Pts:               1,
		PushType:          PushTypeUserUpdate,
		PeerType:          payload.PeerTypeUser,
		PeerId:            base + 7002,
		OperationId:       "v1:test:push:integration",
		PushPartitionId:   12,
		TaskSchemaVersion: payload.PushTaskKafkaMessageSchemaVersion,
		TaskCodec:         PayloadCodecJSON,
		TaskPayload:       []byte(`{"schema_version":1}`),
		Status:            PushTaskStatusPending,
		AvailableAt:       mysqlTestTimeValue(time.Now().Add(-time.Minute)),
		NextRetryAt:       mysqlTestTimeValue(time.Now().Add(-time.Minute)),
		PublishedAt:       int64(0),
	})
	worker := NewPushOutboxWorker(repo, &fakePushPublisher{}, PushOutboxWorkerOptions{
		Interval:          time.Second,
		BatchSize:         10_000,
		PublishingTimeout: time.Minute,
	})

	worker.drain(ctx)

	row, err := repo.models.PushTaskOutboxModel.FindOne(ctx, taskID)
	if err != nil {
		t.Fatalf("FindOne() error = %v", err)
	}
	if row.Status != PushTaskStatusPublished || row.PublishedTopic == "" || row.PublishedOffset < 0 {
		t.Fatalf("push row not published: %+v", row)
	}
}
