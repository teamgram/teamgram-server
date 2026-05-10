package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	dfsapi "github.com/teamgram/teamgram-server/v2/app/service/dfs/dfs"
	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type documentModelNotFound struct {
	model.DocumentsModel
}

func (documentModelNotFound) FindOneByDocumentId(context.Context, int64) (*model.Documents, error) {
	return nil, &model.NotFoundError{Resource: "documents", Key: "document_id=20"}
}

type captureDocumentsModel struct {
	model.DocumentsModel
	inserted []*model.Documents
	byID     map[int64]*model.Documents
}

func (m *captureDocumentsModel) Insert2(_ context.Context, data *model.Documents) (sql.Result, error) {
	m.inserted = append(m.inserted, data)
	return nil, nil
}

func (m *captureDocumentsModel) FindOneByDocumentId(_ context.Context, id int64) (*model.Documents, error) {
	if m.byID != nil {
		if doc := m.byID[id]; doc != nil {
			return doc, nil
		}
	}
	return nil, &model.NotFoundError{Resource: "documents", Key: "document_id"}
}

func TestGetDocumentReturnsStorageError(t *testing.T) {
	r := &Repository{}
	errBoom := errors.New("db down")
	_, err := r.mapDocumentResult(context.Background(), nil, errBoom)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestGetDocumentMapsModelNotFound(t *testing.T) {
	r := &Repository{model: &model.Models{DocumentsModel: documentModelNotFound{}}}
	_, err := r.GetDocument(context.Background(), 20)
	if !errors.Is(err, media.ErrDocumentNotFound) {
		t.Fatalf("expected ErrDocumentNotFound, got %v", err)
	}
	if errors.Is(err, media.ErrMediaStorage) {
		t.Fatalf("expected semantic not found, got storage error: %v", err)
	}
}

func TestGetDocumentReturnsDocument(t *testing.T) {
	r := &Repository{}
	doc := tg.MakeTLDocument(&tg.TLDocument{Id: 20}).ToDocument()
	got, err := r.mapDocumentResult(context.Background(), doc, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	gotDocument, ok := got.ToDocument()
	if !ok {
		t.Fatalf("expected document, got %#v", got)
	}
	if gotDocument.Id != 20 {
		t.Fatalf("expected document id 20, got %d", gotDocument.Id)
	}
}

func TestDocumentFromModelBuildsValidMinimalDocument(t *testing.T) {
	got, err := mapDocumentAggregate(&model.Documents{
		DocumentId: 20,
		AccessHash: 30,
		DcId:       4,
		Date2:      40,
		MimeType:   "image/jpeg",
		FileSize:   50,
	}, nil, nil, []byte("file-reference"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	gotDocument, ok := got.ToDocument()
	if !ok {
		t.Fatalf("expected document, got %#v", got)
	}
	if gotDocument.Id != 20 {
		t.Fatalf("expected document id 20, got %d", gotDocument.Id)
	}
	if gotDocument.FileReference == nil {
		t.Fatal("expected non-nil file_reference")
	}
	if gotDocument.Attributes == nil {
		t.Fatal("expected non-nil required attributes")
	}
	if gotDocument.Thumbs != nil {
		t.Fatalf("expected absent thumbs, got %#v", gotDocument.Thumbs)
	}
	if gotDocument.VideoThumbs != nil {
		t.Fatalf("expected absent video_thumbs, got %#v", gotDocument.VideoThumbs)
	}
	if err := gotDocument.Validate(testLayer); err != nil {
		t.Fatalf("expected valid document: %v", err)
	}
}

func TestUploadedDocumentMediaDispatchesGifMp4Document(t *testing.T) {
	tests := []struct {
		name          string
		mimeType      string
		attrs         []tg.DocumentAttributeClazz
		wantProcessor string
	}{
		{name: "gif", mimeType: "image/gif", attrs: []tg.DocumentAttributeClazz{
			tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}),
			tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "loop.gif"}),
		}, wantProcessor: "gif"},
		{name: "mp4", mimeType: "video/mp4", attrs: []tg.DocumentAttributeClazz{
			tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "clip.mp4"}),
		}, wantProcessor: "mp4"},
		{name: "regular", mimeType: "application/pdf", attrs: []tg.DocumentAttributeClazz{
			tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "report.pdf"}),
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			documents := &captureDocumentsModel{}
			dfsClient := &fakeDfsMediaClient{finalized: dfsapi.MakeTLFileFinalizedObject(&dfsapi.TLFileFinalizedObject{
				ObjectId:        "original-object-" + tt.name,
				UploadSessionId: "ext:77:8:1",
				ReadLease:       []byte("read-lease-" + tt.name),
				Size2:           2048,
				DcId:            2,
			}).ToFileFinalizedObject()}
			processorClient := &fakeMediaProcessorClient{document: testProcessedDocument(t, "processed-object-"+tt.name, tt.mimeType, 4096, tt.attrs, nil)}
			r := &Repository{
				model:                &model.Models{DocumentsModel: documents, FileReferencesModel: newCaptureFileReferencesModel(), PhotoSizesModel: &capturePhotoSizesModel{}, VideoSizesModel: &captureVideoSizesModel{}},
				dfsClient:            dfsClient,
				processorClient:      processorClient,
				fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
				fileReferenceTTL:     time.Hour,
			}

			got, err := r.UploadedDocumentMedia(context.Background(), &media.TLMediaUploadedDocumentMedia{
				OwnerId: 77,
				Media: tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
					File:       tg.MakeTLInputFile(&tg.TLInputFile{Id: 8, Parts: 1, Name: "upload.bin"}),
					MimeType:   tt.mimeType,
					Attributes: tt.attrs,
				}),
			})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got == nil {
				t.Fatal("expected message media document")
			}
			if dfsClient.commitReq == nil {
				t.Fatal("expected dfs CommitUpload request")
			}
			if dfsClient.uploadDocumentReq != nil || dfsClient.uploadGifReq != nil || dfsClient.uploadMp4Req != nil {
				t.Fatalf("expected no legacy document dfs upload calls, got document=%#v gif=%#v mp4=%#v", dfsClient.uploadDocumentReq, dfsClient.uploadGifReq, dfsClient.uploadMp4Req)
			}
			if (processorClient.gifReq != nil) != (tt.wantProcessor == "gif") {
				t.Fatalf("gif processor request mismatch: got %#v want %q", processorClient.gifReq, tt.wantProcessor)
			}
			if (processorClient.mp4Req != nil) != (tt.wantProcessor == "mp4") {
				t.Fatalf("mp4 processor request mismatch: got %#v want %q", processorClient.mp4Req, tt.wantProcessor)
			}
			if len(documents.inserted) != 1 {
				t.Fatalf("expected one saved document, got %#v", documents.inserted)
			}
		})
	}
}

