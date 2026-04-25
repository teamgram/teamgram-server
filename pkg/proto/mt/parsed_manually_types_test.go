package mt

import (
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
