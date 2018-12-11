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
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
)

// session.getPushSessionId user_id:int auth_key_id:long token_type:int = Int64;
func (s *SessionServiceImpl) SessionGetPushSessionId(ctx context.Context, request *mtproto.TLSessionGetPushSessionId) (*mtproto.Int64, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("session.getPushSessionId - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    sessionId := s.AuthSessionModel.GetPushSessionId(request.GetUserId(), request.GetAuthKeyId(), request.GetTokenType())
    reply := &mtproto.TLLong{Data2: &mtproto.Int64_Data{
        V: sessionId,
    }}

    glog.Infof("session.getPushSessionId - reply: {%s}", logger.JsonDebugData(reply))
    return reply.To_Int64(), nil
}
