// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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
	"time"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesUpdatePinnedMessage
// messages.updatePinnedMessage#d2aaf7ec flags:# silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer:InputPeer id:int = Updates;
func (c *MessagesCore) MessagesUpdatePinnedMessage(in *tg.TLMessagesUpdatePinnedMessage) (*tg.Updates, error) {
	var userId int64
	if c.MD != nil {
		userId = c.MD.UserId
	}

	peer := tg.FromInputPeer2(userId, in.Peer)

	switch peer.PeerType {
	case tg.PEER_SELF, tg.PEER_USER, tg.PEER_CHAT:
	case tg.PEER_CHANNEL:
		return nil, tg.ErrEnterpriseIsBlocked
	default:
		return nil, tg.ErrPeerIdInvalid
	}

	// When MsgClient is wired, delegate to msg service.
	if c.svcCtx != nil && c.svcCtx.MsgClient != nil {
		var authKeyId int64
		if c.MD != nil {
			authKeyId = c.MD.AuthId
		}

		return c.svcCtx.MsgClient.MsgUpdatePinnedMessage(c.ctx, &msg.TLMsgUpdatePinnedMessage{
			Silent:    in.Silent,
			Unpin:     in.Unpin,
			PmOneside: in.PmOneside,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			Id:        in.Id,
			UserId:    userId,
			AuthKeyId: authKeyId,
		})
	}

	// Fallback placeholder when MsgClient is not available.
	peerClazz, _ := bffPeerFromInput(c, in.Peer)
	return tg.MakeTLUpdateShort(&tg.TLUpdateShort{
		Update: tg.MakeTLUpdatePinnedMessages(&tg.TLUpdatePinnedMessages{
			Pinned:   !in.Unpin,
			Peer:     peerClazz,
			Messages: []int32{in.Id},
			Pts:      1,
			PtsCount: 1,
		}),
		Date: int32(time.Now().Unix()),
	}).ToUpdates(), nil
}
