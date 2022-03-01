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
	"github.com/teamgram/teamgram-server/app/service/biz/gif/gif"
	"github.com/teamgram/teamgram-server/app/service/biz/gif/internal/dal/dataobject"
)

// GifSaveGif
// gif.saveGif user_id:long gif_id:long = Bool;
func (c *GifCore) GifSaveGif(in *gif.TLGifSaveGif) (*mtproto.Bool, error) {
	c.svcCtx.Dao.SavedGifsDAO.InsertIgnore(c.ctx, &dataobject.SavedGifsDO{
		UserId: in.UserId,
		GifId:  in.GifId,
	})

	return mtproto.BoolTrue, nil
}
