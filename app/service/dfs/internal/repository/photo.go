package repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/ffmpeg2"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/imaging2"
)

type StoredPhotoSize struct {
	Type string
	W    int32
	H    int32
	Size int32
}

type StoredPhoto struct {
	ID    int64
	Sizes []StoredPhotoSize
}

func (r *Repository) NextPhotoID(ctx context.Context) (int64, error) {
	if r == nil || r.idgen == nil {
		return 0, dfs.WrapDfsDownstream("next photo id", errors.New("idgen client unavailable"))
	}
	return r.idgen.NextPhotoID(ctx)
}

func (r *Repository) PutPhotoBytes(ctx context.Context, path string, data []byte) (int64, error) {
	if r == nil || r.objectStore == nil {
		return 0, dfs.WrapDfsStorage("put photo file", errors.New("object store unavailable"))
	}
	size, err := r.objectStore.PutPhotoBytes(ctx, path, data)
	if err != nil {
		return 0, dfs.WrapDfsStorage("put photo file", err)
	}
	return size, nil
}

func (r *Repository) SavePhotoObjects(ctx context.Context, photoID int64, original []byte, ext string, isABC bool, storeOriginal bool) (*StoredPhoto, error) {
	if photoID == 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	if storeOriginal {
		if _, err := r.PutPhotoBytes(ctx, fmt.Sprintf("0/%d.dat", photoID), original); err != nil {
			return nil, err
		}
	}
	resized, err := r.ResizePhoto(ctx, original, ext, isABC)
	if err != nil {
		return nil, err
	}
	if len(resized) == 0 {
		return nil, dfs.ErrDfsImageProcessFailed
	}
	out := &StoredPhoto{ID: photoID, Sizes: make([]StoredPhotoSize, 0, len(resized))}
	for _, size := range resized {
		storedSize, err := r.PutPhotoBytes(ctx, fmt.Sprintf("%s/%d.dat", size.Type, photoID), size.Bytes)
		if err != nil {
			return nil, err
		}
		out.Sizes = append(out.Sizes, StoredPhotoSize{
			Type: size.Type,
			W:    size.W,
			H:    size.H,
			Size: int32(storedSize),
		})
	}
	return out, nil
}

func (r *Repository) GetPhotoFile(ctx context.Context, path string) ([]byte, error) {
	if r == nil || r.objectStore == nil {
		return nil, dfs.WrapDfsStorage("get photo file", errors.New("object store unavailable"))
	}
	data, err := r.objectStore.GetPhotoFile(ctx, path, 0, 0)
	if err != nil {
		return nil, dfs.WrapDfsStorage("get photo file", err)
	}
	return data, nil
}

func (r *Repository) LoadOriginalPhotoBytes(ctx context.Context, photoID int64) ([]byte, error) {
	if photoID == 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	return r.GetPhotoFile(ctx, fmt.Sprintf("0/%d.dat", photoID))
}

func (r *Repository) PutVideoBytes(ctx context.Context, path string, data []byte) (int64, error) {
	if r == nil || r.objectStore == nil {
		return 0, dfs.WrapDfsStorage("put video file", errors.New("object store unavailable"))
	}
	size, err := r.objectStore.PutVideoBytes(ctx, path, data)
	if err != nil {
		return 0, dfs.WrapDfsStorage("put video file", err)
	}
	return size, nil
}

func (r *Repository) SaveProfileVideoObject(ctx context.Context, photoID int64, data []byte) (int64, error) {
	if photoID == 0 {
		return 0, dfs.ErrDfsInvalidArgument
	}
	return r.PutVideoBytes(ctx, fmt.Sprintf("v/%d.dat", photoID), data)
}

func (r *Repository) ResizePhoto(ctx context.Context, original []byte, ext string, isABC bool) ([]imaging2.PhotoSizeBytes, error) {
	if r == nil || r.imaging == nil {
		return nil, dfs.WrapDfsStorage("resize photo", errors.New("imaging processor unavailable"))
	}
	sizes, err := r.imaging.ResizePhoto(ctx, original, ext, isABC)
	if err != nil {
		return nil, dfs.WrapDfsStorage("resize photo", err)
	}
	return sizes, nil
}

func (r *Repository) ConvertToMP4(ctx context.Context, data []byte) ([]byte, error) {
	if r == nil || r.ffmpeg == nil {
		return nil, dfs.WrapDfsStorage("convert video to mp4", errors.New("ffmpeg processor unavailable"))
	}
	out, err := r.ffmpeg.ConvertToMP4(ctx, bytes.NewReader(data))
	if err != nil {
		return nil, dfs.WrapDfsStorage("convert video to mp4", err)
	}
	return out, nil
}

func (r *Repository) ExtractFirstFrame(ctx context.Context, data []byte) ([]byte, error) {
	if r == nil || r.ffmpeg == nil {
		return nil, dfs.WrapDfsStorage("extract first frame", errors.New("ffmpeg processor unavailable"))
	}
	frame, err := r.ffmpeg.ExtractFirstFrame(ctx, bytes.NewReader(data))
	if err != nil {
		return nil, dfs.WrapDfsStorage("extract first frame", err)
	}
	return frame, nil
}

func (r *Repository) GetVideoMetadata(ctx context.Context, data []byte) (*ffmpeg2.VideoMetadata, error) {
	if r == nil || r.ffmpeg == nil {
		return nil, dfs.WrapDfsStorage("get video metadata", errors.New("ffmpeg processor unavailable"))
	}
	metadata, err := r.ffmpeg.GetVideoMetadata(ctx, bytes.NewReader(data))
	if err != nil {
		return nil, dfs.WrapDfsStorage("get video metadata", err)
	}
	return metadata, nil
}
