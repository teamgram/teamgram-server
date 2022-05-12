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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	"math"
)

// MessageSearchV2
// message.searchV2 user_id:long peer_type:int peer_id:long q:string from_id:long min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = Vector<MessageBox>;
func (c *MessageCore) MessageSearchV2(in *message.TLMessageSearchV2) (*message.Vector_MessageBox, error) {
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
		//offset  = in.Offset
		//q       = in.Q
		boxList []*mtproto.MessageBox
		//limit   = in.Limit
	)

	dialogId := mtproto.MakeDialogId(in.UserId, in.PeerType, in.PeerId)
	c.svcCtx.Dao.MessagesDAO.SelectBackwardBySendUserIdOffsetIdLimitWithCB(
		c.ctx,
		in.UserId,
		dialogId.A,
		dialogId.B,
		in.FromId,
		math.MaxInt32,
		in.Limit,
		func(i int, v *dataobject.MessagesDO) {
			boxList = append(boxList, c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, v))
		})

	if boxList == nil {
		boxList = []*mtproto.MessageBox{}
	}

	return &message.Vector_MessageBox{
		Datas: boxList,
	}, nil
}
