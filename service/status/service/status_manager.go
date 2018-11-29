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

package service

import (
	"container/list"
	"github.com/nebula-chat/chatengine/service/status/proto"
	"sync"
)

//var (
//	statuses  = &statusManager{}
//)

//type sessionEntry struct {
//	endpointId  int32
//	userId    int32
//	authKeyId int64
//	expired   int64
//}

type statusManager struct {
	statues map[int32]*list.List
	lock    sync.Mutex
}

func equalSessionEntry(e1, e2 *status.SessionEntry) (b bool) {
	if e1.UserId == e2.UserId && e1.AuthKeyId == e2.AuthKeyId {
		b = true
	}
	return
}

func (s *statusManager) addOrUpdateSession(session *status.SessionEntry) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if slist, ok := s.statues[session.UserId]; !ok {
		slist = list.New()
		slist.PushBack(session)
		s.statues[session.UserId] = slist
	} else {
		var (
			e *list.Element
		)
		for e = slist.Front(); e != nil; e = e.Next() {
			s2, _ := e.Value.(*status.SessionEntry)
			if equalSessionEntry(session, s2) {
				s2.ServerId = session.ServerId
				s2.Expired = session.Expired
				s2.Layer = session.Layer
				break
			}
		}

		if e == nil {
			slist.PushBack(session)
		}
	}
}

func (s *statusManager) removeSession(session *status.SessionEntry) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if slist, ok := s.statues[session.UserId]; !ok {
		return
	} else {
		var (
			e *list.Element
		)
		for e = slist.Front(); e != nil; e = e.Next() {
			s2, _ := e.Value.(*status.SessionEntry)
			if equalSessionEntry(session, s2) {
				slist.Remove(e)
				break
			}
		}
	}
}

func (s *statusManager) querySessionsByUserID(id int32) (sessions []*status.SessionEntry) {
	s.lock.Lock()
	defer s.lock.Unlock()

	sessions = s.querySessionsByUserIDInternal(id)
	return
}

func (s *statusManager) querySessionsByUserIDInternal(id int32) (sessions []*status.SessionEntry) {
	if slist, ok := s.statues[id]; !ok {
		sessions = []*status.SessionEntry{}
	} else {
		sessions = make([]*status.SessionEntry, 0, slist.Len())
		for e := slist.Front(); e != nil; e = e.Next() {
			sessions = append(sessions, e.Value.(*status.SessionEntry))
		}
	}
	return
}

func (s *statusManager) querySessionsListByUserIDList(idList []int32) (sessionsMap map[int32]*status.SessionEntryList) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, id := range idList {
		sessionsMap[id] = &status.SessionEntryList{
			Sessions: s.querySessionsByUserIDInternal(id),
		}
	}

	return
}
