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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	updates2 "github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/update"
	"golang.org/x/net/context"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
)

// Android client有三种场景会调用importContacts:
// 1. 导入通信录
// 2. 导入单个联系人
// contacts.importContacts#2c800be5 contacts:Vector<InputContact> = contacts.ImportedContacts;
func (s *ContactsServiceImpl) ContactsImportContacts(ctx context.Context, request *mtproto.TLContactsImportContacts) (*mtproto.Contacts_ImportedContacts, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.importContacts#2c800be5 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		importedContacts *mtproto.TLContactsImportedContacts
		region 			 string
	)

	if len(request.Contacts) == 0 {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err, ": contacts empty")
		return nil, err
	}

	importedContacts = &mtproto.TLContactsImportedContacts{ Data2: &mtproto.Contacts_ImportedContacts_Data{
		Imported:       []*mtproto.ImportedContact{},
		PopularInvites: []*mtproto.PopularContact{},
		RetryContacts:  []int64{},
	}}

	importList := make([]*mtproto.InputContact_Data, 0, len(request.Contacts))
	importPhoneList := make([]string, 0, len(request.Contacts))
	for _, c2 := range request.GetContacts() {
		c := c2.To_InputPhoneContact()
		// if c.GetPhone() == ""
		// ignore first_name and last_name is emtpy
		if c.GetFirstName() == "" && c.GetLastName() == "" {
			continue
		}
		phone := c.GetPhone()
		pnumber, err := base.MakePhoneNumberUtil(phone, "")
		if err != nil {
			if region == "" {
				region = s.UserModel.GetCountryCodeByUser(md.UserId)
			}
			pnumber, err = base.MakePhoneNumberUtil(phone, region)
			if err != nil {
				continue
			}
		}
		phone = pnumber.GetNormalizeDigits()
		c.SetPhone(phone)
		importPhoneList = append(importPhoneList, phone)
		importList = append(importList, c.GetData2())
	}

	if len(importList) == 0 {
		glog.Warning("")
		//importedContacts = &mtproto.TLContactsImportedContacts{ Data2: &mtproto.Contacts_ImportedContacts_Data{
		//	Imported:       []*mtproto.ImportedContact{},
		//	PopularInvites: []*mtproto.PopularContact{},
		//	RetryContacts:  []int64{},
		//}}
		return importedContacts.To_Contacts_ImportedContacts(), nil
	}

	// TODO(@benqi): retry_contacts - 要求客户端过段时间重试，优化时再来考虑

	// save phone books
	s.ContactModel.BackupPhoneBooks(md.AuthId, importList)

	contactLogic := s.ContactModel.MakeContactLogic(md.UserId)
	importeds, popularInvites, updList := contactLogic.ImportContacts(importList)

	if len(importeds) > 0 {
		selfUpdates := updates2.NewUpdatesLogic(md.UserId)

		var importedUsers []*mtproto.User

		importedIdList := make([]int32, 0, len(importeds))
		for _, i := range importeds {
			importedIdList = append(importedIdList, i.GetData2().GetUserId())
		}

		importedUsers = s.UserModel.GetUserListByIdList(md.UserId, importedIdList)
		importedContacts.SetUsers(importedUsers)

		// sync
		for _, i := range importeds {
			myLink, foreignLink := s.UserModel.GetContactLink(md.UserId, i.GetData2().GetUserId())
			contactLink := &mtproto.TLUpdateContactLink{Data2: &mtproto.Update_Data{
				UserId:      i.GetData2().GetUserId(),
				MyLink:      myLink,
				ForeignLink: foreignLink,
			}}
			selfUpdates.AddUpdate(contactLink.To_Update())
		}
		selfUpdates.AddUsers(importedUsers)
		sync_client.GetSyncClient().SyncUpdatesNotMe(md.UserId, md.AuthId, selfUpdates.ToUpdates())

		// push
		for _, updId := range updList {
			contactUpdates := updates2.NewUpdatesLogic(updId)
			myLink, foreignLink := s.UserModel.GetContactLink(updId, md.UserId)
			contactLink2 := &mtproto.TLUpdateContactLink{Data2: &mtproto.Update_Data{
				UserId:      md.UserId,
				MyLink:      myLink,
				ForeignLink: foreignLink,
			}}
			contactUpdates.AddUpdate(contactLink2.To_Update())

			myUser := s.UserModel.GetUserById(updId, md.UserId)
			contactUpdates.AddUser(myUser)
			sync_client.GetSyncClient().PushUpdates(updId, contactUpdates.ToUpdates())
		}
	}

	importedContacts.SetImported(importeds)
	importedContacts.SetPopularInvites(popularInvites)

	glog.Infof("contacts.importContacts#2c800be5 - reply: {%s}", logger.JsonDebugData(importedContacts))
	return importedContacts.To_Contacts_ImportedContacts(), nil
}
