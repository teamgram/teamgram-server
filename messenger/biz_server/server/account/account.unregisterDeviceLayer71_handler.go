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
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
    "github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
)

// account.unregisterDevice#65c55b40 token_type:int token:string = Bool;
func (s *AccountServiceImpl) AccountUnregisterDeviceLayer71(ctx context.Context, request *mtproto.TLAccountUnregisterDeviceLayer71) (*mtproto.Bool, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("account.unregisterDevice#65c55b40 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // Check token invalid
    // TODO(@benqi): check token format by token_type
    if request.Token == "" {
        err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
        glog.Error(err)
        return nil, err
    }

    // Check token format by token_type
    if request.TokenType < core.TOKEN_TYPE_APNS || request.TokenType > core.TOKEN_TYPE_INTERNAL_PUSH {
        err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
        glog.Error(err)
        return nil, err
    }

    unregistered := s.AccountModel.UnRegisterDevice(md.AuthId, md.UserId)

    glog.Infof("account.unregisterDevice#65c55b40 - reply: {%v}\n", unregistered)
    return mtproto.ToBool(unregistered), nil
}
