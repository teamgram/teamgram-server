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

// MessagesUpdatePinnedMessage
// messages.updatePinnedMessage#d2aaf7ec flags:# silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer:InputPeer id:int = Updates;
func (c *MessagesCore) MessagesUpdatePinnedMessage(in *tg.TLMessagesUpdatePinnedMessage) (*tg.Updates, error) {
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
		action := chatpb.MessageActionPinMessage
		if in.Unpin {
			action = chatpb.MessageActionUnpinAll
		}
		if err := c.checkChatMessageAction(peer.PeerID, action, ""); err != nil {
			return nil, err
		}
	}
	var pinClient updatePinnedMessageClient = c.svcCtx.Repo.MsgClient
	r, err := pinClient.MsgUpdatePinnedMessage(c.ctx, &msg.TLMsgUpdatePinnedMessage{
		UserId:    md.UserId,
		AuthKeyId: md.PermAuthKeyId,
		Silent:    in.Silent,
		Unpin:     in.Unpin,
		PmOneside: in.PmOneside,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerID,
		Id:        in.Id,
	})
	if err != nil {
		c.Logger.Errorf("messages.updatePinnedMessage - msg error: self_user_id: %d, peer_type: %d, peer_id: %d, id: %d, unpin: %t, err: %v",
			md.UserId, peer.PeerType, peer.PeerID, in.Id, in.Unpin, err)
		return nil, mapMsgSendError(err)
	}
	return r, nil
}
