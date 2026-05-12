// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
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

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
)

// ChatCheckChatAccess
// chat.checkChatAccess#fff473b3 self_id:long chat_id:long access_kind:string = ChatAccessCheckResult;
func (c *ChatCore) ChatCheckChatAccess(in *chat.TLChatCheckChatAccess) (*chat.ChatAccessCheckResult, error) {
	if in == nil || in.SelfId <= 0 || in.ChatId <= 0 {
		return nil, chat.ErrParticipantInvalid
	}
	mChat, participant, err := c.loadMessageActionChat(in.ChatId, in.SelfId)
	if err != nil {
		return nil, err
	}
	if mChat == nil || participant == nil || !chat.IsChatMemberStateNormal(participant) {
		return nil, chat.ErrUserNotParticipant
	}
	return chat.MakeTLChatAccessCheckResult(&chat.TLChatAccessCheckResult{
		SelfId:     in.SelfId,
		ChatId:     in.ChatId,
		AccessKind: in.AccessKind,
	}).ToChatAccessCheckResult(), nil
}
