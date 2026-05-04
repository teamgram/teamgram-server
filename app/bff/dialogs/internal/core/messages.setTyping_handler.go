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
	"time"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesSetTyping
// messages.setTyping#58943ee2 flags:# peer:InputPeer top_msg_id:flags.0?int action:SendMessageAction = Bool;
func (c *DialogsCore) MessagesSetTyping(in *tg.TLMessagesSetTyping) (*tg.Bool, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil || in.Peer == nil || in.Action == nil {
		return nil, tg.Err400PeerIdInvalid
	}

	peer := tg.FromInputPeer2(c.MD.UserId, in.Peer)
	if peer == nil || peer.PeerType != tg.PEER_USER || peer.PeerId <= 0 {
		return nil, tg.Err400PeerIdInvalid
	}

	senderUserID := c.MD.UserId
	targetUserID := peer.PeerId
	if c.svcCtx == nil {
		c.logTypingPushFailure("missing service context", senderUserID, targetUserID, nil)
		return tg.BoolTrue, nil
	}
	if c.svcCtx.TypingLimiter != nil && !c.svcCtx.TypingLimiter.Allow(senderUserID, targetUserID, time.Now()) {
		return tg.BoolTrue, nil
	}
	if c.svcCtx.Repo == nil {
		c.logTypingPushFailure("missing repository", senderUserID, targetUserID, nil)
		return tg.BoolTrue, nil
	}
	if c.svcCtx.Repo.SyncClient == nil {
		c.logTypingPushFailure("missing sync client", senderUserID, targetUserID, nil)
		return tg.BoolTrue, nil
	}

	update := tg.MakeTLUpdateUserTyping(&tg.TLUpdateUserTyping{
		UserId: senderUserID,
		Action: in.Action,
	})
	updates := tg.MakeTLUpdateShort(&tg.TLUpdateShort{
		Update: update,
		Date:   int32(time.Now().Unix()),
	})
	if err := c.svcCtx.Repo.PushTypingUpdates(c.ctx, targetUserID, updates); err != nil {
		c.logTypingPushFailure("sync push failed", senderUserID, targetUserID, err)
		return tg.BoolTrue, nil
	}
	return tg.BoolTrue, nil
}

func (c *DialogsCore) logTypingPushFailure(reason string, senderUserID, targetUserID int64, err error) {
	if c.Logger == nil {
		return
	}
	if err != nil {
		c.Logger.Errorf("messages.setTyping - %s: sender_user_id=%d target_user_id=%d err=%v", reason, senderUserID, targetUserID, err)
		return
	}
	c.Logger.Errorf("messages.setTyping - %s: sender_user_id=%d target_user_id=%d", reason, senderUserID, targetUserID)
}
