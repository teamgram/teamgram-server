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
	"math"
	"reflect"
	"time"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

func (c *session) onMsgsAck(ctx context.Context, gatewayId string, msgId int64, seqno int32, request *mtproto.TLMsgsAck) {
	logx.WithContext(ctx).Infof("onMsgsAck - request data: {sess: %s, gatewayId: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId,
		seqno,
		request)

	c.outQueue.OnMsgsAck(request.GetMsgIds(), func(inMsgId int64) {
		// Notes: ignore notifyId
		if inMsgId > math.MaxInt32 {
			// 1. rpc
			c.inQueue.ChangeAckReceived(inMsgId)
		} else {
			// 3. push
			// TODO(@benqi): client received updates, will remove it from cache
		}
	})
}

//// Check Server Salt
/*
 Server Salt
	A (random) 64-bit number periodically (say, every 24 hours) changed (separately for each session)
	at the request of the server.
	All subsequent messages must contain the new salt (although,
	messages with the old salt are still accepted for a further 300 seconds).
	Required to protect against replay attacks and certain tricks associated
	with adjusting the client clock to a moment in the distant future.
*/
func (c *session) checkBadServerSalt(ctx context.Context, gatewayId string, salt int64, msg *mtproto.TLMessage2) bool {
	// Notice of Ignored Error Message
	//
	// Here, error_code can also take on the following values:
	//  48: incorrect server salt (in this case,
	//      the bad_server_salt response is received with the correct salt,
	//      and the message is to be re-sent with it)
	//

	valid := false

	if salt == c.sessList.cacheSalt.GetSalt() {
		valid = true
	} else {
		if c.sessList.cacheSalt != nil {
			if salt == c.sessList.cacheSalt.GetSalt() {
				date := int32(time.Now().Unix())
				if c.sessList.cacheSalt.GetValidUntil()+300 >= date {
					valid = true
				}
			}
		}
	}

	if !valid {
		badServerSalt := mtproto.MakeTLBadServerSalt(&mtproto.BadMsgNotification{
			BadMsgId:      msg.MsgId,
			ErrorCode:     kServerSaltIncorrect,
			BadMsgSeqno:   msg.Seqno,
			NewServerSalt: c.sessList.cacheSalt.GetSalt(),
		}).To_BadMsgNotification()
		logx.WithContext(ctx).Errorf("invalid salt: %d, send badServerSalt: {%v}, cacheSalt: %v", salt, badServerSalt, c.sessList.cacheLastSalt)

		c.sendDirectToGateway(ctx, gatewayId, false, badServerSalt, func(sentRaw *mtproto.TLMessageRawData) {
			// nothing do
		})
		return false
	}

	return valid
}

/**********************************************************************************************************************
** android client source code:
**

  } else if (typeInfo == typeid(TL_bad_msg_notification)) {
      TL_bad_msg_notification *result = (TL_bad_msg_notification *) message;
      if (LOGS_ENABLED) DEBUG_E("bad message notification %d for messageId 0x%" PRIx64 ", seqno %d", result->error_code, result->bad_msg_id, result->bad_msg_seqno);
      switch (result->error_code) {
          case 16:
          case 17:
          case 19:
          case 32:
          case 33:
          case 64: {
              int64_t realId = messageId != 0 ? messageId : containerMessageId;
              if (realId == 0) {
                  realId = innerMsgId;
              }

              if (realId != 0) {
                  int64_t time = (int64_t) (messageId / 4294967296.0 * 1000);
                  int64_t currentTime = getCurrentTimeMillis();
                  timeDifference = (int32_t) ((time - currentTime) / 1000 - currentPingTime / 2);
              }

              datacenter->recreateSessions(HandshakeTypeAll);
              saveConfig();

              lastOutgoingMessageId = 0;
              clearRequestsForDatacenter(datacenter, HandshakeTypeAll);
              break;
          }
          case 20: {
              for (requestsIter iter = runningRequests.begin(); iter != runningRequests.end(); iter++) {
                  Request *request = iter->get();
                  if (request->respondsToMessageId(result->bad_msg_id)) {
                      if (request->completed) {
                          break;
                      }
                      connection->addMessageToConfirm(result->bad_msg_id);
                      request->clear(true);
                      break;
                  }
              }
          }
          default:
              break;
      }
*/
/*
  fun checkSeqNo(message: TL_message): MsgNotificationCode {
    if (queue.isEmpty()) {
      return MsgNotificationCode.OK
    }
    val preEntry = queue.lowerEntry(message.msg_id)
    if (preEntry != null) {
      if (preEntry.value.seq_no > message.seqno) {
        return MsgNotificationCode.MSG_SEQNO_TOO_LOW
      }
    }
    val nextEntry = queue.higherEntry(message.msg_id)
    if (nextEntry != null) {
      if (nextEntry.value.seq_no < message.seqno) {
        return MsgNotificationCode.MSG_SEQNO_TOO_HIGH
      }
    }
    return MsgNotificationCode.OK
  }
*/

