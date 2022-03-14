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

package server

import (
	"container/list"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

type sessionData struct {
	sessionId           int64
	connIdList          *list.List
	pendingHttpDataList *list.List
}

type authSession struct {
	authKey     *authKeyUtil
	sessionList map[int64]sessionData
}

type authSessionManager struct {
	rw       sync.RWMutex
	sessions map[int64]*authSession
}

func NewAuthSessionManager() *authSessionManager {
	return &authSessionManager{
		sessions: make(map[int64]*authSession),
	}
}

func (m *authSessionManager) PushBackHttpData(authKeyId, sessionId int64, raw []byte) {
	m.rw.Lock()
	defer m.rw.Unlock()

	if v, ok := m.sessions[authKeyId]; ok {
		if v2, ok2 := v.sessionList[sessionId]; ok2 {
			v2.pendingHttpDataList.PushFront(raw)
		}
	}
}

func (m *authSessionManager) PopFrontHttpData(authKeyId, sessionId int64) []byte {
	m.rw.Lock()
	defer m.rw.Unlock()

	if v, ok := m.sessions[authKeyId]; ok {
		if v2, ok2 := v.sessionList[sessionId]; ok2 {
			if e := v2.pendingHttpDataList.Front(); e != nil {
				v2.pendingHttpDataList.Remove(e)
				return e.Value.([]byte)
			}
		}
	}
	return nil
}

func (m *authSessionManager) AddNewSession(authKey *authKeyUtil, sessionId int64, connId uint64) (bNew bool) {
	logx.Infof("addNewSession: auth_key_id: %d, session_id: %d, conn_id: %d",
		authKey.AuthKeyId(),
		sessionId,
		connId)

	m.rw.Lock()
	defer m.rw.Unlock()

	if v, ok := m.sessions[authKey.AuthKeyId()]; ok {
		var (
			// sIdx     = -1
			cExisted = false
		)
		if v2, ok2 := v.sessionList[sessionId]; ok2 {
			for e := v2.connIdList.Front(); e != nil; e = e.Next() {
				if e.Value.(uint64) == connId {
					cExisted = true
					break
				}
			}
			if !cExisted {
				v2.connIdList.PushBack(connId)
			}
		} else {
			s := sessionData{
				sessionId:           sessionId,
				connIdList:          list.New(),
				pendingHttpDataList: list.New(),
			}
			s.connIdList.PushBack(connId)
			v.sessionList[sessionId] = s
			bNew = true
		}
	} else {
		s := sessionData{
			sessionId:           sessionId,
			connIdList:          list.New(),
			pendingHttpDataList: list.New(),
		}
		s.connIdList.PushBack(connId)

		m.sessions[authKey.AuthKeyId()] = &authSession{
			authKey: authKey,
			sessionList: map[int64]sessionData{
				sessionId: s,
			},
		}
		bNew = true
	}
	return
}

func (m *authSessionManager) RemoveSession(authKeyId, sessionId int64, connId uint64) (bDeleted bool) {
	logx.Infof("removeSession: auth_key_id: %d, session_id: %d, conn_id: %d",
		authKeyId,
		sessionId,
		connId)

	m.rw.Lock()
	defer m.rw.Unlock()

	if v, ok := m.sessions[authKeyId]; ok {
		if v2, ok2 := v.sessionList[sessionId]; ok2 {
			for e := v2.connIdList.Front(); e != nil; e = e.Next() {
				if e.Value.(uint64) == connId {
					v2.connIdList.Remove(e)
					break
				}
			}
			if v2.connIdList.Len() == 0 {
				delete(v.sessionList, sessionId)
				bDeleted = true
			}
			if len(v.sessionList) == 0 {
				delete(m.sessions, authKeyId)
			}
		}
	}

	return
}

func (m *authSessionManager) FoundSessionConnIdList(authKeyId, sessionId int64) (*authKeyUtil, []uint64) {
	m.rw.RLock()
	defer m.rw.RUnlock()

	if v, ok := m.sessions[authKeyId]; ok {
		if v2, ok2 := v.sessionList[sessionId]; ok2 {
			connIdList := make([]uint64, 0, v2.connIdList.Len())
			for e := v2.connIdList.Back(); e != nil; e = e.Prev() {
				connIdList = append(connIdList, e.Value.(uint64))
			}
			return v.authKey, connIdList
		}
	}

	return nil, nil
}
