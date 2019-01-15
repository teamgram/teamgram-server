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

package server

import (
	"container/list"
	"github.com/golang/glog"
	"reflect"
)

type updatesManager struct {
	genericSessions *list.List
	// *genericSession
	*pushSession
}

func (m *updatesManager) getOnlineGenericSession() *genericSession {
	var lastReceiveTime int64 = 0
	var lastRecentSess *genericSession
	for e := m.genericSessions.Back(); e != nil; e = e.Next() {
		sess := e.Value.(*genericSession)
		if sess.sessionOnline() && sess.lastReceiveTime > lastReceiveTime {
			lastRecentSess = sess
			lastReceiveTime = sess.lastReceiveTime
		}
	}
	return lastRecentSess
}

func (m *updatesManager) onGenericSessionNew(s sessionBase) {
	ss := s.(*genericSession)

	for e := m.genericSessions.Front(); e != nil; e = e.Next() {
		if ss.SessionId() == e.Value.(*genericSession).sessionId {
			*e.Value.(*genericSession) = *ss
			return
		}
	}

	m.genericSessions.PushBack(ss)
}

func (m *updatesManager) onGenericSessionClose(s sessionBase) {
	if ss, ok := s.(*genericSession); ok {
		for e := m.genericSessions.Front(); e != nil; e = e.Next() {
			if ss.SessionId() == e.Value.(*genericSession).sessionId {
				m.genericSessions.Remove(e)
				break
			}
		}
	}
}

func (m *updatesManager) onPushSessionNew(s sessionBase) {
	m.pushSession = s.(*pushSession)
}

func (m *updatesManager) onPushSessionClose() {
	m.pushSession = nil
}

func (m *updatesManager) onUpdatesSyncData(syncMsg *syncData) {
	glog.Infof("onSyncData - generic session: {pts: %d, pts_count: %d, updates: %s}",
		syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))

	genericSess := m.getOnlineGenericSession()
	pushSess := m.pushSession

	if genericSess != nil {
		glog.Infof("updatesManager]>> - generic session: {sess: %s, pts: %d, pts_count: %d, updates: %s}",
			genericSess, syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
		genericSess.onSyncData(syncMsg.cntl, syncMsg.data.obj)
	} else {
		if pushSess != nil && pushSess.sessionOnline() && syncMsg.ptsCount > 0 {
			glog.Infof("updatesManager]]>> - push session: {sess: %s, pts: %d, pts_count: %d, updates: %s}",
				pushSess, syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
			pushSess.onSyncData(syncMsg.cntl)
		} else {
			glog.Infof("updatesManager]]>> - push session: {sess: %s, pts: %d, pts_count: %d, updates: %s}",
				pushSess, syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
		}
	}

	/*
		if pushSess != nil {
			glog.Infof("updatesManager]]>> - push session: {sess: %s, pts: %d, pts_count: %d, updates: %s}",
				pushSess, syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
			pushSess.onSyncData(syncMsg.cntl)

			if genericSess != nil {
				glog.Infof("updatesManager]>> - generic session: {sess: %s, pts: %d, pts_count: %d, updates: %s}",
					genericSess, syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
				genericSess.onSyncData(syncMsg.cntl, syncMsg.data.obj)
			}
		} else {
			if genericSess != nil {
				glog.Infof("updatesManager]]>> - generic session: {sess: %s, pts: %d, pts_count: %d, updates: %s}",
					genericSess, syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
				genericSess.onSyncData(syncMsg.cntl, syncMsg.data.obj)
			}
		}
	*/

	//if m.pushSession.sessionOnline() {
	//	if syncMsg.ptsCount > 0 {
	//		glog.Infof("onSyncData - push session: {pts: %d, pts_count: %d, updates: %s}",
	//			syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
	//		m.pushSession.onSyncData(syncMsg.cntl)
	//
	//		if m.genericSession.sessionOnline() {
	//			glog.Infof("onSyncData - generic session: {pts: %d, pts_count: %d, updates: %s}",
	//				syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
	//			m.genericSession.onSyncData(syncMsg.cntl, syncMsg.data.obj)
	//		}
	//	}
	//} else {
	//	if m.genericSession.sessionOnline() {
	//		glog.Infof("onSyncData - generic session: {pts: %d, pts_count: %d, updates: %s}",
	//			syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
	//		m.genericSession.onSyncData(syncMsg.cntl, syncMsg.data.obj)
	//	}
	//}
}

func (m *updatesManager) onUpdatesSyncRpcResultData(syncMsg *syncRpcResultData) {
	//if m.genericSession != nil {
	//	m.genericSession.onSyncRpcResultData(syncMsg.cntl, syncMsg.data)
	//}
}
