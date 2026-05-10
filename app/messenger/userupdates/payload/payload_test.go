package payload

import (
	"encoding/hex"
	"encoding/json"
	"math"
	"strings"
	"testing"
)

func TestHashBytesUsesSHA256RawDigest(t *testing.T) {
	got := HashBytes([]byte("teamgram"))
	const want = "bb142b781fe5b7b5679b8a3417a26f967aab3e054f836d07872d2c7d1a686547"
	if hex.EncodeToString(got) != want {
		t.Fatalf("HashBytes() = %x, want %s", got, want)
	}
	if len(got) != 32 {
		t.Fatalf("HashBytes() len = %d, want 32", len(got))
	}
}

func TestHashHexUsesSHA256Hex(t *testing.T) {
	got := HashHex([]byte("teamgram"))
	const want = "bb142b781fe5b7b5679b8a3417a26f967aab3e054f836d07872d2c7d1a686547"
	if got != want {
		t.Fatalf("HashHex() = %q, want %q", got, want)
	}
}

func TestRoutingUsesStableBucketAndPartitionCounts(t *testing.T) {
	if BucketCount != 4096 {
		t.Fatalf("BucketCount = %d, want 4096", BucketCount)
	}
	if ReceiverPartitionCount != 256 {
		t.Fatalf("ReceiverPartitionCount = %d, want 256", ReceiverPartitionCount)
	}
	if PushPartitionCount != 256 {
		t.Fatalf("PushPartitionCount = %d, want 256", PushPartitionCount)
	}

	first := RouteUser(123456789)
	second := RouteUser(123456789)
	if first != second {
		t.Fatalf("RouteUser not deterministic: first=%+v second=%+v", first, second)
	}
	if first.BucketID < 0 || first.BucketID >= BucketCount {
		t.Fatalf("bucket out of range: %+v", first)
	}
	if first.ReceiverPartitionID < 0 || first.ReceiverPartitionID >= ReceiverPartitionCount {
		t.Fatalf("receiver partition out of range: %+v", first)
	}
	if first.PushPartitionID < 0 || first.PushPartitionID >= PushPartitionCount {
		t.Fatalf("push partition out of range: %+v", first)
	}
}

func TestOperationIDBuildersAreStableAndFitStorageColumn(t *testing.T) {
	sender := SenderOperationID(math.MaxInt64, math.MaxInt64)
	if !strings.HasPrefix(sender, "v1:msg:") {
		t.Fatalf("sender operation id missing version prefix: %q", sender)
	}
	if len(sender) >= MaxOperationIDLength {
		t.Fatalf("sender operation id length = %d, want < %d", len(sender), MaxOperationIDLength)
	}

	receiver := ReceiverOperationID(math.MaxInt64, math.MaxInt64)
	if !strings.HasPrefix(receiver, "v1:msg:") {
		t.Fatalf("receiver operation id missing version prefix: %q", receiver)
	}
	if len(receiver) >= MaxOperationIDLength {
		t.Fatalf("receiver operation id length = %d, want < %d", len(receiver), MaxOperationIDLength)
	}
}

func TestMessageOperationV1UsesServerOwnedJSON(t *testing.T) {
	op := MessageOperationV1{
		SchemaVersion:      MessageOperationSchemaVersion,
		OperationKind:      OperationKindSendMessage,
		CanonicalMessageID: 1001,
		PeerType:           PeerTypeUser,
		PeerID:             2002,
		PeerSeq:            3,
		FromUserID:         100,
		ToUserID:           200,
		Date:               1710000000,
		Out:                false,
		MessageText:        "hello",
	}
	b, err := json.Marshal(op)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	text := string(b)
	if strings.Contains(text, "@type") || strings.Contains(text, "@id") || strings.Contains(text, "clazz") {
		t.Fatalf("payload JSON contains TL-like fields: %s", text)
	}
	if !strings.Contains(text, `"schema_version":2`) {
		t.Fatalf("payload JSON missing schema_version: %s", text)
	}

	var decoded MessageOperationV1
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if decoded.SchemaVersion != MessageOperationSchemaVersion || decoded.MessageText != "hello" {
		t.Fatalf("decoded operation mismatch: %+v", decoded)
	}
}

