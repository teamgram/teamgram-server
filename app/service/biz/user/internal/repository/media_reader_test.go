package repository

import (
	"context"
	"errors"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeMediaClient struct {
	photoReq    *media.TLMediaGetPhoto
	documentReq *media.TLMediaGetDocument
	err         error
}

func (f *fakeMediaClient) MediaGetPhoto(ctx context.Context, in *media.TLMediaGetPhoto) (*tg.Photo, error) {
	f.photoReq = in
	return &tg.Photo{}, f.err
}

func (f *fakeMediaClient) MediaGetDocument(ctx context.Context, in *media.TLMediaGetDocument) (*tg.Document, error) {
	f.documentReq = in
	return &tg.Document{}, f.err
}

func TestMediaReaderWrapsPhotoErrorWithStorageSemantic(t *testing.T) {
	cause := errors.New("media unavailable")
	reader := &mediaReader{cli: &fakeMediaClient{err: cause}}

	_, err := reader.GetPhoto(context.Background(), 123)

	if !errors.Is(err, userpb.ErrUserStorage) {
		t.Fatalf("expected storage semantic error, got %v", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("expected original cause to be preserved, got %v", err)
	}
}

func TestMediaReaderWrapsDocumentErrorWithStorageSemantic(t *testing.T) {
	cause := errors.New("media unavailable")
	reader := &mediaReader{cli: &fakeMediaClient{err: cause}}

	_, err := reader.GetDocument(context.Background(), 456)

	if !errors.Is(err, userpb.ErrUserStorage) {
		t.Fatalf("expected storage semantic error, got %v", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("expected original cause to be preserved, got %v", err)
	}
}

func TestMediaReaderForwardsRequestIDs(t *testing.T) {
	cli := &fakeMediaClient{}
	reader := &mediaReader{cli: cli}

	if _, err := reader.GetPhoto(context.Background(), 123); err != nil {
		t.Fatalf("GetPhoto returned error: %v", err)
	}
	if _, err := reader.GetDocument(context.Background(), 456); err != nil {
		t.Fatalf("GetDocument returned error: %v", err)
	}

	if cli.photoReq == nil || cli.photoReq.PhotoId != 123 {
		t.Fatalf("expected photo_id 123, got %#v", cli.photoReq)
	}
	if cli.documentReq == nil || cli.documentReq.Id != 456 {
		t.Fatalf("expected document id 456, got %#v", cli.documentReq)
	}
}
