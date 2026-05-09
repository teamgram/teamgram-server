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