func TestMessageOperationV2DecodesWithCurrentFields(t *testing.T) {
	body := []byte(`{"schema_version":2,"operation_kind":"send_message","canonical_message_id":101,"peer_type":1,"peer_id":202,"peer_seq":7,"from_user_id":1010,"to_user_id":2020,"date":1700000000,"out":true,"message_text":"caption","reply_to_user_message_id":55}`)
	var got MessageOperationV1
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got.SchemaVersion != MessageOperationSchemaVersion || got.MessageText != "caption" || got.ReplyToUserMessageID != 55 {
		t.Fatalf("decoded V2 operation mismatch: %+v", got)
	}
}

func TestMessageOperationV3CarriesV2AndMediaFields(t *testing.T) {
	op := MessageOperationV3{
		SchemaVersion:        MessageOperationSchemaVersionV3,
		OperationKind:        OperationKindSendMessage,
		CanonicalMessageID:   101,
		PeerType:             PeerTypeUser,
		PeerID:               202,
		PeerSeq:              7,
		FromUserID:           1010,
		ToUserID:             2020,
		Date:                 1700000000,
		Out:                  true,
		MessageText:          "caption",
		ReplyToUserMessageID: 55,
		MediaRef:             &MediaRefV1{SchemaVersion: MediaRefSchemaVersionV1, Kind: "photo", ID: 333},
		Attrs:                &MessageAttrsV1{SchemaVersion: MessageAttrsSchemaVersionV1, GroupedID: 444, Silent: true},
		ForwardRef:           &ForwardRefV1{SchemaVersion: ForwardRefSchemaVersionV1, FromUserID: 3030, Date: 1700000001, SourceMessageID: 66},
	}
	body, err := json.Marshal(op)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	text := string(body)
	if strings.Contains(text, "@type") || strings.Contains(text, "@id") || strings.Contains(text, "clazz") {
		t.Fatalf("payload JSON contains TL-like fields: %s", text)
	}
	var got MessageOperationV3
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got.SchemaVersion != MessageOperationSchemaVersionV3 || got.MessageText != "caption" || got.ReplyToUserMessageID != 55 {
		t.Fatalf("decoded V3 missing V2 fields: %+v", got)
	}
	if got.MediaRef == nil || got.MediaRef.Kind != "photo" || got.Attrs == nil || got.Attrs.GroupedID != 444 || got.ForwardRef == nil || got.ForwardRef.FromUserID != 3030 {
		t.Fatalf("decoded V3 missing media/attrs/forward: %+v", got)
	}
}

func TestStickerSetRefV1CarriesDiceEmoticon(t *testing.T) {
	ref := StickerSetRefV1{Kind: "dice", Emoticon: "🎲"}
	body, err := json.Marshal(ref)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	if !strings.Contains(string(body), `"emoticon":"🎲"`) {
		t.Fatalf("payload missing emoticon: %s", body)
	}
	var got StickerSetRefV1
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got.Kind != "dice" || got.Emoticon != "🎲" {
		t.Fatalf("decoded sticker set = %+v, want dice emoticon", got)
	}
}

