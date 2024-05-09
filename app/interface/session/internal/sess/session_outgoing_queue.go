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
	"math"

	"github.com/teamgram/proto/mtproto"
)

const (
	maxAckedIdListSize = 500
)

type outboxMsg struct {
	msgId int64 // 1. req_msg_id; 2. 0; 3. < math.MaxInt32
	sent  int64
	state byte
	msg   *mtproto.TLMessageRawData
}

type sessionOutgoingQueue struct {
	minMsgId    int64
	maxMsgId    int64
	oMsgs       *list.List
	ackedIdList []int64
}

func newSessionOutgoingQueue() *sessionOutgoingQueue {
	return &sessionOutgoingQueue{
		minMsgId:    0,
		maxMsgId:    math.MaxInt64,
		oMsgs:       list.New(),
		ackedIdList: make([]int64, 0),
	}
}

func (q *sessionOutgoingQueue) AddRpcResultMsg(reqMsgId int64, result *mtproto.TLMessageRawData) *outboxMsg {
	oMsg := q.Lookup(reqMsgId)
	if oMsg == nil {
		oMsg = &outboxMsg{
			msgId: reqMsgId,
			sent:  0,
			state: ACKNOWLEDGED,
			msg:   result,
		}
		q.oMsgs.PushBack(oMsg)
	}

	q.Shrink()
	return oMsg
}

func (q *sessionOutgoingQueue) AddNotifyMsg(notifyId int64, confirm bool, msg *mtproto.TLMessageRawData) *outboxMsg {
	oMsg := new(outboxMsg)
	oMsg.msgId = notifyId
	oMsg.sent = 0
	if confirm {
		oMsg.state = ACKNOWLEDGED
	} else {
		oMsg.state = NEED_NO_ACK
	}
	oMsg.msg = msg
	q.oMsgs.PushBack(oMsg)

	q.Shrink()
	return oMsg
}

func (q *sessionOutgoingQueue) AddPushUpdates(pushMsgId int64, result *mtproto.TLMessageRawData) *outboxMsg {
	oMsg := q.Lookup(pushMsgId)
	if oMsg == nil {
		oMsg = &outboxMsg{
			msgId: pushMsgId,
			sent:  0,
			state: ACKNOWLEDGED,
			msg:   result,
		}
		q.oMsgs.PushBack(oMsg)
	}

	q.Shrink()
	return oMsg
}

func (q *sessionOutgoingQueue) OnMsgsAck(ackIds []int64, cb func(inMsgId int64)) {
	var next *list.Element
	for _, id := range ackIds {
		for e := q.oMsgs.Front(); e != nil; e = next {
			next = e.Next()
			if id == e.Value.(*outboxMsg).msg.MsgId {
				iMsgId := e.Value.(*outboxMsg).msgId
				q.ackedIdList = append(q.ackedIdList, iMsgId)
				cb(iMsgId)
				q.oMsgs.Remove(e)
			}
		}
	}

	if len(q.ackedIdList) > maxAckedIdListSize {
		q.ackedIdList = q.ackedIdList[len(q.ackedIdList)-maxAckedIdListSize-1:]
	}
}

func (q *sessionOutgoingQueue) Lookup(msgId int64) (oMsg *outboxMsg) {
	for e := q.oMsgs.Front(); e != nil; e = e.Next() {
		if msgId == e.Value.(*outboxMsg).msgId {
			oMsg = e.Value.(*outboxMsg)
		}
	}
	return
}

func (q *sessionOutgoingQueue) Remove(msgId int64) (oMsg *outboxMsg) {
	for e := q.oMsgs.Front(); e != nil; e = e.Next() {
		if msgId == e.Value.(*outboxMsg).msgId {
			oMsg = e.Value.(*outboxMsg)
			q.oMsgs.Remove(e)
			break
		}
	}
	return
}

func (q *sessionOutgoingQueue) Shrink() {
	for q.oMsgs.Len() > maxQueueSize {
		iMsg := q.oMsgs.Remove(q.oMsgs.Front())
		q.minMsgId = iMsg.(*outboxMsg).msgId
	}
}
