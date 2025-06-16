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
	"time"
)

type httpReqItem struct {
	expireTime int64
	resChan    chan interface{}
}

type httpRequestQueue struct {
	q *list.List
}

func newHttpRequestQueue() *httpRequestQueue {
	return &httpRequestQueue{
		q: list.New(),
	}
}

func (q *httpRequestQueue) Push(resChan chan interface{}) {
	q.q.PushBack(&httpReqItem{
		expireTime: time.Now().Unix() + 3,
		resChan:    resChan,
	})
}

func (q *httpRequestQueue) Pop() chan interface{} {
	e := q.q.Front()
	if e != nil {
		q.q.Remove(e)
		return e.Value.(*httpReqItem).resChan
	}

	return nil
}

func (q *httpRequestQueue) PopTimeoutList() []chan interface{} {
	var rList []chan interface{}
	date := time.Now().Unix()
	for e := q.q.Front(); e != nil; e = e.Next() {
		if date >= e.Value.(*httpReqItem).expireTime {
			rList = append(rList, e.Value.(*httpReqItem).resChan)
			q.q.Remove(e)
		} else {
			break
		}
	}
	return rList
}

func (q *httpRequestQueue) Clear() {
	for e := q.q.Front(); e != nil; e = e.Next() {
		close(e.Value.(*httpReqItem).resChan)
	}
}
