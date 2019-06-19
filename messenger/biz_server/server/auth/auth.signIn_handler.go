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
	"github.com/nebula-chat/chatengine/service/auth_session/client"
)

/*
 1. PHONE_NUMBER_UNOCCUPIED ==> setPage(5, true, params, false);
 2. SESSION_PASSWORD_NEEDED ==> invoke rpc: TL_account_getPassword
 3. error:
	if (error.text.contains("PHONE_NUMBER_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidPhoneNumber", R.string.InvalidPhoneNumber));
	} else if (error.text.contains("PHONE_CODE_EMPTY") || error.text.contains("PHONE_CODE_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidCode", R.string.InvalidCode));
	} else if (error.text.contains("PHONE_CODE_EXPIRED")) {
		onBackPressed();
		setPage(0, true, null, true);
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("CodeExpired", R.string.CodeExpired));
	} else if (error.text.startsWith("FLOOD_WAIT")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("FloodWait", R.string.FloodWait));
	} else {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred) + "\n" + error.text);
	}
*/

// auth.signIn#bcd51581 phone_number:string phone_code_hash:string phone_code:string = auth.Authorization;
func (s *AuthServiceImpl) AuthSignIn(ctx context.Context, request *mtproto.TLAuthSignIn) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("auth.signIn#bcd51581 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	//// 3. check number
	//// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	phoneNumber, err := base.CheckAndGetPhoneNumber(request.GetPhoneNumber())
	if err != nil {
		// PHONE_NUMBER_INVALID
		glog.Error(err)
		return nil, err
	}

	if request.PhoneCode == "" {
		err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_EMPTY), "code empty")
		glog.Error(err)
		return nil, err
	}
	// TODO(@benqi): check phoneCode rule: number, length etc ...

	code := s.AuthModel.MakeCodeDataByHash(md.AuthId, phoneNumber, request.PhoneCodeHash)
	phoneRegistered := s.AuthModel.CheckPhoneNumberExist(phoneNumber)

	// , verifyCodeF func(codeHash, code, extraData string) (error)
	err = code.DoSignIn(request.PhoneCode, phoneRegistered, getVerifySmsCodeF())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	// signIn sucess, check phoneRegistered.
	if !phoneRegistered {
		//  not register, next step: auth.singUp
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PHONE_NUMBER_UNOCCUPIED)
		glog.Info("auth.signIn#bcd51581 - not registered, next step auth.signIn, ", err)
		return nil, err
	}

	// do signIn...
	user := s.UserModel.GetMyUserByPhoneNumber(phoneNumber)
	// Bind authKeyId and userId

	auth_session_client.BindAuthKeyUser(md.AuthId, user.Data2.Id)
	// s.AuthModel.BindAuthKeyAndUser(md.AuthId, user.GetId())
	// TODO(@benqi): check and set authKeyId state

	// Check SESSION_PASSWORD_NEEDED
	sessionPasswordNeeded := s.AccountModel.CheckSessionPasswordNeeded(user.Data2.Id)
	if sessionPasswordNeeded {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_SESSION_PASSWORD_NEEDED)
		glog.Info("auth.signIn#bcd51581 - registered, next step auth.checkPassword, ", err)
		return nil, err
	}

	authAuthorization := &mtproto.TLAuthAuthorization{Data2: &mtproto.Auth_Authorization_Data{
		User: user,
	}}

	glog.Infof("auth.signIn#bcd51581 - reply: %s\n", logger.JsonDebugData(authAuthorization))
	return authAuthorization.To_Auth_Authorization(), nil
}
