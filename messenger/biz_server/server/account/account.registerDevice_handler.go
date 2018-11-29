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

// {"token_type":10,"token":"{\"endpoint\":\"https://android.googleapis.com/gcm/send/fFBQfzHLq0I:APA91bF28ucFXm7ZF2T7sf87dKAARGXHzdK3HbK0rrhJMmPZr42amq6B-QASi-mLzOzZ5qsynyvtEOSNNYbvadNKI5LCxmYMhQXkhoh_fpTB0GsYLBjwpElaV68OmTUzN-AFDgWuqMIpQH5XYDZoYQopg-yHHdsxcQ\",\"expirationTime\":null,\"keys\":{\"p256dh\":\"BJLqPVxd2KNAmW_izYz4ha5hN4ZEzXnNbk4__FC-xhmaa2vZD3RRtvgPNphH8ZSM9wF4_vSTJZLzQ5Iv0byZxrY\",\"auth\":\"nBC8C_1cvhSTlEEelbk9kw\"}}","app_sandbox":{"constructor":-1132882121,"data2":{}}}
// account.registerDevice#5cbea590 token_type:int token:string app_sandbox:Bool secret:bytes other_uids:Vector<int> = Bool;
func (s *AccountServiceImpl) AccountRegisterDevice(ctx context.Context, request *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("account.registerDevice#5cbea590 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

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

    registered := s.AccountModel.RegisterDevice(md.AuthId, md.UserId, request)

    glog.Infof("account.registerDevice#5cbea590 - reply: {true}")
    return mtproto.ToBool(registered), nil
}
