package repository

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	affectedOutboxRetryableCode         = "affected_apply_retryable"
	affectedOutboxTerminalCode          = "affected_apply_terminal"
	affectedOutboxWorkerMaxErrorMessage = 240
)

type AffectedOutboxWorker struct {
	repo              affectedOutboxStore
	interval          time.Duration
	batchSize         int32
	processingTimeout time.Duration
	stop              chan struct{}
	wake              chan struct{}
	drainMu           sync.Mutex
	stopping          bool
	stopClosed        bool
	drainDone         chan struct{}
}

type AffectedOutboxWorkerOptions struct {
	Interval          time.Duration
	BatchSize         int32
	ProcessingTimeout time.Duration
}

type affectedOutboxStore interface {
	ListPendingAffectedOutboxes(ctx context.Context, now time.Time, limit int32) ([]AffectedOutboxRow, error)
	TryMarkAffectedOutboxProcessing(ctx context.Context, outboxID int64, now time.Time, deadline time.Time) (bool, error)
	ApplyUserOperation(ctx context.Context, in ApplyUserOperationInput) (*ApplyUserOperationResult, error)
	MarkAffectedOutboxCompleted(ctx context.Context, outboxID int64, processingDeadline time.Time) error
	MarkAffectedOutboxRetryable(ctx context.Context, outboxID int64, processingDeadline time.Time, code string, message string, nextAvailableAt time.Time) error
	MarkAffectedOutboxFailedTerminal(ctx context.Context, outboxID int64, processingDeadline time.Time, code string, message string) error
	ResetExpiredAffectedOutboxes(ctx context.Context, now time.Time, limit int32) (int64, error)
}

func NewAffectedOutboxWorker(repo affectedOutboxStore, options AffectedOutboxWorkerOptions) *AffectedOutboxWorker {
	if options.Interval <= 0 {
		options.Interval = time.Second
	}
	if options.BatchSize <= 0 {
		options.BatchSize = 100
	}
	if options.ProcessingTimeout <= 0 {
		options.ProcessingTimeout = time.Minute
	}
	return &AffectedOutboxWorker{
		repo:              repo,
		interval:          options.Interval,
		batchSize:         options.BatchSize,
		processingTimeout: options.ProcessingTimeout,
		stop:              make(chan struct{}),
		wake:              make(chan struct{}, 1),
	}
}

func (w *AffectedOutboxWorker) Run(ctx context.Context) {
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

func (w *AffectedOutboxWorker) Wake() {
	if w == nil {
		return
	}
	select {
	case w.wake <- struct{}{}:
	default:
	}
}

func (w *AffectedOutboxWorker) Stop() {
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

func (w *AffectedOutboxWorker) Wait() {
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

func (w *AffectedOutboxWorker) runDrain(ctx context.Context) {
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

func (w *AffectedOutboxWorker) drain(ctx context.Context) {
	defer func() {
		if recovered := recover(); recovered != nil {
			logx.WithContext(ctx).Errorf("affected outbox drain panic recovered: %v", recovered)
		}
	}()
	now := time.Now().UTC()
	resetCount, err := w.repo.ResetExpiredAffectedOutboxes(ctx, now, w.batchSize)
	if err != nil {
		logx.WithContext(ctx).Errorf("affected outbox reset expired processing failed: batch_size=%d err=%v", w.batchSize, err)
		return
	}
	if resetCount > 0 {
		logx.WithContext(ctx).Infof("affected outbox reset expired processing: count=%d now=%s batch_size=%d", resetCount, now.Format(time.RFC3339Nano), w.batchSize)
	}

	rows, err := w.repo.ListPendingAffectedOutboxes(ctx, now, w.batchSize)
	if err != nil {
		logx.WithContext(ctx).Errorf("affected outbox list pending failed: batch_size=%d err=%v", w.batchSize, err)
		return
	}
	for _, row := range rows {
		claimTime := time.Now().UTC()
		processingDeadline := claimTime.Add(w.processingTimeout)
		claimed, err := w.repo.TryMarkAffectedOutboxProcessing(ctx, row.OutboxID, claimTime, processingDeadline)
		if err != nil {
			logx.WithContext(ctx).Errorf("affected outbox claim failed: outbox_id=%d user_id=%d operation_id=%s err=%v", row.OutboxID, row.UserID, row.OperationID, err)
			continue
		}
		if !claimed {
			continue
		}
		w.applyRow(ctx, row, processingDeadline)
	}
}

func (w *AffectedOutboxWorker) applyRow(ctx context.Context, row AffectedOutboxRow, processingDeadline time.Time) {
	_, err := w.repo.ApplyUserOperation(ctx, ApplyUserOperationInput{
		UserID:       row.UserID,
		OperationID:  row.OperationID,
		OpType:       row.OpType,
		PeerType:     row.PeerType,
		PeerID:       row.PeerID,
		PayloadCodec: row.PayloadCodec,
		Payload:      row.Payload,
		PayloadHash:  row.PayloadHash,
		BucketID:     row.BucketID,
		PartitionID:  row.PartitionID,
	})
	if err == nil {
		if markErr := w.repo.MarkAffectedOutboxCompleted(ctx, row.OutboxID, processingDeadline); markErr != nil {
			logx.WithContext(ctx).Errorf("affected outbox mark completed failed: outbox_id=%d user_id=%d operation_id=%s err=%v", row.OutboxID, row.UserID, row.OperationID, markErr)
		}
		return
	}

	message := boundedAffectedOutboxErrorMessage(err)
	if errors.Is(err, userupdates.ErrOperationTerminal) || errors.Is(err, userupdates.ErrOperationPayloadConflict) {
		logx.WithContext(ctx).Errorf("affected outbox apply terminal failed: outbox_id=%d user_id=%d operation_id=%s err=%v", row.OutboxID, row.UserID, row.OperationID, err)
		if markErr := w.repo.MarkAffectedOutboxFailedTerminal(ctx, row.OutboxID, processingDeadline, affectedOutboxTerminalCode, message); markErr != nil {
			logx.WithContext(ctx).Errorf("affected outbox mark terminal failed: outbox_id=%d user_id=%d operation_id=%s err=%v", row.OutboxID, row.UserID, row.OperationID, markErr)
		}
		return
	}

	nextAvailableAt := time.Now().UTC().Add(w.interval)
	logx.WithContext(ctx).Errorf("affected outbox apply retryable failed: outbox_id=%d user_id=%d operation_id=%s code=%s err=%v", row.OutboxID, row.UserID, row.OperationID, affectedOutboxRetryableCode, err)
	if markErr := w.repo.MarkAffectedOutboxRetryable(ctx, row.OutboxID, processingDeadline, affectedOutboxRetryableCode, message, nextAvailableAt); markErr != nil {
		logx.WithContext(ctx).Errorf("affected outbox mark retryable failed: outbox_id=%d user_id=%d operation_id=%s err=%v", row.OutboxID, row.UserID, row.OperationID, markErr)
	}
}

func boundedAffectedOutboxErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	message := strings.TrimSpace(err.Error())
	if len(message) <= affectedOutboxWorkerMaxErrorMessage {
		return message
	}
	for len(message) > affectedOutboxWorkerMaxErrorMessage {
		_, size := utf8.DecodeLastRuneInString(message)
		if size <= 0 {
			return message[:affectedOutboxWorkerMaxErrorMessage]
		}
		message = message[:len(message)-size]
	}
	return message
}