// func checkConfirm()
func (c *session) checkBadMsgNotification(ctx context.Context, gatewayId string, excludeMsgIdToo bool, msg *mtproto.TLMessage2) bool {
	// Notice of Ignored Error Message
	//
	// In certain cases, a server may notify a client that its incoming message was ignored for whatever reason.
	// Note that such a notification cannot be generated unless a message is correctly decoded by the server.
	//
	// bad_msg_notification#a7eff811 bad_msg_id:long bad_msg_seqno:int error_code:int = BadMsgNotification;
	// bad_server_salt#edab447b bad_msg_id:long bad_msg_seqno:int error_code:int new_server_salt:long = BadMsgNotification;
	//
	// Here, error_code can also take on the following values:
	//
	//  16: msg_id too low (most likely, client time is wrong;
	//      it would be worthwhile to synchronize it using msg_id notifications
	//      and re-send the original message with the “correct” msg_id or wrap
	//      it in a container with a new msg_id
	//      if the original message had waited too long on the client to be transmitted)
	//  17: msg_id too high (similar to the previous case,
	//      the client time has to be synchronized, and the message re-sent with the correct msg_id)
	//  18: incorrect two lower order msg_id bits (the server expects client message msg_id to be divisible by 4)
	//  19: container msg_id is the same as msg_id of a previously received message (this must never happen)
	//  20: message too old, and it cannot be verified whether the server has received a message with this msg_id or not
	//  32: msg_seqno too low (the server has already received a message with a lower msg_id
	//      but with either a higher or an equal and odd seqno)
	//  33: msg_seqno too high (similarly, there is a message with a higher msg_id
	//      but with either a lower or an equal and odd seqno)
	//  34: an even msg_seqno expected (irrelevant message), but odd received
	//  35: odd msg_seqno expected (relevant message), but even received
	//  48: incorrect server salt (in this case,
	//      the bad_server_salt response is received with the correct salt,
	//      and the message is to be re-sent with it)
	//  64: invalid container.
	//
	// The intention is that error_code values are grouped (error_code >> 4):
	// for example, the codes 0x40 - 0x4f correspond to errors in container decomposition.
	//
	// Notifications of an ignored message do not require acknowledgment (i.e., are irrelevant).
	//
	// Important: if server_salt has changed on the server or if client time is incorrect,
	// any query will result in a notification in the above format.
	// The client must check that it has, in fact,
	// recently sent a message with the specified msg_id, and if that is the case,
	// update its time correction value (the difference between the client’s and the server’s clocks)
	// and the server salt based on msg_id and the server_salt notification,
	// so as to use these to (re)send future messages. In the meantime,
	// the original message (the one that caused the error message to be returned)
	// must also be re-sent with a better msg_id and/or server_salt.
	//
	// In addition, the client can update the server_salt value used to send messages to the server,
	// based on the values of RPC responses or containers carrying an RPC response,
	// provided that this RPC response is actually a match for the query sent recently.
	// (If there is doubt, it is best not to update since there is risk of a replay attack).
	//

	//=============================================================================================
	// TODO(@benqi): Time Synchronization, https://core.telegram.org/mtproto#time-synchronization
	//
	// Time Synchronization
	//
	// If client time diverges widely from server time,
	// a server may start ignoring client messages,
	// or vice versa, because of an invalid message identifier (which is closely related to creation time).
	// Under these circumstances,
	// the server will send the client a special message containing the correct time and
	// a certain 128-bit salt (either explicitly provided by the client in a special RPC synchronization request or
	// equal to the key of the latest message received from the client during the current session).
	// This message could be the first one in a container that includes other messages
	// (if the time discrepancy is significant but does not as yet result in the client’s messages being ignored).
	//
	// Having received such a message or a container holding it,
	// the client first performs a time synchronization (in effect,
	// simply storing the difference between the server’s time
	// and its own to be able to compute the “correct” time in the future)
	// and then verifies that the message identifiers for correctness.
	//
	// Where a correction has been neglected,
	// the client will have to generate a new session to assure the monotonicity of message identifiers.
	//

	var errorCode int32 = 0

	serverTime := time.Now().Unix()
	for {
		clientTime := int64(msg.MsgId / 4294967296.0)
		if !excludeMsgIdToo {
			/********************************************************************
			// tdesktop source code

				int32 serverTime((int32)(msgId >> 32)), clientTime(unixtime());
				bool isReply = ((msgId & 0x03) == 1);
				if (!isReply && ((msgId & 0x03) != 3)) {
					LOG(("MTP Error: bad msg_id %1 in message received").arg(msgId));

					return restartOnError();
				}

				bool badTime = false;
				uint64 mySalt = sessionData->getSalt();
				if (serverTime > clientTime + 60 || serverTime + 300 < clientTime) {
					DEBUG_LOG(("MTP Info: bad server time from msg_id: %1, my time: %2").arg(serverTime).arg(clientTime));
					badTime = true;
				}
			*/
			if clientTime+60 < serverTime {
				errorCode = kMsgIdTooLow
				logx.WithContext(ctx).Errorf("bad server time - {msg_id: %d, clientTime: %d, serverTime: %d}", msg.MsgId, clientTime, serverTime)

				break
			}
			if clientTime > serverTime+300 {
				errorCode = kMsgIdTooHigh
				logx.WithContext(ctx).Errorf("bad server time - {msg_id: %d, clientTime: %d, serverTime: %d}", msg.MsgId, clientTime, serverTime)
				break
			}
		}

		// TODO(@benqi): 检查低32位

		//=================================================================================================
		// Check Message Identifier (msg_id)
		//
		// https://core.telegram.org/mtproto/description#message-identifier-msg-id
		// Message Identifier (msg_id)
		//
		// A (time-dependent) 64-bit number used uniquely to identify a message within a session.
		// Client message identifiers are divisible by 4,
		// server message identifiers modulo 4 yield 1 if the message is a response to a client message, and 3 otherwise.
		// Client message identifiers must increase monotonically (within a single session),
		// the same as server message identifiers, and must approximately equal unixtime*2^32.
		// This way, a message identifier points to the approximate moment in time the message was created.
		// A message is rejected over 300 seconds after it is created or 30 seconds
		// before it is created (this is needed to protect from replay attacks).
		// In this situation,
		// it must be re-sent with a different identifier (or placed in a container with a higher identifier).
		// The identifier of a message container must be strictly greater than those of its nested messages.
		//
		// Important: to counter replay-attacks the lower 32 bits of msg_id passed
		// by the client must not be empty and must present a fractional
		// part of the time point when the message was created.
		//
		if msg.MsgId%4 != 0 {
			errorCode = kMsgIdMod4
			break
		}

		if msg.MsgId < c.inQueue.GetMinMsgId() {
			errorCode = kMsgIdTooOld
			break
		}

		//// TODO(@benqi): check kSeqNoTooHigh and kSeqNoTooLow
		//lMsg := c.inQueue.FindLowerEntry(msg.MsgId)
		//if lMsg != nil && lMsg.seqNo > msg.Seqno {
		//	errorCode = kSeqNoTooLow
		//	break
		//}
		//hMsg := c.inQueue.FindHigherEntry(msg.MsgId)
		//if hMsg != nil && hMsg.seqNo < msg.Seqno {
		//	errorCode = kSeqNoTooHigh
		//	break
		//}

		// check container
		if msgContainer, ok := msg.Object.(*mtproto.TLMsgContainer); ok {
			errorCode = c.checkContainer(ctx, msg.MsgId, msg.Seqno, msgContainer)
			if errorCode != 0 {
				break
			}
		}

		if checkMessageConfirm(msg.Object) {
			if msg.Seqno%2 == 0 {
				// errorCode = kSeqNoNotOdd
				// break
				// log.Errorf("kSeqNoNotOdd: %s", reflect.TypeOf(msg.Object))
			}
		} else {
			if msg.Seqno%2 != 0 {
				// errorCode = kSeqNoNotEven
				// break
				// log.Errorf("kSeqNoNotEven %s", reflect.TypeOf(msg.Object))
			}
		}

		// end
		break
	}

	if errorCode != 0 {
		badMsgNotification := mtproto.MakeTLBadMsgNotification(&mtproto.BadMsgNotification{
			BadMsgId:    msg.MsgId,
			BadMsgSeqno: msg.Seqno,
			ErrorCode:   errorCode,
		}).To_BadMsgNotification()
		logx.WithContext(ctx).Error("errorCode - ", errorCode, ", msg: ", reflect.TypeOf(msg.Object))
		c.sendDirectToGateway(ctx, gatewayId, false, badMsgNotification, func(sentRaw *mtproto.TLMessageRawData) {
			// nothing do
		})
		return false
	}
	return true
}

