package repository

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (r *Repository) GetPhoto(ctx context.Context, id int64) (*tg.Photo, error) {
	if id == 0 {
		return nil, media.ErrPhotoNotFound
	}
	return r.loadPhoto(ctx, id)
}

func (r *Repository) mapPhotoResult(ctx context.Context, photo *tg.Photo, err error) (*tg.Photo, error) {
	if err != nil {
		if isServiceError(err) {
			return nil, err
		}
		return nil, wrapStorage("get photo", err)
	}
	return photo, nil
}

func (r *Repository) loadPhoto(ctx context.Context, id int64) (*tg.Photo, error) {
	do, err := r.model.PhotosModel.FindOneByPhotoId(ctx, id)
	if err != nil {
		if isNotFound(err) {
			return nil, media.ErrPhotoNotFound
		}
		return nil, wrapStorage("load photo", err)
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
