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

// AuthsessionSetAuthKey
// authsession.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt = Bool;
func (c *AuthsessionCore) AuthsessionSetAuthKey(in *authsession.TLAuthsessionSetAuthKey) (*mtproto.Bool, error) {
	var (
		keyInfo = in.GetAuthKey()
		salt    *mtproto.TLFutureSalt
	)

	if in.FutureSalt != nil {
		salt = in.FutureSalt.To_FutureSalt()
	}

	// TODO(@benqi): add key type
	err := c.svcCtx.Dao.InsertAuthKey(
		c.ctx,
		&mtproto.AuthKeyInfo{
			AuthKeyId:          keyInfo.AuthKeyId,
			AuthKeyType:        keyInfo.AuthKeyType,
			AuthKey:            keyInfo.AuthKey,
			PermAuthKeyId:      keyInfo.PermAuthKeyId,
			TempAuthKeyId:      keyInfo.TempAuthKeyId,
			MediaTempAuthKeyId: keyInfo.MediaTempAuthKeyId,
		},
		salt,
		in.ExpiresIn)
	if err != nil {
		c.Logger.Errorf("session.setAuthKey - error: %v", err)
		return mtproto.BoolFalse, nil
	}

	return mtproto.BoolTrue, nil
}
