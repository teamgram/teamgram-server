package dispatch

import (
	"context"
	"errors"
	"fmt"

	bffclient "github.com/teamgram/teamgram-server/v2/app/bff/bff/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
)

type RawDispatcher struct {
	bff *bffclient.BFFProxyClient2
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
		return nil, err
	}

	constructorID, idErr := bin.NewDecoder(payload).PeekClazzID()
	if idErr != nil {
		return nil, fmt.Errorf("raw dispatcher: request constructor: %w", idErr)
	}
	if fake, ok, fakeErr := bffclient.TryReturnRawFakeRpcResult(ctx, md, constructorID); fakeErr != nil {
		return nil, fakeErr
	} else if ok {
		return fake, nil
	}
	return nil, err
}

func isMissingRoute(err error) bool {
	return errors.Is(err, kitex.ErrRawClientNotFound) || errors.Is(err, kitex.ErrRawConstructorNotRegistered)
}
