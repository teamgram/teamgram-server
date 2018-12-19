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

/*
 Android client auth.resendCode#3ef1a9bf, handler error
	if (error.text != null) {
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
		} else if (error.code != -1000) {
			needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred) + "\n" + error.text);
		}
	}
*/

// auth.resendCode#3ef1a9bf phone_number:string phone_code_hash:string = auth.SentCode;
func (s *AuthServiceImpl) AuthResendCode(ctx context.Context, request *mtproto.TLAuthResendCode) (*mtproto.Auth_SentCode, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AuthResendCode - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// 1. check phone code
	if request.PhoneCodeHash == "" {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_BAD_REQUEST), "auth.resendCode#3ef1a9bf: phone code hash empty.")
		glog.Error(err)
		return nil, err
	}

	// 2. check number
	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	phoneNumber, err := base.CheckAndGetPhoneNumber(request.GetPhoneNumber())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	// TODO(@benqi): PHONE_NUMBER_FLOOD
	// <string name="PhoneNumberFlood">Sorry, you have deleted and re-created your account too many times recently.
	//    Please wait for a few days before signing up again.</string>
	//

	// PHONE_NUMBER_BANNED: Banned phone number
	banned := s.AuthModel.CheckBannedByPhoneNumber(phoneNumber)
	if banned {
		err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_BANNED), "auth.sendCode#86aef0ec: phone number banned.")
		glog.Error(err)
		return nil, err
	}

	code := s.AuthModel.MakeCodeDataByHash(md.AuthId, phoneNumber, request.GetPhoneCodeHash())
	err = code.DoReSendCode(getSendSmsFunc())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	// 使用app类型，code统一为123456
	phoneRegistered := s.AuthModel.CheckPhoneNumberExist(phoneNumber)
	authSentCode := code.ToAuthSentCode(phoneRegistered)

	glog.Infof("auth.resendCode#3ef1a9bf - reply: %s", logger.JsonDebugData(authSentCode))
	return authSentCode.To_Auth_SentCode(), nil
}
