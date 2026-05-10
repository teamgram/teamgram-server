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
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/config"
	minioadapter "github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/minio"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/objectstore"
	"github.com/teamgram/teamgram-server/v2/app/service/dfs/internal/repository/spool"
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

func TestReadByLeaseUsesLeaseSizeForEOFAndRangeClamp(t *testing.T) {
	store := newFakeObjectStore()
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9010})

	out, err := repo.PutInternalFile(context.Background(), 1001, "media_derivative", "video.mp4", "video/mp4", []byte("abcdef"))
	if err != nil {
		t.Fatalf("PutInternalFile() error = %v", err)
	}

	store.rangeRequests = nil
	got, storageType, err := repo.ReadByLease(context.Background(), out.ReadLease, out.Size2, 4)
	if err != nil {
		t.Fatalf("ReadByLease() EOF error = %v", err)
	}
	if len(got) != 0 || storageType != storageTypeID(tg.ClazzID_storage_fileMp4) {
		t.Fatalf("ReadByLease() EOF got %q type %#x", got, storageType)
	}
	if len(store.rangeRequests) != 0 {
		t.Fatalf("object store range calls = %d, want 0 at EOF", len(store.rangeRequests))
	}

	got, _, err = repo.ReadByLease(context.Background(), out.ReadLease, 4, 99)
	if err != nil {
		t.Fatalf("ReadByLease() clamp error = %v", err)
	}
	if string(got) != "ef" {
		t.Fatalf("ReadByLease() clamp got %q, want ef", got)
	}
	if len(store.rangeRequests) != 1 || store.rangeRequests[0].Offset != 4 || store.rangeRequests[0].Limit != 2 {
		t.Fatalf("object store range request = %#v, want offset 4 limit 2", store.rangeRequests)
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
	if store.putCounts["documents/objects/"+wantObjectID+".dat"] != 1 {
		t.Fatalf("object put count = %d, want retry to reuse existing deterministic object", store.putCounts["documents/objects/"+wantObjectID+".dat"])
	}
}

func TestCommitUploadLocalSpoolUsesMultipartSegmentsAndPersistsManifests(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9020})
	repo.uploadStateModel = model

	for i, data := range [][]byte{[]byte("aa"), []byte("bb"), []byte("cc")} {
		if err := model.SaveUploadPart(ctx, 1001, 7001, int32(i), data); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	info := &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileName:          "clip.mp4",
		FileTotalParts:    3,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		MimeType:          "video/mp4",
	}
	if err := model.SaveUploadFileInfo(ctx, info); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}

	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 3, Name: "clip.mp4"})
	out, err := repo.CommitUpload(ctx, "upload-segmented", 1001, file, "media_original")
	if err != nil {
		t.Fatalf("CommitUpload() error = %v", err)
	}
	wantObjectID := expectedCommitUploadObjectID("upload-segmented")
	if out.ObjectId != wantObjectID || out.Key != "objects/"+wantObjectID+".dat" || out.Size2 != 6 {
		t.Fatalf("CommitUpload() = %#v", out)
	}
	if got := string(store.objects["documents/objects/"+wantObjectID+".dat"]); got != "aabbcc" {
		t.Fatalf("object bytes = %q, want aabbcc", got)
	}
	if len(store.multipartBegins) != 1 || len(store.multipartUploads) != 2 || len(store.multipartCompletes) != 1 {
		t.Fatalf("multipart calls begins=%d uploads=%d completes=%d, want 1/2/1", len(store.multipartBegins), len(store.multipartUploads), len(store.multipartCompletes))
	}
	states, err := model.LoadSegmentStates(ctx, info)
	if err != nil {
		t.Fatalf("LoadSegmentStates() error = %v", err)
	}
	if len(states) != 2 || states[0].Status != spool.SegmentStatusUploaded || states[1].Status != spool.SegmentStatusUploaded {
		t.Fatalf("segment states = %#v, want two uploaded states", states)
	}

	var manifest objectstore.ObjectManifest
	mustLoadJSON(t, store, "documents", "_meta/uploads/upload-segmented.json", &manifest)
	if manifest.ObjectID != wantObjectID || manifest.Size != 6 || len(manifest.Chunks) == 0 {
		t.Fatalf("upload manifest = %#v", manifest)
	}
	if _, _, err := repo.ReadByLease(ctx, out.ReadLease, 0, 0); err != nil {
		t.Fatalf("ReadByLease() error = %v", err)
	}
}

