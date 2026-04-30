package repository

import (
	"context"
	"errors"
	"testing"
	"time"
)

type fakePushTaskStore struct {
	tasks       []PushTask
	claimed     map[int64]bool
	published   []int64
	failed      []int64
	resetCalls  int
	resetCount  int64
	claimResult bool
}

func (s *fakePushTaskStore) ListPendingPushTasks(ctx context.Context, now time.Time, limit int32) ([]PushTask, error) {
	return s.tasks, nil
}

func (s *fakePushTaskStore) TryMarkPushTaskPublishing(ctx context.Context, taskID int64, now time.Time, leaseExpiresAt time.Time) (bool, error) {
	if s.claimed == nil {
		s.claimed = map[int64]bool{}
	}
	if s.claimed[taskID] {
		return false, nil
	}
	s.claimed[taskID] = true
	if s.claimResult {
		return true, nil
	}
	return true, nil
}

func (s *fakePushTaskStore) MarkPushTaskPublished(ctx context.Context, taskID int64, ack KafkaAck) error {
	s.published = append(s.published, taskID)
	return nil
}

func (s *fakePushTaskStore) MarkPushTaskPublishFailed(ctx context.Context, taskID int64, code string, nextRetryAt time.Time) error {
	s.failed = append(s.failed, taskID)
	return nil
}

func (s *fakePushTaskStore) ResetExpiredPublishingTasks(ctx context.Context, now time.Time, limit int32) (int64, error) {
	s.resetCalls++
	return s.resetCount, nil
}

type fakePushPublisher struct {
	err   error
	calls int
}

func (p *fakePushPublisher) Publish(ctx context.Context, task PushTask) (KafkaAck, error) {
	p.calls++
	if p.err != nil {
		return KafkaAck{}, p.err
	}
	return KafkaAck{Topic: "push", Partition: task.PushPartitionID, Offset: 88}, nil
}

func TestPushOutboxWorkerPublishesPendingTask(t *testing.T) {
	store := &fakePushTaskStore{tasks: []PushTask{{TaskID: 1, UserID: 10, Pts: 1, OperationID: "op", PushPartitionID: 3}}}
	publisher := &fakePushPublisher{}
	worker := NewPushOutboxWorker(store, publisher, PushOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, PublishingTimeout: time.Minute})

	worker.drain(context.Background())

	if publisher.calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", publisher.calls)
	}
	if len(store.published) != 1 || store.published[0] != 1 {
		t.Fatalf("published = %+v", store.published)
	}
	if len(store.failed) != 0 {
		t.Fatalf("failed = %+v", store.failed)
	}
}

func TestPushOutboxWorkerMarksRetryableOnPublishFailure(t *testing.T) {
	store := &fakePushTaskStore{tasks: []PushTask{{TaskID: 2, UserID: 10, Pts: 2, OperationID: "op2", PushPartitionID: 4}}}
	publisher := &fakePushPublisher{err: errors.New("kafka down")}
	worker := NewPushOutboxWorker(store, publisher, PushOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, PublishingTimeout: time.Minute})

	worker.drain(context.Background())

	if len(store.published) != 0 {
		t.Fatalf("published = %+v", store.published)
	}
	if len(store.failed) != 1 || store.failed[0] != 2 {
		t.Fatalf("failed = %+v", store.failed)
	}
}

func TestPushOutboxWorkerRepublishesPendingTaskAfterRestart(t *testing.T) {
	store := &fakePushTaskStore{tasks: []PushTask{{TaskID: 3, UserID: 10, Pts: 3, OperationID: "op3", PushPartitionID: 5}}}
	publisher := &fakePushPublisher{}
	worker := NewPushOutboxWorker(store, publisher, PushOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, PublishingTimeout: time.Minute})

	worker.drain(context.Background())

	if len(store.published) != 1 || store.published[0] != 3 {
		t.Fatalf("published = %+v", store.published)
	}
}

func TestPushOutboxWorkerResetsExpiredPublishingTask(t *testing.T) {
	store := &fakePushTaskStore{resetCount: 2}
	publisher := &fakePushPublisher{}
	worker := NewPushOutboxWorker(store, publisher, PushOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, PublishingTimeout: time.Minute})

	worker.drain(context.Background())

	if store.resetCalls != 1 {
		t.Fatalf("reset calls = %d, want 1", store.resetCalls)
	}
}
