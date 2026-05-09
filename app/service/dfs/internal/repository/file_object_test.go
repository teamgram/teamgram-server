package repository

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/objectstore"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/xkv"
	"github.com/teamgram/teamgram-server/v2/pkg/filelease"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestPutInternalFilePersistsManifestAndSupportsLeaseRead(t *testing.T) {
	store := newFakeObjectStore()
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9001})

	out, err := repo.PutInternalFile(context.Background(), 1001, "media_derivative", "avatar.png", "image/png", []byte("abcdef"))
	if err != nil {
		t.Fatalf("PutInternalFile() error = %v", err)
	}
	if out.ObjectId != "9001" || out.Bucket != "documents" || out.Key != "objects/9001.dat" || out.Size2 != 6 || out.MimeType != "image/png" || out.DcId != 2 {
		t.Fatalf("PutInternalFile() = %#v", out)
	}
	sum := sha256.Sum256([]byte("abcdef"))
	if !bytes.Equal(out.Sha256, sum[:]) {
		t.Fatalf("sha256 mismatch")
	}
	if len(out.ReadLease) == 0 {
		t.Fatal("expected read lease")
	}

	var manifest objectstore.ObjectManifest
	mustLoadJSON(t, store, "documents", "_meta/objects/9001.json", &manifest)
	if manifest.ObjectID != "9001" || manifest.StorageType != storageTypeID(tg.ClazzID_storage_filePng) || len(manifest.Chunks) != 1 {
		t.Fatalf("object manifest = %#v", manifest)
	}

	got, storageType, err := repo.ReadByLease(context.Background(), out.ReadLease, 2, 3)
	if err != nil {
		t.Fatalf("ReadByLease() error = %v", err)
	}
	if string(got) != "cde" || storageType != storageTypeID(tg.ClazzID_storage_filePng) {
		t.Fatalf("ReadByLease() got %q type %#x", got, storageType)
	}

	chunks, err := repo.HashesByLease(context.Background(), out.ReadLease, 0, 0)
	if err != nil {
		t.Fatalf("HashesByLease() error = %v", err)
	}
	if !reflect.DeepEqual(chunks, manifest.Chunks) {
		t.Fatalf("HashesByLease() = %#v, want %#v", chunks, manifest.Chunks)
	}
}

func TestHashesByLeaseReloadsPersistedHashManifestAfterRestart(t *testing.T) {
	store := newFakeObjectStore()
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9005})

	out, err := repo.PutInternalFile(context.Background(), 1001, "media_derivative", "avatar.png", "image/png", []byte("abcdef"))
	if err != nil {
		t.Fatalf("PutInternalFile() error = %v", err)
	}
	var manifest objectstore.ObjectManifest
	mustLoadJSON(t, store, "documents", "_meta/objects/9005.json", &manifest)
	wantChunks := append([]objectstore.HashChunk(nil), manifest.Chunks...)
	manifest.Chunks = nil
	mustStoreJSON(t, store, "documents", "_meta/objects/9005.json", manifest)

	restarted := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9999})
	got, err := restarted.HashesByLease(context.Background(), out.ReadLease, 0, 0)
	if err != nil {
		t.Fatalf("HashesByLease() error = %v", err)
	}
	if !reflect.DeepEqual(got, wantChunks) {
		t.Fatalf("HashesByLease() = %#v, want persisted chunks %#v", got, wantChunks)
	}
}

func TestCommitUploadIsIdempotentForUploadSession(t *testing.T) {
	store := newFakeObjectStore()
	uploads := &fakeUploadStateModel{
		info: &xkv.DfsFileInfo{
			Creator:           1001,
			FileID:            7001,
			FileName:          "clip.mp4",
			FileTotalParts:    1,
			FirstFilePartSize: 5,
			MimeType:          "video/mp4",
		},
		data: []byte("hello"),
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9002})
	repo.uploadStateModel = uploads

	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 1, Name: "clip.mp4"})
	first, err := repo.CommitUpload(context.Background(), "upload-1", 1001, file, "media_original")
	if err != nil {
		t.Fatalf("CommitUpload() error = %v", err)
	}
	second, err := repo.CommitUpload(context.Background(), "upload-1", 1001, file, "media_original")
	if err != nil {
		t.Fatalf("CommitUpload() second error = %v", err)
	}
	if first.ObjectId != second.ObjectId || first.Key != second.Key || !bytes.Equal(first.Sha256, second.Sha256) {
		t.Fatalf("idempotent result mismatch: first=%#v second=%#v", first, second)
	}
	wantObjectID := expectedCommitUploadObjectID("upload-1")
	wantObjectKey := "documents/objects/" + wantObjectID + ".dat"
	if store.putCounts[wantObjectKey] != 1 {
		t.Fatalf("finalized object put count = %d, want 1", store.putCounts[wantObjectKey])
	}
	if store.putCounts["documents/_meta/uploads/upload-1.json"] != 1 {
		t.Fatalf("upload manifest put count = %d, want 1", store.putCounts["documents/_meta/uploads/upload-1.json"])
	}
	if len(second.ReadLease) == 0 {
		t.Fatal("expected fresh read lease")
	}

	if first.ObjectId != wantObjectID {
		t.Fatalf("object id = %q, want %q", first.ObjectId, wantObjectID)
	}

	var manifest objectstore.ObjectManifest
	mustLoadJSON(t, store, "documents", "_meta/uploads/upload-1.json", &manifest)
	if manifest.OwnerID != 1001 || manifest.FileID != 7001 || manifest.Purpose != "media_original" || manifest.FileName != "clip.mp4" {
		t.Fatalf("upload manifest binding = %#v", manifest)
	}
}

