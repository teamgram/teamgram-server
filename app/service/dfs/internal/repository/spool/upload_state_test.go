package spool

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

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

func newTestUploadStateModel(t *testing.T) *UploadStateModel {
	t.Helper()
	model, err := NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	return model
}
