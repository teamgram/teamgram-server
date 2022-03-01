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

// UserGetContentSettings
// user.getContentSettings user_id:long = account.ContentSettings;
func (c *UserCore) UserGetContentSettings(in *user.TLUserGetContentSettings) (*mtproto.Account_ContentSettings, error) {
	var (
		rV = false
	)

	if do, _ := c.svcCtx.Dao.UserSettingsDAO.SelectByKey(c.ctx, in.UserId, "sensitive_enabled"); do != nil {
		if do.Value == "true" {
			rV = true
		}
	}

	return mtproto.MakeTLAccountContentSettings(&mtproto.Account_ContentSettings{
		SensitiveEnabled:   rV,
		SensitiveCanChange: true,
	}).To_Account_ContentSettings(), nil
}
