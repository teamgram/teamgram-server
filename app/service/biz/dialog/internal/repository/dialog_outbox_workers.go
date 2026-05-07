package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DialogOutboxUserupdatesClient interface {
	UserupdatesAppendDialogAuthSeqSideEffect(ctx context.Context, in *userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect) (*userupdates.UserAuthSeqAppendResult, error)
	UserupdatesAppendDialogPtsSideEffect(ctx context.Context, in *userupdates.TLUserupdatesAppendDialogPtsSideEffect) (*userupdates.UserPtsAppendResult, error)
}

type DialogOutboxWorkerOptions struct {
	Owner          string
	BatchSize      int32
	LeaseSeconds   int32
	PollInterval   time.Duration
	BlockedAttempt int32
}

type DialogAuthSeqOutboxWorker struct {
	repo    *Repository
	client  DialogOutboxUserupdatesClient
	options DialogOutboxWorkerOptions
	stop    chan struct{}
	done    chan struct{}
	once    sync.Once
}

type DialogPublicUpdateOutboxWorker struct {
	repo    *Repository
	client  DialogOutboxUserupdatesClient
	options DialogOutboxWorkerOptions
	stop    chan struct{}
	done    chan struct{}
	once    sync.Once
}

func NewDialogAuthSeqOutboxWorker(repo *Repository, client DialogOutboxUserupdatesClient, options DialogOutboxWorkerOptions) *DialogAuthSeqOutboxWorker {
	return &DialogAuthSeqOutboxWorker{
		repo:    repo,
		client:  client,
		options: normalizeOutboxWorkerOptions(options, "dialog-auth-seq-outbox"),
		stop:    make(chan struct{}),
		done:    make(chan struct{}),
	}
}

func NewDialogPublicUpdateOutboxWorker(repo *Repository, client DialogOutboxUserupdatesClient, options DialogOutboxWorkerOptions) *DialogPublicUpdateOutboxWorker {
	return &DialogPublicUpdateOutboxWorker{
		repo:    repo,
		client:  client,
		options: normalizeOutboxWorkerOptions(options, "dialog-public-update-outbox"),
		stop:    make(chan struct{}),
		done:    make(chan struct{}),
	}
}