func TestCommitUploadValidatesSmallFileMD5(t *testing.T) {
	data := []byte("small-file")
	sum := md5.Sum(data)
	repo := newCommitUploadTestRepository(newFakeObjectStore(), &fakeIDGenerator{id: 9008})
	repo.uploadStateModel.(*fakeUploadStateModel).data = data
	file := tg.MakeTLInputFile(&tg.TLInputFile{
		Id:          7001,
		Parts:       1,
		Name:        "small.bin",
		Md5Checksum: fmt.Sprintf("%x", sum),
	})

	out, err := repo.CommitUpload(context.Background(), "upload-small-md5", 1001, file, "media_original")
	if err != nil {
		t.Fatalf("CommitUpload() error = %v", err)
	}
	if out.ObjectId != expectedCommitUploadObjectID("upload-small-md5") {
		t.Fatalf("object id = %q", out.ObjectId)
	}
}

func TestCommitUploadRejectsSmallFileMD5Mismatch(t *testing.T) {
	repo := newCommitUploadTestRepository(newFakeObjectStore(), &fakeIDGenerator{id: 9009})
	file := tg.MakeTLInputFile(&tg.TLInputFile{
		Id:          7001,
		Parts:       1,
		Name:        "small.bin",
		Md5Checksum: "00000000000000000000000000000000",
	})

	_, err := repo.CommitUpload(context.Background(), "upload-small-md5-bad", 1001, file, "media_original")
	if !errors.Is(err, dfs.ErrDfsChecksumInvalid) {
		t.Fatalf("CommitUpload() error = %v, want ErrDfsChecksumInvalid", err)
	}
}

func TestCommitUploadBigFileDoesNotRequireMD5(t *testing.T) {
	repo := newCommitUploadTestRepository(newFakeObjectStore(), &fakeIDGenerator{id: 9013})
	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 1, Name: "big.bin"})

	if _, err := repo.CommitUpload(context.Background(), "upload-big-no-md5", 1001, file, "media_original"); err != nil {
		t.Fatalf("CommitUpload() error = %v", err)
	}
}

func TestCommitUploadRetryAfterUploadManifestFailureReusesObjectID(t *testing.T) {
	store := newFakeObjectStore()
	store.failOnceKeys["documents/_meta/uploads/retry-window.json"] = errors.New("forced upload manifest failure")
	repo := newCommitUploadTestRepository(store, &fakeIDGenerator{id: 9014})
	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 1, Name: "clip.mp4"})

	_, err := repo.CommitUpload(context.Background(), "retry-window", 1001, file, "media_original")
	if !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("CommitUpload() first error = %v, want ErrDfsStorage", err)
	}

	out, err := repo.CommitUpload(context.Background(), "retry-window", 1001, file, "media_original")
	if err != nil {
		t.Fatalf("CommitUpload() retry error = %v", err)
	}
	wantObjectID := expectedCommitUploadObjectID("retry-window")
	if out.ObjectId != wantObjectID || out.Key != "objects/"+wantObjectID+".dat" {
		t.Fatalf("retry result = %#v, want object %q", out, wantObjectID)
	}

	var objectKeys []string
	for key := range store.objects {
		if strings.HasPrefix(key, "documents/objects/") {
			objectKeys = append(objectKeys, key)
		}
	}
	if len(objectKeys) != 1 || objectKeys[0] != "documents/objects/"+wantObjectID+".dat" {
		t.Fatalf("object keys = %#v, want only deterministic retry key", objectKeys)
	}
	if store.putCounts["documents/objects/"+wantObjectID+".dat"] != 2 {
		t.Fatalf("object put count = %d, want retry overwrite count 2", store.putCounts["documents/objects/"+wantObjectID+".dat"])
	}
}

