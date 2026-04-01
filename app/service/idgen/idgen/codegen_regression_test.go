package idgen

import (
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

type badInputId struct{}

func (badInputId) Encode(x *bin.Encoder, layer int32) error { return errors.New("boom") }
func (badInputId) Decode(d *bin.Decoder) error              { return nil }
func (badInputId) InputIdClazzName() string                 { return ClazzName_inputId }

func TestTLIdValsDecodePropagatesVectorErrors(t *testing.T) {
	d := bin.NewDecoder([]byte{1, 2, 3, 4})
	obj := &TLIdVals{ClazzID: 0x1c3baa66}
	if err := obj.Decode(d); err == nil {
		t.Fatalf("expected decode error for malformed int64 vector")
	}
}

func TestTLIdgenGetNextIdValListDecodeRejectsWrongVectorMarker(t *testing.T) {
	x := bin.AcquireEncoderCap(8)
	defer x.Release()
	x.PutClazzID(0xaa85f137)
	x.PutClazzID(iface.ClazzID_boolTrue)

	d := bin.NewDecoder(x.Bytes())
	obj := &TLIdgenGetNextIdValList{}
	if err := obj.Decode(d); err == nil {
		t.Fatalf("expected decode error for non-vector marker")
	}
}

func TestTLIdgenGetNextIdValListEncodePropagatesChildErrors(t *testing.T) {
	obj := &TLIdgenGetNextIdValList{
		Id: []InputIdClazz{badInputId{}},
	}
	x := bin.AcquireEncoderCap(16)
	defer x.Release()

	if err := obj.Encode(x, 0); err == nil {
		t.Fatalf("expected child encode error")
	}
}
