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
	if r == nil || r.db == nil {
		return nil, wrapReadStorage("list pinned dialogs", errors.New("dialog mysql is not configured"))
	}
	orderColumn := "main_pinned_order"
	where := "user_id = ? AND main_pinned_order > 0"
	args := []interface{}{userID}
	if folderID != 0 {
		orderColumn = "folder_pinned_order"
		where = "user_id = ? AND folder_id = ? AND folder_pinned_order > 0"
		args = []interface{}{userID, folderID}
	}
	query := `
SELECT
	user_id, peer_type, peer_id, peer_dialog_id, folder_id, ` + orderColumn + ` AS pin_order
FROM dialog_preferences
WHERE ` + where + `
ORDER BY ` + orderColumn + ` DESC`
	type pinnedDialogRow struct {
		UserID       int64 `db:"user_id"`
		PeerType     int32 `db:"peer_type"`
		PeerID       int64 `db:"peer_id"`
		PeerDialogID int64 `db:"peer_dialog_id"`
		FolderID     int32 `db:"folder_id"`
		PinOrder     int64 `db:"pin_order"`
	}
	var rows []pinnedDialogRow
	if err := r.db.QueryRowsPartial(ctx, &rows, query, args...); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return []DialogRecord{}, nil
		}
		return nil, wrapReadStorage("list pinned dialogs", err)
	}
	records := make([]DialogRecord, 0, len(rows))
	for _, row := range rows {
		records = append(records, DialogRecord{
			UserID:       row.UserID,
			PeerType:     row.PeerType,
			PeerID:       row.PeerID,
			PeerDialogID: row.PeerDialogID,
			FolderID:     row.FolderID,
			Pinned:       row.PinOrder,
			Order:        row.PinOrder,
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
