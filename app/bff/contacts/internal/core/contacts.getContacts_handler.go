// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

/*
   private int getContactsHash(ArrayList<TLRPC.TL_contact> contacts) {
       long acc = 0;
       contacts = new ArrayList<>(contacts);
       Collections.sort(contacts, (tl_contact, tl_contact2) -> {
           if (tl_contact.user_id > tl_contact2.user_id) {
               return 1;
           } else if (tl_contact.user_id < tl_contact2.user_id) {
               return -1;
           }
           return 0;
       });
       int count = contacts.size();
       for (int a = -1; a < count; a++) {
           if (a == -1) {
               acc = ((acc * 20261) + 0x80000000L + getUserConfig().contactsSavedCount) % 0x80000000L;
           } else {
               TLRPC.TL_contact set = contacts.get(a);
               acc = ((acc * 20261) + 0x80000000L + set.user_id) % 0x80000000L;
           }
       }
       return (int) acc;
   }
*/

// ContactsGetContacts
// contacts.getContacts#5dd69e12 hash:long = contacts.Contacts;
func (c *ContactsCore) ContactsGetContacts(in *mtproto.TLContactsGetContacts) (*mtproto.Contacts_Contacts, error) {
	var (
		contacts *mtproto.Contacts_Contacts
	)

	contactList, _ := c.svcCtx.Dao.UserClient.UserGetContactList(c.ctx, &userpb.TLUserGetContactList{
		UserId: c.MD.UserId,
	})

	// 避免查询数据库时IN()条件为empty
	if len(contactList.GetDatas()) > 0 {
		idList := make([]int64, 0, len(contactList.Datas))
		for _, cV := range contactList.Datas {
			idList = append(idList, cV.ContactUserId)
		}

		c.Logger.Infof("contactIdList - {%v}", idList)

		users, _ := c.svcCtx.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
			Id: append([]int64{c.MD.UserId}, idList...),
			To: []int64{c.MD.UserId},
		})
		contacts = mtproto.MakeTLContactsContacts(&mtproto.Contacts_Contacts{
			Contacts:   contactList.ToContacts(),
			SavedCount: 0,
			Users:      users.GetUserListByIdList(c.MD.UserId, idList...),
		}).To_Contacts_Contacts()
	} else {
		contacts = mtproto.MakeTLContactsContacts(&mtproto.Contacts_Contacts{
			Contacts:   []*mtproto.Contact{},
			SavedCount: 0,
			Users:      []*mtproto.User{},
		}).To_Contacts_Contacts()
	}

	return contacts, nil
}
