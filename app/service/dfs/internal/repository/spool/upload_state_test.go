package spool

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/xkv"
)

func TestUploadStateModelReadsSavedPartsInOrder(t *testing.T) {
	ctx := context.Background()
	model := newTestUploadStateModel(t)

	if err := model.SaveUploadPart(ctx, 1001, 2002, 2, []byte("cc")); err != nil {
		t.Fatalf("SaveUploadPart(2) error = %v", err)
	}
	if err := model.SaveUploadPart(ctx, 1001, 2002, 0, []byte("aa")); err != nil {
		t.Fatalf("SaveUploadPart(0) error = %v", err)
	}
	if err := model.SaveUploadPart(ctx, 1001, 2002, 1, []byte("bb")); err != nil {
		t.Fatalf("SaveUploadPart(1) error = %v", err)
	}
	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            2002,
		FileTotalParts:    3,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		Mtime:             1_700_000_000,
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}

	info, err := model.LoadUploadFileInfo(ctx, 1001, 2002)
	if err != nil {
		t.Fatalf("LoadUploadFileInfo() error = %v", err)
	}
	reader, err := model.OpenUploadFileReader(ctx, info)
	if err != nil {
		t.Fatalf("OpenUploadFileReader() error = %v", err)
	}
	got, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}
	if string(got) != "aabbcc" {
		t.Fatalf("OpenUploadFileReader() bytes = %q, want %q", got, "aabbcc")
	}
}

func TestUploadStateModelMissingMetadataReturnsNotFound(t *testing.T) {
	_, err := newTestUploadStateModel(t).LoadUploadFileInfo(context.Background(), 1001, 2002)
	if !errors.Is(err, ErrUploadStateNotFound) {
		t.Fatalf("LoadUploadFileInfo() error = %v, want ErrUploadStateNotFound", err)
	}
}

func TestUploadStateModelMissingPartReturnsNotFound(t *testing.T) {
	ctx := context.Background()
	model := newTestUploadStateModel(t)

	if err := model.SaveUploadPart(ctx, 1001, 2002, 0, []byte("aa")); err != nil {
		t.Fatalf("SaveUploadPart() error = %v", err)
	}
	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            2002,
		FileTotalParts:    2,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}

	_, err := model.OpenUploadFileReader(ctx, &xkv.DfsFileInfo{
		Creator:        1001,
		FileID:         2002,
		FileTotalParts: 2,
	})
	if !errors.Is(err, ErrUploadStateNotFound) {
		t.Fatalf("OpenUploadFileReader() error = %v, want ErrUploadStateNotFound", err)
	}
}

func TestUploadStateModelCacheRefRoundTrip(t *testing.T) {
	ctx := context.Background()
	model := newTestUploadStateModel(t)

	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            2002,
		FileTotalParts:    1,
		FirstFilePartSize: 7,
		Mtime:             1_700_000_000,
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	if err := model.SaveObjectCacheRef(ctx, 3003, 1001, 2002); err != nil {
		t.Fatalf("SaveObjectCacheRef() error = %v", err)
	}

	info, err := model.LoadObjectCacheRef(ctx, 3003)
	if err != nil {
		t.Fatalf("LoadObjectCacheRef() error = %v", err)
	}
	if info.Creator != 1001 || info.FileID != 2002 || info.FirstFilePartSize != 7 {
		t.Fatalf("LoadObjectCacheRef() = %+v, want creator/file/part size", info)
	}
}

func TestUploadStateModelNodeIDFileIsStableAcrossInstances(t *testing.T) {
	root := t.TempDir()
	nodeIDFile := filepath.Join(root, "node_id")
	conf := config.UploadSpoolConf{
		RootDir:         root,
		NodeIDFile:      nodeIDFile,
		LocalShardCount: 4,
	}

	first, err := NewUploadStateModel(conf)
	if err != nil {
		t.Fatalf("NewUploadStateModel(first) error = %v", err)
	}
	second, err := NewUploadStateModel(conf)
	if err != nil {
		t.Fatalf("NewUploadStateModel(second) error = %v", err)
	}
	if first.NodeID() == "" {
		t.Fatal("NodeID() is empty")
	}
	if first.NodeID() != second.NodeID() {
		t.Fatalf("NodeID() changed across instances: %q != %q", first.NodeID(), second.NodeID())
	}
	onDisk, err := os.ReadFile(nodeIDFile)
	if err != nil {
		t.Fatalf("ReadFile(node_id) error = %v", err)
	}
	if string(onDisk) != first.NodeID()+"\n" {
		t.Fatalf("node_id file = %q, want %q", string(onDisk), first.NodeID()+"\n")
	}
}