func TestCommitUploadRejectsExistingManifestOwnerMismatch(t *testing.T) {
	store := newFakeObjectStore()
	repo := newCommitUploadTestRepository(store, &fakeIDGenerator{id: 9010})
	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 1, Name: "clip.mp4"})

	if _, err := repo.CommitUpload(context.Background(), "upload-owner-mismatch", 1001, file, "media_original"); err != nil {
		t.Fatalf("CommitUpload() error = %v", err)
	}
	_, err := repo.CommitUpload(context.Background(), "upload-owner-mismatch", 2002, file, "media_original")
	if !errors.Is(err, dfs.ErrDfsInvalidArgument) {
		t.Fatalf("CommitUpload() owner mismatch error = %v, want ErrDfsInvalidArgument", err)
	}
	wantObjectKey := "documents/objects/" + expectedCommitUploadObjectID("upload-owner-mismatch") + ".dat"
	if store.putCounts[wantObjectKey] != 1 {
		t.Fatalf("finalized object put count = %d, want 1", store.putCounts[wantObjectKey])
	}
}

func TestCommitUploadRejectsExistingManifestFileIDMismatch(t *testing.T) {
	store := newFakeObjectStore()
	repo := newCommitUploadTestRepository(store, &fakeIDGenerator{id: 9011})
	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 1, Name: "clip.mp4"})

	if _, err := repo.CommitUpload(context.Background(), "upload-file-mismatch", 1001, file, "media_original"); err != nil {
		t.Fatalf("CommitUpload() error = %v", err)
	}
	otherFile := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7002, Parts: 1, Name: "clip.mp4"})
	_, err := repo.CommitUpload(context.Background(), "upload-file-mismatch", 1001, otherFile, "media_original")
	if !errors.Is(err, dfs.ErrDfsInvalidArgument) {
		t.Fatalf("CommitUpload() file mismatch error = %v, want ErrDfsInvalidArgument", err)
	}
	wantObjectKey := "documents/objects/" + expectedCommitUploadObjectID("upload-file-mismatch") + ".dat"
	if store.putCounts[wantObjectKey] != 1 {
		t.Fatalf("finalized object put count = %d, want 1", store.putCounts[wantObjectKey])
	}
}

func TestCommitUploadRejectsExistingManifestPurposeMismatch(t *testing.T) {
	store := newFakeObjectStore()
	repo := newCommitUploadTestRepository(store, &fakeIDGenerator{id: 9012})
	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 1, Name: "clip.mp4"})

	if _, err := repo.CommitUpload(context.Background(), "upload-purpose-mismatch", 1001, file, "media_original"); err != nil {
		t.Fatalf("CommitUpload() error = %v", err)
	}
	_, err := repo.CommitUpload(context.Background(), "upload-purpose-mismatch", 1001, file, "media_derivative")
	if !errors.Is(err, dfs.ErrDfsInvalidArgument) {
		t.Fatalf("CommitUpload() purpose mismatch error = %v, want ErrDfsInvalidArgument", err)
	}
	wantObjectKey := "documents/objects/" + expectedCommitUploadObjectID("upload-purpose-mismatch") + ".dat"
	if store.putCounts[wantObjectKey] != 1 {
		t.Fatalf("finalized object put count = %d, want 1", store.putCounts[wantObjectKey])
	}
}

func TestReadByLeaseMapsInvalidAndExpiredTokenToInvalidArgument(t *testing.T) {
	repo := newFileObjectTestRepository(newFakeObjectStore(), &fakeIDGenerator{id: 9006})
	out, err := repo.PutInternalFile(context.Background(), 1001, "media_derivative", "x.bin", "", []byte("abc"))
	if err != nil {
		t.Fatalf("PutInternalFile() error = %v", err)
	}
	tampered := append([]byte(nil), out.ReadLease...)
	tampered[len(tampered)-1] ^= 0xff
	if _, _, err := repo.ReadByLease(context.Background(), tampered, 0, 0); !errors.Is(err, dfs.ErrDfsInvalidArgument) || !errors.Is(err, filelease.ErrInvalidToken) {
		t.Fatalf("ReadByLease() tampered error = %v, want ErrDfsInvalidArgument and ErrInvalidToken", err)
	}

	expired, err := filelease.Sign(repo.readLeaseSecret, filelease.Claims{
		ObjectID:  "9006",
		Bucket:    "documents",
		Key:       "objects/9006.dat",
		ExpiresAt: time.Now().Add(-time.Second).Unix(),
	})
	if err != nil {
		t.Fatalf("Sign() expired error = %v", err)
	}
	if _, _, err := repo.ReadByLease(context.Background(), expired, 0, 0); !errors.Is(err, dfs.ErrDfsInvalidArgument) || !errors.Is(err, filelease.ErrExpired) {
		t.Fatalf("ReadByLease() expired error = %v, want ErrDfsInvalidArgument and ErrExpired", err)
	}
}

