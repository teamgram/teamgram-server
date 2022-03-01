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
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	"math"
)

// MessageGetUnreadMentions
// message.getUnreadMentions user_id:long peer_type:int peer_id:long offset_id:int add_offset:int limit:int min_id:int max_int:int = Vector<MessageBox>;
func (c *MessageCore) MessageGetUnreadMentions(in *message.TLMessageGetUnreadMentions) (*message.Vector_MessageBox, error) {
	var (
		addOffset = in.AddOffset
		limit     = in.Limit
		offsetId  = in.OffsetId
		minId     = in.MinId
		maxId     = in.MaxInt
		boxList   []*mtproto.MessageBox
		peer      = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
	)

	loadType := loadTypeBackward
	if addOffset >= 0 {
		loadType = loadTypeBackward
	} else if addOffset+limit > 0 {
		loadType = loadTypeFirstAroundDate
	} else {
		loadType = loadTypeForward
	}

	if offsetId == 0 {
		offsetId = math.MaxInt32
	}

	switch loadType {
	case loadTypeBackward:
		if offsetId == 0 {
			offsetId = math.MaxInt32
		}
		boxList = c.svcCtx.Dao.GetOffsetIdBackwardUnreadMentions(
			c.ctx,
			in.UserId,
			peer,
			offsetId,
			minId,
			maxId,
			addOffset+limit)
	case loadTypeFirstAroundDate:
		boxList1 := c.svcCtx.Dao.GetOffsetIdForwardUnreadMentions(
			c.ctx,
			in.UserId,
			peer,
			offsetId,
			minId,
			maxId,
			-addOffset)

		for i, j := 0, len(boxList1)-1; i < j; i, j = i+1, j-1 {
			boxList1[i], boxList1[j] = boxList1[j], boxList1[i]
		}
		boxList = append(boxList, boxList1...)
		// 降序
		boxList2 := c.svcCtx.Dao.GetOffsetIdBackwardUnreadMentions(
			c.ctx,
			in.UserId,
			peer,
			offsetId,
			minId,
			maxId,
			limit+addOffset)

		boxList = append(boxList, boxList2...)
	case loadTypeForward:
		boxList = c.svcCtx.Dao.GetOffsetIdForwardUnreadMentions(
			c.ctx,
			in.UserId,
			peer,
			offsetId,
			minId,
			maxId,
			-addOffset)
		for i, j := 0, len(boxList)-1; i < j; i, j = i+1, j-1 {
			boxList[i], boxList[j] = boxList[j], boxList[i]
		}
	}

	return &message.Vector_MessageBox{
		Datas: boxList,
	}, nil
}
