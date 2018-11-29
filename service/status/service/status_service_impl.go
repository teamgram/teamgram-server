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
	"context"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/service/status/proto"
)

type statusServiceImpl struct {
	statuses *statusManager
}

func newStatusServiceImpl() *statusServiceImpl {
	return &statusServiceImpl{
		statuses: &statusManager{
			statues: make(map[int32]*list.List),
		},
	}
}

// rpc SetSessionOnline (SessionEntry) returns (Void);
func (s *statusServiceImpl) SetSessionOnline(ctx context.Context, request *status.SessionEntry) (*status.Void, error) {
	glog.Infof("status.SetSessionOnline - request: %s", logger.JsonDebugData(request))

	s.statuses.addOrUpdateSession(request)
	reply := &status.Void{}

	glog.Infof("status.SetSessionOnline - reply: {%v}", reply)
	return reply, nil
}

// rpc SetSessionOffline (SessionEntry) returns (Void);
func (s *statusServiceImpl) SetSessionOffline(ctx context.Context, request *status.SessionEntry) (*status.Void, error) {
	glog.Infof("status.SetSessionOffline - request: %s", logger.JsonDebugData(request))

	s.statuses.removeSession(request)
	reply := &status.Void{}

	glog.Infof("status.SetSessionOffline - reply: {%v}", reply)
	return reply, nil
}

// rpc GetUserOnlineSessions (Int32) returns (SessionEntryList);
func (s *statusServiceImpl) GetUserOnlineSessions(ctx context.Context, request *status.Int32) (*status.SessionEntryList, error) {
	glog.Infof("status.GetUserOnlineSessions - request: %s", logger.JsonDebugData(request))

	sessions := s.statuses.querySessionsByUserID(request.GetV())
	reply := &status.SessionEntryList{
		Sessions: sessions,
	}

	glog.Infof("status.GetUserOnlineSessions - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}

// rpc GetUsersOnlineSessionsList (Int32List) returns (UsersSessionEntryList);
func (s *statusServiceImpl) GetUsersOnlineSessionsList(ctx context.Context, request *status.Int32List) (*status.UsersSessionEntryList, error) {
	glog.Infof("status.GetUsersOnlineSessionsList - request: %s", logger.JsonDebugData(request))

	sessionMap := s.statuses.querySessionsListByUserIDList(request.GetVlist())
	reply := &status.UsersSessionEntryList{
		UsersSessions: sessionMap,
	}

	glog.Infof("status.GetUsersOnlineSessionsList - reply:: %s", logger.JsonDebugData(reply))
	return nil, nil
}
