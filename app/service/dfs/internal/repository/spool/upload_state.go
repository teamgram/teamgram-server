package spool

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/xkv"
)

const (
	defaultLocalShardCount = 256
	defaultSegmentSize     = 16 * 1024 * 1024
	metadataFileName       = "metadata.json"
)

var ErrUploadStateNotFound = errors.New("dfs upload spool state not found")
var ErrUploadSpoolNotWritable = errors.New("dfs upload spool is not writable")

var nodeIDPattern = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)

const (
	SegmentStatusUploading = "uploading"
	SegmentStatusUploaded  = "uploaded"
)

type SegmentState struct {
	SegmentNo           int64  `json:"segment_no"`
	Status              string `json:"status"`
	MultipartUploadID   string `json:"multipart_upload_id"`
	MultipartPartNumber int    `json:"multipart_part_number"`
	ObjectKey           string `json:"object_key,omitempty"`
	ETag                string `json:"etag,omitempty"`
	Checksum            string `json:"checksum"`
	Size                int64  `json:"size"`
}

type MultipartPart struct {
	PartNumber int
	ETag       string
	Size       int64
}

type BuiltSegment struct {
	State         SegmentState
	Path          string
	FirstBytes    []byte
	AlreadyDone   bool
	MissingPart   bool
	MissingPartNo int32
}

type ReplayedSegment struct {
	FirstBytes    []byte
	Size          int64
	Checksum      string
	MissingPart   bool
	MissingPartNo int32
}

type UploadStateModel struct {
	rootDir         string
	nodeID          string
	segmentSize     int64
	localShardCount int
	partTTLSeconds  int64
	mu              sync.RWMutex
	writable        bool
	draining        bool
	drainReason     string
	probeWritable   func(ctx context.Context, nodeRoot string) error
}

type objectCacheRef struct {
	Creator int64 `json:"creator"`
	FileID  int64 `json:"file_id"`
}

func NewUploadStateModel(conf config.UploadSpoolConf) (*UploadStateModel, error) {
	if conf.RootDir == "" {
		return nil, errors.New("upload spool root dir is empty")
	}
	rootDir, err := filepath.Abs(conf.RootDir)
	if err != nil {
		return nil, fmt.Errorf("upload_spool.NewUploadStateModel abs root: %w", err)
	}
	shardCount := conf.LocalShardCount
	if shardCount <= 0 {
		shardCount = defaultLocalShardCount
	}
	segmentSize := conf.SegmentSize
	if segmentSize <= 0 {
		segmentSize = defaultSegmentSize
	}
	if err := mkdirAllDurable(rootDir, 0o755); err != nil {
		return nil, fmt.Errorf("upload_spool.NewUploadStateModel mkdir root: %w", err)
	}
	nodeID, err := loadOrCreateNodeID(rootDir, conf.NodeIDFile)
	if err != nil {
		return nil, err
	}
	nodeRoot, err := safeJoin(rootDir, nodeID)
	if err != nil {
		return nil, fmt.Errorf("upload_spool.NewUploadStateModel node root: %w", err)
	}
	if err := mkdirAllDurable(nodeRoot, 0o755); err != nil {
		return nil, fmt.Errorf("upload_spool.NewUploadStateModel mkdir node root: %w", err)
	}
	return &UploadStateModel{
		rootDir:         rootDir,
		nodeID:          nodeID,
		segmentSize:     segmentSize,
		localShardCount: shardCount,
		partTTLSeconds:  conf.PartTTLSeconds,
		writable:        true,
		probeWritable:   defaultProbeWritable,
	}, nil
}

func (m *UploadStateModel) NodeID() string {
	if m == nil {
		return ""
	}
	return m.nodeID
}

func (m *UploadStateModel) SegmentSize() int64 {
	if m == nil || m.segmentSize <= 0 {
		return defaultSegmentSize
	}
	return m.segmentSize
}

func (m *UploadStateModel) SaveUploadPart(ctx context.Context, creator, fileID int64, partIndex int32, data []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := m.checkWritable(); err != nil {
		return err
	}
	if partIndex < 0 {
		return fmt.Errorf("upload_spool.SaveUploadPart invalid part index %d", partIndex)
	}
	dir, err := m.sessionDir(creator, fileID)
	if err != nil {
		return err
	}
	if err := mkdirAllDurable(dir, 0o755); err != nil {
		m.markUnwritable(fmt.Sprintf("save upload part mkdir session failed: %v", err))
		return fmt.Errorf("upload_spool.SaveUploadPart mkdir session: %w", err)
	}
	if err := writeFileAtomically(filepath.Join(dir, partFileName(partIndex)), data, 0o644); err != nil {
		m.markUnwritable(fmt.Sprintf("save upload part write failed: %v", err))
		return fmt.Errorf("upload_spool.SaveUploadPart write part: %w", err)
	}
	return nil
}