func TestUploadedDocumentMediaCommitsAndProcessesByMime(t *testing.T) {
	videoAttrs := []tg.DocumentAttributeClazz{
		tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{SupportsStreaming: true, Duration: 3, W: 640, H: 360}),
		tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "clip.mp4"}),
	}
	gifAttrs := append([]tg.DocumentAttributeClazz{}, videoAttrs...)
	gifAttrs = append(gifAttrs, tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}))

	tests := []struct {
		name          string
		mimeType      string
		fileName      string
		attrs         []tg.DocumentAttributeClazz
		finalized     *dfsapi.FileFinalizedObject
		processed     *mediaprocessor.ProcessedDocument
		wantProcessor string
		wantMime      string
		wantObjectID  string
		wantSize      int64
	}{
		{
			name:     "animated gif",
			mimeType: "image/gif",
			fileName: "loop.gif",
			attrs: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}),
				tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "loop.gif"}),
			},
			finalized:     testFinalizedObject("original-gif-object", 111),
			processed:     testProcessedDocument(t, "processed-gif-object", "video/mp4", 222, gifAttrs, testDocumentThumbs()),
			wantProcessor: "gif",
			wantMime:      "video/mp4",
			wantObjectID:  "processed-gif-object",
			wantSize:      222,
		},
		{
			name:          "mp4",
			mimeType:      "video/mp4",
			fileName:      "clip.mp4",
			attrs:         []tg.DocumentAttributeClazz{tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "clip.mp4"})},
			finalized:     testFinalizedObject("original-mp4-object", 333),
			processed:     testProcessedDocument(t, "processed-mp4-object", "video/mp4", 444, videoAttrs, nil),
			wantProcessor: "mp4",
			wantMime:      "video/mp4",
			wantObjectID:  "processed-mp4-object",
			wantSize:      444,
		},
		{
			name:         "plain document",
			mimeType:     "application/pdf",
			fileName:     "report.pdf",
			attrs:        []tg.DocumentAttributeClazz{tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "report.pdf"})},
			finalized:    testFinalizedObject("original-pdf-object", 555),
			wantMime:     "application/pdf",
			wantObjectID: "original-pdf-object",
			wantSize:     555,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			documents := &captureDocumentsModel{}
			photoSizes := &capturePhotoSizesModel{}
			dfsClient := &fakeDfsMediaClient{finalized: tt.finalized}
			processorClient := &fakeMediaProcessorClient{document: tt.processed}
			r := &Repository{
				model:                &model.Models{DocumentsModel: documents, FileReferencesModel: newCaptureFileReferencesModel(), PhotoSizesModel: photoSizes, VideoSizesModel: &captureVideoSizesModel{}},
				dfsClient:            dfsClient,
				processorClient:      processorClient,
				fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
				fileReferenceTTL:     time.Hour,
			}

			got, err := r.UploadedDocumentMedia(context.Background(), &media.TLMediaUploadedDocumentMedia{
				OwnerId: 77,
				Media: tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
					File:       tg.MakeTLInputFile(&tg.TLInputFile{Id: 8, Parts: 1, Name: "upload.bin"}),
					MimeType:   tt.mimeType,
					Attributes: tt.attrs,
				}),
			})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if dfsClient.commitReq == nil {
				t.Fatal("expected dfs CommitUpload request")
			}
			if dfsClient.commitReq.UploadSessionId != "ext:77:8:1" || dfsClient.commitReq.OwnerId != 77 || dfsClient.commitReq.File == nil || dfsClient.commitReq.Purpose != "media_original" {
				t.Fatalf("unexpected dfs commit request: %#v", dfsClient.commitReq)
			}
			if dfsClient.uploadDocumentReq != nil || dfsClient.uploadGifReq != nil || dfsClient.uploadMp4Req != nil {
				t.Fatalf("expected no legacy document dfs upload calls, got document=%#v gif=%#v mp4=%#v", dfsClient.uploadDocumentReq, dfsClient.uploadGifReq, dfsClient.uploadMp4Req)
			}
			if (processorClient.gifReq != nil) != (tt.wantProcessor == "gif") {
				t.Fatalf("gif processor request mismatch: got %#v want %q", processorClient.gifReq, tt.wantProcessor)
			}
			if (processorClient.mp4Req != nil) != (tt.wantProcessor == "mp4") {
				t.Fatalf("mp4 processor request mismatch: got %#v want %q", processorClient.mp4Req, tt.wantProcessor)
			}
			if processorClient.gifReq != nil {
				if processorClient.gifReq.OwnerId != 77 || processorClient.gifReq.ObjectId != tt.finalized.ObjectId || string(processorClient.gifReq.ReadLease) != "read-lease" || processorClient.gifReq.FileName != tt.fileName {
					t.Fatalf("unexpected gif processor request: %#v", processorClient.gifReq)
				}
			}
			if processorClient.mp4Req != nil {
				if processorClient.mp4Req.OwnerId != 77 || processorClient.mp4Req.ObjectId != tt.finalized.ObjectId || string(processorClient.mp4Req.ReadLease) != "read-lease" || processorClient.mp4Req.FileName != tt.fileName || len(processorClient.mp4Req.Attributes) == 0 {
					t.Fatalf("unexpected mp4 processor request: %#v", processorClient.mp4Req)
				}
			}

			mediaDoc, ok := got.ToMessageMediaDocument()
			if !ok || mediaDoc.Document == nil {
				t.Fatalf("expected messageMediaDocument, got %#v", got)
			}
			doc, ok := mediaDoc.Document.(*tg.TLDocument)
			if !ok {
				t.Fatalf("expected TLDocument, got %#v", mediaDoc.Document)
			}
			if len(doc.FileReference) != 25 {
				t.Fatalf("len(file_reference) = %d, want 25", len(doc.FileReference))
			}
			if doc.MimeType != tt.wantMime || doc.Size2 != tt.wantSize {
				t.Fatalf("unexpected returned document mime/size: %#v", doc)
			}
			if len(doc.Attributes) == 0 {
				t.Fatalf("expected returned document attributes, got %#v", doc)
			}
			if len(documents.inserted) != 1 {
				t.Fatalf("expected one saved document, got %#v", documents.inserted)
			}
			row := documents.inserted[0]
			if row.MimeType != tt.wantMime || row.FileSize != tt.wantSize || row.UploadedFileName != tt.fileName || row.FilePath != tt.wantObjectID {
				t.Fatalf("unexpected saved document row: %#v", row)
			}
			if row.FilePath == documentObjectPath(row.DocumentId) {
				t.Fatalf("expected object id file_path, got legacy path %q", row.FilePath)
			}
			if tt.wantProcessor == "gif" {
				if len(photoSizes.inserted) != 1 || photoSizes.inserted[0].SizeType != "m" || photoSizes.inserted[0].FilePath != "gif-thumb-object" {
					t.Fatalf("expected saved document thumb derivative, got %#v", photoSizes.inserted)
				}
			} else if len(photoSizes.inserted) != 0 {
				t.Fatalf("expected no saved document thumbs, got %#v", photoSizes.inserted)
			}
		})
	}
}

