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
	"github.com/zeromicro/go-zero/core/mr"
)

// AuthsessionSetAuthKey
// authsession.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt = Bool;
func (c *AuthsessionCore) AuthsessionSetAuthKey(in *authsession.TLAuthsessionSetAuthKey) (*mtproto.Bool, error) {
	var (
		keyInfo = in.GetAuthKey()
		salt    *mtproto.TLFutureSalt
		err     error
	)

	if in.FutureSalt != nil {
		salt = in.FutureSalt.To_FutureSalt()
	}
	if salt == nil {
		err = c.svcCtx.Dao.SetAuthKeyV2(c.ctx, keyInfo, in.ExpiresIn)
	} else {
		err = mr.Finish(
			func() error {
				return c.svcCtx.Dao.SetAuthKeyV2(c.ctx, keyInfo, in.ExpiresIn)
			},
			func() error {
				return c.svcCtx.Dao.PutSaltCache(c.ctx, keyInfo.AuthKeyId, salt)
			})
	}

	if err != nil {
		c.Logger.Errorf("authsession.setAuthKey - error: %v", err)
		return mtproto.BoolFalse, nil
	}

	return mtproto.BoolTrue, nil
}