func TestCommitUploadLocalSpoolRetryRebuildsDigestForUploadedSegments(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9023})
	repo.uploadStateModel = model

	for i, data := range [][]byte{[]byte("aa"), []byte("bb"), []byte("cc"), []byte("dd")} {
		if err := model.SaveUploadPart(ctx, 1001, 7001, int32(i), data); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	info := &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileName:          "clip.mp4",
		FileTotalParts:    4,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		MimeType:          "video/mp4",
	}
	if err := model.SaveUploadFileInfo(ctx, info); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	uploadID := "mp-retry"
	objectID := expectedCommitUploadObjectID("upload-segmented-retry")
	objectKey := "objects/" + objectID + ".dat"
	store.multipartParts[uploadID] = map[int][]byte{1: []byte("aabb")}
	store.multipartPartETags[uploadID] = map[int]string{1: "etag-mp-retry-1"}
	if err := model.WriteSegmentState(ctx, info, spool.SegmentState{
		SegmentNo:           0,
		Status:              spool.SegmentStatusUploaded,
		MultipartUploadID:   uploadID,
		MultipartPartNumber: 1,
		ObjectKey:           objectKey,
		ETag:                "etag-mp-retry-1",
		Checksum:            fmt.Sprintf("%x", sha256.Sum256([]byte("aabb"))),
		Size:                4,
	}); err != nil {
		t.Fatalf("WriteSegmentState(uploaded) error = %v", err)
	}
	if err := model.WriteSegmentState(ctx, info, spool.SegmentState{
		SegmentNo:           1,
		Status:              spool.SegmentStatusUploading,
		MultipartUploadID:   uploadID,
		MultipartPartNumber: 2,
		ObjectKey:           objectKey,
		Size:                4,
	}); err != nil {
		t.Fatalf("WriteSegmentState(uploading) error = %v", err)
	}

	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 4, Name: "clip.mp4"})
	out, err := repo.CommitUpload(ctx, "upload-segmented-retry", 1001, file, "media_original")
	if err != nil {
		t.Fatalf("CommitUpload() error = %v", err)
	}
	if out.Size2 != 8 {
		t.Fatalf("finalized size = %d, want 8", out.Size2)
	}
	if got := string(store.objects["documents/"+objectKey]); got != "aabbccdd" {
		t.Fatalf("object bytes = %q, want aabbccdd", got)
	}
	var manifest objectstore.ObjectManifest
	mustLoadJSON(t, store, "documents", "_meta/uploads/upload-segmented-retry.json", &manifest)
	wantSHA := sha256.Sum256([]byte("aabbccdd"))
	if manifest.Size != 8 || !bytes.Equal(manifest.SHA256, wantSHA[:]) {
		t.Fatalf("manifest size/sha = %d/%x, want 8/%x", manifest.Size, manifest.SHA256, wantSHA)
	}
	if len(manifest.Chunks) != 1 || manifest.Chunks[0].Offset != 0 || manifest.Chunks[0].Limit != 8 {
		t.Fatalf("manifest chunks = %#v, want one full-file chunk", manifest.Chunks)
	}
	if len(store.multipartBegins) != 0 {
		t.Fatalf("multipart begins = %#v, want retry to reuse existing upload id", store.multipartBegins)
	}
	if len(store.multipartUploads) != 1 || !strings.Contains(store.multipartUploads[0], uploadID+":2:") {
		t.Fatalf("multipart uploads = %#v, want only missing segment uploaded with existing upload id", store.multipartUploads)
	}
}

func TestCommitUploadLocalSpoolRetryReuploadsUploadedSegmentWhenClientPartChanged(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9028})
	repo.uploadStateModel = model

	for i, data := range [][]byte{[]byte("aa"), []byte("bb"), []byte("cc"), []byte("dd")} {
		if err := model.SaveUploadPart(ctx, 1001, 7001, int32(i), data); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	info := &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileName:          "clip.mp4",
		FileTotalParts:    4,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		MimeType:          "video/mp4",
	}
	if err := model.SaveUploadFileInfo(ctx, info); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	uploadID := "mp-overwrite"
	objectID := expectedCommitUploadObjectID("upload-overwritten-uploaded-segment")
	objectKey := "objects/" + objectID + ".dat"
	oldChecksum := fmt.Sprintf("%x", sha256.Sum256([]byte("aabb")))
	store.multipartParts[uploadID] = map[int][]byte{1: []byte("aabb")}
	store.multipartPartETags[uploadID] = map[int]string{1: "etag-old-1"}
	if err := model.WriteSegmentState(ctx, info, spool.SegmentState{
		SegmentNo:           0,
		Status:              spool.SegmentStatusUploaded,
		MultipartUploadID:   uploadID,
		MultipartPartNumber: 1,
		ObjectKey:           objectKey,
		ETag:                "etag-old-1",
		Checksum:            oldChecksum,
		Size:                4,
	}); err != nil {
		t.Fatalf("WriteSegmentState(uploaded) error = %v", err)
	}
	if err := model.SaveUploadPart(ctx, 1001, 7001, 0, []byte("xx")); err != nil {
		t.Fatalf("SaveUploadPart(overwrite 0) error = %v", err)
	}
	if err := model.SaveUploadPart(ctx, 1001, 7001, 1, []byte("yy")); err != nil {
		t.Fatalf("SaveUploadPart(overwrite 1) error = %v", err)
	}

	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 4, Name: "clip.mp4"})
	out, err := repo.CommitUpload(ctx, "upload-overwritten-uploaded-segment", 1001, file, "media_original")
	if err != nil {
		t.Fatalf("CommitUpload() error = %v", err)
	}
	if out.Size2 != 8 {
		t.Fatalf("finalized size = %d, want 8", out.Size2)
	}
	if got := string(store.objects["documents/"+objectKey]); got != "xxyyccdd" {
		t.Fatalf("object bytes = %q, want xxyyccdd", got)
	}
	newChecksum := fmt.Sprintf("%x", sha256.Sum256([]byte("xxyy")))
	if len(store.multipartUploads) != 2 ||
		!strings.Contains(store.multipartUploads[0], uploadID+":1:"+newChecksum) ||
		!strings.Contains(store.multipartUploads[1], uploadID+":2:") {
		t.Fatalf("multipart uploads = %#v, want overwritten part 1 and new part 2 uploaded with existing upload id", store.multipartUploads)
	}
	var manifest objectstore.ObjectManifest
	mustLoadJSON(t, store, "documents", "_meta/uploads/upload-overwritten-uploaded-segment.json", &manifest)
	wantSHA := sha256.Sum256([]byte("xxyyccdd"))
	if manifest.Size != 8 || !bytes.Equal(manifest.SHA256, wantSHA[:]) {
		t.Fatalf("manifest size/sha = %d/%x, want 8/%x", manifest.Size, manifest.SHA256, wantSHA)
	}
	state, err := model.LoadSegmentState(ctx, info, 0)
	if err != nil {
		t.Fatalf("LoadSegmentState() error = %v", err)
	}
	if state.Checksum != newChecksum || state.ETag == "etag-old-1" || state.Status != spool.SegmentStatusUploaded {
		t.Fatalf("segment state = %#v, want uploaded with refreshed checksum/etag", state)
	}
}

