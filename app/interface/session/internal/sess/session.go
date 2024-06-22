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
	"fmt"
	"time"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	kDefaultPingTimeout  = 60
	kPingAddTimeout      = 15
	kCacheSessionTimeout = 3 * 60
	waitMsgAcksTimeout   = 30
)

const (
	kStateNew = iota
	kStateOnline
	kStateOffline
	kStateClose
)

const (
	kSessionStateNew = iota
	kSessionStateCreated
)

const (
	kServerSaltIncorrect = int32(48)
)

const (
	kMsgIdTooLow    = int32(16)
	kMsgIdTooHigh   = int32(17)
	kMsgIdMod4      = int32(18)
	kMsgIdCollision = int32(19)

	kMsgIdTooOld = int32(20)

	kSeqNoTooLow  = int32(32)
	kSeqNoTooHigh = int32(33)
	kSeqNoNotEven = int32(34)
	kSeqNoNotOdd  = int32(35)

	kInvalidContainer = int32(64)
)

var emptyMsgContainer = mtproto.NewTLMsgRawDataContainer()
var androidPushTooLong = mtproto.MakeTLUpdatesTooLong(nil)

type messageData struct {
	confirmFlag  bool
	compressFlag bool
	obj          mtproto.TLObject
}

func (m *messageData) String() string {
	return fmt.Sprintf("{confirmFlag: %v, compressFlag: %v, obj: {%s}}", m.confirmFlag, m.compressFlag, m.obj)
}

type serverIdCtx struct {
	gatewayId       string
	lastReceiveTime int64
}

func (c serverIdCtx) Equal(id string) bool {
	return c.gatewayId == id
}

/*
* tdesktop's SessionData:

PreRequestMap _toSend; // map of request_id -> request, that is waiting to be sent
RequestMap _haveSent; // map of msg_id -> request, that was sent, msDate = 0 for msgs_state_req (no resend / state req), msDate = 0, seqNo = 0 for containers
RequestIdsMap _toResend; // map of msg_id -> request_id, that request_id -> request lies in toSend and is waiting to be resent
ReceivedMsgIds _receivedIds; // set of received msg_id's, for checking new msg_ids
RequestIdsMap _wereAcked; // map of msg_id -> request_id, this msg_ids already were acked or do not need ack
QMap<mtpMsgId, bool> _stateRequest; // set of msg_id's, whose state should be requested

QMap<mtpRequestId, SerializedMessage> _receivedResponses; // map of request_id -> response that should be processed in the main thread
QList<SerializedMessage> _receivedUpdates; // list of updates that should be processed in the main thread
*/
type session struct {
	sessionId       int64
	sessionState    int
	gatewayId       *serverIdCtx
	nextSeqNo       uint32
	firstMsgId      int64
	connState       int
	closeDate       int64
	lastReceiveTime int64
	isAndroidPush   bool
	isGeneric       bool
	inQueue         *sessionInboundQueue
	outQueue        *sessionOutgoingQueue
	pendingQueue    *sessionRpcResultWaitingQueue
	pushQueue       *sessionPushQueue
	sessList        *SessionList
	// isHttp       bool
	// httpQueue    *httpRequestQueue
}

func newSession(sessionId int64, sessList *SessionList) *session {
	sess := &session{
		sessionId:       sessionId,
		gatewayId:       nil,
		sessionState:    kSessionStateNew,
		closeDate:       time.Now().Unix() + kDefaultPingTimeout + kPingAddTimeout,
		connState:       kStateNew,
		lastReceiveTime: time.Now().UnixNano(),
		inQueue:         newSessionInboundQueue(),
		outQueue:        newSessionOutgoingQueue(),
		pendingQueue:    newSessionRpcResultWaitingQueue(),
		pushQueue:       newSessionPushQueue(),
		sessList:        sessList,
		// isHttp:       false,
		// httpQueue:    newHttpRequestQueue(),
	}

	return sess
}

func (c *session) String() string {
	return fmt.Sprintf("{user_id: %d, auth_key_id: %d, session_id: %d, state: %d, conn_state: %d, conn_id_list: %#v}",
		c.sessList.cb.AuthUserId,
		c.sessList.authId,
		c.sessionId,
		c.sessionState,
		c.connState,
		c.gatewayId)
}

