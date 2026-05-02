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

// MessagesReadHistory
// messages.readHistory#e306d3a peer:InputPeer max_id:int = messages.AffectedMessages;
func (c *MessagesCore) MessagesReadHistory(in *tg.TLMessagesReadHistory) (*tg.MessagesAffectedMessages, error) {
	md := c.MD
	if md == nil || md.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}

	peerUserID, ok := resolveUserPeerID(in.Peer, md.UserId)
	if !ok {
		return nil, tg.Err400PeerIdInvalid
	}

	var readClient readHistoryClient = c.svcCtx.Repo.MsgClient
	r, err := readClient.MsgReadHistoryV2(c.ctx, &msg.TLMsgReadHistoryV2{
		UserId:    md.UserId,
		AuthKeyId: md.PermAuthKeyId,
		PeerType:  payload.PeerTypeUser,
		PeerId:    peerUserID,
		MaxId:     in.MaxId,
	})
	if err != nil {
		c.Logger.Errorf("messages.readHistory - msg error: self_user_id: %d, peer_id: %d, max_id: %d, err: %v",
			md.UserId, peerUserID, in.MaxId, err)
		return nil, mapMsgSendError(err)
	}
	return r, nil
}
