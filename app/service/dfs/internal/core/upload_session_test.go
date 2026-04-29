package core

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeUploadStateRepository struct {
	parts map[UploadRef]map[int32][]byte
	infos map[UploadRef]*DfsFileInfo
	refs  map[int64]UploadRef
}

func newFakeUploadStateRepository() *fakeUploadStateRepository {
	return &fakeUploadStateRepository{
		parts: make(map[UploadRef]map[int32][]byte),
		infos: make(map[UploadRef]*DfsFileInfo),
		refs:  make(map[int64]UploadRef),
	}
}

func (f *fakeUploadStateRepository) SaveUploadPart(_ context.Context, ref UploadRef, partIndex int32, data []byte) error {
	if f.parts[ref] == nil {
		f.parts[ref] = make(map[int32][]byte)
	}
	f.parts[ref][partIndex] = append([]byte(nil), data...)
	return nil
}

func (f *fakeUploadStateRepository) SaveUploadFileInfo(_ context.Context, info *DfsFileInfo) error {
	cp := *info
	f.infos[UploadRef{Creator: info.Creator, FileID: info.FileID}] = &cp
	return nil
}

func (f *fakeUploadStateRepository) LoadUploadFileInfo(_ context.Context, creator, fileID int64) (*DfsFileInfo, error) {
	info := f.infos[UploadRef{Creator: creator, FileID: fileID}]
	if info == nil {
		return nil, dfs.ErrDfsFileNotFound
	}
	cp := *info
	return &cp, nil
}

func (f *fakeUploadStateRepository) OpenUploadFileReader(_ context.Context, info *DfsFileInfo) (io.ReadSeeker, error) {
	partMap := f.parts[UploadRef{Creator: info.Creator, FileID: info.FileID}]
	if partMap == nil {
		return nil, dfs.ErrDfsFileNotFound
	}
	var b bytes.Buffer
	for i := int32(0); i < int32(info.FileTotalParts); i++ {
		part, ok := partMap[i]
		if !ok {
			return nil, dfs.ErrDfsFileNotFound
		}
		b.Write(part)
	}
	return bytes.NewReader(b.Bytes()), nil
}

func (f *fakeUploadStateRepository) SaveObjectCacheRef(_ context.Context, objectID int64, info *DfsFileInfo) error {
	f.refs[objectID] = UploadRef{Creator: info.Creator, FileID: info.FileID}
	return nil
}

func (f *fakeUploadStateRepository) LoadObjectCacheRef(ctx context.Context, objectID int64) (*DfsFileInfo, error) {
	ref, ok := f.refs[objectID]
	if !ok {
		return nil, dfs.ErrDfsFileNotFound
	}
	return f.LoadUploadFileInfo(ctx, ref.Creator, ref.FileID)
}

func TestUploadSessionWritePartUpdatesFileInfo(t *testing.T) {
	ctx := context.Background()
	repo := newFakeUploadStateRepository()
	manager := NewUploadSessionManager(repo)
	manager.now = func() time.Time { return time.Unix(1_700_000_000, 0) }
	totalParts := int32(3)

	for _, tc := range []struct {
		part int32
		data []byte
	}{
		{part: 0, data: []byte("first")},
		{part: 1, data: []byte("middle")},
		{part: 2, data: []byte("last")},
	} {
		if err := manager.WritePart(ctx, WritePartCommand{
			Creator:        1001,
			FileID:         2002,
			FilePart:       tc.part,
			Bytes:          tc.data,
			Big:            true,
			FileTotalParts: &totalParts,
		}); err != nil {
			t.Fatalf("WritePart(part %d) error = %v", tc.part, err)
		}
	}

	info, err := repo.LoadUploadFileInfo(ctx, 1001, 2002)
	if err != nil {
		t.Fatalf("LoadUploadFileInfo() error = %v", err)
	}
	if !info.Big {
		t.Fatal("Big = false, want true")
	}
	if info.FileTotalParts != 3 {
		t.Fatalf("FileTotalParts = %d, want 3", info.FileTotalParts)
	}
	if info.FirstFilePartSize != len("first") {
		t.Fatalf("FirstFilePartSize = %d, want %d", info.FirstFilePartSize, len("first"))
	}
	if info.FilePartSize != len("middle") {
		t.Fatalf("FilePartSize = %d, want %d", info.FilePartSize, len("middle"))
	}
	if info.LastFilePartSize != len("last") {
		t.Fatalf("LastFilePartSize = %d, want %d", info.LastFilePartSize, len("last"))
	}
	if info.Mtime != 1_700_000_000 {
		t.Fatalf("Mtime = %d, want 1700000000", info.Mtime)
	}
	if got, want := info.FileSize(), int64(len("first")+len("middle")+len("last")); got != want {
		t.Fatalf("FileSize() = %d, want %d", got, want)
	}
}

func TestUploadSessionRejectsEmptyPart(t *testing.T) {
	ctx := context.Background()
	manager := NewUploadSessionManager(newFakeUploadStateRepository())
	totalParts := int32(1)

	err := manager.WritePart(ctx, WritePartCommand{
		Creator:        1001,
		FileID:         2002,
		FilePart:       0,
		FileTotalParts: &totalParts,
	})
	if !errors.Is(err, dfs.ErrDfsInvalidFilePart) {
		t.Fatalf("WritePart() error = %v, want ErrDfsInvalidFilePart", err)
	}
}

