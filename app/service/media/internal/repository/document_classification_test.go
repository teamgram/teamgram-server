package repository

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestClassifyUploadedDocumentCustomEmojiRejectsWrongMime(t *testing.T) {
	_, err := classifyUploadedDocument(uploadedDocumentForClassification(
		"image/png",
		tg.MakeTLDocumentAttributeCustomEmoji(&tg.TLDocumentAttributeCustomEmoji{Alt: "x"}),
	))
	if !errors.Is(err, media.ErrMediaInvalidUploadedFile) {
		t.Fatalf("expected ErrMediaInvalidUploadedFile, got %v", err)
	}
}

func TestClassifyUploadedDocumentWebMStickerDoesNotInferVideoRenderFlag(t *testing.T) {
	got, err := classifyUploadedDocument(uploadedDocumentForClassification(
		"video/webm",
		tg.MakeTLDocumentAttributeSticker(&tg.TLDocumentAttributeSticker{Alt: "x"}),
	))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Kind != documentKindSticker || got.RenderKind != messageRenderKindFile || got.Video {
		t.Fatalf("unexpected classification: %#v", got)
	}
}

func TestClassifyUploadedDocumentWebMWithVideoAttributeRejectsWrongMime(t *testing.T) {
	_, err := classifyUploadedDocument(uploadedDocumentForClassification(
		"video/webm",
		tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{Duration: 3, W: 640, H: 360}),
	))
	if !errors.Is(err, media.ErrMediaInvalidUploadedFile) {
		t.Fatalf("expected ErrMediaInvalidUploadedFile, got %v", err)
	}
}

func TestClassifyUploadedDocumentGifvRequiresMp4Transform(t *testing.T) {
	got, err := classifyUploadedDocument(uploadedDocumentForClassification(
		" image/gif ",
		tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}),
	))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Kind != documentKindGifv || got.RenderKind != messageRenderKindVideo || got.RequiredTransform != requiredDocumentTransformGifv || !got.Video {
		t.Fatalf("unexpected classification: %#v", got)
	}
}

func TestClassifyUploadedDocumentMp4RequiresMp4Probe(t *testing.T) {
	got, err := classifyUploadedDocument(uploadedDocumentForClassification("VIDEO/MP4"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Kind != documentKindVideo || got.RenderKind != messageRenderKindVideo || got.RequiredTransform != requiredDocumentTransformMp4 || !got.Video {
		t.Fatalf("unexpected classification: %#v", got)
	}
}

func TestClassifyUploadedDocumentForceFileSuppressesRenderFlagsOnly(t *testing.T) {
	tests := []struct {
		name     string
		mimeType string
		attrs    []tg.DocumentAttributeClazz
		wantKind documentKind
	}{
		{
			name:     "voice",
			mimeType: "audio/ogg",
			attrs: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeAudio(&tg.TLDocumentAttributeAudio{Voice: true, Duration: 3}),
			},
			wantKind: documentKindVoice,
		},
		{
			name:     "round",
			mimeType: "video/mp4",
			attrs: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{RoundMessage: true, Duration: 3, W: 640, H: 640}),
			},
			wantKind: documentKindRound,
		},
		{
			name:     "video",
			mimeType: "video/mp4",
			attrs: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeVideo(&tg.TLDocumentAttributeVideo{Duration: 3, W: 640, H: 360}),
			},
			wantKind: documentKindVideo,
		},
		{
			name:     "gifv",
			mimeType: "image/gif",
			attrs: []tg.DocumentAttributeClazz{
				tg.MakeTLDocumentAttributeAnimated(&tg.TLDocumentAttributeAnimated{}),
			},
			wantKind: documentKindGifv,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uploaded := uploadedDocumentForClassification(tt.mimeType, tt.attrs...)
			uploaded.ForceFile = true
			got, err := classifyUploadedDocument(uploaded)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.Kind != tt.wantKind || got.RenderKind != messageRenderKindFile || got.Video || got.Round || got.Voice {
				t.Fatalf("unexpected force-file classification: %#v", got)
			}
		})
	}
}

func uploadedDocumentForClassification(mimeType string, attrs ...tg.DocumentAttributeClazz) *tg.TLInputMediaUploadedDocument {
	return tg.MakeTLInputMediaUploadedDocument(&tg.TLInputMediaUploadedDocument{
		MimeType:   mimeType,
		Attributes: attrs,
	})
}
