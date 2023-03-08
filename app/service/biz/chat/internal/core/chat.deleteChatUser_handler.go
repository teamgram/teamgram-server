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
)

// ChatDeleteChatUser
// chat.deleteChatUser chat_id:long operator_id:long delete_user_id:long = MutableChat;
func (c *ChatCore) ChatDeleteChatUser(in *chat.TLChatDeleteChatUser) (*mtproto.MutableChat, error) {
	var (
		now             = time.Now().Unix()
		chat2           *mtproto.MutableChat
		me, deletedUser *mtproto.ImmutableChatParticipant
		err             error
		chatId          = in.ChatId
		operatorId      = in.OperatorId
		deleteUserId    = in.DeleteUserId
		kicked          = operatorId != deleteUserId
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, chatId)
	if err != nil {
		c.Logger.Errorf("chat.deleteChatUser - error: %v", err)
		return nil, err
	}

	if operatorId == 0 {
		operatorId = chat2.Creator()
	}

	me, _ = chat2.GetImmutableChatParticipant(operatorId)
	if me == nil {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("chat.deleteChatUser - error: %v", err)
		return nil, err
	}

	if kicked {
		if me.State != mtproto.ChatMemberStateNormal {
			err = mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("chat.deleteChatUser - error: %v", err)
			return nil, err
		}

		if !me.CanAdminBanUsers() {
			err = mtproto.ErrChatAdminRequired
			c.Logger.Errorf("chat.deleteChatUser - error: %v", err)
			return nil, err
		}
		//switch me.ChatParticipant.PredicateName {
		//case mtproto.Predicate_chatParticipantCreator:
		//default:
		//	err = mtproto.ErrChatAdminRequired
		//	return nil, err
		//}

		deletedUser, _ = chat2.GetImmutableChatParticipant(deleteUserId)
		if deletedUser == nil {
			// USER_NOT_PARTICIPANT
			err = mtproto.ErrUserNotParticipant
			c.Logger.Errorf("chat.deleteChatUser - error: %v", err)
			return nil, err
		} else if deletedUser.State != mtproto.ChatMemberStateNormal {
			err = mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("chat.deleteChatUser - error: %v", err)
			return nil, err
		}
	} else {
		// left
		deletedUser = me
		if me.State != mtproto.ChatMemberStateNormal {
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}
	}

	_, _, err = c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			// deletedUser.Dialog.TopMessage
			tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
				if kicked {
					_, result.Err = c.svcCtx.Dao.ChatParticipantsDAO.UpdateKickedTx(tx, now, deletedUser.Id)
					if result.Err != nil {
						c.Logger.Errorf("chat.deleteChatUser - error: %v", err)
						return
					}
					deletedUser.State = mtproto.ChatMemberStateKicked
				} else {
					_, result.Err = c.svcCtx.Dao.ChatParticipantsDAO.UpdateLeftTx(tx, now, deletedUser.Id)
					if result.Err != nil {
						c.Logger.Errorf("chat.deleteChatUser - error: %v", err)
						return
					}
					deletedUser.State = mtproto.ChatMemberStateLeft
				}
				chat2.Chat.ParticipantsCount -= 1
				chat2.Chat.Date = now
				chat2.Chat.Version += 1
				_, result.Err = c.svcCtx.Dao.ChatsDAO.UpdateParticipantCountTx(tx, chat2.Chat.ParticipantsCount, chat2.Chat.Id)
				if result.Err != nil {
					c.Logger.Errorf("chat.deleteChatUser - error: %v", err)
					return
				}

				_, result.Err = c.svcCtx.Dao.ChatInviteParticipantsDAO.DeleteTx(tx, chat2.Chat.Id, deleteUserId)
			})
			return 0, 0, tR.Err
		},
		c.svcCtx.Dao.GetChatCacheKey(chat2.Id()),
		c.svcCtx.Dao.GetChatParticipantCacheKey(chat2.Id(), deleteUserId))

	if err != nil {
		c.Logger.Errorf("chat.deleteChatUser - error: %v", err)
		return nil, err
	}

	return chat2, nil
}
