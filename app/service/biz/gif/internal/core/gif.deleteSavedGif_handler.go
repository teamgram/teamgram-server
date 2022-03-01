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
)

// GifDeleteSavedGif
// gif.deleteSavedGif user_id:long gif_id:long = Bool;
func (c *GifCore) GifDeleteSavedGif(in *gif.TLGifDeleteSavedGif) (*mtproto.Bool, error) {
	c.svcCtx.Dao.SavedGifsDAO.Delete(c.ctx, in.UserId, in.GifId)

	return mtproto.BoolTrue, nil
}