func (c *session) setGatewayId(gateId string) {
	if c.gatewayId == nil {
		c.gatewayId = &serverIdCtx{gatewayId: gateId, lastReceiveTime: time.Now().Unix()}
	} else {
		c.gatewayId.gatewayId = gateId
		c.gatewayId.lastReceiveTime = time.Now().Unix()
	}
}

func (c *session) getGatewayId() string {
	if c.gatewayId == nil {
		return ""
	} else {
		return c.gatewayId.gatewayId
	}
}

func (c *session) checkGatewayIdExist(gateId string) bool {
	if c.gatewayId == nil {
		return false
	}

	return c.gatewayId.Equal(gateId)
}

func (c *session) changeConnState(ctx context.Context, state int) {
	c.connState = state
	if c.isAndroidPush || c.isGeneric {
		if state == kStateOnline {
			c.sessList.cb.setOnline(ctx)
		} else if state == kStateOffline {
			c.sessList.cb.trySetOffline(ctx)
		}
	}
}

func (c *session) onSessionConnNew(ctx context.Context, id string) {
	if c.connState != kStateOnline {
		c.changeConnState(ctx, kStateOnline)
		c.setGatewayId(id)
	}
}

func (c *session) onSessionMessageData(ctx context.Context, gatewayId, clientIp string, salt int64, msg *mtproto.TLMessage2) {
	// 1. check salt
	if !c.checkBadServerSalt(ctx, gatewayId, salt, msg) {
		return
	}

	willCloseDate := time.Now().Unix() + kDefaultPingTimeout + kPingAddTimeout
	if willCloseDate > c.closeDate {
		c.closeDate = willCloseDate
	}

	// 2. TODO(@benqi): checkBadMsgNotification
	/*
		const auto needResend = false
			|| (errorCode == 16) // bad msg_id
			|| (errorCode == 17) // bad msg_id
			|| (errorCode == 64); // bad container
		.......

		if (needResend) { // bad msg_id or bad container
		if (serverSalt) sessionData->setSalt(serverSalt);
		unixtimeSet(serverTime, true);

		DEBUG_LOG(("Message Info: unixtime updated, now %1, resending in container...").arg(serverTime));

		resend(resendId, 0, true);
	*/
	// tdesktop: code = 16, 17, 64会重发，我们不要检查container里的msgId和seqNo
	//
	if !c.checkBadMsgNotification(ctx, gatewayId, false, msg) {
		logx.WithContext(ctx).Errorf("badMsgNotification - {sess: %s, conn_id: %s}", c, gatewayId)
		return
	}

	// extract msgs
	var (
		msgs []*mtproto.TLMessage2
	)

	// TODO(@benqi): ignore TLMsgCopy
	if msgContainer, ok := msg.Object.(*mtproto.TLMsgContainer); ok {
		msgs = msgContainer.Messages

		// check
		c.inQueue.AddMsgId(msg.MsgId)
	} else {
		msgs = append(msgs, msg)
	}

	for i := 0; i < len(msgs); i++ {
		if packed, ok := msgs[i].Object.(*mtproto.TLGzipPacked); ok {
			msgs[i] = &mtproto.TLMessage2{
				MsgId:  msgs[i].MsgId,
				Seqno:  msgs[i].Seqno,
				Bytes:  int32(len(packed.PackedData)),
				Object: packed.Obj,
			}
		}
	}

	// check onNewSessionCreated
	minMsgId := msg.MsgId
	for _, m2 := range msgs {
		if minMsgId < m2.MsgId {
			minMsgId = m2.MsgId
		}
	}

	if c.sessionState == kSessionStateNew || minMsgId < c.firstMsgId {
		logx.WithContext(ctx).Infof("onNewSessionCreated - %#v, c: %s", msgs, c)
		c.onNewSessionCreated(ctx, gatewayId, minMsgId)
		if c.firstMsgId != 0 {
			c.firstMsgId = minMsgId
		}
		c.sessionState = kSessionStateCreated
	}

	defer func() {
		c.sessList.cb.sendToRpcQueue(ctx, c.sessList.cb.tmpRpcApiMessageList)
		c.sessList.cb.tmpRpcApiMessageList = []*rpcApiMessage{}

		c.sendQueueToGateway(ctx, gatewayId)
		c.inQueue.Shrink()
	}()

	for _, m2 := range msgs {
		if m2.MsgId < c.firstMsgId {
			continue
		}
		if !c.checkBadMsgNotification(ctx, gatewayId, true, m2) {
			// log.Errorf("badMsgNotification - {sess: %s, conn_id: %s}", c, id)
			continue
		}

		if m2.Object == nil {
			logx.WithContext(ctx).Errorf("obj is nil: %v", m2)
			continue
		}

		switch m2.Object.(type) {
		case *mtproto.TLMsgsAck:
			c.onMsgsAck(ctx, gatewayId, m2.MsgId, m2.Seqno, m2.Object.(*mtproto.TLMsgsAck))

		case *mtproto.TLHttpWait:
			c.onHttpWait(ctx, gatewayId, m2.MsgId, m2.Seqno, m2.Object.(*mtproto.TLHttpWait))

		default:
			inMsg := c.inQueue.AddMsgId(m2.MsgId)
			if inMsg.state == NONE {
				// 第一次收到
				c.processMsg(ctx, gatewayId, clientIp, inMsg, m2.Object)
			} else {
				// TODO(@benqi): resend
				// 已经收到，返回发送状态
				// c.notifyMsgsStateInfo(gatewayId, inMsg)
				continue
			}
		}
	}
}

