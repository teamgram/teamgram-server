package sess

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"
	rpcmetadata "github.com/teamgram/proto/mtproto/rpc/metadata"
	"google.golang.org/grpc/status"
)

type takeoutGuard struct{}

func newTakeoutGuard() *takeoutGuard {
	return &takeoutGuard{}
}

func (g *takeoutGuard) Validate(takeoutId int64) (*rpcmetadata.Takeout, error) {
	if takeoutId == 0 {
		return nil, takeoutRequiredError()
	}

	return &rpcmetadata.Takeout{Id: takeoutId}, nil
}

func (g *takeoutGuard) ValidateWrappedQuery(takeoutId int64, query mtproto.TLObject) (mtproto.TLObject, *rpcmetadata.Takeout, error) {
	takeout, err := g.Validate(takeoutId)
	if err != nil {
		return nil, nil, err
	}

	wrapped, ok := query.(*mtproto.TLInvokeWithMessagesRange)
	if !ok {
		return query, takeout, nil
	}

	inner, err := decodeWrappedQuery(wrapped.GetQuery())
	if err != nil {
		return nil, nil, err
	}

	takeout.Range = toTakeoutMessageRange(wrapped.GetRange())
	return inner, takeout, nil
}

func newTakeoutMetadata(takeoutId int64, msgRange *mtproto.MessageRange) *rpcmetadata.Takeout {
	takeout := &rpcmetadata.Takeout{}
	if takeoutId != 0 {
		takeout.Id = takeoutId
	}
	if msgRange != nil {
		takeout.Range = toTakeoutMessageRange(msgRange)
	}
	if takeout.Id == 0 && takeout.Range == nil {
		return nil
	}
	return takeout
}

func decodeWrappedQuery(queryData []byte) (mtproto.TLObject, error) {
	if len(queryData) == 0 {
		return nil, mtproto.NewRpcError(status.Error(mtproto.ErrBadRequest, "TAKEOUT_QUERY_INVALID"))
	}

	dBuf := mtproto.NewDecodeBuf(queryData)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		return nil, mtproto.NewRpcError(status.Error(mtproto.ErrBadRequest, fmt.Sprintf("TAKEOUT_QUERY_INVALID: %v", dBuf.GetError())))
	}
	if query == nil {
		return nil, mtproto.NewRpcError(status.Error(mtproto.ErrBadRequest, "TAKEOUT_QUERY_INVALID"))
	}
	return query, nil
}

func toTakeoutMessageRange(msgRange *mtproto.MessageRange) *rpcmetadata.TakeoutMessageRange {
	if msgRange == nil {
		return nil
	}

	return &rpcmetadata.TakeoutMessageRange{
		MinId: msgRange.GetMinId(),
		MaxId: msgRange.GetMaxId(),
	}
}

func takeoutRequiredError() *mtproto.TLRpcError {
	return mtproto.NewRpcError(status.Error(mtproto.ErrForbidden, "TAKEOUT_REQUIRED"))
}
