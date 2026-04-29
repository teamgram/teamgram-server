package repository

import (
	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func mapPhotoAggregate(photo *model.Photos, sizes []model.PhotoSizes, videoSizes []model.VideoSizes) (*tg.Photo, error) {
	if photo == nil {
		return nil, media.ErrPhotoNotFound
	}
	return tg.MakeTLPhoto(&tg.TLPhoto{
		HasStickers:   photo.HasStickers,
		Id:            photo.PhotoId,
		AccessHash:    photo.AccessHash,
		FileReference: []byte{},
		Date:          int32(photo.Date2),
		Sizes:         mapPhotoSizes(sizes),
		VideoSizes:    mapVideoSizes(videoSizes),
		DcId:          photo.DcId,
	}).ToPhoto(), nil
}

func mapPhotoSizes(sizes []model.PhotoSizes) []tg.PhotoSizeClazz {
	if len(sizes) == 0 {
		return []tg.PhotoSizeClazz{}
	}
	out := make([]tg.PhotoSizeClazz, 0, len(sizes))
	for _, size := range sizes {
		if size.HasStripped || size.SizeType == "i" {
			out = append(out, tg.MakeTLPhotoStrippedSize(&tg.TLPhotoStrippedSize{
				Type:  size.SizeType,
				Bytes: []byte(size.StrippedBytes),
			}))
			continue
		}
		out = append(out, tg.MakeTLPhotoSize(&tg.TLPhotoSize{
			Type:  size.SizeType,
			W:     size.Width,
			H:     size.Height,
			Size2: size.FileSize,
		}))
	}
	return out
}

func mapVideoSizes(sizes []model.VideoSizes) []tg.VideoSizeClazz {
	if len(sizes) == 0 {
		return nil
	}
	out := make([]tg.VideoSizeClazz, 0, len(sizes))
	for _, size := range sizes {
		startTs := size.VideoStartTs
		out = append(out, tg.MakeTLVideoSize(&tg.TLVideoSize{
			Type:         size.SizeType,
			W:            size.Width,
			H:            size.Height,
			Size2:        size.FileSize,
			VideoStartTs: &startTs,
		}))
	}
	return out
}
