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
	"github.com/teamgram/teamgram-server/app/service/media/media"
)

// MediaGetPhotoSizeList
// media.getPhotoSizeList size_id:long = PhotoSizeList;
func (c *MediaCore) MediaGetPhotoSizeList(in *media.TLMediaGetPhotoSizeList) (*media.PhotoSizeList, error) {
	szList := c.svcCtx.Dao.GetPhotoSizeListV2(c.ctx, in.GetSizeId())
	if szList == nil {
		szList = []*mtproto.PhotoSize{}
	}

	return &media.PhotoSizeList{
		SizeId: in.GetSizeId(),
		Sizes:  szList,
		DcId:   1,
	}, nil
}