func (m *UploadStateModel) SaveUploadFileInfo(ctx context.Context, info *xkv.DfsFileInfo) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := m.checkWritable(); err != nil {
		return err
	}
	if info == nil {
		return nil
	}
	dir, err := m.sessionDir(info.Creator, info.FileID)
	if err != nil {
		return err
	}
	if err := mkdirAllDurable(dir, 0o755); err != nil {
		m.markUnwritable(fmt.Sprintf("save upload file info mkdir session failed: %v", err))
		return fmt.Errorf("upload_spool.SaveUploadFileInfo mkdir session: %w", err)
	}
	data, err := json.Marshal(info)
	if err != nil {
		return fmt.Errorf("upload_spool.SaveUploadFileInfo marshal: %w", err)
	}
	if err := writeFileAtomically(filepath.Join(dir, metadataFileName), data, 0o644); err != nil {
		m.markUnwritable(fmt.Sprintf("save upload file info write failed: %v", err))
		return fmt.Errorf("upload_spool.SaveUploadFileInfo write metadata: %w", err)
	}
	return nil
}

func (m *UploadStateModel) LoadUploadFileInfo(ctx context.Context, creator, fileID int64) (*xkv.DfsFileInfo, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	dir, err := m.sessionDir(creator, fileID)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(filepath.Join(dir, metadataFileName))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrUploadStateNotFound
		}
		return nil, fmt.Errorf("upload_spool.LoadUploadFileInfo read metadata: %w", err)
	}
	var info xkv.DfsFileInfo
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, fmt.Errorf("upload_spool.LoadUploadFileInfo unmarshal metadata: %w", err)
	}
	if info.Creator == 0 {
		info.Creator = creator
	}
	if info.FileID == 0 {
		info.FileID = fileID
	}
	return &info, nil
}

func (m *UploadStateModel) OpenUploadFileReader(ctx context.Context, info *xkv.DfsFileInfo) (io.ReadSeeker, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if info == nil || info.FileTotalParts <= 0 {
		return nil, ErrUploadStateNotFound
	}
	dir, err := m.sessionDir(info.Creator, info.FileID)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	// Task 5a keeps compatibility with the existing io.ReadSeeker contract by
	// buffering all saved parts; streaming segment lifecycle is deferred to 5d.
	for i := int32(0); i < int32(info.FileTotalParts); i++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		part, err := os.ReadFile(filepath.Join(dir, partFileName(i)))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return nil, ErrUploadStateNotFound
			}
			return nil, fmt.Errorf("upload_spool.OpenUploadFileReader read part %d: %w", i, err)
		}
		buf.Write(part)
	}
	return bytes.NewReader(buf.Bytes()), nil
}

func (m *UploadStateModel) SegmentCount(info *xkv.DfsFileInfo) (int64, error) {
	partsPerSegment, err := m.partsPerSegment(info)
	if err != nil {
		return 0, err
	}
	if info.FileTotalParts <= 0 {
		return 0, ErrUploadStateNotFound
	}
	return int64((info.FileTotalParts + partsPerSegment - 1) / partsPerSegment), nil
}

func (m *UploadStateModel) BuildSegment(ctx context.Context, info *xkv.DfsFileInfo, segmentNo int64, uploadID string, partNumber int) (BuiltSegment, error) {
	return m.buildSegment(ctx, info, segmentNo, uploadID, partNumber, false)
}

func (m *UploadStateModel) RebuildSegment(ctx context.Context, info *xkv.DfsFileInfo, segmentNo int64, uploadID string, partNumber int) (BuiltSegment, error) {
	return m.buildSegment(ctx, info, segmentNo, uploadID, partNumber, true)
}

