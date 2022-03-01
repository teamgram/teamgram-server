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

// UserSetContactSignUpNotification
// user.setContactSignUpNotification user_id:long silent:Bool = Bool;
func (c *UserCore) UserSetContactSignUpNotification(in *user.TLUserSetContactSignUpNotification) (*mtproto.Bool, error) {
	var (
		k, v string
	)

	k = "contactSignUpNotification"

	if mtproto.FromBool(in.Silent) {
		v = "false"
	} else {
		v = "true"
	}

	_, _, err := c.svcCtx.Dao.UserSettingsDAO.InsertOrUpdate(c.ctx, &dataobject.UserSettingsDO{
		UserId: in.UserId,
		Key2:   k,
		Value:  v,
	})
	if err != nil {
		c.Logger.Errorf("user.setContactSignUpNotification - error: %v", err)
	}

	return mtproto.BoolTrue, nil
}
