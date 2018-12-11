// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package server

import (
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/golang/glog"
	"reflect"
	"github.com/nebula-chat/chatengine/pkg/util"
	"time"
)

type genericSession struct {
	*session
	*grpc_util.RPCClient
}

func (c *genericSession) onMessageData(id ClientConnID, cntl *zrpc.ZRpcController, salt int64, msg *mtproto.TLMessage2) {
	c.session.processMessageData(id, cntl, salt, msg, func(sessMsg *mtproto.TLMessage2) {
		switch sessMsg.Object.(type) {
		case *mtproto.TLPing:
			// ignore
		case *mtproto.TLPingDelayDisconnect:
			// ignore
		case *TLInvokeAfterMsgExt: // 未知
			invokeAfterMsgExt, _ := sessMsg.Object.(*TLInvokeAfterMsgExt)
			c.onInvokeAfterMsgExt(id, cntl, sessMsg.MsgId, sessMsg.Seqno, invokeAfterMsgExt)
		case *TLInvokeAfterMsgsExt: // 未知
			invokeAfterMsgsExt, _ := sessMsg.Object.(*TLInvokeAfterMsgsExt)
			c.onInvokeAfterMsgsExt(id, cntl, sessMsg.MsgId, sessMsg.Seqno, invokeAfterMsgsExt)
		case *TLInitConnectionExt: // 都有可能
			initConnectionExt, _ := sessMsg.Object.(*TLInitConnectionExt)
			c.onInitConnectionEx(id, cntl, sessMsg.MsgId, sessMsg.Seqno, initConnectionExt)
		case *TLInvokeWithoutUpdatesExt:
			invokeWithoutUpdatesExt, _ := sessMsg.Object.(*TLInvokeWithoutUpdatesExt)
			c.onInvokeWithoutUpdatesExt(id, cntl, sessMsg.MsgId, sessMsg.Seqno, invokeWithoutUpdatesExt)
		case *TLInvokeWithMessagesRangeExt:
			invokeWithMessagesRangeExt, _ := sessMsg.Object.(*TLInvokeWithMessagesRangeExt)
			c.onInvokeWithMessagesRangeExt(id, cntl, sessMsg.MsgId, sessMsg.Seqno, invokeWithMessagesRangeExt)
		case *TLInvokeWithTakeoutExt:
			invokeWithTakeoutExt, _ := sessMsg.Object.(*TLInvokeWithTakeoutExt)
			c.onInvokeWithTakeoutExt(id, cntl, sessMsg.MsgId, sessMsg.Seqno, invokeWithTakeoutExt)
		default:
			c.onRpcRequest(id, cntl, sessMsg.MsgId, sessMsg.Seqno, sessMsg.Object)
		}
	})

	if len(c.pendingMessages) > 0 {
		c.sendPendingMessagesToClient(id, cntl, c.pendingMessages)
		c.pendingMessages = []*pendingMessage{}
	}

 	if len(c.rpcMessages) > 0 {
		c.cb.sendToRpcQueue(&rpcApiMessages{connID: id, cntl: cntl, sessionId: c.sessionId, rpcMessages: c.rpcMessages})
		c.rpcMessages = []*networkApiMessage{}
	}

	//if len(c.syncMessages) > 0 {
	//	c.sendPendingMessagesToClient(id, cntl, c.syncMessages)
	//	c.syncMessages = []*pendingMessage{}
	//}
}

func (c *genericSession) onInitConnectionEx(connID ClientConnID, cntl *zrpc.ZRpcController, msgId int64, seqNo int32, request *TLInitConnectionExt) bool {
	glog.Infof("onInitConnection - request data: {sess: %s, conn_id: %s, md: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		connID,
		cntl,
		msgId,
		seqNo,
		request)
	// glog.Infof("onInitConnection - request: %s", request.String())
	// auth_session_client.BindAuthKeyUser()
	c.cb.setLayer(request.Layer)
	uploadInitConnection(c.cb.getAuthKeyId(), request.Layer, cntl.GetMtprotoMeta().GetClientAddr(), request)
	return c.onRpcRequest(connID, cntl, msgId, seqNo, request.Query)
}

