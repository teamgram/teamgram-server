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
	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetHistory
// messages.getHistory#4423e6c5 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (c *MessagesCore) MessagesGetHistory(in *tg.TLMessagesGetHistory) (*tg.MessagesMessages, error) {
	md := c.MD
	if md == nil || md.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}

	peer, ok := resolveMessagePeer(in.Peer, md.UserId)
	if !ok {
		return nil, tg.Err400PeerIdInvalid
	}
	if peer.PeerType == payload.PeerTypeChat {
		if err := c.checkChatAccess(peer.PeerID, chatpb.ChatAccessGetHistory); err != nil {
			return nil, err
		}
	}

	var historyClient getHistoryClient = c.svcCtx.Repo.MsgClient
	r, err := historyClient.MsgGetHistory(c.ctx, &msg.TLMsgGetHistory{
		UserId:     md.UserId,
		AuthKeyId:  md.PermAuthKeyId,
		PeerType:   peer.PeerType,
		PeerId:     peer.PeerID,
		OffsetId:   in.OffsetId,
		OffsetDate: in.OffsetDate,
		AddOffset:  in.AddOffset,
		Limit:      in.Limit,
		MaxId:      in.MaxId,
		MinId:      in.MinId,
		Hash:       in.Hash,
	})
	if err != nil {
		c.Logger.Errorf("messages.getHistory - msg error: self_user_id: %d, peer_type: %d, peer_id: %d, offset_id: %d, limit: %d, err: %v",
			md.UserId, peer.PeerType, peer.PeerID, in.OffsetId, in.Limit, err)
		return nil, mapMsgSendError(err)
	}
	if err = userprojection.FillMessagesMessagesUsers(c.ctx, c.svcCtx.Repo.UserClient, md.UserId, r, userprojection.MissingStoredReference); err != nil {
		return nil, err
	}
	return r, nil
}
