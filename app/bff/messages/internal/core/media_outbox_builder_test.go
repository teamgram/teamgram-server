package core

import (
	"bytes"
	"context"
	"errors"
	"testing"

	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestResolveMessageMediaInputMediaPhotoUsesMediaGetPhotoReference(t *testing.T) {
	fileReference := []byte("1234567890123456789012345")
	mediaClient := &fakeResolveMediaClient{
		photoResp: tg.MakeTLPhoto(&tg.TLPhoto{
			Id:            42,
			AccessHash:    99,
			FileReference: fileReference,
			Date:          123,
			Sizes:         []tg.PhotoSizeClazz{tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: "m", W: 320, H: 240, Size2: 1000})},
			DcId:          2,
		}).ToPhoto(),
	}

	got, err := resolveMessageMedia(context.Background(), mediaClient, nil, 1001, tg.MakeTLInputMediaPhoto(&tg.TLInputMediaPhoto{
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
	if string(photo.FileReference) != string(fileReference) {
		t.Fatalf("FileReference = %q, want signed media reference", photo.FileReference)
	}
	if len(photo.FileReference) != 25 {
		t.Fatalf("len(photo.FileReference) = %d, want 25", len(photo.FileReference))
	}
	if !bytes.Equal(photo.FileReference, fileReference) {
		t.Fatalf("FileReference = %x, want media reference %x", photo.FileReference, fileReference)
	}
}

func TestResolveMessageMediaInputMediaUploadedPhotoReturns25ByteFileReference(t *testing.T) {
	fileReference := []byte("1234567890123456789012345")
	mediaClient := &fakeResolveMediaClient{
		uploadPhotoResp: tg.MakeTLPhoto(&tg.TLPhoto{
			Id:            43,
			AccessHash:    100,
			FileReference: fileReference,
			Date:          124,
			Sizes:         []tg.PhotoSizeClazz{tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: "m", W: 320, H: 240, Size2: 1000})},
			DcId:          2,
		}).ToPhoto(),
	}

	got, err := resolveMessageMedia(context.Background(), mediaClient, nil, 1001, tg.MakeTLInputMediaUploadedPhoto(&tg.TLInputMediaUploadedPhoto{}))
	if err != nil {
		t.Fatalf("resolveMessageMedia() error = %v", err)
	}
	if mediaClient.uploadPhotoReq == nil || mediaClient.uploadPhotoReq.OwnerId != 1001 {
		t.Fatalf("MediaUploadPhotoFile request = %#v", mediaClient.uploadPhotoReq)
	}
	mediaPhoto, ok := got.(*tg.TLMessageMediaPhoto)
	if !ok {
		t.Fatalf("message media = %#v, want TLMessageMediaPhoto", got)
	}
	photo, ok := mediaPhoto.Photo.(*tg.TLPhoto)
	if !ok {
		t.Fatalf("photo = %#v, want TLPhoto", mediaPhoto.Photo)
	}
	if len(photo.FileReference) != 25 {
		t.Fatalf("len(photo.FileReference) = %d, want 25", len(photo.FileReference))
	}
	if !bytes.Equal(photo.FileReference, fileReference) {
		t.Fatalf("FileReference = %x, want media reference %x", photo.FileReference, fileReference)
	}
}

func TestResolveMessageMediaUploadedDocumentPreservesReturnedVideoCover(t *testing.T) {
	videoTimestamp := int32(19)
	videoCover := tg.MakeTLPhoto(&tg.TLPhoto{
		Id:         303,
		AccessHash: 404,
		Date:       123,
		Sizes:      []tg.PhotoSizeClazz{tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: "m", W: 320, H: 240, Size2: 1000})},
		DcId:       2,
	})
	mediaClient := &fakeResolveMediaClient{
		uploadedDocumentResp: tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
			Document: tg.MakeTLDocument(&tg.TLDocument{
				Id:         42,
				AccessHash: 99,
				MimeType:   "video/mp4",
				Size2:      1000,
				DcId:       2,
			}),
			VideoCover:     videoCover,
			VideoTimestamp: &videoTimestamp,
		}).ToMessageMedia(),
	}

	got, err := resolveMessageMedia(context.Background(), mediaClient, nil, 1001, tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
		File:     tg.MakeTLInputFile(&tg.TLInputFile{Id: 8, Parts: 1, Name: "clip.mp4"}),
		MimeType: "video/mp4",
	}))
	if err != nil {
		t.Fatalf("resolveMessageMedia() error = %v", err)
	}
	if mediaClient.uploadedDocumentReq == nil || mediaClient.uploadedDocumentReq.OwnerId != 1001 {
		t.Fatalf("MediaUploadedDocumentMedia request = %#v", mediaClient.uploadedDocumentReq)
	}
	mediaDoc, ok := got.(*tg.TLMessageMediaDocument)
	if !ok {
		t.Fatalf("message media = %#v, want TLMessageMediaDocument", got)
	}
	if mediaDoc.VideoCover != videoCover {
		t.Fatalf("VideoCover = %#v, want returned cover %#v", mediaDoc.VideoCover, videoCover)
	}
	if mediaDoc.VideoTimestamp == nil || *mediaDoc.VideoTimestamp != videoTimestamp {
		t.Fatalf("VideoTimestamp = %#v, want %d", mediaDoc.VideoTimestamp, videoTimestamp)
	}
}