func TestUploadStateModelRejectsUnsafeNodeIDFile(t *testing.T) {
	tests := []struct {
		name   string
		nodeID string
	}{
		{name: "parent escape", nodeID: "../outside"},
		{name: "slash", nodeID: "a/b"},
		{name: "backslash", nodeID: `a\b`},
		{name: "empty", nodeID: ""},
		{name: "whitespace", nodeID: "   \n\t"},
		{name: "surrounding whitespace", nodeID: " node-1\n"},
		{name: "dotdot", nodeID: ".."},
		{name: "dot", nodeID: "."},
		{name: "space", nodeID: "node id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := t.TempDir()
			nodeIDFile := filepath.Join(root, "node_id")
			if err := os.WriteFile(nodeIDFile, []byte(tt.nodeID), 0o644); err != nil {
				t.Fatalf("WriteFile(node_id) error = %v", err)
			}
			_, err := NewUploadStateModel(config.UploadSpoolConf{
				RootDir:    root,
				NodeIDFile: nodeIDFile,
			})
			if err == nil {
				t.Fatal("NewUploadStateModel() error = nil, want unsafe node id rejection")
			}
		})
	}
}

func TestUploadStateModelPathsRemainUnderRoot(t *testing.T) {
	root := t.TempDir()
	nodeIDFile := filepath.Join(root, "node_id")
	if err := os.WriteFile(nodeIDFile, []byte("node-1\n"), 0o644); err != nil {
		t.Fatalf("WriteFile(node_id) error = %v", err)
	}
	model, err := NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         root,
		NodeIDFile:      nodeIDFile,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}

	for name, pathFn := range map[string]func() (string, error){
		"session":   func() (string, error) { return model.sessionDir(1001, 2002) },
		"cache_ref": func() (string, error) { return model.cacheRefDir(3003) },
	} {
		t.Run(name, func(t *testing.T) {
			got, err := pathFn()
			if err != nil {
				t.Fatalf("pathFn() error = %v", err)
			}
			rel, err := filepath.Rel(root, got)
			if err != nil {
				t.Fatalf("Rel() error = %v", err)
			}
			if rel == "." || rel == ".." || rel == "" || len(rel) >= 2 && rel[:2] == ".." {
				t.Fatalf("path %q escaped root %q", got, root)
			}
		})
	}
}

func TestUploadStateModelScanOnStartKeepsFreshSession(t *testing.T) {
	ctx := context.Background()
	model := newTestUploadStateModelWithConf(t, config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		LocalShardCount: 4,
		PartTTLSeconds:  3600,
	})
	now := time.Unix(1_700_010_000, 0)
	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{
		Creator:        1001,
		FileID:         2002,
		FileTotalParts: 1,
		Mtime:          now.Add(-10 * time.Minute).Unix(),
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}

	if err := model.ScanOnStart(ctx, now); err != nil {
		t.Fatalf("ScanOnStart() error = %v", err)
	}
	if _, err := model.LoadUploadFileInfo(ctx, 1001, 2002); err != nil {
		t.Fatalf("LoadUploadFileInfo() after scan error = %v", err)
	}
	if !model.IsWritable() {
		t.Fatal("IsWritable() = false, want true")
	}
}