func TestCommitUploadLocalSpoolRetryAfterObjectManifestFailureRecoversCompletedObject(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9024})
	repo.uploadStateModel = model

	for i, data := range [][]byte{[]byte("aa"), []byte("bb"), []byte("cc")} {
		if err := model.SaveUploadPart(ctx, 1001, 7001, int32(i), data); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	info := &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileName:          "clip.mp4",
		FileTotalParts:    3,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		MimeType:          "video/mp4",
	}
	if err := model.SaveUploadFileInfo(ctx, info); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	objectID := expectedCommitUploadObjectID("upload-manifest-retry")
	store.failOnceKeys["documents/_meta/objects/"+objectID+".json"] = errors.New("forced object manifest failure")
	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 3, Name: "clip.mp4"})

	if _, err := repo.CommitUpload(ctx, "upload-manifest-retry", 1001, file, "media_original"); !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("CommitUpload() first error = %v, want ErrDfsStorage", err)
	}
	if len(store.multipartBegins) != 1 {
		t.Fatalf("multipart begins after first commit = %#v, want one", store.multipartBegins)
	}

	out, err := repo.CommitUpload(ctx, "upload-manifest-retry", 1001, file, "media_original")
	if err != nil {
		t.Fatalf("CommitUpload() retry error = %v", err)
	}
	if out.ObjectId != objectID || out.Size2 != 6 {
		t.Fatalf("CommitUpload() retry = %#v, want object %q size 6", out, objectID)
	}
	if len(store.multipartBegins) != 1 {
		t.Fatalf("multipart begins after retry = %#v, want no new multipart upload", store.multipartBegins)
	}
	if len(store.multipartCompletes) != 1 {
		t.Fatalf("multipart completes = %#v, want retry to recover completed object without completing again", store.multipartCompletes)
	}
	if got := string(store.objects["documents/objects/"+objectID+".dat"]); got != "aabbcc" {
		t.Fatalf("object bytes = %q, want aabbcc", got)
	}
}

func TestCommitUploadLocalSpoolRetryAfterHashManifestFailureBackfillsHashManifest(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9031})
	repo.uploadStateModel = model

	for i, data := range [][]byte{[]byte("aa"), []byte("bb"), []byte("cc")} {
		if err := model.SaveUploadPart(ctx, 1001, 7001, int32(i), data); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileName:          "clip.mp4",
		FileTotalParts:    3,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		MimeType:          "video/mp4",
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	objectID := expectedCommitUploadObjectID("upload-hash-manifest-retry")
	store.failOnceKeys["documents/_meta/hashes/"+objectID+"/v1.json"] = errors.New("forced hash manifest failure")
	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 3, Name: "clip.mp4"})

	if _, err := repo.CommitUpload(ctx, "upload-hash-manifest-retry", 1001, file, "media_original"); !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("CommitUpload() first error = %v, want ErrDfsStorage", err)
	}
	if len(store.multipartCompletes) != 1 {
		t.Fatalf("multipart completes after first commit = %#v, want one", store.multipartCompletes)
	}

	out, err := repo.CommitUpload(ctx, "upload-hash-manifest-retry", 1001, file, "media_original")
	if err != nil {
		t.Fatalf("CommitUpload() retry error = %v", err)
	}
	if out.ObjectId != objectID || out.Size2 != 6 {
		t.Fatalf("CommitUpload() retry = %#v, want object %q size 6", out, objectID)
	}
	if len(store.multipartCompletes) != 1 {
		t.Fatalf("multipart completes after retry = %#v, want no second complete", store.multipartCompletes)
	}
	var chunks []objectstore.HashChunk
	mustLoadJSON(t, store, "documents", "_meta/hashes/"+objectID+"/v1.json", &chunks)
	if len(chunks) == 0 {
		t.Fatalf("hash manifest chunks = %#v, want backfilled chunks", chunks)
	}
}

