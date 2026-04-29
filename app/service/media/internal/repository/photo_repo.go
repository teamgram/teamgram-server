package repository

import (
	"context"

	dfsapi "github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
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
	var sizes []model.PhotoSizes
	if do.SizeId != 0 {
		sizes, err = r.model.PhotoSizesModel.SelectListByPhotoSizeId(ctx, do.SizeId)
		if err != nil {
			return nil, wrapStorage("load photo sizes", err)
		}
	}
	var videoSizes []model.VideoSizes
	if do.VideoSizeId != 0 {
		videoSizes, err = r.model.VideoSizesModel.SelectListByVideoSizeId(ctx, do.VideoSizeId)
		if err != nil {
			return nil, wrapStorage("load video sizes", err)
		}
	}
	return mapPhotoAggregate(do, sizes, videoSizes)
}

func (r *Repository) GetPhotoByRequest(ctx context.Context, in *media.TLMediaGetPhoto) (*tg.Photo, error) {
	return r.GetPhoto(ctx, in.PhotoId)
}

func photoFromModel(do *model.Photos) *tg.Photo {
	photo, err := mapPhotoAggregate(do, nil, nil)
	if err != nil {
		return nil
	}
	return photo
}

func (r *Repository) UploadPhotoFile(ctx context.Context, in *media.TLMediaUploadPhotoFile) (*tg.Photo, error) {
	if in == nil || in.File == nil {
		return nil, media.ErrMediaInvalidArgument
	}
	if r.dfsClient == nil {
		return nil, wrapMediaDownstream("dfs upload photo", media.ErrMediaDownstream)
	}
	photo, err := r.dfsClient.UploadPhotoFileV2(ctx, &dfsapi.TLDfsUploadPhotoFileV2{
		Creator: in.OwnerId,
		File:    in.File,
	})
	if err != nil {
		return nil, wrapDfsUploadError("dfs upload photo", err)
	}
	if err := r.savePhotoAggregate(ctx, inputFileName(in.File), photo); err != nil {
		return nil, err
	}
	return photo, nil
}

func inputFileName(file tg.InputFileClazz) string {
	switch f := file.(type) {
	case *tg.TLInputFile:
		return f.Name
	case *tg.TLInputFileBig:
		return f.Name
	default:
		return ""
	}
}

func (r *Repository) UploadProfilePhotoFile(ctx context.Context, in *media.TLMediaUploadProfilePhotoFile) (*tg.Photo, error) {
	if in == nil || (in.File == nil && in.Video == nil) {
		return nil, media.ErrMediaInvalidArgument
	}
	if r.dfsClient == nil {
		return nil, wrapMediaDownstream("dfs upload profile photo", media.ErrMediaDownstream)
	}
	photo, err := r.dfsClient.UploadProfilePhotoFileV2(ctx, &dfsapi.TLDfsUploadProfilePhotoFileV2{
		Creator:          in.OwnerId,
		File:             in.File,
		Video:            in.Video,
		VideoStartTs:     in.VideoStartTs,
		VideoEmojiMarkup: in.VideoEmojiMarkup,
	})
	if err != nil {
		return nil, wrapDfsUploadError("dfs upload profile photo", err)
	}
	if err := r.savePhotoAggregate(ctx, "", photo); err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *Repository) UploadedProfilePhoto(ctx context.Context, in *media.TLMediaUploadedProfilePhoto) (*tg.Photo, error) {
	if in == nil || in.PhotoId == 0 {
		return nil, media.ErrMediaInvalidArgument
	}
	if r.dfsClient == nil {
		return nil, wrapMediaDownstream("dfs uploaded profile photo", media.ErrMediaDownstream)
	}
	photo, err := r.dfsClient.UploadedProfilePhoto(ctx, &dfsapi.TLDfsUploadedProfilePhoto{
		Creator: in.OwnerId,
		PhotoId: in.PhotoId,
	})
	if err != nil {
		return nil, wrapDfsUploadError("dfs uploaded profile photo", err)
	}
	if err := r.savePhotoAggregate(ctx, "", photo); err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *Repository) GetPhotoSizeList(ctx context.Context, sizeID int64) (*media.PhotoSizeList, error) {
	if sizeID == 0 {
		return nil, media.ErrMediaInvalidArgument
	}
	sizes, err := r.model.PhotoSizesModel.SelectListByPhotoSizeId(ctx, sizeID)
	if err != nil {
		return nil, wrapStorage("load photo size list", err)
	}
	return media.MakeTLPhotoSizeList(&media.TLPhotoSizeList{SizeId: sizeID, Sizes: mapPhotoSizes(sizes), DcId: 1}).ToPhotoSizeList(), nil
}

