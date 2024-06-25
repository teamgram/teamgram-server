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
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
)

// ChatHideChatJoinRequests
// chat.hideChatJoinRequests flags:# self_id:long chat_id:long approved:flags.0?true link:flags.1?string user_id:flags.2?long = RecentChatInviteRequesters;
func (c *ChatCore) ChatHideChatJoinRequests(in *chat.TLChatHideChatJoinRequests) (*chat.RecentChatInviteRequesters, error) {
	if in.GetUserId() == nil {
		// TODO: not impl
		c.Logger.Errorf("chat.hideChatJoinRequests - error: method ChatHideChatJoinRequests not impl")
		return nil, mtproto.ErrMethodNotImpl
	}

	var (
		joinId = in.GetUserId().GetValue()
	)

	requesters := chat.MakeTLRecentChatInviteRequesters(&chat.RecentChatInviteRequesters{
		RequestsPending:  0,
		RecentRequesters: []int64{},
	}).To_RecentChatInviteRequesters()

	if in.GetApproved() {
		_, err := c.ChatAddChatUser(&chat.TLChatAddChatUser{
			ChatId:    in.ChatId,
			InviterId: in.SelfId,
			UserId:    joinId,
		})
		if err != nil {
			c.Logger.Errorf("chat.importChatInvite - error: %v", err)
			return nil, err
		}
		c.svcCtx.Dao.ChatInviteParticipantsDAO.UpdateApprovedBy(c.ctx, in.SelfId, in.ChatId, joinId)
	} else {
		c.svcCtx.Dao.ChatInviteParticipantsDAO.Delete(c.ctx, in.ChatId, joinId)
	}

	c.svcCtx.Dao.ChatInviteParticipantsDAO.SelectRecentRequestedListWithCB(
		c.ctx,
		in.ChatId,
		func(sz, i int, v *dataobject.ChatInviteParticipantsDO) {
			if joinId != v.UserId {
				requesters.RequestsPending += 1
				requesters.RecentRequesters = append(requesters.RecentRequesters, v.UserId)
			}
		})

	return requesters, nil
}
