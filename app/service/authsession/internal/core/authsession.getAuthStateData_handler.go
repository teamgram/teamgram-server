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

// AuthsessionGetAuthStateData
// authsession.getAuthStateData auth_key_id:long = AuthKeyStateData;
func (c *AuthsessionCore) AuthsessionGetAuthStateData(in *authsession.TLAuthsessionGetAuthStateData) (*authsession.AuthKeyStateData, error) {
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

	return authsession.MakeTLAuthKeyStateData(&authsession.AuthKeyStateData{
		AuthKeyId:            inKeyId,
		KeyState:             int32(cData.ToAuthState()),
		UserId:               cData.UserId(),
		AccessHash:           0,
		Client:               cData.GetClient(),
		AndroidPushSessionId: nil,
	}).To_AuthKeyStateData(), nil
}