func (c *session) processMsg(ctx context.Context, gatewayId, clientIp string, inMsg *inboxMsg, r mtproto.TLObject) {
	switch r.(type) {
	case *mtproto.TLDestroyAuthKey: // 所有连接都有可能
		c.onDestroyAuthKey(ctx, gatewayId, inMsg, r.(*mtproto.TLDestroyAuthKey))
	case *mtproto.TLRpcDropAnswer: // 所有连接都有可能
		c.onRpcDropAnswer(ctx, gatewayId, inMsg, r.(*mtproto.TLRpcDropAnswer))
	case *mtproto.TLGetFutureSalts: // GENERIC
		c.onGetFutureSalts(ctx, gatewayId, inMsg, r.(*mtproto.TLGetFutureSalts))
	case *mtproto.TLPing: // android未用
		c.onPing(ctx, gatewayId, inMsg, r.(*mtproto.TLPing))
	case *mtproto.TLPingDelayDisconnect: // PUSH和GENERIC
		c.onPingDelayDisconnect(ctx, gatewayId, inMsg, r.(*mtproto.TLPingDelayDisconnect))
	case *mtproto.TLDestroySession: // GENERIC
		c.onDestroySession(ctx, gatewayId, inMsg, r.(*mtproto.TLDestroySession))
	case *mtproto.TLMsgsStateReq: // android未用
		c.onMsgsStateReq(ctx, gatewayId, inMsg, r.(*mtproto.TLMsgsStateReq))
	case *mtproto.TLMsgsStateInfo: // android未用
		c.onMsgsStateInfo(ctx, gatewayId, inMsg, r.(*mtproto.TLMsgsStateInfo))
	case *mtproto.TLMsgsAllInfo: // android未用
		c.onMsgsAllInfo(ctx, gatewayId, inMsg, r.(*mtproto.TLMsgsAllInfo))
	case *mtproto.TLMsgResendReq: // 都有可能
		c.onMsgResendReq(ctx, gatewayId, inMsg, r.(*mtproto.TLMsgResendReq))
	case *mtproto.TLMsgDetailedInfo: // 都有可能
		c.onMsgDetailInfo(ctx, gatewayId, inMsg, r.(*mtproto.TLMsgDetailedInfo))
	case *mtproto.TLMsgNewDetailedInfo: // 都有可能
		c.onMsgNewDetailInfo(ctx, gatewayId, inMsg, r.(*mtproto.TLMsgDetailedInfo))
	case *mtproto.TLInvokeWithLayer:
		c.onInvokeWithLayer(ctx, gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeWithLayer))
	case *mtproto.TLInvokeAfterMsg:
		c.onInvokeAfterMsg(ctx, gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeAfterMsg))
	case *mtproto.TLInvokeAfterMsgs:
		c.onInvokeAfterMsgs(ctx, gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeAfterMsgs))
	case *mtproto.TLInvokeWithoutUpdates:
		c.onInvokeWithoutUpdates(ctx, gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeWithoutUpdates))
	case *mtproto.TLInvokeWithMessagesRange:
		c.onInvokeWithMessagesRange(ctx, gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeWithMessagesRange))
	case *mtproto.TLInvokeWithTakeout:
		c.onInvokeWithTakeout(ctx, gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeWithTakeout))
	case *mtproto.TLInvokeWithBusinessConnection:
		c.onInvokeWithBusinessConnection(ctx, gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeWithBusinessConnection))
	case *mtproto.TLInvokeWithGooglePlayIntegrity:
		c.onInvokeWithGooglePlayIntegrity(ctx, gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeWithGooglePlayIntegrity))
	case *mtproto.TLInvokeWithApnsSecret:
		c.onInvokeWithApnsSecret(ctx, gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeWithApnsSecret))
	case *mtproto.TLInitConnection:
		c.onInitConnection(ctx, gatewayId, clientIp, inMsg, r.(*mtproto.TLInitConnection))
	case *mtproto.TLGzipPacked:
		c.onRpcRequest(ctx, gatewayId, clientIp, inMsg, r.(*mtproto.TLGzipPacked).Obj)
	default:
		c.onRpcRequest(ctx, gatewayId, clientIp, inMsg, r)
	}
}

