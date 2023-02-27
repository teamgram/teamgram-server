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
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"time"
)

// ChatCheckChatInvite
// chat.checkChatInvite self_id:long hash:string = ChatInvite;
func (c *ChatCore) ChatCheckChatInvite(in *chat.TLChatCheckChatInvite) (*chat.ChatInviteExt, error) {
	chatInviteDO, err := c.svcCtx.Dao.ChatInvitesDAO.SelectByLink(c.ctx, in.Hash)
	if err != nil {
		c.Logger.Errorf("chat.checkChatInvite - error: %v", err)
		return nil, err
	} else if chatInviteDO == nil {
		c.Logger.Errorf("chat.checkChatInvite - error: not found hash %s", in.Hash)
		err = mtproto.ErrInviteHashInvalid
		return nil, err
	}

	// check expire
	if chatInviteDO.Revoked {
		c.Logger.Errorf("chat.checkChatInvite - error: invite hash %s expired", in.Hash)
		err = mtproto.ErrInviteHashExpired
		return nil, err
	}

	if chatInviteDO.ExpireDate != 0 && time.Now().Unix() > chatInviteDO.ExpireDate {
		c.Logger.Errorf("chat.checkChatInvite - error: invite hash %s expired", in.Hash)
		err = mtproto.ErrInviteHashExpired
		return nil, err
	}

	if chatInviteDO.UsageLimit > 0 {
		// TODO: calc
		sz := c.svcCtx.Dao.CommonDAO.CalcSize(
			c.ctx,
			"chat_invite_participants",
			map[string]interface{}{
				"link": chatInviteDO.Link,
			})
		if sz >= int(chatInviteDO.UsageLimit) {
			err = mtproto.ErrInviteHashExpired
			c.Logger.Errorf("chat.importChatInvite - error: %v", err)
			return nil, err
		}
	}

	mChat, err := c.svcCtx.Dao.GetMutableChat(c.ctx, chatInviteDO.ChatId)
	if err != nil {
		c.Logger.Errorf("chat.checkChatInvite - error: %v", err)
		return nil, err
	} else {
		if mChat.Deactivated() && mChat.GetChat().GetMigratedTo() != nil {
			err = mtproto.ErrMigratedToChannel
			c.Logger.Errorf("chat.checkChatInvite - error: %v", err)
			return nil, err
		}
	}

	admin, _ := mChat.GetImmutableChatParticipant(chatInviteDO.AdminId)
	if admin == nil || !admin.CanInviteUsers() {
		err = mtproto.ErrInviteHashExpired
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	}

	me, _ := mChat.GetImmutableChatParticipant(in.SelfId)
	if me == nil || !me.IsChatMemberStateNormal() {
		rValue := chat.MakeTLChatInvite(&chat.ChatInviteExt{
			RequestNeeded:     chatInviteDO.RequestNeeded,
			Title:             mChat.Title(),
			About:             mtproto.MakeFlagsString(mChat.About()),
			Photo:             mChat.Photo(),
			ParticipantsCount: mChat.Chat.ParticipantsCount,
			Participants:      mChat.ParticipantIdList(),
		}).To_ChatInviteExt()

		return rValue, nil
	} else {
		return chat.MakeTLChatInviteAlready(&chat.ChatInviteExt{
			Chat: mChat,
		}).To_ChatInviteExt(), nil
	}
}
