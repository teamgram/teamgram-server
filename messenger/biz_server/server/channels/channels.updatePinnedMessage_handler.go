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

package channels

import (
    "fmt"
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
)

// channels.updatePinnedMessage#a72ded52 flags:# silent:flags.0?true channel:InputChannel id:int = Updates;
func (s *ChannelsServiceImpl) ChannelsUpdatePinnedMessage(ctx context.Context, request *mtproto.TLChannelsUpdatePinnedMessage) (*mtproto.Updates, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("channels.updatePinnedMessage - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // Sorry: not impl ChannelsUpdatePinnedMessage logic
    glog.Warning("channels.updatePinnedMessage blocked, License key from https://nebula.chat required to unlock enterprise features.")

    return nil, fmt.Errorf("not imp ChannelsUpdatePinnedMessage")
}
