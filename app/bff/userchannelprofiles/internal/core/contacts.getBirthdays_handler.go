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
	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// ContactsGetBirthdays
// contacts.getBirthdays#daeda864 = contacts.ContactBirthdays;
func (c *UserChannelProfilesCore) ContactsGetBirthdays(in *tg.TLContactsGetBirthdays) (*tg.ContactsContactBirthdays, error) {
	selfID, err := requireSelfID(c)
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if err := requireUserClient(c); err != nil {
		return nil, err
	}

	birthdays, err := c.svcCtx.Repo.UserClient.UserGetBirthdays(c.ctx, &userpb.TLUserGetBirthdays{
		UserId: selfID,
	})
	if err != nil {
		return nil, err
	}

	contacts := []tg.ContactBirthdayClazz{}
	if birthdays != nil {
		contacts = birthdays.Datas
	}
	ids := make([]int64, 0, len(contacts))
	for _, birthday := range contacts {
		if birthday != nil && birthday.ContactId > 0 {
			ids = append(ids, birthday.ContactId)
		}
	}
	if len(ids) == 0 {
		return tg.MakeTLContactsContactBirthdays(&tg.TLContactsContactBirthdays{
			Contacts: contacts,
			Users:    []tg.UserClazz{},
		}).ToContactsContactBirthdays(), nil
	}

	users, err := userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, selfID, ids, userprojection.MissingStoredReference)
	if err != nil {
		return nil, err
	}

	return tg.MakeTLContactsContactBirthdays(&tg.TLContactsContactBirthdays{
		Contacts: contacts,
		Users:    users,
	}).ToContactsContactBirthdays(), nil
}
