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
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserGetLastSeen
// user.getLastSeen id:long = LastSeenData;
func (c *UserCore) UserGetLastSeen(in *user.TLUserGetLastSeen) (*user.LastSeenData, error) {
	do, _ := c.svcCtx.Dao.GetLastSeenAt(c.ctx, in.Id)
	if do == nil {
		err := mtproto.ErrUserIdInvalid
		c.Logger.Errorf("user.getLastSeen - error: %d", err)
		return nil, err
	}

	return user.MakeTLLastSeenData(&user.LastSeenData{
		UserId:     in.Id,
		LastSeenAt: do.LastSeenAt,
		Expires:    do.Expires,
	}).To_LastSeenData(), nil
}