func TestCommitUploadLocalSpoolRetryAfterManifestFailureUsesFinalizedSnapshot(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9032})
	repo.uploadStateModel = model

	for i, data := range [][]byte{[]byte("aa"), []byte("bb"), []byte("cc")} {
		if err := model.SaveUploadPart(ctx, 1001, 7001, int32(i), data); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	info := &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileName:          "clip.mp4",
		FileTotalParts:    3,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		MimeType:          "video/mp4",
	}
	if err := model.SaveUploadFileInfo(ctx, info); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	objectID := expectedCommitUploadObjectID("upload-finalizing-snapshot")
	store.failOnceKeys["documents/_meta/objects/"+objectID+".json"] = errors.New("forced object manifest failure")
	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 3, Name: "clip.mp4"})

	if _, err := repo.CommitUpload(ctx, "upload-finalizing-snapshot", 1001, file, "media_original"); !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("CommitUpload() first error = %v, want ErrDfsStorage", err)
	}
	if len(store.multipartCompletes) != 1 {
		t.Fatalf("multipart completes after first commit = %#v, want one", store.multipartCompletes)
	}

	for i, data := range [][]byte{[]byte("xx"), []byte("yy"), []byte("zz")} {
		if err := model.SaveUploadPart(ctx, 1001, 7001, int32(i), data); err != nil {
			t.Fatalf("SaveUploadPart(overwrite %d) error = %v", i, err)
		}
	}

	out, err := repo.CommitUpload(ctx, "upload-finalizing-snapshot", 1001, file, "media_original")
	if err != nil {
		t.Fatalf("CommitUpload() retry error = %v", err)
	}
	if len(store.multipartCompletes) != 1 {
		t.Fatalf("multipart completes after retry = %#v, want no second complete", store.multipartCompletes)
	}
	if got := string(store.objects["documents/objects/"+objectID+".dat"]); got != "aabbcc" {
		t.Fatalf("object bytes = %q, want finalized bytes aabbcc", got)
	}
	wantSHA := sha256.Sum256([]byte("aabbcc"))
	if out.Size2 != 6 || !bytes.Equal(out.Sha256, wantSHA[:]) {
		t.Fatalf("retry result size/sha = %d/%x, want finalized 6/%x", out.Size2, out.Sha256, wantSHA)
	}
	var manifest objectstore.ObjectManifest
	mustLoadJSON(t, store, "documents", "_meta/uploads/upload-finalizing-snapshot.json", &manifest)
	if manifest.Size != 6 || !bytes.Equal(manifest.SHA256, wantSHA[:]) {
		t.Fatalf("upload manifest size/sha = %d/%x, want finalized 6/%x", manifest.Size, manifest.SHA256, wantSHA)
	}
	var chunks []objectstore.HashChunk
	mustLoadJSON(t, store, "documents", "_meta/hashes/"+objectID+"/v1.json", &chunks)
	if len(chunks) != 1 || !bytes.Equal(chunks[0].Hash, wantSHA[:]) {
		t.Fatalf("hash manifest = %#v, want finalized object hash", chunks)
	}
}

func TestCommitUploadLocalSpoolRejectsMultipartObjectChangedAfterComplete(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	store.afterCompleteMultipart = func(objects map[string][]byte, fullKey string) {
		objects[fullKey] = []byte("corrupt")
	}
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9033})
	repo.uploadStateModel = model

	for i, data := range [][]byte{[]byte("aa"), []byte("bb"), []byte("cc")} {
		if err := model.SaveUploadPart(ctx, 1001, 7001, int32(i), data); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileName:          "clip.mp4",
		FileTotalParts:    3,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		MimeType:          "video/mp4",
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 3, Name: "clip.mp4"})

	_, err = repo.CommitUpload(ctx, "upload-corrupt-after-complete", 1001, file, "media_original")
	if !errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("CommitUpload() error = %v, want ErrDfsStorage", err)
	}
	objectID := expectedCommitUploadObjectID("upload-corrupt-after-complete")
	if _, ok := store.objects["documents/_meta/objects/"+objectID+".json"]; ok {
		t.Fatalf("object manifest was written after object verification failure")
	}
	if _, ok := store.objects["documents/_meta/uploads/upload-corrupt-after-complete.json"]; ok {
		t.Fatalf("upload manifest was written after object verification failure")
	}
	if _, ok := store.objects["documents/_meta/hashes/"+objectID+"/v1.json"]; ok {
		t.Fatalf("hash manifest was written after object verification failure")
	}
}

func TestCommitUploadLocalSpoolRejectsSmallMD5MismatchBeforeSinglePut(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9029})
	repo.uploadStateModel = model

	if err := model.SaveUploadPart(ctx, 1001, 7001, 0, []byte("bad")); err != nil {
		t.Fatalf("SaveUploadPart() error = %v", err)
	}
	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileName:          "bad.bin",
		FileTotalParts:    1,
		FirstFilePartSize: 3,
		FilePartSize:      3,
		LastFilePartSize:  3,
		MimeType:          "application/octet-stream",
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}

	file := tg.MakeTLInputFile(&tg.TLInputFile{
		Id:          7001,
		Parts:       1,
		Name:        "bad.bin",
		Md5Checksum: "00000000000000000000000000000000",
	})
	_, err = repo.CommitUpload(ctx, "upload-single-bad-md5", 1001, file, "media_original")
	if !errors.Is(err, dfs.ErrDfsChecksumInvalid) {
		t.Fatalf("CommitUpload() error = %v, want ErrDfsChecksumInvalid", err)
	}
	objectKey := "documents/objects/" + expectedCommitUploadObjectID("upload-single-bad-md5") + ".dat"
	if store.putCounts[objectKey] != 0 {
		t.Fatalf("object put count = %d, want checksum rejection before put", store.putCounts[objectKey])
	}
}