func TestUploadedDocumentMediaReturnsVideoFlagsForMp4(t *testing.T) {
	got := uploadTestDocumentMedia(t, false)
	mediaDoc, ok := got.ToMessageMediaDocument()
	if !ok {
		t.Fatalf("expected messageMediaDocument, got %#v", got)
	}
	if !mediaDoc.Video || mediaDoc.Round || mediaDoc.Voice {
		t.Fatalf("unexpected messageMediaDocument flags: %#v", mediaDoc)
	}
}

func TestUploadedDocumentMediaReturnsNoVideoFlagForForceFileMp4(t *testing.T) {
	got := uploadTestDocumentMedia(t, true)
	mediaDoc, ok := got.ToMessageMediaDocument()
	if !ok {
		t.Fatalf("expected messageMediaDocument, got %#v", got)
	}
	if mediaDoc.Video || mediaDoc.Round || mediaDoc.Voice {
		t.Fatalf("unexpected force-file messageMediaDocument flags: %#v", mediaDoc)
	}
}

func TestUploadedDocumentMediaViaLegacyDFSCallsLegacyWrappers(t *testing.T) {
	tests := []struct {
		name         string
		mimeType     string
		attrs        []tg.DocumentAttributeClazz
		wantDocument bool
		wantGif      bool
		wantMp4      bool
	}{
		{
			name:         "regular",
			mimeType:     "application/pdf",
			attrs:        []tg.DocumentAttributeClazz{tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "report.pdf"})},
			wantDocument: true,
		},
		{
			name:     "gif",
			mimeType: "image/gif",
			attrs: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}),
				tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "loop.gif"}),
			},
			wantGif: true,
		},
		{
			name:     "mp4",
			mimeType: "video/mp4",
			attrs:    []tg.DocumentAttributeClazz{tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "clip.mp4"})},
			wantMp4:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			documents := &captureDocumentsModel{}
			photoSizes := &capturePhotoSizesModel{}
			dfsClient := &fakeDfsMediaClient{document: testDocumentWithThumbs(808)}
			r := &Repository{
				model:                &model.Models{DocumentsModel: documents, FileReferencesModel: newCaptureFileReferencesModel(), PhotoSizesModel: photoSizes, VideoSizesModel: &captureVideoSizesModel{}},
				dfsClient:            dfsClient,
				fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
				fileReferenceTTL:     time.Hour,
			}

			got, err := r.UploadedDocumentMediaViaLegacyDFS(context.Background(), &media.TLMediaUploadedDocumentMedia{
				OwnerId: 77,
				Media: tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
					File:       tg.MakeTLInputFile(&tg.TLInputFile{Id: 8, Parts: 1, Name: "upload.bin"}),
					MimeType:   tt.mimeType,
					Attributes: tt.attrs,
				}),
			})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got == nil {
				t.Fatal("expected message media document")
			}
			if (dfsClient.uploadDocumentReq != nil) != tt.wantDocument || (dfsClient.uploadGifReq != nil) != tt.wantGif || (dfsClient.uploadMp4Req != nil) != tt.wantMp4 {
				t.Fatalf("legacy request mismatch: document=%#v gif=%#v mp4=%#v", dfsClient.uploadDocumentReq, dfsClient.uploadGifReq, dfsClient.uploadMp4Req)
			}
			if dfsClient.commitReq != nil {
				t.Fatalf("expected legacy path not to commit upload, got %#v", dfsClient.commitReq)
			}
			if len(documents.inserted) != 1 || documents.inserted[0].DocumentId != 808 {
				t.Fatalf("expected saved legacy document row, got %#v", documents.inserted)
			}
			if len(photoSizes.inserted) != 1 || photoSizes.inserted[0].PhotoSizeId != 808 {
				t.Fatalf("expected saved legacy document thumb, got %#v", photoSizes.inserted)
			}
		})
	}
}

