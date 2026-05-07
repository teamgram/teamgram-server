package repository

import (
	"context"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushOutboxWorker struct {
	repo              pushTaskOutboxStore
	publisher         pushTaskPublisher
	interval          time.Duration
	batchSize         int32
	publishingTimeout time.Duration
	stop              chan struct{}
	wake              chan struct{}
	drainMu           sync.Mutex
	stopping          bool
	stopClosed        bool
	drainDone         chan struct{}
	counters          PushOutboxWorkerCounters
}

type PushOutboxWorkerCounters interface {
	IncDrainStart()
	IncResetStaleSuccess()
	IncResetStaleError()
	IncClaimEmpty()
	IncClaimSuccess()
	IncPublishSuccess()
	IncPublishError()
	IncMarkPublishedError()
	IncMarkRetryableError()
	IncPanicRecovered()
}

type noopPushOutboxWorkerCounters struct{}

func (noopPushOutboxWorkerCounters) IncDrainStart()         {}
func (noopPushOutboxWorkerCounters) IncResetStaleSuccess()  {}
func (noopPushOutboxWorkerCounters) IncResetStaleError()    {}
func (noopPushOutboxWorkerCounters) IncClaimEmpty()         {}
func (noopPushOutboxWorkerCounters) IncClaimSuccess()       {}
func (noopPushOutboxWorkerCounters) IncPublishSuccess()     {}
func (noopPushOutboxWorkerCounters) IncPublishError()       {}
func (noopPushOutboxWorkerCounters) IncMarkPublishedError() {}
func (noopPushOutboxWorkerCounters) IncMarkRetryableError() {}
func (noopPushOutboxWorkerCounters) IncPanicRecovered()     {}

type PushOutboxWorkerOptions struct {
	Interval          time.Duration
	BatchSize         int32
	PublishingTimeout time.Duration
}

func NewPushOutboxWorker(repo pushTaskOutboxStore, publisher pushTaskPublisher, options PushOutboxWorkerOptions) *PushOutboxWorker {
	if options.Interval <= 0 {
		options.Interval = time.Second
	}
	if options.BatchSize <= 0 {
		options.BatchSize = 100
	}
	if options.PublishingTimeout <= 0 {
		options.PublishingTimeout = time.Minute
	}
	return &PushOutboxWorker{
		repo:              repo,
		publisher:         publisher,
		interval:          options.Interval,
		batchSize:         options.BatchSize,
		publishingTimeout: options.PublishingTimeout,
		stop:              make(chan struct{}),
		wake:              make(chan struct{}, 1),
		counters:          noopPushOutboxWorkerCounters{},
	}
}

func (w *PushOutboxWorker) WithCounters(counters PushOutboxWorkerCounters) *PushOutboxWorker {
	if counters == nil {
		w.counters = noopPushOutboxWorkerCounters{}
		return w
	}
	w.counters = counters
	return w
}

func (w *PushOutboxWorker) ensureCounters() {
	if w.counters == nil {
		w.counters = noopPushOutboxWorkerCounters{}
	}
}

func (w *PushOutboxWorker) Wait() {
	if w == nil {
		return
	}
	w.drainMu.Lock()
	done := w.drainDone
	w.drainMu.Unlock()
	if done != nil {
		<-done
	}
}

type pushTaskOutboxStore interface {
	ListPendingPushTasks(ctx context.Context, now time.Time, limit int32) ([]PushTask, error)
	TryMarkPushTaskPublishing(ctx context.Context, taskID int64, now time.Time, leaseExpiresAt time.Time) (bool, error)
	MarkPushTaskPublished(ctx context.Context, taskID int64, ack KafkaAck) error
	MarkPushTaskPublishFailed(ctx context.Context, taskID int64, code string, nextRetryAt time.Time) error
	ResetExpiredPublishingTasks(ctx context.Context, now time.Time, limit int32) (int64, error)
}

type pushTaskPublisher interface {
	Publish(ctx context.Context, task PushTask) (KafkaAck, error)
}

func (w *PushOutboxWorker) Run(ctx context.Context) {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-w.stop:
			return
		case <-w.wake:
			w.runDrain(ctx)
		case <-ticker.C:
			w.runDrain(ctx)
		}
	}
}

