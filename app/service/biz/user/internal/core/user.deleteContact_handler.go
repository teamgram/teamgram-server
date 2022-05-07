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

// UserDeleteContact
// user.deleteContact user_id:long id:long = Bool;
func (c *UserCore) UserDeleteContact(in *user.TLUserDeleteContact) (*mtproto.Bool, error) {
	c.svcCtx.Dao.DeleteUserContact(c.ctx, in.GetUserId(), in.GetId())

	return mtproto.BoolTrue, nil
}
