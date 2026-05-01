package dispatch

import (
	"context"
	"errors"
	"testing"

	bffclient "github.com/teamgram/teamgram-server/v2/app/bff/bff/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeRawClient struct {
	method      string
	reqPayload  []byte
	respPayload []byte
	err         error
}

func (c *fakeRawClient) Call(ctx context.Context, method string, request, response interface{}) error {
	c.method = method
	c.reqPayload = append([]byte(nil), request.(*codec.RawTLObject).Payload...)
	if c.err != nil {
		return c.err
	}
	response.(*codec.RawTLObject).Payload = append([]byte(nil), c.respPayload...)
	return nil
}

func TestRawDispatchSuccess(t *testing.T) {
	const reqClazzID = uint32(0xfeed3001)
	iface.RegisterClazzName("raw_dispatch_test", 0, reqClazzID)
	iface.RegisterClazzIDName("raw_dispatch_test", reqClazzID)
	iface.RegisterRPCContextTuple("TLRawDispatchTest", "/tg.RPCRawDispatch/raw.dispatchTest", func() interface{} {
		return codec.NewRawTLObject(nil)
	})
	reqPayload := []byte{0x01, 0x30, 0xed, 0xfe}
	respPayload := []byte{0x02, 0x30, 0xed, 0xfe}
	rawClient := &fakeRawClient{respPayload: respPayload}
	dispatcher := NewRawDispatcher(&bffclient.BFFProxyClient2{
		RawClients: map[string]kitex.Client{"RPCRawDispatch": rawClient},
	})

	got, err := dispatcher.Invoke(context.Background(), &metadata.RpcMetadata{AuthId: 1}, reqPayload)
	if err != nil {
		t.Fatalf("Invoke() error = %v", err)
	}
	if string(got) != string(respPayload) {
		t.Fatalf("Invoke() = %x, want %x", got, respPayload)
	}
	if rawClient.method != "/tg.RPCRawDispatch/raw.dispatchTest" || string(rawClient.reqPayload) != string(reqPayload) {
		t.Fatalf("raw call method=%q payload=%x", rawClient.method, rawClient.reqPayload)
	}
}

func TestRawDispatchMissingClientUsesFakeFallback(t *testing.T) {
	req := encodeRaw(t, &tg.TLHelpGetTermsOfServiceUpdate{})
	dispatcher := NewRawDispatcher(&bffclient.BFFProxyClient2{RawClients: map[string]kitex.Client{}})
	got, err := dispatcher.Invoke(context.Background(), &metadata.RpcMetadata{}, req)
	if err != nil {
		t.Fatalf("Invoke() error = %v", err)
	}
	obj, err := iface.DecodeObject(bin.NewDecoder(got))
	if err != nil {
		t.Fatalf("DecodeObject() error = %v", err)
	}
	if _, ok := obj.(*tg.TLHelpTermsOfServiceUpdateEmpty); !ok {
		t.Fatalf("DecodeObject() = %T", obj)
	}
}

func TestRawDispatchInfrastructureError(t *testing.T) {
	const reqClazzID = uint32(0xfeed3002)
	iface.RegisterClazzName("raw_dispatch_infra", 0, reqClazzID)
	iface.RegisterClazzIDName("raw_dispatch_infra", reqClazzID)
	iface.RegisterRPCContextTuple("TLRawDispatchInfra", "/tg.RPCRawDispatchInfra/raw.dispatchInfra", func() interface{} {
		return codec.NewRawTLObject(nil)
	})
	want := errors.New("connection refused")
	rawClient := &fakeRawClient{err: want}
	dispatcher := NewRawDispatcher(&bffclient.BFFProxyClient2{
		RawClients: map[string]kitex.Client{"RPCRawDispatchInfra": rawClient},
	})
	_, err := dispatcher.Invoke(context.Background(), nil, []byte{0x02, 0x30, 0xed, 0xfe})
	if !errors.Is(err, want) {
		t.Fatalf("Invoke() error = %v, want %v", err, want)
	}
}

func TestRawDispatchMissingConstructorUsesSentinel(t *testing.T) {
	dispatcher := NewRawDispatcher(&bffclient.BFFProxyClient2{RawClients: map[string]kitex.Client{}})
	_, err := dispatcher.Invoke(context.Background(), nil, []byte{0x99, 0x99, 0xed, 0xfe})
	if !errors.Is(err, kitex.ErrRawConstructorNotRegistered) {
		t.Fatalf("Invoke() error = %v, want ErrRawConstructorNotRegistered", err)
	}
}

func encodeRaw(t *testing.T, obj interface {
	Encode(*bin.Encoder, int32) error
}) []byte {
	t.Helper()
	x := bin.NewEncoder()
	defer x.End()
	if err := obj.Encode(x, 0); err != nil {
		x.Reset()
		if err2 := obj.Encode(x, 224); err2 != nil {
			t.Fatalf("Encode(%T) error = %v", obj, err)
		}
	}
	return append([]byte(nil), x.Bytes()...)
}