func (w *PushOutboxWorker) Wake() {
	if w == nil {
		return
	}
	select {
	case w.wake <- struct{}{}:
	default:
	}
}

func (w *PushOutboxWorker) Stop() {
	if w == nil {
		return
	}
	w.drainMu.Lock()
	w.stopping = true
	if !w.stopClosed {
		close(w.stop)
		w.stopClosed = true
	}
	w.drainMu.Unlock()
}

func (w *PushOutboxWorker) runDrain(ctx context.Context) {
	if w == nil {
		return
	}
	done := make(chan struct{})
	w.drainMu.Lock()
	if w.stopping {
		w.drainMu.Unlock()
		return
	}
	w.drainDone = done
	w.drainMu.Unlock()
	defer func() {
		w.drainMu.Lock()
		if w.drainDone == done {
			w.drainDone = nil
		}
		w.drainMu.Unlock()
		close(done)
	}()
	w.drain(ctx)
}

func (w *PushOutboxWorker) drain(ctx context.Context) {
	w.ensureCounters()
	defer func() {
		if recovered := recover(); recovered != nil {
			w.counters.IncPanicRecovered()
			logx.WithContext(ctx).Errorf("push outbox drain panic recovered: %v", recovered)
		}
	}()
	w.counters.IncDrainStart()
	now := time.Now().UTC()
	if w.publishingTimeout > 0 {
		resetCount, err := w.repo.ResetExpiredPublishingTasks(ctx, now, w.batchSize)
		if err != nil {
			w.counters.IncResetStaleError()
			logx.WithContext(ctx).Errorf("push outbox reset expired publishing failed: batch_size=%d err=%v", w.batchSize, err)
			return
		}
		if resetCount > 0 {
			w.counters.IncResetStaleSuccess()
			logx.WithContext(ctx).Infof("push outbox reset expired publishing: count=%d now=%s batch_size=%d", resetCount, now.Format(time.RFC3339Nano), w.batchSize)
		}
	}

	tasks, err := w.repo.ListPendingPushTasks(ctx, now, w.batchSize)
	if err != nil {
		logx.WithContext(ctx).Errorf("push outbox list pending failed: batch_size=%d err=%v", w.batchSize, err)
		return
	}
	if len(tasks) == 0 {
		w.counters.IncClaimEmpty()
	}
	for _, task := range tasks {
		claimTime := time.Now().UTC()
		leaseExpiresAt := claimTime.Add(w.publishingTimeout)
		claimed, err := w.repo.TryMarkPushTaskPublishing(ctx, task.TaskID, claimTime, leaseExpiresAt)
		if err != nil {
			logx.WithContext(ctx).Errorf("push outbox claim failed: task_id=%d user_id=%d pts=%d operation_id=%s err=%v", task.TaskID, task.UserID, task.Pts, task.OperationID, err)
			continue
		}
		if !claimed {
			continue
		}
		w.counters.IncClaimSuccess()

		ack, err := w.publisher.Publish(ctx, task)
		if err != nil {
			w.counters.IncPublishError()
			nextRetryAt := time.Now().UTC().Add(w.interval)
			logx.WithContext(ctx).Errorf("push outbox publish failed: task_id=%d user_id=%d pts=%d operation_id=%s code=push_publish_failed err=%v", task.TaskID, task.UserID, task.Pts, task.OperationID, err)
			if markErr := w.repo.MarkPushTaskPublishFailed(ctx, task.TaskID, "push_publish_failed", nextRetryAt); markErr != nil {
				w.counters.IncMarkRetryableError()
				logx.WithContext(ctx).Errorf("push outbox mark publish failed failed: task_id=%d user_id=%d pts=%d operation_id=%s err=%v", task.TaskID, task.UserID, task.Pts, task.OperationID, markErr)
			}
			continue
		}
		w.counters.IncPublishSuccess()
		if err := w.repo.MarkPushTaskPublished(ctx, task.TaskID, ack); err != nil {
			w.counters.IncMarkPublishedError()
			logx.WithContext(ctx).Errorf("push outbox mark published failed: task_id=%d user_id=%d pts=%d operation_id=%s topic=%s partition=%d offset=%d err=%v", task.TaskID, task.UserID, task.Pts, task.OperationID, ack.Topic, ack.Partition, ack.Offset, err)
		}
	}
}
