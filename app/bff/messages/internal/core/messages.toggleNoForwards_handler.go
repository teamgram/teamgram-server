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

// MessagesToggleNoForwards
// messages.toggleNoForwards#b2081a35 flags:# peer:InputPeer enabled:Bool request_msg_id:flags.0?int = Updates;
func (c *MessagesCore) MessagesToggleNoForwards(in *tg.TLMessagesToggleNoForwards) (*tg.Updates, error) {
	selfID := selfID(c.MD)
	peer := tg.FromInputPeer2(selfID, in.Peer)
	if peer.PeerType != tg.PEER_CHAT {
		return nil, tg.Err400PeerIdInvalid
	}

	enabled := tg.FromBool(&tg.Bool{Clazz: in.Enabled})
	mutableChat, err := c.svcCtx.Repo.ChatClient.ChatToggleNoForwards(c.ctx, &chatpb.TLChatToggleNoForwards{
		ChatId:     peer.PeerId,
		OperatorId: selfID,
		Enabled:    tg.ToBool(enabled).Clazz,
	})
	if err != nil {
		return nil, mapChatError(err)
	}

	return updatesWithChat(mutableChat, selfID), nil
}