func TestCommitUploadLocalSpoolRejectsSmallMD5MismatchBeforeMultipartBegin(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9030})
	repo.uploadStateModel = model

	for i, data := range [][]byte{[]byte("aa"), []byte("bb"), []byte("cc"), []byte("dd")} {
		if err := model.SaveUploadPart(ctx, 1001, 7001, int32(i), data); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileName:          "bad.bin",
		FileTotalParts:    4,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		MimeType:          "application/octet-stream",
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}

	file := tg.MakeTLInputFile(&tg.TLInputFile{
		Id:          7001,
		Parts:       4,
		Name:        "bad.bin",
		Md5Checksum: "00000000000000000000000000000000",
	})
	_, err = repo.CommitUpload(ctx, "upload-multipart-bad-md5", 1001, file, "media_original")
	if !errors.Is(err, dfs.ErrDfsChecksumInvalid) {
		t.Fatalf("CommitUpload() error = %v, want ErrDfsChecksumInvalid", err)
	}
	if len(store.multipartBegins) != 0 || len(store.multipartUploads) != 0 || len(store.multipartCompletes) != 0 {
		t.Fatalf("multipart calls begins=%#v uploads=%#v completes=%#v, want checksum rejection before remote staging", store.multipartBegins, store.multipartUploads, store.multipartCompletes)
	}
}

func TestCommitUploadLocalSpoolConflictingMultipartUploadIDsReturnsStorage(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9027})
	repo.uploadStateModel = model
	for i, data := range [][]byte{[]byte("aa"), []byte("bb"), []byte("cc")} {
		if err := model.SaveUploadPart(ctx, 1001, 7001, int32(i), data); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	info := &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileName:          "clip.mp4",
		FileTotalParts:    3,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		MimeType:          "video/mp4",
	}
	if err := model.SaveUploadFileInfo(ctx, info); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	objectKey := "objects/" + expectedCommitUploadObjectID("upload-conflict") + ".dat"
	for _, state := range []spool.SegmentState{
		{SegmentNo: 0, Status: spool.SegmentStatusUploaded, MultipartUploadID: "mp-a", MultipartPartNumber: 1, ObjectKey: objectKey, ETag: "etag-a", Size: 4},
		{SegmentNo: 1, Status: spool.SegmentStatusUploading, MultipartUploadID: "mp-b", MultipartPartNumber: 2, ObjectKey: objectKey, Size: 2},
	} {
		if err := model.WriteSegmentState(ctx, info, state); err != nil {
			t.Fatalf("WriteSegmentState(%d) error = %v", state.SegmentNo, err)
		}
	}

	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 3, Name: "clip.mp4"})
	_, err = repo.CommitUpload(ctx, "upload-conflict", 1001, file, "media_original")
	if !errors.Is(err, dfs.ErrDfsStorage) || !strings.Contains(err.Error(), "conflicting multipart upload ids") {
		t.Fatalf("CommitUpload() error = %v, want storage conflict", err)
	}
	if len(store.multipartBegins) != 0 {
		t.Fatalf("multipart begins = %#v, want no new multipart upload on conflict", store.multipartBegins)
	}
}

func TestCommitUploadLocalSpoolMissingPartReturnsFilePartXMissing(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9021})
	repo.uploadStateModel = model

	if err := model.SaveUploadPart(ctx, 1001, 7001, 0, []byte("aa")); err != nil {
		t.Fatalf("SaveUploadPart(0) error = %v", err)
	}
	if err := model.SaveUploadPart(ctx, 1001, 7001, 2, []byte("cc")); err != nil {
		t.Fatalf("SaveUploadPart(2) error = %v", err)
	}
	if err := model.SaveUploadFileInfo(ctx, &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileName:          "clip.mp4",
		FileTotalParts:    3,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		MimeType:          "video/mp4",
	}); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}

	file := tg.MakeTLInputFileBig(&tg.TLInputFileBig{Id: 7001, Parts: 3, Name: "clip.mp4"})
	_, err = repo.CommitUpload(ctx, "upload-missing-part", 1001, file, "media_original")
	var missing *dfs.MissingUploadPartError
	if !errors.As(err, &missing) || missing.Part != 1 {
		t.Fatalf("CommitUpload() error = %v, want MissingUploadPartError part 1", err)
	}
	if errors.Is(err, dfs.ErrDfsFileNotFound) || errors.Is(err, dfs.ErrDfsStorage) {
		t.Fatalf("CommitUpload() error = %v, should not be DFS file-not-found/storage", err)
	}
	if len(store.multipartBegins) != 0 {
		t.Fatalf("multipart begins = %#v, want missing part before remote staging", store.multipartBegins)
	}
}

