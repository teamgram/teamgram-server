package repository

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) GetPhoto(ctx context.Context, id int64) (*tg.Photo, error) {
	if id == 0 {
		return nil, fmt.Errorf("media: photo id empty")
	}
	return r.loadPhoto(ctx, id)
}

func (r *Repository) mapPhotoResult(ctx context.Context, photo *tg.Photo, err error) (*tg.Photo, error) {
	if err != nil {
		return nil, fmt.Errorf("media get photo: %w", err)
	}
	return photo, nil
}

func (r *Repository) loadPhoto(ctx context.Context, id int64) (*tg.Photo, error) {
	do, err := r.model.PhotosModel.FindOneByPhotoId(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("media load photo %d: %w", id, err)
	}
	return photoFromModel(do), nil
}

func (r *Repository) GetPhotoByRequest(ctx context.Context, in *media.TLMediaGetPhoto) (*tg.Photo, error) {
	return r.GetPhoto(ctx, in.PhotoId)
}

func photoFromModel(do *model.Photos) *tg.Photo {
	if do == nil {
		return nil
	}
	return tg.MakeTLPhoto(&tg.TLPhoto{
		HasStickers:   do.HasStickers,
		Id:            do.PhotoId,
		AccessHash:    do.AccessHash,
		FileReference: []byte{},
		Date:          int32(do.Date2),
		Sizes:         []tg.PhotoSizeClazz{},
		DcId:          do.DcId,
	}).ToPhoto()
}
