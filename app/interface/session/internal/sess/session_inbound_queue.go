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
	"container/list"
	"fmt"
	"math"
)

/*
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
/*
 1    = nothing is known about the message (msg_id too low, the other party may have forgotten it)
 2    = message not received (msg_id falls within the range of stored identifiers;
	    however, the other party has certainly not received a message like that)
 3    = message not received (msg_id too high; however, the other party has certainly not received it yet)
 4    = message received (note that this response is also at the same time a receipt acknowledgment)
 +8   = message already acknowledged
 +16  = message not requiring acknowledgment
 +32  = RPC query contained in message being processed or processing already complete
 +64  = content-related response to message already generated
 +128 = other party knows for a fact that message is already received
*/
const (
	NONE                  byte = 0
	UNKNOWN               byte = 1
	NOT_RECEIVED          byte = 2
	NOT_RECEIVED_SURE     byte = 3
	RECEIVED              byte = 4
	ACKNOWLEDGED          byte = 8
	NEED_NO_ACK           byte = 16
	RPC_PROCESSING        byte = 32
	RESPONSE_GENERATED    byte = 64
	RESPONSE_ACKNOWLEDGED byte = 128
)

/*
const (
	inStateNone             = 0 // created
	inStateReceived         = 1 // received from client
	inStateRunning          = 2 // invoke api
	inStateWaitReplyTimeout = 3 // invoke timeout
	inStateInvoked          = 4 // invoke ok, send to client
	inStatePushSync         = 5 // invoke ok, send to client
	inStateAck              = 6 // received client ack
	inStateWaitAckTimeout   = 7 // wait ack timeout
	inStateError            = 8 // invalid error
	inStateEnd              = 9 // end state
	inStateIgnore           = 10
) */

/*
enum {
	MTPIdsBufferSize = 400, // received msgIds and wereAcked msgIds count stored
	MTPCheckResendTimeout = 10000, // how much time passed from send till we resend request or check it's state, in ms
	MTPCheckResendWaiting = 1000, // how much time to wait for some more requests, when resending request or checking it's state, in ms
	MTPAckSendWaiting = 10000, // how much time to wait for some more requests, when sending msg acks
	MTPResendThreshold = 1, // how much ints should message contain for us not to resend, but to check it's state
	MTPContainerLives = 600, // container lives 10 minutes in haveSent map

	MTPKillFileSessionTimeout = 5000, // how much time without upload / download causes additional session kill

	MaxUsersPerInvite = 100, // max users in one super group invite request

	MTPChannelGetDifferenceLimit = 100,

	MaxSelectedItems = 100,

	MaxPhoneCodeLength = 4, // max length of country phone code
	MaxPhoneTailLength = 32, // rest of the phone number, without country code (seen 12 at least), need more for service numbers

	MaxScrollSpeed = 37, // 37px per 15ms while select-by-drag
	FingerAccuracyThreshold = 3, // touch flick ignore 3px
	MaxScrollAccelerated = 4000, // 4000px per second
	MaxScrollFlick = 2500, // 2500px per second

	LocalEncryptIterCount = 4000, // key derivation iteration count
	LocalEncryptNoPwdIterCount = 4, // key derivation iteration count without pwd (not secure anyway)
	LocalEncryptSaltSize = 32, // 256 bit

	AnimationTimerDelta = 7,
	ClipThreadsCount = 8,
	AverageGifSize = 320 * 240,
	WaitBeforeGifPause = 200, // wait 200ms for gif draw before pausing it
	RecentInlineBotsLimit = 10,

	AVBlockSize = 4096, // 4Kb for ffmpeg blocksize

	AutoSearchTimeout = 900, // 0.9 secs
	SearchPerPage = 50,
	SearchManyPerPage = 100,
	LinksOverviewPerPage = 12,
	MediaOverviewStartPerPage = 5,

	AudioVoiceMsgMaxLength = 100 * 60, // 100 minutes
	AudioVoiceMsgUpdateView = 100, // 100ms
	AudioVoiceMsgChannels = 2, // stereo
	AudioVoiceMsgBufferSize = 256 * 1024, // 256 Kb buffers (1.3 - 3.0 secs)

	StickerMaxSize = 2048, // 2048x2048 is a max image size for sticker

	MaxZoomLevel = 7, // x8
	ZoomToScreenLevel = 1024, // just constant

	ShortcutsCountLimit = 256, // how many shortcuts can be in json file

	PreloadHeightsCount = 3, // when 3 screens to scroll left make a preload request

	SearchPeopleLimit = 5,
	UsernameCheckTimeout = 200,

	MaxMessageSize = 4096,

	WriteMapTimeout = 1000,

	SetOnlineAfterActivity = 30, // user with hidden last seen stays online for such amount of seconds in the interface

	ServiceUserId = 777000,
	WebPageUserId = 701000,

	CacheBackgroundTimeout = 3000, // cache background scaled image after 3s
	BackgroundsInRow = 3,

	UpdateDelayConstPart = 8 * 3600, // 8 hour min time between update check requests
	UpdateDelayRandPart = 8 * 3600, // 8 hour max - min time between update check requests

	WrongPasscodeTimeout = 1500,
	SessionsShortPollTimeout = 60000,

	ChoosePeerByDragTimeout = 1000, // 1 second mouse not moved to choose dialog when dragging a file
};
*/

