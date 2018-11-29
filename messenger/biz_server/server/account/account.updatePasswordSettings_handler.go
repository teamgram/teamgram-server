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
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

// account.updatePasswordSettings#fa7c4b86 current_password_hash:bytes new_settings:account.PasswordInputSettings = Bool;
func (s *AccountServiceImpl) AccountUpdatePasswordSettings(ctx context.Context, request *mtproto.TLAccountUpdatePasswordSettings) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.updatePasswordSettings#fa7c4b86 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	passwordInputSetting := request.NewSettings.To_AccountPasswordInputSettings()
	_ = passwordInputSetting
	// TODO(@benqi): check request invalid

	passwordLogic, err := s.AccountModel.MakePasswordData(md.UserId)
	_ = passwordLogic
	if err == nil {
		//err = passwordLogic.UpdatePasswordSetting(request.CurrentPasswordHash,
		//	passwordInputSetting.GetNewSalt(),
		//	passwordInputSetting.GetNewPasswordHash(),
		//	passwordInputSetting.GetHint(),
		//	passwordInputSetting.GetEmail())

		// 未注册：error_message: "EMAIL_UNCONFIRMED" [STRING],
	}

	if err != nil {
		glog.Error("account.updatePasswordSettings#fa7c4b86 - error: ", err)
		return nil, err
	}

	reply := mtproto.ToBool(true)
	glog.Infof("account.getPassword#548a30f5 - reply: {}", logger.JsonDebugData(reply))
	return reply, nil
}
