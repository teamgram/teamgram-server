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
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
)

// ChatGetRecentChatInviteRequesters
// chat.getRecentChatInviteRequesters self_id:long chat_id:long = RecentChatInviteRequesters;
func (c *ChatCore) ChatGetRecentChatInviteRequesters(in *chat.TLChatGetRecentChatInviteRequesters) (*chat.RecentChatInviteRequesters, error) {
	rValue := chat.MakeTLRecentChatInviteRequesters(&chat.RecentChatInviteRequesters{
		RequestsPending:  0,
		RecentRequesters: []int64{},
	}).To_RecentChatInviteRequesters()

	doList, _ := c.svcCtx.Dao.ChatInviteParticipantsDAO.SelectRecentRequestedListWithCB(
		c.ctx,
		in.GetChatId(),
		func(sz, i int, v *dataobject.ChatInviteParticipantsDO) {
			rValue.RecentRequesters = append(rValue.RecentRequesters, v.UserId)
		})

	rValue.RequestsPending = int32(len(doList))

	return rValue, nil
}
