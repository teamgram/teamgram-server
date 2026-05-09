package core

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/objectstore"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeFileObjectRepository struct {
	commitResult *dfs.FileFinalizedObject
	putResult    *dfs.FileFinalizedObject
	readData     []byte
	readType     int32
	hashes       []objectstore.HashChunk
	err          error

	commitCalled  bool
	commitUpload  string
	commitOwner   int64
	commitFile    tg.InputFileClazz
	commitPurpose string

	putCalled  bool
	putOwner   int64
	putPurpose string
	putName    string
	putMime    string
	putBytes   []byte

	readCalled bool
	readLease  []byte
	readOffset int64
	readLimit  int32

	hashCalled bool
	hashLease  []byte
	hashOffset int64
	hashLimit  int32
}

func (f *fakeFileObjectRepository) CommitUpload(_ context.Context, uploadSessionID string, ownerID int64, file tg.InputFileClazz, purpose string) (*dfs.FileFinalizedObject, error) {
	f.commitCalled = true
	f.commitUpload = uploadSessionID
	f.commitOwner = ownerID
	f.commitFile = file
	f.commitPurpose = purpose
	return f.commitResult, f.err
}

func (f *fakeFileObjectRepository) PutInternalFile(_ context.Context, ownerID int64, purpose, fileName, mimeType string, data []byte) (*dfs.FileFinalizedObject, error) {
	f.putCalled = true
	f.putOwner = ownerID
	f.putPurpose = purpose
	f.putName = fileName
	f.putMime = mimeType
	f.putBytes = append([]byte(nil), data...)
	return f.putResult, f.err
}

func (f *fakeFileObjectRepository) ReadByLease(_ context.Context, readLease []byte, offset int64, limit int32) ([]byte, int32, error) {
	f.readCalled = true
	f.readLease = append([]byte(nil), readLease...)
	f.readOffset = offset
	f.readLimit = limit
	return f.readData, f.readType, f.err
}

func (f *fakeFileObjectRepository) HashesByLease(_ context.Context, readLease []byte, offset int64, limit int32) ([]objectstore.HashChunk, error) {
	f.hashCalled = true
	f.hashLease = append([]byte(nil), readLease...)
	f.hashOffset = offset
	f.hashLimit = limit
	return f.hashes, f.err
}

func TestDfsCommitUploadInvalidInputReturnsInvalidArgument(t *testing.T) {
	core := newFileObjectTestCore(&fakeFileObjectRepository{})
	validFile := tg.MakeTLInputFile(&tg.TLInputFile{Id: 10, Parts: 1, Name: "a.bin"})

	tests := []struct {
		name string
		in   *dfs.TLDfsCommitUpload
	}{
		{name: "nil", in: nil},
		{name: "empty session", in: &dfs.TLDfsCommitUpload{OwnerId: 1001, File: validFile, Purpose: "media_original"}},
		{name: "zero owner", in: &dfs.TLDfsCommitUpload{UploadSessionId: "upload-1", File: validFile, Purpose: "media_original"}},
		{name: "nil file", in: &dfs.TLDfsCommitUpload{UploadSessionId: "upload-1", OwnerId: 1001, Purpose: "media_original"}},
		{name: "empty purpose", in: &dfs.TLDfsCommitUpload{UploadSessionId: "upload-1", OwnerId: 1001, File: validFile}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := core.DfsCommitUpload(tt.in); !errors.Is(err, dfs.ErrDfsInvalidArgument) {
				t.Fatalf("DfsCommitUpload() error = %v, want ErrDfsInvalidArgument", err)
			}
		})
	}
}

func TestDfsCommitUploadDelegatesAndReturnsFinalizedObject(t *testing.T) {
	file := tg.MakeTLInputFile(&tg.TLInputFile{Id: 10, Parts: 1, Name: "a.bin"})
	want := dfs.MakeTLFileFinalizedObject(&dfs.TLFileFinalizedObject{ObjectId: "obj-1"})
	repo := &fakeFileObjectRepository{commitResult: want}
	core := newFileObjectTestCore(repo)

	got, err := core.DfsCommitUpload(&dfs.TLDfsCommitUpload{
		UploadSessionId: "upload-1",
		OwnerId:         1001,
		File:            file,
		Purpose:         "media_original",
	})
	if err != nil {
		t.Fatalf("DfsCommitUpload() error = %v", err)
	}
	if got != want {
		t.Fatalf("DfsCommitUpload() = %#v, want same result pointer", got)
	}
	if !repo.commitCalled || repo.commitUpload != "upload-1" || repo.commitOwner != 1001 || repo.commitFile != file || repo.commitPurpose != "media_original" {
		t.Fatalf("CommitUpload call = %#v", repo)
	}
}