func TestRepositoryScanSpoolOnStartReconcilesUploadedMultipartPart(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		LocalShardCount: 4,
		PartTTLSeconds:  0,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9025})
	repo.uploadStateModel = model

	info := &xkv.DfsFileInfo{Creator: 1001, FileID: 7001, FileTotalParts: 1, FirstFilePartSize: 4, Mtime: 1_700_000_000}
	if err := model.SaveUploadFileInfo(ctx, info); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	if err := model.SaveUploadPart(ctx, 1001, 7001, 0, []byte("aabb")); err != nil {
		t.Fatalf("SaveUploadPart() error = %v", err)
	}
	segmentPath, err := model.SegmentPath(info, 0)
	if err != nil {
		t.Fatalf("SegmentPath() error = %v", err)
	}
	if err := os.MkdirAll(filepath.Dir(segmentPath), 0o755); err != nil {
		t.Fatalf("MkdirAll(segment dir) error = %v", err)
	}
	if err := os.WriteFile(segmentPath, []byte("aabb"), 0o644); err != nil {
		t.Fatalf("WriteFile(segment) error = %v", err)
	}
	uploadID := "mp-recovered"
	objectKey := "objects/recovered.dat"
	store.multipartParts[uploadID] = map[int][]byte{1: []byte("aabb")}
	store.multipartPartETags[uploadID] = map[int]string{1: "etag-recovered-1"}
	if err := model.WriteSegmentState(ctx, info, spool.SegmentState{
		SegmentNo:           0,
		Status:              spool.SegmentStatusUploading,
		MultipartUploadID:   uploadID,
		MultipartPartNumber: 1,
		ObjectKey:           objectKey,
		Size:                4,
		Checksum:            fmt.Sprintf("%x", sha256.Sum256([]byte("aabb"))),
	}); err != nil {
		t.Fatalf("WriteSegmentState() error = %v", err)
	}

	if err := repo.ScanSpoolOnStart(ctx); err != nil {
		t.Fatalf("ScanSpoolOnStart() error = %v", err)
	}
	state, err := model.LoadSegmentState(ctx, info, 0)
	if err != nil {
		t.Fatalf("LoadSegmentState() error = %v", err)
	}
	if state.Status != spool.SegmentStatusUploaded || state.ETag != "etag-recovered-1" || state.Size != 4 {
		t.Fatalf("segment state = %#v, want uploaded reconciled state", state)
	}
	if _, err := os.Stat(segmentPath); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("segment bytes stat error = %v, want not exist", err)
	}
}

func TestRepositoryScanSpoolOnStartDoesNotReconcileWhenLocalChecksumChanged(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		SegmentSize:     4,
		LocalShardCount: 4,
		PartTTLSeconds:  0,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9029})
	repo.uploadStateModel = model

	info := &xkv.DfsFileInfo{
		Creator:           1001,
		FileID:            7001,
		FileTotalParts:    2,
		FirstFilePartSize: 2,
		FilePartSize:      2,
		LastFilePartSize:  2,
		Mtime:             1_700_000_000,
	}
	if err := model.SaveUploadFileInfo(ctx, info); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	for i, data := range [][]byte{[]byte("xx"), []byte("yy")} {
		if err := model.SaveUploadPart(ctx, 1001, 7001, int32(i), data); err != nil {
			t.Fatalf("SaveUploadPart(%d) error = %v", i, err)
		}
	}
	segmentPath, err := model.SegmentPath(info, 0)
	if err != nil {
		t.Fatalf("SegmentPath() error = %v", err)
	}
	if err := os.MkdirAll(filepath.Dir(segmentPath), 0o755); err != nil {
		t.Fatalf("MkdirAll(segment dir) error = %v", err)
	}
	if err := os.WriteFile(segmentPath, []byte("aabb"), 0o644); err != nil {
		t.Fatalf("WriteFile(segment) error = %v", err)
	}
	uploadID := "mp-recovered-mismatch"
	objectKey := "objects/recovered-mismatch.dat"
	store.multipartParts[uploadID] = map[int][]byte{1: []byte("aabb")}
	store.multipartPartETags[uploadID] = map[int]string{1: "etag-recovered-1"}
	if err := model.WriteSegmentState(ctx, info, spool.SegmentState{
		SegmentNo:           0,
		Status:              spool.SegmentStatusUploading,
		MultipartUploadID:   uploadID,
		MultipartPartNumber: 1,
		ObjectKey:           objectKey,
		ETag:                "etag-recovered-1",
		Checksum:            fmt.Sprintf("%x", sha256.Sum256([]byte("aabb"))),
		Size:                4,
	}); err != nil {
		t.Fatalf("WriteSegmentState() error = %v", err)
	}

	if err := repo.ScanSpoolOnStart(ctx); err != nil {
		t.Fatalf("ScanSpoolOnStart() error = %v", err)
	}
	state, err := model.LoadSegmentState(ctx, info, 0)
	if err != nil {
		t.Fatalf("LoadSegmentState() error = %v", err)
	}
	if state.Status != spool.SegmentStatusUploading || state.ETag != "etag-recovered-1" {
		t.Fatalf("segment state = %#v, want still uploading", state)
	}
	if data, err := os.ReadFile(segmentPath); err != nil || string(data) != "aabb" {
		t.Fatalf("segment bytes = %q, err = %v; want retained aabb", data, err)
	}
}

