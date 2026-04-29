package xkv

import (
	"context"
	"io"
	"testing"

	"github.com/teamgram/marmota/pkg/stores/kv"
)

type fakeUploadStateExtStore struct {
	kv.ExtStore
	hashes  map[string]map[string]string
	values  map[string]string
	expires map[string]int
}

func newFakeUploadStateExtStore() *fakeUploadStateExtStore {
	return &fakeUploadStateExtStore{
		hashes:  make(map[string]map[string]string),
		values:  make(map[string]string),
		expires: make(map[string]int),
	}
}

func (f *fakeUploadStateExtStore) HsetCtx(_ context.Context, key, field, value string) error {
	if f.hashes[key] == nil {
		f.hashes[key] = make(map[string]string)
	}
	f.hashes[key][field] = value
	return nil
}

func (f *fakeUploadStateExtStore) HgetCtx(_ context.Context, key, field string) (string, error) {
	if f.hashes[key] == nil {
		return "", nil
	}
	return f.hashes[key][field], nil
}

func (f *fakeUploadStateExtStore) HgetallCtx(_ context.Context, key string) (map[string]string, error) {
	src := f.hashes[key]
	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst, nil
}

func (f *fakeUploadStateExtStore) HlenCtx(_ context.Context, key string) (int, error) {
	return len(f.hashes[key]), nil
}

func (f *fakeUploadStateExtStore) ExpireCtx(_ context.Context, key string, seconds int) error {
	f.expires[key] = seconds
	return nil
}

func (f *fakeUploadStateExtStore) SetexCtx(_ context.Context, key, value string, seconds int) error {
	f.values[key] = value
	f.expires[key] = seconds
	return nil
}

func (f *fakeUploadStateExtStore) GetCtx(_ context.Context, key string) (string, error) {
	return f.values[key], nil
}

func TestUploadStateModelPersistsMasterCompatibleKeysAndTTL(t *testing.T) {
	ctx := context.Background()
	store := newFakeUploadStateExtStore()
	model := NewUploadStateModel(store)

	if err := model.SaveUploadPart(ctx, 1001, 2002, 0, []byte("first")); err != nil {
		t.Fatalf("SaveUploadPart() error = %v", err)
	}
	if got := store.hashes["file_1001_2002"]["0"]; got != "first" {
		t.Fatalf("part hash value = %q, want %q", got, "first")
	}
	if got := store.expires["file_1001_2002"]; got != uploadStateTTLSeconds {
		t.Fatalf("part ttl = %d, want %d", got, uploadStateTTLSeconds)
	}

	if err := model.SaveUploadFileInfo(ctx, &DfsFileInfo{
		Creator:           1001,
		FileID:            2002,
		Big:               true,
		FileName:          "photo.jpg",
		FileTotalParts:    3,
		FirstFilePartSize: 5,
		FilePartSize:      6,
		LastFilePartSize:  4,
		MimeType:          "image/jpeg",
		Mtime:             1_700_000_000,
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	infoHash := store.hashes["file_info_1001_2002"]
	for _, field := range []string{"big", "file_name", "file_total_parts", "first_file_part_size", "file_part_size", "last_file_part_size", "mtime"} {
		if infoHash[field] == "" {
			t.Fatalf("file info field %q not persisted", field)
		}
	}
	if _, ok := infoHash["mime_type"]; ok {
		t.Fatal("mime_type was persisted, want master ToArgs behavior")
	}
	if got := store.expires["file_info_1001_2002"]; got != uploadStateTTLSeconds {
		t.Fatalf("info ttl = %d, want %d", got, uploadStateTTLSeconds)
	}
}

func TestUploadStateModelReadsPartsInOrder(t *testing.T) {
	ctx := context.Background()
	store := newFakeUploadStateExtStore()
	model := NewUploadStateModel(store)

	for i, data := range []string{"aa", "bb", "cc"} {
		if err := model.SaveUploadPart(ctx, 1001, 2002, int32(i), []byte(data)); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	reader, err := model.OpenUploadFileReader(ctx, &DfsFileInfo{
		Creator:        1001,
		FileID:         2002,
		FileTotalParts: 3,
	})
	if err != nil {
		t.Fatalf("OpenUploadFileReader() error = %v", err)
	}
	got, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}
	if string(got) != "aabbcc" {
		t.Fatalf("OpenUploadFileReader() bytes = %q, want %q", got, "aabbcc")
	}
}

func TestUploadStateModelDerivesTotalPartsWhenMissing(t *testing.T) {
	ctx := context.Background()
	store := newFakeUploadStateExtStore()
	model := NewUploadStateModel(store)

	for i, data := range []string{"aa", "bbb"} {
		if err := model.SaveUploadPart(ctx, 1001, 2002, int32(i), []byte(data)); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	if err := model.SaveUploadFileInfo(ctx, &DfsFileInfo{
		Creator:           1001,
		FileID:            2002,
		FirstFilePartSize: 2,
		Mtime:             1_700_000_000,
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}

	info, err := model.LoadUploadFileInfo(ctx, 1001, 2002)
	if err != nil {
		t.Fatalf("LoadUploadFileInfo() error = %v", err)
	}
	if info.FileTotalParts != 2 {
		t.Fatalf("FileTotalParts = %d, want 2", info.FileTotalParts)
	}
	if info.LastFilePartSize != 3 {
		t.Fatalf("LastFilePartSize = %d, want 3", info.LastFilePartSize)
	}
}

func TestUploadStateModelCacheRefRoundTripUsesMasterEncoding(t *testing.T) {
	ctx := context.Background()
	store := newFakeUploadStateExtStore()
	model := NewUploadStateModel(store)

	if err := model.SaveUploadFileInfo(ctx, &DfsFileInfo{
		Creator:           1001,
		FileID:            2002,
		FileTotalParts:    1,
		FirstFilePartSize: 7,
		Mtime:             1_700_000_000,
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	if err := model.SaveObjectCacheRef(ctx, 3003, 1001, 2002); err != nil {
		t.Fatalf("SaveObjectCacheRef() error = %v", err)
	}
	if got := store.values["cache_file_info_3003"]; got != "1001_2002" {
		t.Fatalf("cache ref value = %q, want %q", got, "1001_2002")
	}
	if got := store.expires["cache_file_info_3003"]; got != uploadStateCacheRefTTLSeconds {
		t.Fatalf("cache ref ttl = %d, want %d", got, uploadStateCacheRefTTLSeconds)
	}

	info, err := model.LoadObjectCacheRef(ctx, 3003)
	if err != nil {
		t.Fatalf("LoadObjectCacheRef() error = %v", err)
	}
	if info.Creator != 1001 || info.FileID != 2002 || info.FirstFilePartSize != 7 {
		t.Fatalf("LoadObjectCacheRef() = %+v, want creator/file/part size", info)
	}
}
