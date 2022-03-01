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

// UserSetAccountDaysTTL
// user.setAccountDaysTTL user_id:int ttl:int = Bool;
func (c *UserCore) UserSetAccountDaysTTL(in *user.TLUserSetAccountDaysTTL) (*mtproto.Bool, error) {
	c.svcCtx.Dao.UsersDAO.UpdateAccountDaysTTL(c.ctx, in.Ttl, in.UserId)

	return mtproto.BoolTrue, nil
}
