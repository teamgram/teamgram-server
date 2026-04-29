package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeDocumentRepository struct {
	nextDocumentID  int64
	nextEncryptedID int64
	documents       map[int64][]byte
	encrypted       map[int64][]byte
}

func newFakeDocumentRepository() *fakeDocumentRepository {
	return &fakeDocumentRepository{
		nextDocumentID:  7001,
		nextEncryptedID: 8001,
		documents:       make(map[int64][]byte),
		encrypted:       make(map[int64][]byte),
	}
}

func (f *fakeDocumentRepository) NextDocumentID(context.Context) (int64, error) {
	id := f.nextDocumentID
	f.nextDocumentID++
	return id, nil
}

func (f *fakeDocumentRepository) NextEncryptedFileID(context.Context) (int64, error) {
	id := f.nextEncryptedID
	f.nextEncryptedID++
	return id, nil
}

func (f *fakeDocumentRepository) SaveDocumentObject(_ context.Context, documentID int64, data []byte) (int64, error) {
	f.documents[documentID] = append([]byte(nil), data...)
	return int64(len(data)), nil
}

func (f *fakeDocumentRepository) SaveEncryptedObject(_ context.Context, fileID int64, data []byte) (int64, error) {
	f.encrypted[fileID] = append([]byte(nil), data...)
	return int64(len(data)), nil
}

func TestDfsUploadDocumentFileV2CachesUploadRefAndReturnsDocument(t *testing.T) {
	core, uploads, docs := newDocumentTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 2002, []byte("document-bytes"))

	out, err := core.DfsUploadDocumentFileV2(&dfs.TLDfsUploadDocumentFileV2{
		Creator: 1001,
		Media: tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
			File:     tg.MakeTLInputFile(&tg.TLInputFile{Id: 2002, Parts: 1, Name: "report.pdf"}),
			MimeType: "application/pdf",
			Attributes: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}),
				tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: ""}),
				tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "report.pdf"}),
			},
		}),
	})
	if err != nil {
		t.Fatalf("DfsUploadDocumentFileV2() error = %v", err)
	}
	document, ok := out.ToDocument()
	if !ok {
		t.Fatalf("Document clazz = %s, want document", out.ClazzName())
	}
	if document.Id != 7001 {
		t.Fatalf("document.Id = %d, want 7001", document.Id)
	}
	if document.MimeType != "application/pdf" || document.Size2 != int64(len("document-bytes")) {
		t.Fatalf("document mime/size = %s/%d", document.MimeType, document.Size2)
	}
	if len(document.Attributes) != 1 {
		t.Fatalf("len(document.Attributes) = %d, want 1", len(document.Attributes))
	}
	if string(docs.documents[7001]) != "document-bytes" {
		t.Fatalf("stored document = %q, want document-bytes", docs.documents[7001])
	}
	info, err := uploads.LoadObjectCacheRef(context.Background(), 7001)
	if err != nil {
		t.Fatalf("LoadObjectCacheRef() error = %v", err)
	}
	if info.Creator != 1001 || info.FileID != 2002 {
		t.Fatalf("cache ref = %+v, want creator/file 1001/2002", info)
	}
}

func TestDfsUploadEncryptedFileV2StoresEncryptedObject(t *testing.T) {
	core, uploads, docs := newDocumentTestCore(t)
	writeUploadedTestFile(t, uploads, 1001, 3003, []byte("secret-bytes"))

	out, err := core.DfsUploadEncryptedFileV2(&dfs.TLDfsUploadEncryptedFileV2{
		Creator: 1001,
		File: tg.MakeTLInputEncryptedFileUploaded(&tg.TLInputEncryptedFileUploaded{
			Id:             3003,
			Parts:          1,
			Md5Checksum:    "dbfc252afe1c84d312028ac2d558a2c3",
			KeyFingerprint: 1234,
		}),
	})
	if err != nil {
		t.Fatalf("DfsUploadEncryptedFileV2() error = %v", err)
	}
	encrypted, ok := out.ToEncryptedFile()
	if !ok {
		t.Fatalf("EncryptedFile clazz = %s, want encryptedFile", out.ClazzName())
	}
	if encrypted.Id != 8001 {
		t.Fatalf("encrypted.Id = %d, want 8001", encrypted.Id)
	}
	if encrypted.Size2 != int64(len("secret-bytes")) {
		t.Fatalf("encrypted.Size2 = %d, want %d", encrypted.Size2, len("secret-bytes"))
	}
	if encrypted.KeyFingerprint != 0 {
		t.Fatalf("encrypted.KeyFingerprint = %d, want 0", encrypted.KeyFingerprint)
	}
	if string(docs.encrypted[8001]) != "secret-bytes" {
		t.Fatalf("stored encrypted = %q, want secret-bytes", docs.encrypted[8001])
	}
	info, err := uploads.LoadObjectCacheRef(context.Background(), 8001)
	if err != nil {
		t.Fatalf("LoadObjectCacheRef() error = %v", err)
	}
	if info.Creator != 1001 || info.FileID != 3003 {
		t.Fatalf("cache ref = %+v, want creator/file 1001/3003", info)
	}
}

func newDocumentTestCore(t *testing.T) (*DfsCore, *UploadSessionManager, *fakeDocumentRepository) {
	t.Helper()
	uploadRepo := newFakeUploadStateRepository()
	uploads := NewUploadSessionManager(uploadRepo)
	docs := newFakeDocumentRepository()
	return &DfsCore{
		ctx:                  context.Background(),
		uploadSessionManager: uploads,
		documentRepository:   docs,
	}, uploads, docs
}
