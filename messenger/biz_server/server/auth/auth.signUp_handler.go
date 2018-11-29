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
  Android client auth.signUp#1b067634, handler error
	if (error.text.contains("PHONE_NUMBER_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidPhoneNumber", R.string.InvalidPhoneNumber));
	} else if (error.text.contains("PHONE_CODE_EMPTY") || error.text.contains("PHONE_CODE_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidCode", R.string.InvalidCode));
	} else if (error.text.contains("PHONE_CODE_EXPIRED")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("CodeExpired", R.string.CodeExpired));
	} else if (error.text.contains("FIRSTNAME_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidFirstName", R.string.InvalidFirstName));
	} else if (error.text.contains("LASTNAME_INVALID")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("InvalidLastName", R.string.InvalidLastName));
	} else {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), error.text);
	}

*/
// auth.signUp#1b067634 phone_number:string phone_code_hash:string phone_code:string first_name:string last_name:string = auth.Authorization;
func (s *AuthServiceImpl) AuthSignUp(ctx context.Context, request *mtproto.TLAuthSignUp) (*mtproto.Auth_Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthSignUp - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// 1. check number
	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	//phoneNumber, err := base.CheckAndGetPhoneNumber(request.GetPhoneNumber())
	//if err != nil {
	//	glog.Error(err)
	//	return nil, err
	//}

	pnumber, err := base.MakePhoneNumberUtil(request.GetPhoneNumber(), "")
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	phoneNumber := pnumber.GetNormalizeDigits()

	if request.GetPhoneCode() == "" {
		err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_EMPTY), "phone code empty")
		glog.Error(err)
		return nil, err
	}

	// TODO(@benqi): regist name ruler
	// check first name invalid
	if request.GetFirstName() == "" {
		err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_FIRSTNAME_INVALID), "first name invalid")
		glog.Error(err)
		return nil, err
	}

	// check first name invalid
	//if request.GetLastName() == "" {
	//	err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_LASTNAME_INVALID), "auth.signUp#1b067634: last name invalid")
	//	glog.Error(err)
	//	return nil, err
	//}

	// TODO(@benqi): PHONE_NUMBER_FLOOD
	// <string name="PhoneNumberFlood">Sorry, you have deleted and re-created your account too many times recently.
	//    Please wait for a few days before signing up again.</string>
	//

	phoneRegistered := s.AuthModel.CheckPhoneNumberExist(phoneNumber)
	if phoneRegistered {
		err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_OCCUPIED), "phone number occuiped.")
		glog.Error(err)
		return nil, err
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	code := s.AuthModel.MakeCodeDataByHash(md.AuthId, phoneNumber, request.PhoneCodeHash)
	// phoneRegistered := auth.CheckPhoneNumberExist(phoneNumber)
	err = code.DoSignUp(request.PhoneCode)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	user := s.UserModel.CreateNewUser(phoneNumber, pnumber.GetRegionCode(), request.FirstName, request.LastName)

	// bind auth_key and user_id
	auth_session_client.BindAuthKeyUser(md.AuthId, user.GetId())

	// TODO(@benqi): check and set authKeyId state
	// TODO(@benqi): 修改那些将我的phoneNumber加到他们的联系人列表里的联系人的状态

	// TODO(@benqi): 创建新帐号后执行的事件
	s.UserModel.CreateNewUserPassword(user.GetId())

	authAuthorization := &mtproto.TLAuthAuthorization{Data2: &mtproto.Auth_Authorization_Data{
		User: user.To_User(),
	}}

	glog.Infof("auth.signUp#1b067634 - reply: %s\n", logger.JsonDebugData(authAuthorization))
	return authAuthorization.To_Auth_Authorization(), nil
}
