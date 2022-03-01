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

// UserCheckContact
// user.checkContact user_id:long id:long = Bool;
func (c *UserCore) UserCheckContact(in *user.TLUserCheckContact) (*mtproto.Bool, error) {
	contact, _ := c.UserGetContact(&user.TLUserGetContact{
		UserId: in.UserId,
		Id:     in.Id,
	})

	return mtproto.ToBool(contact != nil), nil
}