func (c *session) onSessionConnClose(ctx context.Context, id string) {
	if c.checkGatewayIdExist(id) {
		c.gatewayId = nil
		c.changeConnState(ctx, kStateOffline)
	}
}

func (c *session) sessionOnline() bool {
	return c.connState == kStateOnline
}

func (c *session) sessionClosed() bool {
	return c.connState == kStateClose
}

// ============================================================================================
// return false, will delete this clientSession
func (c *session) onTimer(ctx context.Context) bool {
	date := time.Now().Unix()
	// log.Debugf("onTimer - c: %s, outQ len: %d", c.String(), c.outQueue.oMsgs.Len())
	gatewayId := c.getGatewayId()

	timeoutIdList := c.pendingQueue.OnTimer()
	for _, id := range timeoutIdList {
		c.sendRpcResult(
			ctx,
			&mtproto.TLRpcResult{
				ReqMsgId: id,
				Result: &mtproto.TLRpcError{Data2: &mtproto.RpcError{
					ErrorCode:    -503,
					ErrorMessage: "Timeout",
				}},
			})
	}

	//httpTimeOutList := c.httpQueue.PopTimeoutList()
	//if len(httpTimeOutList) > 0 {
	//	logx.WithContext(ctx).Infof("timeoutList: %d", len(httpTimeOutList))
	//}
	//for _, ch := range httpTimeOutList {
	//	c.sendHttpDirectToGateway(ctx, ch, false, emptyMsgContainer, func(sentRaw *mtproto.TLMessageRawData) {
	//		//
	//	})
	//}

	if c.connState == kStateOnline {
		if date >= c.closeDate {
			// log.Debugf("closeDate: %s", c.String())
			c.changeConnState(context.Background(), kStateOffline)
		} else {
			c.sendQueueToGateway(ctx, gatewayId)
		}
	} else if c.connState == kStateOffline || c.connState == kStateNew {
		if date >= c.closeDate+kCacheSessionTimeout {
			c.changeConnState(context.Background(), kStateClose)
		}
	}
	return true
}

func (c *session) generateMessageSeqNo(increment bool) int32 {
	value := c.nextSeqNo
	if increment {
		c.nextSeqNo++
		return int32(value*2 + 1)
	} else {
		return int32(value * 2)
	}
}

func (c *session) sendRpcResultToQueue(ctx context.Context, gatewayId string, reqMsgId int64, result mtproto.TLObject) {
	rpcResult := &mtproto.TLRpcResult{
		ReqMsgId: reqMsgId,
		Result:   result,
	}
	x := mtproto.NewEncodeBuf(500)
	rpcResult.Encode(x, c.sessList.cb.Layer())
	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(true),
		Seqno: c.generateMessageSeqNo(true),
		Bytes: int32(x.GetOffset()),
		Body:  x.GetBuf(),
	}
	c.outQueue.AddRpcResultMsg(reqMsgId, rawMsg)
	// cb(rawMsg)
}

func (c *session) sendPushRpcResultToQueue(gatewayId string, reqMsgId int64, result []byte) {
	//rpcResult := &mtproto.TLRpcResult{
	//	ReqMsgId: reqMsgId,
	//	Result:   result,
	//}
	//b := rpcResult.Encode(c.cb.getLayer())
	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(true),
		Seqno: c.generateMessageSeqNo(true),
		Bytes: int32(len(result)),
		Body:  result,
	}
	c.outQueue.AddRpcResultMsg(reqMsgId, rawMsg)
	// cb(rawMsg)
}

