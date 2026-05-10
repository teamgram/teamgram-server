package core

import (
	"errors"
	"reflect"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestUploadGetFileHashesUsesMediaLocatorAndReadLease(t *testing.T) {
	location := tg.MakeTLInputDocumentFileLocation(&tg.TLInputDocumentFileLocation{
		Id:            1001,
		AccessHash:    2002,
		FileReference: []byte("file-reference"),
	})
	hashes := []tg.FileHashClazz{
		tg.MakeTLFileHash(&tg.TLFileHash{Offset: 0, Limit: 1024, Hash: []byte("hash-1")}),
		tg.MakeTLFileHash(&tg.TLFileHash{Offset: 1024, Limit: 512, Hash: []byte("hash-2")}),
	}
	dfsClient := &fakeFilesDfsClient{
		hashResp: &dfs.VectorFileHash{Datas: hashes},
	}
	mediaClient := &fakeFilesMediaClient{
		resolveResp: mediapb.MakeTLMediaResolvedFileObject(&mediapb.TLMediaResolvedFileObject{
			ObjectId:        "object-1001",
			ReadLease:       []byte("read-lease"),
			Size2:           1536,
			MimeType:        "application/octet-stream",
			DcId:            1,
			StorageFileType: testStorageFileType(tg.ClazzID_storage_filePartial),
		}).ToMediaResolvedFileObject(),
	}
	core := newUploadGetFileTestCore(dfsClient, mediaClient, false)

	got, err := core.UploadGetFileHashes(&tg.TLUploadGetFileHashes{
		Location: location,
		Offset:   1024,
	})
	if err != nil {
		t.Fatalf("UploadGetFileHashes() error = %v", err)
	}
	if mediaClient.resolveReq == nil || mediaClient.resolveReq.Location != location || mediaClient.resolveReq.ViewerId != 1001 {
		t.Fatalf("resolve request = %#v", mediaClient.resolveReq)
	}
	if dfsClient.hashReq == nil || string(dfsClient.hashReq.ReadLease) != "read-lease" || dfsClient.hashReq.Offset != 1024 || dfsClient.hashReq.Limit != 0 {
		t.Fatalf("hash request = %#v", dfsClient.hashReq)
	}
	if dfsClient.readReq != nil || dfsClient.downloadReq != nil {
		t.Fatalf("unexpected read/download calls: read=%#v download=%#v", dfsClient.readReq, dfsClient.downloadReq)
	}
	if got == nil || !reflect.DeepEqual(got.Datas, hashes) {
		t.Fatalf("hashes = %#v, want %#v", got, hashes)
	}
}

func TestUploadGetFileHashesUnsupportedLocationReturnsLocationInvalid(t *testing.T) {
	location := tg.MakeTLInputEncryptedFileLocation(&tg.TLInputEncryptedFileLocation{Id: 1001, AccessHash: 2002})
	dfsClient := &fakeFilesDfsClient{}
	mediaClient := &fakeFilesMediaClient{}

	_, err := newUploadGetFileTestCore(dfsClient, mediaClient, true).UploadGetFileHashes(&tg.TLUploadGetFileHashes{Location: location})
	if !errors.Is(err, tg.ErrLocationInvalid) && err != tg.ErrLocationInvalid {
		t.Fatalf("UploadGetFileHashes() error = %v, want LOCATION_INVALID", err)
	}
	if mediaClient.resolveReq != nil || dfsClient.hashReq != nil || dfsClient.downloadReq != nil {
		t.Fatalf("unexpected calls: resolve=%#v hash=%#v download=%#v", mediaClient.resolveReq, dfsClient.hashReq, dfsClient.downloadReq)
	}
}

func TestUploadGetFileHashesPropagatesDfsStorageFailure(t *testing.T) {
	location := tg.MakeTLInputPhotoFileLocation(&tg.TLInputPhotoFileLocation{
		Id:            1001,
		AccessHash:    2002,
		FileReference: []byte("file-reference"),
		ThumbSize:     "m",
	})
	storageErr := errors.New("storage unavailable")
	dfsClient := &fakeFilesDfsClient{hashErr: storageErr}
	mediaClient := &fakeFilesMediaClient{
		resolveResp: mediapb.MakeTLMediaResolvedFileObject(&mediapb.TLMediaResolvedFileObject{
			ObjectId:        "photo-1001",
			ReadLease:       []byte("read-lease"),
			Size2:           1536,
			MimeType:        "image/jpeg",
			DcId:            1,
			StorageFileType: testStorageFileType(tg.ClazzID_storage_fileJpeg),
		}).ToMediaResolvedFileObject(),
	}

	_, err := newUploadGetFileTestCore(dfsClient, mediaClient, false).UploadGetFileHashes(&tg.TLUploadGetFileHashes{Location: location})
	if !errors.Is(err, storageErr) {
		t.Fatalf("UploadGetFileHashes() error = %v, want storage error", err)
	}
	if dfsClient.hashReq == nil {
		t.Fatalf("hash request was not issued")
	}
}
