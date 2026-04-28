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

// MessagesCheckChatInvite
// messages.checkChatInvite#3eadb1bb hash:string = ChatInvite;
func (c *ChatInvitesCore) MessagesCheckChatInvite(in *tg.TLMessagesCheckChatInvite) (*tg.ChatInvite, error) {
	if in.Hash == "" {
		return nil, tg.ErrInviteHashEmpty
	}
	if !validChatInviteHash(in.Hash) {
		return nil, tg.ErrInviteHashInvalid
	}

	selfID := selfID(c.MD)
	invite, err := c.svcCtx.Repo.ChatClient.ChatCheckChatInvite(c.ctx, &chatpb.TLChatCheckChatInvite{
		SelfId: selfID,
		Hash:   in.Hash,
	})
	if err != nil {
		return nil, mapChatError(err)
	}

	return c.projectChatInviteExt(invite, selfID)
}
