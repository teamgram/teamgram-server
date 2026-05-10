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

func TestResolveMessageMediaUploadedDocumentForwardsFullInputToMedia(t *testing.T) {
	videoTimestamp := int32(19)
	ttlSeconds := int32(30)
	file := tg.MakeTLInputFile(&tg.TLInputFile{Id: 8, Parts: 2, Name: "clip.mp4", Md5Checksum: "sum"})
	thumb := tg.MakeTLInputFile(&tg.TLInputFile{Id: 9, Parts: 1, Name: "thumb.jpg"})
	attributes := []tg.DocumentAttributeClazz{
		tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{SupportsStreaming: true, Duration: 3, W: 640, H: 360}),
		tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "clip.mp4"}),
	}
	stickers := []tg.InputDocumentClazz{
		tg.MakeTLInputDocument(&tg.TLInputDocument{Id: 77, AccessHash: 88, FileReference: []byte("sticker-reference")}),
	}
	videoCover := tg.MakeTLInputPhoto(&tg.TLInputPhoto{Id: 303, AccessHash: 404, FileReference: []byte("cover-reference")})
	input := tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
		NosoundVideo:   true,
		ForceFile:      true,
		Spoiler:        true,
		File:           file,
		Thumb:          thumb,
		MimeType:       "video/mp4",
		Attributes:     attributes,
		Stickers:       stickers,
		VideoCover:     videoCover,
		VideoTimestamp: &videoTimestamp,
		TtlSeconds:     &ttlSeconds,
	})
	returnedMedia := tg.MakeTLMessageMediaDocument(&tg.TLMessageMediaDocument{
		Document: tg.MakeTLDocument(&tg.TLDocument{
			Id:         42,
			AccessHash: 99,
			MimeType:   "video/mp4",
			Size2:      1000,
			DcId:       2,
		}),
		TtlSeconds: &ttlSeconds,
	}).ToMessageMedia()
	mediaClient := &fakeResolveMediaClient{uploadedDocumentResp: returnedMedia}

	got, err := resolveMessageMedia(context.Background(), mediaClient, nil, 1001, input)
	if err != nil {
		t.Fatalf("resolveMessageMedia() error = %v", err)
	}
	if got != returnedMedia.Clazz {
		t.Fatalf("message media = %#v, want returned media %#v", got, returnedMedia.Clazz)
	}
	if mediaClient.uploadedDocumentReq == nil {
		t.Fatal("MediaUploadedDocumentMedia request is nil")
	}
	if mediaClient.uploadedDocumentReq.OwnerId != 1001 {
		t.Fatalf("OwnerId = %d, want 1001", mediaClient.uploadedDocumentReq.OwnerId)
	}
	forwarded, ok := mediaClient.uploadedDocumentReq.Media.(*tg.TLInputMediaUploadedDocument)
	if !ok {
		t.Fatalf("Media = %T, want *tg.TLInputMediaUploadedDocument", mediaClient.uploadedDocumentReq.Media)
	}
	if forwarded != input {
		t.Fatalf("Media pointer = %#v, want original input %#v", forwarded, input)
	}
	if forwarded.File != file || forwarded.Thumb != thumb || forwarded.VideoCover != videoCover {
		t.Fatalf("forwarded file fields = file:%#v thumb:%#v cover:%#v", forwarded.File, forwarded.Thumb, forwarded.VideoCover)
	}
	if forwarded.MimeType != "video/mp4" {
		t.Fatalf("MimeType = %q, want video/mp4", forwarded.MimeType)
	}
	if len(forwarded.Attributes) != len(attributes) || forwarded.Attributes[0] != attributes[0] || forwarded.Attributes[1] != attributes[1] {
		t.Fatalf("Attributes = %#v, want %#v", forwarded.Attributes, attributes)
	}
	if len(forwarded.Stickers) != len(stickers) || forwarded.Stickers[0] != stickers[0] {
		t.Fatalf("Stickers = %#v, want %#v", forwarded.Stickers, stickers)
	}
	if forwarded.VideoTimestamp != &videoTimestamp || forwarded.TtlSeconds != &ttlSeconds {
		t.Fatalf("optional timestamps = video:%#v ttl:%#v", forwarded.VideoTimestamp, forwarded.TtlSeconds)
	}
	if !forwarded.ForceFile || !forwarded.Spoiler || !forwarded.NosoundVideo {
		t.Fatalf("flags = force_file:%t spoiler:%t nosound_video:%t", forwarded.ForceFile, forwarded.Spoiler, forwarded.NosoundVideo)
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

func TestResolveMessageMediaUploadedDocumentPropagatesMediaTransformError(t *testing.T) {
	mediaClient := &fakeResolveMediaClient{uploadedDocumentErr: mediapb.ErrMediaInvalidUploadedFile}

	got, err := resolveMessageMedia(context.Background(), mediaClient, nil, 1001, tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
		File:     tg.MakeTLInputFile(&tg.TLInputFile{Id: 8, Parts: 1, Name: "clip.mp4"}),
		MimeType: "video/mp4",
	}))
	if !errors.Is(err, mediapb.ErrMediaInvalidUploadedFile) {
		t.Fatalf("resolveMessageMedia() error = %v, want ErrMediaInvalidUploadedFile", err)
	}
	if got != nil {
		t.Fatalf("message media = %#v, want nil", got)
	}
}

func TestMapMediaResolveError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want error
	}{
		{name: "nil", err: nil, want: nil},
		{name: "media empty", err: tg.ErrMediaEmpty, want: tg.ErrMediaEmpty},
		{name: "invalid uploaded file", err: mediapb.ErrMediaInvalidUploadedFile, want: tg.ErrMediaEmpty},
		{name: "wrapped invalid uploaded file", err: errors.Join(errors.New("transform failed"), mediapb.ErrMediaInvalidUploadedFile), want: tg.ErrMediaEmpty},
		{name: "downstream", err: mediapb.ErrMediaDownstream, want: tg.ErrInternalServerError},
		{name: "storage", err: mediapb.ErrMediaStorage, want: tg.ErrInternalServerError},
		{name: "unknown", err: errors.New("boom"), want: tg.ErrInternalServerError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := mapMediaResolveError(test.err)
			if got != test.want {
				t.Fatalf("mapMediaResolveError() = %v, want %v", got, test.want)
			}
		})
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