func TestUploadStateModelScanOnStartRemovesExpiredSession(t *testing.T) {
	ctx := context.Background()
	model := newTestUploadStateModelWithConf(t, config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		LocalShardCount: 4,
		PartTTLSeconds:  3600,
	})
	now := time.Unix(1_700_010_000, 0)
	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{
		Creator:        1001,
		FileID:         2002,
		FileTotalParts: 1,
		Mtime:          now.Add(-2 * time.Hour).Unix(),
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}

	if err := model.ScanOnStart(ctx, now); err != nil {
		t.Fatalf("ScanOnStart() error = %v", err)
	}
	if _, err := model.LoadUploadFileInfo(ctx, 1001, 2002); !errors.Is(err, ErrUploadStateNotFound) {
		t.Fatalf("LoadUploadFileInfo() after scan error = %v, want ErrUploadStateNotFound", err)
	}
}

func TestUploadStateModelScanOnStartUsesFileMtimeForMissingAndCorruptMetadata(t *testing.T) {
	ctx := context.Background()
	model := newTestUploadStateModelWithConf(t, config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		LocalShardCount: 4,
		PartTTLSeconds:  3600,
	})
	now := time.Unix(1_700_010_000, 0)

	missingFreshDir, err := model.sessionDir(1001, 2002)
	if err != nil {
		t.Fatalf("sessionDir(missing fresh) error = %v", err)
	}
	if err := mkdirAllDurable(missingFreshDir, 0o755); err != nil {
		t.Fatalf("mkdir missing fresh session error = %v", err)
	}
	missingFreshPart := filepath.Join(missingFreshDir, partFileName(0))
	if err := os.WriteFile(missingFreshPart, []byte("fresh"), 0o644); err != nil {
		t.Fatalf("WriteFile(missing fresh part) error = %v", err)
	}
	freshTime := now.Add(-10 * time.Minute)
	if err := os.Chtimes(missingFreshPart, freshTime, freshTime); err != nil {
		t.Fatalf("Chtimes(missing fresh part) error = %v", err)
	}
	if err := os.Chtimes(missingFreshDir, freshTime, freshTime); err != nil {
		t.Fatalf("Chtimes(missing fresh dir) error = %v", err)
	}

	missingExpiredDir, err := model.sessionDir(5005, 6006)
	if err != nil {
		t.Fatalf("sessionDir(missing expired) error = %v", err)
	}
	if err := mkdirAllDurable(missingExpiredDir, 0o755); err != nil {
		t.Fatalf("mkdir missing expired session error = %v", err)
	}
	missingExpiredPart := filepath.Join(missingExpiredDir, partFileName(0))
	if err := os.WriteFile(missingExpiredPart, []byte("expired"), 0o644); err != nil {
		t.Fatalf("WriteFile(missing expired part) error = %v", err)
	}
	expiredTime := now.Add(-2 * time.Hour)
	if err := os.Chtimes(missingExpiredPart, expiredTime, expiredTime); err != nil {
		t.Fatalf("Chtimes(missing expired part) error = %v", err)
	}
	if err := os.Chtimes(missingExpiredDir, expiredTime, expiredTime); err != nil {
		t.Fatalf("Chtimes(missing expired dir) error = %v", err)
	}

	corruptFreshDir, err := model.sessionDir(7007, 8008)
	if err != nil {
		t.Fatalf("sessionDir(corrupt fresh) error = %v", err)
	}
	if err := mkdirAllDurable(corruptFreshDir, 0o755); err != nil {
		t.Fatalf("mkdir corrupt fresh session error = %v", err)
	}
	corruptFreshMetadata := filepath.Join(corruptFreshDir, metadataFileName)
	if err := os.WriteFile(corruptFreshMetadata, []byte("{not-json"), 0o644); err != nil {
		t.Fatalf("WriteFile(corrupt fresh metadata) error = %v", err)
	}
	if err := os.Chtimes(corruptFreshMetadata, freshTime, freshTime); err != nil {
		t.Fatalf("Chtimes(corrupt fresh metadata) error = %v", err)
	}
	if err := os.Chtimes(corruptFreshDir, freshTime, freshTime); err != nil {
		t.Fatalf("Chtimes(corrupt fresh dir) error = %v", err)
	}

	corruptExpiredDir, err := model.sessionDir(3003, 4004)
	if err != nil {
		t.Fatalf("sessionDir(corrupt expired) error = %v", err)
	}
	if err := mkdirAllDurable(corruptExpiredDir, 0o755); err != nil {
		t.Fatalf("mkdir corrupt expired session error = %v", err)
	}
	corruptMetadata := filepath.Join(corruptExpiredDir, metadataFileName)
	if err := os.WriteFile(corruptMetadata, []byte("{not-json"), 0o644); err != nil {
		t.Fatalf("WriteFile(corrupt metadata) error = %v", err)
	}
	if err := os.Chtimes(corruptMetadata, expiredTime, expiredTime); err != nil {
		t.Fatalf("Chtimes(corrupt metadata) error = %v", err)
	}
	if err := os.Chtimes(corruptExpiredDir, expiredTime, expiredTime); err != nil {
		t.Fatalf("Chtimes(corrupt expired dir) error = %v", err)
	}

	if err := model.ScanOnStart(ctx, now); err != nil {
		t.Fatalf("ScanOnStart() error = %v", err)
	}
	if _, err := os.Stat(missingFreshDir); err != nil {
		t.Fatalf("missing metadata fresh session stat error = %v", err)
	}
	if _, err := os.Stat(missingExpiredDir); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("missing metadata expired session stat error = %v, want not exist", err)
	}
	if _, err := os.Stat(corruptFreshDir); err != nil {
		t.Fatalf("corrupt metadata fresh session stat error = %v", err)
	}
	if _, err := os.Stat(corruptExpiredDir); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("corrupt metadata expired session stat error = %v, want not exist", err)
	}
}

