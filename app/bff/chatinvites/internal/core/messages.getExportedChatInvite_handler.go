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

// MessagesGetExportedChatInvite
// messages.getExportedChatInvite#73746f5c peer:InputPeer link:string = messages.ExportedChatInvite;
func (c *ChatInvitesCore) MessagesGetExportedChatInvite(in *tg.TLMessagesGetExportedChatInvite) (*tg.MessagesExportedChatInvite, error) {
	selfID := selfID(c.MD)
	peer := tg.FromInputPeer2(selfID, in.Peer)
	if peer.PeerType != tg.PEER_CHAT {
		return nil, tg.Err400PeerIdInvalid
	}

	invite, err := c.svcCtx.Repo.ChatClient.ChatGetExportedChatInvite(c.ctx, &chatpb.TLChatGetExportedChatInvite{
		ChatId: peer.PeerId,
		Link:   in.Link,
	})
	if err != nil {
		return nil, mapChatError(err)
	}

	users, err := c.fetchUserClazzes(append([]int64{selfID}, adminIDsFromInvites([]tg.ExportedChatInviteClazz{exportedInviteClazz(invite)})...), selfID)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessagesExportedChatInvite(&tg.TLMessagesExportedChatInvite{
		Invite: exportedInviteClazz(invite),
		Users:  users,
	}).ToMessagesExportedChatInvite(), nil
}
