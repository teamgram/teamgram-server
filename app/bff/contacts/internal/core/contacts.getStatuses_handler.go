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

// ContactsGetStatuses
// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
func (c *ContactsCore) ContactsGetStatuses(in *mtproto.TLContactsGetStatuses) (*mtproto.Vector_ContactStatus, error) {
	cList, _ := c.svcCtx.Dao.UserClient.UserGetContactList(c.ctx, &userpb.TLUserGetContactList{
		UserId: c.MD.UserId,
	})

	rList := &mtproto.Vector_ContactStatus{
		Datas: make([]*mtproto.ContactStatus, 0, len(cList.GetDatas())),
	}

	idList := make([]int64, 0, len(cList.GetDatas()))
	for _, id := range cList.GetDatas() {
		idList = append(idList, id.ContactUserId)
	}

	lastSeenList, _ := c.svcCtx.Dao.UserGetLastSeens(c.ctx, &userpb.TLUserGetLastSeens{
		Id: idList,
	})

	for _, v := range lastSeenList.GetDatas() {
		rList.Datas = append(rList.Datas, mtproto.MakeTLContactStatus(&mtproto.ContactStatus{
			UserId: v.UserId,
			Status: userpb.MakeUserStatus(v.LastSeenAt, true),
		}).To_ContactStatus())
	}

	return rList, nil
}