func (c *genericSession) onInvokeAfterMsgExt(connID ClientConnID, cntl *zrpc.ZRpcController, msgId int64, seqNo int32, request *TLInvokeAfterMsgExt) bool {
	glog.Infof("onInvokeAfterMsgExt - request data: {sess: %s, conn_id: %s, md: %s, msg_id: %d, seq_no: %d, request: {%v}}",
		c,
		connID,
		cntl,
		msgId,
		seqNo,
		request)

	//		if invokeAfterMsg.GetQuery() == nil {
	//			glog.Errorf("invokeAfterMsg Query is nil, query: {%v}", invokeAfterMsg)
	//			return
	//		}
	//
	//		dbuf := mtproto.NewDecodeBuf(invokeAfterMsg.Query)
	//		query := dbuf.Object()
	//		if query == nil {
	//			glog.Errorf("Decode query error: %s", hex.EncodeToString(invokeAfterMsg.Query))
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

	return c.onRpcRequest(connID, cntl, msgId, seqNo, request.Query)
}

func (c *genericSession) onInvokeAfterMsgsExt(connID ClientConnID, cntl *zrpc.ZRpcController, msgId int64, seqNo int32, request *TLInvokeAfterMsgsExt) bool {
	glog.Infof("onInvokeAfterMsgsExt - request data: {sess: %s, conn_id: %s, md: %s, msg_id: %d, seq_no: %d, request: {%v}}",
		c,
		connID,
		cntl,
		msgId,
		seqNo,
		request)
	//		if invokeAfterMsgs.GetQuery() == nil {
	//			glog.Errorf("invokeAfterMsgs Query is nil, query: {%v}", invokeAfterMsgs)
	//			return
	//		}
	//
	//		dbuf := mtproto.NewDecodeBuf(invokeAfterMsgs.Query)
	//		query := dbuf.Object()
	//		if query == nil {
	//			glog.Errorf("Decode query error: %s", hex.EncodeToString(invokeAfterMsgs.Query))
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

	return c.onRpcRequest(connID, cntl, msgId, seqNo, request.Query)
}

func (c *genericSession) onInvokeWithoutUpdatesExt(connID ClientConnID, cntl *zrpc.ZRpcController, msgId int64, seqNo int32, request *TLInvokeWithoutUpdatesExt) bool {
	glog.Infof("onInvokeWithoutUpdatesExt - request data: {sess: %s, conn_id: %s, md: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		connID,
		cntl,
		msgId,
		seqNo,
		reflect.TypeOf(request))

	return c.onRpcRequest(connID, cntl, msgId, seqNo, request.Query)
}

func (c *genericSession) onInvokeWithMessagesRangeExt(connID ClientConnID, cntl *zrpc.ZRpcController, msgId int64, seqNo int32, request *TLInvokeWithMessagesRangeExt) bool {
	glog.Infof("onInvokeWithMessagesRangeExt - request data: {sess: %s, conn_id: %s, md: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		connID,
		cntl,
		msgId,
		seqNo,
		reflect.TypeOf(request))

	return c.onRpcRequest(connID, cntl, msgId, seqNo, request.Query)
}

func (c *genericSession) onInvokeWithTakeoutExt(connID ClientConnID, cntl *zrpc.ZRpcController, msgId int64, seqNo int32, request *TLInvokeWithTakeoutExt) bool {
	glog.Infof("onInvokeWithTakeout - request data: {sess: %s, conn_id: %s, md: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		connID,
		cntl,
		msgId,
		seqNo,
		reflect.TypeOf(request))

	return c.onRpcRequest(connID, cntl, msgId, seqNo, request.Query)
}

