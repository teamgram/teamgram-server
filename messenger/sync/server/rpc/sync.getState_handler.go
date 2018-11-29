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
    "time"
)

// sync.getState auth_key_id:long user_id:int = updates.State;
func (s *SyncServiceImpl) SyncGetState(ctx context.Context, request *mtproto.TLSyncGetState) (*mtproto.Updates_State, error) {
    glog.Infof("sync.getState - request: {%s}", logger.JsonDebugData(request))

    // state := s.UpdateModel.GetServerUpdatesState(request.GetAuthKeyId(), request.GetUserId())
    state := &mtproto.TLUpdatesState{Data2: &mtproto.Updates_State_Data{
        Pts:         int32(s.CurrentPtsId(request.UserId)),
        Qts:         0,
        Seq:         -1,
        Date:        int32(time.Now().Unix()), // TODO(@benqi): do.Date2???
        UnreadCount: 0,
    }}

    glog.Infof("getServerUpdatesState - reply: %s", logger.JsonDebugData(state))
    return state.To_Updates_State(), nil
}
