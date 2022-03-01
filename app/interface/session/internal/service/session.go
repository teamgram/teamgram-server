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
	"fmt"
	"time"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

//const (
//	kConnUnknown = 0
//	kTcpConn     = 1
//	kHttpConn    = 2
//	kPushConn    = 3
//)

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

//////////////////////////////////////////////////////////////////////////////////////////////////////////
type sessionCallback interface {
	getCacheSalt() *mtproto.TLFutureSalt

	getAuthKeyId() int64
	getTempAuthKeyId() int64

	getUserId() int64
	setUserId(userId int64)

	getLayer() int32
	setLayer(layer int32)

	setClient(c string)
	getClient() string

	setLangpack(c string)
	getLangpack() string

	destroySession(sessionId int64) bool

	sendToRpcQueue(rpcMessage *rpcApiMessage)

	onBindPushSessionId(sessionId int64)
	setOnline()
	trySetOffline()
}

/** tdesktop's SessionData:

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
	gatewayIdList   []serverIdCtx
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
	isHttp          bool
	httpQueue       *httpRequestQueue
	cb              sessionCallback
	*authSessions
}

func newSession(sessionId int64, sesses *authSessions) *session {
	// var sess *sessionHandler
	sess := &session{
		sessionId:       sessionId,
		gatewayIdList:   make([]serverIdCtx, 0, 1),
		sessionState:    kSessionStateNew,
		closeDate:       time.Now().Unix() + kDefaultPingTimeout + kPingAddTimeout,
		connState:       kStateNew,
		lastReceiveTime: time.Now().UnixNano(),
		inQueue:         newSessionInboundQueue(),
		outQueue:        newSessionOutgoingQueue(),
		pendingQueue:    newSessionRpcResultWaitingQueue(),
		pushQueue:       newSessionPushQueue(),
		isHttp:          false,
		httpQueue:       newHttpRequestQueue(),
		cb:              sesses,
		authSessions:    sesses,
	}

	return sess
}

func (c *session) String() string {
	return fmt.Sprintf("{user_id: %d, auth_key_id: %d, session_id: %d, state: %d, conn_state: %d, conn_id_list: %#v}",
		c.authSessions.AuthUserId,
		c.authSessions.authKeyId,
		c.sessionId,
		c.sessionState,
		c.connState,
		c.gatewayIdList)
}

func (c *session) addGatewayId(gateId string) {
	for _, id := range c.gatewayIdList {
		if id.Equal(gateId) {
			return
		}
	}
	c.gatewayIdList = append(c.gatewayIdList, serverIdCtx{gatewayId: gateId, lastReceiveTime: time.Now().Unix()})
}

func (c *session) getGatewayId() string {
	if len(c.gatewayIdList) > 0 {
		// TODO(@benqi): rand or by new lastReceiveTime
		return c.gatewayIdList[len(c.gatewayIdList)-1].gatewayId
	} else {
		return ""
	}
}

func (c *session) checkGatewayIdExist(gateId string) bool {
	for _, id := range c.gatewayIdList {
		if id.Equal(gateId) {
			return true
		}
	}
	return false
}

func (c *session) changeConnState(state int) {
	c.connState = state
	if c.isAndroidPush || c.isGeneric {
		if state == kStateOnline {
			c.cb.setOnline()
		} else if state == kStateOffline {
			c.cb.trySetOffline()
		}
	}
}

func (c *session) onSessionConnNew(id string) {
	if c.connState != kStateOnline {
		// c.sessionState = kSessionStateNew
		c.changeConnState(kStateOnline)
		c.addGatewayId(id)
	}
}

func (c *session) onSessionMessageData(gatewayId, clientIp string, salt int64, msg *mtproto.TLMessage2) {
	// 1. check salt
	if !c.checkBadServerSalt(gatewayId, salt, msg) {
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
	if !c.checkBadMsgNotification(gatewayId, false, msg) {
		// log.Errorf("badMsgNotification - {sess: %s, conn_id: %s}", c, gatewayId)
		return
	}

	// extract msgs
	var msgs []*mtproto.TLMessage2
	// TODO(@benqi): ignore TLMsgCopy
	if msgContainer, ok := msg.Object.(*mtproto.TLMsgContainer); ok {
		for _, m2 := range msgContainer.Messages {
			msgs = append(msgs, &mtproto.TLMessage2{
				MsgId:  m2.MsgId,
				Seqno:  m2.Seqno,
				Bytes:  m2.Bytes,
				Object: m2.Object,
			})
		}

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
		logx.Infof("onNewSessionCreated - %#v, c: %s", msgs, c)
		c.onNewSessionCreated(gatewayId, minMsgId)
		if c.firstMsgId != 0 {
			c.firstMsgId = minMsgId
		}
		c.sessionState = kSessionStateCreated
		// return
	}

	defer func() {
		c.sendQueueToGateway(gatewayId)
		c.inQueue.Shrink()

		//pendings := make([]*outboxMsg, 0)
		//for e := c.outQueue.oMsgs.Front(); e != nil; e = e.Next() {
		//	if e.Value.(*outboxMsg).sent == false {
		//		e.Value.(*outboxMsg).sent = true
		//		pendings = append(pendings, e.Value.(*outboxMsg))
		//	}
		//}
		//
		//if len(pendings) == 1 {
		//	c.sendRawDirectToGateway(gatewayId, pendings[0].msg)
		//} else if len(pendings) > 1 {
		//	msgContainer := &mtproto.TLMsgRawDataContainer{
		//		Messages: make([]*mtproto.TLMessageRawData, 0, len(pendings)),
		//	}
		//	for _, m := range pendings {
		//		msgContainer.Messages = append(msgContainer.Messages, m.msg)
		//	}
		//
		//	c.sendDirectToGateway(gatewayId, false, msgContainer, func(sentRaw *mtproto.TLMessageRawData) {
		//		// TODO(@benqi):
		//	})
		//}
	}()

	for _, m2 := range msgs {
		if m2.MsgId < c.firstMsgId {
			continue
		}
		if !c.checkBadMsgNotification(gatewayId, true, m2) {
			// log.Errorf("badMsgNotification - {sess: %s, conn_id: %s}", c, id)
			continue
		}

		if m2.Object == nil {
			logx.Errorf("obj is nil: %v", m2)
			continue
		}

		switch m2.Object.(type) {
		case *mtproto.TLMsgsAck:
			c.onMsgsAck(gatewayId, m2.MsgId, m2.Seqno, m2.Object.(*mtproto.TLMsgsAck))

		case *mtproto.TLHttpWait:
			c.onHttpWait(gatewayId, m2.MsgId, m2.Seqno, m2.Object.(*mtproto.TLHttpWait))

		default:
			inMsg := c.inQueue.AddMsgId(m2.MsgId)
			if inMsg.state == NONE {
				// 第一次收到
				c.processMsg(gatewayId, clientIp, inMsg, m2.Object)
			} else {
				// TODO(@benqi): resend
				// 已经收到，返回发送状态
				// c.notifyMsgsStateInfo(gatewayId, inMsg)
				continue
			}
		}
	}
}

func (c *session) processMsg(gatewayId, clientIp string, inMsg *inboxMsg, r mtproto.TLObject) {
	switch r.(type) {
	case *mtproto.TLDestroyAuthKey: // 所有连接都有可能
		c.onDestroyAuthKey(gatewayId, inMsg, r.(*mtproto.TLDestroyAuthKey))
	case *mtproto.TLRpcDropAnswer: // 所有连接都有可能
		c.onRpcDropAnswer(gatewayId, inMsg, r.(*mtproto.TLRpcDropAnswer))
	case *mtproto.TLGetFutureSalts: // GENERIC
		c.onGetFutureSalts(gatewayId, inMsg, r.(*mtproto.TLGetFutureSalts))
	case *mtproto.TLPing: // android未用
		c.onPing(gatewayId, inMsg, r.(*mtproto.TLPing))
	case *mtproto.TLPingDelayDisconnect: // PUSH和GENERIC
		c.onPingDelayDisconnect(gatewayId, inMsg, r.(*mtproto.TLPingDelayDisconnect))
	case *mtproto.TLDestroySession: // GENERIC
		c.onDestroySession(gatewayId, inMsg, r.(*mtproto.TLDestroySession))
	case *mtproto.TLMsgsStateReq: // android未用
		c.onMsgsStateReq(gatewayId, inMsg, r.(*mtproto.TLMsgsStateReq))
	case *mtproto.TLMsgsStateInfo: // android未用
		c.onMsgsStateInfo(gatewayId, inMsg, r.(*mtproto.TLMsgsStateInfo))
	case *mtproto.TLMsgsAllInfo: // android未用
		c.onMsgsAllInfo(gatewayId, inMsg, r.(*mtproto.TLMsgsAllInfo))
	case *mtproto.TLMsgResendReq: // 都有可能
		c.onMsgResendReq(gatewayId, inMsg, r.(*mtproto.TLMsgResendReq))
	case *mtproto.TLMsgDetailedInfo: // 都有可能
		c.onMsgDetailInfo(gatewayId, inMsg, r.(*mtproto.TLMsgDetailedInfo))
	case *mtproto.TLMsgNewDetailedInfo: // 都有可能
		c.onMsgNewDetailInfo(gatewayId, inMsg, r.(*mtproto.TLMsgDetailedInfo))
	case *mtproto.TLInvokeWithLayer:
		c.onInvokeWithLayer(gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeWithLayer))
	case *mtproto.TLInvokeAfterMsg:
		c.onInvokeAfterMsg(gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeAfterMsg))
	case *mtproto.TLInvokeAfterMsgs:
		c.onInvokeAfterMsgs(gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeAfterMsgs))
	case *mtproto.TLInvokeWithoutUpdates:
		c.onInvokeWithoutUpdates(gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeWithoutUpdates))
	case *mtproto.TLInvokeWithMessagesRange:
		c.onInvokeWithMessagesRange(gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeWithMessagesRange))
	case *mtproto.TLInvokeWithTakeout:
		c.onInvokeWithTakeout(gatewayId, clientIp, inMsg, r.(*mtproto.TLInvokeWithTakeout))
	case *mtproto.TLGzipPacked:
		c.onRpcRequest(gatewayId, clientIp, inMsg, r.(*mtproto.TLGzipPacked).Obj)
	default:
		c.onRpcRequest(gatewayId, clientIp, inMsg, r)
	}
}

func (c *session) onSessionConnClose(id string) {
	var (
		idx = -1
	)

	for i, cId := range c.gatewayIdList {
		if cId.Equal(id) {
			idx = i
			break
		}
	}

	if idx != -1 {
		c.gatewayIdList = append(c.gatewayIdList[:idx], c.gatewayIdList[idx+1:]...)
	}

	if len(c.gatewayIdList) == 0 {
		c.changeConnState(kStateOffline)
	}
}

func (c *session) sessionOnline() bool {
	return c.connState == kStateOnline
}

func (c *session) sessionClosed() bool {
	return c.connState == kStateClose
}

//============================================================================================
// return false, will delete this clientSession
func (c *session) onTimer() bool {
	date := time.Now().Unix()
	// log.Debugf("onTimer - c: %s, outQ len: %d", c.String(), c.outQueue.oMsgs.Len())
	gatewayId := c.getGatewayId()

	timeoutIdList := c.pendingQueue.OnTimer()
	for _, id := range timeoutIdList {
		c.sendRpcResult(&mtproto.TLRpcResult{
			ReqMsgId: id,
			Result: &mtproto.TLRpcError{Data2: &mtproto.RpcError{
				ErrorCode:    -503,
				ErrorMessage: "Timeout",
			}},
		})
	}

	httpTimeOutList := c.httpQueue.PopTimeoutList()
	if len(httpTimeOutList) > 0 {
		logx.Infof("timeoutList: %d", len(httpTimeOutList))
	}
	for _, ch := range httpTimeOutList {
		c.sendHttpDirectToGateway(ch, false, emptyMsgContainer, func(sentRaw *mtproto.TLMessageRawData) {
			//
		})
	}

	if c.connState == kStateOnline {
		if date >= c.closeDate {
			// log.Debugf("closeDate: %s", c.String())
			c.changeConnState(kStateOffline)
		} else {
			c.sendQueueToGateway(gatewayId)
		}
	} else if c.connState == kStateOffline || c.connState == kStateNew {
		if date >= c.closeDate+kCacheSessionTimeout {
			c.changeConnState(kStateClose)
		}
	}
	return true
}

//============================================================================================
func (c *session) encodeMessage(messageId int64, confirm bool, tl mtproto.TLObject) ([]byte, error) {
	salt := c.cb.getCacheSalt().GetSalt()
	seqNo := c.generateMessageSeqNo(confirm)

	if messageId == 0 {
		messageId = nextMessageId(false)
	}

	return SerializeToBuffer(salt, c.sessionId, &mtproto.TLMessage2{
		MsgId:  messageId,
		Seqno:  seqNo,
		Object: tl,
	}, c.cb.getLayer()), nil
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

func (c *session) sendRpcResultToQueue(gatewayId string, reqMsgId int64, result mtproto.TLObject) {
	rpcResult := &mtproto.TLRpcResult{
		ReqMsgId: reqMsgId,
		Result:   result,
	}
	b := rpcResult.Encode(c.cb.getLayer())
	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(true),
		Seqno: c.generateMessageSeqNo(true),
		Bytes: int32(len(b)),
		Body:  b,
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

func (c *session) sendPushToQueue(gatewayId string, pushMsgId int64, pushMsg mtproto.TLObject) {
	rawBytes := pushMsg.Encode(c.cb.getLayer())
	if len(rawBytes) > 256 {
		gzipPacked := &mtproto.TLGzipPacked{
			PackedData: rawBytes,
		}
		rawBytes = gzipPacked.Encode(c.cb.getLayer())
	}

	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(false),
		Seqno: c.generateMessageSeqNo(true),
		Bytes: int32(len(rawBytes)),
		Body:  rawBytes,
	}
	c.outQueue.AddPushUpdates(pushMsgId, rawMsg)
}

func (c *session) sendRawToQueue(gatewayId string, msgId int64, confirm bool, rawMsg mtproto.TLObject) {
	b := rawMsg.Encode(c.cb.getLayer())
	rawMsg2 := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(false),
		Seqno: c.generateMessageSeqNo(confirm),
		Bytes: int32(len(b)),
		Body:  b,
	}
	c.outQueue.AddNotifyMsg(msgId, confirm, rawMsg2)
	// cb(rawMsg2)
}

func (c *session) sendHttpDirectToGateway(ch chan interface{}, confirm bool, obj mtproto.TLObject, cb func(sentRaw *mtproto.TLMessageRawData)) (bool, error) {
	salt := c.cb.getCacheSalt().GetSalt()
	b := obj.Encode(c.cb.getLayer())

	rawMsg := &mtproto.TLMessageRawData{
		MsgId: nextMessageId(false),
		Seqno: c.generateMessageSeqNo(confirm),
		Bytes: int32(len(b)),
		Body:  b,
	}

	rB, err := c.SendHttpDataToGateway(
		context.Background(),
		ch,
		c.cb.getTempAuthKeyId(),
		salt,
		c.sessionId,
		rawMsg)

	if err != nil {
		logx.Errorf("sendHttpDirectToGateway - %v", err)
	}

	cb(rawMsg)
	// c.httpTimeOut = 0
	return rB, err
}

func (c *session) sendDirectToGateway(gatewayId string, confirm bool, obj mtproto.TLObject, cb func(sentRaw *mtproto.TLMessageRawData)) (bool, error) {
	salt := c.cb.getCacheSalt().GetSalt()
	b := obj.Encode(c.cb.getLayer())

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

	if !c.isHttp {
		rB, err = c.SendDataToGateway(
			context.Background(),
			gatewayId,
			c.cb.getTempAuthKeyId(),
			salt,
			c.sessionId,
			rawMsg)
	} else {
		if ch := c.httpQueue.Pop(); ch != nil {
			rB, err = c.SendHttpDataToGateway(
				context.Background(),
				ch,
				c.cb.getTempAuthKeyId(),
				salt,
				c.sessionId,
				rawMsg)
		}
	}

	if err != nil {
		logx.Errorf("sendToClient - %v", err)
	}

	cb(rawMsg)
	// c.httpTimeOut = 0
	return rB, err
}

func (c *session) sendRawDirectToGateway(gatewayId string, raw *mtproto.TLMessageRawData) (bool, error) {
	salt := c.cb.getCacheSalt().GetSalt()

	var (
		rB  bool
		err error
	)
	if !c.isHttp {
		rB, err = c.SendDataToGateway(
			context.Background(),
			gatewayId,
			c.cb.getTempAuthKeyId(),
			salt,
			c.sessionId,
			raw)
	} else {
		if ch := c.httpQueue.Pop(); ch != nil {
			rB, err = c.SendHttpDataToGateway(
				context.Background(),
				ch,
				c.cb.getTempAuthKeyId(),
				salt,
				c.sessionId,
				raw)
		}
	}

	if err != nil {
		logx.Errorf("sendRawDirectToGateway - %v", err)
	}
	return rB, err
}

func (c *session) sendQueueToGateway(gatewayId string) {
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
	)
	for e := c.outQueue.oMsgs.Front(); e != nil; e = e.Next() {
		if e.Value.(*outboxMsg).sent == 0 || time.Now().Unix() >= e.Value.(*outboxMsg).sent+waitMsgAcksTimeout {
			pendings = append(pendings, e.Value.(*outboxMsg))
		}
	}

	if len(pendings) == 1 {
		logx.Infof("sendRawDirectToGateway - pendings[0]")
		b, err = c.sendRawDirectToGateway(gatewayId, pendings[0].msg)

	} else if len(pendings) > 1 {
		msgContainer := &mtproto.TLMsgRawDataContainer{
			Messages: make([]*mtproto.TLMessageRawData, 0, len(pendings)),
		}
		for _, m := range pendings {
			msgContainer.Messages = append(msgContainer.Messages, m.msg)
		}

		logx.Infof("sendRawDirectToGateway - TLMsgRawDataContainer")
		b, err = c.sendDirectToGateway(gatewayId, false, msgContainer, func(sentRaw *mtproto.TLMessageRawData) {
			// TODO(@benqi):
		})

		// TODO(@benqi): setOffline
	}

	// log.Debugf("err: %v, b: %v", err, b)
	if err == nil && b {
		for _, m := range pendings {
			logx.Infof("need_no_ack: %d", m.msgId)
			if m.state == NEED_NO_ACK {
				c.outQueue.Remove(m.msgId)
			} else {
				m.sent = time.Now().Unix()
			}
		}
	}
}
