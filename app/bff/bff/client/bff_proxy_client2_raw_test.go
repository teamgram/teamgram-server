package bffproxyclient

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

type fakeRawBFFClient struct {
	method      string
	reqPayload  []byte
	respPayload []byte
}

func (c *fakeRawBFFClient) Call(ctx context.Context, method string, request, response interface{}) error {
	c.method = method
	c.reqPayload = append([]byte(nil), request.(*codec.RawTLObject).Payload...)
	response.(*codec.RawTLObject).Payload = append([]byte(nil), c.respPayload...)
	return nil
}

func TestBFFProxyClient2InvokeRawRoutesByServiceName(t *testing.T) {
	const reqClazzID = uint32(0xfeed2001)
	iface.RegisterClazzName("bff_raw_test", 0, reqClazzID)
	iface.RegisterClazzIDName("bff_raw_test", reqClazzID)
	iface.RegisterRPCContextTuple("TLBffRawTest", "/tg.RPCBffRaw/bff.rawTest", func() interface{} {
		return codec.NewRawTLObject(nil)
	})

	reqPayload := []byte{0x01, 0x20, 0xed, 0xfe, 0x11}
	respPayload := []byte{0x02, 0x20, 0xed, 0xfe, 0x22}
	rawClient := &fakeRawBFFClient{respPayload: respPayload}
	cli := &BFFProxyClient2{
		RawClients: map[string]kitex.Client{
			"RPCBffRaw": rawClient,
		},
	}

	resp, err := cli.InvokeRawContext(context.Background(), &metadata.RpcMetadata{}, reqPayload)
	if err != nil {
		t.Fatalf("InvokeRawContext() error = %v", err)
	}

	if rawClient.method != "/tg.RPCBffRaw/bff.rawTest" {
		t.Fatalf("method = %q, want %q", rawClient.method, "/tg.RPCBffRaw/bff.rawTest")
	}
	if string(rawClient.reqPayload) != string(reqPayload) {
		t.Fatalf("request payload = %x, want %x", rawClient.reqPayload, reqPayload)
	}
	if string(resp.Payload) != string(respPayload) {
		t.Fatalf("response payload = %x, want %x", resp.Payload, respPayload)
	}
}

func TestBFFProxyClient2InvokeRawMethodUsesExplicitRoute(t *testing.T) {
	reqPayload := []byte{0x03, 0x20, 0xed, 0xfe, 0x11}
	respPayload := []byte{0x04, 0x20, 0xed, 0xfe, 0x22}
	rawClient := &fakeRawBFFClient{respPayload: respPayload}
	cli := &BFFProxyClient2{
		RawClients: map[string]kitex.Client{
			"RPCBffRawExplicit": rawClient,
		},
	}

	resp, err := cli.InvokeRawMethodContext(
		context.Background(),
		&metadata.RpcMetadata{},
		"RPCBffRawExplicit",
		"/tg.RPCBffRawExplicit/bff.rawExplicit",
		reqPayload)
	if err != nil {
		t.Fatalf("InvokeRawMethodContext() error = %v", err)
	}

	if rawClient.method != "/tg.RPCBffRawExplicit/bff.rawExplicit" {
		t.Fatalf("method = %q, want %q", rawClient.method, "/tg.RPCBffRawExplicit/bff.rawExplicit")
	}
	if string(resp.Payload) != string(respPayload) {
		t.Fatalf("response payload = %x, want %x", resp.Payload, respPayload)
	}
}