func TestDfsPutFileInvalidInputReturnsInvalidArgument(t *testing.T) {
	core := newFileObjectTestCore(&fakeFileObjectRepository{})

	tests := []struct {
		name string
		in   *dfs.TLDfsPutFile
	}{
		{name: "nil", in: nil},
		{name: "zero owner", in: &dfs.TLDfsPutFile{Purpose: "media_derivative", FileName: "a.bin", Bytes: []byte("a")}},
		{name: "empty purpose", in: &dfs.TLDfsPutFile{OwnerId: 1001, FileName: "a.bin", Bytes: []byte("a")}},
		{name: "empty file name", in: &dfs.TLDfsPutFile{OwnerId: 1001, Purpose: "media_derivative", Bytes: []byte("a")}},
		{name: "empty bytes", in: &dfs.TLDfsPutFile{OwnerId: 1001, Purpose: "media_derivative", FileName: "a.bin"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := core.DfsPutFile(tt.in); !errors.Is(err, dfs.ErrDfsInvalidArgument) {
				t.Fatalf("DfsPutFile() error = %v, want ErrDfsInvalidArgument", err)
			}
		})
	}
}

func TestDfsPutFileDelegatesAndReturnsFinalizedObject(t *testing.T) {
	want := dfs.MakeTLFileFinalizedObject(&dfs.TLFileFinalizedObject{ObjectId: "obj-2"})
	repo := &fakeFileObjectRepository{putResult: want}
	core := newFileObjectTestCore(repo)

	got, err := core.DfsPutFile(&dfs.TLDfsPutFile{
		OwnerId:  1001,
		Purpose:  "media_derivative",
		FileName: "a.bin",
		MimeType: "application/octet-stream",
		Bytes:    []byte("payload"),
	})
	if err != nil {
		t.Fatalf("DfsPutFile() error = %v", err)
	}
	if got != want {
		t.Fatalf("DfsPutFile() = %#v, want same result pointer", got)
	}
	if !repo.putCalled || repo.putOwner != 1001 || repo.putPurpose != "media_derivative" || repo.putName != "a.bin" || repo.putMime != "application/octet-stream" || string(repo.putBytes) != "payload" {
		t.Fatalf("PutInternalFile call = %#v", repo)
	}
}

func TestDfsGetFileByReadLeaseInvalidInputReturnsInvalidArgument(t *testing.T) {
	core := newFileObjectTestCore(&fakeFileObjectRepository{})

	tests := []struct {
		name string
		in   *dfs.TLDfsGetFileByReadLease
	}{
		{name: "nil", in: nil},
		{name: "empty lease", in: &dfs.TLDfsGetFileByReadLease{}},
		{name: "negative offset", in: &dfs.TLDfsGetFileByReadLease{ReadLease: []byte("lease"), Offset: -1}},
		{name: "negative limit", in: &dfs.TLDfsGetFileByReadLease{ReadLease: []byte("lease"), Limit: -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := core.DfsGetFileByReadLease(tt.in); !errors.Is(err, dfs.ErrDfsInvalidArgument) {
				t.Fatalf("DfsGetFileByReadLease() error = %v, want ErrDfsInvalidArgument", err)
			}
		})
	}
}

func TestDfsGetFileByReadLeaseReturnsUploadFile(t *testing.T) {
	repo := &fakeFileObjectRepository{
		readData: []byte("abcdef"),
		readType: int32(tg.ClazzID_storage_filePng),
	}
	core := newFileObjectTestCore(repo)

	got, err := core.DfsGetFileByReadLease(&dfs.TLDfsGetFileByReadLease{
		ReadLease: []byte("lease"),
		Offset:    2,
		Limit:     3,
	})
	if err != nil {
		t.Fatalf("DfsGetFileByReadLease() error = %v", err)
	}
	file, ok := got.ToUploadFile()
	if !ok {
		t.Fatalf("UploadFile clazz = %s, want upload.file", got.ClazzName())
	}
	if string(file.Bytes) != "abcdef" {
		t.Fatalf("Bytes = %q, want abcdef", file.Bytes)
	}
	if _, ok := file.Type.(*tg.TLStorageFilePng); !ok {
		t.Fatalf("Type = %T, want storage.filePng", file.Type)
	}
	if !repo.readCalled || string(repo.readLease) != "lease" || repo.readOffset != 2 || repo.readLimit != 3 {
		t.Fatalf("ReadByLease call = %#v", repo)
	}
}

