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

// AuthsessionGetLangCode
// authsession.getLangCode auth_key_id:long = String;
func (c *AuthsessionCore) AuthsessionGetLangCode(in *authsession.TLAuthsessionGetLangCode) (*mtproto.String, error) {
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

	langCode := c.svcCtx.Dao.GetLangCode(c.ctx, keyData.PermAuthKeyId)

	return mtproto.MakeTLString(&mtproto.String{
		V: langCode,
	}).To_String(), nil
}
