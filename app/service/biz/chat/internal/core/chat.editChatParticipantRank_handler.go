// Copyright (c) 2026 The Teamgram Authors (https://teamgram.net).
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

package core

import (
	"context"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChatEditChatParticipantRank
// chat.editChatParticipantRank self_id:long chat_id:long participant:long rank:string = MutableChat;
func (c *ChatCore) ChatEditChatParticipantRank(in *chat.TLChatEditChatParticipantRank) (*mtproto.MutableChat, error) {
	var (
		now         = time.Now().Unix()
		chat2       *mtproto.MutableChat
		me          *mtproto.ImmutableChatParticipant
		participant *mtproto.ImmutableChatParticipant
		err         error
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId)
	if err != nil {
		c.Logger.Errorf("chat.editChatParticipantRank - error: %v", err)
		return nil, mtproto.ErrChatIdInvalid
	}

	me, _ = chat2.GetImmutableChatParticipant(in.SelfId)
	if me == nil || me.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrUserNotParticipant
		c.Logger.Errorf("chat.editChatParticipantRank - error: %v", err)
		return nil, err
	}

	participant, _ = chat2.GetImmutableChatParticipant(in.Participant)
	if participant == nil || participant.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("chat.editChatParticipantRank - error: %v", err)
		return nil, err
	}

	if !canEditChatParticipantRank(me, in.Participant) {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.editChatParticipantRank - error: %v", err)
		return nil, err
	}

	if participant.Rank == in.Rank {
		err = mtproto.ErrChatNotModified
		c.Logger.Errorf("chat.editChatParticipantRank - error: %v", err)
		return nil, err
	}

	_, _, err = c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err2 := c.svcCtx.Dao.ChatParticipantsDAO.UpdateRank(ctx, in.Rank, participant.Id)
			return 0, rowsAffected, err2
		},
		c.svcCtx.Dao.GetChatCacheKey(in.ChatId),
		c.svcCtx.Dao.GetChatParticipantCacheKey(in.ChatId, in.Participant))

	if err != nil {
		c.Logger.Errorf("chat.editChatParticipantRank - error: %v", err)
		return nil, err
	}

	participant.Rank = in.Rank
	chat2.Chat.Version += 1
	chat2.Chat.Date = now

	return chat2, nil
}

func canEditChatParticipantRank(operator *mtproto.ImmutableChatParticipant, participantID int64) bool {
	if operator == nil || operator.State != mtproto.ChatMemberStateNormal {
		return false
	}
	if operator.IsChatMemberCreator() {
		return true
	}
	return operator.UserId == participantID &&
		operator.IsChatMemberAdmin() &&
		operator.AdminRights.GetManageRanks()
}
