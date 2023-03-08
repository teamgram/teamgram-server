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
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChatEditChatTitle
// chat.editChatTitle chat_id:long edit_user_id:long title:string = MutableChat;
func (c *ChatCore) ChatEditChatTitle(in *chat.TLChatEditChatTitle) (*mtproto.MutableChat, error) {
	var (
		now                = time.Now().Unix()
		chat2              *mtproto.MutableChat
		me                 *mtproto.ImmutableChatParticipant
		err                error
		chatId, editUserId = in.ChatId, in.EditUserId
	)

	if in.Title == "" {
		err = mtproto.ErrChatTitleEmpty
		c.Logger.Errorf("chat.editChatTitle - error: %v", err)
		return nil, err
	}

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, chatId, editUserId)
	if err != nil {
		c.Logger.Errorf("chat.editChatTitle - error: %v", err)
		return nil, err
	}
	if chat2.Chat.Title == in.Title {
		err = mtproto.ErrChatNotModified
		c.Logger.Errorf("chat.editChatTitle - error: %v", err)
		return nil, err
	}

	me, _ = chat2.GetImmutableChatParticipant(editUserId)
	if me == nil || me.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("chat.editChatTitle - error: %v", err)
		return nil, err
	}

	// TODO(@benqi): check
	// 400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this
	if !me.CanChangeInfo() {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.editChatTitle - error: %v", err)
		return nil, err
	}

	_, _, err = c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowAffected, err2 := c.svcCtx.Dao.ChatsDAO.UpdateTitle(c.ctx, in.Title, chatId)
			return 0, rowAffected, err2
		},
		c.svcCtx.Dao.GetChatCacheKey(chatId))
	if err != nil {
		c.Logger.Errorf("chat.editChatTitle - error: %v", err)
		return nil, err
	}

	chat2.Chat.Title = in.Title
	chat2.Chat.Version += 1
	chat2.Chat.Date = now
	return chat2, nil
}
