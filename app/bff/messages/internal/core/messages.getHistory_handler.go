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
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
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

	peerUser, ok := in.Peer.(*tg.TLInputPeerUser)
	if !ok {
		return nil, tg.Err400PeerIdInvalid
	}

	var historyClient getHistoryClient = c.svcCtx.Repo.MsgClient
	r, err := historyClient.MsgGetHistory(c.ctx, &msg.TLMsgGetHistory{
		UserId:     md.UserId,
		AuthKeyId:  md.PermAuthKeyId,
		PeerType:   payload.PeerTypeUser,
		PeerId:     peerUser.UserId,
		OffsetId:   in.OffsetId,
		OffsetDate: in.OffsetDate,
		AddOffset:  in.AddOffset,
		Limit:      in.Limit,
		MaxId:      in.MaxId,
		MinId:      in.MinId,
		Hash:       in.Hash,
	})
	if err != nil {
		c.Logger.Errorf("messages.getHistory - msg error: self_user_id: %d, peer_id: %d, offset_id: %d, limit: %d, err: %v",
			md.UserId, peerUser.UserId, in.OffsetId, in.Limit, err)
		return nil, mapMsgSendError(err)
	}
	return r, nil
}
