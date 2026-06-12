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

func TestWrapDfsUploadErrorMapping(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want error
	}{
		{name: "invalid argument", err: dfs.ErrDfsInvalidArgument, want: media.ErrMediaInvalidArgument},
		{name: "remote invalid argument", err: errors.New("remote or network error: biz error: dfs: invalid argument"), want: media.ErrMediaInvalidArgument},
		{name: "checksum invalid", err: dfs.ErrDfsChecksumInvalid, want: media.ErrMediaChecksumInvalid},
		{name: "invalid uploaded file", err: dfs.ErrDfsInvalidFilePart, want: media.ErrMediaInvalidUploadedFile},
		{name: "missing upload part", err: &dfs.MissingUploadPartError{Part: 3}, want: media.ErrMediaInvalidUploadedFile},
		{name: "downstream", err: dfs.ErrDfsStorage, want: media.ErrMediaDownstream},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := wrapDfsUploadError("dfs upload", tt.err)
			if !errors.Is(err, tt.want) {
				t.Fatalf("error = %v, want %v", err, tt.want)
			}
			if !errors.Is(err, tt.err) {
				t.Fatalf("error = %v, want dfs cause %v", err, tt.err)
			}
		})
	}
}