func (m *UploadStateModel) buildSegment(ctx context.Context, info *xkv.DfsFileInfo, segmentNo int64, uploadID string, partNumber int, force bool) (BuiltSegment, error) {
	if err := ctx.Err(); err != nil {
		return BuiltSegment{}, err
	}
	if info == nil || info.FileTotalParts <= 0 || segmentNo < 0 || uploadID == "" || partNumber <= 0 {
		return BuiltSegment{}, ErrUploadStateNotFound
	}
	existing, err := m.LoadSegmentState(ctx, info, segmentNo)
	if !force && err == nil && existing.Status == SegmentStatusUploaded && existing.MultipartUploadID == uploadID {
		return BuiltSegment{State: existing, AlreadyDone: true}, nil
	} else if err != nil && !errors.Is(err, ErrUploadStateNotFound) {
		return BuiltSegment{}, err
	}
	partsPerSegment, err := m.partsPerSegment(info)
	if err != nil {
		return BuiltSegment{}, err
	}
	startPart := int(segmentNo) * partsPerSegment
	if startPart >= info.FileTotalParts {
		return BuiltSegment{}, ErrUploadStateNotFound
	}
	endPart := startPart + partsPerSegment
	if endPart > info.FileTotalParts {
		endPart = info.FileTotalParts
	}
	path, err := m.SegmentPath(info, segmentNo)
	if err != nil {
		return BuiltSegment{}, err
	}
	if err := mkdirAllDurable(filepath.Dir(path), 0o755); err != nil {
		return BuiltSegment{}, fmt.Errorf("upload_spool.BuildSegment mkdir segment: %w", err)
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), "."+filepath.Base(path)+".tmp-*")
	if err != nil {
		return BuiltSegment{}, fmt.Errorf("upload_spool.BuildSegment create temp: %w", err)
	}
	tmpName := tmp.Name()
	closed := false
	defer func() {
		if !closed {
			_ = tmp.Close()
		}
		_ = os.Remove(tmpName)
	}()
	sum := sha256.New()
	var size int64
	var firstBytes []byte
	dir, err := m.sessionDir(info.Creator, info.FileID)
	if err != nil {
		return BuiltSegment{}, err
	}
	for partIndex := startPart; partIndex < endPart; partIndex++ {
		if err := ctx.Err(); err != nil {
			return BuiltSegment{}, err
		}
		partPath := filepath.Join(dir, partFileName(int32(partIndex)))
		part, err := os.Open(partPath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return BuiltSegment{MissingPart: true, MissingPartNo: int32(partIndex)}, ErrUploadStateNotFound
			}
			return BuiltSegment{}, fmt.Errorf("upload_spool.BuildSegment open part %d: %w", partIndex, err)
		}
		n, copyErr := copyToWriters(tmp, sum, part, &firstBytes)
		closeErr := part.Close()
		if copyErr != nil {
			return BuiltSegment{}, fmt.Errorf("upload_spool.BuildSegment copy part %d: %w", partIndex, copyErr)
		}
		if closeErr != nil {
			return BuiltSegment{}, fmt.Errorf("upload_spool.BuildSegment close part %d: %w", partIndex, closeErr)
		}
		size += n
	}
	if size <= 0 {
		return BuiltSegment{}, ErrUploadStateNotFound
	}
	if err := tmp.Sync(); err != nil {
		return BuiltSegment{}, fmt.Errorf("upload_spool.BuildSegment sync segment: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return BuiltSegment{}, fmt.Errorf("upload_spool.BuildSegment close segment: %w", err)
	}
	closed = true
	if err := os.Rename(tmpName, path); err != nil {
		return BuiltSegment{}, fmt.Errorf("upload_spool.BuildSegment rename segment: %w", err)
	}
	if err := syncDir(filepath.Dir(path)); err != nil {
		return BuiltSegment{}, fmt.Errorf("upload_spool.BuildSegment sync segment dir: %w", err)
	}
	state := SegmentState{
		SegmentNo:           segmentNo,
		Status:              SegmentStatusUploading,
		MultipartUploadID:   uploadID,
		MultipartPartNumber: partNumber,
		Checksum:            hex.EncodeToString(sum.Sum(nil)),
		Size:                size,
	}
	return BuiltSegment{State: state, Path: path, FirstBytes: firstBytes}, nil
}

func (m *UploadStateModel) ReplaySegment(ctx context.Context, info *xkv.DfsFileInfo, segmentNo int64, dst io.Writer) (ReplayedSegment, error) {
	if err := ctx.Err(); err != nil {
		return ReplayedSegment{}, err
	}
	if info == nil || info.FileTotalParts <= 0 || segmentNo < 0 || dst == nil {
		return ReplayedSegment{}, ErrUploadStateNotFound
	}
	partsPerSegment, err := m.partsPerSegment(info)
	if err != nil {
		return ReplayedSegment{}, err
	}
	startPart := int(segmentNo) * partsPerSegment
	if startPart >= info.FileTotalParts {
		return ReplayedSegment{}, ErrUploadStateNotFound
	}
	endPart := startPart + partsPerSegment
	if endPart > info.FileTotalParts {
		endPart = info.FileTotalParts
	}
	dir, err := m.sessionDir(info.Creator, info.FileID)
	if err != nil {
		return ReplayedSegment{}, err
	}
	sum := sha256.New()
	var size int64
	var firstBytes []byte
	for partIndex := startPart; partIndex < endPart; partIndex++ {
		if err := ctx.Err(); err != nil {
			return ReplayedSegment{}, err
		}
		partPath := filepath.Join(dir, partFileName(int32(partIndex)))
		part, err := os.Open(partPath)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return ReplayedSegment{MissingPart: true, MissingPartNo: int32(partIndex)}, ErrUploadStateNotFound
			}
			return ReplayedSegment{}, fmt.Errorf("upload_spool.ReplaySegment open part %d: %w", partIndex, err)
		}
		n, copyErr := copyToWriters(dst, sum, part, &firstBytes)
		closeErr := part.Close()
		if copyErr != nil {
			return ReplayedSegment{}, fmt.Errorf("upload_spool.ReplaySegment copy part %d: %w", partIndex, copyErr)
		}
		if closeErr != nil {
			return ReplayedSegment{}, fmt.Errorf("upload_spool.ReplaySegment close part %d: %w", partIndex, closeErr)
		}
		size += n
	}
	if size <= 0 {
		return ReplayedSegment{}, ErrUploadStateNotFound
	}
	return ReplayedSegment{
		FirstBytes: firstBytes,
		Size:       size,
		Checksum:   hex.EncodeToString(sum.Sum(nil)),
	}, nil
}

