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

// MessagesGetExportedChatInvites
// messages.getExportedChatInvites#a2b5a3f6 flags:# revoked:flags.3?true peer:InputPeer admin_id:InputUser offset_date:flags.2?int offset_link:flags.2?string limit:int = messages.ExportedChatInvites;
func (c *ChatInvitesCore) MessagesGetExportedChatInvites(in *tg.TLMessagesGetExportedChatInvites) (*tg.MessagesExportedChatInvites, error) {
	selfID := selfID(c.MD)
	peer := tg.FromInputPeer2(selfID, in.Peer)
	if peer.PeerType != tg.PEER_CHAT {
		return nil, tg.Err400PeerIdInvalid
	}

	adminID := tg.FromInputUser(selfID, in.AdminId).PeerId
	invites, err := c.svcCtx.Repo.ChatClient.ChatGetExportedChatInvites(c.ctx, &chatpb.TLChatGetExportedChatInvites{
		ChatId:     peer.PeerId,
		AdminId:    adminID,
		Revoked:    in.Revoked,
		OffsetDate: in.OffsetDate,
		OffsetLink: in.OffsetLink,
		Limit:      in.Limit,
	})
	if err != nil {
		return nil, mapChatError(err)
	}

	data := []tg.ExportedChatInviteClazz{}
	if invites != nil {
		data = invites.Datas
	}
	users, err := c.fetchUserClazzes(adminIDsFromInvites(data), selfID)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessagesExportedChatInvites(&tg.TLMessagesExportedChatInvites{
		Count:   int32(len(data)),
		Invites: data,
		Users:   users,
	}), nil
}
