package repository

import (
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func mapPhotoAggregate(photo *model.Photos, sizes []model.PhotoSizes, videoSizes []model.VideoSizes, fileReference []byte) (*tg.Photo, error) {
	if photo == nil {
		return nil, media.ErrPhotoNotFound
	}
	photoSizes, err := mapPhotoSizes(sizes)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLPhoto(&tg.TLPhoto{
		HasStickers:   photo.HasStickers,
		Id:            photo.PhotoId,
		AccessHash:    photo.AccessHash,
		FileReference: fileReference,
		Date:          int32(photo.Date2),
		Sizes:         photoSizes,
		VideoSizes:    mapVideoSizes(videoSizes),
		DcId:          photo.DcId,
	}).ToPhoto(), nil
}

func mapPhotoSizes(sizes []model.PhotoSizes) ([]tg.PhotoSizeClazz, error) {
	if len(sizes) == 0 {
		return []tg.PhotoSizeClazz{}, nil
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
		if size.CachedType == photoSizeCachedTypeProgressive {
			var progressiveSizes []int32
			if err := json.Unmarshal([]byte(size.CachedBytes), &progressiveSizes); err != nil {
				return nil, fmt.Errorf("%w: decode progressive photo size %d/%s: %w", media.ErrMediaStorage, size.PhotoSizeId, size.SizeType, err)
			}
			out = append(out, tg.MakeTLPhotoSizeProgressive(&tg.TLPhotoSizeProgressive{
				Type:  size.SizeType,
				W:     size.Width,
				H:     size.Height,
				Sizes: progressiveSizes,
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
	return out, nil
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
