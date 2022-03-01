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
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
)

// AuthsessionUnbindAuthKeyUser
// authsession.unbindAuthKeyUser auth_key_id:long user_id:long = Bool;
func (c *AuthsessionCore) AuthsessionUnbindAuthKeyUser(in *authsession.TLAuthsessionUnbindAuthKeyUser) (*mtproto.Bool, error) {
	var (
		unBindKeyId = in.AuthKeyId
	)

	if unBindKeyId != 0 {
		keyData, err := c.svcCtx.Dao.GetAuthKey(c.ctx, unBindKeyId)
		if err != nil {
			c.Logger.Errorf("session.unbindAuthKeyUser - error: %v", err)
			return nil, err
		} else if keyData == nil || keyData.PermAuthKeyId == 0 {
			err = mtproto.ErrAuthKeyInvalid
			return nil, err
		} else {
			unBindKeyId = keyData.PermAuthKeyId
		}
	}

	c.svcCtx.Dao.UnbindAuthUser(c.ctx, unBindKeyId, in.UserId)

	return mtproto.BoolTrue, nil
}
