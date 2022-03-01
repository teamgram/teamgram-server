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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetContact
// user.getContact user_id:long id:long = ContactData;
func (c *UserCore) UserGetContact(in *user.TLUserGetContact) (*user.ContactData, error) {
	do, _ := c.svcCtx.Dao.UserContactsDAO.SelectContact(c.ctx, in.UserId, in.Id)
	if do == nil {
		err := mtproto.ErrContactIdInvalid
		c.Logger.Errorf("user.getContact - error: %v", err)
		return nil, err
	}

	return user.MakeTLContactData(&user.ContactData{
		UserId:        in.UserId,
		ContactUserId: in.Id,
		FirstName:     mtproto.MakeFlagsString(do.ContactFirstName),
		LastName:      mtproto.MakeFlagsString(do.ContactLastName),
		MutualContact: do.Mutual,
	}).To_ContactData(), nil
}
