/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
)

// ChatAddChatUser
// chat.addChatUser chat_id:long inviter_id:long user_id:long = MutableChat;
func (c *ChatCore) ChatAddChatUser(in *chat.TLChatAddChatUser) (*chat.MutableChat, error) {
	var (
		now                       = time.Now().Unix()
		chat2                     *chat.MutableChat
		me, willAdd               *chat.ImmutableChatParticipant
		err                       error
		chatId, inviterId, userId = in.ChatId, in.InviterId, in.UserId
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, chatId, inviterId, userId)
	if err != nil {
		return nil, err
	}

	me, _ = chat2.GetImmutableChatParticipant(inviterId)
	if me == nil || (me.State != mtproto.ChatMemberStateNormal && !me.IsChatMemberCreator()) {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	willAdd, _ = chat2.GetImmutableChatParticipant(userId)
	if willAdd != nil && willAdd.State == mtproto.ChatMemberStateNormal {
		err = mtproto.ErrUserAlreadyParticipant
		return nil, err
	}

	// TODO(@benqi): check
	// 400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this
	if !me.CanInviteUsers() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		chatParticipantDO := &dataobject.ChatParticipantsDO{
			ChatId:          chat2.Chat.Id,
			UserId:          userId,
			ParticipantType: mtproto.ChatMemberNormal,
			InviterUserId:   inviterId,
			InvitedAt:       now,
		}
		if chat2.Chat.Creator == userId {
			chatParticipantDO.ParticipantType = mtproto.ChatMemberCreator
		}
		if willAdd == nil {
			lastInsertId, _, err := c.svcCtx.Dao.ChatParticipantsDAO.InsertTx(tx, chatParticipantDO)
			if err != nil {
				result.Err = err
				return
			}
			chatParticipantDO.Id = lastInsertId
			willAdd = c.svcCtx.Dao.MakeImmutableChatParticipant(chatParticipantDO)
		} else {
			chatParticipantDO.Id = willAdd.Id
			if _, err = c.svcCtx.Dao.ChatParticipantsDAO.UpdateTx(tx, chatParticipantDO.ParticipantType, inviterId, now, chatParticipantDO.Id); err != nil {
				result.Err = err
				return
			}
		}
		chat2.Chat.ParticipantsCount += 1
		chat2.Chat.Version += 1
		chat2.Chat.Date = now
		_, result.Err = c.svcCtx.Dao.ChatsDAO.UpdateParticipantCountTx(tx, chat2.Chat.ParticipantsCount, chatId)
	})

	if tR.Err != nil {
		return nil, tR.Err
	}
	return chat2, nil
}
