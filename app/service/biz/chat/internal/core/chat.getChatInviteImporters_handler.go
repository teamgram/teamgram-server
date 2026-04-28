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

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/chat/internal/repository"
)

// ChatGetChatInviteImporters
// chat.getChatInviteImporters flags:# self_id:long chat_id:long requested:flags.0?true link:flags.1?string q:flags.2?string offset_date:int offset_user:long limit:int = Vector<ChatInviteImporter>;
func (c *ChatCore) ChatGetChatInviteImporters(in *chat.TLChatGetChatInviteImporters) (*chat.VectorChatInviteImporter, error) {
	if _, err := c.requireCanInvite(in.ChatId, in.SelfId); err != nil {
		return nil, err
	}
	link := ""
	if in.Link != nil {
		link = *in.Link
	}
	importers, err := c.inviteRepository().GetChatInviteImporters(c.ctx, repository.ChatInviteImporterQuery{
		ChatID:     in.ChatId,
		Link:       link,
		Requested:  in.Requested,
		OffsetDate: in.OffsetDate,
		OffsetUser: in.OffsetUser,
		Limit:      in.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &chat.VectorChatInviteImporter{Datas: importers}, nil
}
