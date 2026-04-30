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
	if !strings.Contains(text, `"schema_version":1`) {
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

func TestOperationResponseAndEventCarrySchemaVersion(t *testing.T) {
	resp := OperationResponseV1{SchemaVersion: OperationResponseSchemaVersion, Pts: 1, PtsCount: 1}
	if resp.SchemaVersion != 1 {
		t.Fatalf("response schema version = %d, want 1", resp.SchemaVersion)
	}

	event := MessageEventV1{SchemaVersion: MessageEventSchemaVersion, CanonicalMessageID: 1, MessageText: "hello"}
	if event.SchemaVersion != 1 {
		t.Fatalf("event schema version = %d, want 1", event.SchemaVersion)
	}
}
