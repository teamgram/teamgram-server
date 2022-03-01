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
	"github.com/teamgram/teamgram-server/app/service/biz/banned/internal/dal/dataobject"
)

// BannedBan
// banned.ban phone:string expires:int reason:string = Bool;
func (c *BannedCore) BannedBan(in *banned.TLBannedBan) (*mtproto.Bool, error) {
	c.svcCtx.BannedDAO.InsertOrUpdate(c.ctx, &dataobject.BannedDO{
		Phone:        in.Phone,
		BannedTime:   time.Now().Unix(),
		Expires:      int64(in.Expires),
		BannedReason: in.Reason,
		Log:          "ban",
		State:        1,
	})

	return mtproto.BoolTrue, nil
}