func TestMediaRefV2CarriesFullDocumentProjection(t *testing.T) {
	videoStartTs := 1.25
	videoTimestamp := int32(7)
	ref := MediaRefV1{
		SchemaVersion: MediaRefSchemaVersionV2,
		Kind:          "document",
		ID:            555,
		AccessHash:    666,
		FileReference: []byte("doc-ref"),
		Date:          1_700_000_000,
		DcID:          4,
		MimeType:      "video/mp4",
		Size:          98765,
		DocumentThumbs: []PhotoSizeRefV1{
			{Kind: "size", Type: "m", W: 320, H: 200, Size: 1234},
		},
		DocumentVideoThumbs: []VideoSizeRefV1{
			{Kind: "size", Type: "v", W: 320, H: 200, Size: 4567, VideoStartTs: &videoStartTs},
		},
		DocumentAttributes: []DocumentAttributeRefV1{
			{Kind: "filename", FileName: "clip.mp4"},
			{Kind: "video", W: 1280, H: 720, DurationFloat: 3.5, SupportsStreaming: true, VideoStartTs: &videoStartTs},
			{Kind: "audio", Duration: 3, Title: stringPtrForTest("title"), Performer: stringPtrForTest("performer"), Waveform: []byte{1, 2, 3}, Voice: true},
			{Kind: "sticker", Alt: ":)", StickerSet: &StickerSetRefV1{Kind: "id", ID: 1001, AccessHash: 2002}, Mask: true, MaskCoords: &MaskCoordsRefV1{N: 1, X: 0.5, Y: 0.25, Zoom: 1.5}},
			{Kind: "custom_emoji", Alt: ":)", StickerSet: &StickerSetRefV1{Kind: "id", ID: 3003, AccessHash: 4004}, Free: true, TextColor: true},
			{Kind: "has_stickers"},
		},
		DocumentMediaFlags: &DocumentMediaFlagsV1{Video: true, Spoiler: true},
		VideoCover: &PhotoRefV1{
			ID:            777,
			AccessHash:    888,
			FileReference: []byte("cover-ref"),
			Date:          1_700_000_001,
			DcID:          5,
			Sizes: []PhotoSizeRefV1{
				{Kind: "size", Type: "x", W: 640, H: 360, Size: 4321},
			},
		},
		VideoTimestamp: &videoTimestamp,
	}

	body, err := json.Marshal(ref)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	var got MediaRefV1
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got.SchemaVersion != MediaRefSchemaVersionV2 || got.Kind != "document" || got.ID != 555 || got.AccessHash != 666 ||
		string(got.FileReference) != "doc-ref" || got.Date != 1_700_000_000 || got.DcID != 4 || got.MimeType != "video/mp4" || got.Size != 98765 {
		t.Fatalf("document media core fields = %+v, want full document identity preserved", got)
	}
	if len(got.DocumentThumbs) != 1 {
		t.Fatalf("DocumentThumbs len = %d, want 1", len(got.DocumentThumbs))
	}
	thumb := got.DocumentThumbs[0]
	if thumb.Kind != "size" || thumb.Type != "m" || thumb.W != 320 || thumb.H != 200 || thumb.Size != 1234 {
		t.Fatalf("DocumentThumbs[0] = %+v, want size m 320x200/1234", thumb)
	}
	if len(got.DocumentAttributes) != 6 {
		t.Fatalf("DocumentAttributes len = %d, want 6", len(got.DocumentAttributes))
	}
	for i, want := range []string{"filename", "video", "audio", "sticker", "custom_emoji", "has_stickers"} {
		if got.DocumentAttributes[i].Kind != want {
			t.Fatalf("DocumentAttributes[%d].Kind = %q, want %q", i, got.DocumentAttributes[i].Kind, want)
		}
	}
	filename := got.DocumentAttributes[0]
	if filename.FileName != "clip.mp4" {
		t.Fatalf("filename attr FileName = %q, want clip.mp4", filename.FileName)
	}
	video := got.DocumentAttributes[1]
	if video.DurationFloat != 3.5 || video.W != 1280 || video.H != 720 || !video.SupportsStreaming {
		t.Fatalf("video attr = %+v, want duration/w/h/supports_streaming preserved", video)
	}
	if video.VideoStartTs == nil || *video.VideoStartTs != videoStartTs {
		t.Fatalf("video attr VideoStartTs = %v, want %v", video.VideoStartTs, videoStartTs)
	}
	audio := got.DocumentAttributes[2]
	if audio.Duration != 3 || audio.Title == nil || *audio.Title != "title" || audio.Performer == nil || *audio.Performer != "performer" || !audio.Voice {
		t.Fatalf("audio attr = %+v, want duration/title/performer/voice preserved", audio)
	}
	if string(audio.Waveform) != string([]byte{1, 2, 3}) {
		t.Fatalf("audio attr Waveform = %v, want [1 2 3]", audio.Waveform)
	}
	sticker := got.DocumentAttributes[3]
	if sticker.Alt != ":)" || sticker.StickerSet == nil || sticker.StickerSet.Kind != "id" || sticker.StickerSet.ID != 1001 || sticker.StickerSet.AccessHash != 2002 || !sticker.Mask {
		t.Fatalf("sticker attr = %+v, want alt/stickerset/mask preserved", sticker)
	}
	if sticker.MaskCoords == nil || sticker.MaskCoords.N != 1 || sticker.MaskCoords.X != 0.5 || sticker.MaskCoords.Y != 0.25 || sticker.MaskCoords.Zoom != 1.5 {
		t.Fatalf("sticker attr MaskCoords = %+v, want exact mask coords", sticker.MaskCoords)
	}
	customEmoji := got.DocumentAttributes[4]
	if customEmoji.Alt != ":)" || customEmoji.StickerSet == nil || customEmoji.StickerSet.Kind != "id" || customEmoji.StickerSet.ID != 3003 || customEmoji.StickerSet.AccessHash != 4004 || !customEmoji.Free || !customEmoji.TextColor {
		t.Fatalf("custom emoji attr = %+v, want alt/stickerset/free/text_color preserved", customEmoji)
	}
	if len(got.DocumentVideoThumbs) != 1 {
		t.Fatalf("DocumentVideoThumbs len = %d, want 1", len(got.DocumentVideoThumbs))
	}
	videoThumb := got.DocumentVideoThumbs[0]
	if videoThumb.Kind != "size" || videoThumb.Type != "v" || videoThumb.W != 320 || videoThumb.H != 200 || videoThumb.Size != 4567 {
		t.Fatalf("DocumentVideoThumbs[0] = %+v, want size v 320x200/4567", videoThumb)
	}
	if videoThumb.VideoStartTs == nil || *videoThumb.VideoStartTs != videoStartTs {
		t.Fatalf("DocumentVideoThumbs[0].VideoStartTs = %v, want %v", videoThumb.VideoStartTs, videoStartTs)
	}
	if got.DocumentMediaFlags == nil || !got.DocumentMediaFlags.Video || !got.DocumentMediaFlags.Spoiler {
		t.Fatalf("DocumentMediaFlags = %+v, want video spoiler", got.DocumentMediaFlags)
	}
	if got.VideoCover == nil || got.VideoCover.ID != 777 || got.VideoCover.AccessHash != 888 ||
		string(got.VideoCover.FileReference) != "cover-ref" || got.VideoCover.Date != 1_700_000_001 || got.VideoCover.DcID != 5 {
		t.Fatalf("VideoCover = %+v, want full photo 777", got.VideoCover)
	}
	if len(got.VideoCover.Sizes) != 1 {
		t.Fatalf("VideoCover.Sizes len = %d, want 1", len(got.VideoCover.Sizes))
	}
	coverSize := got.VideoCover.Sizes[0]
	if coverSize.Kind != "size" || coverSize.Type != "x" || coverSize.W != 640 || coverSize.H != 360 || coverSize.Size != 4321 {
		t.Fatalf("VideoCover.Sizes[0] = %+v, want size x 640x360/4321", coverSize)
	}
	if got.VideoTimestamp == nil || *got.VideoTimestamp != videoTimestamp {
		t.Fatalf("VideoTimestamp = %v, want %d", got.VideoTimestamp, videoTimestamp)
	}
}

