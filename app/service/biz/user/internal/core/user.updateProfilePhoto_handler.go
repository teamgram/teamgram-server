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

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserUpdateProfilePhoto
// user.updateProfilePhoto user_id:long id:long = Int64;
func (c *UserCore) UserUpdateProfilePhoto(in *user.TLUserUpdateProfilePhoto) (*mtproto.Int64, error) {
	var (
		id          = in.GetId()
		meId        = in.GetUserId()
		mainPhotoId = id
	)

	if id == 0 {
		mainPhotoId, _ = c.svcCtx.Dao.UsersDAO.SelectProfilePhoto(c.ctx, meId)
		// _, _ = c.svcCtx.Dao.UserProfilePhotosDAO.SelectList(c.ctx, meId)
		if mainPhotoId > 0 {
			nextPhotoId, _ := c.svcCtx.Dao.UserProfilePhotosDAO.SelectNext(c.ctx, meId, []int64{mainPhotoId})
			sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
				c.svcCtx.Dao.UserProfilePhotosDAO.DeleteTx(tx, meId, []int64{mainPhotoId})
				c.svcCtx.Dao.UsersDAO.UpdateProfilePhotoTx(tx, nextPhotoId, meId)
			})
		}
	} else {
		sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			c.svcCtx.Dao.UserProfilePhotosDAO.InsertOrUpdateTx(tx, &dataobject.UserProfilePhotosDO{
				UserId:  meId,
				PhotoId: mainPhotoId,
				Date2:   time.Now().Unix(),
			})
			c.svcCtx.Dao.UsersDAO.UpdateProfilePhotoTx(tx, mainPhotoId, meId)
		})
	}

	return &mtproto.Int64{
		V: mainPhotoId,
	}, nil
}
