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
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
	"time"
)

// account.updateStatus#6628562c offline:Bool = Bool;
func (s *AccountServiceImpl) AccountUpdateStatus(ctx context.Context, request *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.updateStatus#6628562c - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// var status *mtproto.UserStatus

	offline := mtproto.FromBool(request.GetOffline())
	if offline {
		//	// pc端：离开应用程序激活状态（点击其他应用程序）
		//	statusOffline := &mtproto.TLUserStatusOffline{Data2: &mtproto.UserStatus_Data{
		//		WasOnline: int32(time.Now().Unix()),
		//	}}
		//	status = statusOffline.To_UserStatus()
		//} else {
		// pc端：客户端应用程序激活（点击客户端窗口）
		now := time.Now().Unix()
		//statusOnline := &mtproto.TLUserStatusOnline{Data2: &mtproto.UserStatus_Data{
		//	Expires: int32(now + 60),
		//}}
		//status = statusOnline.To_UserStatus()
		s.UserModel.UpdateUserStatus(md.UserId, now)
	}

	// push to other contacts.
	contactIDList := s.UserModel.GetContactUserIDList(md.UserId)
	for _, id := range contactIDList {
		if md.UserId == id {
			// why??
			continue
		}

		blocked := s.UserModel.IsBlockedByUser(md.UserId, id)
		if blocked {
			continue
		}

		//if !s.AccountModel.CheckShowStatus(md.UserId, id,true) {
		//	continue
		//}

		// TODO(@benqi): dialog or contact???
		status := s.UserModel.GetUserStatus2(id, md.UserId, true, false)

		updateUserStatus := &mtproto.TLUpdateUserStatus{Data2: &mtproto.Update_Data{
			UserId: md.UserId,
			Status: status,
		}}
		updates := &mtproto.TLUpdateShort{Data2: &mtproto.Updates_Data{
			Update: updateUserStatus.To_Update(),
			Date:   int32(time.Now().Unix()),
		}}

		// log.Debugf("updateStatus - toId:{%d}", id)
		sync_client.GetSyncClient().PushUpdates(id, updates.To_Updates())
	}

	glog.Infof("account.updateStatus#6628562c - reply: {true}")
	return mtproto.ToBool(true), nil
}