func (c *session) sendPushToQueue(ctx context.Context, gatewayId string, pushMsgId int64, pushMsg mtproto.TLObject) {
	x := mtproto.NewEncodeBuf(512)
	pushMsg.Encode(x, c.sessList.cb.Layer())
	rawBytes := x.GetBuf()
	if x.GetOffset() > 256 {
		gzipPacked := &mtproto.TLGzipPacked{
			PackedData: rawBytes,
		}
		x2 := mtproto.NewEncodeBuf(512)
		gzipPacked.Encode(x2, c.sessList.cb.Layer())
		rawBytes = x2.GetBuf()
	}

	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(false),
		Seqno: c.generateMessageSeqNo(true),
		Bytes: int32(len(rawBytes)),
		Body:  rawBytes,
	}
	c.outQueue.AddPushUpdates(pushMsgId, rawMsg)
}

func (c *session) sendRawToQueue(ctx context.Context, gatewayId string, msgId int64, confirm bool, rawMsg mtproto.TLObject) {
	x := mtproto.NewEncodeBuf(512)
	rawMsg.Encode(x, c.sessList.cb.Layer())
	b := x.GetBuf()
	rawMsg2 := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(false),
		Seqno: c.generateMessageSeqNo(confirm),
		Bytes: int32(len(b)),
		Body:  b,
	}
	c.outQueue.AddNotifyMsg(msgId, confirm, rawMsg2)
}

func (c *session) sendHttpDirectToGateway(ctx context.Context, ch chan interface{}, confirm bool, obj mtproto.TLObject, cb func(sentRaw *mtproto.TLMessageRawData)) (bool, error) {
	if c.connState != kStateOnline {
		return false, nil
	}

	x := mtproto.NewEncodeBuf(512)
	salt := c.sessList.cacheSalt.GetSalt()
	obj.Encode(x, c.sessList.cb.Layer())
	b := x.GetBuf()
	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(false),
		Seqno: c.generateMessageSeqNo(confirm),
		Bytes: int32(len(b)),
		Body:  b,
	}

	rB, err := c.sessList.cb.cb.Dao.SendHttpDataToGateway(
		ctx,
		ch,
		c.sessList.authId,
		salt,
		c.sessionId,
		rawMsg)

	if err != nil {
		logx.WithContext(ctx).Errorf("sendHttpDirectToGateway - %v", err)
	}

	if cb != nil {
		cb(rawMsg)
	}
	// c.httpTimeOut = 0
	return rB, err
}

func (c *session) sendDirectToGateway(ctx context.Context, gatewayId string, confirm bool, obj mtproto.TLObject, cb func(sentRaw *mtproto.TLMessageRawData)) (bool, error) {
	if c.connState != kStateOnline {
		return false, nil
	}

	x := mtproto.NewEncodeBuf(512)
	salt := c.sessList.cacheSalt.GetSalt()
	obj.Encode(x, c.sessList.cb.Layer())
	b := x.GetBuf()

	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(false),
		Seqno: c.generateMessageSeqNo(confirm),
		Bytes: int32(len(b)),
		Body:  b,
	}

	var (
		rB  bool
		err error
	)

	//if !c.isHttp {
	rB, err = c.sessList.cb.cb.Dao.SendDataToGateway(
		ctx,
		gatewayId,
		c.sessList.authId,
		salt,
		c.sessionId,
		rawMsg)
	//} else {
	//	if ch := c.httpQueue.Pop(); ch != nil {
	//		rB, err = c.sessList.cb.cb.Dao.SendHttpDataToGateway(
	//			ctx,
	//			ch,
	//			c.sessList.authId,
	//			salt,
	//			c.sessionId,
	//			rawMsg)
	//	}
	//}

	if err != nil {
		logx.WithContext(ctx).Errorf("sendToClient - %v", err)
	}

	if cb != nil {
		cb(rawMsg)
	}
	// c.httpTimeOut = 0
	return rB, err
}

