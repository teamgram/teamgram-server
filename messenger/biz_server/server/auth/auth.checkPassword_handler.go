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
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

/*
	if (error.text.equals("PASSWORD_HASH_INVALID")) {
		onPasscodeError(true);
	} else if (error.text.startsWith("FLOOD_WAIT")) {
		int time = Utilities.parseInt(error.text);
		String timeString;
		if (time < 60) {
			timeString = LocaleController.formatPluralString("Seconds", time);
		} else {
			timeString = LocaleController.formatPluralString("Minutes", time / 60);
		}
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.formatString("FloodWaitTime", R.string.FloodWaitTime, timeString));
	} else {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), error.text);
	}
*/

// 客户端调用auth.signIn时返回SESSION_PASSWORD_NEEDED时会触发

// auth.checkPassword#d18b4d16 password:InputCheckPasswordSRP = auth.Authorization;
func (s *AuthServiceImpl) AuthCheckPassword(ctx context.Context, request *mtproto.TLAuthCheckPassword) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.checkPassword#a63011e - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

/*
	var (
		err error
	)

	password := request.GetPassword()

	if len(request.PasswordHash) == 0 {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
		glog.Error(err)
		return nil, err
	}

	passwordLogic, err := s.AccountModel.MakePasswordData(md.UserId)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	ok := passwordLogic.CheckPassword(request.PasswordHash)
	if !ok {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
		glog.Error(err)
		return nil, err
	}

	user := s.UserModel.GetUserById(md.UserId, md.UserId)
	authAuthorization := &mtproto.TLAuthAuthorization{Data2: &mtproto.Auth_Authorization_Data{
		User: user.To_User(),
	}}

	glog.Infof("auth.checkPassword#a63011e - reply: %s\n", logger.JsonDebugData(authAuthorization))
	return authAuthorization.To_Auth_Authorization(), nil
 */

	return nil, fmt.Errorf("not impl auth.checkPassword#d18b4d16")
}
