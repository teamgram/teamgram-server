// Copyright 2024 Teamgram Authors
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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// ContactsGetBirthdays
// contacts.getBirthdays#daeda864 = contacts.ContactBirthdays;
func (c *UserProfileCore) ContactsGetBirthdays(in *mtproto.TLContactsGetBirthdays) (*mtproto.Contacts_ContactBirthdays, error) {
	cBirthdays, err := c.svcCtx.Dao.UserClient.UserGetBirthdays(c.ctx, &user.TLUserGetBirthdays{
		UserId: c.MD.UserId,
	})

	if err != nil {
		c.Logger.Errorf("contacts.getBirthdays - error: %v", err)
		return nil, err
	}

	var (
		idList []int64

		rV = mtproto.MakeTLContactsContactBirthdays(&mtproto.Contacts_ContactBirthdays{
			Contacts: []*mtproto.ContactBirthday{},
			Users:    []*mtproto.User{},
		}).To_Contacts_ContactBirthdays()
	)

	for _, c := range cBirthdays.GetDatas() {
		idList = append(idList, c.ContactId)
	}
	if len(idList) == 0 {
		return rV, nil
	}

	cUsers, err := c.svcCtx.Dao.UserClient.UserGetMutableUsersV2(c.ctx, &user.TLUserGetMutableUsersV2{
		Id: idList,
	})
	if err != nil {
		c.Logger.Errorf("contacts.getBirthdays - error: %v", err)
	} else {
		rV.Contacts = cBirthdays.GetDatas()
		rV.Users = cUsers.GetUserListByIdList(c.MD.UserId, idList...)
	}

	return rV, nil
}
