package repository

import (
	"context"
	"errors"

	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

func (r *Repository) ListDialogs(ctx context.Context, userID int64, excludePinned bool, folderID int32) ([]DialogRecord, error) {
	if r == nil || r.model == nil || r.model.DialogsModel == nil {
		return nil, wrapReadStorage("list dialogs", errors.New("dialog repository models not initialized"))
	}

	var (
		rows []model.Dialogs
		err  error
	)
	switch {
	case excludePinned && folderID != 0:
		rows, err = r.model.DialogsModel.SelectExcludeFolderPinnedDialogs(ctx, userID)
	case excludePinned:
		rows, err = r.model.DialogsModel.SelectExcludePinnedDialogs(ctx, userID)
	default:
		rows, err = r.model.DialogsModel.SelectDialogs(ctx, userID, folderID)
	}
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return []DialogRecord{}, nil
		}
		return nil, wrapReadStorage("list dialogs", err)
	}
	return mapDialogRecords(rows), nil
}

func (r *Repository) CountDialogs(ctx context.Context, userID int64, excludePinned bool, folderID int32) (int32, error) {
	records, err := r.ListDialogs(ctx, userID, excludePinned, folderID)
	if err != nil {
		return 0, err
	}
	return countDialogRecords(records), nil
}

func (r *Repository) GetDialogByPeer(ctx context.Context, userID int64, peerType int32, peerID int64) (*DialogRecord, error) {
	if r == nil || r.model == nil || r.model.DialogsModel == nil {
		return nil, wrapReadStorage("get dialog by peer", errors.New("dialog repository models not initialized"))
	}
	row, err := r.model.DialogsModel.SelectDialog(ctx, userID, peerType, peerID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, dialogpb.ErrDialogNotFound
		}
		return nil, wrapReadStorage("get dialog by peer", err)
	}
	record := mapDialogRecord(*row)
	return &record, nil
}

func (r *Repository) ListDialogsByPeerDialogIDs(ctx context.Context, userID int64, ids []int64) ([]DialogRecord, error) {
	if len(ids) == 0 {
		return []DialogRecord{}, nil
	}
	if r == nil || r.model == nil || r.model.DialogsModel == nil {
		return nil, wrapReadStorage("list dialogs by ids", errors.New("dialog repository models not initialized"))
	}
	rows, err := r.model.DialogsModel.SelectPeerDialogList(ctx, userID, ids)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return []DialogRecord{}, nil
		}
		return nil, wrapReadStorage("list dialogs by ids", err)
	}
	return mapDialogRecords(rows), nil
}

func countDialogRecords(records []DialogRecord) int32 {
	return int32(len(records))
}

func findDialogInRecords(records []DialogRecord, match func(DialogRecord) bool) (DialogRecord, bool) {
	for _, record := range records {
		if match(record) {
			return record, true
		}
	}
	return DialogRecord{}, false
}

func wrapReadStorage(op string, err error) error {
	return dialogpb.WrapDialogStorage(op, err)
}
