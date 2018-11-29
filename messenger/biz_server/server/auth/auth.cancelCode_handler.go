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

package auth

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"golang.org/x/net/context"
)

// 客户端不处理错误码

// auth.cancelCode#1f040578 phone_number:string phone_code_hash:string = Bool;
func (s *AuthServiceImpl) AuthCancelCode(ctx context.Context, request *mtproto.TLAuthCancelCode) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.cancelCode#1f040578 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// 1. check phone code
	if request.PhoneCodeHash == "" {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_BAD_REQUEST), "auth.resendCode#3ef1a9bf: phone code hash empty.")
		return nil, err
	}

	// 2. check number
	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	phoneNumber, err := base.CheckAndGetPhoneNumber(request.GetPhoneNumber())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	code := s.AuthModel.MakeCancelCodeData(md.AuthId, phoneNumber, request.PhoneCodeHash)
	canceled := mtproto.ToBool(code.DoCancelCode())

	glog.Infof("auth.cancelCode#1f040578 -  - reply: %s", logger.JsonDebugData(canceled))
	return canceled, nil
}
