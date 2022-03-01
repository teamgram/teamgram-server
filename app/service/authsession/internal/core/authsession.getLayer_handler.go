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

// AuthsessionGetLayer
// authsession.getLayer auth_key_id:long = Int32;
func (c *AuthsessionCore) AuthsessionGetLayer(in *authsession.TLAuthsessionGetLayer) (*mtproto.Int32, error) {
	keyData, err := c.svcCtx.Dao.GetAuthKey(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.getUserId - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	layer := c.svcCtx.GetApiLayer(c.ctx, keyData.PermAuthKeyId)

	return mtproto.MakeTLInt32(&mtproto.Int32{
		V: layer,
	}).To_Int32(), nil
}
