package repository

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/xkv"
)

type UploadRef struct {
	Creator int64
	FileID  int64
}

type DfsFileInfo struct {
	Creator           int64
	FileID            int64
	Big               bool
	FileName          string
	FileTotalParts    int
	FirstFilePartSize int
	FilePartSize      int
	LastFilePartSize  int
	MimeType          string
	Mtime             int64
}

func (i *DfsFileInfo) FileSize() int64 {
	if i == nil || i.FileTotalParts <= 0 {
		return 0
	}
	if i.FileTotalParts == 1 {
		return int64(i.FirstFilePartSize)
	}
	if i.FileTotalParts == 2 {
		if i.LastFilePartSize > 0 {
			return int64(i.FirstFilePartSize + i.LastFilePartSize)
		}
		return int64(i.FirstFilePartSize + i.FilePartSize)
	}
	return int64(i.FirstFilePartSize + (i.FileTotalParts-2)*i.FilePartSize + i.LastFilePartSize)
}

func (r *Repository) SaveUploadPart(ctx context.Context, ref UploadRef, partIndex int32, data []byte) error {
	if r == nil || r.uploadStateModel == nil {
		return dfs.WrapDfsStorage("save upload part", errUploadStateStoreUnavailable)
	}
	if err := r.uploadStateModel.SaveUploadPart(ctx, ref.Creator, ref.FileID, partIndex, data); err != nil {
		return dfs.WrapDfsStorage("save upload part", err)
	}
	return nil
}

func (r *Repository) SaveUploadFileInfo(ctx context.Context, info *DfsFileInfo) error {
	if r == nil || r.uploadStateModel == nil {
		return dfs.WrapDfsStorage("save upload file info", errUploadStateStoreUnavailable)
	}
	if err := r.uploadStateModel.SaveUploadFileInfo(ctx, toXKVDfsFileInfo(info)); err != nil {
		return dfs.WrapDfsStorage("save upload file info", err)
	}
	return nil
}

func (r *Repository) LoadUploadFileInfo(ctx context.Context, creator, fileID int64) (*DfsFileInfo, error) {
	if r == nil || r.uploadStateModel == nil {
		return nil, dfs.WrapDfsStorage("load upload file info", errUploadStateStoreUnavailable)
	}
	info, err := r.uploadStateModel.LoadUploadFileInfo(ctx, creator, fileID)
	if err != nil {
		return nil, mapUploadStateError("load upload file info", err)
	}
	return fromXKVDfsFileInfo(info), nil
}

func (r *Repository) OpenUploadFileReader(ctx context.Context, info *DfsFileInfo) (io.ReadSeeker, error) {
	if r == nil || r.uploadStateModel == nil {
		return nil, dfs.WrapDfsStorage("open upload file reader", errUploadStateStoreUnavailable)
	}
	reader, err := r.uploadStateModel.OpenUploadFileReader(ctx, toXKVDfsFileInfo(info))
	if err != nil {
		return nil, mapUploadStateError("open upload file reader", err)
	}
	return reader, nil
}

func (r *Repository) SaveObjectCacheRef(ctx context.Context, objectID int64, info *DfsFileInfo) error {
	if r == nil || r.uploadStateModel == nil {
		return dfs.WrapDfsStorage("save object cache ref", errUploadStateStoreUnavailable)
	}
	if info == nil {
		return dfs.ErrDfsFileNotFound
	}
	if err := r.uploadStateModel.SaveObjectCacheRef(ctx, objectID, info.Creator, info.FileID); err != nil {
		return dfs.WrapDfsStorage("save object cache ref", err)
	}
	return nil
}

func (r *Repository) LoadObjectCacheRef(ctx context.Context, objectID int64) (*DfsFileInfo, error) {
	if r == nil || r.uploadStateModel == nil {
		return nil, dfs.WrapDfsStorage("load object cache ref", errUploadStateStoreUnavailable)
	}
	info, err := r.uploadStateModel.LoadObjectCacheRef(ctx, objectID)
	if err != nil {
		return nil, mapUploadStateError("load object cache ref", err)
	}
	return fromXKVDfsFileInfo(info), nil
}

var errUploadStateStoreUnavailable = errors.New("upload state store unavailable")

func mapUploadStateError(op string, err error) error {
	if errors.Is(err, xkv.ErrUploadStateNotFound) {
		return dfs.ErrDfsFileNotFound
	}
	return dfs.WrapDfsStorage(op, err)
}

func toXKVDfsFileInfo(info *DfsFileInfo) *xkv.DfsFileInfo {
	if info == nil {
		return nil
	}
	return &xkv.DfsFileInfo{
		Creator:           info.Creator,
		FileID:            info.FileID,
		Big:               info.Big,
		FileName:          info.FileName,
		FileTotalParts:    info.FileTotalParts,
		FirstFilePartSize: info.FirstFilePartSize,
		FilePartSize:      info.FilePartSize,
		LastFilePartSize:  info.LastFilePartSize,
		MimeType:          info.MimeType,
		Mtime:             info.Mtime,
	}
}

func fromXKVDfsFileInfo(info *xkv.DfsFileInfo) *DfsFileInfo {
	if info == nil {
		return nil
	}
	return &DfsFileInfo{
		Creator:           info.Creator,
		FileID:            info.FileID,
		Big:               info.Big,
		FileName:          info.FileName,
		FileTotalParts:    info.FileTotalParts,
		FirstFilePartSize: info.FirstFilePartSize,
		FilePartSize:      info.FilePartSize,
		LastFilePartSize:  info.LastFilePartSize,
		MimeType:          info.MimeType,
		Mtime:             info.Mtime,
	}
}

func (i *DfsFileInfo) String() string {
	if i == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DfsFileInfo{creator:%d file_id:%d total_parts:%d size:%d}", i.Creator, i.FileID, i.FileTotalParts, i.FileSize())
}
