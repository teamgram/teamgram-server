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

// AuthsessionGetLangCode
// authsession.getLangCode auth_key_id:long = String;
func (c *AuthsessionCore) AuthsessionGetLangCode(in *authsession.TLAuthsessionGetLangCode) (*mtproto.String, error) {
	keyData, err := c.svcCtx.Dao.GetAuthKey(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.getLangCode - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		c.Logger.Errorf("session.getLangCode - not found keyId %d", in.GetAuthKeyId())
		return nil, fmt.Errorf("not found keyId")
	}

	langCode := c.svcCtx.Dao.GetLangCode(c.ctx, keyData.PermAuthKeyId)

	return mtproto.MakeTLString(&mtproto.String{
		V: langCode,
	}).To_String(), nil
}
