// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// MessagesToggleNoForwards
// messages.toggleNoForwards#b11eafa2 peer:InputPeer enabled:Bool = Updates;
func (c *MessagesCore) MessagesToggleNoForwards(in *mtproto.TLMessagesToggleNoForwards) (*mtproto.Updates, error) {
	var (
		peer     = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		rUpdates *mtproto.Updates
	)

	if !peer.IsChat() {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.toggleNoForwards - error: %v", err)
		return nil, err
	}

	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatToggleNoForwards(c.ctx, &chatpb.TLChatToggleNoForwards{
		ChatId:     peer.PeerId,
		OperatorId: c.MD.UserId,
		Enabled:    in.Enabled,
	})
	if err != nil {
		c.Logger.Errorf("messages.toggleNoForwards - error: %v", err)
		return nil, err
	}

	rUpdates = mtproto.MakeUpdatesByUpdatesChats(
		[]*mtproto.Chat{chat.ToUnsafeChat(c.MD.UserId)},
		mtproto.MakeTLUpdateChat(&mtproto.Update{
			ChatId_INT64: peer.PeerId,
		}).To_Update())

	return rUpdates, nil
}
