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

// ChatMigratedToChannel
// chat.migratedToChannel chat:MutableChat id:long access_hash:long = Bool;
func (c *ChatCore) ChatMigratedToChannel(in *chat.TLChatMigratedToChannel) (*mtproto.Bool, error) {
	var (
		chat2 = in.Chat
		_     = chat2
	)

	keys := []string{c.svcCtx.Dao.GetChatCacheKey(chat2.Id())}
	chat2.Walk(func(userId int64, participant *mtproto.ImmutableChatParticipant) error {
		keys = append(keys, c.svcCtx.Dao.GetChatParticipantCacheKey(participant.ChatId, participant.UserId))
		return nil
	})

	_, _, err := c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			tR := sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
				_, err := c.svcCtx.Dao.ChatsDAO.UpdateMigratedToTx(tx, in.Id, in.AccessHash, in.Chat.Id())
				if err != nil {
					result.Err = err
					return
				}
				c.svcCtx.Dao.ChatParticipantsDAO.UpdateStateByChatIdTx(tx, mtproto.ChatMemberStateMigrated, in.Chat.Id())
			})
			return 0, 0, tR.Err
		},
		keys...)
	if err != nil {
		c.Logger.Errorf("chat.migratedToChannel - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
