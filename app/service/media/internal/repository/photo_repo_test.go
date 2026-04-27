package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const testLayer = 223

func TestGetPhotoReturnsStorageError(t *testing.T) {
	r := &Repository{}
	errBoom := errors.New("db down")
	_, err := r.mapPhotoResult(context.Background(), nil, errBoom)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestGetPhotoReturnsPhoto(t *testing.T) {
	r := &Repository{}
	photo := tg.MakeTLPhoto(&tg.TLPhoto{Id: 10}).ToPhoto()
	got, err := r.mapPhotoResult(context.Background(), photo, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	gotPhoto, ok := got.ToPhoto()
	if !ok {
		t.Fatalf("expected photo, got %#v", got)
	}
	if gotPhoto.Id != 10 {
		t.Fatalf("expected photo id 10, got %d", gotPhoto.Id)
	}
}

func TestPhotoFromModelBuildsValidMinimalPhoto(t *testing.T) {
	got := photoFromModel(&model.Photos{
		PhotoId:     10,
		AccessHash:  20,
		HasStickers: true,
		DcId:        4,
		Date2:       30,
	})
	gotPhoto, ok := got.ToPhoto()
	if !ok {
		t.Fatalf("expected photo, got %#v", got)
	}
	if gotPhoto.Id != 10 {
		t.Fatalf("expected photo id 10, got %d", gotPhoto.Id)
	}
	if gotPhoto.FileReference == nil {
		t.Fatal("expected non-nil file_reference")
	}
	if gotPhoto.Sizes == nil {
		t.Fatal("expected non-nil required sizes")
	}
	if gotPhoto.VideoSizes != nil {
		t.Fatalf("expected absent video_sizes, got %#v", gotPhoto.VideoSizes)
	}
	if err := gotPhoto.Validate(testLayer); err != nil {
		t.Fatalf("expected valid photo: %v", err)
	}
}
