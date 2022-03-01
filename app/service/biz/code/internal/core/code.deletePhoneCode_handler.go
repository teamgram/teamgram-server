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
	"github.com/teamgram/teamgram-server/app/service/biz/code/code"
)

// CodeDeletePhoneCode
// code.deletePhoneCode auth_key_id:long phone:string phone_code_hash:string = Bool;
func (c *CodeCore) CodeDeletePhoneCode(in *code.TLCodeDeletePhoneCode) (*mtproto.Bool, error) {
	if err := c.svcCtx.Dao.DeleteCachePhoneCode(c.ctx, in.AuthKeyId, in.Phone); err != nil {
		c.Logger.Errorf("code.deletePhoneCode - error: %v", err)
	}

	return mtproto.BoolTrue, nil
}
