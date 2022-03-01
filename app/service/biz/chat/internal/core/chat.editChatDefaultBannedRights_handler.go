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
	"math"
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChatEditChatDefaultBannedRights
// chat.editChatDefaultBannedRights chat_id:long operator_id:long banned_rights:ChatBannedRights = MutableChat;
func (c *ChatCore) ChatEditChatDefaultBannedRights(in *chat.TLChatEditChatDefaultBannedRights) (*chat.MutableChat, error) {
	var (
		now   = time.Now().Unix()
		chat2 *chat.MutableChat
		me    *chat.ImmutableChatParticipant
		err   error
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId, in.OperatorId)
	if err != nil {
		return nil, err
	}

	me, _ = chat2.GetImmutableChatParticipant(in.OperatorId)
	if me == nil || me.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	if me.IsChatMemberNormal() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	// b, d := mtproto.FromChatBannedRights(in.BannedRights.To_ChatBannedRights())
	// _ = d
	bannedRights := in.BannedRights
	if bannedRights.UntilDate == 0 {
		bannedRights.UntilDate = math.MaxInt32
	}

	_, err = c.svcCtx.Dao.ChatsDAO.UpdateDefaultBannedRights(c.ctx, int64(bannedRights.ToBannedRights()), in.ChatId)
	if err != nil {
		return nil, err
	}

	chat2.Chat.DefaultBannedRights = bannedRights
	chat2.Chat.Version += 1
	chat2.Chat.Date = now

	return chat2, nil
}
