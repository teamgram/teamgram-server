package dispatch

import (
	"context"
	"errors"
	"fmt"

	bffclient "github.com/teamgram/teamgram-server/v2/app/bff/bff/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface/ecode"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type RawDispatcher struct {
	bff *bffclient.BFFProxyClient2
}

type RPCError struct {
	err *tg.TLRpcError
}

func NewRPCError(err error) *RPCError {
	return &RPCError{err: tg.NewRpcError(err)}
}

func (e *RPCError) Error() string {
	if e == nil || e.err == nil {
		return "raw dispatcher rpc error"
	}
	return e.err.Error()
}

func (e *RPCError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

func (e *RPCError) RPCError() *tg.TLRpcError {
	if e == nil {
		return nil
	}
	return e.err
}

func NewRawDispatcher(bff *bffclient.BFFProxyClient2) *RawDispatcher {
	return &RawDispatcher{bff: bff}
}

func (d *RawDispatcher) Invoke(ctx context.Context, md *metadata.RpcMetadata, payload []byte) ([]byte, error) {
	if d == nil || d.bff == nil {
		return nil, fmt.Errorf("raw dispatcher: bff client is nil")
	}
	resp, err := d.bff.InvokeRawContext(ctx, md, payload)
	if err == nil {
		if resp == nil {
			return nil, fmt.Errorf("raw dispatcher: nil raw response")
		}
		return resp.Payload, nil
	}
	if !isMissingRoute(err) {
		if isTLRPCError(err) {
			return nil, NewRPCError(err)
		}
		return nil, err
	}

	constructorID, idErr := bin.NewDecoder(payload).PeekClazzID()
	if idErr != nil {
		return nil, fmt.Errorf("raw dispatcher: request constructor: %w", idErr)
	}
	if fake, ok, fakeErr := bffclient.TryReturnRawFakeRpcResult(ctx, md, constructorID); fakeErr != nil {
		if isTLRPCError(fakeErr) {
			return nil, NewRPCError(fakeErr)
		}
		return nil, fakeErr
	} else if ok {
		return fake, nil
	}
	return nil, err
}

func isMissingRoute(err error) bool {
	return errors.Is(err, kitex.ErrRawClientNotFound) || errors.Is(err, kitex.ErrRawConstructorNotRegistered)
}

func isTLRPCError(err error) bool {
	var rpcErr *tg.TLRpcError
	if errors.As(err, &rpcErr) && rpcErr != nil {
		return true
	}
	var codeErr ecode.CodeError
	return errors.As(err, &codeErr) && codeErr != nil
}
