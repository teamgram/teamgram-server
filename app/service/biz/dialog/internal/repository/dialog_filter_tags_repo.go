package repository

import (
	"context"
	"errors"

	dialogpb "github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/internal/repository/model"
)

func (r *Repository) GetDialogFilterTagsEnabled(ctx context.Context, userID int64) (bool, error) {
	models, err := r.requireModels()
	if err != nil {
		return false, err
	}
	row, err := models.DialogFilterTagsModel.Select(ctx, userID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return false, nil
		}
		return false, storageError("get dialog filter tags", err)
	}
	return row.Enabled, nil
}

func (r *Repository) SetDialogFilterTagsEnabled(ctx context.Context, userID int64, enabled bool) error {
	models, err := r.requireModels()
	if err != nil {
		return err
	}
	if userID == 0 {
		return dialogpb.ErrDialogInvalid
	}
	if _, _, err := models.DialogFilterTagsModel.InsertOrUpdate(ctx, &model.DialogFilterTags{
		UserId:  userID,
		Enabled: enabled,
	}); err != nil {
		return storageError("set dialog filter tags", err)
	}
	return nil
}