func TestSignReadLeaseMapsEmptySecretToStorage(t *testing.T) {
	repo := newFileObjectTestRepository(newFakeObjectStore(), &fakeIDGenerator{id: 9007})
	repo.readLeaseSecret = ""

	_, err := repo.signReadLease(objectstore.ObjectManifest{ObjectID: "9007", Bucket: "documents", Key: "objects/9007.dat"})
	if !errors.Is(err, dfs.ErrDfsStorage) || !errors.Is(err, filelease.ErrEmptySecret) {
		t.Fatalf("signReadLease() error = %v, want ErrDfsStorage and ErrEmptySecret", err)
	}
}

func TestCommitUploadRejectsOwnerMismatch(t *testing.T) {
	repo := newFileObjectTestRepository(newFakeObjectStore(), &fakeIDGenerator{id: 9003})
	repo.uploadStateModel = &fakeUploadStateModel{
		info: &xkv.DfsFileInfo{Creator: 1002, FileID: 7001, FileTotalParts: 1, FirstFilePartSize: 5},
		data: []byte("hello"),
	}
	file := tg.MakeTLInputFile(&tg.TLInputFile{Id: 7001, Parts: 1, Name: "x.jpg", Md5Checksum: "ignored"})

	_, err := repo.CommitUpload(context.Background(), "upload-2", 1001, file, "media_original")
	if !errors.Is(err, dfs.ErrDfsInvalidArgument) {
		t.Fatalf("CommitUpload() error = %v, want ErrDfsInvalidArgument", err)
	}
}

func TestReadByLeaseMapsObjectMissToFileNotFound(t *testing.T) {
	store := newFakeObjectStore()
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9004})
	out, err := repo.PutInternalFile(context.Background(), 1001, "media_derivative", "x.bin", "", []byte("abc"))
	if err != nil {
		t.Fatalf("PutInternalFile() error = %v", err)
	}
	delete(store.objects, "documents/objects/9004.dat")

	_, _, err = repo.ReadByLease(context.Background(), out.ReadLease, 0, 0)
	if !errors.Is(err, dfs.ErrDfsFileNotFound) {
		t.Fatalf("ReadByLease() error = %v, want ErrDfsFileNotFound", err)
	}
	if errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("ReadByLease() error = %v, should not be ErrDfsStorage", err)
	}
}

type fakeIDGenerator struct {
	id  int64
	err error
}

func (f *fakeIDGenerator) NextPhotoID(context.Context) (int64, error) { return 0, nil }
func (f *fakeIDGenerator) NextEncryptedFileID(context.Context) (int64, error) {
	return 0, nil
}
func (f *fakeIDGenerator) NextDocumentID(context.Context) (int64, error) {
	if f.err != nil {
		return 0, f.err
	}
	return f.id, nil
}

type fakeObjectStore struct {
	objects      map[string][]byte
	putCounts    map[string]int
	failOnceKeys map[string]error
}

func newFakeObjectStore() *fakeObjectStore {
	return &fakeObjectStore{
		objects:      make(map[string][]byte),
		putCounts:    make(map[string]int),
		failOnceKeys: make(map[string]error),
	}
}

