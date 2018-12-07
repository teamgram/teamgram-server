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
    "golang.org/x/net/context"
    "github.com/nebula-chat/chatengine/mtproto"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto/rpc"
)

// sync.pushUpdates#5c612649 user_id:int updates:Updates = Bool;
func (s *SyncServiceImpl) SyncPushUpdates(ctx context.Context, request *mtproto.TLSyncPushUpdates) (*mtproto.Bool, error) {
    glog.Infof("sync.pushUpdates#5c612649 - request: {%s}", logger.JsonDebugData(request))

    isMessage, err := s.processUpdatesRequest(request.GetUserId(), request.GetUpdates())
    if err == nil {
        userId := request.GetUserId()
        cntl := zrpc.NewController()
        pushData := request.GetUpdates().Encode()
        s.pushUpdatesToSession(syncTypeUser, userId, 0, 0, cntl, pushData, 0, isMessage)
    } else {
        glog.Error(err)
        return mtproto.ToBool(false), nil
    }

    glog.Infof("sync.pushUpdates#5c612649 - reply: {true}",)
    return mtproto.ToBool(true), nil
}
