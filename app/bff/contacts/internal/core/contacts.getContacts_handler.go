// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// ContactsGetContacts
// contacts.getContacts#5dd69e12 hash:long = contacts.Contacts;
func (c *ContactsCore) ContactsGetContacts(in *tg.TLContactsGetContacts) (*tg.ContactsContacts, error) {
	contactList, err := c.svcCtx.Repo.UserClient.UserGetContactList(c.ctx, &userpb.TLUserGetContactList{
		UserId: c.MD.UserId,
	})
	if err != nil {
		return nil, err
	}
	if contactList == nil || len(contactList.Datas) == 0 {
		return tg.MakeTLContactsContacts(&tg.TLContactsContacts{
			Contacts:   []tg.ContactClazz{},
			SavedCount: 0,
			Users:      []tg.UserClazz{},
		}).ToContactsContacts(), nil
	}

	idList := make([]int64, 0, len(contactList.Datas))
	for _, contact := range contactList.Datas {
		if contact != nil {
			idList = append(idList, contact.ContactUserId)
		}
	}

	users, err := c.svcCtx.Repo.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: append([]int64{c.MD.UserId}, idList...),
		To: []int64{c.MD.UserId},
	})
	if err != nil {
		return nil, err
	}

	var immutableUsers []tg.ImmutableUserClazz
	if users != nil {
		immutableUsers = users.Datas
	}

	return tg.MakeTLContactsContacts(&tg.TLContactsContacts{
		Contacts:   contactDatasToContacts(contactList.Datas),
		SavedCount: 0,
		Users:      projectUsersByIDs(immutableUsers, idList),
	}).ToContactsContacts(), nil
}
