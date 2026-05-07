package repository

import (
	"context"
	"time"

	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
	"github.com/zeromicro/go-zero/core/logx"
)

type DialogSideEffectWorker struct {
	repo      dialogSideEffectStore
	dialog    dialogSideEffectDialogClient
	interval  time.Duration
	batchSize int32
	stop      chan struct{}
}

type DialogSideEffectWorkerOptions struct {
	Interval  time.Duration
	BatchSize int32
}

type dialogSideEffectStore interface {
	ClaimDialogSideEffectsByKind(ctx context.Context, kind string, now time.Time, limit int32) ([]DialogSideEffect, error)
	MarkDialogSideEffectCompleted(ctx context.Context, sideEffectID int64) error
	MarkDialogSideEffectRetryableFailure(ctx context.Context, sideEffectID int64, errCode string, now time.Time) error
	MarkDialogSideEffectBlocked(ctx context.Context, sideEffectID int64, errCode string) error
}

type dialogSideEffectDialogClient interface {
	DialogUpsertSavedDialogFromMessage(ctx context.Context, in *dialogpb.TLDialogUpsertSavedDialogFromMessage) (*tg.Bool, error)
}

func NewDialogSideEffectWorker(repo dialogSideEffectStore, dialog dialogSideEffectDialogClient, options DialogSideEffectWorkerOptions) *DialogSideEffectWorker {
	if options.Interval <= 0 {
		options.Interval = time.Second
	}
	if options.BatchSize <= 0 {
		options.BatchSize = OutboxWorkerBatchSize
	}
	return &DialogSideEffectWorker{
		repo:      repo,
		dialog:    dialog,
		interval:  options.Interval,
		batchSize: options.BatchSize,
		stop:      make(chan struct{}),
	}
}

func (w *DialogSideEffectWorker) Run(ctx context.Context) {
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

func (w *DialogSideEffectWorker) Stop() {
	if w == nil {
		return
	}
	select {
	case <-w.stop:
	default:
		close(w.stop)
	}
}

func (w *DialogSideEffectWorker) drain(ctx context.Context) {
	if w == nil || w.repo == nil || w.dialog == nil {
		return
	}
	now := time.Now().UTC()
	rows, err := w.repo.ClaimDialogSideEffectsByKind(ctx, DialogSideEffectKindUpsertSavedDialogFromMessage, now, w.batchSize)
	if err != nil {
		logx.WithContext(ctx).Errorf("dialog side effect claim failed: kind=%s batch_size=%d err=%v", DialogSideEffectKindUpsertSavedDialogFromMessage, w.batchSize, err)
		return
	}
	for _, row := range rows {
		if err := w.publishSavedDialog(ctx, row); err != nil {
			logx.WithContext(ctx).Errorf("dialog side effect publish failed: side_effect_id=%d kind=%s user_id=%d peer_type=%d peer_id=%d err=%v",
				row.SideEffectID, row.Kind, row.UserID, row.PeerType, row.PeerID, err)
			if markErr := w.repo.MarkDialogSideEffectRetryableFailure(ctx, row.SideEffectID, "dialog_saved_side_effect_failed", time.Now().UTC()); markErr != nil {
				logx.WithContext(ctx).Errorf("dialog side effect mark retryable failed: side_effect_id=%d err=%v", row.SideEffectID, markErr)
			}
			continue
		}
		if err := w.repo.MarkDialogSideEffectCompleted(ctx, row.SideEffectID); err != nil {
			logx.WithContext(ctx).Errorf("dialog side effect mark completed failed: side_effect_id=%d err=%v", row.SideEffectID, err)
		}
	}
}

func (w *DialogSideEffectWorker) publishSavedDialog(ctx context.Context, row DialogSideEffect) error {
	if row.Kind != DialogSideEffectKindUpsertSavedDialogFromMessage {
		return w.repo.MarkDialogSideEffectBlocked(ctx, row.SideEffectID, "unsupported_dialog_side_effect_kind")
	}
	_, err := w.dialog.DialogUpsertSavedDialogFromMessage(ctx, &dialogpb.TLDialogUpsertSavedDialogFromMessage{
		UserId:                row.UserID,
		PeerType:              row.PeerType,
		PeerId:                row.PeerID,
		TopPeerSeq:            row.SourcePeerSeq,
		TopCanonicalMessageId: row.SourceCanonicalMessageID,
		TopMessageDate:        int32(row.SourceMessageDate),
		Payload:               row.Payload,
	})
	return err
}
