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
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesEditChatTitle
// messages.editChatTitle#73783ffd chat_id:long title:string = Updates;
func (c *ChatsCore) MessagesEditChatTitle(in *tg.TLMessagesEditChatTitle) (*tg.Updates, error) {
	selfID := selfID(c.MD)
	mutableChat, err := c.svcCtx.Repo.ChatClient.ChatEditChatTitle(c.ctx, &chatpb.TLChatEditChatTitle{
		ChatId:     in.ChatId,
		EditUserId: selfID,
		Title:      in.Title,
	})
	if err != nil {
		return nil, mapChatError(err)
	}

	return updatesWithChat(mutableChat, selfID), nil
}
