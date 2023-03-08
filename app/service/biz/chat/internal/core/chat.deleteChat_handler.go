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
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChatDeleteChat
// chat.deleteChat chat_id:long operator_id:long = MutableChat;
func (c *ChatCore) ChatDeleteChat(in *chat.TLChatDeleteChat) (*mtproto.MutableChat, error) {
	mChat, err := c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId)
	if err != nil {
		c.Logger.Errorf("chat.deleteChat - error: %v", err)
		return nil, err
	}

	if in.OperatorId == 0 {
		in.OperatorId = mChat.Creator()
	}

	if mChat.Creator() != in.OperatorId {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.deleteChat - error: %v", err)
		return nil, err
	}

	keys := []string{c.svcCtx.Dao.GetChatCacheKey(in.ChatId)}
	mChat.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
		keys = append(keys, c.svcCtx.Dao.GetChatParticipantCacheKey(participant.ChatId, participant.UserId))
		return nil
	})

	_, _, err = c.svcCtx.Dao.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
				// kicked
				c.svcCtx.Dao.ChatParticipantsDAO.UpdateStateByChatIdTx(tx, mtproto.ChatMemberStateKicked, in.ChatId)
				c.svcCtx.Dao.ChatsDAO.UpdateParticipantCountTx(tx, 0, in.ChatId)
				c.svcCtx.Dao.ChatsDAO.UpdateDeactivatedTx(tx, false, in.ChatId)
			})
			return 0, 0, tR.Err
		},
		keys...)
	if err != nil {
		c.Logger.Errorf("chat.deleteChat - error: %v", err)
		return nil, err
	}

	return mChat, nil
}
