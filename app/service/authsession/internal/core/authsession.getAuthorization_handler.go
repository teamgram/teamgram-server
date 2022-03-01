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

// AuthsessionGetAuthorization
// authsession.getAuthorization auth_key_id:long = Authorization;
func (c *AuthsessionCore) AuthsessionGetAuthorization(in *authsession.TLAuthsessionGetAuthorization) (*mtproto.Authorization, error) {
	myKeyData, err := c.svcCtx.Dao.GetAuthKey(c.ctx, in.AuthKeyId)
	if err != nil {
		c.Logger.Errorf("session.getAuthorization - error: %v", err)
		return nil, err
	} else if myKeyData == nil || myKeyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	authorization, err := c.svcCtx.Dao.GetAuthorization(c.ctx, myKeyData.PermAuthKeyId)
	if err != nil {
		c.Logger.Errorf("session.getAuthorization - error: %v", err)
		return nil, err
	}

	return authorization, nil
}
