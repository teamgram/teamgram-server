package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
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
	got := documentFromModel(&model.Documents{
		DocumentId: 20,
		AccessHash: 30,
		DcId:       4,
		Date2:      40,
		MimeType:   "image/jpeg",
		FileSize:   50,
	})
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
		name        string
		mimeType    string
		attrs       []tg.DocumentAttributeClazz
		wantGif     bool
		wantMp4     bool
		wantRegular bool
	}{
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
			attrs: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "clip.mp4"}),
			},
			wantMp4: true,
		},
		{
			name:     "regular",
			mimeType: "application/pdf",
			attrs: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "report.pdf"}),
			},
			wantRegular: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			documents := &captureDocumentsModel{}
			dfsClient := &fakeDfsMediaClient{document: testDocumentWithThumbs(500)}
			r := &Repository{
				model:     &model.Models{DocumentsModel: documents, PhotoSizesModel: &capturePhotoSizesModel{}, VideoSizesModel: &captureVideoSizesModel{}},
				dfsClient: dfsClient,
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
			if (dfsClient.uploadGifReq != nil) != tt.wantGif {
				t.Fatalf("gif request mismatch: got %#v want %v", dfsClient.uploadGifReq, tt.wantGif)
			}
			if (dfsClient.uploadMp4Req != nil) != tt.wantMp4 {
				t.Fatalf("mp4 request mismatch: got %#v want %v", dfsClient.uploadMp4Req, tt.wantMp4)
			}
			if (dfsClient.uploadDocumentReq != nil) != tt.wantRegular {
				t.Fatalf("regular request mismatch: got %#v want %v", dfsClient.uploadDocumentReq, tt.wantRegular)
			}
			if len(documents.inserted) != 1 || documents.inserted[0].DocumentId != 500 {
				t.Fatalf("expected saved document 500, got %#v", documents.inserted)
			}
		})
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
			ThumbId:    700,
			Attributes: `[{"_name":"documentAttributeFilename","_object":{"file_name":"report.pdf"}}]`,
		},
	}}
	photoSizes := &capturePhotoSizesModel{byID: []model.PhotoSizes{{PhotoSizeId: 700, SizeType: "m", Width: 320, Height: 240, FileSize: 1000}}}
	r := &Repository{model: &model.Models{DocumentsModel: documents, PhotoSizesModel: photoSizes, VideoSizesModel: &captureVideoSizesModel{}}}

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
}

func TestGetDocumentListPreservesRequestedOrder(t *testing.T) {
	documents := &captureDocumentsModel{byID: map[int64]*model.Documents{
		1: {DocumentId: 1, AccessHash: 10, DcId: 1},
		2: {DocumentId: 2, AccessHash: 20, DcId: 1},
	}}
	r := &Repository{model: &model.Models{DocumentsModel: documents, PhotoSizesModel: &capturePhotoSizesModel{}, VideoSizesModel: &captureVideoSizesModel{}}}

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
