package repository

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMapDocumentIncludesThumbsVideoThumbsAndAttributes(t *testing.T) {
	got, err := mapDocumentAggregate(
		&model.Documents{
			DocumentId: 20,
			AccessHash: 30,
			DcId:       1,
			Date2:      40,
			MimeType:   "application/pdf",
			FileSize:   50,
			Attributes: `[{"_name":"documentAttributeFilename","_object":{"file_name":"report.pdf"}}]`,
		},
		[]model.PhotoSizes{{SizeType: "m", Width: 320, Height: 200, FileSize: 123}},
		[]model.VideoSizes{{SizeType: "v", Width: 320, Height: 200, FileSize: 456, VideoStartTs: 2}},
	)
	if err != nil {
		t.Fatalf("mapDocumentAggregate() error = %v", err)
	}
	doc, ok := got.ToDocument()
	if !ok {
		t.Fatalf("Document clazz = %s, want document", got.ClazzName())
	}
	if doc.Id != 20 || doc.MimeType != "application/pdf" {
		t.Fatalf("document = %+v, want id 20 pdf", doc)
	}
	if len(doc.Thumbs) != 1 {
		t.Fatalf("len(Thumbs) = %d, want 1", len(doc.Thumbs))
	}
	if len(doc.VideoThumbs) != 1 {
		t.Fatalf("len(VideoThumbs) = %d, want 1", len(doc.VideoThumbs))
	}
	if len(doc.Attributes) != 1 {
		t.Fatalf("len(Attributes) = %d, want 1", len(doc.Attributes))
	}
	if _, ok := doc.Attributes[0].(*tg.TLDocumentAttributeFilename); !ok {
		t.Fatalf("attribute = %T, want filename", doc.Attributes[0])
	}
}

func TestMapDocumentReturnsDecodeErrorForInvalidLegacyAttributes(t *testing.T) {
	_, err := mapDocumentAggregate(&model.Documents{DocumentId: 20, Attributes: `[`}, nil, nil)
	if err == nil {
		t.Fatal("mapDocumentAggregate() error = nil, want decode error")
	}
}
