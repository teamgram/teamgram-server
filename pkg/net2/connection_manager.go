// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package net2

import (
	"sync"
)

const connectionMapNum = 32

type ConnectionManager struct {
	connectionMaps [connectionMapNum]connectionMap
	disposeFlag    bool
	disposeOnce    sync.Once
	disposeWait    sync.WaitGroup
}

type connectionMap struct {
	sync.RWMutex
	conns    map[uint64]Connection
	disposed bool
}

func NewConnectionManager() *ConnectionManager {
	manager := &ConnectionManager{}
	for i := 0; i < len(manager.connectionMaps); i++ {
		manager.connectionMaps[i].conns = make(map[uint64]Connection)
	}
	return manager
}

func (manager *ConnectionManager) Dispose() {
	manager.disposeOnce.Do(func() {
		manager.disposeFlag = true
		for i := 0; i < connectionMapNum; i++ {
			conns := &manager.connectionMaps[i]
			conns.Lock()
			conns.disposed = true
			for _, conn := range conns.conns {
				conn.Close()
			}
			conns.Unlock()
		}
		manager.disposeWait.Wait()
	})
}

func (manager *ConnectionManager) GetConnection(connID uint64) Connection {
	conns := &manager.connectionMaps[connID%connectionMapNum]
	conns.RLock()
	defer conns.RUnlock()

	conn, _ := conns.conns[connID]
	return conn
}

func (manager *ConnectionManager) putConnection(conn Connection) {
	conns := &manager.connectionMaps[conn.GetConnID()%connectionMapNum]

	conns.Lock()
	defer conns.Unlock()

	if conns.disposed {
		conn.Close()
		return
	}

	conns.conns[conn.GetConnID()] = conn
	manager.disposeWait.Add(1)
}

func (manager *ConnectionManager) delConnection(conn Connection) {
	if manager.disposeFlag {
		manager.disposeWait.Done()
		return
	}

	conns := &manager.connectionMaps[conn.GetConnID()%connectionMapNum]

	conns.Lock()
	defer conns.Unlock()

	delete(conns.conns, conn.GetConnID())
	manager.disposeWait.Done()
}

//type Manager struct {
//	sessionMaps [sessionMapNum]sessionMap
//	disposeOnce sync.Once
//	disposeWait sync.WaitGroup
//}
//
//type sessionMap struct {
//	sync.RWMutex
//	sessions map[uint64]*Session
//	disposed bool
//}
//
//func NewManager() *Manager {
//	manager := &Manager{}
//	for i := 0; i < len(manager.sessionMaps); i++ {
//		manager.sessionMaps[i].sessions = make(map[uint64]*Session)
//	}
//	return manager
//}
//
//func (manager *Manager) Dispose() {
//	manager.disposeOnce.Do(func() {
//		for i := 0; i < sessionMapNum; i++ {
//			smap := &manager.sessionMaps[i]
//			smap.Lock()
//			smap.disposed = true
//			for _, session := range smap.sessions {
//				session.Close()
//			}
//			smap.Unlock()
//		}
//		manager.disposeWait.Wait()
//	})
//}
//
//func (manager *Manager) NewSession(codec Codec, sendChanSize int) *Session {
//	session := newSession(manager, codec, sendChanSize)
//	manager.putSession(session)
//	return session
//}
//
//func (manager *Manager) GetSession(sessionID uint64) *Session {
//	smap := &manager.sessionMaps[sessionID%sessionMapNum]
//	smap.RLock()
//	defer smap.RUnlock()
//
//	session, _ := smap.sessions[sessionID]
//	return session
//}
//
//func (manager *Manager) putSession(session *Session) {
//	smap := &manager.sessionMaps[session.id%sessionMapNum]
//
//	smap.Lock()
//	defer smap.Unlock()
//
//	if smap.disposed {
//		session.Close()
//		return
//	}
//
//	smap.sessions[session.id] = session
//	manager.disposeWait.Add(1)
//}
//
//func (manager *Manager) delSession(session *Session) {
//	smap := &manager.sessionMaps[session.id%sessionMapNum]
//
//	smap.Lock()
//	defer smap.Unlock()
//
//	delete(smap.sessions, session.id)
//	manager.disposeWait.Done()
//}
