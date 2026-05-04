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

// ContactsGetStatuses
// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
func (c *ContactsCore) ContactsGetStatuses(in *tg.TLContactsGetStatuses) (*tg.VectorContactStatus, error) {
	contactList, err := c.svcCtx.Repo.UserClient.UserGetContactList(c.ctx, &userpb.TLUserGetContactList{
		UserId: c.MD.UserId,
	})
	if err != nil {
		return nil, err
	}

	var contactDatas []tg.ContactDataClazz
	if contactList != nil {
		contactDatas = contactList.Datas
	}

	idList := make([]int64, 0, len(contactDatas))
	for _, contact := range contactDatas {
		if contact != nil {
			idList = append(idList, contact.ContactUserId)
		}
	}

	lastSeenList, err := c.svcCtx.Repo.UserClient.UserGetLastSeens(c.ctx, &userpb.TLUserGetLastSeens{Id: idList})
	if err != nil {
		return nil, err
	}

	var lastSeens []userpb.LastSeenDataClazz
	if lastSeenList != nil {
		lastSeens = lastSeenList.Datas
	}

	statuses := make([]tg.ContactStatusClazz, 0, len(lastSeens))
	for _, v := range lastSeens {
		if v == nil {
			continue
		}
		statuses = append(statuses, tg.MakeTLContactStatus(&tg.TLContactStatus{
			UserId: v.UserId,
			Status: makeUserStatus(v.LastSeenAt, true),
		}).ToContactStatus())
	}
	return &tg.VectorContactStatus{Datas: statuses}, nil
}
