package core

import (
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestMessageEventV3ToTLMessageProjectsFullUploadedDocumentContract(t *testing.T) {
	videoStartTs := 1.25
	videoTimestamp := int32(7)
	got, err := messageEventV3ToTLMessage(payload.MessageEventV3{
		SchemaVersion:      payload.MessageEventSchemaVersionV3,
		EventKind:          payload.EventKindNewMessage,
		CanonicalMessageID: 103,
		PeerSeq:            11,
		MessageID:          79,
		PeerType:           payload.PeerTypeChat,
		PeerID:             202,
		FromUserID:         101,
		ToUserID:           202,
		Date:               1700000000,
		Out:                true,
		MessageText:        "document",
		MediaRef: &payload.MediaRefV1{
			SchemaVersion: payload.MediaRefSchemaVersionV2,
			Kind:          "document",
			ID:            555,
			AccessHash:    666,
			FileReference: []byte("doc-ref"),
			Date:          1700000000,
			DcID:          4,
			Size:          98765,
			MimeType:      "video/mp4",
			DocumentVideoThumbs: []payload.VideoSizeRefV1{
				{Type: "v", W: 320, H: 200, Size: 4567, VideoStartTs: &videoStartTs},
			},
			DocumentAttributes: []payload.DocumentAttributeRefV1{
				{Kind: "filename", FileName: "clip.mp4"},
				{Kind: "video", W: 1280, H: 720, DurationFloat: 3.5, SupportsStreaming: true, VideoStartTs: &videoStartTs},
				{Kind: "sticker", Alt: ":)", StickerSetID: 1001, StickerSetAccessHash: 2002, MaskCoords: &payload.MaskCoordsRefV1{N: 1, X: 0.5, Y: 0.25, Zoom: 1.5}},
				{Kind: "custom_emoji", Alt: ":)", StickerSetID: 3003, StickerSetAccessHash: 4004, Free: true, TextColor: true},
				{Kind: "has_stickers"},
			},
			DocumentMediaFlags: payload.DocumentMediaFlagsV1{Video: true, Spoiler: true},
			VideoCover: &payload.PhotoRefV1{
				ID:            777,
				AccessHash:    888,
				FileReference: []byte("cover-ref"),
				Date:          1700000001,
				DcID:          5,
				Sizes: []payload.PhotoSizeRefV1{
					{Kind: "size", Type: "x", W: 640, H: 360, Size: 4321},
				},
			},
			VideoTimestamp: &videoTimestamp,
		},
	})
	if err != nil {
		t.Fatalf("messageEventV3ToTLMessage() error = %v", err)
	}
	message, ok := got.(*tg.TLMessage)
	if !ok {
		t.Fatalf("message = %T, want *tg.TLMessage", got)
	}
	media, ok := message.Media.(*tg.TLMessageMediaDocument)
	if !ok {
		t.Fatalf("media = %T, want *tg.TLMessageMediaDocument", message.Media)
	}
	if !media.Video || !media.Spoiler {
		t.Fatalf("messageMediaDocument flags video=%v spoiler=%v, want true/true", media.Video, media.Spoiler)
	}
	doc, ok := media.Document.(*tg.TLDocument)
	if !ok {
		t.Fatalf("document = %T, want *tg.TLDocument", media.Document)
	}
	if len(doc.VideoThumbs) != 1 {
		t.Fatalf("VideoThumbs len = %d, want 1", len(doc.VideoThumbs))
	}
	if _, ok := doc.VideoThumbs[0].(*tg.TLVideoSize); !ok {
		t.Fatalf("VideoThumbs[0] = %T, want *tg.TLVideoSize", doc.VideoThumbs[0])
	}
	if !hasCoreProjectionDocumentAttribute[*tg.TLDocumentAttributeSticker](doc.Attributes) ||
		!hasCoreProjectionDocumentAttribute[*tg.TLDocumentAttributeCustomEmoji](doc.Attributes) ||
		!hasCoreProjectionDocumentAttribute[*tg.TLDocumentAttributeHasStickers](doc.Attributes) {
		t.Fatalf("document attrs = %#v, want sticker/custom_emoji/has_stickers", doc.Attributes)
	}
	if media.VideoTimestamp == nil || *media.VideoTimestamp != videoTimestamp {
		t.Fatalf("VideoTimestamp = %v, want %d", media.VideoTimestamp, videoTimestamp)
	}
	if _, ok := media.VideoCover.(*tg.TLPhoto); !ok {
		t.Fatalf("VideoCover = %T, want *tg.TLPhoto", media.VideoCover)
	}
}

func hasCoreProjectionDocumentAttribute[T tg.DocumentAttributeClazz](attrs []tg.DocumentAttributeClazz) bool {
	for _, attr := range attrs {
		if _, ok := attr.(T); ok {
			return true
		}
	}
	return false
}
