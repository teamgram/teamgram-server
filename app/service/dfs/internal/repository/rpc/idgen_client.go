package rpc

import (
	"context"
	"errors"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	idgenclient "github.com/teamgram/teamgram-server/v2/app/service/idgen/client"
	"github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
)

type IDGenerator interface {
	NextPhotoID(ctx context.Context) (int64, error)
	NextDocumentID(ctx context.Context) (int64, error)
	NextEncryptedFileID(ctx context.Context) (int64, error)
}

type IDGenClient struct {
	client idgenclient.IdgenClient
}

func NewIDGenClient(client idgenclient.IdgenClient) *IDGenClient {
	return &IDGenClient{client: client}
}

func (c *IDGenClient) NextPhotoID(ctx context.Context) (int64, error) {
	return c.nextID(ctx, "next photo id")
}

func (c *IDGenClient) NextDocumentID(ctx context.Context) (int64, error) {
	return c.nextID(ctx, "next document id")
}

func (c *IDGenClient) NextEncryptedFileID(ctx context.Context) (int64, error) {
	return c.nextID(ctx, "next encrypted file id")
}

func (c *IDGenClient) nextID(ctx context.Context, op string) (int64, error) {
	if c == nil || c.client == nil {
		return 0, dfs.WrapDfsDownstream(op, errors.New("idgen client unavailable"))
	}
	id, err := c.client.IdgenNextId(ctx, &idgen.TLIdgenNextId{})
	if err != nil {
		return 0, dfs.WrapDfsDownstream(op, err)
	}
	if id == nil {
		return 0, dfs.WrapDfsDownstream(op, errors.New("idgen returned nil"))
	}
	return id.V, nil
}