func (c *genericSession) onRpcRequest(connID ClientConnID, cntl *zrpc.ZRpcController, msgId int64, seqNo int32, object mtproto.TLObject) bool {
	glog.Infof("onRpcRequest - request data: {sess: %s, conn_id: %s, md: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		connID,
		cntl,
		msgId,
		seqNo,
		reflect.TypeOf(object))

	// TODO(@benqi): sync AuthUserId??
	requestMessage := &mtproto.TLMessage2{
		MsgId:  msgId,
		Seqno:  seqNo,
		Object: object,
	}

	switch object.(type) {
	case *mtproto.TLAccountRegisterDevice:
		registerDevice, _ := object.(*mtproto.TLAccountRegisterDevice)
		if registerDevice.TokenType == 7 {
			pushSessionId, err := util.StringToUint64(registerDevice.GetToken())
			if err == nil {
				c.cb.onBindPushSessionId(int64(pushSessionId))
				putCachePushSessionId(c.cb.getAuthKeyId(), int64(pushSessionId))
			}
		}
	case *mtproto.TLAccountRegisterDeviceLayer71:
		registerDevice, _ := object.(*mtproto.TLAccountRegisterDeviceLayer71)
		if registerDevice.TokenType == 7 {
			pushSessionId, err := util.StringToUint64(registerDevice.GetToken())
			if err == nil {
				c.cb.onBindPushSessionId(int64(pushSessionId))
				putCachePushSessionId(c.cb.getAuthKeyId(), int64(pushSessionId))
			}
		}
	}

	// reqMsgId := msgId
	for e := c.apiMessages.Front(); e != nil; e = e.Next() {
		//v, _ := e.Value.(*networkApiMessage)
		//if v.rpcRequest.MsgId == msgId {
		//	if v.state >= kNetworkMessageStateInvoked {
		//		// c.pendingMessages = append(c.pendingMessages, makePendingMessage(v.rpcMsgId, true, v.rpcResult))
		//		return false
		//	}
		//}
	}

	//if c.sessionType == kSessionUnknown {
	//	c.sessionType = getSessionType(object)
	//	// c.manager.setUserOnline(c.sessionId, connID)
	//}

	if c.cb.getUserId() == 0 {
		if !checkRpcWithoutLogin(object) {
			authUserId := getCacheUserID(c.cb.getAuthKeyId())
			if authUserId == 0 {
				glog.Error("not found authUserId by authKeyId: ", c.cb.getAuthKeyId())
				// 401
				rpcError := &mtproto.TLRpcError{Data2: &mtproto.RpcError_Data{
					ErrorCode: 401,
					ErrorMessage: "AUTH_KEY_INVALID",
				}}
				_ = rpcError
				// c.pendingMessages = append(c.pendingMessages, makePendingMessage(0, true, &mtproto.TLRpcResult{ReqMsgId: msgId, Result: rpcError}))
				return false
			} else {
				c.cb.setUserId(authUserId)
			}
		}
	}

	//// updates
	//if checkRpcUpdatesType(object) {
	//	// c.manager.setUserOnline(c.sessionId, connID)
	//	glog.Infof("onRpcRequest - isUpdate: {connID: {%v}, rpcMethod: {%T}}", connID, object)
	//	c.isUpdates = true
	//	// c.manager.updatesSession.SubscribeUpdates(c, connID)
	//
	//	// subscribe
	//	// c.manager.updatesSession.SubscribeUpdates(c, connID)
	//}

	//if checkRpcPushType(object) {
	//	glog.Infof("onRpcRequest - isPush: {connID: {%v}, rpcMethod: {%T}}", connID, object)
	//	c.isPush = true
	//}

	apiMessage := &networkApiMessage{
		date:       time.Now().Unix(),
		rpcRequest: requestMessage,
		state:      kNetworkMessageStateReceived,
	}
	glog.Info("onRpcRequest - ", apiMessage)

	c.apiMessages.PushBack(apiMessage)
	c.rpcMessages = append(c.rpcMessages, apiMessage)
	// c.cb.sendToRpcQueue(&rpcApiMessage{connID: connID, sessionId: c.sessionId, rpcMessage: apiMessage})

	return true
}

