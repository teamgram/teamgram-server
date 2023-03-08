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
	"math"
	"time"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChatEditChatDefaultBannedRights
// chat.editChatDefaultBannedRights chat_id:long operator_id:long banned_rights:ChatBannedRights = MutableChat;
func (c *ChatCore) ChatEditChatDefaultBannedRights(in *chat.TLChatEditChatDefaultBannedRights) (*mtproto.MutableChat, error) {
	var (
		now   = time.Now().Unix()
		chat2 *mtproto.MutableChat
		me    *mtproto.ImmutableChatParticipant
		err   error
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId, in.OperatorId)
	if err != nil {
		return nil, err
	}

	me, _ = chat2.GetImmutableChatParticipant(in.OperatorId)
	if me == nil || me.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("chat.editChatDefaultBannedRights - error: %v", err)
		return nil, err
	}

	if me.IsChatMemberNormal() {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.editChatDefaultBannedRights - error: %v", err)
		return nil, err
	}

	// b, d := mtproto.FromChatBannedRights(in.BannedRights.To_ChatBannedRights())
	// _ = d
	bannedRights := in.BannedRights
	if bannedRights.UntilDate == 0 {
		bannedRights.UntilDate = math.MaxInt32
	}

	_, _, err = c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			affected, err2 := c.svcCtx.Dao.ChatsDAO.UpdateDefaultBannedRights(c.ctx, int64(bannedRights.ToBannedRights()), in.ChatId)
			return 0, affected, err2
		},
		c.svcCtx.Dao.GetChatCacheKey(in.ChatId))
	if err != nil {
		c.Logger.Errorf("chat.editChatDefaultBannedRights - error: %v", err)
		return nil, err
	}

	chat2.Chat.DefaultBannedRights = bannedRights
	chat2.Chat.Version += 1
	chat2.Chat.Date = now

	return chat2, nil
}