func (w *DialogAuthSeqOutboxWorker) Run(ctx context.Context) {
	defer close(w.done)
	ticker := time.NewTicker(w.options.PollInterval)
	defer ticker.Stop()

	for {
		if err := w.Drain(ctx); err != nil {
			logx.Errorf("dialog auth seq outbox drain failed: %v", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-w.stop:
			return
		case <-ticker.C:
		}
	}
}

func (w *DialogAuthSeqOutboxWorker) Stop() {
	w.once.Do(func() { close(w.stop) })
}

func (w *DialogAuthSeqOutboxWorker) Wait() {
	<-w.done
}

func (w *DialogAuthSeqOutboxWorker) Drain(ctx context.Context) error {
	if w == nil || w.repo == nil || w.client == nil {
		return nil
	}
	now := time.Now().UTC()
	rows, err := w.repo.ClaimDialogAuthSeqOutbox(ctx, w.options.Owner, now, now.Add(time.Duration(w.options.LeaseSeconds)*time.Second), w.options.BatchSize)
	if err != nil {
		return err
	}
	for i := range rows {
		if err := w.processRow(ctx, rows[i]); err != nil {
			return err
		}
	}
	return nil
}

func (w *DialogAuthSeqOutboxWorker) processRow(ctx context.Context, row model.DialogAuthSeqOutbox) error {
	if row.TargetAuthPolicy == TargetAuthPolicyNotSourcePermAuthKey && row.SourcePermAuthKeyId == 0 {
		return w.repo.MarkDialogAuthSeqOutboxBlocked(ctx, row.OutboxId, "invalid_target_auth_policy", "source perm auth key is required")
	}
	result, err := w.client.UserupdatesAppendDialogAuthSeqSideEffect(ctx, &userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect{
		UserId:               row.UserId,
		SourcePermAuthKeyId:  row.SourcePermAuthKeyId,
		OperationId:          row.OperationId,
		TargetAuthPolicy:     row.TargetAuthPolicy,
		PublicUpdateType:     row.EventType,
		PeerType:             row.PeerType,
		PeerId:               row.PeerId,
		PayloadSchemaVersion: row.PayloadSchemaVersion,
		Payload:              row.Payload,
		PayloadHash:          row.PayloadHash,
	})
	if err != nil {
		logx.Errorf("dialog auth seq outbox publish failed: outbox_id=%d user_id=%d err=%v", row.OutboxId, row.UserId, err)
		return w.markAuthSeqFailure(ctx, row, "userupdates_append_auth_seq", err)
	}
	return w.repo.MarkDialogAuthSeqOutboxPublished(ctx, row.OutboxId, result.Seq, result.Date)
}

func (w *DialogAuthSeqOutboxWorker) markAuthSeqFailure(ctx context.Context, row model.DialogAuthSeqOutbox, kind string, err error) error {
	next := nextOutboxRetry(row.AttemptCount + 1)
	now := unixNow()
	if row.AttemptCount+1 >= w.options.BlockedAttempt || outboxRetryAgeExceeded(row.NextRetryAt, now) {
		return w.repo.MarkDialogAuthSeqOutboxBlocked(ctx, row.OutboxId, kind, err.Error())
	}
	return w.repo.MarkDialogAuthSeqOutboxRetryable(ctx, row.OutboxId, OutboxRetryState{
		AttemptCount:     row.AttemptCount + 1,
		NextRetryAt:      now + int64(next/time.Second),
		LastErrorKind:    kind,
		LastErrorMessage: err.Error(),
	})
}

func (w *DialogPublicUpdateOutboxWorker) Run(ctx context.Context) {
	defer close(w.done)
	ticker := time.NewTicker(w.options.PollInterval)
	defer ticker.Stop()

	for {
		if err := w.Drain(ctx); err != nil {
			logx.Errorf("dialog public update outbox drain failed: %v", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-w.stop:
			return
		case <-ticker.C:
		}
	}
}

func (w *DialogPublicUpdateOutboxWorker) Stop() {
	w.once.Do(func() { close(w.stop) })
}

func (w *DialogPublicUpdateOutboxWorker) Wait() {
	<-w.done
}

func (w *DialogPublicUpdateOutboxWorker) Drain(ctx context.Context) error {
	if w == nil || w.repo == nil || w.client == nil {
		return nil
	}
	now := time.Now().UTC()
	rows, err := w.repo.ClaimDialogPublicUpdateOutbox(ctx, w.options.Owner, now, now.Add(time.Duration(w.options.LeaseSeconds)*time.Second), w.options.BatchSize)
	if err != nil {
		return err
	}
	for i := range rows {
		if err := w.processRow(ctx, rows[i]); err != nil {
			return err
		}
	}
	return nil
}

func (w *DialogPublicUpdateOutboxWorker) processRow(ctx context.Context, row model.DialogPublicUpdateOutbox) error {
	if row.TargetAuthPolicy == TargetAuthPolicyNotSourcePermAuthKey && row.SourcePermAuthKeyId == 0 {
		return w.repo.MarkDialogPublicUpdateOutboxBlocked(ctx, row.OutboxId, "invalid_target_auth_policy", "source perm auth key is required")
	}

	switch row.DeliveryPath {
	case DeliveryPathUserupdatesPTS:
		result, err := w.client.UserupdatesAppendDialogPtsSideEffect(ctx, &userupdates.TLUserupdatesAppendDialogPtsSideEffect{
			UserId:               row.TargetUserId,
			SourcePermAuthKeyId:  row.SourcePermAuthKeyId,
			OperationId:          row.OperationId,
			TargetAuthPolicy:     row.TargetAuthPolicy,
			PublicUpdateType:     row.PublicUpdateType,
			PeerType:             row.PeerType,
			PeerId:               row.PeerId,
			PayloadSchemaVersion: row.PayloadSchemaVersion,
			Payload:              row.Payload,
			PayloadHash:          row.PayloadHash,
		})
		if err != nil {
			logx.Errorf("dialog public update outbox pts publish failed: outbox_id=%d target_user_id=%d err=%v", row.OutboxId, row.TargetUserId, err)
			return w.markPublicUpdateFailure(ctx, row, "userupdates_append_pts", err)
		}
		return w.repo.MarkDialogPublicUpdateOutboxPublishedPTS(ctx, row.OutboxId, result.Pts, result.PtsCount)
	case DeliveryPathUserupdatesAuthSeq:
		result, err := w.client.UserupdatesAppendDialogAuthSeqSideEffect(ctx, &userupdates.TLUserupdatesAppendDialogAuthSeqSideEffect{
			UserId:               row.TargetUserId,
			SourcePermAuthKeyId:  row.SourcePermAuthKeyId,
			OperationId:          row.OperationId,
			TargetAuthPolicy:     row.TargetAuthPolicy,
			PublicUpdateType:     row.PublicUpdateType,
			PeerType:             row.PeerType,
			PeerId:               row.PeerId,
			PayloadSchemaVersion: row.PayloadSchemaVersion,
			Payload:              row.Payload,
			PayloadHash:          row.PayloadHash,
		})
		if err != nil {
			logx.Errorf("dialog public update outbox auth seq publish failed: outbox_id=%d target_user_id=%d err=%v", row.OutboxId, row.TargetUserId, err)
			return w.markPublicUpdateFailure(ctx, row, "userupdates_append_auth_seq", err)
		}
		return w.repo.MarkDialogPublicUpdateOutboxPublishedAuthSeq(ctx, row.OutboxId, result.Seq, result.Date)
	default:
		return w.repo.MarkDialogPublicUpdateOutboxBlocked(ctx, row.OutboxId, "unsupported_delivery_path", fmt.Sprintf("unsupported delivery path %q", row.DeliveryPath))
	}
}

func (w *DialogPublicUpdateOutboxWorker) markPublicUpdateFailure(ctx context.Context, row model.DialogPublicUpdateOutbox, kind string, err error) error {
	next := nextOutboxRetry(row.AttemptCount + 1)
	now := unixNow()
	if row.AttemptCount+1 >= w.options.BlockedAttempt || outboxRetryAgeExceeded(row.NextRetryAt, now) {
		return w.repo.MarkDialogPublicUpdateOutboxBlocked(ctx, row.OutboxId, kind, err.Error())
	}
	return w.repo.MarkDialogPublicUpdateOutboxRetryable(ctx, row.OutboxId, OutboxRetryState{
		AttemptCount:     row.AttemptCount + 1,
		NextRetryAt:      now + int64(next/time.Second),
		LastErrorKind:    kind,
		LastErrorMessage: err.Error(),
	})
}

func normalizeOutboxWorkerOptions(options DialogOutboxWorkerOptions, defaultOwner string) DialogOutboxWorkerOptions {
	if options.Owner == "" {
		options.Owner = fmt.Sprintf("%s-%d", defaultOwner, time.Now().UnixNano())
	}
	if options.BatchSize <= 0 {
		options.BatchSize = DefaultOutboxWorkerBatchSize
	}
	if options.LeaseSeconds <= 0 {
		options.LeaseSeconds = DefaultOutboxWorkerLeaseSeconds
	}
	if options.PollInterval <= 0 {
		options.PollInterval = time.Duration(DefaultOutboxWorkerPollSeconds) * time.Second
	}
	if options.BlockedAttempt <= 0 {
		options.BlockedAttempt = OutboxWorkerBlockedAttempts
	}
	return options
}

func nextOutboxRetry(attempt int32) time.Duration {
	if attempt < 1 {
		attempt = 1
	}
	seconds := int64(InitialRetryDelaySeconds)
	for i := int32(1); i < attempt && seconds < OutboxWorkerMaxRetryDelay; i++ {
		seconds *= 2
	}
	if seconds > OutboxWorkerMaxRetryDelay {
		seconds = OutboxWorkerMaxRetryDelay
	}
	return time.Duration(seconds) * time.Second
}

func outboxRetryAgeExceeded(firstRetryAt int64, now int64) bool {
	if firstRetryAt <= 0 {
		return false
	}
	return now-firstRetryAt >= OutboxWorkerBlockedAgeSeconds
}
