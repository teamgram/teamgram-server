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
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesDeleteHistory
// messages.deleteHistory#b08f922a flags:# just_clear:flags.0?true revoke:flags.1?true peer:InputPeer max_id:int min_date:flags.2?int max_date:flags.3?int = messages.AffectedHistory;
func (c *MessagesCore) MessagesDeleteHistory(in *tg.TLMessagesDeleteHistory) (*tg.MessagesAffectedHistory, error) {
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

		return c.svcCtx.MsgClient.MsgDeleteHistory(c.ctx, &msg.TLMsgDeleteHistory{
			UserId:    userId,
			AuthKeyId: authKeyId,
			PeerType:  peer.PeerType,
			PeerId:    peer.PeerId,
			JustClear: in.JustClear,
			Revoke:    in.Revoke,
			MaxId:     in.MaxId,
		})
	}

	// Fallback placeholder when MsgClient is not available.
	return makeBffAffectedHistoryPlaceholder(in.MaxId, 1), nil
}
