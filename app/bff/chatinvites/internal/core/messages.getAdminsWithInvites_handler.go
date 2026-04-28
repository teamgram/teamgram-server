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

// MessagesGetAdminsWithInvites
// messages.getAdminsWithInvites#3920e6ef peer:InputPeer = messages.ChatAdminsWithInvites;
func (c *ChatInvitesCore) MessagesGetAdminsWithInvites(in *tg.TLMessagesGetAdminsWithInvites) (*tg.MessagesChatAdminsWithInvites, error) {
	selfID := selfID(c.MD)
	peer := tg.FromInputPeer2(selfID, in.Peer)
	if peer.PeerType != tg.PEER_CHAT {
		return nil, tg.Err400PeerIdInvalid
	}

	admins, err := c.svcCtx.Repo.ChatClient.ChatGetAdminsWithInvites(c.ctx, &chatpb.TLChatGetAdminsWithInvites{
		SelfId: selfID,
		ChatId: peer.PeerId,
	})
	if err != nil {
		return nil, mapChatError(err)
	}

	data := []tg.ChatAdminWithInvitesClazz{}
	if admins != nil {
		data = admins.Datas
	}
	users, err := c.fetchUserClazzes(adminIDsFromAdmins(data), selfID)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessagesChatAdminsWithInvites(&tg.TLMessagesChatAdminsWithInvites{
		Admins: data,
		Users:  users,
	}), nil
}
