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
				{Kind: "size", Type: "v", W: 320, H: 200, Size: 4567, VideoStartTs: &videoStartTs},
			},
			DocumentAttributes: []payload.DocumentAttributeRefV1{
				{Kind: "filename", FileName: "clip.mp4"},
				{Kind: "video", W: 1280, H: 720, DurationFloat: 3.5, SupportsStreaming: true, VideoStartTs: &videoStartTs},
				{Kind: "sticker", Alt: ":)", StickerSetKind: "id", StickerSetID: 1001, StickerSetAccessHash: 2002, Mask: true, MaskCoords: &payload.MaskCoordsRefV1{N: 1, X: 0.5, Y: 0.25, Zoom: 1.5}},
				{Kind: "custom_emoji", Alt: ":)", StickerSetKind: "id", StickerSetID: 3003, StickerSetAccessHash: 4004, Free: true, TextColor: true},
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
	videoThumb, ok := doc.VideoThumbs[0].(*tg.TLVideoSize)
	if !ok {
		t.Fatalf("VideoThumbs[0] = %T, want *tg.TLVideoSize", doc.VideoThumbs[0])
	}
	if videoThumb.Type != "v" || videoThumb.W != 320 || videoThumb.H != 200 || videoThumb.Size2 != 4567 {
		t.Fatalf("VideoThumbs[0] = %#v, want videoSize v 320x200/4567", videoThumb)
	}
	if videoThumb.VideoStartTs == nil || *videoThumb.VideoStartTs != videoStartTs {
		t.Fatalf("VideoThumbs[0].VideoStartTs = %v, want %v", videoThumb.VideoStartTs, videoStartTs)
	}
	assertCoreProjectedDocumentAttributes(t, doc.Attributes, videoStartTs)
	if media.VideoTimestamp == nil || *media.VideoTimestamp != videoTimestamp {
		t.Fatalf("VideoTimestamp = %v, want %d", media.VideoTimestamp, videoTimestamp)
	}
	videoCover, ok := media.VideoCover.(*tg.TLPhoto)
	if !ok {
		t.Fatalf("VideoCover = %T, want *tg.TLPhoto", media.VideoCover)
	}
	assertCoreProjectedVideoCover(t, videoCover)
}

func assertCoreProjectedDocumentAttributes(t *testing.T, attrs []tg.DocumentAttributeClazz, videoStartTs float64) {
	t.Helper()
	filename, hasFilename := findCoreProjectionDocumentAttribute[*tg.TLDocumentAttributeFilename](attrs)
	video, hasVideo := findCoreProjectionDocumentAttribute[*tg.TLDocumentAttributeVideo](attrs)
	sticker, hasSticker := findCoreProjectionDocumentAttribute[*tg.TLDocumentAttributeSticker](attrs)
	customEmoji, hasCustomEmoji := findCoreProjectionDocumentAttribute[*tg.TLDocumentAttributeCustomEmoji](attrs)
	_, hasStickers := findCoreProjectionDocumentAttribute[*tg.TLDocumentAttributeHasStickers](attrs)
	if !hasFilename || !hasVideo || !hasSticker || !hasCustomEmoji || !hasStickers {
		t.Fatalf("document attrs = %#v, want filename/video/sticker/custom_emoji/has_stickers", attrs)
	}
	if filename.FileName != "clip.mp4" {
		t.Fatalf("filename attr FileName = %q, want clip.mp4", filename.FileName)
	}
	if video.Duration != 3.5 || video.W != 1280 || video.H != 720 || !video.SupportsStreaming {
		t.Fatalf("video attr = %#v, want duration/w/h/supports_streaming preserved", video)
	}
	if video.VideoStartTs == nil || *video.VideoStartTs != videoStartTs {
		t.Fatalf("video attr VideoStartTs = %v, want %v", video.VideoStartTs, videoStartTs)
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

func assertCoreProjectedVideoCover(t *testing.T, cover *tg.TLPhoto) {
	t.Helper()
	if cover.Id != 777 || cover.AccessHash != 888 || string(cover.FileReference) != "cover-ref" || cover.Date != 1700000001 || cover.DcId != 5 {
		t.Fatalf("VideoCover = %#v, want full photo 777", cover)
	}
	if len(cover.Sizes) != 1 {
		t.Fatalf("VideoCover.Sizes len = %d, want 1", len(cover.Sizes))
	}
	size, ok := cover.Sizes[0].(*tg.TLPhotoSize)
	if !ok {
		t.Fatalf("VideoCover.Sizes[0] = %T, want *tg.TLPhotoSize", cover.Sizes[0])
	}
	if size.Type != "x" || size.W != 640 || size.H != 360 || size.Size2 != 4321 {
		t.Fatalf("VideoCover.Sizes[0] = %#v, want photoSize x 640x360/4321", size)
	}
}

func findCoreProjectionDocumentAttribute[T tg.DocumentAttributeClazz](attrs []tg.DocumentAttributeClazz) (T, bool) {
	for _, attr := range attrs {
		if got, ok := attr.(T); ok {
			return got, true
		}
	}
	var zero T
	return zero, false
}
