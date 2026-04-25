package codec

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
)

func TestRawTLObjectEncodeWritesPayload(t *testing.T) {
	payload := []byte{0xe2, 0x06, 0x05, 0x9f, 0x04, 0x00, 0x00, 0x00}
	obj := NewRawTLObject(payload)

	enc := bin.NewEncoder()
	defer enc.End()
	if err := obj.Encode(enc, 0); err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	if got := enc.Bytes(); string(got) != string(payload) {
		t.Fatalf("Encode() bytes = %x, want %x", got, payload)
	}
}

func TestRawTLObjectDecodeCapturesRemainingPayload(t *testing.T) {
	payload := []byte{0x1e, 0xa5, 0x3b, 0x2e, 0x01, 0x02, 0x03, 0x04}
	obj := NewRawTLObject(nil)

	dec := bin.NewDecoder(payload)
	if err := obj.Decode(dec); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	if got := obj.Payload; string(got) != string(payload) {
		t.Fatalf("Decode() payload = %x, want %x", got, payload)
	}
	if dec.Remaining() != 0 {
		t.Fatalf("Decode() left %d unread bytes, want 0", dec.Remaining())
	}
}

func TestRawTLObjectConstructorID(t *testing.T) {
	obj := NewRawTLObject([]byte{0xe2, 0x06, 0x05, 0x9f})

	got, err := obj.ConstructorID()
	if err != nil {
		t.Fatalf("ConstructorID() error = %v", err)
	}
	if got != 0x9f0506e2 {
		t.Fatalf("ConstructorID() = %#x, want %#x", got, uint32(0x9f0506e2))
	}
}

func TestRawTLObjectConstructorIDRejectsShortPayload(t *testing.T) {
	obj := NewRawTLObject([]byte{0x01, 0x02, 0x03})

	_, err := obj.ConstructorID()
	if !errors.Is(err, bin.ErrUnexpectedEOF) {
		t.Fatalf("ConstructorID() error = %v, want %v", err, bin.ErrUnexpectedEOF)
	}
}
