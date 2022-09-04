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

// UserUpdateLastSeen
// user.updateLastSeen id:long last_seen_at:long expires:int = Bool;
func (c *UserCore) UserUpdateLastSeen(in *user.TLUserUpdateLastSeen) (*mtproto.Bool, error) {
	if in.GetId() > 0 {
		c.svcCtx.Dao.PutLastSeenAt(
			c.ctx,
			in.GetId(),
			in.GetLastSeenAt(),
			in.GetExpires())
	}

	return mtproto.BoolTrue, nil
}