func (f *fakeObjectStore) PutObjectReader(_ context.Context, bucket, key string, r io.Reader) (int64, error) {
	fullKey := bucket + "/" + key
	if err := f.failOnceKeys[fullKey]; err != nil {
		delete(f.failOnceKeys, fullKey)
		return 0, err
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	f.objects[fullKey] = append([]byte(nil), data...)
	f.putCounts[fullKey]++
	return int64(len(data)), nil
}

func (f *fakeObjectStore) PutObjectBytes(ctx context.Context, bucket, key string, data []byte) (int64, error) {
	return f.PutObjectReader(ctx, bucket, key, bytes.NewReader(data))
}

func (f *fakeObjectStore) GetObjectRange(_ context.Context, bucket, key string, offset int64, limit int32) ([]byte, error) {
	if offset < 0 || limit < 0 {
		return nil, errors.New("invalid range")
	}
	data, ok := f.objects[bucket+"/"+key]
	if !ok {
		return nil, minioadapter.ErrObjectNotFound
	}
	if offset >= int64(len(data)) {
		return []byte{}, nil
	}
	end := int64(len(data))
	if limit > 0 && offset+int64(limit) < end {
		end = offset + int64(limit)
	}
	return append([]byte(nil), data[offset:end]...), nil
}

func (f *fakeObjectStore) PutPhotoBytes(context.Context, string, []byte) (int64, error) {
	return 0, nil
}
func (f *fakeObjectStore) PutPhotoReader(context.Context, string, io.Reader) (int64, error) {
	return 0, nil
}
func (f *fakeObjectStore) GetPhotoFile(context.Context, string, int64, int32) ([]byte, error) {
	return nil, nil
}
func (f *fakeObjectStore) PutVideoBytes(context.Context, string, []byte) (int64, error) {
	return 0, nil
}
func (f *fakeObjectStore) GetVideoFile(context.Context, string, int64, int32) ([]byte, error) {
	return nil, nil
}
func (f *fakeObjectStore) PutDocumentReader(context.Context, string, io.Reader) (int64, error) {
	return 0, nil
}
func (f *fakeObjectStore) GetDocumentFile(context.Context, string, int64, int32) ([]byte, error) {
	return nil, nil
}
func (f *fakeObjectStore) PutEncryptedFileReader(context.Context, string, io.Reader) (int64, error) {
	return 0, nil
}
func (f *fakeObjectStore) GetEncryptedFile(context.Context, string, int64, int32) ([]byte, error) {
	return nil, nil
}

type fakeUploadStateModel struct {
	info *xkv.DfsFileInfo
	data []byte
}

func (f *fakeUploadStateModel) SaveUploadPart(context.Context, int64, int64, int32, []byte) error {
	return nil
}
func (f *fakeUploadStateModel) SaveUploadFileInfo(context.Context, *xkv.DfsFileInfo) error {
	return nil
}
func (f *fakeUploadStateModel) LoadUploadFileInfo(_ context.Context, creator, fileID int64) (*xkv.DfsFileInfo, error) {
	if f.info == nil || f.info.FileID != fileID {
		return nil, xkv.ErrUploadStateNotFound
	}
	return f.info, nil
}
func (f *fakeUploadStateModel) OpenUploadFileReader(context.Context, *xkv.DfsFileInfo) (io.ReadSeeker, error) {
	return bytes.NewReader(f.data), nil
}
func (f *fakeUploadStateModel) SaveObjectCacheRef(context.Context, int64, int64, int64) error {
	return nil
}
func (f *fakeUploadStateModel) LoadObjectCacheRef(context.Context, int64) (*xkv.DfsFileInfo, error) {
	return nil, xkv.ErrUploadStateNotFound
}

func newFileObjectTestRepository(store *fakeObjectStore, idgen *fakeIDGenerator) *Repository {
	return &Repository{
		objectStore:         store,
		idgen:               idgen,
		documentsBucket:     "documents",
		manifestKeys:        objectstore.ManifestKeys{MetaPrefix: "_meta"},
		readLeaseSecret:     "test-secret",
		readLeaseTTLSeconds: int64(time.Hour / time.Second),
		localDCID:           2,
	}
}

func newCommitUploadTestRepository(store *fakeObjectStore, idgen *fakeIDGenerator) *Repository {
	repo := newFileObjectTestRepository(store, idgen)
	repo.uploadStateModel = &fakeUploadStateModel{
		info: &xkv.DfsFileInfo{
			Creator:           1001,
			FileID:            7001,
			FileName:          "clip.mp4",
			FileTotalParts:    1,
			FirstFilePartSize: 5,
			MimeType:          "video/mp4",
		},
		data: []byte("hello"),
	}
	return repo
}

func mustLoadJSON(t *testing.T, store *fakeObjectStore, bucket, key string, out any) {
	t.Helper()
	data, ok := store.objects[bucket+"/"+key]
	if !ok {
		t.Fatalf("missing object %s/%s", bucket, key)
	}
	if err := json.Unmarshal(data, out); err != nil {
		t.Fatalf("json unmarshal %s/%s: %v", bucket, key, err)
	}
}

func mustStoreJSON(t *testing.T, store *fakeObjectStore, bucket, key string, v any) {
	t.Helper()
	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("json marshal %s/%s: %v", bucket, key, err)
	}
	store.objects[bucket+"/"+key] = data
}

func expectedCommitUploadObjectID(uploadSessionID string) string {
	sum := sha256.Sum256([]byte(uploadSessionID))
	return fmt.Sprintf("upload-%x", sum)
}
