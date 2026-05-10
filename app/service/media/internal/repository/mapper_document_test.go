package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/media/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
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

func TestEncodeDocumentAttributesForStorageRoundTripsStickerAndCustomEmoji(t *testing.T) {
	attrs := repositoryDocumentAttributesWithStickerCustomEmoji()
	raw, err := encodeDocumentAttributesForStorage(attrs)
	if err != nil {
		t.Fatalf("encodeDocumentAttributesForStorage() error = %v", err)
	}
	var envelope documentAttributeStorageEnvelope
	if err := json.Unmarshal([]byte(raw), &envelope); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if envelope.SchemaVersion != documentAttributeStorageVersionV2 {
		t.Fatalf("SchemaVersion = %d, want %d", envelope.SchemaVersion, documentAttributeStorageVersionV2)
	}
	if envelope.Encoding != documentAttributeStorageTLVector {
		t.Fatalf("Encoding = %q, want %q", envelope.Encoding, documentAttributeStorageTLVector)
	}
	if envelope.Layer != documentAttributeVectorLayer {
		t.Fatalf("Layer = %d, want %d", envelope.Layer, documentAttributeVectorLayer)
	}
	if _, err := base64.StdEncoding.DecodeString(envelope.Bytes); err != nil {
		t.Fatalf("base64.DecodeString() error = %v", err)
	}
	decodedVector, err := decodeDocumentAttributesFromStorage(raw)
	if err != nil {
		t.Fatalf("decodeDocumentAttributesFromStorage() error = %v", err)
	}
	assertRepositoryDocumentAttributesLossless(t, decodedVector)
}

func TestDocumentAttributePersistencePreservesStickerCustomEmojiAndHasStickers(t *testing.T) {
	raw, err := encodeDocumentAttributesForStorage(repositoryDocumentAttributesWithStickerCustomEmoji())
	if err != nil {
		t.Fatalf("encodeDocumentAttributesForStorage() error = %v", err)
	}
	got, err := mapDocumentAggregate(
		&model.Documents{
			DocumentId: 20,
			AccessHash: 30,
			DcId:       1,
			Date2:      40,
			MimeType:   "image/webp",
			FileSize:   50,
			Attributes: raw,
		},
		nil,
		nil,
		[]byte("file-reference"),
	)
	if err != nil {
		t.Fatalf("mapDocumentAggregate() error = %v", err)
	}
	doc, ok := got.ToDocument()
	if !ok {
		t.Fatalf("Document clazz = %s, want document", got.ClazzName())
	}
	assertRepositoryDocumentAttributesLossless(t, doc.Attributes)
}

func TestDecodeDocumentAttributesFromStorageAcceptsLegacyJSON(t *testing.T) {
	decodedLegacy, err := decodeDocumentAttributesFromStorage(`[{"_name":"documentAttributeFilename","_object":{"file_name":"report.pdf"}}]`)
	if err != nil {
		t.Fatalf("decodeDocumentAttributesFromStorage() error = %v", err)
	}
	filename, hasFilename := findRepositoryDocumentAttribute[*tg.TLDocumentAttributeFilename](decodedLegacy)
	if !hasFilename {
		t.Fatalf("decoded legacy attrs = %#v, want filename", decodedLegacy)
	}
	if filename.FileName != "report.pdf" {
		t.Fatalf("filename attr FileName = %q, want report.pdf", filename.FileName)
	}
}