const (
	maxQueueSize = 400
)

type inboxMsg struct {
	msgId int64
	seqNo int32
	state byte
}

func (m *inboxMsg) ChangeState(s byte) {
	m.state = s
}

func (m *inboxMsg) DebugString() string {
	return fmt.Sprintf("{msg_id:%d, seqno: %d, state: %d}", m.msgId, m.seqNo, m.state)
}

func newInboxMsg(msgId int64) *inboxMsg {
	r := new(inboxMsg)
	r.msgId = msgId
	r.state = NONE
	return r
}

type sessionInboundQueue struct {
	firstMsgId int64
	minMsgId   int64
	maxMsgId   int64
	msgIds     *list.List
}

func newSessionInboundQueue() *sessionInboundQueue {
	q := new(sessionInboundQueue)
	q.msgIds = list.New()
	q.firstMsgId = 0
	q.minMsgId = 0
	q.maxMsgId = math.MaxInt64
	return q
}

func (q *sessionInboundQueue) AddMsgId(msgId int64) (r *inboxMsg) {
	// TODO(@benqi): resize 100

	if msgId < q.minMsgId {
		q.minMsgId = msgId
	}
	if msgId > q.maxMsgId {
		q.maxMsgId = msgId
	}

	for e := q.msgIds.Front(); e != nil; e = e.Next() {
		if e.Value.(*inboxMsg).msgId > msgId {
			r = newInboxMsg(msgId)
			q.msgIds.InsertBefore(r, e)
			return
		} else if e.Value.(*inboxMsg).msgId == msgId {
			r = e.Value.(*inboxMsg)
			return
		}
	}
	r = newInboxMsg(msgId)
	q.msgIds.PushBack(r)

	return
}

func (q *sessionInboundQueue) GetMinMsgId() int64 {
	return q.minMsgId
}

func (q *sessionInboundQueue) GetMaxMsgId() int64 {
	return q.maxMsgId
}

func (q *sessionInboundQueue) ChangeAckReceived(msgId int64) {
	for e := q.msgIds.Front(); e != nil; e = e.Next() {
		if e.Value.(*inboxMsg).msgId == msgId {
			e.Value.(*inboxMsg).state = RECEIVED | RESPONSE_ACKNOWLEDGED
		}
	}
}

func (q *sessionInboundQueue) Lookup(msgId int64) (iMsg *inboxMsg) {
	for e := q.msgIds.Front(); e != nil; e = e.Next() {
		if msgId == e.Value.(*inboxMsg).msgId {
			iMsg = e.Value.(*inboxMsg)
			return
		}
	}
	return
}

func (q *sessionInboundQueue) Shrink() {
	for q.msgIds.Len() > maxQueueSize {
		iMsg := q.msgIds.Remove(q.msgIds.Front())
		q.minMsgId = iMsg.(*inboxMsg).msgId
	}
}

func (q *sessionInboundQueue) FindLowerEntry(msgId int64) (iMsg *inboxMsg) {
	for e := q.msgIds.Back(); e != nil; e = e.Prev() {
		if msgId >= e.Value.(*inboxMsg).msgId {
			return e.Value.(*inboxMsg)
		}
	}
	return nil
}

func (q *sessionInboundQueue) FindHigherEntry(msgId int64) (iMsg *inboxMsg) {
	for e := q.msgIds.Front(); e != nil; e = e.Next() {
		if e.Value.(*inboxMsg).msgId >= msgId {
			return e.Value.(*inboxMsg)
		}
	}
	return nil
}

func (q *sessionInboundQueue) Length() int {
	return q.msgIds.Len()
}
