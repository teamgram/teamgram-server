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
)

// ChatGetExportedChatInvites
// chat.getExportedChatInvites flags:# chat_id:long admin_id:long revoked:flags.3?true offset_date:flags.2?int offset_link:flags.2?string limit:int = Vector<ExportedChatInvite>;
func (c *ChatCore) ChatGetExportedChatInvites(in *chat.TLChatGetExportedChatInvites) (*chat.VectorExportedChatInvite, error) {
	invites, err := c.inviteRepository().GetExportedChatInvites(c.ctx, in.ChatId, in.AdminId, in.Revoked, in.OffsetDate, in.OffsetLink, in.Limit)
	if err != nil {
		return nil, err
	}
	return &chat.VectorExportedChatInvite{Datas: invites}, nil
}
