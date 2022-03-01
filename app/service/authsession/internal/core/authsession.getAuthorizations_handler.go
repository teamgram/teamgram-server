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

// AuthsessionGetAuthorizations
// authsession.getAuthorizations user_id:long exclude_auth_keyId:long = account.Authorizations;
func (c *AuthsessionCore) AuthsessionGetAuthorizations(in *authsession.TLAuthsessionGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	myKeyData, err := c.svcCtx.Dao.GetAuthKey(c.ctx, in.GetExcludeAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.getAuthorizations - error: %v", err)
		return nil, err
	} else if myKeyData == nil || myKeyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	authorizationList := c.svcCtx.Dao.GetAuthorizations(c.ctx, in.GetUserId(), myKeyData.PermAuthKeyId)

	return mtproto.MakeTLAccountAuthorizations(&mtproto.Account_Authorizations{
		Authorizations: authorizationList,
	}).To_Account_Authorizations(), nil
}
