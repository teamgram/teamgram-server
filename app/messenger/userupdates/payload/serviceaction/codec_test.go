package serviceaction

import (
	"fmt"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestEncodeDecodeRoundTripChatEditTitle(t *testing.T) {
	ref, err := Encode(tg.MakeTLMessageActionChatEditTitle(&tg.TLMessageActionChatEditTitle{Title: "renamed"}))
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	if ref.SchemaVersion != payload.ServiceActionSchemaVersionV1 || ref.Codec != payload.ServiceActionCodecTLBinary || ref.Layer != payload.ServiceActionLayer || len(ref.ActionPayload) == 0 {
		t.Fatalf("ref = %+v, want TL-binary service action ref", ref)
	}
	got, err := Decode(ref)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	title, ok := got.(*tg.TLMessageActionChatEditTitle)
	if !ok || title.Title != "renamed" {
		t.Fatalf("decoded action = %T %+v, want chat edit title", got, got)
	}
}

func TestEncodeDecodeRoundTripPhoneCallPreservesFlags(t *testing.T) {
	duration := int32(42)
	ref, err := Encode(tg.MakeTLMessageActionPhoneCall(&tg.TLMessageActionPhoneCall{
		Video:    true,
		CallId:   9001,
		Duration: &duration,
	}))
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	got, err := Decode(ref)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	call, ok := got.(*tg.TLMessageActionPhoneCall)
	if !ok || !call.Video || call.CallId != 9001 || call.Duration == nil || *call.Duration != 42 {
		t.Fatalf("decoded call = %T %+v, want video call duration", got, got)
	}
}

func TestDecodeRejectsNonMessageActionPayload(t *testing.T) {
	ref, err := encodeObjectForTest(tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1001}))
	if err != nil {
		t.Fatalf("encodeObjectForTest() error = %v", err)
	}
	_, err = Decode(ref)
	if err == nil {
		t.Fatal("Decode() error = nil, want non-message-action failure")
	}
}

func TestDecodeRejectsUnsupportedCodec(t *testing.T) {
	ref, err := Encode(tg.MakeTLMessageActionEmpty(&tg.TLMessageActionEmpty{}))
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	ref.Codec = 99
	_, err = Decode(ref)
	if err == nil {
		t.Fatal("Decode() error = nil, want unsupported codec failure")
	}
}

func TestDecodeRejectsTrailingBytes(t *testing.T) {
	ref, err := Encode(tg.MakeTLMessageActionEmpty(&tg.TLMessageActionEmpty{}))
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	ref.ActionPayload = append(ref.ActionPayload, 0)
	_, err = Decode(ref)
	if err == nil {
		t.Fatal("Decode() error = nil, want trailing bytes failure")
	}
}

func encodeObjectForTest(obj iface.TLObject) (*payload.ServiceActionRefV1, error) {
	body, err := iface.EncodeObject(obj, payload.ServiceActionLayer)
	if err != nil {
		return nil, fmt.Errorf("encode object for test: %w", err)
	}
	return &payload.ServiceActionRefV1{
		SchemaVersion: payload.ServiceActionSchemaVersionV1,
		Codec:         payload.ServiceActionCodecTLBinary,
		Layer:         payload.ServiceActionLayer,
		ActionPayload: append([]byte(nil), body...),
	}, nil
}
