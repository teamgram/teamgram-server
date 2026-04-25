package kitex

import (
	"context"
	"errors"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/codec"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

var ErrRawClientNotFound = errors.New("raw rpc client not found")

type RawRPCResponse struct {
	Payload []byte
}

func (m RawRPCResponse) ConstructorID() (uint32, error) {
	return bin.NewDecoder(m.Payload).PeekClazzID()
}

type RawInvoker struct {
	clients map[string]Client
}

func NewRawInvoker(clients map[string]Client) *RawInvoker {
	return &RawInvoker{clients: clients}
}

func (m *RawInvoker) InvokeRaw(ctx context.Context, rpcMetaData *metadata.RpcMetadata, reqPayload []byte) (*RawRPCResponse, error) {
	req := codec.NewRawTLObject(reqPayload)
	clazzID, err := req.ConstructorID()
	if err != nil {
		return nil, fmt.Errorf("raw rpc request constructor: %w", err)
	}

	tuple := iface.FindRPCContextTupleByClazzID(clazzID)
	if tuple == nil {
		return nil, fmt.Errorf("raw rpc constructor %#x is not registered", clazzID)
	}

	return m.InvokeRawMethod(ctx, rpcMetaData, tuple.ServiceName(), tuple.Method, reqPayload)
}

func (m *RawInvoker) InvokeRawMethod(ctx context.Context, rpcMetaData *metadata.RpcMetadata, serviceName, methodName string, reqPayload []byte) (*RawRPCResponse, error) {
	client, ok := m.clients[serviceName]
	if !ok || client == nil {
		return nil, fmt.Errorf("%w: %s", ErrRawClientNotFound, serviceName)
	}

	if rpcMetaData != nil {
		var err error
		ctx, err = metadata.RpcMetadataToOutgoing(ctx, rpcMetaData)
		if err != nil {
			return nil, err
		}
	}

	req := codec.NewRawTLObject(reqPayload)
	resp := codec.NewRawTLObject(nil)
	if err := client.Call(ctx, methodName, req, resp); err != nil {
		return nil, err
	}

	return &RawRPCResponse{
		Payload: resp.Payload,
	}, nil
}
