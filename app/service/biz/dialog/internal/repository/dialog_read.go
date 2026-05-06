package repository

import (
	"context"
	"errors"

	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

func (r *Repository) ListDialogs(ctx context.Context, userID int64, excludePinned bool, folderID int32) ([]DialogRecord, error) {
	return nil, dialogpb.ErrDeprecatedMethod
}

func (r *Repository) CountDialogs(ctx context.Context, userID int64, excludePinned bool, folderID int32) (int32, error) {
	return 0, dialogpb.ErrDeprecatedMethod
}

func (r *Repository) ListPinnedDialogs(ctx context.Context, userID int64, folderID int32) ([]DialogRecord, error) {
	if r == nil || r.model == nil {
		return nil, wrapReadStorage("list pinned dialogs", errors.New("dialog models are not configured"))
	}
	var (
		rows     []model.DialogPreferences
		pinOrder func(model.DialogPreferences) int64
		err      error
	)
	if folderID != 0 {
		rows, err = r.model.DialogPreferencesModel.SelectFolderPinned(ctx, userID, folderID)
		pinOrder = func(row model.DialogPreferences) int64 { return row.FolderPinnedOrder }
	} else {
		rows, err = r.model.DialogPreferencesModel.SelectMainPinned(ctx, userID)
		pinOrder = func(row model.DialogPreferences) int64 { return row.MainPinnedOrder }
	}
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return []DialogRecord{}, nil
		}
		return nil, wrapReadStorage("list pinned dialogs", err)
	}
	records := make([]DialogRecord, 0, len(rows))
	for _, row := range rows {
		records = append(records, DialogRecord{
			UserID:       row.UserId,
			PeerType:     row.PeerType,
			PeerID:       row.PeerId,
			PeerDialogID: row.PeerDialogId,
			FolderID:     row.FolderId,
			Pinned:       pinOrder(row),
			Order:        pinOrder(row),
		})
	}
	return records, nil
}

func (r *Repository) GetDialogByPeer(ctx context.Context, userID int64, peerType int32, peerID int64) (*DialogRecord, error) {
	return nil, dialogpb.ErrDeprecatedMethod
}

func (r *Repository) ListDialogsByPeerDialogIDs(ctx context.Context, userID int64, ids []int64) ([]DialogRecord, error) {
	return nil, dialogpb.ErrDeprecatedMethod
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
