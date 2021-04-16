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
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// sync.syncChannelUpdates flags:# channel_id:int user_id:int auth_key_id:long server_id:flags.0?int updates:Updates = Bool;
func (s *SyncServiceImpl) SyncSyncChannelUpdates(ctx context.Context, request *mtproto.TLSyncSyncChannelUpdates) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("ync.syncChannelUpdate - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	pts, ptsCount, err := s.processChannelUpdatesRequest(request.GetUserId(), request.GetChannelId(), request.GetUpdates())
	if err == nil {
		userId := request.UserId
		authKeyId := request.GetAuthKeyId()
		cntl := zrpc.NewController()
		pushData := request.GetUpdates().Encode()
		serverId := request.GetServerId()
		if request.GetServerId() == 0 {
			s.pushUpdatesToSession(syncTypeUserNotMe, userId, authKeyId, 0, cntl, pushData, 0, pts, ptsCount)
		} else {
			s.pushUpdatesToSession(syncTypeUserMe, userId, authKeyId, 0, cntl, pushData, serverId, pts, ptsCount)
		}
	} else {
		glog.Error(err)
		return mtproto.ToBool(false), nil
	}

	glog.Infof("sync.syncChannelUpdates - reply: {true}")
	return mtproto.ToBool(true), nil
}