func TestMessageOperationV1SideEffectFieldsAffectPayloadHash(t *testing.T) {
	base := MessageOperationV1{
		SchemaVersion:      MessageOperationSchemaVersion,
		OperationKind:      OperationKindSendMessage,
		CanonicalMessageID: 1001,
		PeerType:           PeerTypeUser,
		PeerID:             2002,
		PeerSeq:            3,
		FromUserID:         100,
		ToUserID:           200,
		Date:               1710000000,
		Out:                true,
		MessageText:        "hello",
	}
	withSideEffect := base
	withSideEffect.ClearDraft = true
	withSideEffect.SourcePermAuthKeyID = 9001
	withSideEffect.ClearDraftBeforeDate = 1709999990
	withSideEffect.SavedDialogSideEffect = true

	baseBody, err := json.Marshal(base)
	if err != nil {
		t.Fatalf("Marshal(base) error = %v", err)
	}
	sideEffectBody, err := json.Marshal(withSideEffect)
	if err != nil {
		t.Fatalf("Marshal(withSideEffect) error = %v", err)
	}
	text := string(sideEffectBody)
	for _, want := range []string{
		`"clear_draft":true`,
		`"source_perm_auth_key_id":9001`,
		`"clear_draft_before_date":1709999990`,
		`"saved_dialog_side_effect":true`,
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("side-effect payload missing %s: %s", want, text)
		}
	}
	if hex.EncodeToString(HashBytes(baseBody)) == hex.EncodeToString(HashBytes(sideEffectBody)) {
		t.Fatalf("side-effect fields did not affect payload hash")
	}
}

func stringPtrForTest(s string) *string {
	return &s
}

func TestOperationResponseAndEventCarrySchemaVersion(t *testing.T) {
	resp := OperationResponseV1{SchemaVersion: OperationResponseSchemaVersion, Pts: 1, PtsCount: 1}
	if resp.SchemaVersion != OperationResponseSchemaVersion {
		t.Fatalf("response schema version = %d, want %d", resp.SchemaVersion, OperationResponseSchemaVersion)
	}

	event := MessageEventV1{SchemaVersion: MessageEventSchemaVersion, CanonicalMessageID: 1, MessageText: "hello"}
	if event.SchemaVersion != MessageEventSchemaVersion {
		t.Fatalf("event schema version = %d, want %d", event.SchemaVersion, MessageEventSchemaVersion)
	}
}

