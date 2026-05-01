package mtproto

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/mt"
)

func WrapRPCResult(reqMsgId int64, result []byte) ([]byte, error) {
	if result == nil {
		return nil, fmt.Errorf("wrap rpc_result: result is nil")
	}
	x := bin.NewEncoder()
	defer x.End()
	x.PutClazzID(mt.ClazzID_rpc_result)
	x.PutInt64(reqMsgId)
	x.PutRaw(result)
	return append([]byte(nil), x.Bytes()...), nil
}

func WrapRPCError(reqMsgId int64, code int32, message string) ([]byte, error) {
	errObj := mt.MakeTLRpcError(&mt.TLRpcError{
		ErrorCode:    code,
		ErrorMessage: message,
	})
	x := bin.NewEncoder()
	defer x.End()
	if err := errObj.Encode(x, 0); err != nil {
		return nil, fmt.Errorf("wrap rpc_error: %w", err)
	}
	return WrapRPCResult(reqMsgId, x.Bytes())
}
