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

var _ *tg.Bool

// MsgUpdatePinnedMessage
// msg.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
func (c *MsgCore) MsgUpdatePinnedMessage(in *msg.TLMsgUpdatePinnedMessage) (*tg.Updates, error) {
	peer := tg.PeerClazz(tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: in.PeerId}))
	switch in.PeerType {
	case tg.PEER_CHAT:
		peer = tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: in.PeerId})
	case tg.PEER_CHANNEL:
		peer = tg.MakeTLPeerChannel(&tg.TLPeerChannel{ChannelId: in.PeerId})
	case tg.PEER_SELF, tg.PEER_USER:
	default:
		return nil, tg.ErrPeerIdInvalid
	}

	return tg.MakeTLUpdateShort(&tg.TLUpdateShort{
		Update: tg.MakeTLUpdatePinnedMessages(&tg.TLUpdatePinnedMessages{
			Pinned:   !in.Unpin,
			Peer:     peer,
			Messages: []int32{in.Id},
			Pts:      1,
			PtsCount: 1,
		}),
		Date: int32(time.Now().Unix()),
	}).ToUpdates(), nil
}
