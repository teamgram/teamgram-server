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

// MessagesGetChatInviteImporters
// messages.getChatInviteImporters#df04dd4e flags:# requested:flags.0?true subscription_expired:flags.3?true peer:InputPeer link:flags.1?string q:flags.2?string offset_date:int offset_user:InputUser limit:int = messages.ChatInviteImporters;
func (c *ChatInvitesCore) MessagesGetChatInviteImporters(in *tg.TLMessagesGetChatInviteImporters) (*tg.MessagesChatInviteImporters, error) {
	if in.Q != nil && *in.Q != "" && in.Link != nil && *in.Link != "" {
		return nil, tg.ErrSearchWithLinkNotSupported
	}

	selfID := selfID(c.MD)
	peer := tg.FromInputPeer2(selfID, in.Peer)
	if peer.PeerType != tg.PEER_CHAT {
		return nil, tg.Err400PeerIdInvalid
	}

	limit := in.Limit
	if limit == 0 {
		limit = 50
	}
	importers, err := c.svcCtx.Repo.ChatClient.ChatGetChatInviteImporters(c.ctx, &chatpb.TLChatGetChatInviteImporters{
		SelfId:     selfID,
		ChatId:     peer.PeerId,
		Requested:  in.Requested,
		Link:       in.Link,
		Q:          in.Q,
		OffsetDate: in.OffsetDate,
		OffsetUser: tg.FromInputUser(selfID, in.OffsetUser).PeerId,
		Limit:      limit,
	})
	if err != nil {
		return nil, mapChatError(err)
	}

	data := []tg.ChatInviteImporterClazz{}
	if importers != nil {
		data = importers.Datas
	}
	users, err := c.fetchUserClazzes(userIDsFromImporters(data), selfID)
	if err != nil {
		return nil, err
	}
	return tg.MakeTLMessagesChatInviteImporters(&tg.TLMessagesChatInviteImporters{
		Count:     int32(len(data)),
		Importers: data,
		Users:     users,
	}), nil
}
