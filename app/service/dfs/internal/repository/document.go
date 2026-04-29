package repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
)

func (r *Repository) NextDocumentID(ctx context.Context) (int64, error) {
	if r == nil || r.idgen == nil {
		return 0, dfs.WrapDfsDownstream("next document id", errors.New("idgen client unavailable"))
	}
	return r.idgen.NextDocumentID(ctx)
}

func (r *Repository) NextEncryptedFileID(ctx context.Context) (int64, error) {
	if r == nil || r.idgen == nil {
		return 0, dfs.WrapDfsDownstream("next encrypted file id", errors.New("idgen client unavailable"))
	}
	return r.idgen.NextEncryptedFileID(ctx)
}

func (r *Repository) SaveDocumentObject(ctx context.Context, documentID int64, data []byte) (int64, error) {
	if r == nil || r.objectStore == nil {
		return 0, dfs.WrapDfsStorage("put document file", errors.New("object store unavailable"))
	}
	if documentID == 0 {
		return 0, dfs.ErrDfsInvalidArgument
	}
	size, err := r.objectStore.PutDocumentReader(ctx, fmt.Sprintf("%d.dat", documentID), bytes.NewReader(data))
	if err != nil {
		return 0, dfs.WrapDfsStorage("put document file", err)
	}
	return size, nil
}

func (r *Repository) SaveEncryptedObject(ctx context.Context, fileID int64, data []byte) (int64, error) {
	if r == nil || r.objectStore == nil {
		return 0, dfs.WrapDfsStorage("put encrypted file", errors.New("object store unavailable"))
	}
	if fileID == 0 {
		return 0, dfs.ErrDfsInvalidArgument
	}
	size, err := r.objectStore.PutEncryptedFileReader(ctx, fmt.Sprintf("%d.dat", fileID), bytes.NewReader(data))
	if err != nil {
		return 0, dfs.WrapDfsStorage("put encrypted file", err)
	}
	return size, nil
}
