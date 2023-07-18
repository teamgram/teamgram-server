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
	"context"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
)

// ChatAddChatUser
// chat.addChatUser chat_id:long inviter_id:long user_id:long = MutableChat;
func (c *ChatCore) ChatAddChatUser(in *chat.TLChatAddChatUser) (*mtproto.MutableChat, error) {
	var (
		now                       = time.Now().Unix()
		chat2                     *mtproto.MutableChat
		me, willAdd               *mtproto.ImmutableChatParticipant
		err                       error
		chatId, inviterId, userId = in.ChatId, in.InviterId, in.UserId
	)

	if inviterId != 0 {
		chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, chatId, inviterId, userId)
		if err != nil {
			c.Logger.Errorf("chat.addChatUser - error: %v", err)
			return nil, err
		} else {
			if chat2.Deactivated() && chat2.GetChat().GetMigratedTo() != nil {
				err = mtproto.ErrMigratedToChannel
				c.Logger.Errorf("chat.addChatUser - error: %v", err)
				return nil, err
			}
		}

		me, _ = chat2.GetImmutableChatParticipant(inviterId)
		if me == nil || (me.State != mtproto.ChatMemberStateNormal && !me.IsChatMemberCreator()) {
			err = mtproto.ErrInputUserDeactivated
			c.Logger.Errorf("chat.addChatUser - error: %v", err)
			return nil, err
		}
	} else {
		chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, chatId, userId)
		if err != nil {
			c.Logger.Errorf("chat.addChatUser - error: %v", err)
			return nil, err
		}
		inviterId = chat2.Creator()
	}

	if chat2.ParticipantsCount() >= 200 {
		err = mtproto.ErrUsersTooFew
		c.Logger.Errorf("chat.addChatUser - error: %v", err)
		return nil, err
	}

	willAdd, _ = chat2.GetImmutableChatParticipant(userId)
	if willAdd != nil && willAdd.State == mtproto.ChatMemberStateNormal {
		err = mtproto.ErrUserAlreadyParticipant
		c.Logger.Errorf("chat.addChatUser - error: %v", err)
		return nil, err
	}

	if me != nil {
		// TODO(@benqi): check
		// 400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this
		if !me.CanInviteUsers() &&
			!chat2.DefaultBannedRights().CanInviteUsers(int32(time.Now().Unix())) {
			err = mtproto.ErrChatAdminRequired
			c.Logger.Errorf("chat.addChatUser - error: %v", err)
			return nil, err
		}
	}

	_, _, err = c.svcCtx.Dao.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
				chatParticipantDO := &dataobject.ChatParticipantsDO{
					ChatId:          chat2.Chat.Id,
					UserId:          userId,
					ParticipantType: mtproto.ChatMemberNormal,
					InviterUserId:   inviterId,
					InvitedAt:       now,
					Date2:           now,
					IsBot:           in.GetIsBot(),
				}
				if chat2.Chat.Creator == userId {
					chatParticipantDO.ParticipantType = mtproto.ChatMemberCreator
				}
				if willAdd == nil {
					lastInsertId, _, err2 := c.svcCtx.Dao.ChatParticipantsDAO.InsertTx(tx, chatParticipantDO)
					if err2 != nil {
						result.Err = err2
						return
					}
					chatParticipantDO.Id = lastInsertId
					willAdd = c.svcCtx.Dao.MakeImmutableChatParticipant(chatParticipantDO)
				} else {
					chatParticipantDO.Id = willAdd.Id
					_, err2 := c.svcCtx.Dao.ChatParticipantsDAO.UpdateTx(
						tx,
						chatParticipantDO.ParticipantType,
						inviterId,
						now,
						in.GetIsBot(),
						chatParticipantDO.Id)
					if err != nil {
						result.Err = err2
						return
					}
				}
				chat2.Chat.ParticipantsCount += 1
				chat2.Chat.Version += 1
				chat2.Chat.Date = now
				_, result.Err = c.svcCtx.Dao.ChatsDAO.UpdateParticipantCountTx(tx, chat2.Chat.ParticipantsCount, chatId)
			})
			return 0, 0, tR.Err
		},
		c.svcCtx.Dao.GetChatCacheKey(chat2.Id()),
		c.svcCtx.Dao.GetChatParticipantCacheKey(chat2.Id(), userId))

	if err != nil {
		c.Logger.Errorf("chat.addChatUser - error: %v", err)
		return nil, err
	}

	return chat2, nil
}
