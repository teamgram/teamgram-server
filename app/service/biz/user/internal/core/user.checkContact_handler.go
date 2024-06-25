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
	"github.com/teamgram/marmota/pkg/container2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserCheckContact
// user.checkContact user_id:long id:long = Bool;
func (c *UserCore) UserCheckContact(in *user.TLUserCheckContact) (*mtproto.Bool, error) {
	cacheUserData := c.svcCtx.Dao.GetCacheUserData(c.ctx, in.GetUserId())
	//_, idList := c.svcCtx.Dao.GetUserContactIdList(c.ctx, in.GetUserId())
	isContact := container2.ContainsInt64(cacheUserData.GetContactIdList(), in.GetId())

	return mtproto.ToBool(isContact), nil
}
