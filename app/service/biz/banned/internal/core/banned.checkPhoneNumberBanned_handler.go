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
	"github.com/teamgram/teamgram-server/app/service/biz/banned/banned"
)

// BannedCheckPhoneNumberBanned
// banned.checkPhoneNumberBanned phone:string = Bool;
func (c *BannedCore) BannedCheckPhoneNumberBanned(in *banned.TLBannedCheckPhoneNumberBanned) (*mtproto.Bool, error) {
	// TODO: not impl
	do, _ := c.svcCtx.BannedDAO.CheckBannedByPhone(c.ctx, in.Phone)

	return mtproto.ToBool(do != nil), nil
}
