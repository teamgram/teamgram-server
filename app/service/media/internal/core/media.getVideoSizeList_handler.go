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

// MediaGetVideoSizeList
// media.getVideoSizeList size_id = VideoSizeList;
func (c *MediaCore) MediaGetVideoSizeList(in *media.TLMediaGetVideoSizeList) (*media.VideoSizeList, error) {
	szList := c.svcCtx.Dao.GetVideoSizeList(c.ctx, in.GetSizeId())
	if szList == nil {
		szList = []*mtproto.VideoSize{}
	}
	videoSizeList := &media.VideoSizeList{
		SizeId: in.GetSizeId(),
		Sizes:  szList,
		DcId:   1,
	}

	return videoSizeList, nil
}
