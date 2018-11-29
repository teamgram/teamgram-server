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
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/update"
)

// account.updateProfile#78515775 flags:# first_name:flags.0?string last_name:flags.1?string about:flags.2?string = User;
func (s *AccountServiceImpl) AccountUpdateProfile(ctx context.Context, request *mtproto.TLAccountUpdateProfile) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.updateProfile#78515775 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	user := s.UserModel.GetUserById(md.UserId, md.UserId)

	var isUpdateAbout = request.FirstName == "" && request.LastName == ""
	if isUpdateAbout {
		//// about长度<70并且可以为emtpy
		if len(request.About) > 70 {
			// TODO(@benqi): return error
		}

		s.AccountModel.UpdateAbout(md.UserId, request.About)
	} else {
		s.AccountModel.UpdateFirstAndLastName(md.UserId, request.FirstName, request.LastName)
		user.Data2.FirstName = request.FirstName
		user.Data2.LastName = request.LastName

		// sync to other sessions
		// updateUserName#a7332b73 user_id:int first_name:string last_name:string username:string = Update;
		updateUserName := &mtproto.TLUpdateUserName{Data2: &mtproto.Update_Data{
			UserId:    md.UserId,
			FirstName: user.Data2.FirstName,
			LastName:  user.Data2.LastName,
			Username:  user.Data2.Username,
		}}

		syncUpdates := updates.NewUpdatesLogicByUpdate(md.UserId, updateUserName.To_Update())
		sync_client.GetSyncClient().SyncUpdatesNotMe(md.UserId, md.AuthId, syncUpdates.ToUpdateShort())
	}

	glog.Infof("account.updateProfile#78515775 - reply: {%v}", user)
	return user, nil
}
