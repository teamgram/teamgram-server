package rpc

import (
	"context"

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type MediaReader interface {
	GetChatPhoto(ctx context.Context, photoID int64) (*tg.Photo, error)
}

type mediaClient interface {
	MediaGetPhoto(ctx context.Context, in *media.TLMediaGetPhoto) (*tg.Photo, error)
}

type mediaReader struct {
	cli mediaClient
}

func NewMediaReader(cli mediaClient) MediaReader {
	return &mediaReader{cli: cli}
}

func (m *mediaReader) GetChatPhoto(ctx context.Context, photoID int64) (*tg.Photo, error) {
	photo, err := m.cli.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{PhotoId: photoID})
	if err != nil {
		return nil, chatpb.WrapChatStorage("media.GetChatPhoto", err)
	}
	return photo, nil
}
