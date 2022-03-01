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
	"fmt"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
)

// AuthsessionBindAuthKeyUser
// authsession.bindAuthKeyUser auth_key_id:long user_id:long = Int64;
func (c *AuthsessionCore) AuthsessionBindAuthKeyUser(in *authsession.TLAuthsessionBindAuthKeyUser) (*mtproto.Int64, error) {
	keyData, err := c.svcCtx.Dao.GetAuthKey(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.bindAuthKeyUser - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	hash := c.svcCtx.Dao.BindAuthKeyUser(c.ctx, keyData.PermAuthKeyId, in.GetUserId())

	return &mtproto.Int64{V: hash}, nil
}
