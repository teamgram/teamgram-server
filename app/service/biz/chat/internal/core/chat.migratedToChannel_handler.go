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
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChatMigratedToChannel
// chat.migratedToChannel chat:MutableChat id:long access_hash:long = Bool;
func (c *ChatCore) ChatMigratedToChannel(in *chat.TLChatMigratedToChannel) (*mtproto.Bool, error) {
	var (
		chat = in.Chat
		_    = chat
	)

	_ = sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, err := c.svcCtx.Dao.ChatsDAO.UpdateMigratedToTx(tx, in.Id, in.AccessHash, in.Chat.Id())
		if err != nil {
			result.Err = err
			return
		}
		c.svcCtx.Dao.ChatParticipantsDAO.UpdateStateByChatIdTx(tx, mtproto.ChatMemberStateMigrated, in.Chat.Id())
	})

	return mtproto.BoolTrue, nil
}
