package core

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/objectstore"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fileObjectRepository interface {
	CommitUpload(ctx context.Context, uploadSessionID string, ownerID int64, file tg.InputFileClazz, purpose string) (*dfs.FileFinalizedObject, error)
	PutInternalFile(ctx context.Context, ownerID int64, purpose, fileName, mimeType string, data []byte) (*dfs.FileFinalizedObject, error)
	ReadByLease(ctx context.Context, readLease []byte, offset int64, limit int32) ([]byte, int32, error)
	HashesByLease(ctx context.Context, readLease []byte, offset int64, limit int32) ([]objectstore.HashChunk, error)
}

func (c *DfsCore) fileObjects() fileObjectRepository {
	if c.fileObjectRepository != nil {
		return c.fileObjectRepository
	}
	return c.svcCtx.Repo
}
