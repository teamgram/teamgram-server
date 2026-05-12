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
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesDeleteHistory
// messages.deleteHistory#b08f922a flags:# just_clear:flags.0?true revoke:flags.1?true peer:InputPeer max_id:int min_date:flags.2?int max_date:flags.3?int = messages.AffectedHistory;
func (c *MessagesCore) MessagesDeleteHistory(in *tg.TLMessagesDeleteHistory) (*tg.MessagesAffectedHistory, error) {
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
		action := chatpb.MessageActionDeleteLocal
		if in.Revoke {
			action = chatpb.MessageActionDeleteRevoke
		}
		if err := c.checkChatMessageAction(peer.PeerID, action, ""); err != nil {
			return nil, err
		}
	}
	var deleteClient deleteHistoryClient = c.svcCtx.Repo.MsgClient
	r, err := deleteClient.MsgDeleteHistory(c.ctx, &msg.TLMsgDeleteHistory{
		UserId:    md.UserId,
		AuthKeyId: md.PermAuthKeyId,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerID,
		JustClear: in.JustClear,
		Revoke:    in.Revoke,
		MaxId:     in.MaxId,
	})
	if err != nil {
		c.Logger.Errorf("messages.deleteHistory - msg error: self_user_id: %d, peer_type: %d, peer_id: %d, max_id: %d, err: %v", md.UserId, peer.PeerType, peer.PeerID, in.MaxId, err)
		return nil, mapMsgSendError(err)
	}
	return r, nil
}
