package rpc

import (
	"context"
	"errors"
	"fmt"

	mediaprocessorclient "github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/client"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
)

type MediaProcessorClient interface {
	ProcessPhoto(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessPhoto) (*mediaprocessor.ProcessedPhoto, error)
	ProcessGif(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessGif) (*mediaprocessor.ProcessedDocument, error)
	ProcessMp4(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessMp4) (*mediaprocessor.ProcessedDocument, error)
}

type MediaProcessorRPCClient struct {
	client mediaprocessorclient.MediaProcessorClient
}

func NewMediaProcessorClient(client mediaprocessorclient.MediaProcessorClient) *MediaProcessorRPCClient {
	return &MediaProcessorRPCClient{client: client}
}

func (c *MediaProcessorRPCClient) ProcessPhoto(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessPhoto) (*mediaprocessor.ProcessedPhoto, error) {
	if c == nil || c.client == nil {
		return nil, errMediaProcessorClientUnavailable("process photo")
	}
	return c.client.MediaProcessorProcessPhoto(ctx, in)
}

func (c *MediaProcessorRPCClient) ProcessGif(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessGif) (*mediaprocessor.ProcessedDocument, error) {
	if c == nil || c.client == nil {
		return nil, errMediaProcessorClientUnavailable("process gif")
	}
	return c.client.MediaProcessorProcessGif(ctx, in)
}

func (c *MediaProcessorRPCClient) ProcessMp4(ctx context.Context, in *mediaprocessor.TLMediaProcessorProcessMp4) (*mediaprocessor.ProcessedDocument, error) {
	if c == nil || c.client == nil {
		return nil, errMediaProcessorClientUnavailable("process mp4")
	}
	return c.client.MediaProcessorProcessMp4(ctx, in)
}

func errMediaProcessorClientUnavailable(op string) error {
	return fmt.Errorf("%s: %w", op, errors.New("mediaprocessor client unavailable"))
}
