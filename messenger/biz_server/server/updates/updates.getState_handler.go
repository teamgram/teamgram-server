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

package updates

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
)

// 执行getState后，获取最新的pts, qts and seq
// updates.getState#edd4882a = updates.State;
func (s *UpdatesServiceImpl) UpdatesGetState(ctx context.Context, request *mtproto.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("updates.getState#edd4882a  - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	state, err := sync_client.GetSyncClient().SyncGetState(md.AuthId, md.UserId)

	glog.Infof("updates.getState#edd4882a  - reply: %s", logger.JsonDebugData(state))
	return state, err
}
