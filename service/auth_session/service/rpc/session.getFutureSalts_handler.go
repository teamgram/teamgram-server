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

const (
    kDefaultSaltNum = 32
)

// session.getFutureSalts auth_key_id:long num:int = FutureSalts;
func (s *SessionServiceImpl) SessionGetFutureSalts(ctx context.Context, request *mtproto.TLSessionGetFutureSalts) (*mtproto.FutureSalts, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("session.getFutureSalts - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    num := request.GetNum()
    if num == 0 {
        num = kDefaultSaltNum
    }
    futureSalts, err := s.AuthSessionModel.GetFutureSalts(request.GetAuthKeyId(), num)
    if err != nil {
        glog.Error("session.getFutureSalts - ", err)
        return nil, err
    }

    glog.Infof("session.getFutureSalts - reply: {%s}", logger.JsonDebugData(futureSalts))
    return futureSalts.To_FutureSalts(), nil
}
