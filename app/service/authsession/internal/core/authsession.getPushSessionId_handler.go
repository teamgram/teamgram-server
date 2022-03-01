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

// AuthsessionGetPushSessionId
// authsession.getPushSessionId user_id:long auth_key_id:long token_type:int = Int64;
func (c *AuthsessionCore) AuthsessionGetPushSessionId(in *authsession.TLAuthsessionGetPushSessionId) (*mtproto.Int64, error) {
	sessionId := c.svcCtx.Dao.GetPushSessionId(c.ctx, in.GetUserId(), in.GetAuthKeyId(), in.GetTokenType())

	return mtproto.MakeTLInt64(&mtproto.Int64{
		V: sessionId,
	}).To_Int64(), nil
}
