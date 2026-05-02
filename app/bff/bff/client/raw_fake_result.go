package bffproxyclient

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
)

func TryReturnRawFakeRpcResult(ctx context.Context, md *metadata.RpcMetadata, payload []byte) ([]byte, bool, error) {
	_ = ctx
	_ = md

	constructorID, err := bin.NewDecoder(payload).PeekClazzID()
	if err != nil {
		return nil, false, fmt.Errorf("raw fake result constructor: %w", err)
	}
	if !iface.CheckClazzID(constructorID) {
		return nil, false, nil
	}
	obj, err := iface.DecodeObject(bin.NewDecoder(payload))
	if err != nil {
		return nil, true, fmt.Errorf("decode raw fake request: %w", err)
	}
	result, err := new(BFFProxyClient2).TryReturnFakeRpcResult(obj)
	if err != nil {
		return nil, true, err
	}
	if result == nil {
		return nil, false, nil
	}
	x := bin.NewEncoder()
	defer x.End()
	if err := result.Encode(x, 0); err != nil {
		x.Reset()
		if err2 := result.Encode(x, 224); err2 != nil {
			return nil, true, fmt.Errorf("encode raw fake result: %w", err)
		}
	}
	return append([]byte(nil), x.Bytes()...), true, nil
}