func (m *UploadStateModel) WriteSegmentState(ctx context.Context, info *xkv.DfsFileInfo, state SegmentState) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if info == nil {
		return ErrUploadStateNotFound
	}
	data, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("upload_spool.WriteSegmentState marshal: %w", err)
	}
	path, err := m.segmentStatePath(info, state.SegmentNo)
	if err != nil {
		return err
	}
	if err := writeFileAtomically(path, data, 0o644); err != nil {
		return fmt.Errorf("upload_spool.WriteSegmentState write: %w", err)
	}
	return nil
}

func (m *UploadStateModel) LoadSegmentState(ctx context.Context, info *xkv.DfsFileInfo, segmentNo int64) (SegmentState, error) {
	if err := ctx.Err(); err != nil {
		return SegmentState{}, err
	}
	path, err := m.segmentStatePath(info, segmentNo)
	if err != nil {
		return SegmentState{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return SegmentState{}, ErrUploadStateNotFound
		}
		return SegmentState{}, fmt.Errorf("upload_spool.LoadSegmentState read: %w", err)
	}
	var state SegmentState
	if err := json.Unmarshal(data, &state); err != nil {
		return SegmentState{}, fmt.Errorf("upload_spool.LoadSegmentState unmarshal: %w", err)
	}
	return state, nil
}

func (m *UploadStateModel) LoadSegmentStates(ctx context.Context, info *xkv.DfsFileInfo) ([]SegmentState, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if info == nil {
		return nil, ErrUploadStateNotFound
	}
	dir, err := m.sessionDir(info.Creator, info.FileID)
	if err != nil {
		return nil, err
	}
	return loadSegmentStatesFromDir(dir)
}

func (m *UploadStateModel) SegmentPath(info *xkv.DfsFileInfo, segmentNo int64) (string, error) {
	if info == nil {
		return "", ErrUploadStateNotFound
	}
	dir, err := m.sessionDir(info.Creator, info.FileID)
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, segmentFileName(segmentNo)), nil
}

func (m *UploadStateModel) DeleteSegmentBytes(ctx context.Context, info *xkv.DfsFileInfo, segmentNo int64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	path, err := m.SegmentPath(info, segmentNo)
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("upload_spool.DeleteSegmentBytes remove: %w", err)
	}
	if err := syncDir(filepath.Dir(path)); err != nil {
		return fmt.Errorf("upload_spool.DeleteSegmentBytes sync dir: %w", err)
	}
	return nil
}

func (m *UploadStateModel) SaveObjectCacheRef(ctx context.Context, objectID, creator, fileID int64) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := m.checkWritable(); err != nil {
		return err
	}
	dir, err := m.cacheRefDir(objectID)
	if err != nil {
		return err
	}
	if err := mkdirAllDurable(dir, 0o755); err != nil {
		m.markUnwritable(fmt.Sprintf("save object cache ref mkdir failed: %v", err))
		return fmt.Errorf("upload_spool.SaveObjectCacheRef mkdir cache ref: %w", err)
	}
	data, err := json.Marshal(objectCacheRef{Creator: creator, FileID: fileID})
	if err != nil {
		return fmt.Errorf("upload_spool.SaveObjectCacheRef marshal: %w", err)
	}
	if err := writeFileAtomically(filepath.Join(dir, strconv.FormatInt(objectID, 10)+".json"), data, 0o644); err != nil {
		m.markUnwritable(fmt.Sprintf("save object cache ref write failed: %v", err))
		return fmt.Errorf("upload_spool.SaveObjectCacheRef write cache ref: %w", err)
	}
	return nil
}

func (m *UploadStateModel) LoadObjectCacheRef(ctx context.Context, objectID int64) (*xkv.DfsFileInfo, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	dir, err := m.cacheRefDir(objectID)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(filepath.Join(dir, strconv.FormatInt(objectID, 10)+".json"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrUploadStateNotFound
		}
		return nil, fmt.Errorf("upload_spool.LoadObjectCacheRef read cache ref: %w", err)
	}
	var ref objectCacheRef
	if err := json.Unmarshal(data, &ref); err != nil {
		return nil, fmt.Errorf("upload_spool.LoadObjectCacheRef unmarshal cache ref: %w", err)
	}
	return m.LoadUploadFileInfo(ctx, ref.Creator, ref.FileID)
}

func (m *UploadStateModel) ScanOnStart(ctx context.Context, now time.Time) error {
	if err := m.ScanOnStartWithoutCleanup(ctx, now); err != nil {
		return err
	}
	return m.CleanupExpiredUploadSessions(ctx, now)
}

func (m *UploadStateModel) ScanOnStartWithoutCleanup(ctx context.Context, _ time.Time) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	nodeRoot, err := safeJoin(m.rootDir, m.nodeID)
	if err != nil {
		return fmt.Errorf("upload_spool.ScanOnStart node root: %w", err)
	}
	if err := m.probeWritable(ctx, nodeRoot); err != nil {
		m.markUnwritable(fmt.Sprintf("startup writable probe failed: %v", err))
		return fmt.Errorf("upload_spool.ScanOnStart probe writable: %w", err)
	}
	m.markWritable()
	if err := m.cleanupStaleAtomicTemps(ctx); err != nil {
		return err
	}
	if err := m.cleanupUploadedSegmentBytes(ctx); err != nil {
		return err
	}
	return nil
}