func TestResolveMessageMediaInputDocumentPreservesVideoCover(t *testing.T) {
	videoTimestamp := int32(23)
	videoCover := tg.MakeTLPhoto(&tg.TLPhoto{
		Id:         303,
		AccessHash: 404,
		Date:       123,
		Sizes:      []tg.PhotoSizeClazz{tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: "m", W: 320, H: 240, Size2: 1000})},
		DcId:       2,
	}).ToPhoto()
	mediaClient := &fakeResolveMediaClient{
		documentResp: tg.MakeTLDocument(&tg.TLDocument{
			Id:         42,
			AccessHash: 99,
			MimeType:   "video/mp4",
			Size2:      1000,
			DcId:       2,
		}).ToDocument(),
		photoResp: videoCover,
	}

	got, err := resolveMessageMedia(context.Background(), mediaClient, nil, 1001, tg.MakeTLInputMediaDocument(&tg.TLInputMediaDocument{
		Id:             tg.MakeTLInputDocument(&tg.TLInputDocument{Id: 42, AccessHash: 99}),
		VideoCover:     tg.MakeTLInputPhoto(&tg.TLInputPhoto{Id: 303, AccessHash: 404, FileReference: []byte("input-reference")}),
		VideoTimestamp: &videoTimestamp,
	}))
	if err != nil {
		t.Fatalf("resolveMessageMedia() error = %v", err)
	}
	if mediaClient.documentReq == nil || mediaClient.documentReq.Id != 42 {
		t.Fatalf("MediaGetDocument request = %#v", mediaClient.documentReq)
	}
	if mediaClient.photoReq == nil || mediaClient.photoReq.PhotoId != 303 {
		t.Fatalf("MediaGetPhoto request = %#v", mediaClient.photoReq)
	}
	mediaDoc, ok := got.(*tg.TLMessageMediaDocument)
	if !ok {
		t.Fatalf("message media = %#v, want TLMessageMediaDocument", got)
	}
	if mediaDoc.VideoCover != videoCover.Clazz {
		t.Fatalf("VideoCover = %#v, want resolved cover %#v", mediaDoc.VideoCover, videoCover.Clazz)
	}
	if mediaDoc.VideoTimestamp == nil || *mediaDoc.VideoTimestamp != videoTimestamp {
		t.Fatalf("VideoTimestamp = %#v, want %d", mediaDoc.VideoTimestamp, videoTimestamp)
	}
}

type fakeResolveMediaClient struct {
	photoReq             *mediapb.TLMediaGetPhoto
	photoResp            *tg.Photo
	photoErr             error
	uploadPhotoReq       *mediapb.TLMediaUploadPhotoFile
	uploadPhotoResp      *tg.Photo
	uploadPhotoErr       error
	uploadedDocumentReq  *mediapb.TLMediaUploadedDocumentMedia
	uploadedDocumentResp *tg.MessageMedia
	uploadedDocumentErr  error
	documentReq          *mediapb.TLMediaGetDocument
	documentResp         *tg.Document
	documentErr          error
}

func (f *fakeResolveMediaClient) MediaUploadPhotoFile(_ context.Context, in *mediapb.TLMediaUploadPhotoFile) (*tg.Photo, error) {
	f.uploadPhotoReq = in
	return f.uploadPhotoResp, f.uploadPhotoErr
}

func (f *fakeResolveMediaClient) MediaGetPhoto(_ context.Context, in *mediapb.TLMediaGetPhoto) (*tg.Photo, error) {
	f.photoReq = in
	return f.photoResp, f.photoErr
}

func (f *fakeResolveMediaClient) MediaGetPhotoSizeList(context.Context, *mediapb.TLMediaGetPhotoSizeList) (*mediapb.PhotoSizeList, error) {
	return nil, errors.New("unexpected MediaGetPhotoSizeList")
}

func (f *fakeResolveMediaClient) MediaUploadedDocumentMedia(_ context.Context, in *mediapb.TLMediaUploadedDocumentMedia) (*tg.MessageMedia, error) {
	f.uploadedDocumentReq = in
	return f.uploadedDocumentResp, f.uploadedDocumentErr
}

func (f *fakeResolveMediaClient) MediaGetDocument(_ context.Context, in *mediapb.TLMediaGetDocument) (*tg.Document, error) {
	f.documentReq = in
	return f.documentResp, f.documentErr
}
