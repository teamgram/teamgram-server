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
	_, idList := c.svcCtx.Dao.GetUserContactIdList(c.ctx, in.GetUserId())
	isContact, _ := container2.Contains(in.GetId(), idList)

	return mtproto.ToBool(isContact), nil
}