func (m *UploadStateModel) CleanupExpiredUploadSessions(ctx context.Context, now time.Time) error {
	return m.CleanupExpiredUploadSessionsWithAbort(ctx, now, nil)
}

func (m *UploadStateModel) CleanupExpiredUploadSessionsWithAbort(ctx context.Context, now time.Time, abort func(ctx context.Context, objectKey, uploadID string) error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if m.partTTLSeconds <= 0 {
		return nil
	}
	ttl := time.Duration(m.partTTLSeconds) * time.Second
	return m.forEachSessionDir(ctx, func(sessionDir string) error {
		lastActivity, err := m.sessionLastActivity(sessionDir)
		if err != nil {
			return err
		}
		if !now.After(lastActivity.Add(ttl)) {
			return nil
		}
		if abort != nil {
			if err := m.abortSessionMultipartUploads(ctx, sessionDir, abort); err != nil {
				return err
			}
		}
		if err := os.RemoveAll(sessionDir); err != nil {
			return fmt.Errorf("upload_spool.CleanupExpiredUploadSessions remove session %s: %w", sessionDir, err)
		}
		if err := syncDir(filepath.Dir(sessionDir)); err != nil {
			return fmt.Errorf("upload_spool.CleanupExpiredUploadSessions sync parent %s: %w", filepath.Dir(sessionDir), err)
		}
		return nil
	})
}

func (m *UploadStateModel) abortSessionMultipartUploads(ctx context.Context, sessionDir string, abort func(ctx context.Context, objectKey, uploadID string) error) error {
	states, err := loadSegmentStatesFromDir(sessionDir)
	if err != nil {
		if errors.Is(err, ErrUploadStateNotFound) {
			return nil
		}
		return err
	}
	seen := make(map[string]struct{})
	for _, state := range states {
		if state.MultipartUploadID == "" {
			continue
		}
		objectKey := state.ObjectKey
		if objectKey == "" {
			objectKey = "objects/" + state.MultipartUploadID + ".dat"
		}
		seenKey := objectKey + "\x00" + state.MultipartUploadID
		if _, ok := seen[seenKey]; ok {
			continue
		}
		seen[seenKey] = struct{}{}
		if err := abort(ctx, objectKey, state.MultipartUploadID); err != nil {
			return fmt.Errorf("upload_spool.CleanupExpiredUploadSessions abort multipart upload_id=%s: %w", state.MultipartUploadID, err)
		}
	}
	return nil
}

func (m *UploadStateModel) ReconcileUploadingMultipartSegments(ctx context.Context, list func(ctx context.Context, objectKey, uploadID string) ([]MultipartPart, error)) error {
	if list == nil {
		return nil
	}
	listed := make(map[string][]MultipartPart)
	return m.forEachSessionDir(ctx, func(sessionDir string) error {
		states, err := loadSegmentStatesFromDir(sessionDir)
		if err != nil {
			if errors.Is(err, ErrUploadStateNotFound) {
				return nil
			}
			return err
		}
		for _, state := range states {
			if err := ctx.Err(); err != nil {
				return err
			}
			if state.Status != SegmentStatusUploading || state.MultipartUploadID == "" || state.ObjectKey == "" || state.MultipartPartNumber <= 0 {
				continue
			}
			cacheKey := state.ObjectKey + "\x00" + state.MultipartUploadID
			parts, ok := listed[cacheKey]
			if !ok {
				parts, err = list(ctx, state.ObjectKey, state.MultipartUploadID)
				if err != nil {
					return fmt.Errorf("upload_spool.ReconcileUploadingMultipartSegments list upload_id=%s: %w", state.MultipartUploadID, err)
				}
				listed[cacheKey] = parts
			}
			part, ok := matchingMultipartPart(parts, state)
			if !ok {
				continue
			}
			info, err := loadUploadFileInfoFromDir(sessionDir)
			if err != nil {
				if errors.Is(err, ErrUploadStateNotFound) {
					continue
				}
				return err
			}
			replayed, err := m.ReplaySegment(ctx, info, state.SegmentNo, io.Discard)
			if err != nil {
				if errors.Is(err, ErrUploadStateNotFound) {
					continue
				}
				return fmt.Errorf("upload_spool.ReconcileUploadingMultipartSegments replay segment %d: %w", state.SegmentNo, err)
			}
			if state.Checksum == "" || replayed.Checksum != state.Checksum {
				continue
			}
			state.Status = SegmentStatusUploaded
			state.ETag = part.ETag
			state.Size = part.Size
			statePath := filepath.Join(sessionDir, segmentStateFileName(state.SegmentNo))
			data, err := json.Marshal(state)
			if err != nil {
				return fmt.Errorf("upload_spool.ReconcileUploadingMultipartSegments marshal state: %w", err)
			}
			if err := writeFileAtomically(statePath, data, 0o644); err != nil {
				return fmt.Errorf("upload_spool.ReconcileUploadingMultipartSegments write state: %w", err)
			}
			segmentPath := filepath.Join(sessionDir, segmentFileName(state.SegmentNo))
			if err := os.Remove(segmentPath); err != nil && !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("upload_spool.ReconcileUploadingMultipartSegments remove segment %s: %w", segmentPath, err)
			}
			if err := syncDir(sessionDir); err != nil {
				return fmt.Errorf("upload_spool.ReconcileUploadingMultipartSegments sync session %s: %w", sessionDir, err)
			}
		}
		return nil
	})
}