func TestUploadStateModelScanOnStartRemovesStaleAtomicTempFiles(t *testing.T) {
	ctx := context.Background()
	model := newTestUploadStateModelWithConf(t, config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		LocalShardCount: 4,
		PartTTLSeconds:  0,
	})
	dir, err := model.sessionDir(1001, 2002)
	if err != nil {
		t.Fatalf("sessionDir() error = %v", err)
	}
	if err := mkdirAllDurable(dir, 0o755); err != nil {
		t.Fatalf("mkdir session error = %v", err)
	}
	tmpPath := filepath.Join(dir, "."+partFileName(0)+".tmp-leftover")
	if err := os.WriteFile(tmpPath, []byte("partial"), 0o644); err != nil {
		t.Fatalf("WriteFile(temp) error = %v", err)
	}

	if err := model.ScanOnStart(ctx, time.Unix(1_700_010_000, 0)); err != nil {
		t.Fatalf("ScanOnStart() error = %v", err)
	}
	if _, err := os.Stat(tmpPath); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("temp stat error = %v, want not exist", err)
	}
}

func TestUploadStateModelScanOnStartDoesNotTreatCacheRefsAsSessions(t *testing.T) {
	ctx := context.Background()
	model := newTestUploadStateModelWithConf(t, config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		LocalShardCount: 4,
		PartTTLSeconds:  1,
	})
	if err := model.SaveObjectCacheRef(ctx, 3003, 1001, 2002); err != nil {
		t.Fatalf("SaveObjectCacheRef() error = %v", err)
	}
	cacheDir, err := model.cacheRefDir(3003)
	if err != nil {
		t.Fatalf("cacheRefDir() error = %v", err)
	}

	if err := model.ScanOnStart(ctx, time.Unix(1_700_010_000, 0)); err != nil {
		t.Fatalf("ScanOnStart() error = %v", err)
	}
	if _, err := os.Stat(cacheDir); err != nil {
		t.Fatalf("cache_refs stat error = %v", err)
	}
}

func TestUploadStateModelProbeFailureMarksUnwritable(t *testing.T) {
	model := newTestUploadStateModel(t)
	probeErr := errors.New("probe failed")
	model.probeWritable = func(context.Context, string) error {
		return probeErr
	}

	err := model.ScanOnStart(context.Background(), time.Unix(1_700_010_000, 0))
	if !errors.Is(err, probeErr) {
		t.Fatalf("ScanOnStart() error = %v, want probe error", err)
	}
	if model.IsWritable() {
		t.Fatal("IsWritable() = true, want false")
	}
	if err := model.SaveUploadPart(context.Background(), 1001, 2002, 0, []byte("data")); err == nil {
		t.Fatal("SaveUploadPart() error = nil, want unwritable rejection")
	}
}