/*
	 // tdesktop ------------------------------------------------------------------------------------------

		case mtpc_msgs_state_req: {
			if (badTime) {
				DEBUG_LOG(("Message Info: skipping with bad time..."));
				return HandleResult::Ignored;
			}
			MTPMsgsStateReq msg;
			msg.read(from, end);
			auto &ids = msg.c_msgs_state_req().vmsg_ids.v;
			auto idsCount = ids.size();
			DEBUG_LOG(("Message Info: msgs_state_req received, ids: %1").arg(LogIdsVector(ids)));
			if (!idsCount) return HandleResult::Success;

			QByteArray info(idsCount, Qt::Uninitialized);
			{
				QReadLocker lock(sessionData->receivedIdsMutex());
				auto &receivedIds = sessionData->receivedIdsSet();
				auto minRecv = receivedIds.min();
				auto maxRecv = receivedIds.max();

				QReadLocker locker(sessionData->wereAckedMutex());
				const auto &wereAcked = sessionData->wereAckedMap();
				const auto wereAckedEnd = wereAcked.cend();

				for (uint32 i = 0, l = idsCount; i < l; ++i) {
					char state = 0;
					uint64 reqMsgId = ids[i].v;
					if (reqMsgId < minRecv) {
						state |= 0x01;
					} else if (reqMsgId > maxRecv) {
						state |= 0x03;
					} else {
						auto msgIdState = receivedIds.lookup(reqMsgId);
						if (msgIdState == ReceivedMsgIds::State::NotFound) {
							state |= 0x02;
						} else {
							state |= 0x04;
							if (wereAcked.constFind(reqMsgId) != wereAckedEnd) {
								state |= 0x80; // we know, that server knows, that we received request
							}
							if (msgIdState == ReceivedMsgIds::State::NeedsAck) { // need ack, so we sent ack
								state |= 0x08;
							} else {
								state |= 0x10;
							}
						}
					}
					info[i] = state;
				}
			}
			emit sendMsgsStateInfoAsync(msgId, info);
		} return HandleResult::Success;
*/
func (c *session) onMsgsStateReq(ctx context.Context, gatewayId string, msgId *inboxMsg, request *mtproto.TLMsgsStateReq) {
	logx.WithContext(ctx).Debugf("onMsgsStateReq - request data: {sess: %s, gatewayId: %s, md: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	// Request for Message Status Information
	//
	// If either party has not received information on the status of its outgoing messages for a while,
	// it may explicitly request it from the other party:
	//
	// msgs_state_req#da69fb52 msg_ids:Vector long = MsgsStateReq;
	// The response to the query contains the following information:
	//
	// Informational Message regarding Status of Messages
	// msgs_state_info#04deb57d req_msg_id:long info:string = MsgsStateInfo;
	//
	// Here, info is a string that contains exactly one byte of message status for
	// each message from the incoming msg_ids list:
	//
	// 1 = nothing is known about the message (msg_id too low, the other party may have forgotten it)
	// 2 = message not received (msg_id falls within the range of stored identifiers; however,
	// 	   the other party has certainly not received a message like that)
	// 3 = message not received (msg_id too high; however, the other party has certainly not received it yet)
	// 4 = message received (note that this response is also at the same time a receipt acknowledgment)
	// +8 = message already acknowledged
	// +16 = message not requiring acknowledgment
	// +32 = RPC query contained in message being processed or processing already complete
	// +64 = content-related response to message already generated
	// +128 = other party knows for a fact that message is already received
	//
	// This response does not require an acknowledgment.
	// It is an acknowledgment of the relevant msgs_state_req, in and of itself.
	//
	// Note that if it turns out suddenly that the other party does not have a message
	// that looks like it has been sent to it, the message can simply be re-sent.
	// Even if the other party should receive two copies of the message at the same time,
	// the duplicate will be ignored. (If too much time has passed,
	// and the original msg_id is not longer valid, the message is to be wrapped in msg_copy).
	//

	msgIds := request.GetMsgIds()
	info := make([]byte, len(msgIds))
	for i := 0; i < len(msgIds); i++ {
		if msgIds[i] < c.inQueue.GetMinMsgId() {
			info[i] = UNKNOWN
			continue
		}
		if msgIds[i] > c.inQueue.GetMaxMsgId() {
			info[i] = NOT_RECEIVED_SURE
			continue
		}

		iMsgId := c.inQueue.Lookup(msgIds[i])
		if iMsgId == nil {
			info[i] = NOT_RECEIVED
			continue
		}

		info[i] = iMsgId.state
	}

	msgsStateInfo := mtproto.MakeTLMsgsStateInfo(&mtproto.MsgsStateInfo{
		ReqMsgId: msgId.msgId,
		Info:     string(info),
	}).To_MsgsStateInfo()

	c.sendRawToQueue(ctx, gatewayId, msgId.msgId, false, msgsStateInfo)
	msgId.state = RECEIVED | NEED_NO_ACK
}

/*
****************************************************************************************************************
// resend
// 注意：服务端要处理这种情况，和文档不太一样
//

	mtpRequestId Session::resend(quint64 msgId, qint64 msCanWait, bool forceContainer, bool sendMsgStateInfo) {
		SecureRequest request;
		{
			QWriteLocker locker(data.haveSentMutex());
			auto &haveSent = data.haveSentMap();

			auto i = haveSent.find(msgId);
			if (i == haveSent.end()) {
				if (sendMsgStateInfo) {
					char cantResend[2] = {1, 0};
					DEBUG_LOG(("Message Info: cant resend %1, request not found").arg(msgId));

					auto info = std::string(cantResend, cantResend + 1);
					return _instance->sendProtocolMessage(
						dcWithShift,
						MTPMsgsStateInfo(
							MTP_msgs_state_info(
								MTP_long(msgId),
								MTP_string(std::move(info)))));
				}
				return 0;
			}

			request = i.value();
			haveSent.erase(i);
		}
		if (request.isSentContainer()) { // for container just resend all messages we can
			DEBUG_LOG(("Message Info: resending container from haveSent, msgId %1").arg(msgId));
			const mtpMsgId *ids = (const mtpMsgId *)(request->constData() + 8);
			for (uint32 i = 0, l = (request->size() - 8) >> 1; i < l; ++i) {
				resend(ids[i], 10, true);
			}
			return 0xFFFFFFFF;
		} else if (!request.isStateRequest()) {
			request->msDate = forceContainer ? 0 : getms(true);
			sendPrepared(request, msCanWait, false);
			{
				QWriteLocker locker(data.toResendMutex());
				data.toResendMap().insert(msgId, request->requestId);
			}
			return request->requestId;
		} else {
			return 0;
		}
	}
*/
func (c *session) onMsgsStateInfo(ctx context.Context, gatewayId string, msgId *inboxMsg, request *mtproto.TLMsgsStateInfo) {
	logx.WithContext(ctx).Infof("onMsgsStateInfo - request data: {sess: %s, gatewayId: %s, md: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	// 1. handle other
	// 1) hand request.
	reqMsgId := request.GetReqMsgId()
	oMsg := c.outQueue.Lookup(reqMsgId)
	if oMsg == nil {
		logx.WithContext(ctx).Errorf("not found reqMsgId - %d", reqMsgId)
		return
	}

	var (
		msgIds []int64
		info   = []byte(request.GetInfo())
	)

	dBuf := mtproto.NewDecodeBuf(oMsg.msg.Body)
	r := dBuf.Object()
	switch r.(type) {
	case *mtproto.TLMsgsStateReq:
		msgIds = r.(*mtproto.TLMsgsStateReq).GetMsgIds()
	case *mtproto.TLMsgResendReq:
		msgIds = r.(*mtproto.TLMsgResendReq).GetMsgIds()
	default:
		// TODO(@benqi): process tdektop client: resend
		logx.WithContext(ctx).Errorf("not found reqMsgId - %d", reqMsgId)
		return
	}

	if len(msgIds) != len(info) {
		logx.WithContext(ctx).Errorf("invalid msgIds, len(msgIds) != len(info)")
		return
	}

	ackIds := make([]int64, 0, len(msgIds))
	ackIds = append(ackIds, reqMsgId)
	resendIds := make([]int64, 0, len(msgIds))
	for i := 0; i < len(msgIds); i++ {
		if info[i] == UNKNOWN || info[i] == NOT_RECEIVED_SURE {
			// drop
		} else if info[i] == NOT_RECEIVED {
			resendIds = append(resendIds, msgIds[i])
			// resend
		} else {
			// remove
			ackIds = append(ackIds, msgIds[i])
		}
	}
	c.outQueue.OnMsgsAck(ackIds, func(inMsgId int64) {
		// Notes: ignore notifyId
		if inMsgId <= math.MaxInt32 {
			// 3. push
			// TODO(@benqi): client received updates, will remove it from cache
		}
	})

	// TODO(@benqi): resend
	if len(resendIds) > 0 {
		//
	}

	// 2. no ack
	msgId.state = RECEIVED | NEED_NO_ACK
}

func (c *session) onMsgsAllInfo(ctx context.Context, gatewayId string, msgId *inboxMsg, request *mtproto.TLMsgsAllInfo) {
	logx.WithContext(ctx).Infof("onMsgsAllInfo - request data: {sess: %s, conn_id: %s, md: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	// Voluntary Communication of Status of Messages
	//
	// Either party may voluntarily inform the other party of the status of the messages transmitted by the other party.
	//
	// msgs_all_info#8cc0d131 msg_ids:Vector long info:string = MsgsAllInfo
	//
	// All message codes known to this party are enumerated,
	// with the exception of those for which the +128 and the +16 flags are set.
	// However, if the +32 flag is set but not +64, then the message status will still be communicated.
	//
	// This message does not require an acknowledgment.
	//

	var (
		msgIds = request.GetMsgIds()
		info   = []byte(request.GetInfo())
	)

	ackIds := make([]int64, 0, len(msgIds))
	resendIds := make([]int64, 0, len(msgIds))

	for i := 0; i < len(msgIds); i++ {
		if info[i] == UNKNOWN || info[i] == NOT_RECEIVED_SURE {
			// drop
		} else if info[i] == NOT_RECEIVED {
			resendIds = append(resendIds, msgIds[i])
			// resend
		} else {
			// remove
			ackIds = append(ackIds, msgIds[i])
		}
	}
	c.outQueue.OnMsgsAck(ackIds, func(inMsgId int64) {
		// Notes: ignore notifyId
		if inMsgId <= math.MaxInt32 {
			// 3. push
			// TODO(@benqi): client received updates, will remove it from cache
		}
	})

	// TODO(@benqi): resend
	if len(resendIds) > 0 {
		//
	}

	// 2. no ack
	msgId.state = RECEIVED | NEED_NO_ACK
}

func (c *session) onMsgResendReq(ctx context.Context, gatewayId string, msgId *inboxMsg, request *mtproto.TLMsgResendReq) {
	logx.WithContext(ctx).Errorf("onMsgResendReq - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	// Explicit Request to Re-Send Messages
	//
	// msg_resend_req#7d861a08 msg_ids:Vector long = MsgResendReq;
	//
	// The remote party immediately responds by re-sending the requested messages,
	// normally using the same connection that was used to transmit the query.
	// If at least one message with requested msg_id does not exist or has already been forgotten,
	// or has been sent by the requesting party (known from parity),
	// MsgsStateInfo is returned for all messages requested
	// as if the MsgResendReq query had been a MsgsStateReq query as well.
	//

	var (
		msgIds         = request.GetMsgIds()
		info           = make([]byte, len(msgIds))
		bMsgsStateInfo = false
	)

	for i := 0; i < len(msgIds); i++ {
		if msgIds[i] < c.inQueue.GetMinMsgId() {
			bMsgsStateInfo = true
			break
		}
		if msgIds[i] > c.inQueue.GetMaxMsgId() {
			bMsgsStateInfo = true
			break
		}

		iMsgId := c.inQueue.Lookup(msgIds[i])
		if iMsgId == nil {
			bMsgsStateInfo = true
			break
		}
	}

	if bMsgsStateInfo {
		for i := 0; i < len(msgIds); i++ {
			if msgIds[i] < c.inQueue.GetMinMsgId() {
				info[i] = UNKNOWN
				continue
			}
			if msgIds[i] > c.inQueue.GetMaxMsgId() {
				info[i] = NOT_RECEIVED_SURE
				continue
			}

			// TODO(@benqi):
			iMsgId := c.inQueue.Lookup(msgIds[i])
			if iMsgId == nil {
				info[i] = NOT_RECEIVED
				continue
			}

			info[i] = iMsgId.state
		}

		msgsStateInfo := mtproto.MakeTLMsgsStateInfo(&mtproto.MsgsStateInfo{
			ReqMsgId: msgId.msgId,
			Info:     string(info),
		}).To_MsgsStateInfo()

		c.sendRawToQueue(ctx, gatewayId, msgId.msgId, false, msgsStateInfo)
		msgId.state = RECEIVED | NEED_NO_ACK
	} else {
		for i := 0; i < len(msgIds); i++ {
			// resend
		}
	}
}

func (c *session) onMsgDetailInfo(ctx context.Context, gatewayId string, msgId *inboxMsg, request *mtproto.TLMsgDetailedInfo) {
	logx.WithContext(ctx).Errorf("onMsgDetailInfo - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	// NOTE(@benqi): not received by server
}

func (c *session) onMsgNewDetailInfo(ctx context.Context, gatewayId string, msgId *inboxMsg, request *mtproto.TLMsgDetailedInfo) {
	logx.WithContext(ctx).Errorf("onMsgNewDetailInfo - request data: {sess: %s, conn_id: %s, msg_id: %d, seq_no: %d, request: {%s}}",
		c,
		gatewayId,
		msgId.msgId,
		msgId.seqNo,
		request)

	// NOTE(@benqi): not received by server
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////
func (c *session) notifyMsgsStateInfo(ctx context.Context, gatewayId string, inMsg *inboxMsg) {
	// TODO(@benqi): if aced and < resendSize, send rsp.
	msgsStateInfo := mtproto.MakeTLMsgsStateInfo(&mtproto.MsgsStateInfo{
		ReqMsgId: inMsg.msgId,
		Info:     string([]byte{inMsg.state}),
	})
	c.sendDirectToGateway(ctx, gatewayId, false, msgsStateInfo, func(sentRaw *mtproto.TLMessageRawData) {
		// nothing do
	})
}

func (c *session) notifyMsgsStateReq() {
	// TODO(@benqi):
}

func (c *session) notifyMsgsAllInfo() {
	// TODO(@benqi):
}

// Extended Voluntary Communication of Status of One Message
//
// Normally used by the server to respond to the receipt of a duplicate msg_id,
// especially if a response to the message has already been generated and the response is large.
// If the response is small, the server may re-send the answer itself instead.
// This message can also be used as a notification instead of resending a large message.
//
// msg_detailed_info#276d3ec6 msg_id:long answer_msg_id:long bytes:int status:int = MsgDetailedInfo;
// msg_new_detailed_info#809db6df answer_msg_id:long bytes:int status:int = MsgDetailedInfo;
//
// The second version is used to notify of messages that were created on the server
// not in response to an RPC query (such as notifications of new messages)
// and were transmitted to the client some time ago, but not acknowledged.
//
// Currently, status is always zero. This may change in future.
//
// This message does not require an acknowledgment.
func (c *session) notifyMsgDetailedInfo(inMsg *inboxMsg) {
}

func (c *session) notifyNewMsgDetailedInfo() {
	// TODO(@benqi):

}

// FIXME(@benqi): 看起来像已经废弃了
// case *mtproto.TLMsgResendA
func (c *session) notifyMsgResendAnsSeq() {
	// TODO(@benqi): not impl

	// Explicit Request to Re-Send Answers
	//
	// msg_resend_ans_req#8610baeb msg_ids:Vector long = MsgResendReq;
	//
	// The remote party immediately responds by re-sending answers to the requested messages,
	// normally using the same connection that was used to transmit the query.
	// MsgsStateInfo is returned for all messages requested
	// as if the MsgResendReq query had been a MsgsStateReq query as well.
	//
}
