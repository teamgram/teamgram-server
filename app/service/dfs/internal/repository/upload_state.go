package repository

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/spool"
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

type uploadStateModel interface {
	SaveUploadPart(ctx context.Context, creator, fileID int64, partIndex int32, data []byte) error
	SaveUploadFileInfo(ctx context.Context, info *xkv.DfsFileInfo) error
	LoadUploadFileInfo(ctx context.Context, creator, fileID int64) (*xkv.DfsFileInfo, error)
	OpenUploadFileReader(ctx context.Context, info *xkv.DfsFileInfo) (io.ReadSeeker, error)
	SaveObjectCacheRef(ctx context.Context, objectID, creator, fileID int64) error
	LoadObjectCacheRef(ctx context.Context, objectID int64) (*xkv.DfsFileInfo, error)
}

type uploadStateLifecycleModel interface {
	ScanOnStart(ctx context.Context, now time.Time) error
	CleanupExpiredUploadSessions(ctx context.Context, now time.Time) error
	MarkDraining(reason string)
	IsWritable() bool
}

type uploadStateMultipartCleanupModel interface {
	CleanupExpiredUploadSessionsWithAbort(ctx context.Context, now time.Time, abort func(ctx context.Context, objectKey, uploadID string) error) error
}

type uploadStateMultipartReconcileModel interface {
	ReconcileUploadingMultipartSegments(ctx context.Context, list func(ctx context.Context, objectKey, uploadID string) ([]spool.MultipartPart, error)) error
}

type uploadStateStartupScanModel interface {
	ScanOnStartWithoutCleanup(ctx context.Context, now time.Time) error
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
	if err := r.checkUploadNodeWritable("save upload part"); err != nil {
		return err
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
	if err := r.checkUploadNodeWritable("save upload file info"); err != nil {
		return err
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
	if err := r.checkUploadNodeWritable("save object cache ref"); err != nil {
		return err
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

func (r *Repository) ScanSpoolOnStart(ctx context.Context) error {
	if r == nil {
		return nil
	}
	now := time.Now()
	if startup, ok := r.uploadStateModel.(uploadStateStartupScanModel); ok {
		if err := startup.ScanOnStartWithoutCleanup(ctx, now); err != nil {
			return dfs.WrapDfsStorage("scan upload spool on start", err)
		}
		if err := r.ReconcileUploadingMultipartSegments(ctx); err != nil {
			return err
		}
		return r.CleanupExpiredUploadSessions(ctx, now)
	}
	if lifecycle, ok := r.uploadStateModel.(uploadStateLifecycleModel); ok {
		if err := lifecycle.ScanOnStart(ctx, now); err != nil {
			return dfs.WrapDfsStorage("scan upload spool on start", err)
		}
	}
	return nil
}

func (r *Repository) ReconcileUploadingMultipartSegments(ctx context.Context) error {
	if r == nil {
		return nil
	}
	reconcile, ok := r.uploadStateModel.(uploadStateMultipartReconcileModel)
	if !ok {
		return nil
	}
	list := func(ctx context.Context, objectKey, uploadID string) ([]spool.MultipartPart, error) {
		if r.objectStore == nil {
			return nil, errUploadStateStoreUnavailable
		}
		parts, err := r.objectStore.ListMultipartParts(ctx, r.documentsBucket, objectKey, uploadID)
		if err != nil {
			return nil, err
		}
		out := make([]spool.MultipartPart, 0, len(parts))
		for _, part := range parts {
			out = append(out, spool.MultipartPart{
				PartNumber: part.PartNumber,
				ETag:       part.ETag,
				Size:       part.Size,
			})
		}
		return out, nil
	}
	if err := reconcile.ReconcileUploadingMultipartSegments(ctx, list); err != nil {
		return dfs.WrapDfsStorage("reconcile uploading multipart segments", err)
	}
	return nil
}

func (r *Repository) CleanupExpiredUploadSessions(ctx context.Context, now time.Time) error {
	if r == nil {
		return nil
	}
	if cleanup, ok := r.uploadStateModel.(uploadStateMultipartCleanupModel); ok {
		abort := func(ctx context.Context, objectKey, uploadID string) error {
			if r.objectStore == nil {
				return errUploadStateStoreUnavailable
			}
			return r.objectStore.AbortMultipartUpload(ctx, r.documentsBucket, objectKey, uploadID)
		}
		if err := cleanup.CleanupExpiredUploadSessionsWithAbort(ctx, now, abort); err != nil {
			return dfs.WrapDfsStorage("cleanup expired upload sessions", err)
		}
		return nil
	}
	if lifecycle, ok := r.uploadStateModel.(uploadStateLifecycleModel); ok {
		if err := lifecycle.CleanupExpiredUploadSessions(ctx, now); err != nil {
			return dfs.WrapDfsStorage("cleanup expired upload sessions", err)
		}
	}
	return nil
}

func (r *Repository) MarkUploadNodeDraining(reason string) {
	if r == nil {
		return
	}
	r.uploadNodeMu.Lock()
	r.uploadNodeDraining = true
	r.uploadDrainReason = strings.TrimSpace(reason)
	r.uploadNodeMu.Unlock()
	if lifecycle, ok := r.uploadStateModel.(uploadStateLifecycleModel); ok {
		lifecycle.MarkDraining(reason)
	}
}

func (r *Repository) IsUploadNodeWritable() bool {
	if r == nil {
		return false
	}
	r.uploadNodeMu.RLock()
	draining := r.uploadNodeDraining
	r.uploadNodeMu.RUnlock()
	if draining {
		return false
	}
	if lifecycle, ok := r.uploadStateModel.(uploadStateLifecycleModel); ok {
		return lifecycle.IsWritable()
	}
	return true
}

var errUploadStateStoreUnavailable = errors.New("upload state store unavailable")

func (r *Repository) checkUploadNodeWritable(op string) error {
	if r.IsUploadNodeWritable() {
		return nil
	}
	reason := ""
	if r != nil {
		r.uploadNodeMu.RLock()
		reason = r.uploadDrainReason
		r.uploadNodeMu.RUnlock()
	}
	if reason == "" {
		reason = "upload node is not writable"
	}
	return dfs.WrapDfsStorage(op, fmt.Errorf("%w: %s", spool.ErrUploadSpoolNotWritable, reason))
}

func mapUploadStateError(op string, err error) error {
	if errors.Is(err, xkv.ErrUploadStateNotFound) || errors.Is(err, spool.ErrUploadStateNotFound) {
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
