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

/*
  reset my account?
  delete成功，转到注册页面，失败处理:

	if (error.text.equals("2FA_RECENT_CONFIRM")) {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("ResetAccountCancelledAlert", R.string.ResetAccountCancelledAlert));
	} else if (error.text.startsWith("2FA_CONFIRM_WAIT_")) {
		Bundle params = new Bundle();
		params.putString("phoneFormated", requestPhone);
		params.putString("phoneHash", phoneHash);
		params.putString("code", phoneCode);
		params.putInt("startTime", ConnectionsManager.getInstance().getCurrentTime());
		params.putInt("waitTime", Utilities.parseInt(error.text.replace("2FA_CONFIRM_WAIT_", "")));
		setPage(8, true, params, false);
	} else {
		needShowAlert(LocaleController.getString("AppName", R.string.AppName), error.text);
	}
*/

// account.deleteAccount#418d4e0b reason:string = Bool;
func (s *AccountServiceImpl) AccountDeleteAccount(ctx context.Context, request *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("AccountDeleteAccount - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	deletedOk := s.UserModel.DeleteUser(md.UserId, request.GetReason())

	// TODO(@benqi): 1. Clear account data 2. Kickoff other client

	glog.Infof("AccountDeleteAccount - reply: {%v}", deletedOk)
	return mtproto.ToBool(deletedOk), nil
}