func matchingMultipartPart(parts []MultipartPart, state SegmentState) (MultipartPart, bool) {
	for _, part := range parts {
		if part.PartNumber != state.MultipartPartNumber {
			continue
		}
		if state.Size > 0 && part.Size != state.Size {
			continue
		}
		if state.ETag != "" && part.ETag != state.ETag {
			continue
		}
		return part, true
	}
	return MultipartPart{}, false
}

func (m *UploadStateModel) MarkDraining(reason string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.draining = true
	m.writable = false
	m.drainReason = strings.TrimSpace(reason)
}

func (m *UploadStateModel) IsWritable() bool {
	if m == nil {
		return false
	}
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.writable && !m.draining
}

func (m *UploadStateModel) checkWritable() error {
	if m.IsWritable() {
		return nil
	}
	m.mu.RLock()
	reason := m.drainReason
	m.mu.RUnlock()
	if reason == "" {
		return ErrUploadSpoolNotWritable
	}
	return fmt.Errorf("%w: %s", ErrUploadSpoolNotWritable, reason)
}

func (m *UploadStateModel) markWritable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.draining {
		return
	}
	m.writable = true
	m.drainReason = ""
}

func (m *UploadStateModel) markUnwritable(reason string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.writable = false
	if m.draining {
		return
	}
	if strings.TrimSpace(reason) != "" {
		m.drainReason = strings.TrimSpace(reason)
	}
}

func (m *UploadStateModel) cleanupStaleAtomicTemps(ctx context.Context) error {
	return m.forEachSessionDir(ctx, func(sessionDir string) error {
		return filepath.WalkDir(sessionDir, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return fmt.Errorf("upload_spool.ScanOnStart walk temp %s: %w", path, err)
			}
			if err := ctx.Err(); err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			name := d.Name()
			if !strings.HasPrefix(name, ".") || !strings.Contains(name, ".tmp-") {
				return nil
			}
			if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("upload_spool.ScanOnStart remove temp %s: %w", path, err)
			}
			if err := syncDir(filepath.Dir(path)); err != nil {
				return fmt.Errorf("upload_spool.ScanOnStart sync temp parent %s: %w", filepath.Dir(path), err)
			}
			return nil
		})
	})
}

func (m *UploadStateModel) cleanupUploadedSegmentBytes(ctx context.Context) error {
	return m.forEachSessionDir(ctx, func(sessionDir string) error {
		states, err := loadSegmentStatesFromDir(sessionDir)
		if err != nil {
			if errors.Is(err, ErrUploadStateNotFound) {
				return nil
			}
			return err
		}
		for _, state := range states {
			if err := ctx.Err(); err != nil {
				return err
			}
			if state.Status != SegmentStatusUploaded {
				continue
			}
			path := filepath.Join(sessionDir, segmentFileName(state.SegmentNo))
			if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("upload_spool.ScanOnStart remove uploaded segment %s: %w", path, err)
			}
			if err := syncDir(sessionDir); err != nil {
				return fmt.Errorf("upload_spool.ScanOnStart sync uploaded segment dir %s: %w", sessionDir, err)
			}
		}
		return nil
	})
}

func (m *UploadStateModel) forEachSessionDir(ctx context.Context, fn func(sessionDir string) error) error {
	nodeRoot, err := safeJoin(m.rootDir, m.nodeID)
	if err != nil {
		return fmt.Errorf("upload_spool.forEachSessionDir node root: %w", err)
	}
	shards, err := os.ReadDir(nodeRoot)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("upload_spool.forEachSessionDir read node root: %w", err)
	}
	for _, shard := range shards {
		if err := ctx.Err(); err != nil {
			return err
		}
		if !shard.IsDir() || shard.Name() == "cache_refs" {
			continue
		}
		shardDir, err := safeJoin(nodeRoot, shard.Name())
		if err != nil {
			return fmt.Errorf("upload_spool.forEachSessionDir shard path: %w", err)
		}
		sessions, err := os.ReadDir(shardDir)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return fmt.Errorf("upload_spool.forEachSessionDir read shard %s: %w", shardDir, err)
		}
		for _, session := range sessions {
			if err := ctx.Err(); err != nil {
				return err
			}
			if !session.IsDir() {
				continue
			}
			sessionDir, err := safeJoin(shardDir, session.Name())
			if err != nil {
				return fmt.Errorf("upload_spool.forEachSessionDir session path: %w", err)
			}
			if err := fn(sessionDir); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *UploadStateModel) partsPerSegment(info *xkv.DfsFileInfo) (int, error) {
	if info == nil {
		return 0, ErrUploadStateNotFound
	}
	partSize := info.FilePartSize
	if partSize <= 0 {
		partSize = info.FirstFilePartSize
	}
	if partSize <= 0 {
		partSize = info.LastFilePartSize
	}
	if partSize <= 0 {
		return 0, ErrUploadStateNotFound
	}
	parts := int(m.SegmentSize() / int64(partSize))
	if parts <= 0 {
		parts = 1
	}
	return parts, nil
}