func TestUploadStateModelWriteFailureMarksUnwritable(t *testing.T) {
	ctx := context.Background()
	model := newTestUploadStateModel(t)
	if err := model.SaveUploadPart(ctx, 3003, 4004, 0, []byte("ok")); err != nil {
		t.Fatalf("SaveUploadPart(preexisting) error = %v", err)
	}
	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{
		Creator:           3003,
		FileID:            4004,
		FileTotalParts:    1,
		FirstFilePartSize: 2,
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo(preexisting) error = %v", err)
	}
	if err := model.SaveObjectCacheRef(ctx, 5005, 3003, 4004); err != nil {
		t.Fatalf("SaveObjectCacheRef(preexisting) error = %v", err)
	}
	sessionDir, err := model.sessionDir(1001, 2002)
	if err != nil {
		t.Fatalf("sessionDir() error = %v", err)
	}
	if err := mkdirAllDurable(filepath.Dir(sessionDir), 0o755); err != nil {
		t.Fatalf("mkdir session parent error = %v", err)
	}
	if err := os.WriteFile(sessionDir, []byte("path conflict"), 0o644); err != nil {
		t.Fatalf("WriteFile(session path conflict) error = %v", err)
	}

	err = model.SaveUploadPart(ctx, 1001, 2002, 0, []byte("data"))
	if err == nil {
		t.Fatal("SaveUploadPart() error = nil, want write failure")
	}
	if errors.Is(err, ErrUploadSpoolNotWritable) {
		t.Fatalf("SaveUploadPart() error = %v, should be original write failure", err)
	}
	if model.IsWritable() {
		t.Fatal("IsWritable() = true after write failure, want false")
	}
	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{Creator: 1001, FileID: 2002}); !errors.Is(err, ErrUploadSpoolNotWritable) {
		t.Fatalf("SaveUploadFileInfo() after write failure error = %v, want ErrUploadSpoolNotWritable", err)
	}

	info, err := model.LoadUploadFileInfo(ctx, 3003, 4004)
	if err != nil {
		t.Fatalf("LoadUploadFileInfo() after write failure error = %v", err)
	}
	reader, err := model.OpenUploadFileReader(ctx, info)
	if err != nil {
		t.Fatalf("OpenUploadFileReader() after write failure error = %v", err)
	}
	got, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll() after write failure error = %v", err)
	}
	if string(got) != "ok" {
		t.Fatalf("OpenUploadFileReader() after write failure bytes = %q, want %q", got, "ok")
	}
	if _, err := model.LoadObjectCacheRef(ctx, 5005); err != nil {
		t.Fatalf("LoadObjectCacheRef() after write failure error = %v", err)
	}
}

func TestUploadStateModelDrainReasonSurvivesUnwritableMark(t *testing.T) {
	model := newTestUploadStateModel(t)
	model.MarkDraining("planned drain")
	model.markUnwritable("fsync failed")

	err := model.SaveUploadPart(context.Background(), 1001, 2002, 0, []byte("data"))
	if !errors.Is(err, ErrUploadSpoolNotWritable) {
		t.Fatalf("SaveUploadPart() error = %v, want ErrUploadSpoolNotWritable", err)
	}
	if !strings.Contains(err.Error(), "planned drain") {
		t.Fatalf("SaveUploadPart() error = %v, want planned drain reason", err)
	}
	if strings.Contains(err.Error(), "fsync failed") {
		t.Fatalf("SaveUploadPart() error = %v, should keep drain reason", err)
	}
}

func newTestUploadStateModel(t *testing.T) *UploadStateModel {
	t.Helper()
	return newTestUploadStateModelWithConf(t, config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		LocalShardCount: 4,
	})
}

func newTestUploadStateModelWithConf(t *testing.T, conf config.UploadSpoolConf) *UploadStateModel {
	t.Helper()
	model, err := NewUploadStateModel(conf)
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	return model
}
