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
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// UserDeleteProfilePhotos
// user.deleteProfilePhotos user_id:long id:Vector<long> = Int64;
func (c *UserCore) UserDeleteProfilePhotos(in *user.TLUserDeleteProfilePhotos) (*mtproto.Int64, error) {
	var (
		meId = in.GetUserId()
	)

	mainPhotoId, err := c.svcCtx.Dao.UsersDAO.SelectProfilePhoto(c.ctx, meId)
	if err != nil {
		c.Logger.Errorf("user.deleteProfilePhotos - error: %v", err)
	}

	if mainPhotoId > 0 {
		if b := container2.ContainsInt64(in.GetId(), mainPhotoId); b {
			c.svcCtx.Dao.UserProfilePhotosDAO.Delete(c.ctx, meId, in.GetId())
		} else {
			nextMainId, _ := c.svcCtx.Dao.UserProfilePhotosDAO.SelectNext(c.ctx, meId, in.Id)
			sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
				c.svcCtx.Dao.UserProfilePhotosDAO.DeleteTx(tx, meId, in.Id)
				c.svcCtx.Dao.UsersDAO.UpdateProfilePhotoTx(tx, nextMainId, meId)
				//
				mainPhotoId = nextMainId
			})
		}
	} else {
		c.Logger.Errorf("user.deleteProfilePhotos - error: mainPhotoId invalid")
	}

	return &mtproto.Int64{
		V: mainPhotoId,
	}, nil
}
