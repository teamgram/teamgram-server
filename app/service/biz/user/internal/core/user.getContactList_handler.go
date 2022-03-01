/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetContactList
// user.getContactList user_id:long = Vector<ContactData>;
func (c *UserCore) UserGetContactList(in *user.TLUserGetContactList) (*user.Vector_ContactData, error) {
	rValList := &user.Vector_ContactData{}

	c.svcCtx.UserContactsDAO.SelectUserContactsWithCB(
		c.ctx,
		in.UserId,
		func(i int, v *dataobject.UserContactsDO) {
			rValList.Datas = append(rValList.Datas, user.MakeTLContactData(&user.ContactData{
				UserId:        in.UserId,
				ContactUserId: v.ContactUserId,
				FirstName:     mtproto.MakeFlagsString(v.ContactFirstName),
				LastName:      mtproto.MakeFlagsString(v.ContactLastName),
				MutualContact: v.Mutual,
			}).To_ContactData())
		})

	return rValList, nil
}
