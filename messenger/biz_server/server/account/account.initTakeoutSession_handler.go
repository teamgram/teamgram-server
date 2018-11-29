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

package account

import (
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "fmt"
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
)

// account.initTakeoutSession#f05b4804 flags:# contacts:flags.0?true message_users:flags.1?true message_chats:flags.2?true message_megagroups:flags.3?true message_channels:flags.4?true files:flags.5?true file_max_size:flags.5?int = account.Takeout;
func (s *AccountServiceImpl) AccountInitTakeoutSession(ctx context.Context, request *mtproto.TLAccountInitTakeoutSession) (*mtproto.Account_Takeout, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("AccountInitTakeoutSession - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // TODO(@benqi): Impl AccountInitTakeoutSession logic

    return nil, fmt.Errorf("Not impl AccountInitTakeoutSession")
}
