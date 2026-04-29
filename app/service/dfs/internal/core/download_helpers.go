package core

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type downloadRepository interface {
	GetPhotoObject(ctx context.Context, path string, offset int64, limit int32) ([]byte, error)
	GetVideoObject(ctx context.Context, path string, offset int64, limit int32) ([]byte, error)
	GetDocumentObject(ctx context.Context, path string, offset int64, limit int32) ([]byte, error)
	GetEncryptedObject(ctx context.Context, path string, offset int64, limit int32) ([]byte, error)
}

type HttpUploadedFile struct {
	Name       string
	Modtime    time.Time
	ReadSeeker io.ReadSeeker
}

func (c *DfsCore) downloads() downloadRepository {
	if c.downloadRepository != nil {
		return c.downloadRepository
	}
	return c.svcCtx.Repo
}

func makeUploadFile(storageType int32, data []byte) *tg.UploadFile {
	if data == nil {
		data = []byte{}
	}
	return tg.MakeTLUploadFile(&tg.TLUploadFile{
		Type:  makeStorageFileType(storageType),
		Mtime: int32(time.Now().Unix()),
		Bytes: data,
	}).ToUploadFile()
}

func makeStorageFileType(storageType int32) tg.StorageFileTypeClazz {
	switch uint32(storageType) {
	case tg.ClazzID_storage_filePartial:
		return tg.MakeTLStorageFilePartial(&tg.TLStorageFilePartial{})
	case tg.ClazzID_storage_fileJpeg:
		return tg.MakeTLStorageFileJpeg(&tg.TLStorageFileJpeg{})
	case tg.ClazzID_storage_fileGif:
		return tg.MakeTLStorageFileGif(&tg.TLStorageFileGif{})
	case tg.ClazzID_storage_filePng:
		return tg.MakeTLStorageFilePng(&tg.TLStorageFilePng{})
	case tg.ClazzID_storage_filePdf:
		return tg.MakeTLStorageFilePdf(&tg.TLStorageFilePdf{})
	case tg.ClazzID_storage_fileMp3:
		return tg.MakeTLStorageFileMp3(&tg.TLStorageFileMp3{})
	case tg.ClazzID_storage_fileMov:
		return tg.MakeTLStorageFileMov(&tg.TLStorageFileMov{})
	case tg.ClazzID_storage_fileMp4:
		return tg.MakeTLStorageFileMp4(&tg.TLStorageFileMp4{})
	case tg.ClazzID_storage_fileWebp:
		return tg.MakeTLStorageFileWebp(&tg.TLStorageFileWebp{})
	default:
		return tg.MakeTLStorageFileUnknown(&tg.TLStorageFileUnknown{})
	}
}

func storageTypeFromAccessHash(accessHash int64) int32 {
	return minioadapter.StorageTypeFromAccessHash(accessHash)
}

func photoSizeIsVideo(size string) bool {
	return size == "v" || size == "u"
}

func stickerSetID(stickerSet tg.InputStickerSetClazz) (int64, bool) {
	if set, ok := stickerSet.(*tg.TLInputStickerSetID); ok {
		return set.Id, set.Id != 0
	}
	return 0, false
}

func objectPath(size string, id int64) string {
	return fmt.Sprintf("%s/%d.dat", size, id)
}

func compatibleBytes(data []byte, err error) ([]byte, error) {
	if err == nil {
		return data, nil
	}
	if errors.Is(err, dfs.ErrDfsFileNotFound) || errors.Is(err, dfs.ErrDfsStorage) {
		return []byte{}, nil
	}
	return nil, err
}

func (c *DfsCore) compatibleDownloadBytes(op string, data []byte, err error) ([]byte, error) {
	if err == nil {
		return data, nil
	}
	if errors.Is(err, dfs.ErrDfsFileNotFound) || errors.Is(err, dfs.ErrDfsStorage) {
		c.logNonFatalError(op, err)
		return []byte{}, nil
	}
	return nil, err
}

func (c *DfsCore) OpenHttpUploadedFile(ctx context.Context, creator, fileID int64) (*HttpUploadedFile, error) {
	if creator == 0 || fileID == 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	info, err := c.uploadSessions().LoadUploadedFileInfo(ctx, creator, fileID)
	if err != nil {
		return nil, err
	}
	reader, err := c.uploadSessions().OpenUploadedFile(ctx, creator, fileID)
	if err != nil {
		return nil, err
	}
	modtime := time.Unix(info.Mtime, 0)
	if info.Mtime == 0 {
		modtime = time.Now()
	}
	return &HttpUploadedFile{
		Name:       info.FileName,
		Modtime:    modtime,
		ReadSeeker: reader,
	}, nil
}
