package bffproxyclient

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestRawFakeReturnsEncodedTLBytes(t *testing.T) {
	payload, ok, err := TryReturnRawFakeRpcResult(context.Background(), nil, encodeRawFakeTL(t, &tg.TLHelpGetTermsOfServiceUpdate{}))
	if err != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() error = %v", err)
	}
	if !ok {
		t.Fatal("TryReturnRawFakeRpcResult() ok = false")
	}
	obj, err := iface.DecodeObject(bin.NewDecoder(payload))
	if err != nil {
		t.Fatalf("DecodeObject() error = %v", err)
	}
	if _, ok := obj.(*tg.TLHelpTermsOfServiceUpdateEmpty); !ok {
		t.Fatalf("DecodeObject() = %T, want *tg.TLHelpTermsOfServiceUpdateEmpty", obj)
	}
}

func TestRawFakeDecodesRequestFields(t *testing.T) {
	payload, ok, err := TryReturnRawFakeRpcResult(context.Background(), nil, encodeRawFakeTL(t, &tg.TLLangpackGetDifference{
		LangPack:    "tdesktop",
		LangCode:    "zh-hans",
		FromVersion: 7,
	}))
	if err != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() error = %v", err)
	}
	if !ok {
		t.Fatal("TryReturnRawFakeRpcResult() ok = false")
	}
	obj, err := iface.DecodeObject(bin.NewDecoder(payload))
	if err != nil {
		t.Fatalf("DecodeObject() error = %v", err)
	}
	diff, ok := obj.(*tg.TLLangPackDifference)
	if !ok {
		t.Fatalf("DecodeObject() = %T, want *tg.TLLangPackDifference", obj)
	}
	if diff.LangCode != "zh-hans" || diff.FromVersion != 7 || diff.Version != 7 {
		t.Fatalf("difference = %#v", diff)
	}
}

func TestRawFakeUnknownConstructor(t *testing.T) {
	x := bin.NewEncoder()
	x.PutClazzID(0xfeed9999)
	payload, ok, err := TryReturnRawFakeRpcResult(context.Background(), nil, append([]byte(nil), x.Bytes()...))
	x.End()
	if err != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() error = %v", err)
	}
	if ok || payload != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() = %x, %v; want nil, false", payload, ok)
	}
}

func encodeRawFakeTL(t *testing.T, obj interface {
	Encode(*bin.Encoder, int32) error
}) []byte {
	t.Helper()
	x := bin.NewEncoder()
	defer x.End()
	if err := obj.Encode(x, 224); err != nil {
		x.Reset()
		if err2 := obj.Encode(x, 0); err2 != nil {
			t.Fatalf("Encode(%T) error = %v", obj, err)
		}
	}
	return append([]byte(nil), x.Bytes()...)
}