func TestUploadedDocumentMediaViaLegacyDFSRejectsNilLegacyDocument(t *testing.T) {
	r := &Repository{
		model:     &model.Models{DocumentsModel: &captureDocumentsModel{}, PhotoSizesModel: &capturePhotoSizesModel{}, VideoSizesModel: &captureVideoSizesModel{}},
		dfsClient: &fakeDfsMediaClient{},
	}

	_, err := r.UploadedDocumentMediaViaLegacyDFS(context.Background(), &media.TLMediaUploadedDocumentMedia{
		OwnerId: 77,
		Media: tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
			File:       tg.MakeTLInputFile(&tg.TLInputFile{Id: 8, Parts: 1, Name: "upload.bin"}),
			MimeType:   "application/pdf",
			Attributes: []tg.DocumentAttributeClazz{tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "report.pdf"})},
		}),
	})
	if !errors.Is(err, media.ErrMediaInvalidUploadedFile) {
		t.Fatalf("expected ErrMediaInvalidUploadedFile, got %v", err)
	}
}

func TestUploadedDocumentMediaRejectsZeroSizePlainDocument(t *testing.T) {
	documents := &captureDocumentsModel{}
	photoSizes := &capturePhotoSizesModel{}
	dfsClient := &fakeDfsMediaClient{finalized: testFinalizedObject("zero-size-pdf-object", 0)}
	r := &Repository{
		model:                &model.Models{DocumentsModel: documents, FileReferencesModel: newCaptureFileReferencesModel(), PhotoSizesModel: photoSizes, VideoSizesModel: &captureVideoSizesModel{}},
		dfsClient:            dfsClient,
		fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
		fileReferenceTTL:     time.Hour,
	}

	_, err := r.UploadedDocumentMedia(context.Background(), &media.TLMediaUploadedDocumentMedia{
		OwnerId: 77,
		Media: tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
			File:       tg.MakeTLInputFile(&tg.TLInputFile{Id: 8, Parts: 1, Name: "upload.bin"}),
			MimeType:   "application/pdf",
			Attributes: []tg.DocumentAttributeClazz{tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "report.pdf"})},
		}),
	})
	if !errors.Is(err, media.ErrMediaInvalidUploadedFile) {
		t.Fatalf("expected ErrMediaInvalidUploadedFile, got %v", err)
	}
	if len(documents.inserted) != 0 || len(photoSizes.inserted) != 0 {
		t.Fatalf("expected no persisted rows, got documents=%#v photo_sizes=%#v", documents.inserted, photoSizes.inserted)
	}
}