func (m *UploadStateModel) segmentStatePath(info *xkv.DfsFileInfo, segmentNo int64) (string, error) {
	if info == nil {
		return "", ErrUploadStateNotFound
	}
	dir, err := m.sessionDir(info.Creator, info.FileID)
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, segmentStateFileName(segmentNo)), nil
}

func loadSegmentStatesFromDir(dir string) ([]SegmentState, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrUploadStateNotFound
		}
		return nil, fmt.Errorf("upload_spool.loadSegmentStatesFromDir read dir: %w", err)
	}
	var states []SegmentState
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasPrefix(entry.Name(), "segment-") || !strings.HasSuffix(entry.Name(), ".state.json") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			return nil, fmt.Errorf("upload_spool.loadSegmentStatesFromDir read %s: %w", entry.Name(), err)
		}
		var state SegmentState
		if err := json.Unmarshal(data, &state); err != nil {
			return nil, fmt.Errorf("upload_spool.loadSegmentStatesFromDir unmarshal %s: %w", entry.Name(), err)
		}
		states = append(states, state)
	}
	sort.Slice(states, func(i, j int) bool { return states[i].SegmentNo < states[j].SegmentNo })
	return states, nil
}

func loadUploadFileInfoFromDir(dir string) (*xkv.DfsFileInfo, error) {
	data, err := os.ReadFile(filepath.Join(dir, metadataFileName))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ErrUploadStateNotFound
		}
		return nil, fmt.Errorf("upload_spool.loadUploadFileInfoFromDir read metadata: %w", err)
	}
	var info xkv.DfsFileInfo
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, fmt.Errorf("upload_spool.loadUploadFileInfoFromDir unmarshal metadata: %w", err)
	}
	return &info, nil
}

func copyToWriters(dst io.Writer, checksum io.Writer, src io.Reader, firstBytes *[]byte) (int64, error) {
	buf := make([]byte, 32*1024)
	var total int64
	for {
		n, readErr := src.Read(buf)
		if n > 0 {
			chunk := buf[:n]
			if len(*firstBytes) < 512 {
				need := 512 - len(*firstBytes)
				if need > n {
					need = n
				}
				*firstBytes = append(*firstBytes, chunk[:need]...)
			}
			if _, err := dst.Write(chunk); err != nil {
				return total, err
			}
			if _, err := checksum.Write(chunk); err != nil {
				return total, err
			}
			total += int64(n)
		}
		if readErr == io.EOF {
			return total, nil
		}
		if readErr != nil {
			return total, readErr
		}
	}
}

func (m *UploadStateModel) sessionLastActivity(sessionDir string) (time.Time, error) {
	metadataPath := filepath.Join(sessionDir, metadataFileName)
	if data, err := os.ReadFile(metadataPath); err == nil {
		var info xkv.DfsFileInfo
		if err := json.Unmarshal(data, &info); err == nil && info.Mtime > 0 {
			return time.Unix(info.Mtime, 0), nil
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return time.Time{}, fmt.Errorf("upload_spool.sessionLastActivity read metadata: %w", err)
	}
	return newestModTime(sessionDir)
}

func newestModTime(root string) (time.Time, error) {
	var newest time.Time
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("upload_spool.newestModTime walk %s: %w", path, err)
		}
		info, err := d.Info()
		if err != nil {
			return fmt.Errorf("upload_spool.newestModTime stat %s: %w", path, err)
		}
		if info.ModTime().After(newest) {
			newest = info.ModTime()
		}
		return nil
	})
	if err != nil {
		return time.Time{}, err
	}
	return newest, nil
}

func (m *UploadStateModel) sessionDir(creator, fileID int64) (string, error) {
	shard := int64(0)
	if m.localShardCount > 0 {
		shard = int64(uploadShard(creator, fileID) % uint64(m.localShardCount))
	}
	return safeJoin(m.rootDir, m.nodeID, strconv.FormatInt(shard, 10), fmt.Sprintf("%d_%d", creator, fileID))
}

func (m *UploadStateModel) cacheRefDir(objectID int64) (string, error) {
	shard := uint64(objectID)
	if m.localShardCount > 0 {
		shard = shard % uint64(m.localShardCount)
	}
	return safeJoin(m.rootDir, m.nodeID, "cache_refs", strconv.FormatUint(shard, 10))
}

