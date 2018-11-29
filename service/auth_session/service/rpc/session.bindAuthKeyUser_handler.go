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
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
)

// session.bindAuthKeyUser auth_key_id:long user_id:int = Bool;
func (s *SessionServiceImpl) SessionBindAuthKeyUser(ctx context.Context, request *mtproto.TLSessionBindAuthKeyUser) (*mtproto.Bool, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("session.bindAuthKeyUser - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    reply := s.AuthSessionModel.BindAuthKeyUser(request.GetAuthKeyId(), request.GetUserId())

    glog.Infof("session.bindAuthKeyUser - reply: {%v}", logger.JsonDebugData(reply))
    return mtproto.ToBool(reply), nil
}
