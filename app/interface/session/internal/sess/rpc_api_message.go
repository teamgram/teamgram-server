// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package sess

import (
	"context"
	"fmt"

	"github.com/teamgram/proto/mtproto"
)

type rpcApiMessage struct {
	ctx       context.Context
	sessList  *SessionList
	sessionId int64
	clientIp  string
	reqMsgId  int64
	reqMsg    mtproto.TLObject
	rpcResult *mtproto.TLRpcResult
}

func (m *rpcApiMessage) MoveRpcResult() *mtproto.TLRpcResult {
	v := m.rpcResult
	m.rpcResult = nil
	return v
}

func (m *rpcApiMessage) TryGetRpcResultError() (*mtproto.TLRpcError, bool) {
	if m.rpcResult != nil && m.rpcResult.Result != nil {
		r := m.rpcResult.Result
		switch t := r.(type) {
		case *mtproto.TLRpcError:
			return t, true
		}
	}

	return nil, false
}

func (m *rpcApiMessage) DebugString() string {
	if m.rpcResult == nil {
		return fmt.Sprintf("{trace_id: %d, session_id: %d, req_msg_id: %d, req_msg: %s}",
			m.sessionId,
			m.reqMsgId,
			m.reqMsg)
	} else {
		return fmt.Sprintf("{trace_id: %d, session_id: %d, req_msg_id: %d, req_msg: %s, rpc_result: %s}",
			m.sessionId,
			m.reqMsgId,
			m.reqMsg,
			m.rpcResult.Result)
	}
}