func TestDfsGetFileHashesByReadLeaseInvalidInputReturnsInvalidArgument(t *testing.T) {
	core := newFileObjectTestCore(&fakeFileObjectRepository{})

	tests := []struct {
		name string
		in   *dfs.TLDfsGetFileHashesByReadLease
	}{
		{name: "nil", in: nil},
		{name: "empty lease", in: &dfs.TLDfsGetFileHashesByReadLease{}},
		{name: "negative offset", in: &dfs.TLDfsGetFileHashesByReadLease{ReadLease: []byte("lease"), Offset: -1}},
		{name: "negative limit", in: &dfs.TLDfsGetFileHashesByReadLease{ReadLease: []byte("lease"), Limit: -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := core.DfsGetFileHashesByReadLease(tt.in); !errors.Is(err, dfs.ErrDfsInvalidArgument) {
				t.Fatalf("DfsGetFileHashesByReadLease() error = %v, want ErrDfsInvalidArgument", err)
			}
		})
	}
}

func TestDfsGetFileHashesByReadLeaseReturnsVector(t *testing.T) {
	repo := &fakeFileObjectRepository{
		hashes: []objectstore.HashChunk{
			{Offset: 0, Limit: 4, Hash: []byte("hash-1")},
			{Offset: 4, Limit: 2, Hash: []byte("hash-2")},
		},
	}
	core := newFileObjectTestCore(repo)

	got, err := core.DfsGetFileHashesByReadLease(&dfs.TLDfsGetFileHashesByReadLease{
		ReadLease: []byte("lease"),
		Offset:    4,
		Limit:     8,
	})
	if err != nil {
		t.Fatalf("DfsGetFileHashesByReadLease() error = %v", err)
	}
	if !repo.hashCalled || string(repo.hashLease) != "lease" || repo.hashOffset != 4 || repo.hashLimit != 8 {
		t.Fatalf("HashesByLease call = %#v", repo)
	}
	want := []tg.FileHashClazz{
		tg.MakeTLFileHash(&tg.TLFileHash{Offset: 0, Limit: 4, Hash: []byte("hash-1")}),
		tg.MakeTLFileHash(&tg.TLFileHash{Offset: 4, Limit: 2, Hash: []byte("hash-2")}),
	}
	if !reflect.DeepEqual(got.Datas, want) {
		t.Fatalf("Hashes = %#v, want %#v", got.Datas, want)
	}
}

func TestDfsFileObjectHandlersPropagateRepositoryErrors(t *testing.T) {
	repoErr := dfs.WrapDfsStorage("test", errors.New("disk failed"))
	file := tg.MakeTLInputFile(&tg.TLInputFile{Id: 10, Parts: 1, Name: "a.bin"})

	t.Run("commit", func(t *testing.T) {
		core := newFileObjectTestCore(&fakeFileObjectRepository{err: repoErr})
		_, err := core.DfsCommitUpload(&dfs.TLDfsCommitUpload{UploadSessionId: "upload-1", OwnerId: 1001, File: file, Purpose: "media_original"})
		if !errors.Is(err, repoErr) || !errors.Is(err, dfs.ErrDfsStorage) {
			t.Fatalf("DfsCommitUpload() error = %v, want repo storage error", err)
		}
	})
	t.Run("put", func(t *testing.T) {
		core := newFileObjectTestCore(&fakeFileObjectRepository{err: repoErr})
		_, err := core.DfsPutFile(&dfs.TLDfsPutFile{OwnerId: 1001, Purpose: "media_derivative", FileName: "a.bin", Bytes: []byte("payload")})
		if !errors.Is(err, repoErr) || !errors.Is(err, dfs.ErrDfsStorage) {
			t.Fatalf("DfsPutFile() error = %v, want repo storage error", err)
		}
	})
	t.Run("read", func(t *testing.T) {
		core := newFileObjectTestCore(&fakeFileObjectRepository{err: repoErr})
		_, err := core.DfsGetFileByReadLease(&dfs.TLDfsGetFileByReadLease{ReadLease: []byte("lease")})
		if !errors.Is(err, repoErr) || !errors.Is(err, dfs.ErrDfsStorage) {
			t.Fatalf("DfsGetFileByReadLease() error = %v, want repo storage error", err)
		}
	})
	t.Run("hashes", func(t *testing.T) {
		core := newFileObjectTestCore(&fakeFileObjectRepository{err: repoErr})
		_, err := core.DfsGetFileHashesByReadLease(&dfs.TLDfsGetFileHashesByReadLease{ReadLease: []byte("lease")})
		if !errors.Is(err, repoErr) || !errors.Is(err, dfs.ErrDfsStorage) {
			t.Fatalf("DfsGetFileHashesByReadLease() error = %v, want repo storage error", err)
		}
	})
}

func newFileObjectTestCore(repo fileObjectRepository) *DfsCore {
	return &DfsCore{
		ctx:                  context.Background(),
		fileObjectRepository: repo,
	}
}
