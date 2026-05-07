package repository

import (
	"context"
	"errors"
	"sync/atomic"
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
	calls atomic.Int32
}

func (p *fakePushPublisher) Publish(ctx context.Context, task PushTask) (KafkaAck, error) {
	p.calls.Add(1)
	if p.err != nil {
		return KafkaAck{}, p.err
	}
	return KafkaAck{Topic: "push", Partition: task.PushPartitionID, Offset: 88}, nil
}

type fakePushOutboxWorkerCounters struct {
	drainStart        int
	resetStaleSuccess int
	resetStaleError   int
	claimEmpty        int
	claimSuccess      int
	publishSuccess    int
	publishError      int
	markPublishedErr  int
	markRetryableErr  int
	panicRecovered    int
}

func (c *fakePushOutboxWorkerCounters) IncDrainStart()         { c.drainStart++ }
func (c *fakePushOutboxWorkerCounters) IncResetStaleSuccess()  { c.resetStaleSuccess++ }
func (c *fakePushOutboxWorkerCounters) IncResetStaleError()    { c.resetStaleError++ }
func (c *fakePushOutboxWorkerCounters) IncClaimEmpty()         { c.claimEmpty++ }
func (c *fakePushOutboxWorkerCounters) IncClaimSuccess()       { c.claimSuccess++ }
func (c *fakePushOutboxWorkerCounters) IncPublishSuccess()     { c.publishSuccess++ }
func (c *fakePushOutboxWorkerCounters) IncPublishError()       { c.publishError++ }
func (c *fakePushOutboxWorkerCounters) IncMarkPublishedError() { c.markPublishedErr++ }
func (c *fakePushOutboxWorkerCounters) IncMarkRetryableError() { c.markRetryableErr++ }
func (c *fakePushOutboxWorkerCounters) IncPanicRecovered()     { c.panicRecovered++ }

type panicPushTaskStore struct{}

func (panicPushTaskStore) ListPendingPushTasks(context.Context, time.Time, int32) ([]PushTask, error) {
	panic("list pending panic")
}
func (panicPushTaskStore) TryMarkPushTaskPublishing(context.Context, int64, time.Time, time.Time) (bool, error) {
	return false, nil
}
func (panicPushTaskStore) MarkPushTaskPublished(context.Context, int64, KafkaAck) error { return nil }
func (panicPushTaskStore) MarkPushTaskPublishFailed(context.Context, int64, string, time.Time) error {
	return nil
}
func (panicPushTaskStore) ResetExpiredPublishingTasks(context.Context, time.Time, int32) (int64, error) {
	return 0, nil
}

func TestPushOutboxWorkerPublishesPendingTask(t *testing.T) {
	store := &fakePushTaskStore{tasks: []PushTask{{TaskID: 1, UserID: 10, Pts: 1, OperationID: "op", PushPartitionID: 3}}}
	publisher := &fakePushPublisher{}
	worker := NewPushOutboxWorker(store, publisher, PushOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, PublishingTimeout: time.Minute})

	worker.drain(context.Background())

	if calls := publisher.calls.Load(); calls != 1 {
		t.Fatalf("publisher calls = %d, want 1", calls)
	}
	if len(store.published) != 1 || store.published[0] != 1 {
		t.Fatalf("published = %+v", store.published)
	}
	if len(store.failed) != 0 {
		t.Fatalf("failed = %+v", store.failed)
	}
}

