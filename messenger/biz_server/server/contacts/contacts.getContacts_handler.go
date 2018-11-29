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
	"golang.org/x/net/context"
)

// contacts.getContacts#c023849f hash:int = contacts.Contacts;
func (s *ContactsServiceImpl) ContactsGetContacts(ctx context.Context, request *mtproto.TLContactsGetContacts) (*mtproto.Contacts_Contacts, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("contacts.getContacts#c023849f - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		contacts *mtproto.Contacts_Contacts
	)
	contactLogic := s.ContactModel.MakeContactLogic(md.UserId)

	contactList := contactLogic.GetContactList()
	// 避免查询数据库时IN()条件为empty
	if len(contactList) > 0 {
		idList := make([]int32, 0, len(contactList))
		cList := make([]*mtproto.Contact, 0, len(contactList))
		for _, c := range contactList {
			idList = append(idList, c.ContactUserId)
			c2 := &mtproto.Contact{
				Constructor: mtproto.TLConstructor_CRC32_contact,
				Data2: &mtproto.Contact_Data{
					UserId: c.ContactUserId,
					Mutual: mtproto.ToBool(c.Mutual == 1),
				},
			}
			cList = append(cList, c2)
		}

		glog.Infof("contactIdList - {%v}", idList)

		users := s.UserModel.GetUserListByIdList(md.UserId, idList)
		contacts = &mtproto.Contacts_Contacts{
			Constructor: mtproto.TLConstructor_CRC32_contacts_contacts,
			Data2: &mtproto.Contacts_Contacts_Data{
				Contacts:   cList,
				SavedCount: 0,
				Users:      users,
			},
		}
	} else {
		contacts = mtproto.NewTLContactsContacts().To_Contacts_Contacts()
	}

	glog.Infof("contacts.getContacts#c023849f - reply: %s\n", logger.JsonDebugData(contacts))
	return contacts, nil
}
