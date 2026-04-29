package rpc

import (
	"context"
	"errors"
	"fmt"

	dfsclient "github.com/teamgram/teamgram-server/v2/app/service/dfs/client"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type DfsMediaClient interface {
	UploadPhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadPhotoFileV2) (*tg.Photo, error)
	UploadProfilePhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadProfilePhotoFileV2) (*tg.Photo, error)
	UploadDocumentFileV2(ctx context.Context, in *dfs.TLDfsUploadDocumentFileV2) (*tg.Document, error)
	UploadGifDocumentMedia(ctx context.Context, in *dfs.TLDfsUploadGifDocumentMedia) (*tg.Document, error)
	UploadMp4DocumentMedia(ctx context.Context, in *dfs.TLDfsUploadMp4DocumentMedia) (*tg.Document, error)
	UploadedProfilePhoto(ctx context.Context, in *dfs.TLDfsUploadedProfilePhoto) (*tg.Photo, error)
}

type DFSClient struct {
	client dfsclient.DfsClient
}

func NewDFSClient(client dfsclient.DfsClient) *DFSClient {
	return &DFSClient{client: client}
}

func (c *DFSClient) UploadPhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadPhotoFileV2) (*tg.Photo, error) {
	if c == nil || c.client == nil {
		return nil, errDFSClientUnavailable("upload photo")
	}
	return c.client.DfsUploadPhotoFileV2(ctx, in)
}

func (c *DFSClient) UploadProfilePhotoFileV2(ctx context.Context, in *dfs.TLDfsUploadProfilePhotoFileV2) (*tg.Photo, error) {
	if c == nil || c.client == nil {
		return nil, errDFSClientUnavailable("upload profile photo")
	}
	return c.client.DfsUploadProfilePhotoFileV2(ctx, in)
}

func (c *DFSClient) UploadDocumentFileV2(ctx context.Context, in *dfs.TLDfsUploadDocumentFileV2) (*tg.Document, error) {
	if c == nil || c.client == nil {
		return nil, errDFSClientUnavailable("upload document")
	}
	return c.client.DfsUploadDocumentFileV2(ctx, in)
}

func (c *DFSClient) UploadGifDocumentMedia(ctx context.Context, in *dfs.TLDfsUploadGifDocumentMedia) (*tg.Document, error) {
	if c == nil || c.client == nil {
		return nil, errDFSClientUnavailable("upload gif document")
	}
	return c.client.DfsUploadGifDocumentMedia(ctx, in)
}

func (c *DFSClient) UploadMp4DocumentMedia(ctx context.Context, in *dfs.TLDfsUploadMp4DocumentMedia) (*tg.Document, error) {
	if c == nil || c.client == nil {
		return nil, errDFSClientUnavailable("upload mp4 document")
	}
	return c.client.DfsUploadMp4DocumentMedia(ctx, in)
}

func (c *DFSClient) UploadedProfilePhoto(ctx context.Context, in *dfs.TLDfsUploadedProfilePhoto) (*tg.Photo, error) {
	if c == nil || c.client == nil {
		return nil, errDFSClientUnavailable("uploaded profile photo")
	}
	return c.client.DfsUploadedProfilePhoto(ctx, in)
}

func errDFSClientUnavailable(op string) error {
	return fmt.Errorf("%s: %w", op, errors.New("dfs client unavailable"))
}