func loadOrCreateNodeID(rootDir, nodeIDFile string) (string, error) {
	if nodeIDFile == "" {
		nodeIDFile = "node_id"
	}
	if !filepath.IsAbs(nodeIDFile) {
		nodeIDFile = filepath.Join(rootDir, nodeIDFile)
	}
	data, err := os.ReadFile(nodeIDFile)
	if err == nil {
		nodeID := parseNodeIDFile(data)
		if err := validateNodeID(nodeID); err != nil {
			return "", fmt.Errorf("upload_spool.NewUploadStateModel invalid node id file %s: %w", nodeIDFile, err)
		}
		return nodeID, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("upload_spool.NewUploadStateModel read node id: %w", err)
	}
	nodeID, err := randomNodeID()
	if err != nil {
		return "", err
	}
	if err := mkdirAllDurable(filepath.Dir(nodeIDFile), 0o755); err != nil {
		return "", fmt.Errorf("upload_spool.NewUploadStateModel mkdir node id dir: %w", err)
	}
	if err := writeFileAtomically(nodeIDFile, []byte(nodeID+"\n"), 0o644); err != nil {
		return "", fmt.Errorf("upload_spool.NewUploadStateModel write node id: %w", err)
	}
	return nodeID, nil
}

func parseNodeIDFile(data []byte) string {
	nodeID := strings.TrimSuffix(string(data), "\n")
	return strings.TrimSuffix(nodeID, "\r")
}

func validateNodeID(nodeID string) error {
	if nodeID == "" || strings.TrimSpace(nodeID) != nodeID {
		return errors.New("node id is empty or has surrounding whitespace")
	}
	if nodeID == "." || nodeID == ".." {
		return fmt.Errorf("node id %q is not allowed", nodeID)
	}
	if strings.ContainsAny(nodeID, `/\`) || !nodeIDPattern.MatchString(nodeID) {
		return fmt.Errorf("node id %q contains unsafe characters", nodeID)
	}
	return nil
}

func randomNodeID() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", fmt.Errorf("upload_spool.NewUploadStateModel random node id: %w", err)
	}
	return hex.EncodeToString(b[:]), nil
}

func writeFileAtomically(path string, data []byte, perm os.FileMode) error {
	if err := mkdirAllDurable(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), "."+filepath.Base(path)+".tmp-*")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	closed := false
	defer func() {
		if !closed {
			_ = tmp.Close()
		}
		_ = os.Remove(tmpName)
	}()
	if _, err := tmp.Write(data); err != nil {
		return err
	}
	if err := tmp.Chmod(perm); err != nil {
		return err
	}
	if err := tmp.Sync(); err != nil {
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	closed = true
	if err := os.Rename(tmpName, path); err != nil {
		return err
	}
	return syncDir(filepath.Dir(path))
}

func mkdirAllDurable(path string, perm os.FileMode) error {
	path = filepath.Clean(path)
	var syncParents []string
	for dir := path; ; dir = filepath.Dir(dir) {
		if _, err := os.Stat(dir); err == nil {
			break
		} else if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		parent := filepath.Dir(dir)
		if parent != dir {
			syncParents = append(syncParents, parent)
		}
		if parent == dir {
			break
		}
	}
	if err := os.MkdirAll(path, perm); err != nil {
		return err
	}
	for i := len(syncParents) - 1; i >= 0; i-- {
		if err := syncDir(syncParents[i]); err != nil {
			return err
		}
	}
	return nil
}

func syncDir(dir string) error {
	f, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer f.Close()
	return f.Sync()
}

func defaultProbeWritable(ctx context.Context, nodeRoot string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(nodeRoot, ".writable-probe-*")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	closed := false
	defer func() {
		if !closed {
			_ = tmp.Close()
		}
		_ = os.Remove(tmpName)
	}()
	if _, err := tmp.Write([]byte("ok")); err != nil {
		return err
	}
	if err := tmp.Sync(); err != nil {
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	closed = true
	if err := os.Remove(tmpName); err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	return syncDir(nodeRoot)
}

func safeJoin(root string, elems ...string) (string, error) {
	root = filepath.Clean(root)
	parts := append([]string{root}, elems...)
	path := filepath.Clean(filepath.Join(parts...))
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return "", err
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) || filepath.IsAbs(rel) {
		return "", fmt.Errorf("path %q escapes root %q", path, root)
	}
	return path, nil
}

func uploadShard(creator, fileID int64) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(strconv.FormatInt(creator, 10)))
	_, _ = h.Write([]byte{':'})
	_, _ = h.Write([]byte(strconv.FormatInt(fileID, 10)))
	return h.Sum64()
}

func partFileName(partIndex int32) string {
	return fmt.Sprintf("part_%09d.bin", partIndex)
}

func segmentFileName(segmentNo int64) string {
	return fmt.Sprintf("segment-%06d.tmp", segmentNo)
}

func segmentStateFileName(segmentNo int64) string {
	return fmt.Sprintf("segment-%06d.state.json", segmentNo)
}
