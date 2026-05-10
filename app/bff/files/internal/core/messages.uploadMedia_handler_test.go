package core

import (
	"bytes"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMessagesUploadMediaInputMediaPhotoUsesMediaGetPhotoReference(t *testing.T) {
	fileReference := []byte("1234567890123456789012345")
	mediaClient := &fakeFilesMediaClient{
		photoResp: tg.MakeTLPhoto(&tg.TLPhoto{
			Id:            42,
			AccessHash:    99,
			FileReference: fileReference,
			Date:          123,
			Sizes:         []tg.PhotoSizeClazz{tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: "m", W: 320, H: 240, Size2: 1000})},
			DcId:          2,
		}).ToPhoto(),
	}
	core := newUploadGetFileTestCore(&fakeFilesDfsClient{}, mediaClient, false)

	got, err := core.MessagesUploadMedia(&tg.TLMessagesUploadMedia{
		Media: tg.MakeTLInputMediaPhoto(&tg.TLInputMediaPhoto{
			Id: tg.MakeTLInputPhoto(&tg.TLInputPhoto{Id: 42, AccessHash: 99, FileReference: []byte("input-reference")}),
		}),
	})
	if err != nil {
		t.Fatalf("MessagesUploadMedia() error = %v", err)
	}
	if mediaClient.photoReq == nil || mediaClient.photoReq.PhotoId != 42 {
		t.Fatalf("MediaGetPhoto request = %#v", mediaClient.photoReq)
	}
	mediaPhoto, ok := got.ToMessageMediaPhoto()
	if !ok {
		t.Fatalf("message media = %#v, want messageMediaPhoto", got)
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

func TestMessagesUploadMediaInputMediaUploadedPhotoReturns25ByteFileReference(t *testing.T) {
	fileReference := []byte("1234567890123456789012345")
	mediaClient := &fakeFilesMediaClient{
		uploadPhotoResp: tg.MakeTLPhoto(&tg.TLPhoto{
			Id:            43,
			AccessHash:    100,
			FileReference: fileReference,
			Date:          124,
			Sizes:         []tg.PhotoSizeClazz{tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: "m", W: 320, H: 240, Size2: 1000})},
			DcId:          2,
		}).ToPhoto(),
	}
	core := newUploadGetFileTestCore(&fakeFilesDfsClient{}, mediaClient, false)

	got, err := core.MessagesUploadMedia(&tg.TLMessagesUploadMedia{
		Media: tg.MakeTLInputMediaUploadedPhoto(&tg.TLInputMediaUploadedPhoto{}),
	})
	if err != nil {
		t.Fatalf("MessagesUploadMedia() error = %v", err)
	}
	if mediaClient.uploadPhotoReq == nil {
		t.Fatal("MediaUploadPhotoFile was not called")
	}
	mediaPhoto, ok := got.ToMessageMediaPhoto()
	if !ok {
		t.Fatalf("message media = %#v, want messageMediaPhoto", got)
	}
	photo, ok := mediaPhoto.Photo.(*tg.TLPhoto)
	if !ok {
		t.Fatalf("photo = %#v, want TLPhoto", mediaPhoto.Photo)
	}
	if len(photo.FileReference) != 25 {
		t.Fatalf("len(photo.FileReference) = %d, want 25", len(photo.FileReference))
	}
}
