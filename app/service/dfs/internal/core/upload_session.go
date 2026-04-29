package core

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository"
)

type UploadRef = repository.UploadRef
type DfsFileInfo = repository.DfsFileInfo

type uploadStateRepository interface {
	SaveUploadPart(ctx context.Context, ref UploadRef, partIndex int32, data []byte) error
	SaveUploadFileInfo(ctx context.Context, info *DfsFileInfo) error
	LoadUploadFileInfo(ctx context.Context, creator, fileID int64) (*DfsFileInfo, error)
	OpenUploadFileReader(ctx context.Context, info *DfsFileInfo) (io.ReadSeeker, error)
	SaveObjectCacheRef(ctx context.Context, objectID int64, info *DfsFileInfo) error
	LoadObjectCacheRef(ctx context.Context, objectID int64) (*DfsFileInfo, error)
}

type UploadSessionManager struct {
	repo uploadStateRepository
	now  func() time.Time
}

func NewUploadSessionManager(repo uploadStateRepository) *UploadSessionManager {
	return &UploadSessionManager{
		repo: repo,
		now:  time.Now,
	}
}

type WritePartCommand struct {
	Creator        int64
	FileID         int64
	FilePart       int32
	Bytes          []byte
	Big            bool
	FileTotalParts *int32
}

func (m *UploadSessionManager) WritePart(ctx context.Context, cmd WritePartCommand) error {
	if err := m.ensureRepo("write upload part"); err != nil {
		return err
	}
	if len(cmd.Bytes) == 0 {
		return dfs.ErrDfsInvalidFilePart
	}
	if cmd.Creator == 0 || cmd.FileID == 0 || cmd.FilePart < 0 {
		return dfs.ErrDfsInvalidFilePart
	}
	if cmd.FileTotalParts != nil && *cmd.FileTotalParts <= 0 {
		return dfs.ErrDfsInvalidFilePart
	}

	ref := UploadRef{Creator: cmd.Creator, FileID: cmd.FileID}
	if err := m.repo.SaveUploadPart(ctx, ref, cmd.FilePart, cmd.Bytes); err != nil {
		return wrapUploadStorage("save upload part", err)
	}

	info, err := m.repo.LoadUploadFileInfo(ctx, cmd.Creator, cmd.FileID)
	if err != nil && !errors.Is(err, dfs.ErrDfsFileNotFound) {
		return wrapUploadStorage("load upload file info", err)
	}
	if info == nil {
		info = &DfsFileInfo{Creator: cmd.Creator, FileID: cmd.FileID}
	}

	info.Creator = cmd.Creator
	info.FileID = cmd.FileID
	info.Big = cmd.Big
	info.Mtime = m.now().Unix()
	if cmd.FileTotalParts != nil {
		info.FileTotalParts = int(*cmd.FileTotalParts)
	} else if seenParts := int(cmd.FilePart + 1); seenParts > info.FileTotalParts {
		info.FileTotalParts = seenParts
	}
	partSize := len(cmd.Bytes)
	if cmd.FilePart == 0 {
		info.FirstFilePartSize = partSize
	}
	if cmd.FilePart == 1 {
		info.FilePartSize = partSize
	}
	if cmd.FileTotalParts != nil && cmd.FilePart+1 == *cmd.FileTotalParts {
		info.LastFilePartSize = partSize
	} else if int(cmd.FilePart+1) == info.FileTotalParts {
		info.LastFilePartSize = partSize
	}

	if err := m.repo.SaveUploadFileInfo(ctx, info); err != nil {
		return wrapUploadStorage("save upload file info", err)
	}
	return nil
}

func (m *UploadSessionManager) OpenUploadedFile(ctx context.Context, creator, fileID int64) (io.ReadSeeker, error) {
	if err := m.ensureRepo("open uploaded file"); err != nil {
		return nil, err
	}
	info, err := m.repo.LoadUploadFileInfo(ctx, creator, fileID)
	if err != nil {
		return nil, wrapUploadStorage("load upload file info", err)
	}
	reader, err := m.repo.OpenUploadFileReader(ctx, info)
	if err != nil {
		return nil, wrapUploadStorage("open upload file reader", err)
	}
	return reader, nil
}

func (m *UploadSessionManager) LoadUploadedFileInfo(ctx context.Context, creator, fileID int64) (*DfsFileInfo, error) {
	if err := m.ensureRepo("load uploaded file info"); err != nil {
		return nil, err
	}
	info, err := m.repo.LoadUploadFileInfo(ctx, creator, fileID)
	if err != nil {
		return nil, wrapUploadStorage("load upload file info", err)
	}
	return info, nil
}

func (m *UploadSessionManager) ReadUploadedFileRange(ctx context.Context, creator, fileID int64, offset, limit int64) ([]byte, error) {
	if offset < 0 || limit < 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	reader, err := m.OpenUploadedFile(ctx, creator, fileID)
	if err != nil {
		return nil, err
	}
	if _, err := reader.Seek(offset, io.SeekStart); err != nil {
		return nil, wrapUploadStorage("seek upload file reader", err)
	}
	if limit == 0 {
		return io.ReadAll(reader)
	}
	return io.ReadAll(io.LimitReader(reader, limit))
}

func (m *UploadSessionManager) SaveObjectCacheRef(ctx context.Context, objectID, creator, fileID int64) error {
	if err := m.ensureRepo("save object cache ref"); err != nil {
		return err
	}
	if objectID == 0 || creator == 0 || fileID == 0 {
		return dfs.ErrDfsInvalidArgument
	}
	info, err := m.repo.LoadUploadFileInfo(ctx, creator, fileID)
	if err != nil {
		return wrapUploadStorage("load upload file info", err)
	}
	if err := m.repo.SaveObjectCacheRef(ctx, objectID, info); err != nil {
		return wrapUploadStorage("save object cache ref", err)
	}
	return nil
}

func (m *UploadSessionManager) LoadObjectCacheRef(ctx context.Context, objectID int64) (*DfsFileInfo, error) {
	if err := m.ensureRepo("load object cache ref"); err != nil {
		return nil, err
	}
	if objectID == 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	info, err := m.repo.LoadObjectCacheRef(ctx, objectID)
	if err != nil {
		return nil, wrapUploadStorage("load object cache ref", err)
	}
	return info, nil
}

func (c *DfsCore) uploadSessions() *UploadSessionManager {
	if c.uploadSessionManager != nil {
		return c.uploadSessionManager
	}
	return NewUploadSessionManager(c.svcCtx.Repo)
}

func (m *UploadSessionManager) ensureRepo(op string) error {
	if m == nil || m.repo == nil {
		return dfs.WrapDfsStorage(op, errors.New("upload state repository unavailable"))
	}
	return nil
}

func wrapUploadStorage(op string, err error) error {
	if err == nil || errors.Is(err, dfs.ErrDfsFileNotFound) || errors.Is(err, dfs.ErrDfsInvalidArgument) || errors.Is(err, dfs.ErrDfsInvalidFilePart) {
		return err
	}
	if errors.Is(err, dfs.ErrDfsStorage) {
		return err
	}
	return dfs.WrapDfsStorage(op, err)
}
