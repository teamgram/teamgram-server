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

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
)

// ChatEditChatAdmin
// chat.editChatAdmin chat_id:long operator_id:long edit_chat_admin_id:long is_admin:Bool = MutableChat;
func (c *ChatCore) ChatEditChatAdmin(in *chat.TLChatEditChatAdmin) (*chat.MutableChat, error) {
	var (
		now           = time.Now().Unix()
		chat2         *chat.MutableChat
		me, editAdmin *chat.ImmutableChatParticipant
		err           error
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId, in.OperatorId, in.EditChatAdminId)
	if err != nil {
		return nil, err
	}

	me, _ = chat2.GetImmutableChatParticipant(in.OperatorId)
	if me == nil || me.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	editAdmin, _ = chat2.GetImmutableChatParticipant(in.EditChatAdminId)
	if editAdmin != nil && editAdmin.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	if !me.CanAdminAddAdmins() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	if mtproto.FromBool(in.IsAdmin) {
		_, err = c.svcCtx.Dao.ChatParticipantsDAO.UpdateParticipantType(c.ctx, mtproto.ChatMemberAdmin, editAdmin.Id)
		if err != nil {
			return nil, err
		}
		editAdmin.AdminRights = mtproto.MakeDefaultChatAdminRights()
		editAdmin.ParticipantType = mtproto.ChatMemberAdmin
	} else {
		_, err = c.svcCtx.Dao.ChatParticipantsDAO.UpdateParticipantType(c.ctx, mtproto.ChatMemberNormal, editAdmin.Id)
		if err != nil {
			return nil, err
		}
		editAdmin.AdminRights = nil
		editAdmin.ParticipantType = mtproto.ChatMemberNormal
	}

	chat2.Chat.Version += 1
	chat2.Chat.Date = now
	return chat2, nil
}
