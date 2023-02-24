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
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// MediaGetPhotoSizeList
// media.getPhotoSizeList size_id:long = PhotoSizeList;
func (c *MediaCore) MediaGetPhotoSizeList(in *media.TLMediaGetPhotoSizeList) (*media.PhotoSizeList, error) {
	cData, err := c.svcCtx.Dao.GetCachePhotoData(c.ctx, in.GetSizeId())
	if err != nil {
		c.Logger.Errorf("media.getPhotoSizeList - error: %v", err)
		return nil, err
	}

	return cData.ToPhotoSizeList(), nil
}
