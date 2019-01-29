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
	"github.com/golang/glog"
	"reflect"
)

type updatesManager struct {
	//genericSessions *list.List
	// *genericSession
	//*pushSession
	sessions *authSessions
}

func (m *updatesManager) getOnlineGenericSession() *genericSession {
	s := m.sessions.getOnlineGenericSession()
	if s != nil {
		return s.(*genericSession)
	}
	return nil
}

func (m *updatesManager) getOnlinePushSession() *pushSession {
	s := m.sessions.getOnlinePushSession()
	if s != nil {
		return s.(*pushSession)
	}
	return nil
}

func (m *updatesManager) onUpdatesSyncData(syncMsg *syncData) {
	glog.Infof("onSyncData - generic session: {pts: %d, pts_count: %d, updates: %s}",
		syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))

	genericSess := m.getOnlineGenericSession()
	// pushSess := m.pushSession

	if genericSess != nil {
		glog.Infof("updatesManager]>> - generic session: {sess: %s, pts: %d, pts_count: %d, updates: %s}",
			genericSess, syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
		genericSess.onSyncData(syncMsg.cntl, syncMsg.data.obj)
	} else {
		pushSess := m.getOnlinePushSession()
		if pushSess != nil {
			glog.Infof("updatesManager]]>> - push session: {sess: %s, pts: %d, pts_count: %d, updates: %s}",
				pushSess, syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
			if syncMsg.ptsCount > 0 {
				pushSess.onSyncData(syncMsg.cntl)
			}
		} else {
			glog.Infof("updatesManager]]>> - push session: {sess: nil, pts: %d, pts_count: %d, updates: %s}",
				syncMsg.pts, syncMsg.ptsCount, reflect.TypeOf(syncMsg.data.obj))
		}
	}
}

func (m *updatesManager) onUpdatesSyncRpcResultData(syncMsg *syncRpcResultData) {
	//if m.genericSession != nil {
	//	m.genericSession.onSyncRpcResultData(syncMsg.cntl, syncMsg.data)
	//}
}