func TestGetDocumentLoadsThumbsAndAttributes(t *testing.T) {
	documents := &captureDocumentsModel{byID: map[int64]*model.Documents{
		700: {
			DocumentId: 700,
			AccessHash: 10,
			DcId:       1,
			MimeType:   "application/pdf",
			FileSize:   123,
			FilePath:   "document-object-700",
			ThumbId:    700,
			Attributes: `[{"_name":"documentAttributeFilename","_object":{"file_name":"report.pdf"}}]`,
		},
	}}
	photoSizes := &capturePhotoSizesModel{byID: []model.PhotoSizes{{PhotoSizeId: 700, SizeType: "m", Width: 320, Height: 240, FileSize: 1000}}}
	r := &Repository{
		model:                &model.Models{DocumentsModel: documents, FileReferencesModel: newCaptureFileReferencesModel(), PhotoSizesModel: photoSizes, VideoSizesModel: &captureVideoSizesModel{}},
		fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
		fileReferenceTTL:     time.Hour,
	}

	got, err := r.GetDocument(context.Background(), 700)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	doc, ok := got.ToDocument()
	if !ok {
		t.Fatalf("expected document, got %#v", got)
	}
	if len(doc.Thumbs) != 1 || len(doc.Attributes) != 1 {
		t.Fatalf("expected thumbs and attributes, got %#v %#v", doc.Thumbs, doc.Attributes)
	}
	claims, err := r.fileReferenceService.Validate(context.Background(), doc.FileReference, r)
	if err != nil {
		t.Fatalf("expected valid loaded document file_reference: %v", err)
	}
	if claims.OriginDomain != "document" || claims.MediaID != 700 || claims.AccessHash != 10 || claims.ObjectID != "document-object-700" {
		t.Fatalf("unexpected loaded document file_reference claims: %#v", claims)
	}
}