func (r *Repository) GetPhotoSizeListList(ctx context.Context, ids []int64) (*media.VectorPhotoSizeList, error) {
	out := &media.VectorPhotoSizeList{Datas: make([]media.PhotoSizeListClazz, 0, len(ids))}
	for _, id := range ids {
		list, err := r.GetPhotoSizeList(ctx, id)
		if err != nil {
			return nil, err
		}
		out.Datas = append(out.Datas, list)
	}
	return out, nil
}

func (r *Repository) GetVideoSizeList(ctx context.Context, sizeID int64) (*media.VideoSizeList, error) {
	if sizeID == 0 {
		return nil, media.ErrMediaInvalidArgument
	}
	sizes, err := r.model.VideoSizesModel.SelectListByVideoSizeId(ctx, sizeID)
	if err != nil {
		return nil, wrapStorage("load video size list", err)
	}
	return media.MakeTLVideoSizeList(&media.TLVideoSizeList{SizeId: sizeID, Sizes: mapVideoSizes(sizes), DcId: 1}).ToVideoSizeList(), nil
}

func (r *Repository) savePhotoAggregate(ctx context.Context, inputFileName string, photo *tg.Photo) error {
	if r == nil || r.model == nil || photo == nil {
		return nil
	}
	do, ok := photo.ToPhoto()
	if !ok {
		return media.ErrMediaInvalidArgument
	}
	photoRow := &model.Photos{
		PhotoId:       do.Id,
		AccessHash:    do.AccessHash,
		HasStickers:   do.HasStickers,
		DcId:          do.DcId,
		Date2:         int64(do.Date),
		InputFileName: inputFileName,
	}
	if len(do.Sizes) > 0 {
		photoRow.SizeId = do.Id
	}
	if len(do.VideoSizes) > 0 {
		photoRow.HasVideo = true
		photoRow.VideoSizeId = do.Id
	}
	if _, err := r.model.PhotosModel.Insert2(ctx, photoRow); err != nil {
		return wrapStorage("save photo", err)
	}
	for _, size := range do.Sizes {
		if err := r.savePhotoSize(ctx, do.Id, size); err != nil {
			return err
		}
	}
	for _, size := range do.VideoSizes {
		if err := r.saveVideoSize(ctx, do.Id, size); err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) savePhotoSize(ctx context.Context, id int64, size tg.PhotoSizeClazz) error {
	switch s := size.(type) {
	case *tg.TLPhotoSize:
		_, _, err := r.model.PhotoSizesModel.Insert(ctx, &model.PhotoSizes{PhotoSizeId: id, SizeType: s.Type, Width: s.W, Height: s.H, FileSize: s.Size2})
		if err != nil {
			return wrapStorage("save photo size", err)
		}
	case *tg.TLPhotoStrippedSize:
		_, _, err := r.model.PhotoSizesModel.Insert(ctx, &model.PhotoSizes{PhotoSizeId: id, SizeType: s.Type, HasStripped: true, StrippedBytes: string(s.Bytes)})
		if err != nil {
			return wrapStorage("save stripped photo size", err)
		}
	}
	return nil
}

func (r *Repository) saveVideoSize(ctx context.Context, id int64, size tg.VideoSizeClazz) error {
	if s, ok := size.(*tg.TLVideoSize); ok {
		var startTs float64
		if s.VideoStartTs != nil {
			startTs = *s.VideoStartTs
		}
		_, _, err := r.model.VideoSizesModel.Insert(ctx, &model.VideoSizes{VideoSizeId: id, SizeType: s.Type, Width: s.W, Height: s.H, FileSize: s.Size2, VideoStartTs: startTs})
		if err != nil {
			return wrapStorage("save video size", err)
		}
	}
	return nil
}
