package mt

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
)

type failingTLObject struct {
	err error
}

func (f *failingTLObject) ClazzName() string { return "failing" }

func (f *failingTLObject) Encode(*bin.Encoder, int32) error { return f.err }

func (f *failingTLObject) Decode(*bin.Decoder) error { return f.err }

func TestTLMsgContainerEncodePropagatesChildError(t *testing.T) {
	container := &TLMsgContainer{
		Messages: []*TLMessage2{
			{
				MsgId:  1,
				Seqno:  1,
				Object: &failingTLObject{err: errors.New("encode boom")},
			},
		},
	}

	err := container.Encode(bin.NewEncoder(), 0)
	if err == nil {
		t.Fatal("expected encode error")
	}
	if !strings.Contains(err.Error(), "unable to encode msg_container: field messages element 0") {
		t.Fatalf("expected contextual msg_container error, got %q", err.Error())
	}
}

func TestTLMessage2DecodeWrapsBodyDecodeError(t *testing.T) {
	x := bin.NewEncoder()
	x.PutInt64(1)
	x.PutInt32(2)
	x.PutInt32(4)
	x.PutClazzID(0xffffffff)

	var m TLMessage2
	err := m.Decode(bin.NewDecoder(x.Bytes()))
	if err == nil {
		t.Fatal("expected decode error")
	}
	if !strings.Contains(err.Error(), "unable to decode message2: field body") {
		t.Fatalf("expected body decode context, got %q", err.Error())
	}
}

func TestTLMsgContainerDecodeRejectsInvalidLength(t *testing.T) {
	x := bin.NewEncoder()
	x.PutInt(-1)

	var m TLMsgContainer
	err := m.Decode(bin.NewDecoder(x.Bytes()))
	if err == nil {
		t.Fatal("expected invalid length error")
	}
	if !strings.Contains(err.Error(), "field messages length") {
		t.Fatalf("expected messages length context, got %q", err.Error())
	}
}

func TestTLMessageRawDataDecodeConsumesDeclaredBody(t *testing.T) {
	body := bin.NewEncoder()
	body.PutInt32(7)

	x := bin.NewEncoder()
	x.PutInt64(1)
	x.PutInt32(2)
	x.PutInt32(8)
	x.PutClazzID(0x11223344)
	x.Put(body.Bytes())
	x.PutInt32(99)

	d := bin.NewDecoder(x.Bytes())
	var m TLMessageRawData
	if err := m.Decode(d); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}
	if m.ClazzID != 0x11223344 {
		t.Fatalf("ClazzID = %#x, want %#x", m.ClazzID, uint32(0x11223344))
	}
	if !bytes.Equal(m.Body, body.Bytes()) {
		t.Fatalf("Body = %x, want %x", m.Body, body.Bytes())
	}
	if d.Offset() != 24 {
		t.Fatalf("decoder offset = %d, want 24", d.Offset())
	}
	if d.Remaining() != 4 {
		t.Fatalf("decoder remaining = %d, want 4", d.Remaining())
	}
}

func TestTLMessageRawDataEncodeWritesConstructorAndDerivedLength(t *testing.T) {
	body := bin.NewEncoder()
	body.PutInt32(7)

	x := bin.NewEncoder()
	m := &TLMessageRawData{
		MsgId:   1,
		Seqno:   2,
		ClazzID: 0x11223344,
		Body:    body.Bytes(),
	}
	if err := m.Encode(x, 0); err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	d := bin.NewDecoder(x.Bytes())
	if _, err := d.Int64(); err != nil {
		t.Fatalf("msg_id decode error = %v", err)
	}
	if _, err := d.Int32(); err != nil {
		t.Fatalf("seqno decode error = %v", err)
	}
	size, err := d.Int32()
	if err != nil {
		t.Fatalf("bytes decode error = %v", err)
	}
	if size != 8 {
		t.Fatalf("bytes = %d, want 8", size)
	}
	clazzID, err := d.ClazzID()
	if err != nil {
		t.Fatalf("clazz_id decode error = %v", err)
	}
	if clazzID != 0x11223344 {
		t.Fatalf("clazz_id = %#x, want %#x", clazzID, uint32(0x11223344))
	}
	if !bytes.Equal(d.Raw(), body.Bytes()) {
		t.Fatalf("body = %x, want %x", d.Raw(), body.Bytes())
	}
}
