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

package service

import (
	"context"
	"reflect"
	"strconv"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

func (c *session) onInvokeWithLayer(gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithLayer) {
	logx.Infof("onInvokeWithLayer - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request.DebugString())

	if request.GetQuery() == nil {
		logx.Errorf("invokeWithLayer Query is nil, query: {%s}", request.DebugString())
		// pack.errMsgIDList = append(pack.errMsgIDList, msgId)
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.Errorf("decode buf is nil, query: %v", query)
		return
	}

	initConnection, ok := query.(*mtproto.TLInitConnection)
	if !ok {
		logx.Errorf("need initConnection, but query is : %v", query)
	}

	c.cb.setLayer(request.Layer)
	c.cb.setClient(initConnection.LangPack)

	c.PutUploadInitConnection(context.Background(), c.cb.getAuthKeyId(), request.Layer, clientIp, initConnection)

	dBuf = mtproto.NewDecodeBuf(initConnection.GetQuery())
	query = dBuf.Object()
	if dBuf.GetError() != nil {
		logx.Errorf("dBuf query error: %s", dBuf.GetError().Error())
		return
	}

	if query == nil {
		logx.Errorf("decode buf is nil, query: %v", query)
		return
	}

	logx.Infof("processMsg - data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, query: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		query.DebugString())

	c.processMsg(gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeAfterMsg(gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeAfterMsg) {
	logx.Infof("onInvokeAfterMsg - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%v}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	if request.GetQuery() == nil {
		logx.Errorf("invokeAfterMsg Query is nil, query: {%s}", request.DebugString())
		// pack.errMsgIDList = append(pack.errMsgIDList, msgId)
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.Errorf("decode buf is nil, query: %v", query)
		return
	}

	//		if invokeAfterMsg.GetQuery() == nil {
	//			logx.Errorf("invokeAfterMsg Query is nil, query: {%v}", invokeAfterMsg)
	//			return
	//		}
	//
	//		dbuf := mtproto.NewDecodeBuf(invokeAfterMsg.Query)
	//		query := dbuf.Object()
	//		if query == nil {
	//			logx.Errorf("Decode query error: %s", hex.EncodeToString(invokeAfterMsg.Query))
	//			return
	//		}
	//
	//		var found = false
	//		for j := 0; j < i; j++ {
	//			if messages[j].MsgId == invokeAfterMsg.MsgId {
	//				messages[i].Object = query
	//				found = true
	//				break
	//			}
	//		}
	//
	//		if !found {
	//			for j := i + 1; j < len(messages); j++ {
	//				if messages[j].MsgId == invokeAfterMsg.MsgId {
	//					// c.messages[i].Object = query
	//					messages[i].Object = query
	//					found = true
	//					messages = append(messages, messages[i])
	//
	//					// set messages[i] = nil, will ignore this.
	//					messages[i] = nil
	//					break
	//				}
	//			}
	//		}
	//
	//		if !found {
	//			// TODO(@benqi): backup message, wait.
	//
	//			messages[i].Object = query
	//		}

	c.processMsg(gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeAfterMsgs(gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeAfterMsgs) {
	logx.Infof("onInvokeAfterMsgs - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%v}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	if request.GetQuery() == nil {
		logx.Errorf("invokeAfterMsgs Query is nil, query: {%s}", request.DebugString())
		// pack.errMsgIDList = append(pack.errMsgIDList, msgId)
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.Errorf("decode buf is nil, query: %v", query)
		return
	}

	//		if invokeAfterMsgs.GetQuery() == nil {
	//			logx.Errorf("invokeAfterMsgs Query is nil, query: {%v}", invokeAfterMsgs)
	//			return
	//		}
	//
	//		dbuf := mtproto.NewDecodeBuf(invokeAfterMsgs.Query)
	//		query := dbuf.Object()
	//		if query == nil {
	//			logx.Errorf("Decode query error: %s", hex.EncodeToString(invokeAfterMsgs.Query))
	//			return
	//		}
	//
	//		if len(invokeAfterMsgs.MsgIds) == 0 {
	//			// TODO(@benqi): invalid msgIds, ignore??
	//
	//			messages[i].Object = query
	//		} else {
	//			var maxMsgId = invokeAfterMsgs.MsgIds[0]
	//			for j := 1; j < len(invokeAfterMsgs.MsgIds); j++ {
	//				if maxMsgId > invokeAfterMsgs.MsgIds[j] {
	//					maxMsgId = invokeAfterMsgs.MsgIds[j]
	//				}
	//			}
	//
	//
	//			var found = false
	//			for j := 0; j < i; j++ {
	//				if messages[j].MsgId == maxMsgId {
	//					messages[i].Object = query
	//					found = true
	//					break
	//				}
	//			}
	//
	//			if !found {
	//				for j := i + 1; j < len(messages); j++ {
	//					if messages[j].MsgId == maxMsgId {
	//						// c.messages[i].Object = query
	//						messages[i].Object = query
	//						found = true
	//						messages = append(messages, messages[i])
	//
	//						// set messages[i] = nil, will ignore this.
	//						messages[i] = nil
	//						break
	//					}
	//				}
	//			}
	//
	//			if !found {
	//				// TODO(@benqi): backup message, wait.
	//
	//				messages[i].Object = query
	//			}

	c.processMsg(gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithoutUpdates(gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithoutUpdates) {
	logx.Infof("onInvokeWithoutUpdates - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		reflect.TypeOf(request))

	if request.GetQuery() == nil {
		logx.Errorf("invokeWithoutUpdates Query is nil, query: {%s}", request.DebugString())
		// pack.errMsgIDList = append(pack.errMsgIDList, msgId)
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithMessagesRange(gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithMessagesRange) {
	logx.Infof("onInvokeWithMessagesRange - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.msgId,
		reflect.TypeOf(request))

	if request.GetQuery() == nil {
		logx.Errorf("invokeWithMessagesRange Query is nil, query: {%s}", request.DebugString())
		// pack.errMsgIDList = append(pack.errMsgIDList, msgId)
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(gatewayId, clientIp, msgId, query)
}

func (c *session) onInvokeWithTakeout(gatewayId, clientIp string, msgId *inboxMsg, request *mtproto.TLInvokeWithTakeout) {
	logx.Infof("onInvokeWithTakeout - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.msgId,
		reflect.TypeOf(request))

	if request.GetQuery() == nil {
		logx.Errorf("invokeWithTakeout Query is nil, query: {%s}", request.DebugString())
		// pack.errMsgIDList = append(pack.errMsgIDList, msgId)
		return
	}

	dBuf := mtproto.NewDecodeBuf(request.Query)
	query := dBuf.Object()
	if dBuf.GetError() != nil {
		logx.Errorf("dBuf query error: %v", dBuf.GetError())
		return
	}

	if query == nil {
		logx.Errorf("decode buf is nil, query: %v", query)
		return
	}

	c.processMsg(gatewayId, clientIp, msgId, query)
}

func (c *session) onRpcRequest(gatewayId, clientIp string, msgId *inboxMsg, query mtproto.TLObject) bool {
	logx.Infof("onRpcRequest - request data: {sess: %s, gatewayId: %s, msg_id: %d, seq_no: %d, request: {%s}}",
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
			pushSessionId, err := strconv.ParseInt(registerDevice.GetToken(), 10, 64)
			if err == nil {
				c.cb.onBindPushSessionId(pushSessionId)
				c.PutCachePushSessionId(context.Background(), c.cb.getAuthKeyId(), int64(pushSessionId))
			}
		}
	case *mtproto.TLUpdatesGetState:
		if !c.isGeneric {
			c.isGeneric = true
			c.cb.setOnline()
		}
	case *mtproto.TLUpdatesGetDifference:
		if !c.isGeneric {
			c.isGeneric = true
			c.cb.setOnline()
		}
	case *mtproto.TLUpdatesGetChannelDifference:
		if !c.isGeneric {
			c.isGeneric = true
			c.cb.setOnline()
		}
		//case *mtproto.TLAuthBindTempAuthKey:
		//	res, err := c.AuthSessionRpcClient.AuthBindTempAuthKey(context.Background(), query.(*mtproto.TLAuthBindTempAuthKey))
		//	if err != nil {
		//		logx.Errorf("bindTempAuthKey error - %v", err)
		//		err = mtproto.ErrInternelServerError
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
		if !c.isGeneric {
			c.isGeneric = true
			c.cb.setOnline()
		}
	case *mtproto.TLUsersGetUsers:
		logx.Infof("user.getUsers: %s", query.DebugString())
	}

	if c.cb.getUserId() == 0 {
		if !checkRpcWithoutLogin(query) {
			authUserId, _ := c.GetCacheUserID(context.Background(), c.cb.getAuthKeyId())
			if authUserId == 0 {
				logx.Errorf("not found authUserId by authKeyId: %d", c.cb.getAuthKeyId())
				// 401
				rpcError := &mtproto.TLRpcError{Data2: &mtproto.RpcError{
					ErrorCode:    401,
					ErrorMessage: "AUTH_KEY_INVALID",
				}}
				c.sendRpcResultToQueue(gatewayId, msgId.msgId, rpcError)
				msgId.state = RECEIVED | RESPONSE_GENERATED
				return false
			} else {
				c.cb.setUserId(authUserId)
			}
		}
	}

	msgId.state = RECEIVED | RPC_PROCESSING
	c.cb.sendToRpcQueue(&rpcApiMessage{
		sessionId: c.sessionId,
		clientIp:  clientIp,
		reqMsgId:  msgId.msgId,
		reqMsg:    query,
	})

	return true
}

func (c *session) onRpcResult(rpcResult *rpcApiMessage) {
	defer func() {
		if _, ok := rpcResult.reqMsg.(*mtproto.TLAuthLogOut3E72BA19); ok {
			c.DeleteByAuthKeyId(c.cb.getAuthKeyId())
		}
		if _, ok := rpcResult.reqMsg.(*mtproto.TLAuthLogOut5717DA40); ok {
			c.DeleteByAuthKeyId(c.cb.getAuthKeyId())
		}
	}()

	if rpcErr, ok := rpcResult.rpcResult.Result.(*mtproto.TLRpcError); ok {
		if rpcErr.GetErrorCode() == int32(mtproto.ErrNotReturnClient) {
			logx.Infof("recv NOTRETURN_CLIENT")
			c.pendingQueue.Add(rpcResult.reqMsgId)
			return
		}
	}

	c.sendRpcResult(rpcResult.rpcResult)
}

func (c *session) sendRpcResult(rpcResult *mtproto.TLRpcResult) {
	// TODO(@benqi): lookup inBoxMsg
	msgId := c.inQueue.Lookup(rpcResult.ReqMsgId)
	if msgId == nil {
		logx.Errorf("not found msgId, maybe removed: %d", rpcResult.ReqMsgId)
		return
	}

	gatewayId := c.getGatewayId()
	c.sendRpcResultToQueue(gatewayId, msgId.msgId, rpcResult.Result)
	msgId.state = RECEIVED | ACKNOWLEDGED

	if gatewayId == "" {
		logx.Errorf("gatewayId is empty, send delay...")
	} else {
		c.sendQueueToGateway(gatewayId)
	}
}
