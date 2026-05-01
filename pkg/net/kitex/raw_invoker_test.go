package kitex

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

type fakeRawClient struct {
	method      string
	reqPayload  []byte
	respPayload []byte
}

func (c *fakeRawClient) Call(ctx context.Context, method string, request, response interface{}) error {
	c.method = method
	req := request.(*codec.RawTLObject)
	c.reqPayload = append([]byte(nil), req.Payload...)

	resp := response.(*codec.RawTLObject)
	resp.Payload = append(resp.Payload[:0], c.respPayload...)
	return nil
}

func TestRawInvokerInvokeRawRoutesByConstructor(t *testing.T) {
	const reqClazzID = uint32(0xfeed1001)
	iface.RegisterClazzName("raw_invoke_test", 0, reqClazzID)
	iface.RegisterClazzIDName("raw_invoke_test", reqClazzID)
	iface.RegisterRPCContextTuple("TLRawInvokeTest", "/tg.RPCRawTest/raw.invokeTest", func() interface{} {
		return codec.NewRawTLObject(nil)
	})

	reqPayload := []byte{0x01, 0x10, 0xed, 0xfe, 0x11, 0x22}
	respPayload := []byte{0x02, 0x10, 0xed, 0xfe, 0x33, 0x44}
	client := &fakeRawClient{respPayload: respPayload}
	invoker := NewRawInvoker(map[string]Client{
		"RPCRawTest": client,
	})

	resp, err := invoker.InvokeRaw(context.Background(), &metadata.RpcMetadata{}, reqPayload)
	if err != nil {
		t.Fatalf("InvokeRaw() error = %v", err)
	}

	if client.method != "/tg.RPCRawTest/raw.invokeTest" {
		t.Fatalf("method = %q, want %q", client.method, "/tg.RPCRawTest/raw.invokeTest")
	}
	if string(client.reqPayload) != string(reqPayload) {
		t.Fatalf("request payload = %x, want %x", client.reqPayload, reqPayload)
	}
	if string(resp.Payload) != string(respPayload) {
		t.Fatalf("response payload = %x, want %x", resp.Payload, respPayload)
	}
}

func TestRawInvokerInvokeRawMethodUsesExplicitRoute(t *testing.T) {
	reqPayload := []byte{0x05, 0x10, 0xed, 0xfe, 0x11, 0x22}
	respPayload := []byte{0x06, 0x10, 0xed, 0xfe, 0x33, 0x44}
	client := &fakeRawClient{respPayload: respPayload}
	invoker := NewRawInvoker(map[string]Client{
		"RPCRawExplicit": client,
	})

	resp, err := invoker.InvokeRawMethod(
		context.Background(),
		&metadata.RpcMetadata{},
		"RPCRawExplicit",
		"/tg.RPCRawExplicit/raw.explicit",
		reqPayload)
	if err != nil {
		t.Fatalf("InvokeRawMethod() error = %v", err)
	}

	if client.method != "/tg.RPCRawExplicit/raw.explicit" {
		t.Fatalf("method = %q, want %q", client.method, "/tg.RPCRawExplicit/raw.explicit")
	}
	if string(client.reqPayload) != string(reqPayload) {
		t.Fatalf("request payload = %x, want %x", client.reqPayload, reqPayload)
	}
	if string(resp.Payload) != string(respPayload) {
		t.Fatalf("response payload = %x, want %x", resp.Payload, respPayload)
	}
}

func TestRawInvokerInvokeRawRejectsUnknownConstructor(t *testing.T) {
	invoker := NewRawInvoker(nil)

	_, err := invoker.InvokeRaw(context.Background(), nil, []byte{0x03, 0x10, 0xed, 0xfe})
	if err == nil {
		t.Fatal("InvokeRaw() error = nil, want error")
	}
	if !errors.Is(err, ErrRawConstructorNotRegistered) {
		t.Fatalf("InvokeRaw() error = %v, want %v", err, ErrRawConstructorNotRegistered)
	}
}

func TestRawInvokerInvokeRawRejectsMissingClient(t *testing.T) {
	const reqClazzID = uint32(0xfeed1004)
	iface.RegisterClazzName("raw_missing_client", 0, reqClazzID)
	iface.RegisterClazzIDName("raw_missing_client", reqClazzID)
	iface.RegisterRPCContextTuple("TLRawMissingClient", "/tg.RPCRawMissing/raw.missingClient", func() interface{} {
		return codec.NewRawTLObject(nil)
	})

	invoker := NewRawInvoker(nil)
	_, err := invoker.InvokeRaw(context.Background(), nil, []byte{0x04, 0x10, 0xed, 0xfe})
	if err == nil {
		t.Fatal("InvokeRaw() error = nil, want error")
	}
	if !errors.Is(err, ErrRawClientNotFound) {
		t.Fatalf("InvokeRaw() error = %v, want %v", err, ErrRawClientNotFound)
	}
}

func TestRawRPCResponseConstructorID(t *testing.T) {
	resp := RawRPCResponse{Payload: []byte{0x02, 0x10, 0xed, 0xfe}}

	got, err := resp.ConstructorID()
	if err != nil {
		t.Fatalf("ConstructorID() error = %v", err)
	}
	if got != 0xfeed1002 {
		t.Fatalf("ConstructorID() = %#x, want %#x", got, uint32(0xfeed1002))
	}
}
