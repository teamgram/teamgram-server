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

// AuthsessionGetUserId
// authsession.getUserId auth_key_id:long = Int64;
func (c *AuthsessionCore) AuthsessionGetUserId(in *authsession.TLAuthsessionGetUserId) (*mtproto.Int64, error) {
	keyData, err := c.svcCtx.Dao.GetAuthKey(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.getUserId - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	userId := c.svcCtx.Dao.GetAuthKeyUserId(c.ctx, keyData.PermAuthKeyId)

	return mtproto.MakeTLInt64(&mtproto.Int64{
		V: userId,
	}).To_Int64(), nil
}