func TestDecodeDocumentAttributesFromStorageRejectsMalformedEnvelope(t *testing.T) {
	tests := []struct {
		name string
		raw  string
	}{
		{
			name: "invalid JSON",
			raw:  `{`,
		},
		{
			name: "unsupported schema",
			raw:  `{"schema_version":1,"encoding":"tl_object_vector","layer":224,"bytes":""}`,
		},
		{
			name: "unsupported encoding",
			raw:  `{"schema_version":2,"encoding":"json","layer":224,"bytes":""}`,
		},
		{
			name: "invalid base64",
			raw:  `{"schema_version":2,"encoding":"tl_object_vector","layer":224,"bytes":"%%%"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := decodeDocumentAttributesFromStorage(tt.raw)
			if !errors.Is(err, media.ErrMediaStorage) {
				t.Fatalf("decodeDocumentAttributesFromStorage() error = %v, want ErrMediaStorage", err)
			}
		})
	}
}

func TestSaveDocumentAggregateStoresAttributeEnvelope(t *testing.T) {
	documents := &captureDocumentsModel{}
	repo := &Repository{model: &model.Models{DocumentsModel: documents}}
	doc := tg.MakeTLDocument(&tg.TLDocument{
		Id:         20,
		AccessHash: 30,
		DcId:       1,
		Date:       40,
		MimeType:   "image/webp",
		Size2:      50,
		Attributes: repositoryDocumentAttributesWithStickerCustomEmoji(),
	}).ToDocument()
	if err := repo.saveDocumentAggregateWithPaths(context.Background(), "sticker.webp", doc, "document-object", nil); err != nil {
		t.Fatalf("saveDocumentAggregateWithPaths() error = %v", err)
	}
	if len(documents.inserted) != 1 {
		t.Fatalf("Insert2() rows = %d, want 1", len(documents.inserted))
	}
	row := documents.inserted[0]
	var envelope documentAttributeStorageEnvelope
	if err := json.Unmarshal([]byte(row.Attributes), &envelope); err != nil {
		t.Fatalf("json.Unmarshal(row.Attributes) error = %v", err)
	}
	if envelope.SchemaVersion != documentAttributeStorageVersionV2 || envelope.Encoding != documentAttributeStorageTLVector || envelope.Layer != documentAttributeVectorLayer {
		t.Fatalf("envelope = %#v, want v2 tl vector layer", envelope)
	}
	decoded, err := decodeDocumentAttributesFromStorage(row.Attributes)
	if err != nil {
		t.Fatalf("decodeDocumentAttributesFromStorage() error = %v", err)
	}
	assertRepositoryDocumentAttributesLossless(t, decoded)
}

func repositoryDocumentAttributesWithStickerCustomEmoji() []tg.DocumentAttributeClazz {
	return []tg.DocumentAttributeClazz{
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
}

func assertRepositoryDocumentAttributesLossless(t *testing.T, attrs []tg.DocumentAttributeClazz) {
	t.Helper()
	filename, hasFilename := findRepositoryDocumentAttribute[*tg.TLDocumentAttributeFilename](attrs)
	sticker, hasSticker := findRepositoryDocumentAttribute[*tg.TLDocumentAttributeSticker](attrs)
	customEmoji, hasCustomEmoji := findRepositoryDocumentAttribute[*tg.TLDocumentAttributeCustomEmoji](attrs)
	_, hasStickers := findRepositoryDocumentAttribute[*tg.TLDocumentAttributeHasStickers](attrs)
	if !hasFilename || !hasSticker || !hasCustomEmoji || !hasStickers {
		t.Fatalf("attrs = %#v, want filename/sticker/custom_emoji/has_stickers", attrs)
	}
	if filename.FileName != "sticker.webp" {
		t.Fatalf("filename attr FileName = %q, want sticker.webp", filename.FileName)
	}
	stickerSet, ok := sticker.Stickerset.(*tg.TLInputStickerSetID)
	if !ok || stickerSet.Id != 1001 || stickerSet.AccessHash != 2002 {
		t.Fatalf("sticker stickerset = %#v, want inputStickerSetID 1001/2002", sticker.Stickerset)
	}
	maskCoords := sticker.MaskCoords
	if maskCoords == nil || maskCoords.N != 1 || maskCoords.X != 0.5 || maskCoords.Y != 0.25 || maskCoords.Zoom != 1.5 {
		t.Fatalf("sticker mask coords = %#v, want exact TLMaskCoords", sticker.MaskCoords)
	}
	if sticker.Alt != ":)" || !sticker.Mask {
		t.Fatalf("sticker attr = %#v, want alt and mask preserved", sticker)
	}
	customStickerSet, ok := customEmoji.Stickerset.(*tg.TLInputStickerSetID)
	if !ok || customStickerSet.Id != 3003 || customStickerSet.AccessHash != 4004 {
		t.Fatalf("custom emoji stickerset = %#v, want inputStickerSetID 3003/4004", customEmoji.Stickerset)
	}
	if customEmoji.Alt != ":)" || !customEmoji.Free || !customEmoji.TextColor {
		t.Fatalf("custom emoji attr = %#v, want alt/free/text_color preserved", customEmoji)
	}
}

func findRepositoryDocumentAttribute[T tg.DocumentAttributeClazz](attrs []tg.DocumentAttributeClazz) (T, bool) {
	for _, attr := range attrs {
		if got, ok := attr.(T); ok {
			return got, true
		}
	}
	var zero T
	return zero, false
}