func TestRepositoryScanSpoolOnStartReturnsStorageWhenMultipartListFails(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	store.listMultipartPartsErr = errors.New("list failed")
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		LocalShardCount: 4,
		PartTTLSeconds:  0,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9026})
	repo.uploadStateModel = model

	info := &xkv.DfsFileInfo{Creator: 1001, FileID: 7001, FileTotalParts: 1, FirstFilePartSize: 4, Mtime: 1_700_000_000}
	if err := model.SaveUploadFileInfo(ctx, info); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	if err := model.WriteSegmentState(ctx, info, spool.SegmentState{
		SegmentNo:           0,
		Status:              spool.SegmentStatusUploading,
		MultipartUploadID:   "mp-list-fail",
		MultipartPartNumber: 1,
		ObjectKey:           "objects/list-fail.dat",
		Size:                4,
	}); err != nil {
		t.Fatalf("WriteSegmentState() error = %v", err)
	}

	err = repo.ScanSpoolOnStart(ctx)
	if !errors.Is(err, dfs.ErrDfsStorage) || !strings.Contains(err.Error(), "list failed") {
		t.Fatalf("ScanSpoolOnStart() error = %v, want storage list failure", err)
	}
	state, loadErr := model.LoadSegmentState(ctx, info, 0)
	if loadErr != nil {
		t.Fatalf("LoadSegmentState() error = %v", loadErr)
	}
	if state.Status != spool.SegmentStatusUploading {
		t.Fatalf("segment state = %#v, want uploading preserved", state)
	}
}

