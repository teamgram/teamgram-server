package rpc

import (
	"context"
	"errors"
	"testing"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeMediaClient struct {
	photoID int64
	err     error
}

func (f *fakeMediaClient) MediaGetPhoto(ctx context.Context, in *media.TLMediaGetPhoto) (*tg.Photo, error) {
	f.photoID = in.PhotoId
	if f.err != nil {
		return nil, f.err
	}
	return tg.MakeTLPhotoEmpty(nil).ToPhoto(), nil
}

func TestMediaReaderForwardsPhotoID(t *testing.T) {
	cli := &fakeMediaClient{}
	reader := NewMediaReader(cli)
	if _, err := reader.GetChatPhoto(context.Background(), 77); err != nil {
		t.Fatalf("GetChatPhoto error: %v", err)
	}
	if cli.photoID != 77 {
		t.Fatalf("photo id = %d", cli.photoID)
	}
}

func TestMediaReaderReturnsStorageErrorWithCause(t *testing.T) {
	cause := errors.New("media down")
	reader := NewMediaReader(&fakeMediaClient{err: cause})
	if _, err := reader.GetChatPhoto(context.Background(), 77); !errors.Is(err, chatpb.ErrChatStorage) || !errors.Is(err, cause) {
		t.Fatalf("expected storage error with cause, got %v", err)
	}
}
