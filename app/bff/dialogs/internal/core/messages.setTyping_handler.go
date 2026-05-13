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

	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
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
	if peer == nil || peer.PeerId <= 0 {
		return nil, tg.Err400PeerIdInvalid
	}

	senderUserID := c.MD.UserId
	switch peer.PeerType {
	case tg.PEER_USER:
		return c.pushUserTyping(senderUserID, peer.PeerId, in.Action)
	case tg.PEER_CHAT:
		return c.pushChatTyping(senderUserID, peer.PeerId, in.Action)
	default:
		return nil, tg.Err400PeerIdInvalid
	}
}

func (c *DialogsCore) pushUserTyping(senderUserID, targetUserID int64, action tg.SendMessageActionClazz) (*tg.Bool, error) {
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
		Action: action,
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

func (c *DialogsCore) pushChatTyping(senderUserID, chatID int64, action tg.SendMessageActionClazz) (*tg.Bool, error) {
	if c.svcCtx == nil {
		c.logTypingPushFailure("missing service context", senderUserID, chatID, nil)
		return tg.BoolTrue, nil
	}
	if c.svcCtx.TypingLimiter != nil && !c.svcCtx.TypingLimiter.Allow(senderUserID, chatID, time.Now()) {
		return tg.BoolTrue, nil
	}
	if c.svcCtx.Repo == nil {
		c.logTypingPushFailure("missing repository", senderUserID, chatID, nil)
		return tg.BoolTrue, nil
	}
	if c.svcCtx.Repo.ChatClient == nil {
		c.logTypingPushFailure("missing chat client", senderUserID, chatID, nil)
		return tg.BoolTrue, nil
	}
	if c.svcCtx.Repo.SyncClient == nil {
		c.logTypingPushFailure("missing sync client", senderUserID, chatID, nil)
		return tg.BoolTrue, nil
	}

	participants, err := c.svcCtx.Repo.ChatClient.ChatGetChatParticipantIdList(c.ctx, &chatpb.TLChatGetChatParticipantIdList{
		ChatId: chatID,
	})
	if err != nil {
		c.logTypingPushFailure("chat participants query failed", senderUserID, chatID, err)
		return tg.BoolTrue, nil
	}
	if participants == nil {
		c.logTypingPushFailure("chat participants query returned nil", senderUserID, chatID, nil)
		return tg.BoolTrue, nil
	}

	updates := tg.MakeTLUpdateShort(&tg.TLUpdateShort{
		Update: tg.MakeTLUpdateChatUserTyping(&tg.TLUpdateChatUserTyping{
			ChatId: chatID,
			FromId: tg.MakePeerUser(senderUserID),
			Action: action,
		}),
		Date: int32(time.Now().Unix()),
	})
	for _, userID := range participants.Datas {
		if userID == senderUserID || userID <= 0 {
			continue
		}
		if err := c.svcCtx.Repo.PushTypingUpdates(c.ctx, userID, updates); err != nil {
			c.logTypingPushFailure("sync push failed", senderUserID, userID, err)
		}
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
