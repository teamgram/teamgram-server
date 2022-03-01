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

// AuthsessionGetLangPack
// authsession.getLangPack auth_key_id:long = String;
func (c *AuthsessionCore) AuthsessionGetLangPack(in *authsession.TLAuthsessionGetLangPack) (*mtproto.String, error) {
	keyData, err := c.svcCtx.Dao.GetAuthKey(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.getLangPack - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	langPack := c.svcCtx.GetLangPack(c.ctx, keyData.PermAuthKeyId)

	return mtproto.MakeTLString(&mtproto.String{
		V: langPack,
	}).To_String(), nil
}
