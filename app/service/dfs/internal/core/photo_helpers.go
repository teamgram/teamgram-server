package core

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"path"
	"strings"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type photoRepository interface {
	NextPhotoID(ctx context.Context) (int64, error)
	SavePhotoObjects(ctx context.Context, photoID int64, original []byte, ext string, isABC bool, storeOriginal bool) (*repository.StoredPhoto, error)
	LoadOriginalPhotoBytes(ctx context.Context, photoID int64) ([]byte, error)
	SaveProfileVideoObject(ctx context.Context, photoID int64, data []byte) (int64, error)
	ConvertToMP4(ctx context.Context, data []byte) ([]byte, error)
	ExtractFirstFrame(ctx context.Context, data []byte) ([]byte, error)
}

type uploadedPhotoFile struct {
	id          int64
	parts       int32
	name        string
	md5Checksum string
}

func (c *DfsCore) photos() photoRepository {
	if c.photoRepository != nil {
		return c.photoRepository
	}
	return c.svcCtx.Repo
}

func (c *DfsCore) buildPhotoFromUpload(creator int64, file *uploadedPhotoFile, isABC bool, date int32, videoSizes []tg.VideoSizeClazz) (*tg.Photo, error) {
	if file == nil {
		return nil, dfs.ErrDfsInvalidArgument
	}
	if err := checkFileParts(file.parts); err != nil {
		return nil, err
	}
	reader, err := c.uploadSessions().OpenUploadedFile(c.ctx, creator, file.id)
	if err != nil {
		return nil, err
	}
	data, err := readAllSeeker(reader)
	if err != nil {
		return nil, dfs.WrapDfsStorage("read uploaded photo", err)
	}
	if err := checkMD5(data, file.md5Checksum); err != nil {
		return nil, err
	}
	return c.buildPhotoFromBytes(data, fileExt(file.name), isABC, date, videoSizes, true)
}

func (c *DfsCore) buildPhotoFromBytes(data []byte, ext string, isABC bool, date int32, videoSizes []tg.VideoSizeClazz, storeOriginal bool) (*tg.Photo, error) {
	repo := c.photos()
	if repo == nil {
		return nil, dfs.WrapDfsStorage("build photo", errors.New("photo repository unavailable"))
	}
	photoID, err := repo.NextPhotoID(c.ctx)
	if err != nil {
		return nil, err
	}
	stored, err := repo.SavePhotoObjects(c.ctx, photoID, data, ext, isABC, storeOriginal)
	if err != nil {
		if errors.Is(err, dfs.ErrDfsImageProcessFailed) {
			return nil, fmt.Errorf("%w: %w", dfs.ErrDfsImageProcessFailed, err)
		}
		return nil, err
	}
	if stored == nil || len(stored.Sizes) == 0 {
		return nil, dfs.ErrDfsImageProcessFailed
	}
	sizes := make([]tg.PhotoSizeClazz, 0, len(stored.Sizes))
	for _, size := range stored.Sizes {
		sizes = append(sizes, tg.MakeTLPhotoSize(&tg.TLPhotoSize{
			Type:  size.Type,
			W:     size.W,
			H:     size.H,
			Size2: size.Size,
		}))
	}
	return tg.MakeTLPhoto(&tg.TLPhoto{
		HasStickers:   false,
		Id:            photoID,
		AccessHash:    minioadapter.MakeAccessHash(storageFileTypeConstructor(ext), rand32()),
		FileReference: []byte{},
		Date:          date,
		Sizes:         sizes,
		VideoSizes:    videoSizes,
		DcId:          1,
	}).ToPhoto(), nil
}

func inputFile(in tg.InputFileClazz) (*uploadedPhotoFile, error) {
	switch f := in.(type) {
	case *tg.TLInputFile:
		return &uploadedPhotoFile{id: f.Id, parts: f.Parts, name: f.Name, md5Checksum: f.Md5Checksum}, nil
	case *tg.TLInputFileBig:
		return &uploadedPhotoFile{id: f.Id, parts: f.Parts, name: f.Name}, nil
	default:
		return nil, dfs.ErrDfsInvalidArgument
	}
}

func checkFileParts(parts int32) error {
	if parts < 1 || parts > 3000 {
		return dfs.ErrDfsInvalidFilePart
	}
	return nil
}

func checkMD5(data []byte, expected string) error {
	if expected == "" {
		return nil
	}
	if fmt.Sprintf("%x", md5.Sum(data)) != strings.ToLower(expected) {
		return dfs.ErrDfsChecksumInvalid
	}
	return nil
}

func fileExt(name string) string {
	ext := strings.ToLower(path.Ext(name))
	if ext == "" {
		return ".partial"
	}
	return ext
}

func storageFileTypeConstructor(ext string) int32 {
	switch strings.ToLower(ext) {
	case ".jpeg", ".jpg":
		return storageConstructor(tg.ClazzID_storage_fileJpeg)
	case ".gif":
		return storageConstructor(tg.ClazzID_storage_fileGif)
	case ".png":
		return storageConstructor(tg.ClazzID_storage_filePng)
	case ".pdf":
		return storageConstructor(tg.ClazzID_storage_filePdf)
	case ".mp3":
		return storageConstructor(tg.ClazzID_storage_fileMp3)
	case ".mov":
		return storageConstructor(tg.ClazzID_storage_fileMov)
	case ".mp4":
		return storageConstructor(tg.ClazzID_storage_fileMp4)
	case ".webp":
		return storageConstructor(tg.ClazzID_storage_fileWebp)
	default:
		return storageConstructor(tg.ClazzID_storage_filePartial)
	}
}

func storageConstructor(id uint32) int32 {
	return int32(id)
}

func rand32() uint32 {
	return rand.Uint32()
}

func nowUnix() int32 {
	return int32(time.Now().Unix())
}

func readAllSeeker(r io.ReadSeeker) ([]byte, error) {
	if r == nil {
		return nil, dfs.ErrDfsFileNotFound
	}
	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	return io.ReadAll(r)
}
