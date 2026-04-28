// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"

// ChatGetAdminsWithInvites
// chat.getAdminsWithInvites self_id:long chat_id:long = Vector<ChatAdminWithInvites>;
func (c *ChatCore) ChatGetAdminsWithInvites(in *chat.TLChatGetAdminsWithInvites) (*chat.VectorChatAdminWithInvites, error) {
	mChat, err := c.requireCanInvite(in.ChatId, in.SelfId)
	if err != nil {
		return nil, err
	}

	adminIDs := make([]int64, 0, len(mChat.ChatParticipants))
	for _, participant := range mChat.ChatParticipants {
		if participant == nil || !chat.CanInviteUsers(participant) {
			continue
		}
		adminIDs = append(adminIDs, participant.UserId)
	}
	out, err := c.inviteRepository().GetAdminsWithInvites(c.ctx, in.ChatId, adminIDs)
	if err != nil {
		return nil, err
	}
	return &chat.VectorChatAdminWithInvites{Datas: out}, nil
}
