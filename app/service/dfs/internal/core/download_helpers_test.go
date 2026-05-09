package core

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
)

func TestCompatibleDownloadBytesDoesNotSwallowStorage(t *testing.T) {
	core := &DfsCore{}
	got, err := core.compatibleDownloadBytes("download object", []byte("ignored"), dfs.ErrDfsStorage)
	if err == nil {
		t.Fatal("expected storage error")
	}
	if len(got) != 0 {
		t.Fatalf("expected no bytes on storage error, got %q", got)
	}
	if !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("expected ErrDfsStorage, got %v", err)
	}
}

func TestCompatibleDownloadBytesKeepsFileNotFoundCompatibility(t *testing.T) {
	core := &DfsCore{}
	got, err := core.compatibleDownloadBytes("download object", nil, dfs.ErrDfsFileNotFound)
	if err != nil {
		t.Fatalf("expected compatible empty success, got %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("expected empty bytes on file miss, got %q", got)
	}
}

func TestCompatibleBytesDoesNotSwallowStorage(t *testing.T) {
	got, err := compatibleBytes([]byte("ignored"), dfs.ErrDfsStorage)
	if err == nil {
		t.Fatal("expected storage error")
	}
	if len(got) != 0 {
		t.Fatalf("expected no bytes on storage error, got %q", got)
	}
	if !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("expected ErrDfsStorage, got %v", err)
	}
}

func TestCompatibleBytesKeepsFileNotFoundCompatibility(t *testing.T) {
	got, err := compatibleBytes(nil, dfs.ErrDfsFileNotFound)
	if err != nil {
		t.Fatalf("expected compatible empty success, got %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("expected empty bytes on file miss, got %q", got)
	}
}
