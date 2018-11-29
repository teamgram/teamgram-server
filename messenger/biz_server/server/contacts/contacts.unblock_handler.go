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

package contacts

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	updates2 "github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/update"
	"golang.org/x/net/context"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
)

// contacts.unblock#e54100bd id:InputUser = Bool;
func (s *ContactsServiceImpl) ContactsUnblock(ctx context.Context, request *mtproto.TLContactsUnblock) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.unblock#e54100bd - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		blockedId int32
		id        = request.Id
	)

	switch id.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUserSelf:
		blockedId = md.UserId
	case mtproto.TLConstructor_CRC32_inputUser:
		// Check access hash
		if ok := s.UserModel.CheckAccessHashByUserId(id.GetData2().GetUserId(), id.GetData2().GetAccessHash()); !ok {
			// TODO(@benqi): Add ACCESS_HASH_INVALID codes
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			glog.Error(err, ": is access_hash error")
			return nil, err
		}

		blockedId = id.GetData2().GetUserId()
		// TODO(@benqi): contact exist
	default:
		// mtproto.TLConstructor_CRC32_inputUserEmpty:
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err, ": is inputUserEmpty")
		return nil, err
	}

	contactLogic := s.ContactModel.MakeContactLogic(md.UserId)
	unBlocked := contactLogic.UnBlockUser(blockedId)

	if unBlocked {
		// Sync unblocked: updateUserBlocked
		updateUserUnBlocked := &mtproto.TLUpdateUserBlocked{Data2: &mtproto.Update_Data{
			UserId:  blockedId,
			Blocked: mtproto.ToBool(false),
		}}

		unBlockedUpdates := updates2.NewUpdatesLogic(md.UserId)
		unBlockedUpdates.AddUpdate(updateUserUnBlocked.To_Update())
		unBlockedUpdates.AddUser(s.UserModel.GetUserById(md.UserId, blockedId))

		// TODO(@benqi): handle seq
		sync_client.GetSyncClient().PushUpdates(blockedId, unBlockedUpdates.ToUpdates())
	}

	glog.Infof("contacts.unblock#e54100bd - reply: {%v}", unBlocked)
	return mtproto.ToBool(unBlocked), nil
}
