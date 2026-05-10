package core

import (
	"context"
	"errors"
	"testing"

	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestResolveMessageMediaInputMediaPhotoUsesMediaGetPhotoReference(t *testing.T) {
	mediaClient := &fakeResolveMediaClient{
		photoResp: tg.MakeTLPhoto(&tg.TLPhoto{
			Id:            42,
			AccessHash:    99,
			FileReference: []byte("file-reference"),
			Date:          123,
			Sizes:         []tg.PhotoSizeClazz{tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: "m", W: 320, H: 240, Size2: 1000})},
			DcId:          2,
		}).ToPhoto(),
	}

	got, err := resolveMessageMedia(context.Background(), mediaClient, 1001, tg.MakeTLInputMediaPhoto(&tg.TLInputMediaPhoto{
		Id: tg.MakeTLInputPhoto(&tg.TLInputPhoto{Id: 42, AccessHash: 99, FileReference: []byte("input-reference")}),
	}))
	if err != nil {
		t.Fatalf("resolveMessageMedia() error = %v", err)
	}
	if mediaClient.photoReq == nil || mediaClient.photoReq.PhotoId != 42 {
		t.Fatalf("MediaGetPhoto request = %#v", mediaClient.photoReq)
	}
	mediaPhoto, ok := got.(*tg.TLMessageMediaPhoto)
	if !ok {
		t.Fatalf("message media = %#v, want TLMessageMediaPhoto", got)
	}
	photo, ok := mediaPhoto.Photo.(*tg.TLPhoto)
	if !ok {
		t.Fatalf("photo = %#v, want TLPhoto", mediaPhoto.Photo)
	}
	if string(photo.FileReference) != "file-reference" {
		t.Fatalf("FileReference = %q, want signed media reference", photo.FileReference)
	}
}

type fakeResolveMediaClient struct {
	photoReq  *mediapb.TLMediaGetPhoto
	photoResp *tg.Photo
	photoErr  error
}

func (f *fakeResolveMediaClient) MediaUploadPhotoFile(context.Context, *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
	return nil, errors.New("unexpected MediaUploadPhotoFile")
}

func (f *fakeResolveMediaClient) MediaGetPhoto(_ context.Context, in *mediapb.TLMediaGetPhoto) (*tg.Photo, error) {
	f.photoReq = in
	return f.photoResp, f.photoErr
}

func (f *fakeResolveMediaClient) MediaGetPhotoSizeList(context.Context, *mediapb.TLMediaGetPhotoSizeList) (*mediapb.PhotoSizeList, error) {
	return nil, errors.New("unexpected MediaGetPhotoSizeList")
}

func (f *fakeResolveMediaClient) MediaUploadedDocumentMedia(context.Context, *mediapb.TLMediaUploadedDocumentMedia) (*tg.MessageMedia, error) {
	return nil, errors.New("unexpected MediaUploadedDocumentMedia")
}

func (f *fakeResolveMediaClient) MediaGetDocument(context.Context, *mediapb.TLMediaGetDocument) (*tg.Document, error) {
	return nil, errors.New("unexpected MediaGetDocument")
}
