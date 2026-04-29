package repository

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMapPhotoIncludesPhotoAndVideoSizes(t *testing.T) {
	got, err := mapPhotoAggregate(
		&model.Photos{PhotoId: 10, AccessHash: 20, DcId: 1, Date2: 30, HasStickers: true},
		[]model.PhotoSizes{{SizeType: "m", Width: 320, Height: 200, FileSize: 123}},
		[]model.VideoSizes{{SizeType: "v", Width: 320, Height: 200, FileSize: 456, VideoStartTs: 1.5}},
	)
	if err != nil {
		t.Fatalf("mapPhotoAggregate() error = %v", err)
	}
	photo, ok := got.ToPhoto()
	if !ok {
		t.Fatalf("Photo clazz = %s, want photo", got.ClazzName())
	}
	if photo.Id != 10 || !photo.HasStickers {
		t.Fatalf("photo = %+v, want id 10 has stickers", photo)
	}
	if len(photo.Sizes) != 1 {
		t.Fatalf("len(Sizes) = %d, want 1", len(photo.Sizes))
	}
	if _, ok := photo.Sizes[0].(*tg.TLPhotoSize); !ok {
		t.Fatalf("photo size type = %T, want TLPhotoSize", photo.Sizes[0])
	}
	if len(photo.VideoSizes) != 1 {
		t.Fatalf("len(VideoSizes) = %d, want 1", len(photo.VideoSizes))
	}
}

func TestMapPhotoReturnsNotFoundForNilPhotoRow(t *testing.T) {
	_, err := mapPhotoAggregate(nil, nil, nil)
	if !errors.Is(err, media.ErrPhotoNotFound) {
		t.Fatalf("mapPhotoAggregate() error = %v, want ErrPhotoNotFound", err)
	}
}
