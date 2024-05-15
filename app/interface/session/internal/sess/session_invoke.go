// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package sess

import (
	"context"
	"reflect"
	"strconv"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/logx"
)

func (c *session) onInvokeWithLayer(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithLayer) {
	logx.WithContext(ctx).Infof("onInvokeWithLayer - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request.DebugString())

	if request.GetQuery() == nil {
		logx.WithContext(ctx).Errorf("invokeWithLayer Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.sessList.cb.onUpdateLayer(ctx, request.Layer)

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeAfterMsg(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeAfterMsg) {
	logx.WithContext(ctx).Infof("onInvokeAfterMsg - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%v}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	if request.GetQuery() == nil {
		logx.WithContext(ctx).Errorf("invokeAfterMsg Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeAfterMsgs(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeAfterMsgs) {
	logx.WithContext(ctx).Infof("onInvokeAfterMsgs - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%v}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	if request.GetQuery() == nil {
		logx.WithContext(ctx).Errorf("invokeAfterMsgs Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	/*
		if invokeAfterMsgs.GetQuery() == nil {
			logx.Errorf("invokeAfterMsgs Query is nil, query: {%v}", invokeAfterMsgs)
			return
		}

		dbuf := mtproto.NewDecodeBuf(invokeAfterMsgs.Query)
		query := dbuf.Object()
		if query == nil {
			logx.Errorf("Decode query error: %s", hex.EncodeToString(invokeAfterMsgs.Query))
			return
		}

		if len(invokeAfterMsgs.MsgIds) == 0 {
			// TODO(@benqi): invalid msgIds, ignore??

			messages[i].Object = query
		} else {
			var maxMsgId = invokeAfterMsgs.MsgIds[0]
			for j := 1; j < len(invokeAfterMsgs.MsgIds); j++ {
				if maxMsgId > invokeAfterMsgs.MsgIds[j] {
					maxMsgId = invokeAfterMsgs.MsgIds[j]
				}
			}


			var found = false
			for j := 0; j < i; j++ {
				if messages[j].MsgId == maxMsgId {
					messages[i].Object = query
					found = true
					break
				}
			}

			if !found {
				for j := i + 1; j < len(messages); j++ {
					if messages[j].MsgId == maxMsgId {
						// c.messages[i].Object = query
						messages[i].Object = query
						found = true
						messages = append(messages, messages[i])

						// set messages[i] = nil, will ignore this.
						messages[i] = nil
						break
					}
				}
			}

			if !found {
				// TODO(@benqi): backup message, wait.

				messages[i].Object = query
			}
		}
	*/

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithoutUpdates(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithoutUpdates) {
	logx.WithContext(ctx).Infof("onInvokeWithoutUpdates - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		reflect.TypeOf(request))

	if request.GetQuery() == nil {
		logx.WithContext(ctx).Errorf("invokeWithoutUpdates Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithMessagesRange(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithMessagesRange) {
	logx.WithContext(ctx).Infof("onInvokeWithMessagesRange - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.msgId,
		reflect.TypeOf(request))

	if request.GetQuery() == nil {
		logx.WithContext(ctx).Errorf("invokeWithMessagesRange Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithTakeout(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithTakeout) {
	logx.WithContext(ctx).Infof("onInvokeWithTakeout - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.msgId,
		reflect.TypeOf(request))

	if request.GetQuery() == nil {
		logx.WithContext(ctx).Errorf("invokeWithTakeout Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithBusinessConnection(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithBusinessConnection) {
	logx.WithContext(ctx).Infof("onInvokeWithBusinessConnection - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.msgId,
		reflect.TypeOf(request))

	if request.GetQuery() == nil {
		logx.WithContext(ctx).Errorf("invokeWithBusinessConnection Query is nil, query: {%s}", request.DebugString())
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onInitConnection(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInitConnection) {
	logx.WithContext(ctx).Infof("onInitConnection - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.msgId,
		reflect.TypeOf(request))

	c.sessList.cb.onUpdateInitConnection(ctx, clientIp, request)

	dBuf := mtproto.NewDecodeBuf(request.GetQuery())
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.WithContext(ctx).Errorf("dBuf query error: %s", dBuf.GetError().Error())
		return
	}

	if query == nil {
		logx.WithContext(ctx).Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(ctx, gatewayId, clientIp, msgId, query)
}

func (c *session) onRpcRequest(ctx context.Context, gatewayId, clientIp string, msgId *inboxMsg, query mtproto.TLObject) bool {
	logx.WithContext(ctx).Infof("onRpcRequest - request data: {sess: %s, gatewayId: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		reflect.TypeOf(query))

	// TODO(@benqi): sync AuthUserId??
	//requestMessage := &mtproto.TLMessage2{
	//	MsgId:  msgId.msgId,
	//	Seqno:  msgId.seqNo,
	//	Object: object,
	//}

	switch query.(type) {
	case *mtproto.TLAccountRegisterDevice:
		registerDevice, _ := query.(*mtproto.TLAccountRegisterDevice)
		if registerDevice.TokenType == 7 {
			pushSessionId, _ := strconv.ParseInt(registerDevice.GetToken(), 10, 64)
			c.sessList.cb.onBindPushSessionId(ctx, c.sessList, pushSessionId)
		}
	case *mtproto.TLUpdatesGetState:
		c.sessList.cb.onSetMainUpdatesSession(ctx, c)
	case *mtproto.TLUpdatesGetDifference:
		c.sessList.cb.onSetMainUpdatesSession(ctx, c)
	case *mtproto.TLUpdatesGetChannelDifference:
		c.sessList.cb.onSetMainUpdatesSession(ctx, c)
		//case *mtproto.TLAuthBindTempAuthKey:
		//	res, err := c.AuthSessionRpcClient.AuthBindTempAuthKey(context.Background(), query.(*mtproto.TLAuthBindTempAuthKey))
		//	if err != nil {
		//		logx.WithContext(ctx).Errorf("bindTempAuthKey error - %v", err)
		//		err = mtproto.ErrInternalServerError
		//		c.sendRpcResultToQueue(gatewayId, msgId.msgId, &mtproto.RpcError{
		//			ErrorCode:    500,
		//			ErrorMessage: "INTERNAL_SERVER_ERROR",
		//		})
		//	} else {
		//		c.sendRpcResultToQueue(gatewayId, msgId.msgId, res)
		//	}
		//	msgId.state = RECEIVED | RESPONSE_GENERATED
		//	return false

		//case *mtproto.TLAuthSignIn:
		//	if !c.isGeneric {
		//		c.isGeneric = true
		//		c.cb.setOnline()
		//	}
		//case *mtproto.TLAuthSignUp:
		//	if !c.isGeneric {
		//		c.isGeneric = true
		//		c.cb.setOnline()
		//	}
	case *mtproto.TLAccountUpdateStatus:
		c.sessList.cb.onSetMainUpdatesSession(ctx, c)
	case *mtproto.TLUsersGetUsers:
		// logx.WithContext(ctx).Infof("user.getUsers: %s", query.DebugString())
	}

	switch c.sessList.cb.state {
	case mtproto.AuthStateNew:
	case mtproto.AuthStatePermBound:
	case mtproto.AuthStateWaitInit:
	case mtproto.AuthStateUnauthorized:
	case mtproto.AuthStateNeedPassword:
	case mtproto.AuthStateNormal:
		//logx.WithContext(ctx).Errorf("not found authUserId by authKeyId: %d", c.sessList.cb.authKeyId)
		//if !checkRpcWithoutLogin(query) {
		//	// 401
		//	rpcError := &mtproto.TLRpcError{Data2: &mtproto.RpcError{
		//		ErrorCode:    401,
		//		ErrorMessage: "AUTH_KEY_UNREGISTERED",
		//	}}
		//	c.sendRpcResultToQueue(ctx, gatewayId, msgId.msgId, rpcError)
		//	msgId.state = RECEIVED | RESPONSE_GENERATED
		//}
		//return false
	case mtproto.AuthStateLogout:
		logx.WithContext(ctx).Errorf("authUserId is logout: %d", c.sessList.cb.authKeyId)
	case mtproto.AuthStateDeleted:
	default:
		logx.WithContext(ctx).Errorf("unknown state: %d", c.sessList.cb.state)
	}

	msgId.state = RECEIVED | RPC_PROCESSING

	c.sessList.cb.tmpRpcApiMessageList = append(
		c.sessList.cb.tmpRpcApiMessageList,
		&rpcApiMessage{
			ctx:       contextx.ValueOnlyFrom(ctx),
			sessList:  c.sessList,
			sessionId: c.sessionId,
			clientIp:  clientIp,
			reqMsgId:  msgId.msgId,
			reqMsg:    query,
		})

	return true
}

func (c *session) onRpcResult(ctx context.Context, rpcResult *rpcApiMessage) {
	rpcErr, ok := rpcResult.TryGetRpcResultError()
	if ok {
		if rpcErr.GetErrorCode() == int32(mtproto.ErrNotReturnClient) {
			logx.WithContext(ctx).Debugf("recv NOTRETURN_CLIENT")
			c.pendingQueue.Add(rpcResult.reqMsgId)
			return
		}
	}

	switch request := rpcResult.reqMsg.(type) {
	case *mtproto.TLAuthBindTempAuthKey:
		if ok {
			_ = request
			if rpcErr.Message() == "ENCRYPTED_MESSAGE_INVALID" {
				// c.sessList.cb.cb.Dao.PutCacheUserId(context.Background(), c.sessList.cb.authKeyId, 0)
				c.sessList.cb.cb.DeleteByAuthKeyId(c.sessList.authId)
				c.sessList.cb.AuthUserId = 0
			}
		} else {
			c.sessList.cb.cb.Dao.PutCachePermAuthKeyId(ctx, c.sessList.authId, request.PermAuthKeyId)
		}
	default:
	}

	defer func() {
		switch rpcResult.reqMsg.(type) {
		case *mtproto.TLAuthLogOut:
			c.sessList.cb.cb.DeleteByAuthKeyId(c.sessList.cb.authKeyId)
		default:
		}
	}()

	c.sendRpcResult(ctx, rpcResult.MoveRpcResult())
}

func (c *session) sendRpcResult(ctx context.Context, rpcResult *mtproto.TLRpcResult) {
	// TODO(@benqi): lookup inBoxMsg
	msgId := c.inQueue.Lookup(rpcResult.ReqMsgId)
	if msgId == nil {
		logx.WithContext(ctx).Errorf("not found msgId, maybe removed: %d", rpcResult.ReqMsgId)
		return
	}

	gatewayId := c.getGatewayId()
	c.sendRpcResultToQueue(ctx, gatewayId, msgId.msgId, rpcResult.Result)
	msgId.state = RECEIVED | ACKNOWLEDGED

	if gatewayId == "" {
		logx.WithContext(ctx).Errorf("gatewayId is empty, send delay...")
	} else {
		c.sendQueueToGateway(ctx, gatewayId)
	}
}
