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

// AuthsessionSetClientSessionInfo
// authsession.setClientSessionInfo data:ClientSession = Bool;
func (c *AuthsessionCore) AuthsessionSetClientSessionInfo(in *authsession.TLAuthsessionSetClientSessionInfo) (*mtproto.Bool, error) {
	clientSession := in.GetData()
	if clientSession == nil {
		err := mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("session.setClientSessionInfo - error: %v", err)
		return nil, err
	}

	var (
		inKeyId = clientSession.GetAuthKeyId()
	)

	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, inKeyId)
	if err != nil {
		c.Logger.Errorf("queryAuthKeyV2(%d) is error: %v", inKeyId, err)
		return nil, err
	} else if keyData.PermAuthKeyId == 0 {
		c.Logger.Errorf("queryAuthKeyV2(%d) - PermAuthKeyId is empty", inKeyId)
		return nil, mtproto.ErrAuthKeyPermEmpty
	}

	clientSession.AuthKeyId = keyData.PermAuthKeyId
	err = c.svcCtx.Dao.SetClientSessionInfo(c.ctx, clientSession)
	if err != nil {
		c.Logger.Errorf("setClientSessionInfo - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
