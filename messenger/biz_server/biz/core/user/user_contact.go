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

package user

import (
	// "github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/mtproto"
)

const (
	CONTACT_UNKNOWN 	= 0			// 非联系人
	CONTACT_NONE 		= 1			// 非联系人
	CONTACT_HAS_PHONE 	= 2			// 非联系人有电话
	CONTACT 			= 3			// 联系人
)

func (m *UserModel) GetContactUserIDList(userId int32) []int32 {
	contactsDOList := m.dao.UserContactsDAO.SelectUserContacts(userId)
	idList := make([]int32, 0, len(contactsDOList))

	for _, do := range contactsDOList {
		idList = append(idList, do.ContactUserId)
	}
	return idList
}

func (m *UserModel) GetContactLink(userId, contactId int32) (myLink, foreignLink *mtproto.ContactLink) {
	if userId == contactId {
		myLink = mtproto.NewTLContactLinkContact().To_ContactLink()
		foreignLink = mtproto.NewTLContactLinkContact().To_ContactLink()
	} else {
		myContactDO := m.dao.UserContactsDAO.SelectByContactId(userId, contactId)
		foreignContactDO := m.dao.UserContactsDAO.SelectByContactId(contactId, userId)

		if myContactDO == nil || myContactDO.IsDeleted == 1 {
			if myContactDO == nil {
				myLink = mtproto.NewTLContactLinkNone().To_ContactLink()
			} else {
				myLink = mtproto.NewTLContactLinkHasPhone().To_ContactLink()
			}

			if foreignContactDO == nil {
				foreignLink = mtproto.NewTLContactLinkUnknown().To_ContactLink()
			} else {
				foreignLink = mtproto.NewTLContactLinkHasPhone().To_ContactLink()
			}
		} else {
			myLink = mtproto.NewTLContactLinkContact().To_ContactLink()
			if foreignContactDO == nil {
				foreignLink = mtproto.NewTLContactLinkNone().To_ContactLink()
			} else {
				if foreignContactDO.IsDeleted == 1 {
					foreignLink = mtproto.NewTLContactLinkHasPhone().To_ContactLink()
				} else {
					foreignLink = mtproto.NewTLContactLinkContact().To_ContactLink()
				}
			}
		}
	}

	return
}

