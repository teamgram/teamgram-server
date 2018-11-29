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

// account.registerDevice#637ea878 token_type:int token:string = Bool;
func (s *AccountServiceImpl) AccountRegisterDeviceLayer71(ctx context.Context, request *mtproto.TLAccountRegisterDeviceLayer71) (*mtproto.Bool, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("account.registerDevice#637ea878 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    // Check token format by token_type
    // TODO(@benqi): check token format by token_type
    if request.Token == "" {
        err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
        glog.Error(err)
        return nil, err
    }

    // TODO(@benqi): check toke_type invalid
    if request.TokenType < core.TOKEN_TYPE_APNS || request.TokenType > core.TOKEN_TYPE_MAXSIZE {
        // glog.Error("request.TokenType: ", request.TokenType)
        err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
        glog.Error(err)
        return nil, err
    }

    registerDevice := &mtproto.TLAccountRegisterDevice{
        TokenType:  request.TokenType,
        Token:      request.Token,
        AppSandbox: mtproto.ToBool(false),
        Secret:     []byte{},
        OtherUids:  []int32{},
    }

    registered := s.AccountModel.RegisterDevice(md.AuthId, md.UserId, registerDevice)

    glog.Infof("account.registerDevice#637ea878 - reply: {true}")
    return mtproto.ToBool(registered), nil
}
