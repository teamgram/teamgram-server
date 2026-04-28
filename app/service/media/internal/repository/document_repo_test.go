package repository

import (
	"context"
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