func TestGetDocumentListPreservesRequestedOrder(t *testing.T) {
	documents := &captureDocumentsModel{byID: map[int64]*model.Documents{
		1: {DocumentId: 1, AccessHash: 10, DcId: 1},
		2: {DocumentId: 2, AccessHash: 20, DcId: 1},
	}}
	r := &Repository{
		model:                &model.Models{DocumentsModel: documents, FileReferencesModel: newCaptureFileReferencesModel(), PhotoSizesModel: &capturePhotoSizesModel{}, VideoSizesModel: &captureVideoSizesModel{}},
		fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
		fileReferenceTTL:     time.Hour,
	}

	got, err := r.GetDocumentList(context.Background(), []int64{2, 99, 1})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Datas) != 2 {
		t.Fatalf("expected two existing documents, got %d", len(got.Datas))
	}
	first, _ := got.Datas[0].(*tg.TLDocument)
	second, _ := got.Datas[1].(*tg.TLDocument)
	if first.Id != 2 || second.Id != 1 {
		t.Fatalf("expected requested order [2,1], got [%d,%d]", first.Id, second.Id)
	}
}

func testDocumentWithThumbs(id int64) *tg.Document {
	return tg.MakeTLDocument(&tg.TLDocument{
		Id:         id,
		AccessHash: 99,
		Date:       10,
		MimeType:   "application/octet-stream",
		Size2:      2048,
		DcId:       1,
		Thumbs: []tg.PhotoSizeClazz{
			tg.MakeTLPhotoSize(&tg.TLPhotoSize{Type: "m", W: 320, H: 240, Size2: 1000}),
		},
		Attributes: []tg.DocumentAttributeClazz{
			tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "upload.bin"}),
		},
	}).ToDocument()
}

