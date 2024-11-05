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
	"errors"

	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/authsession/authsession"
)

// AuthsessionGetPushSessionId
// authsession.getPushSessionId user_id:long auth_key_id:long token_type:int = Int64;
func (c *AuthsessionCore) AuthsessionGetPushSessionId(in *authsession.TLAuthsessionGetPushSessionId) (*mtproto.Int64, error) {
	var (
		inKeyId = in.GetAuthKeyId()
	)

	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, inKeyId)
	if err != nil {
		c.Logger.Errorf("queryAuthKeyV2(%d) is error: %v", inKeyId, err)
		return nil, err
	} else if keyData.PermAuthKeyId == 0 {
		c.Logger.Errorf("queryAuthKeyV2(%d) - PermAuthKeyId is empty", inKeyId)
		return nil, mtproto.ErrAuthKeyPermEmpty
	}

	cData, err := c.svcCtx.GetCacheAuthData(c.ctx, keyData.PermAuthKeyId)
	if err != nil {
		if !errors.Is(err, sqlc.ErrNotFound) {
			c.Logger.Errorf("authsession.getAuthStateData - error: %v", err)
			return nil, err
		}
	}

	return mtproto.MakeTLInt64(&mtproto.Int64{
		V: cData.AndroidPushSessionId(),
	}).To_Int64(), nil
}
