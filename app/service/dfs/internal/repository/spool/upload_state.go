package spool

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/xkv"
)

const (
	defaultLocalShardCount = 256
	metadataFileName       = "metadata.json"
)

var ErrUploadStateNotFound = errors.New("dfs upload spool state not found")
var ErrUploadSpoolNotWritable = errors.New("dfs upload spool is not writable")

var nodeIDPattern = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)

type UploadStateModel struct {
	rootDir         string
	nodeID          string
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
	return m.CleanupExpiredUploadSessions(ctx, now)
}

func (m *UploadStateModel) CleanupExpiredUploadSessions(ctx context.Context, now time.Time) error {
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
		if err := os.RemoveAll(sessionDir); err != nil {
			return fmt.Errorf("upload_spool.CleanupExpiredUploadSessions remove session %s: %w", sessionDir, err)
		}
		if err := syncDir(filepath.Dir(sessionDir)); err != nil {
			return fmt.Errorf("upload_spool.CleanupExpiredUploadSessions sync parent %s: %w", filepath.Dir(sessionDir), err)
		}
		return nil
	})
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