func TestUploadSessionInfersTotalPartsWhenOptionalTotalMissing(t *testing.T) {
	ctx := context.Background()
	manager := NewUploadSessionManager(newFakeUploadStateRepository())

	for i, data := range []string{"aa", "bb"} {
		if err := manager.WritePart(ctx, WritePartCommand{
			Creator:  1001,
			FileID:   2002,
			FilePart: int32(i),
			Bytes:    []byte(data),
		}); err != nil {
			t.Fatalf("WritePart(part %d) error = %v", i, err)
		}
	}

	reader, err := manager.OpenUploadedFile(ctx, 1001, 2002)
	if err != nil {
		t.Fatalf("OpenUploadedFile() error = %v", err)
	}
	got, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}
	if string(got) != "aabb" {
		t.Fatalf("OpenUploadedFile bytes = %q, want %q", got, "aabb")
	}
}

func TestUploadSessionReaderPreservesPartOrder(t *testing.T) {
	ctx := context.Background()
	manager := NewUploadSessionManager(newFakeUploadStateRepository())
	totalParts := int32(3)

	for _, part := range []struct {
		index int32
		data  string
	}{
		{index: 2, data: "cc"},
		{index: 0, data: "aa"},
		{index: 1, data: "bb"},
	} {
		if err := manager.WritePart(ctx, WritePartCommand{
			Creator:        1001,
			FileID:         2002,
			FilePart:       part.index,
			Bytes:          []byte(part.data),
			FileTotalParts: &totalParts,
		}); err != nil {
			t.Fatalf("WritePart(part %d) error = %v", part.index, err)
		}
	}

	reader, err := manager.OpenUploadedFile(ctx, 1001, 2002)
	if err != nil {
		t.Fatalf("OpenUploadedFile() error = %v", err)
	}
	got, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}
	if string(got) != "aabbcc" {
		t.Fatalf("OpenUploadedFile bytes = %q, want %q", got, "aabbcc")
	}
}

func TestUploadSessionReaderSupportsOffsetLimitAcrossParts(t *testing.T) {
	ctx := context.Background()
	manager := NewUploadSessionManager(newFakeUploadStateRepository())
	totalParts := int32(3)

	for i, data := range []string{"abc", "def", "ghi"} {
		if err := manager.WritePart(ctx, WritePartCommand{
			Creator:        1001,
			FileID:         2002,
			FilePart:       int32(i),
			Bytes:          []byte(data),
			FileTotalParts: &totalParts,
		}); err != nil {
			t.Fatalf("WritePart(part %d) error = %v", i, err)
		}
	}

	got, err := manager.ReadUploadedFileRange(ctx, 1001, 2002, 2, 5)
	if err != nil {
		t.Fatalf("ReadUploadedFileRange() error = %v", err)
	}
	if string(got) != "cdefg" {
		t.Fatalf("ReadUploadedFileRange() = %q, want %q", got, "cdefg")
	}
}

func TestUploadSessionCacheRefRoundTrip(t *testing.T) {
	ctx := context.Background()
	manager := NewUploadSessionManager(newFakeUploadStateRepository())
	totalParts := int32(1)

	if err := manager.WritePart(ctx, WritePartCommand{
		Creator:        1001,
		FileID:         2002,
		FilePart:       0,
		Bytes:          []byte("payload"),
		FileTotalParts: &totalParts,
	}); err != nil {
		t.Fatalf("WritePart() error = %v", err)
	}
	if err := manager.SaveObjectCacheRef(ctx, 3003, 1001, 2002); err != nil {
		t.Fatalf("SaveObjectCacheRef() error = %v", err)
	}

	info, err := manager.LoadObjectCacheRef(ctx, 3003)
	if err != nil {
		t.Fatalf("LoadObjectCacheRef() error = %v", err)
	}
	if info.Creator != 1001 || info.FileID != 2002 || info.FileSize() != int64(len("payload")) {
		t.Fatalf("LoadObjectCacheRef() = %+v, want creator/file/file size", info)
	}
}

func TestUploadSessionNilRepositoryReturnsStorageError(t *testing.T) {
	ctx := context.Background()
	manager := NewUploadSessionManager(nil)
	if err := manager.SaveObjectCacheRef(ctx, 3003, 1001, 2002); !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("SaveObjectCacheRef() error = %v, want ErrDfsStorage", err)
	}
	if _, err := manager.LoadObjectCacheRef(ctx, 3003); !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("LoadObjectCacheRef() error = %v, want ErrDfsStorage", err)
	}
}

func TestDfsWriteFilePartDataWritesUploadSessionAndReturnsBoolTrue(t *testing.T) {
	ctx := context.Background()
	repo := newFakeUploadStateRepository()
	core := &DfsCore{
		ctx:                  ctx,
		uploadSessionManager: NewUploadSessionManager(repo),
	}
	totalParts := int32(1)

	got, err := core.DfsWriteFilePartData(&dfs.TLDfsWriteFilePartData{
		Creator:        1001,
		FileId:         2002,
		FilePart:       0,
		Bytes:          []byte("payload"),
		Big:            true,
		FileTotalParts: &totalParts,
	})
	if err != nil {
		t.Fatalf("DfsWriteFilePartData() error = %v", err)
	}
	if got != tg.BoolTrue {
		t.Fatalf("DfsWriteFilePartData() = %v, want BoolTrue", got)
	}
	info, err := repo.LoadUploadFileInfo(ctx, 1001, 2002)
	if err != nil {
		t.Fatalf("LoadUploadFileInfo() error = %v", err)
	}
	if !info.Big || info.FileTotalParts != 1 || info.FirstFilePartSize != len("payload") {
		t.Fatalf("written info = %+v, want mapped generated fields", info)
	}
}
