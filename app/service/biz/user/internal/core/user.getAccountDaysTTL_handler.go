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

// UserGetAccountDaysTTL
// user.getAccountDaysTTL user_id:int = AccountDaysTTL;
func (c *UserCore) UserGetAccountDaysTTL(in *user.TLUserGetAccountDaysTTL) (*mtproto.AccountDaysTTL, error) {
	userDO, _ := c.svcCtx.Dao.UsersDAO.SelectAccountDaysTTL(c.ctx, in.UserId)
	if userDO == nil {
		err := mtproto.ErrUserIdInvalid
		c.Logger.Errorf("user.getAccountDaysTTL - error: %v", err)
		return nil, err
	}

	return mtproto.MakeTLAccountDaysTTL(&mtproto.AccountDaysTTL{
		Days: userDO.AccountDaysTtl,
	}).To_AccountDaysTTL(), nil
}
