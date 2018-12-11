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
	"time"
)

type uploadSession struct {
	*session
	*grpc_util.RPCClient
}

func (c *uploadSession) onMessageData(id ClientConnID, cntl *zrpc.ZRpcController, salt int64, msg *mtproto.TLMessage2) {
	c.session.processMessageData(id, cntl, salt, msg, func(sessMsg *mtproto.TLMessage2) {
		glog.Infof("onRpcRequest - request data: {sess: %s, conn_id: %s, md: %s, msg_id: %d, seq_no: %d, request: {%s}}",
			c,
			id,
			cntl,
			msg.MsgId,
			msg.Seqno,
			reflect.TypeOf(msg.Object))

		// TODO(@benqi): sync AuthUserId??
		requestMessage := msg

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
		glog.Info("onRpcRequest - ", apiMessage)
		// c.apiMessages = append(c.apiMessages, apiMessage)
		c.apiMessages.PushBack(apiMessage)

		c.rpcMessages = append(c.rpcMessages, apiMessage)
		// c.manager.rpcQueue.Push(&rpcApiMessage{connID: connID, sessionId: c.sessionId, rpcMessage: apiMessage})

		return
	})


	if len(c.pendingMessages) > 0 {
		c.sendPendingMessagesToClient(id, cntl, c.pendingMessages)
		c.pendingMessages = []*pendingMessage{}
	}
}

func (c *uploadSession) onInvokeRpcRequest(authUserId int32, authKeyId int64, layer int32, requests *rpcApiMessages) []*networkApiMessage {
	glog.Infof("onRpcRequest - receive data: {sess: %s, session_id: %d, conn_id: %d, md: %s, data: {%v}}",
		c, requests.sessionId, requests.connID, requests.cntl, requests.rpcMessages)

	return invokeRpcRequest(authUserId, authKeyId, layer, requests, func() *grpc_util.RPCClient{ return c.RPCClient })
}

func (c *uploadSession) onRpcResult(rpcResults *rpcApiMessages) {
	msgList := c.pendingMessages
	c.pendingMessages = []*pendingMessage{}
	for _, m := range rpcResults.rpcMessages {
		msgList = append(msgList, &pendingMessage{mtproto.GenerateMessageId(), true, m.rpcResult})
	}
	if len(msgList) > 0 {
		c.sendPendingMessagesToClient(rpcResults.connID, rpcResults.cntl, msgList)
	}
}