func (c *genericSession) onInvokeRpcRequest(authUserId int32, authKeyId int64, layer int32, requests *rpcApiMessages) []*networkApiMessage {
	glog.Infof("onRpcRequest - receive data: {sess: %s, session_id: %d, conn_id: %d, md: %s, data: {%v}}",
		c, requests.sessionId, requests.connID, requests.cntl, requests.rpcMessages)

	return invokeRpcRequest(authUserId, authKeyId, layer, requests, func() *grpc_util.RPCClient{ return c.RPCClient })
	//rpcMessageList := make([]*networkApiMessage, 0, len(requests.rpcMessages))
	//
	//for i := 0; i < len(requests.rpcMessages); i++ {
	//	var (
	//		err         error
	//		rpcResult   mtproto.TLObject
	//	)
	//
	//	// 初始化metadata
	//	rpcMetadata := &grpc_util.RpcMetadata{
	//		ServerId:        getServerID(),
	//		NetlibSessionId: int64(requests.connID.clientConnID),
	//		AuthId:          authKeyId,
	//		SessionId:       requests.sessionId,
	//		TraceId:         requests.cntl.RpcMeta.GetRequest().GetTraceId(),
	//		SpanId:          getUUID(),
	//		ReceiveTime:     time.Now().Unix(),
	//		UserId:          authUserId,
	//		ClientMsgId:     requests.rpcMessages[i].rpcRequest.MsgId,
	//		Layer:           layer,
	//	}
	//
	//	//if s.Layer == 0 {
	//	//	s.Layer = getCacheApiLayer(s.authKeyId)
	//	//}
	//	// rpcMetadata.Layer = s.Layer
	//
	//	// TODO(@benqi): change state.
	//	requests.rpcMessages[i].state = kNetworkMessageStateRunning
	//
	//	// TODO(@benqi): rpc proxy
	//	if checkRpcUploadRequest(requests.rpcMessages[i].rpcRequest.Object) {
	//		rpcResult, err = s.nbfsRPCClient.Invoke(rpcMetadata, requests.rpcMessages[i].rpcRequest.Object)
	//	} else if checkRpcDownloadRequest(requests.rpcMessages[i].rpcRequest.Object) {
	//		rpcResult, err = s.nbfsRPCClient.Invoke(rpcMetadata, requests.rpcMessages[i].rpcRequest.Object)
	//	} else {
	//		rpcResult, err = s.bizRPCClient.Invoke(rpcMetadata, requests.rpcMessages[i].rpcRequest.Object)
	//	}
	//
	//	reply := &mtproto.TLRpcResult{
	//		ReqMsgId: requests.rpcMessages[i].rpcRequest.MsgId,
	//	}
	//
	//	if err != nil {
	//		glog.Error(err)
	//		rpcErr, _ := err.(*mtproto.TLRpcError)
	//		if rpcErr.GetErrorCode() == int32(mtproto.TLRpcErrorCodes_NOTRETURN_CLIENT) {
	//			continue
	//		}
	//		reply.Result = rpcErr
	//	} else {
	//		// glog.Infof("OnMessage - rpc_result: {%v}\n", rpcResult)
	//		reply.Result = rpcResult
	//	}
	//
	//	requests.rpcMessages[i].state = kNetworkMessageStateInvoked
	//	requests.rpcMessages[i].rpcResult = reply
	//
	//	rpcMessageList = append(rpcMessageList, requests.rpcMessages[i])
	//
	//	if _, ok := requests.rpcMessages[i].rpcRequest.Object.(*mtproto.TLAuthLogOut); ok {
	//		glog.Info("authLogOut - ", rpcMetadata)
	//		putCacheUserId(authKeyId, 0)
	//		break
	//	}
	//}
	//
	//return rpcMessageList
}

func (c *genericSession) onRpcResult(rpcResults *rpcApiMessages) {
	var hasAuthLogout = false
	msgList := c.pendingMessages
	c.pendingMessages = []*pendingMessage{}
	for _, m := range rpcResults.rpcMessages {
		msgList = append(msgList, &pendingMessage{mtproto.GenerateMessageId(), true, m.rpcResult})
		if _, ok := m.rpcRequest.Object.(*mtproto.TLAuthLogOut); ok {
			hasAuthLogout = true
			break
		}
	}
	if len(msgList) > 0 {
		c.sendPendingMessagesToClient(rpcResults.connID, rpcResults.cntl, msgList)
	}

	if hasAuthLogout {
		deleteClientSessionManager(c.cb.getAuthKeyId())
	}
}

func (c *genericSession) onSyncData(cntl *zrpc.ZRpcController, obj mtproto.TLObject) {
	glog.Info("onSyncData - ", cntl)

	if c.sessionOnline() {
		id := c.connIds.Back()
		if id != nil {
			syncMessage := &pendingMessage{
				messageId: mtproto.GenerateMessageId(),
				confirm:   true,
				tl:        obj,
			}
			c.syncMessages = append(c.syncMessages, syncMessage)

			glog.Infof("onSyncData - sendPending {sess: {%v}, connID: {%v}}, pushObj: {%s}, connLen: {%d}", c, id.Value, reflect.TypeOf(obj), c.connIds.Len())
			c.sendPendingMessagesToClient(id.Value.(ClientConnID), cntl, c.syncMessages)
			c.syncMessages = []*pendingMessage{}
		} else {
			// glog.Info("id is nil")
		}
	}
}

func (c *genericSession) onSyncRpcResultData(cntl *zrpc.ZRpcController, data []byte) {
	// TODO(@benqi):
	glog.Info("onSyncRpcResultData - ", cntl)
}