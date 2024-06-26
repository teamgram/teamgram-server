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
	"math/rand"
	"time"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

// ///////////////////////////////////////////////////////////////////////////////////////
func (c *session) checkContainer(ctx context.Context, msgId int64, seqno int32, container *mtproto.TLMsgContainer) int32 {
	if c.inQueue.Lookup(msgId) != nil {
		logx.WithContext(ctx).Errorf("checkContainer - msgId collision: {msg_id: %d, seqno: %d}", msgId, seqno)
		return kMsgIdCollision
	}
	//for e := c.msgIds.Front(); e != nil; e = e.Next() {
	//	if e.Value.(int64) == msgId {
	//		return kMsgIdCollision
	//	}
	//}

	if len(container.Messages) == 0 {
		return 0
	}

	for _, v := range container.Messages {
		// container.Seqno >= v.Seqno
		if v.Seqno > seqno {
			logx.WithContext(ctx).Errorf("checkContainer - v.seqno(%s) > seqno({msg_id: %d, seqno: %d})", v, msgId, seqno)
			return kInvalidContainer
		}
		if v.MsgId >= msgId {
			logx.WithContext(ctx).Errorf("checkContainer - v.MsgId(%s) > msgId({msg_id: %d, seqno: %d})", v, msgId, seqno)
			return kInvalidContainer
		}

		if _, ok := v.Object.(*mtproto.TLMsgContainer); ok {
			logx.WithContext(ctx).Errorf("checkContainer - is container: %v", v)
			return kInvalidContainer
		}
	}

	return 0
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////
// ============================================================================================
func (c *session) onNewSessionCreated(ctx context.Context, gatewayId string, msgId int64) {
	logx.WithContext(ctx).Infof("onNewSessionCreated - request data: %d", msgId)
	newSessionCreated := mtproto.MakeTLNewSessionCreated(&mtproto.NewSession{
		FirstMsgId: msgId,
		UniqueId:   rand.Int63(),
		ServerSalt: c.sessList.cacheSalt.GetSalt(),
	})

	logx.WithContext(ctx).Infof("onNewSessionCreated - reply: {%v}", newSessionCreated)

	c.sendDirectToGateway(ctx, gatewayId, true, newSessionCreated, func(sentRaw *mtproto.TLMessageRawData) {
		id2 := c.sessList.cb.getNextNotifyId()
		sentMsg := c.outQueue.AddNotifyMsg(id2, true, sentRaw)
		sentMsg.sent = 0
	})
}

/*
	tdesktop

	void Instance::Private::performKeyDestroy(ShiftedDcId shiftedDcId) {
		Expects(isKeysDestroyer());

		_instance->send(MTPDestroy_auth_key(), rpcDone([this, shiftedDcId](const MTPDestroyAuthKeyRes &result) {
			switch (result.type()) {
			case mtpc_destroy_auth_key_ok: LOG(("MTP Info: key %1 destroyed.").arg(shiftedDcId)); break;
			case mtpc_destroy_auth_key_fail: {
				LOG(("MTP Error: key %1 destruction fail, leave it for now.").arg(shiftedDcId));
				killSession(shiftedDcId);
			} break;
			case mtpc_destroy_auth_key_none: LOG(("MTP Info: key %1 already destroyed.").arg(shiftedDcId)); break;
			}
			emit _instance->keyDestroyed(shiftedDcId);
		}), rpcFail([this, shiftedDcId](const RPCError &error) {
			LOG(("MTP Error: key %1 destruction resulted in error: %2").arg(shiftedDcId).arg(error.type()));
			emit _instance->keyDestroyed(shiftedDcId);
			return true;
		}), shiftedDcId);
	}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
destroy_auth_key#d1435160 = DestroyAuthKeyRes;

destroy_auth_key_ok#f660e1d4 = DestroyAuthKeyRes;
destroy_auth_key_none#0a9f2259 = DestroyAuthKeyRes;
destroy_auth_key_fail#ea109b13 = DestroyAuthKeyRes;
*/
func (c *session) onDestroyAuthKey(ctx context.Context, gatewayId string, msgId *inboxMsg, destroyAuthKey *mtproto.TLDestroyAuthKey) {
	logx.WithContext(ctx).Infof("onDestroyAuthKey - request data: {sess: %s, gatewayId: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		destroyAuthKey)

	//// TODO(@benqi): FIXME
	// destroy_auth_key_ok#f660e1d4 = DestroyAuthKeyRes;
	// destroy_auth_key_none#0a9f2259 = DestroyAuthKeyRes;
	// destroy_auth_key_fail#ea109b13 = DestroyAuthKeyRes;
	res := mtproto.MakeTLDestroyAuthKeyOk(nil).To_DestroyAuthKeyRes()
	c.sendRpcResultToQueue(ctx, gatewayId, msgId.msgId, res)
	msgId.state = RECEIVED | ACKNOWLEDGED
}

func (c *session) onPing(ctx context.Context, gatewayId string, msgId *inboxMsg, ping *mtproto.TLPing) {
	logx.WithContext(ctx).Infof("onPing - request data: {sess: %s, gatewayId: %s, msg_id: %s, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId,
		msgId,
		ping)

	pong := &mtproto.TLPong{Data2: &mtproto.Pong{
		MsgId:  msgId.msgId,
		PingId: ping.PingId,
	}}

	c.sendRawToQueue(ctx, gatewayId, msgId.msgId, false, pong)
	msgId.state = RECEIVED | NEED_NO_ACK

	c.closeDate = time.Now().Unix() + kDefaultPingTimeout + kPingAddTimeout
}

func (c *session) onPingDelayDisconnect(ctx context.Context, gatewayId string, msgId *inboxMsg, pingDelayDisconnect *mtproto.TLPingDelayDisconnect) {
	logx.WithContext(ctx).Infof("onPingDelayDisconnect - request data: {sess: %s, gatewayId: %s, msg_id: %s, request: {%s}}",
		c,
		gatewayId,
		msgId,
		pingDelayDisconnect)

	pong := &mtproto.TLPong{Data2: &mtproto.Pong{
		MsgId:  msgId.msgId,
		PingId: pingDelayDisconnect.PingId,
	}}

	c.sendRawToQueue(ctx, gatewayId, msgId.msgId, false, pong)
	msgId.state = RECEIVED | NEED_NO_ACK

	willCloseDate := time.Now().Unix() + int64(pingDelayDisconnect.DisconnectDelay) + kPingAddTimeout
	if willCloseDate > c.closeDate {
		c.closeDate = willCloseDate
	}
}

/*
## HTTP Wait/Long Poll
The following special service query not requiring an acknowledgement
(which must be transmitted only through an HTTP connection) is used to
enable the server to send messages in the future to the client using HTTP protocol:

```
http_wait#9299359f max_delay:int wait_after:int max_wait:int = HttpWait;
```

When such a message (or a container carrying such a message) is received,
the server either waits max_delay milliseconds,
whereupon it forwards all the messages that it is holding on to the client
if there is at least one message queued in session
(if needed, by placing them into a container to which acknowledgments may also be added);
or else waits no more than max_wait milliseconds until such a message is available.
If a message never appears, an empty container is transmitted.

The max_delay parameter denotes the maximum number of milliseconds
that has elapsed between the first message for this session
and the transmission of an HTTP response.
The wait_after parameter works as follows:
after the receipt of the latest message for a particular session,
the server waits another wait_after milliseconds in case there are more messages.
If there are no additional messages,
the result is transmitted (a container with all the messages).
If more messages appear, the wait_after timer is reset.

At the same time, the max_delay parameter has higher priority than wait_after,
and max_wait has higher priority than max_delay.

This message does not require a response or an acknowledgement.
If the container transmitted over HTTP carries several such messages,
the behavior is undefined (in fact, the latest parameter will be used).

If no http_wait is present in container,
default values max_delay=0 (milliseconds),
wait_after=0 (milliseconds),
and max_wait=25000 (milliseconds) are used.

If the client’s ping of the server takes a long time,
it may make sense to set max_delay to a value that is comparable in magnitude to ping time.
*/
func (c *session) onHttpWait(ctx context.Context, gatewayId string, msgId int64, seqNo int32, request *mtproto.TLHttpWait) {
	logx.WithContext(ctx).Infof("onHttpWait - request data: {sess: %s, gatewayId: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId,
		seqNo,
		request)

	_ = request.GetMaxDelay()
	_ = request.GetWaitAfter()
	t := request.GetMaxWait() / 1000
	if t == 0 {
		t = 1
	}
	// c.httpTimeOut = time.Now().Unix() + int64(t)
	// c.isUpdates = true
	// c.manager.setUserOnline(c.sessionId, connID)
	// c.manager.updatesSession.SubscribeHttpUpdates(c, connID)
}

func (c *session) onDestroySession(ctx context.Context, gatewayId string, msgId *inboxMsg, request *mtproto.TLDestroySession) {
	logx.WithContext(ctx).Infof("onDestroySession - request data: {sess: %s, gatewayId: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	// Request to Destroy Session
	//
	// Used by the client to notify the server that it may
	// forget the data from a different session belonging to the same user (i. e. with the same auth_key_id).
	// The result of this being applied to the current session is undefined.
	//
	// destroy_session#e7512126 session_id:long = DestroySessionRes;
	// destroy_session_ok#e22045fc session_id:long = DestroySessionRes;
	// destroy_session_none#62d350c9 session_id:long = DestroySessionRes;
	//

	if request.SessionId == c.sessionId {
		// The result of this being applied to the current session is undefined.
		logx.WithContext(ctx).Error("the result of this being applied to the current session is undefined.")

		// TODO(@benqi): handle error???
		return
	}

	if c.sessList.destroySession(request.GetSessionId()) {
		destroySessionOk := mtproto.MakeTLDestroySessionOk(&mtproto.DestroySessionRes{
			SessionId: request.SessionId,
		}).To_DestroySessionRes()
		c.sendRawToQueue(ctx, gatewayId, msgId.msgId, false, destroySessionOk)
	} else {
		destroySessionNone := mtproto.MakeTLDestroySessionNone(&mtproto.DestroySessionRes{
			SessionId: request.SessionId,
		}).To_DestroySessionRes()
		c.sendRawToQueue(ctx, gatewayId, msgId.msgId, false, destroySessionNone)
	}

	msgId.state = RECEIVED | NEED_NO_ACK
}

func (c *session) onGetFutureSalts(ctx context.Context, gatewayId string, msgId *inboxMsg, request *mtproto.TLGetFutureSalts) {
	logx.WithContext(ctx).Infof("onGetFutureSalts - request data: {sess: %s, gateway_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	salts, err := c.sessList.cb.cb.Dao.GetFutureSalts(context.Background(), c.sessList.authId, request.Num)
	if err != nil {
		logx.WithContext(ctx).Errorf("getFutureSalts error: %v", err)
		return
	}

	futureSalts := mtproto.MakeTLFutureSalts(&mtproto.FutureSalts{
		ReqMsgId: msgId.msgId,
		Now:      int32(time.Now().Unix()),
		Salts:    salts,
	}).To_FutureSalts()

	logx.WithContext(ctx).Infof("onGetFutureSalts - reply data: %s", futureSalts)
	c.sendRawToQueue(ctx, gatewayId, msgId.msgId, false, futureSalts)
	msgId.state = RECEIVED | NEED_NO_ACK
}

// sendToClient:
//
//	rpc_answer_unknown#5e2ad36e = RpcDropAnswer;
//	rpc_answer_dropped_running#cd78e586 = RpcDropAnswer;
//	rpc_answer_dropped#a43ad8b7 msg_id:long seq_no:int bytes:int = RpcDropAnswer;
//
// and both of these responses require an acknowledgment from the client.
func (c *session) onRpcDropAnswer(ctx context.Context, gatewayId string, msgId *inboxMsg, request *mtproto.TLRpcDropAnswer) {
	logx.WithContext(ctx).Infof("onRpcDropAnswer - request data: {sess: %s, gatewayId: %s, msg_id: %d, seq_no: %d, request: {%v}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	rpcAnswer := &mtproto.RpcDropAnswer{}

	var found = false
	//for e := c.apiMessages.Front(); e != nil; e = e.Next() {
	//	v, _ := e.Value.(*networkApiMessage)
	//	if v.rpcRequest.MsgId == request.ReqMsgId {
	//		if v.state == kNetworkMessageStateReceived {
	//			rpcAnswer.Constructor = mtproto.CRC32_rpc_answer_dropped
	//			rpcAnswer.MsgId = request.ReqMsgId
	//			// TODO(@benqi): set seqno and bytes
	//			// rpcAnswer.Data2.SeqNo = 0
	//			// rpcAnswer.Data2.Bytes = 0
	//		} else if v.state == kNetworkMessageStateInvoked {
	//			rpcAnswer.Constructor = mtproto.CRC32_rpc_answer_dropped_running
	//		} else {
	//			rpcAnswer.Constructor = mtproto.CRC32_rpc_answer_unknown
	//		}
	//		found = true
	//		break
	//	}
	//}

	if !found {
		rpcAnswer.Constructor = mtproto.CRC32_rpc_answer_unknown
	}

	// android client code:
	/*
		 if (notifyServer) {
			TL_rpc_drop_answer *dropAnswer = new TL_rpc_drop_answer();
			dropAnswer->req_msg_id = request->messageId;
			sendRequest(dropAnswer, nullptr, nullptr, RequestFlagEnableUnauthorized | RequestFlagWithoutLogin | RequestFlagFailOnServerErrors, request->datacenterId, request->connectionType, true);
		 }
	*/

	rpcAnswer = &mtproto.RpcDropAnswer{
		PredicateName: mtproto.Predicate_rpc_answer_unknown,
	}

	c.sendRpcResultToQueue(ctx, gatewayId, msgId.msgId, rpcAnswer)
	msgId.state = RECEIVED | ACKNOWLEDGED
}
