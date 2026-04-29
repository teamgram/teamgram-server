package repository

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
)

func TestWrapMediaInvalidUploadedFile(t *testing.T) {
	err := wrapMediaInvalidUploadedFile("dfs upload", dfs.ErrDfsFileNotFound)
	if !errors.Is(err, media.ErrMediaInvalidUploadedFile) {
		t.Fatalf("error = %v, want ErrMediaInvalidUploadedFile", err)
	}
	if !errors.Is(err, dfs.ErrDfsFileNotFound) {
		t.Fatalf("error = %v, want dfs cause", err)
	}
}

func TestWrapMediaDownstream(t *testing.T) {
	err := wrapMediaDownstream("dfs upload", dfs.ErrDfsStorage)
	if !errors.Is(err, media.ErrMediaDownstream) {
		t.Fatalf("error = %v, want ErrMediaDownstream", err)
	}
	if !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("error = %v, want dfs cause", err)
	}
}
