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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessageGetOutboxReadDate
// message.getOutboxReadDate user_id:long peer_type:int peer_id:long msg_id:int = Vector<ReadParticipantDate>;
func (c *MessageCore) MessageGetOutboxReadDate(in *message.TLMessageGetOutboxReadDate) (*message.Vector_ReadParticipantDate, error) {
	var (
		dateList []*mtproto.ReadParticipantDate
	)

	_, err := c.svcCtx.Dao.MessageReadOutboxDAO.SelectListWithCB(
		c.ctx,
		in.UserId,
		in.PeerId,
		in.MsgId,
		func(sz, i int, v *dataobject.MessageReadOutboxDO) {
			if i == 0 {
				dateList = []*mtproto.ReadParticipantDate{
					mtproto.MakeTLReadParticipantDate(&mtproto.ReadParticipantDate{
						UserId: v.ReadUserId,
						Date:   int32(v.ReadOutboxMaxDate),
					}).To_ReadParticipantDate(),
				}
			}
		})
	if err != nil {
		c.Logger.Errorf("message.getOutboxReadDate - error: %v", err)
		return nil, err
	} else if dateList == nil {
		dateList = []*mtproto.ReadParticipantDate{}
	}

	return &message.Vector_ReadParticipantDate{
		Datas: dateList,
	}, nil
}
