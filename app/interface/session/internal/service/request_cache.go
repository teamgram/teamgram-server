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
	"errors"
	"fmt"
	"sync"

	"github.com/teamgram/proto/mtproto"
)

const mapNum = 8

type requestMap struct {
	mutex    sync.Mutex
	respChan map[int64]chan mtproto.TLObject
}

type RequestManager struct {
	requestMaps [mapNum]requestMap
	current     uint32
}

func NewRequestManager() *RequestManager {
	manager := &RequestManager{}
	for i := 0; i < mapNum; i++ {
		manager.requestMaps[i].respChan = make(map[int64]chan mtproto.TLObject)
	}

	return manager
}

func (r *RequestManager) cache(id int64, msg chan mtproto.TLObject) {
	m := &r.requestMaps[uint64(id)%mapNum]
	m.mutex.Lock()
	m.respChan[id] = msg
	m.mutex.Unlock()
}

func (r *RequestManager) shoot(id int64, msg mtproto.TLObject) (err error) {
	m := &r.requestMaps[uint64(id)%mapNum]

	m.mutex.Lock()
	respChans, exist := m.respChan[id]
	if exist {
		delete(m.respChan, id)
		m.mutex.Unlock()
		//如果此时无法写入，说明resp_chan已经不可用，要立马返回
		select {
		case respChans <- msg:
			//只使用一次，写入完成后关闭；
			close(respChans)
		default:
			err = errors.New(fmt.Sprint("Default fail !!!!! request id : ", id))
		}
	} else {
		m.mutex.Unlock()
		err = errors.New("default fail ")
	}
	return
}

func (r *RequestManager) dispose(id int64) {
	m := &r.requestMaps[uint64(id)%mapNum]
	m.mutex.Lock()
	respChans, exist := m.respChan[id]
	if exist {
		delete(m.respChan, id)
		m.mutex.Unlock()
		close(respChans)
	} else {
		//r.disposeWait.Done()
		m.mutex.Unlock()
	}
}