func TestCleanupExpiredUploadSessionsAbortsMultipartBeforeRemovingLocalState(t *testing.T) {
	ctx := context.Background()
	store := newFakeObjectStore()
	model, err := spool.NewUploadStateModel(config.UploadSpoolConf{
		RootDir:         t.TempDir(),
		NodeIDFile:      "node_id",
		LocalShardCount: 4,
		PartTTLSeconds:  10,
	})
	if err != nil {
		t.Fatalf("NewUploadStateModel() error = %v", err)
	}
	repo := newFileObjectTestRepository(store, &fakeIDGenerator{id: 9022})
	repo.uploadStateModel = model

	expired := time.Unix(1_700_000_000, 0)
	info := &xkv.DfsFileInfo{Creator: 1001, FileID: 7001, FileTotalParts: 1, Mtime: expired.Unix()}
	if err := model.SaveUploadFileInfo(ctx, info); err != nil {
		t.Fatalf("SaveUploadFileInfo() error = %v", err)
	}
	if err := model.WriteSegmentState(ctx, info, spool.SegmentState{
		SegmentNo:           0,
		Status:              spool.SegmentStatusUploading,
		MultipartUploadID:   "mp-expired",
		MultipartPartNumber: 1,
		Size:                4,
	}); err != nil {
		t.Fatalf("WriteSegmentState() error = %v", err)
	}

	if err := repo.CleanupExpiredUploadSessions(ctx, expired.Add(time.Hour)); err != nil {
		t.Fatalf("CleanupExpiredUploadSessions() error = %v", err)
	}
	if len(store.multipartAborts) != 1 || store.multipartAborts[0] != "documents/objects/mp-expired.dat:mp-expired" {
		t.Fatalf("multipart aborts = %#v, want expired upload abort", store.multipartAborts)
	}
	if _, err := model.LoadUploadFileInfo(ctx, 1001, 7001); !errors.Is(err, spool.ErrUploadStateNotFound) {
		t.Fatalf("LoadUploadFileInfo() after cleanup error = %v, want not found", err)
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
	objects                map[string][]byte
	putCounts              map[string]int
	failOnceKeys           map[string]error
	nextMultipartID        int
	multipartParts         map[string]map[int][]byte
	multipartPartETags     map[string]map[int]string
	multipartBegins        []string
	multipartUploads       []string
	multipartCompletes     []string
	multipartAborts        []string
	completedMultipartIDs  map[string]struct{}
	abortedMultipartIDs    map[string]struct{}
	rangeRequests          []objectRangeRequest
	listMultipartPartsErr  error
	afterCompleteMultipart func(objects map[string][]byte, fullKey string)
}

type objectRangeRequest struct {
	Bucket string
	Key    string
	Offset int64
	Limit  int32
}

func newFakeObjectStore() *fakeObjectStore {
	return &fakeObjectStore{
		objects:               make(map[string][]byte),
		putCounts:             make(map[string]int),
		failOnceKeys:          make(map[string]error),
		multipartParts:        make(map[string]map[int][]byte),
		multipartPartETags:    make(map[string]map[int]string),
		completedMultipartIDs: make(map[string]struct{}),
		abortedMultipartIDs:   make(map[string]struct{}),
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

func (f *fakeObjectStore) GetObjectReader(_ context.Context, bucket, key string) (io.ReadCloser, error) {
	data, ok := f.objects[bucket+"/"+key]
	if !ok {
		return nil, minioadapter.ErrObjectNotFound
	}
	return io.NopCloser(bytes.NewReader(append([]byte(nil), data...))), nil
}

func (f *fakeObjectStore) GetObjectRange(_ context.Context, bucket, key string, offset int64, limit int32) ([]byte, error) {
	f.rangeRequests = append(f.rangeRequests, objectRangeRequest{Bucket: bucket, Key: key, Offset: offset, Limit: limit})
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

func (f *fakeObjectStore) BeginMultipartUpload(_ context.Context, bucket, key string) (string, error) {
	f.nextMultipartID++
	uploadID := fmt.Sprintf("mp-%d", f.nextMultipartID)
	full := bucket + "/" + key
	f.multipartBegins = append(f.multipartBegins, full+":"+uploadID)
	f.multipartParts[uploadID] = make(map[int][]byte)
	f.multipartPartETags[uploadID] = make(map[int]string)
	return uploadID, nil
}

func (f *fakeObjectStore) UploadMultipartPart(_ context.Context, bucket, key, uploadID string, partNumber int, r io.Reader, size int64, checksum string) (minioadapter.MultipartPart, error) {
	if _, done := f.completedMultipartIDs[uploadID]; done {
		return minioadapter.MultipartPart{}, fmt.Errorf("no such multipart upload %s", uploadID)
	}
	if _, aborted := f.abortedMultipartIDs[uploadID]; aborted {
		return minioadapter.MultipartPart{}, fmt.Errorf("no such multipart upload %s", uploadID)
	}
	if _, ok := f.multipartParts[uploadID]; !ok {
		return minioadapter.MultipartPart{}, fmt.Errorf("no such multipart upload %s", uploadID)
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return minioadapter.MultipartPart{}, err
	}
	if int64(len(data)) != size {
		return minioadapter.MultipartPart{}, fmt.Errorf("part size = %d, want %d", len(data), size)
	}
	etag := fmt.Sprintf("etag-%s-%d", uploadID, partNumber)
	f.multipartUploads = append(f.multipartUploads, bucket+"/"+key+":"+uploadID+":"+strconv.Itoa(partNumber)+":"+checksum)
	if f.multipartParts[uploadID] == nil {
		f.multipartParts[uploadID] = make(map[int][]byte)
	}
	if f.multipartPartETags[uploadID] == nil {
		f.multipartPartETags[uploadID] = make(map[int]string)
	}
	f.multipartParts[uploadID][partNumber] = append([]byte(nil), data...)
	f.multipartPartETags[uploadID][partNumber] = etag
	return minioadapter.MultipartPart{PartNumber: partNumber, ETag: etag, Size: size, Checksum: checksum}, nil
}

func (f *fakeObjectStore) ListMultipartParts(_ context.Context, _, _ string, uploadID string) ([]minioadapter.MultipartPart, error) {
	if f.listMultipartPartsErr != nil {
		return nil, f.listMultipartPartsErr
	}
	partData := f.multipartParts[uploadID]
	parts := make([]minioadapter.MultipartPart, 0, len(partData))
	for partNumber, data := range partData {
		parts = append(parts, minioadapter.MultipartPart{
			PartNumber: partNumber,
			ETag:       f.multipartPartETags[uploadID][partNumber],
			Size:       int64(len(data)),
		})
	}
	return parts, nil
}

func (f *fakeObjectStore) CompleteMultipartUpload(_ context.Context, bucket, key, uploadID string, parts []minioadapter.MultipartPart) (int64, error) {
	if _, done := f.completedMultipartIDs[uploadID]; done {
		return 0, fmt.Errorf("no such multipart upload %s", uploadID)
	}
	if _, aborted := f.abortedMultipartIDs[uploadID]; aborted {
		return 0, fmt.Errorf("no such multipart upload %s", uploadID)
	}
	partData, ok := f.multipartParts[uploadID]
	if !ok {
		return 0, fmt.Errorf("no such multipart upload %s", uploadID)
	}
	var out bytes.Buffer
	for _, part := range parts {
		data, ok := partData[part.PartNumber]
		if !ok {
			return 0, fmt.Errorf("missing uploaded part %d", part.PartNumber)
		}
		out.Write(data)
	}
	full := bucket + "/" + key
	f.objects[full] = append([]byte(nil), out.Bytes()...)
	if f.afterCompleteMultipart != nil {
		f.afterCompleteMultipart(f.objects, full)
	}
	f.putCounts[full]++
	f.multipartCompletes = append(f.multipartCompletes, full+":"+uploadID)
	f.completedMultipartIDs[uploadID] = struct{}{}
	delete(f.multipartParts, uploadID)
	delete(f.multipartPartETags, uploadID)
	return int64(out.Len()), nil
}

func (f *fakeObjectStore) AbortMultipartUpload(_ context.Context, bucket, key, uploadID string) error {
	f.multipartAborts = append(f.multipartAborts, bucket+"/"+key+":"+uploadID)
	f.abortedMultipartIDs[uploadID] = struct{}{}
	delete(f.multipartParts, uploadID)
	delete(f.multipartPartETags, uploadID)
	return nil
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
