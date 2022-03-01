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
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserDeleteContact
// user.deleteContact user_id:long id:long = Bool;
func (c *UserCore) UserDeleteContact(in *user.TLUserDeleteContact) (*mtproto.Bool, error) {
	// A 删除 B
	// 如果AB is mutual，则BA设置为非mutual

	sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		c.svcCtx.Dao.UserContactsDAO.DeleteContactsTx(tx, in.UserId, []int64{in.Id})
		c.svcCtx.Dao.UserContactsDAO.UpdateMutualTx(tx, false, in.Id, in.UserId)
	})

	return mtproto.ToBool(true), nil

}
