package iface

import (
	"errors"
	"strings"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
)

const (
	testVectorGoodClazzID = 0xfeed1001
	testVectorBadClazzID  = 0xfeed1002
)

type testVectorGoodClazz interface {
	TLObject
	testVectorGoodClazz()
}

type testVectorGood struct{}

func (*testVectorGood) Encode(*bin.Encoder, int32) error { return nil }
func (*testVectorGood) Decode(*bin.Decoder) error        { return nil }
func (*testVectorGood) testVectorGoodClazz()             {}

type testVectorBad struct{}

func (*testVectorBad) Encode(*bin.Encoder, int32) error { return nil }
func (*testVectorBad) Decode(*bin.Decoder) error        { return nil }

func init() {
	RegisterClazzID(testVectorGoodClazzID, func() TLObject { return &testVectorGood{} })
	RegisterClazzID(testVectorBadClazzID, func() TLObject { return &testVectorBad{} })
}

func TestDecodeObjectListReturnsErrorForWrongElementType(t *testing.T) {
	x := bin.NewEncoder()
	x.PutClazzID(ClazzID_vector)
	x.PutInt(1)
	x.PutClazzID(testVectorBadClazzID)

	_, err := DecodeObjectList[testVectorGoodClazz](bin.NewDecoder(x.Bytes()))
	if err == nil {
		t.Fatal("DecodeObjectList() error = nil, want wrong element type error")
	}
	if !strings.Contains(err.Error(), "vector[0]") {
		t.Fatalf("DecodeObjectList() error = %q, want element index", err)
	}
}

func TestDecodeBoolFalseConsumesConstructor(t *testing.T) {
	x := bin.NewEncoder()
	EncodeBool(x, false)
	x.PutInt32(42)

	d := bin.NewDecoder(x.Bytes())
	got, err := DecodeBool(d)
	if err != nil {
		t.Fatalf("DecodeBool() error = %v", err)
	}
	if got {
		t.Fatal("DecodeBool() = true, want false")
	}
	if d.Offset() != Size32 {
		t.Fatalf("DecodeBool() offset = %d, want %d", d.Offset(), Size32)
	}
	next, err := d.Int32()
	if err != nil {
		t.Fatalf("Int32() after DecodeBool() error = %v", err)
	}
	if next != 42 {
		t.Fatalf("Int32() after DecodeBool() = %d, want 42", next)
	}
}

func TestDecodeInt32ListRejectsImpossibleCountBeforeAllocation(t *testing.T) {
	x := bin.NewEncoder()
	x.PutClazzID(ClazzID_vector)
	x.PutInt(3)

	_, err := DecodeInt32List(bin.NewDecoder(x.Bytes()))
	if err == nil {
		t.Fatal("DecodeInt32List() error = nil, want invalid length")
	}
	var invalidLength *bin.InvalidLengthError
	if !errors.As(err, &invalidLength) {
		t.Fatalf("DecodeInt32List() error = %T %v, want InvalidLengthError", err, err)
	}
}