func TestOperationResponseCarriesUserMessageID(t *testing.T) {
	resp := OperationResponseV2{
		SchemaVersion: OperationResponseSchemaVersion,
		OperationID:   "op",
		Pts:           11,
		PtsCount:      1,
		EventType:     EventKindNewMessage,
		UserMessageID: 101,
	}
	if resp.UserMessageID != 101 {
		t.Fatalf("user_message_id = %d, want 101", resp.UserMessageID)
	}
}

func TestMessageEventV2PublicIDs(t *testing.T) {
	event := MessageEventV2{
		SchemaVersion:        MessageEventSchemaVersion,
		EventKind:            EventKindNewMessage,
		CanonicalMessageID:   200,
		PeerSeq:              9,
		MessageID:            101,
		ReplyToUserMessageID: 88,
	}
	if event.MessageID == event.PeerSeq {
		t.Fatalf("message_id must be public id, got peer_seq=%d", event.PeerSeq)
	}
	if event.ReplyToUserMessageID != 88 {
		t.Fatalf("reply public id = %d, want 88", event.ReplyToUserMessageID)
	}
}

func TestMessageEventV2DeleteMessagesCarriesPublicIDs(t *testing.T) {
	event := MessageEventV2{
		SchemaVersion:        MessageEventSchemaVersion,
		EventKind:            OperationKindDeleteMessages,
		PeerType:             PeerTypeUser,
		PeerID:               1002,
		Date:                 1_772_000_000,
		DeleteUserMessageIDs: []int64{107, 108},
	}
	body, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}
	var got MessageEventV2
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("unmarshal event: %v", err)
	}
	if len(got.DeleteUserMessageIDs) != 2 || got.DeleteUserMessageIDs[0] != 107 || got.DeleteUserMessageIDs[1] != 108 {
		t.Fatalf("delete ids = %v", got.DeleteUserMessageIDs)
	}
}

func TestReceiverKafkaMessageV1RoundTrip(t *testing.T) {
	op := ReceiverOperationEnvelopeV1{
		UserID:       1001,
		BucketID:     9,
		PartitionID:  9,
		OperationID:  "v1:msg:2001:receiver:1001:in",
		OpType:       OpTypeSendMessage,
		PeerType:     PeerTypeUser,
		PeerID:       2002,
		PayloadCodec: PayloadCodecJSON,
		Payload:      []byte(`{"schema_version":1}`),
		PayloadHash:  HashBytes([]byte(`{"schema_version":1}`)),
	}
	msg := ReceiverKafkaMessageV1{
		SchemaVersion: ReceiverKafkaMessageSchemaVersion,
		Operation:     op,
		Attempt:       0,
	}
	body, err := MarshalReceiverKafkaMessage(msg)
	if err != nil {
		t.Fatalf("MarshalReceiverKafkaMessage() error = %v", err)
	}
	got, err := UnmarshalReceiverKafkaMessage(body)
	if err != nil {
		t.Fatalf("UnmarshalReceiverKafkaMessage() error = %v", err)
	}
	if got.SchemaVersion != ReceiverKafkaMessageSchemaVersion {
		t.Fatalf("schema version = %d", got.SchemaVersion)
	}
	if got.Operation.OperationID != op.OperationID {
		t.Fatalf("operation id = %q", got.Operation.OperationID)
	}
}

func TestReceiverKafkaMessageRejectsUnknownSchema(t *testing.T) {
	_, err := UnmarshalReceiverKafkaMessage([]byte(`{"schema_version":999}`))
	if err == nil {
		t.Fatal("expected unknown schema error")
	}
}

func TestPushTaskKafkaMessageV1RoundTrip(t *testing.T) {
	msg := PushTaskKafkaMessageV1{
		SchemaVersion: PushTaskKafkaMessageSchemaVersion,
		TaskID:        3001,
		UserID:        1001,
		Pts:           42,
		PushType:      1,
		PeerType:      PeerTypeUser,
		PeerID:        2002,
		OperationID:   "v1:msg:2001:receiver:1001:in",
		Payload:       []byte(`{"schema_version":1}`),
	}
	body, err := MarshalPushTaskKafkaMessage(msg)
	if err != nil {
		t.Fatalf("MarshalPushTaskKafkaMessage() error = %v", err)
	}
	got, err := UnmarshalPushTaskKafkaMessage(body)
	if err != nil {
		t.Fatalf("UnmarshalPushTaskKafkaMessage() error = %v", err)
	}
	if got.TaskID != msg.TaskID || got.UserID != msg.UserID || got.Pts != msg.Pts {
		t.Fatalf("push task = %+v", got)
	}
}
