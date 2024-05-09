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

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

type pushMsg struct {
	msgId int64
	msg   mtproto.TLObject
}

type sessionPushQueue struct {
	q *list.List
}

func newSessionPushQueue() *sessionPushQueue {
	return &sessionPushQueue{
		q: list.New(),
	}
}

func (q *sessionPushQueue) Add(msgId int64, pushMsg2 mtproto.TLObject) {
	logx.Infof("add msgId: %d", msgId)
	q.q.PushBack(&pushMsg{
		msgId: msgId,
		msg:   pushMsg2,
	})
}

func (q *sessionPushQueue) Remove(msgId int64) {
	logx.Infof("remove msgId: %d", msgId)
	for e := q.q.Front(); e != nil; e = e.Next() {
		if msgId == e.Value.(*pushMsg).msgId {
			q.q.Remove(e)
			break
		}
	}
}
