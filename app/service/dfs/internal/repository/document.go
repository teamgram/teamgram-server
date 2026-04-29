package repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/ffmpeg2"
)

type StoredDocumentThumb struct {
	Type  string
	W     int32
	H     int32
	Size  int32
	Bytes []byte
}

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

func (r *Repository) GetDocumentObject(ctx context.Context, path string, offset int64, limit int32) ([]byte, error) {
	if r == nil || r.objectStore == nil {
		return nil, dfs.WrapDfsStorage("get document file", errors.New("object store unavailable"))
	}
	data, err := r.objectStore.GetDocumentFile(ctx, path, offset, limit)
	if err != nil {
		return nil, dfs.WrapDfsStorage("get document file", err)
	}
	return data, nil
}

func (r *Repository) GetEncryptedObject(ctx context.Context, path string, offset int64, limit int32) ([]byte, error) {
	if r == nil || r.objectStore == nil {
		return nil, dfs.WrapDfsStorage("get encrypted file", errors.New("object store unavailable"))
	}
	data, err := r.objectStore.GetEncryptedFile(ctx, path, offset, limit)
	if err != nil {
		return nil, dfs.WrapDfsStorage("get encrypted file", err)
	}
	return data, nil
}

func (r *Repository) SaveDocumentThumbs(ctx context.Context, documentID int64, image []byte, ext string) ([]StoredDocumentThumb, error) {
	if documentID == 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	stripped, err := r.EncodeDocumentStrippedThumb(ctx, image)
	if err != nil {
		return nil, err
	}
	resized, err := r.ResizePhoto(ctx, image, ext, false)
	if err != nil {
		return nil, err
	}
	if len(resized) == 0 {
		return nil, dfs.ErrDfsImageProcessFailed
	}
	medium := resized[0]
	for _, size := range resized {
		if size.Type == "m" {
			medium = size
			break
		}
	}
	storedSize, err := r.PutPhotoBytes(ctx, fmt.Sprintf("%s/%d.dat", medium.Type, documentID), medium.Bytes)
	if err != nil {
		return nil, err
	}
	return []StoredDocumentThumb{
		{Type: "i", Bytes: stripped},
		{Type: medium.Type, W: medium.W, H: medium.H, Size: int32(storedSize)},
	}, nil
}

func (r *Repository) EncodeDocumentStrippedThumb(ctx context.Context, data []byte) ([]byte, error) {
	if r == nil || r.imaging == nil {
		return nil, dfs.WrapDfsStorage("encode document stripped thumb", errors.New("imaging processor unavailable"))
	}
	stripped, err := r.imaging.EncodeStripped(ctx, data)
	if err != nil {
		return nil, dfs.WrapDfsStorage("encode document stripped thumb", err)
	}
	return stripped, nil
}

func (r *Repository) ConvertDocumentToMP4(ctx context.Context, data []byte) ([]byte, error) {
	return r.ConvertToMP4(ctx, data)
}

func (r *Repository) ExtractDocumentFrame(ctx context.Context, data []byte) ([]byte, error) {
	return r.ExtractFirstFrame(ctx, data)
}

func (r *Repository) GetDocumentVideoMetadata(ctx context.Context, data []byte) (*ffmpeg2.VideoMetadata, error) {
	return r.GetVideoMetadata(ctx, data)
}