func (c *session) sendRawDirectToGateway(ctx context.Context, gatewayId string, raw *mtproto.TLMessageRawData) (bool, error) {
	if c.connState != kStateOnline {
		return false, nil
	}

	salt := c.sessList.cacheSalt.GetSalt()

	var (
		rB  bool
		err error
	)

	//if !c.isHttp {
	rB, err = c.sessList.cb.cb.Dao.SendDataToGateway(
		ctx,
		gatewayId,
		c.sessList.authId,
		salt,
		c.sessionId,
		raw)
	//} else {
	//	if ch := c.httpQueue.Pop(); ch != nil {
	//		rB, err = c.sessList.cb.cb.Dao.SendHttpDataToGateway(
	//			ctx,
	//			ch,
	//			c.sessList.authId,
	//			salt,
	//			c.sessionId,
	//			raw)
	//	}
	//}

	if err != nil {
		logx.WithContext(ctx).Errorf("sendRawDirectToGateway - %v", err)
	}
	return rB, err
}

func (c *session) sendQueueToGateway(ctx context.Context, gatewayId string) {
	if gatewayId == "" {
		return
	}

	if c.outQueue.oMsgs.Len() == 0 {
		return
	}

	var (
		pendings = make([]*outboxMsg, 0)
		b        = false
		err      error
		sentTime = time.Now().Unix()
	)
	for e := c.outQueue.oMsgs.Front(); e != nil; e = e.Next() {
		if e.Value.(*outboxMsg).sent == 0 || time.Now().Unix() >= e.Value.(*outboxMsg).sent+waitMsgAcksTimeout {
			pendings = append(pendings, e.Value.(*outboxMsg))
		}
	}

	if len(pendings) == 1 {
		logx.WithContext(ctx).Infof("sendRawDirectToGateway - pendings[0]")
		b, err = c.sendRawDirectToGateway(ctx, gatewayId, pendings[0].msg)
		// log.Debugf("err: %v, b: %v", err, b)
		if err != nil || !b {
			return
		}

		for _, m := range pendings {
			if m.state == NEED_NO_ACK {
				logx.WithContext(ctx).Infof("need_no_ack: %d", m.msgId)
				c.outQueue.Remove(m.msgId)
			} else {
				logx.WithContext(ctx).Infof("pending sent: %d", m.msgId)
				m.sent = sentTime
			}
		}
	} else if len(pendings) > 1 {
		var (
			split = 16
		)
		for i := 0; i < len(pendings)/split; i++ {
			msgContainer := &mtproto.TLMsgRawDataContainer{
				Messages: make([]*mtproto.TLMessageRawData, 0, split),
			}
			for _, m := range pendings[i*split : (i+1)*split] {
				msgContainer.Messages = append(msgContainer.Messages, m.msg)
			}
			logx.WithContext(ctx).Infof("sendRawDirectToGateway - TLMsgRawDataContainer")
			b, err = c.sendDirectToGateway(ctx, gatewayId, false, msgContainer, func(sentRaw *mtproto.TLMessageRawData) {
				// TODO(@benqi):
				// nothing do
			})
			// log.Debugf("err: %v, b: %v", err, b)
			if err != nil || !b {
				continue
			}

			for _, m := range pendings[i*split : (i+1)*split] {
				if m.state == NEED_NO_ACK {
					logx.WithContext(ctx).Infof("need_no_ack: %d", m.msgId)
					c.outQueue.Remove(m.msgId)
				} else {
					logx.WithContext(ctx).Infof("pending sent: %d", m.msgId)
					m.sent = sentTime
				}
			}
		}
		if (len(pendings) % split) > 0 {
			msgContainer := &mtproto.TLMsgRawDataContainer{
				Messages: make([]*mtproto.TLMessageRawData, 0, split),
			}
			for _, m := range pendings[split*(len(pendings)/split):] {
				msgContainer.Messages = append(msgContainer.Messages, m.msg)
			}
			logx.WithContext(ctx).Infof("sendRawDirectToGateway - TLMsgRawDataContainer")
			b, err = c.sendDirectToGateway(ctx, gatewayId, false, msgContainer, func(sentRaw *mtproto.TLMessageRawData) {
				// TODO(@benqi):
				// nothing do
			})
			// log.Debugf("err: %v, b: %v", err, b)
			if err != nil || !b {
				return
			}
			for _, m := range pendings[split*(len(pendings)/split):] {
				if m.state == NEED_NO_ACK {
					logx.WithContext(ctx).Infof("need_no_ack: %d", m.msgId)
					c.outQueue.Remove(m.msgId)
				} else {
					logx.WithContext(ctx).Infof("pending sent: %d", m.msgId)
					m.sent = sentTime
				}
			}
		}
	}
}
