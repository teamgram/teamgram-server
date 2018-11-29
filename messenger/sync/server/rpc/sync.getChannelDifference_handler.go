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
    "golang.org/x/net/context"
    "fmt"
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
)

// sync.getChannelDifference flags:# auth_key_id:long user_id:int force:flags.0?true channel:InputChannel filter:ChannelMessagesFilter pts:int limit:int = updates.ChannelDifference;
func (s *SyncServiceImpl) SyncGetChannelDifference(ctx context.Context, request *mtproto.TLSyncGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("sync.getChannelDifference - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): Impl SyncGetChannelDifference logic

    return nil, fmt.Errorf("not impl SyncGetChannelDifference")
}
