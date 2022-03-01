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
	"github.com/teamgram/teamgram-server/app/service/biz/auth/auth"
)

// AuthExportLoginToken
// auth.exportLoginToken api_id:int api_hash:string except_ids:Vector<long> = auth.LoginToken;
func (c *AuthCore) AuthExportLoginToken(in *auth.TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error) {
	// TODO: not impl
	c.Logger.Errorf("auth.exportLoginToken - error: method AuthExportLoginToken not impl")

	return nil, mtproto.ErrMethodNotImpl
}
