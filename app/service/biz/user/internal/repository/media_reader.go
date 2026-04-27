package repository

import (
	"context"

	mediaclient "github.com/teamgram/teamgram-server/v2/app/service/media/client"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type MediaReader interface {
	GetPhoto(ctx context.Context, photoID int64) (*tg.Photo, error)
	GetDocument(ctx context.Context, documentID int64) (*tg.Document, error)
}

type mediaClient interface {
	MediaGetPhoto(ctx context.Context, in *media.TLMediaGetPhoto) (*tg.Photo, error)
	MediaGetDocument(ctx context.Context, in *media.TLMediaGetDocument) (*tg.Document, error)
}

type mediaReader struct {
	cli mediaClient
}

func NewMediaReader(cli mediaclient.MediaClient) MediaReader {
	return &mediaReader{cli: cli}
}

func (m *mediaReader) GetPhoto(ctx context.Context, photoID int64) (*tg.Photo, error) {
	photo, err := m.cli.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{PhotoId: photoID})
	if err != nil {
		return nil, wrapStorage("media.GetPhoto", err)
	}
	return photo, nil
}

func (m *mediaReader) GetDocument(ctx context.Context, documentID int64) (*tg.Document, error) {
	document, err := m.cli.MediaGetDocument(ctx, &media.TLMediaGetDocument{Id: documentID})
	if err != nil {
		return nil, wrapStorage("media.GetDocument", err)
	}
	return document, nil
}
