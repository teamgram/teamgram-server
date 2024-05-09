/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teagramio (teagram.io@gmail.com)
 */

package sess

import (
	"container/list"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type rpcResultWaiting struct {
	msgId int64
	date  int64
}

type sessionRpcResultWaitingQueue struct {
	q *list.List
}

func newSessionRpcResultWaitingQueue() *sessionRpcResultWaitingQueue {
	return &sessionRpcResultWaitingQueue{
		q: list.New(),
	}
}

func (q *sessionRpcResultWaitingQueue) Add(msgId int64) {
	logx.Infof("add msgId: %d", msgId)
	q.q.PushBack(&rpcResultWaiting{
		msgId: msgId,
		date:  time.Now().Unix() + 5,
	})
}

func (q *sessionRpcResultWaitingQueue) Remove(msgId int64) {
	logx.Infof("remove msgId: %d", msgId)
	for e := q.q.Front(); e != nil; e = e.Next() {
		if msgId == e.Value.(*rpcResultWaiting).msgId {
			q.q.Remove(e)
			break
		}
	}
}

func (q *sessionRpcResultWaitingQueue) OnTimer() (msgIdList []int64) {
	date := time.Now().Unix()
	for e := q.q.Front(); e != nil; e = e.Next() {
		if date >= e.Value.(*rpcResultWaiting).date {
			logx.Infof("onTimer msgId: %v", e.Value)
			msgIdList = append(msgIdList, e.Value.(*rpcResultWaiting).msgId)
			q.q.Remove(e)
		}
	}
	return
}
