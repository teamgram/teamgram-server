package repository

import (
	"context"
	"errors"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

func (r *Repository) ListDialogFilters(ctx context.Context, userID int64) ([]DialogFilterRecord, error) {
	models, err := r.requireModels()
	if err != nil {
		return nil, err
	}
	rows, err := models.DialogFiltersModel.SelectList(ctx, userID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return []DialogFilterRecord{}, nil
		}
		return nil, storageError("list dialog filters", err)
	}
	return mapDialogFilterRecords(rows), nil
}

func (r *Repository) GetDialogFilter(ctx context.Context, userID int64, filterID int32) (*DialogFilterRecord, error) {
	models, err := r.requireModels()
	if err != nil {
		return nil, err
	}
	row, err := models.DialogFiltersModel.Select(ctx, userID, filterID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, dialogpb.ErrDialogNotFound
		}
		return nil, storageError("get dialog filter", err)
	}
	record := mapDialogFilterRecord(*row)
	return &record, nil
}

func (r *Repository) GetDialogFilterBySlug(ctx context.Context, userID int64, slug string) (*DialogFilterRecord, error) {
	models, err := r.requireModels()
	if err != nil {
		return nil, err
	}
	row, err := models.DialogFiltersModel.SelectBySlug(ctx, userID, slug)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, dialogpb.ErrDialogNotFound
		}
		return nil, storageError("get dialog filter by slug", err)
	}
	record := mapDialogFilterRecord(*row)
	return &record, nil
}

func (r *Repository) SaveDialogFilter(ctx context.Context, in SaveDialogFilterInput) (*DialogFilterRecord, error) {
	if in.SourcePermAuthKeyID == 0 {
		return nil, dialogpb.ErrSourceAuthKeyRequired
	}
	db, err := r.requireDB()
	if err != nil {
		return nil, err
	}
	if len(in.FilterPayload) == 0 {
		in.FilterPayload = []byte(`{"schema_version":1}`)
	}
	var saved DialogFilterRecord
	err = db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		outbox := authSeqOutboxInput{
			OutboxID:            in.OutboxID,
			UserID:              in.UserID,
			SourcePermAuthKeyID: in.SourcePermAuthKeyID,
			OperationID:         in.OperationID,
			EventType:           in.EventType,
			PeerType:            0,
			PeerID:              int64(in.DialogFilterID),
			Payload:             in.FilterPayload,
		}
		duplicate, err := authSeqOutboxDuplicateExists(txModels, outbox)
		if err != nil || duplicate {
			return err
		}
		row := &model.DialogFilters{
			UserId:              in.UserID,
			DialogFilterId:      in.DialogFilterID,
			Slug:                in.Slug,
			Title:               in.Title,
			OrderValue:          in.OrderValue,
			Enabled:             in.Enabled,
			Deleted:             false,
			FilterSchemaVersion: in.FilterSchemaVersion,
			FilterPayload:       in.FilterPayload,
		}
		if row.FilterSchemaVersion == 0 {
			row.FilterSchemaVersion = 1
		}
		if _, _, err := txModels.DialogFiltersModel.InsertOrUpdate(row); err != nil {
			return storageError("save dialog filter", err)
		}
		if _, err := insertAuthSeqOutbox(txModels, outbox); err != nil {
			return err
		}
		saved = mapDialogFilterRecord(*row)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &saved, nil
}

func (r *Repository) DeleteDialogFilter(ctx context.Context, in DeleteDialogFilterInput) error {
	if in.SourcePermAuthKeyID == 0 {
		return dialogpb.ErrSourceAuthKeyRequired
	}
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	return db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		outbox := authSeqOutboxInput{
			OutboxID:            in.OutboxID,
			UserID:              in.UserID,
			SourcePermAuthKeyID: in.SourcePermAuthKeyID,
			OperationID:         in.OperationID,
			EventType:           in.EventType,
			PeerType:            0,
			PeerID:              int64(in.DialogFilterID),
		}
		duplicate, err := authSeqOutboxDuplicateExists(txModels, outbox)
		if err != nil || duplicate {
			return err
		}
		if _, err := txModels.DialogFiltersModel.Clear(in.UserID, in.DialogFilterID); err != nil {
			return storageError("delete dialog filter", err)
		}
		_, err = insertAuthSeqOutbox(txModels, outbox)
		return err
	})
}

func (r *Repository) ReorderDialogFilters(ctx context.Context, in ReorderDialogFiltersInput) error {
	if in.SourcePermAuthKeyID == 0 {
		return dialogpb.ErrSourceAuthKeyRequired
	}
	db, err := r.requireDB()
	if err != nil {
		return err
	}
	return db.Transact(ctx, func(tx *sqlx.Tx) error {
		txModels := r.model.WithTx(tx)
		outbox := authSeqOutboxInput{
			OutboxID:            in.OutboxID,
			UserID:              in.UserID,
			SourcePermAuthKeyID: in.SourcePermAuthKeyID,
			OperationID:         in.OperationID,
			EventType:           in.EventType,
			PeerType:            0,
			PeerID:              0,
		}
		duplicate, err := authSeqOutboxDuplicateExists(txModels, outbox)
		if err != nil || duplicate {
			return err
		}
		for i, id := range in.Order {
			if _, err := txModels.DialogFiltersModel.UpdateOrder(int64(i+1), in.UserID, id); err != nil {
				return storageError("reorder dialog filters", err)
			}
		}
		_, err = insertAuthSeqOutbox(txModels, outbox)
		return err
	})
}

func mapDialogFilterRecords(rows []model.DialogFilters) []DialogFilterRecord {
	if len(rows) == 0 {
		return []DialogFilterRecord{}
	}
	out := make([]DialogFilterRecord, 0, len(rows))
	for _, row := range rows {
		out = append(out, mapDialogFilterRecord(row))
	}
	return out
}

func mapDialogFilterRecord(row model.DialogFilters) DialogFilterRecord {
	return DialogFilterRecord{
		UserID:              row.UserId,
		DialogFilterID:      row.DialogFilterId,
		Slug:                row.Slug,
		Title:               row.Title,
		OrderValue:          row.OrderValue,
		Enabled:             row.Enabled,
		FilterSchemaVersion: row.FilterSchemaVersion,
		FilterPayload:       row.FilterPayload,
	}
}
