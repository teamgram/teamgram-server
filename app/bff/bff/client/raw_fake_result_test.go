package bffproxyclient

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestRawFakeReturnsEncodedTLBytes(t *testing.T) {
	payload, ok, err := TryReturnRawFakeRpcResult(context.Background(), nil, tg.ClazzID_help_getTermsOfServiceUpdate)
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

func TestRawFakeUnknownConstructor(t *testing.T) {
	payload, ok, err := TryReturnRawFakeRpcResult(context.Background(), nil, 0xfeed9999)
	if err != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() error = %v", err)
	}
	if ok || payload != nil {
		t.Fatalf("TryReturnRawFakeRpcResult() = %x, %v; want nil, false", payload, ok)
	}
}
