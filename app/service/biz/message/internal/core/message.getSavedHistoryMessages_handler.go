// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"math"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessageGetSavedHistoryMessages
// message.getSavedHistoryMessages user_id:long peer_type:int peer_id:long offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = MessageBoxList;
func (c *MessageCore) MessageGetSavedHistoryMessages(in *message.TLMessageGetSavedHistoryMessages) (*mtproto.MessageBoxList, error) {
	var (
		selfUserId = in.UserId
		peer       = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
		addOffset  = in.AddOffset
		limit      = in.Limit
		offsetId   = in.OffsetId
		minId      = in.MinId
		maxId      = in.MaxId
		hash       = in.Hash
		boxList    []*mtproto.MessageBox
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
		// c.svcCtx.Dao.MessageClient.MessageGet
		boxList = c.svcCtx.Dao.GetOffsetIdBackwardSavedHistoryMessages(c.ctx, selfUserId, peer, offsetId, minId, maxId, addOffset+limit, hash)
	case loadTypeFirstAroundDate:
		boxList1 := c.svcCtx.GetOffsetIdForwardSavedHistoryMessages(c.ctx, selfUserId, peer, offsetId, minId, maxId, -addOffset, hash)
		for i, j := 0, len(boxList1)-1; i < j; i, j = i+1, j-1 {
			boxList1[i], boxList1[j] = boxList1[j], boxList1[i]
		}
		boxList = append(boxList, boxList1...)
		// 降序
		boxList2 := c.svcCtx.Dao.GetOffsetIdBackwardSavedHistoryMessages(c.ctx, selfUserId, peer, offsetId, minId, maxId, limit+addOffset, hash)
		// log.Infof("%v", messages2)
		boxList = append(boxList, boxList2...)
	case loadTypeForward:
		boxList = c.svcCtx.Dao.GetOffsetIdForwardSavedHistoryMessages(c.ctx, selfUserId, peer, offsetId, minId, maxId, -addOffset, hash)
		for i, j := 0, len(boxList)-1; i < j; i, j = i+1, j-1 {
			boxList[i], boxList[j] = boxList[j], boxList[i]
		}
	}

	count := c.svcCtx.Dao.CommonDAO.CalcSize(
		c.ctx,
		c.svcCtx.Dao.MessagesDAO.CalcTableName(selfUserId),
		map[string]interface{}{
			"user_id":         selfUserId,
			"saved_peer_type": peer.PeerType,
			"saved_peer_id":   peer.PeerId,
			"deleted":         0,
		},
	)

	return mtproto.MakeTLMessageBoxListSlice(&mtproto.MessageBoxList{
		BoxList: boxList,
		Count:   int32(count),
	}).To_MessageBoxList(), nil
}
