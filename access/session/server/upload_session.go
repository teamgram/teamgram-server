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
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"reflect"
	"time"
)

type uploadSession struct {
	*session
	*grpc_util.RPCClient
}

func (c *uploadSession) onMessageData(id ClientConnID, cntl *zrpc.ZRpcController, salt int64, msg *mtproto.TLMessage2) {
	c.session.processMessageData(id, cntl, salt, msg, func(sessMsg *mtproto.TLMessage2) {
		glog.Infof("uploadSession]]>> onRpcRequest - request data: {sess: %s, conn_id: %s, md: %s, msg_id: %d, seq_no: %d, request: {%s}}",
			c,
			id,
			cntl,
			sessMsg.MsgId,
			sessMsg.Seqno,
			reflect.TypeOf(sessMsg.Object))

		// TODO(@benqi): sync AuthUserId??
		var requestMessage *mtproto.TLMessage2

		switch sessMsg.Object.(type) {
		case *TLInvokeWithoutUpdatesExt:
			invokeWithoutUpdatesExt, _ := sessMsg.Object.(*TLInvokeWithoutUpdatesExt)
			requestMessage = &mtproto.TLMessage2{
				MsgId:  sessMsg.MsgId,
				Seqno:  sessMsg.Seqno,
				Object: invokeWithoutUpdatesExt.Query,
			}
		default:
			requestMessage = sessMsg
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

		apiMessage := &networkApiMessage{
			date:       time.Now().Unix(),
			rpcRequest: requestMessage,
			state:      kNetworkMessageStateReceived,
		}
		glog.Info("uploadSession]]>> onRpcRequest - ", apiMessage)
		c.apiMessages.PushBack(apiMessage)
		c.rpcMessages = append(c.rpcMessages, apiMessage)

		return
	})

	if len(c.pendingMessages) > 0 {
		c.sendPendingMessagesToClient(id, cntl, c.pendingMessages)
		c.pendingMessages = []*pendingMessage{}
	}

	if len(c.rpcMessages) > 0 {
		c.cb.sendToRpcQueue(&rpcApiMessages{connID: id, cntl: cntl, sessionId: c.sessionId, rpcMessages: c.rpcMessages})
		c.rpcMessages = []*networkApiMessage{}
	}

}

func (c *uploadSession) onInvokeRpcRequest(authUserId int32, authKeyId int64, layer int32, requests *rpcApiMessages) []*networkApiMessage {
	glog.Infof("uploadSession]]>> onInvokeRpcRequest - receive data: {sess: %s, session_id: %d, conn_id: %d, md: %s, data: {%v}}",
		c, requests.sessionId, requests.connID, requests.cntl, requests.rpcMessages)

	return invokeRpcRequest(authUserId, authKeyId, layer, requests, func() *grpc_util.RPCClient { return c.RPCClient })
}

func (c *uploadSession) onRpcResult(rpcResults *rpcApiMessages) {
	msgList := c.pendingMessages
	c.pendingMessages = []*pendingMessage{}
	for _, m := range rpcResults.rpcMessages {
		glog.Infof("uploadSession]]>> onRpcResult - rpcResults: {sess: %s, result: {%s}}",
			c, reflect.TypeOf(m.rpcRequest))
		msgList = append(msgList, &pendingMessage{mtproto.GenerateMessageId(), true, m.rpcResult})
	}
	if len(msgList) > 0 {
		c.sendPendingMessagesToClient(rpcResults.connID, rpcResults.cntl, msgList)
	}
}
