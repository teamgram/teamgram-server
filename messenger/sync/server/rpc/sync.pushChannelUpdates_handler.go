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

package rpc

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	zrpc "github.com/nebula-chat/chatengine/mtproto/rpc"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// sync.pushChannelUpdates channel_id:int user_id:int updates:Updates = Bool;
func (s *SyncServiceImpl) SyncPushChannelUpdates(ctx context.Context, request *mtproto.TLSyncPushChannelUpdates) (*mtproto.Bool, error) {
	glog.Infof("sync.pushChannelUpdates - request: {%s}", logger.JsonDebugData(request))
	pts, ptsCount, err := s.processChannelUpdatesRequest(request.GetUserId(), request.GetChannelId(), request.GetUpdates())
	if err == nil {
		userId := request.GetUserId()
		cntl := zrpc.NewController()
		pushData := request.GetUpdates().Encode()
		s.pushUpdatesToSession(syncTypeUser, userId, 0, 0, cntl, pushData, 0, pts, ptsCount)
	}
	glog.Infof("sync.pushChannelUpdates - reply: {true}")
	return mtproto.ToBool(true), nil
}
