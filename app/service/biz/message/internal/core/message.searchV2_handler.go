// Copyright 2022 Teamgram Authors
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
	"fmt"
	"math"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

/*
{
    "constructor": "CRC32_messages_search",
    "peer": {
        "predicate_name": "inputPeerChat",
        "constructor": "CRC32_inputPeerChat",
        "chat_id": "10053"
    },
    "from_id": {
        "predicate_name": "inputPeerUser",
        "constructor": "CRC32_inputPeerUser",
        "user_id": "136817689",
        "access_hash": "4630984627545308247"
    },
    "filter": {
        "predicate_name": "inputMessagesFilterEmpty",
        "constructor": "CRC32_inputMessagesFilterEmpty"
    },
    "offset_id": 94,
    "limit": 21
}
*/

// MessageSearchV2
// message.searchV2 user_id:long peer_type:int peer_id:long q:string from_id:long min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = Vector<MessageBox>;
func (c *MessageCore) MessageSearchV2(in *message.TLMessageSearchV2) (*mtproto.MessageBoxList, error) {
	if in.FromId == 0 {
		return c.MessageSearch(&message.TLMessageSearch{
			UserId:   in.UserId,
			PeerType: in.PeerType,
			PeerId:   in.PeerId,
			Q:        in.Q,
			Offset:   in.OffsetId,
			Limit:    in.Limit,
		})
	}

	var (
		offset = in.OffsetId
		//q       = in.Q
		boxList []*mtproto.MessageBox
		//limit   = in.Limit
	)

	if offset == 0 {
		offset = math.MaxInt32
	}

	dialogId := mtproto.MakeDialogId(in.UserId, in.PeerType, in.PeerId)
	c.svcCtx.Dao.MessagesDAO.SelectBackwardBySendUserIdOffsetIdLimitWithCB(
		c.ctx,
		in.UserId,
		dialogId.A,
		dialogId.B,
		in.FromId,
		math.MaxInt32,
		in.Limit,
		func(sz, i int, v *dataobject.MessagesDO) {
			boxList = append(boxList, c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, v))
		})

	if boxList == nil {
		boxList = []*mtproto.MessageBox{}
	}

	return mtproto.MakeTLMessageBoxListSlice(&mtproto.MessageBoxList{
		BoxList: boxList,
		Count:   int32(c.calcSize(in.UserId, in.FromId, dialogId)),
	}).To_MessageBoxList(), nil
}

func (c *MessageCore) calcSize(id, fromId int64, did mtproto.DialogID) int {
	where := fmt.Sprintf("user_id = %d AND (dialog_id1 = %d AND dialog_id2 = %d) AND sender_user_id = %d AND deleted = 0",
		id,
		did.A,
		did.B,
		fromId)

	return c.svcCtx.Dao.CommonDAO.CalcSizeByWhere(
		c.ctx,
		c.svcCtx.Dao.MessagesDAO.CalcTableName(id),
		where)
}
