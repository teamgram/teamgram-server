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

// contacts.deleteContact#8e953744 id:InputUser = contacts.Link;
func (s *ContactsServiceImpl) ContactsDeleteContact(ctx context.Context, request *mtproto.TLContactsDeleteContact) (*mtproto.Contacts_Link, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.deleteContact#8e953744 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		deleteId int32
		id       = request.Id
	)

	switch id.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputUserSelf:
		deleteId = md.UserId
	case mtproto.TLConstructor_CRC32_inputUser:
		// Check access hash
		if ok := s.UserModel.CheckAccessHashByUserId(id.GetData2().GetUserId(), id.GetData2().GetAccessHash()); !ok {
			// TODO(@benqi): Add ACCESS_HASH_INVALID codes
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			glog.Error(err, ": is access_hash error")
			return nil, err
		}

		deleteId = id.GetData2().GetUserId()
		// TODO(@benqi): contact exist
	default:
		// mtproto.TLConstructor_CRC32_inputUserEmpty:
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err, ": is inputUserEmpty")
		return nil, err
	}

	// selfUser := user2.GetUserById(md.UserId, md.UserId)
	deleteUser := s.UserModel.GetUserById(md.UserId, deleteId)

	contactLogic := s.ContactModel.MakeContactLogic(md.UserId)
	needUpdate := contactLogic.DeleteContact(deleteId, deleteUser.Data2.MutualContact)

	selfUpdates := updates2.NewUpdatesLogic(md.UserId)
	myLink, foreignLink := s.UserModel.GetContactLink(md.UserId, deleteId)
	contactLink := &mtproto.TLUpdateContactLink{Data2: &mtproto.Update_Data{
		UserId:      deleteId,
		MyLink:      myLink,
		ForeignLink: foreignLink,
	}}
	selfUpdates.AddUpdate(contactLink.To_Update())
	selfUpdates.AddUser(deleteUser)

	// TODO(@benqi): handle seq
	sync_client.GetSyncClient().SyncUpdatesNotMe(md.UserId, md.AuthId, selfUpdates.ToUpdates())

	// TODO(@benqi): 推给联系人逻辑需要再考虑考虑
	if needUpdate {
		// TODO(@benqi): push to contact user update contact link
		contactUpdates := updates2.NewUpdatesLogic(deleteUser.Data2.Id)
		myLink, foreignLink := s.UserModel.GetContactLink(deleteId, md.UserId)
		contactLink2 := &mtproto.TLUpdateContactLink{Data2: &mtproto.Update_Data{
			UserId:      md.UserId,
			MyLink:      myLink,
			ForeignLink: foreignLink,
		}}
		contactUpdates.AddUpdate(contactLink2.To_Update())

		selfUser := s.UserModel.GetUserById(deleteId, md.UserId)
		contactUpdates.AddUser(selfUser)
		sync_client.GetSyncClient().PushUpdates(deleteId, contactUpdates.ToUpdates())
	}

	////////////////////////////////////////////////////////////////////////////////////////
	contactsLink := &mtproto.TLContactsLink{Data2: &mtproto.Contacts_Link_Data{
		MyLink:      contactLink.Data2.MyLink,
		ForeignLink: contactLink.Data2.ForeignLink,
		User:        deleteUser,
	}}

	glog.Infof("contacts.deleteContact#8e953744 - reply: %s", logger.JsonDebugData(contactsLink))
	return contactsLink.To_Contacts_Link(), nil
}