func TestPushOutboxWorkerWakeDrainsBeforeInterval(t *testing.T) {
	store := &fakePushTaskStore{tasks: []PushTask{{TaskID: 1, UserID: 10, Pts: 1, OperationID: "op", PushPartitionID: 3}}}
	publisher := &fakePushPublisher{}
	worker := NewPushOutboxWorker(store, publisher, PushOutboxWorkerOptions{Interval: time.Hour, BatchSize: 10, PublishingTimeout: time.Minute})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go worker.Run(ctx)

	worker.Wake()

	deadline := time.After(time.Second)
	for {
		if publisher.calls.Load() > 0 {
			return
		}
		select {
		case <-deadline:
			t.Fatal("worker did not publish after wake")
		case <-time.After(10 * time.Millisecond):
		}
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

func TestPushOutboxWorkerCountersForPublishSuccessAndFailure(t *testing.T) {
	successCounters := &fakePushOutboxWorkerCounters{}
	successStore := &fakePushTaskStore{tasks: []PushTask{{TaskID: 1, UserID: 10, Pts: 1, OperationID: "op", PushPartitionID: 3}}}
	successWorker := NewPushOutboxWorker(successStore, &fakePushPublisher{}, PushOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, PublishingTimeout: time.Minute}).WithCounters(successCounters)
	successWorker.drain(context.Background())
	if successCounters.drainStart != 1 || successCounters.claimSuccess != 1 || successCounters.publishSuccess != 1 {
		t.Fatalf("success counters drain=%d claim=%d publish=%d, want 1/1/1", successCounters.drainStart, successCounters.claimSuccess, successCounters.publishSuccess)
	}

	errorCounters := &fakePushOutboxWorkerCounters{}
	errorStore := &fakePushTaskStore{tasks: []PushTask{{TaskID: 2, UserID: 10, Pts: 2, OperationID: "op2", PushPartitionID: 4}}}
	errorWorker := NewPushOutboxWorker(errorStore, &fakePushPublisher{err: errors.New("kafka down")}, PushOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, PublishingTimeout: time.Minute}).WithCounters(errorCounters)
	errorWorker.drain(context.Background())
	if errorCounters.publishError != 1 {
		t.Fatalf("publish error counter = %d, want 1", errorCounters.publishError)
	}
}

func TestPushOutboxWorkerRecoverDrainPanic(t *testing.T) {
	counters := &fakePushOutboxWorkerCounters{}
	worker := NewPushOutboxWorker(panicPushTaskStore{}, &fakePushPublisher{}, PushOutboxWorkerOptions{Interval: time.Second, BatchSize: 10, PublishingTimeout: time.Minute}).WithCounters(counters)

	worker.drain(context.Background())

	if counters.panicRecovered != 1 {
		t.Fatalf("panic counter = %d, want 1", counters.panicRecovered)
	}
}

type blockingPushTaskStore struct {
	started chan struct{}
	release chan struct{}
	done    atomic.Bool
}

func (s *blockingPushTaskStore) ListPendingPushTasks(ctx context.Context, now time.Time, limit int32) ([]PushTask, error) {
	close(s.started)
	<-s.release
	s.done.Store(true)
	return nil, nil
}
func (s *blockingPushTaskStore) TryMarkPushTaskPublishing(context.Context, int64, time.Time, time.Time) (bool, error) {
	return false, nil
}
func (s *blockingPushTaskStore) MarkPushTaskPublished(context.Context, int64, KafkaAck) error {
	return nil
}
func (s *blockingPushTaskStore) MarkPushTaskPublishFailed(context.Context, int64, string, time.Time) error {
	return nil
}
func (s *blockingPushTaskStore) ResetExpiredPublishingTasks(context.Context, time.Time, int32) (int64, error) {
	return 0, nil
}

func TestPushOutboxWorkerWaitsForDrain(t *testing.T) {
	store := &blockingPushTaskStore{
		started: make(chan struct{}),
		release: make(chan struct{}),
	}
	worker := NewPushOutboxWorker(store, &fakePushPublisher{}, PushOutboxWorkerOptions{Interval: time.Hour, BatchSize: 10, PublishingTimeout: time.Minute})

	go worker.runDrain(context.Background())
	<-store.started

	done := make(chan struct{})
	go func() {
		worker.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Fatal("Wait returned while drain was still running")
	case <-time.After(20 * time.Millisecond):
	}

	close(store.release)

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Wait did not return after drain completed")
	}
	if !store.done.Load() {
		t.Fatal("drain did not complete")
	}
}

func TestPushOutboxWorkerStopPreventsNewDrain(t *testing.T) {
	store := &fakePushTaskStore{tasks: []PushTask{{TaskID: 1, UserID: 10, Pts: 1, OperationID: "op", PushPartitionID: 3}}}
	publisher := &fakePushPublisher{}
	worker := NewPushOutboxWorker(store, publisher, PushOutboxWorkerOptions{Interval: time.Hour, BatchSize: 10, PublishingTimeout: time.Minute})

	worker.Stop()
	worker.runDrain(context.Background())

	if calls := publisher.calls.Load(); calls != 0 {
		t.Fatalf("publisher calls = %d, want 0 after Stop", calls)
	}
}
