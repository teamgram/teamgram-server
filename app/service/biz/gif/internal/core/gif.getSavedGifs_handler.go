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
	"github.com/teamgram/teamgram-server/app/service/biz/gif/gif"
)

// GifGetSavedGifs
// gif.getSavedGifs user_id:long = Vector<long>;
func (c *GifCore) GifGetSavedGifs(in *gif.TLGifGetSavedGifs) (*gif.Vector_Long, error) {
	idList, _ := c.svcCtx.Dao.SavedGifsDAO.SelectAll(c.ctx, in.UserId)
	if idList == nil {
		idList = []int64{}
	}

	return &gif.Vector_Long{
		Datas: idList,
	}, nil
}
