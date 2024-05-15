// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
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

package server

import (
	"sync"
)

type authSessionManager struct {
	rw       sync.RWMutex
	sessions map[int64]map[int64]int64
}

func NewAuthSessionManager() *authSessionManager {
	return &authSessionManager{
		sessions: make(map[int64]map[int64]int64),
	}
}

func (m *authSessionManager) AddNewSession(authId int64, sessionId int64, connId int64) (bNew bool) {
	m.rw.Lock()
	defer m.rw.Unlock()

	if v, ok := m.sessions[authId]; ok {
		v[sessionId] = connId
	} else {
		m.sessions[authId] = map[int64]int64{
			sessionId: connId,
		}
		bNew = true
	}

	return
}

func (m *authSessionManager) RemoveSession(authKeyId, sessionId int64, connId int64) (bDeleted bool) {
	m.rw.Lock()
	defer m.rw.Unlock()

	if v, ok := m.sessions[authKeyId]; ok {
		if _, ok2 := v[sessionId]; ok2 {
			bDeleted = true
			delete(v, sessionId)
			if len(v) == 0 {
				delete(m.sessions, authKeyId)
			}
		}
	}

	return
}

func (m *authSessionManager) FoundSessionConnId(authKeyId, sessionId int64) (int64, bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()

	if v, ok := m.sessions[authKeyId]; ok {
		v2, ok2 := v[sessionId]
		return v2, ok2
	}

	return 0, false
}
