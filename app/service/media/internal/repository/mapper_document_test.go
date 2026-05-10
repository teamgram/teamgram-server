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
		[]byte("file-reference"),
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
	_, err := mapDocumentAggregate(&model.Documents{DocumentId: 20, Attributes: `[`}, nil, nil, nil)
	if err == nil {
		t.Fatal("mapDocumentAggregate() error = nil, want decode error")
	}
}

func TestDocumentAttributePersistencePreservesStickerCustomEmojiAndHasStickers(t *testing.T) {
	attrs := []tg.DocumentAttributeClazz{
		tg.MakeTLDocumentAttributeFilename(&tg.TLDocumentAttributeFilename{FileName: "sticker.webp"}),
		tg.MakeTLDocumentAttributeSticker(&tg.TLDocumentAttributeSticker{
			Mask: true,
			Alt:  ":)",
			Stickerset: tg.MakeTLInputStickerSetID(&tg.TLInputStickerSetID{
				Id:         1001,
				AccessHash: 2002,
			}),
			MaskCoords: tg.MakeTLMaskCoords(&tg.TLMaskCoords{N: 1, X: 0.5, Y: 0.25, Zoom: 1.5}),
		}),
		tg.MakeTLDocumentAttributeCustomEmoji(&tg.TLDocumentAttributeCustomEmoji{
			Free:      true,
			TextColor: true,
			Alt:       ":)",
			Stickerset: tg.MakeTLInputStickerSetID(&tg.TLInputStickerSetID{
				Id:         3003,
				AccessHash: 4004,
			}),
		}),
		tg.MakeTLDocumentAttributeHasStickers(&tg.TLDocumentAttributeHasStickers{}),
	}

	vector, err := encodeDocumentAttributeVector(attrs)
	if err != nil {
		t.Fatalf("encodeDocumentAttributeVector() error = %v", err)
	}
	decodedVector, err := decodeDocumentAttributeVector(vector)
	if err != nil {
		t.Fatalf("decodeDocumentAttributeVector() error = %v", err)
	}
	if len(decodedVector) != len(attrs) {
		t.Fatalf("decoded vector attrs len = %d, want %d", len(decodedVector), len(attrs))
	}

	legacy, err := encodeLegacyDocumentAttributes(decodedVector)
	if err != nil {
		t.Fatalf("encodeLegacyDocumentAttributes() error = %v", err)
	}
	decodedLegacy, err := decodeLegacyDocumentAttributes(legacy)
	if err != nil {
		t.Fatalf("decodeLegacyDocumentAttributes() error = %v", err)
	}
	if !hasRepositoryDocumentAttribute[*tg.TLDocumentAttributeFilename](decodedLegacy) ||
		!hasRepositoryDocumentAttribute[*tg.TLDocumentAttributeSticker](decodedLegacy) ||
		!hasRepositoryDocumentAttribute[*tg.TLDocumentAttributeCustomEmoji](decodedLegacy) ||
		!hasRepositoryDocumentAttribute[*tg.TLDocumentAttributeHasStickers](decodedLegacy) {
		t.Fatalf("decoded legacy attrs = %#v, want filename/sticker/custom_emoji/has_stickers", decodedLegacy)
	}
}

func hasRepositoryDocumentAttribute[T tg.DocumentAttributeClazz](attrs []tg.DocumentAttributeClazz) bool {
	for _, attr := range attrs {
		if _, ok := attr.(T); ok {
			return true
		}
	}
	return false
}