func testFinalizedObject(objectID string, size int64) *dfsapi.FileFinalizedObject {
	return dfsapi.MakeTLFileFinalizedObject(&dfsapi.TLFileFinalizedObject{
		ObjectId:        objectID,
		UploadSessionId: "ext:77:8:1",
		ReadLease:       []byte("read-lease"),
		Size2:           size,
		DcId:            2,
	}).ToFileFinalizedObject()
}

func testProcessedDocument(t *testing.T, objectID, mimeType string, size int64, attrs []tg.DocumentAttributeClazz, thumbs []mediaprocessor.ProcessorDerivativeClazz) *mediaprocessor.ProcessedDocument {
	t.Helper()
	return mediaprocessor.MakeTLProcessedDocument(&mediaprocessor.TLProcessedDocument{
		ObjectId:   objectID,
		MimeType:   mimeType,
		Size2:      size,
		Attributes: testDocumentAttributeVectorBytes(t, attrs),
		Thumbs:     thumbs,
	}).ToProcessedDocument()
}

func uploadTestDocumentMedia(t *testing.T, forceFile bool) *tg.MessageMedia {
	t.Helper()
	attrs := []tg.DocumentAttributeClazz{
		tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{SupportsStreaming: true, Duration: 3, W: 640, H: 360}),
		tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "clip.mp4"}),
	}
	documents := &captureDocumentsModel{}
	dfsClient := &fakeDfsMediaClient{finalized: testFinalizedObject("original-mp4-object", 333)}
	processorClient := &fakeMediaProcessorClient{document: testProcessedDocument(t, "processed-mp4-object", "video/mp4", 444, attrs, nil)}
	r := &Repository{
		model:                &model.Models{DocumentsModel: documents, FileReferencesModel: newCaptureFileReferencesModel(), PhotoSizesModel: &capturePhotoSizesModel{}, VideoSizesModel: &captureVideoSizesModel{}},
		dfsClient:            dfsClient,
		processorClient:      processorClient,
		fileReferenceService: NewFileReferenceService([]byte("test-secret"), func() time.Time { return time.Unix(1700000000, 0) }),
		fileReferenceTTL:     time.Hour,
	}

	got, err := r.UploadedDocumentMedia(context.Background(), &media.TLMediaUploadedDocumentMedia{
		OwnerId: 77,
		Media: tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
			ForceFile:  forceFile,
			File:       tg.MakeTLInputFile(&tg.TLInputFile{Id: 8, Parts: 1, Name: "clip.mp4"}),
			MimeType:   "video/mp4",
			Attributes: attrs,
		}),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if processorClient.mp4Req == nil {
		t.Fatal("expected mp4 processor request")
	}
	if len(documents.inserted) != 1 {
		t.Fatalf("expected one saved document, got %#v", documents.inserted)
	}
	return got
}

func testDocumentAttributeVectorBytes(t *testing.T, attrs []tg.DocumentAttributeClazz) []byte {
	t.Helper()
	x := bin.NewEncoder()
	if err := iface.EncodeObjectList(x, attrs, 224); err != nil {
		t.Fatalf("encode test document attrs: %v", err)
	}
	return x.Clone()
}

func testDocumentThumbs() []mediaprocessor.ProcessorDerivativeClazz {
	return []mediaprocessor.ProcessorDerivativeClazz{
		mediaprocessor.MakeTLProcessorDerivative(&mediaprocessor.TLProcessorDerivative{
			Kind:     "document_thumb",
			ObjectId: "gif-thumb-object",
			FileName: "loop_thumb.jpg",
			MimeType: "image/jpeg",
			Size2:    1234,
			Width:    320,
			Height:   180,
			Bytes:    []byte("thumb"),
		}),
	}
}
