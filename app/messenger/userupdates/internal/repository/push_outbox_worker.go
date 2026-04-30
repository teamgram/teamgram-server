package repository

import (
	"context"
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
}

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
		case <-ticker.C:
			w.drain(ctx)
		}
	}
}

func (w *PushOutboxWorker) Stop() {
	if w == nil {
		return
	}
	select {
	case <-w.stop:
	default:
		close(w.stop)
	}
}

func (w *PushOutboxWorker) drain(ctx context.Context) {
	now := time.Now().UTC()
	if w.publishingTimeout > 0 {
		resetCount, err := w.repo.ResetExpiredPublishingTasks(ctx, now, w.batchSize)
		if err != nil {
			logx.WithContext(ctx).Errorf("push outbox reset expired publishing failed: batch_size=%d err=%v", w.batchSize, err)
			return
		}
		if resetCount > 0 {
			logx.WithContext(ctx).Infof("push outbox reset expired publishing: count=%d now=%s batch_size=%d", resetCount, now.Format(time.RFC3339Nano), w.batchSize)
		}
	}

	tasks, err := w.repo.ListPendingPushTasks(ctx, now, w.batchSize)
	if err != nil {
		logx.WithContext(ctx).Errorf("push outbox list pending failed: batch_size=%d err=%v", w.batchSize, err)
		return
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

		ack, err := w.publisher.Publish(ctx, task)
		if err != nil {
			nextRetryAt := time.Now().UTC().Add(w.interval)
			logx.WithContext(ctx).Errorf("push outbox publish failed: task_id=%d user_id=%d pts=%d operation_id=%s code=push_publish_failed err=%v", task.TaskID, task.UserID, task.Pts, task.OperationID, err)
			if markErr := w.repo.MarkPushTaskPublishFailed(ctx, task.TaskID, "push_publish_failed", nextRetryAt); markErr != nil {
				logx.WithContext(ctx).Errorf("push outbox mark publish failed failed: task_id=%d user_id=%d pts=%d operation_id=%s err=%v", task.TaskID, task.UserID, task.Pts, task.OperationID, markErr)
			}
			continue
		}
		if err := w.repo.MarkPushTaskPublished(ctx, task.TaskID, ack); err != nil {
			logx.WithContext(ctx).Errorf("push outbox mark published failed: task_id=%d user_id=%d pts=%d operation_id=%s topic=%s partition=%d offset=%d err=%v", task.TaskID, task.UserID, task.Pts, task.OperationID, ack.Topic, ack.Partition, ack.Offset, err)
		}
	}
}
