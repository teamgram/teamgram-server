package core

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeDownloadRepository struct {
	photos    map[string][]byte
	videos    map[string][]byte
	documents map[string][]byte
	encrypted map[string][]byte
}

func newFakeDownloadRepository() *fakeDownloadRepository {
	return &fakeDownloadRepository{
		photos:    make(map[string][]byte),
		videos:    make(map[string][]byte),
		documents: make(map[string][]byte),
		encrypted: make(map[string][]byte),
	}
}

func (f *fakeDownloadRepository) GetPhotoObject(_ context.Context, path string, offset int64, limit int32) ([]byte, error) {
	return rangeBytes(f.photos[path], offset, limit)
}

func (f *fakeDownloadRepository) GetVideoObject(_ context.Context, path string, offset int64, limit int32) ([]byte, error) {
	return rangeBytes(f.videos[path], offset, limit)
}

func (f *fakeDownloadRepository) GetDocumentObject(_ context.Context, path string, offset int64, limit int32) ([]byte, error) {
	return rangeBytes(f.documents[path], offset, limit)
}

func (f *fakeDownloadRepository) GetEncryptedObject(_ context.Context, path string, offset int64, limit int32) ([]byte, error) {
	return rangeBytes(f.encrypted[path], offset, limit)
}

func rangeBytes(data []byte, offset int64, limit int32) ([]byte, error) {
	if data == nil {
		return nil, dfs.ErrDfsFileNotFound
	}
	if offset < 0 || limit < 0 {
		return nil, dfs.ErrDfsInvalidArgument
	}
	if offset >= int64(len(data)) {
		return []byte{}, nil
	}
	out := data[offset:]
	if limit > 0 && int(limit) < len(out) {
		out = out[:limit]
	}
	return append([]byte(nil), out...), nil
}

func TestDfsDownloadFileDocumentUsesCacheThenMinio(t *testing.T) {
	core, uploads, downloads := newDownloadTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 2002, []byte("cached-document"))
	if err := uploads.SaveObjectCacheRef(context.Background(), 7001, 1001, 2002); err != nil {
		t.Fatalf("SaveObjectCacheRef() error = %v", err)
	}
	downloads.documents["7001.dat"] = []byte("minio-document")

	out, err := core.DfsDownloadFile(&dfs.TLDfsDownloadFile{
		Location: tg.MakeTLInputDocumentFileLocation(&tg.TLInputDocumentFileLocation{
			Id:         7001,
			AccessHash: minioadapter.MakeAccessHash(storageConstructor(tg.ClazzID_storage_filePdf), 1),
		}),
		Offset: 2,
		Limit:  4,
	})
	if err != nil {
		t.Fatalf("DfsDownloadFile() error = %v", err)
	}
	file, ok := out.ToUploadFile()
	if !ok {
		t.Fatalf("UploadFile clazz = %s, want upload.file", out.ClazzName())
	}
	if string(file.Bytes) != "ched" {
		t.Fatalf("Bytes = %q, want cached range", file.Bytes)
	}
	if _, ok := file.Type.(*tg.TLStorageFilePdf); !ok {
		t.Fatalf("Type = %T, want storage.filePdf", file.Type)
	}
}

func TestDfsDownloadFilePhotoSizeReadsPhotoBucket(t *testing.T) {
	core, _, downloads := newDownloadTestCore(t)
	downloads.photos["m/9001.dat"] = []byte("photo-medium")

	out, err := core.DfsDownloadFile(&dfs.TLDfsDownloadFile{
		Location: tg.MakeTLInputPhotoFileLocation(&tg.TLInputPhotoFileLocation{
			Id:         9001,
			AccessHash: minioadapter.MakeAccessHash(storageConstructor(tg.ClazzID_storage_fileJpeg), 1),
			ThumbSize:  "m",
		}),
		Offset: 0,
		Limit:  0,
	})
	if err != nil {
		t.Fatalf("DfsDownloadFile(photo) error = %v", err)
	}
	file, _ := out.ToUploadFile()
	if string(file.Bytes) != "photo-medium" {
		t.Fatalf("Bytes = %q, want photo-medium", file.Bytes)
	}
	if _, ok := file.Type.(*tg.TLStorageFileJpeg); !ok {
		t.Fatalf("Type = %T, want storage.fileJpeg", file.Type)
	}
}

func TestDfsDownloadFileDocumentThumbReturnsJpegType(t *testing.T) {
	core, _, downloads := newDownloadTestCore(t)
	downloads.photos["m/7001.dat"] = []byte("document-thumb")

	out, err := core.DfsDownloadFile(&dfs.TLDfsDownloadFile{
		Location: tg.MakeTLInputDocumentFileLocation(&tg.TLInputDocumentFileLocation{
			Id:         7001,
			AccessHash: minioadapter.MakeAccessHash(storageConstructor(tg.ClazzID_storage_filePdf), 1),
			ThumbSize:  "m",
		}),
	})
	if err != nil {
		t.Fatalf("DfsDownloadFile(document thumb) error = %v", err)
	}
	file, _ := out.ToUploadFile()
	if string(file.Bytes) != "document-thumb" {
		t.Fatalf("Bytes = %q, want document-thumb", file.Bytes)
	}
	if _, ok := file.Type.(*tg.TLStorageFileJpeg); !ok {
		t.Fatalf("Type = %T, want storage.fileJpeg", file.Type)
	}
}

func TestDfsDownloadFileInvalidConstructorReturnsInvalidArgument(t *testing.T) {
	core, _, _ := newDownloadTestCore(t)
	_, err := core.DfsDownloadFile(&dfs.TLDfsDownloadFile{
		Location: tg.MakeTLInputTakeoutFileLocation(&tg.TLInputTakeoutFileLocation{}),
	})
	if !errors.Is(err, dfs.ErrDfsInvalidArgument) {
		t.Fatalf("DfsDownloadFile() error = %v, want ErrDfsInvalidArgument", err)
	}
}

func TestDfsDownloadFileMinioMissReturnsEmptyBytesForCompatibleBranches(t *testing.T) {
	core, _, _ := newDownloadTestCore(t)
	out, err := core.DfsDownloadFile(&dfs.TLDfsDownloadFile{
		Location: tg.MakeTLInputDocumentFileLocation(&tg.TLInputDocumentFileLocation{
			Id:         7001,
			AccessHash: minioadapter.MakeAccessHash(storageConstructor(tg.ClazzID_storage_filePdf), 1),
		}),
	})
	if err != nil {
		t.Fatalf("DfsDownloadFile() error = %v, want compatible empty success", err)
	}
	file, _ := out.ToUploadFile()
	if !bytes.Equal(file.Bytes, []byte{}) {
		t.Fatalf("Bytes = %q, want empty", file.Bytes)
	}
}

func newDownloadTestCore(t *testing.T) (*DfsCore, *UploadSessionManager, *fakeDownloadRepository) {
	t.Helper()
	uploadRepo := newFakeUploadStateRepository()
	uploads := NewUploadSessionManager(uploadRepo)
	downloads := newFakeDownloadRepository()
	return &DfsCore{
		ctx:                  context.Background(),
		uploadSessionManager: uploads,
		downloadRepository:   downloads,
	}, uploads, downloads
}
