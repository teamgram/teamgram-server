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
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/banned/banned"
)

// BannedUnBan
// banned.unBan phone:string = Bool;
func (c *BannedCore) BannedUnBan(in *banned.TLBannedUnBan) (*mtproto.Bool, error) {
	c.svcCtx.BannedDAO.Update(c.ctx, time.Now().Unix(), "unBan", 0, in.Phone)

	return mtproto.BoolTrue, nil
}
