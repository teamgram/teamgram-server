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

// AuthsessionSetClientSessionInfo
// authsession.setClientSessionInfo data:ClientSession = Bool;
func (c *AuthsessionCore) AuthsessionSetClientSessionInfo(in *authsession.TLAuthsessionSetClientSessionInfo) (*mtproto.Bool, error) {
	clientSession := in.GetData()
	if clientSession == nil {
		err := mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("session.setClientSessionInfo - error: %v", err)
		return nil, err
	}

	keyData, err := c.svcCtx.Dao.GetAuthKey(c.ctx, clientSession.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.setClientSessionInfo - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	clientSession.AuthKeyId = keyData.PermAuthKeyId
	r := c.svcCtx.Dao.SetClientSessionInfo(c.ctx, clientSession)

	return mtproto.ToBool(r), nil
}
