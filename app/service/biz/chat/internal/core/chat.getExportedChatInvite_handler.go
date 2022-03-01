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
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
)

// ChatGetExportedChatInvite
// chat.getExportedChatInvite chat_id:long link:string = ExportedChatInvite;
func (c *ChatCore) ChatGetExportedChatInvite(in *chat.TLChatGetExportedChatInvite) (*mtproto.ExportedChatInvite, error) {
	chatInviteDO, err := c.svcCtx.Dao.ChatInvitesDAO.SelectByLink(c.ctx, in.Link)
	if err != nil {
		c.Logger.Errorf("chat.getExportedChatInvite - error: %v", err)
		return nil, err
	} else if chatInviteDO == nil {
		err = mtproto.ErrChatLinkExists
		c.Logger.Errorf("chat.getExportedChatInvite - error: %v", err)
		return nil, err
	}

	return c.makeChatInviteExported(chatInviteDO), nil
}

func (c *ChatCore) makeChatInviteExported(chatInviteDO *dataobject.ChatInvitesDO) *mtproto.ExportedChatInvite {
	rValue := mtproto.MakeTLChatInviteExported(&mtproto.ExportedChatInvite{
		Revoked:       chatInviteDO.Revoked,
		Permanent:     chatInviteDO.Permanent,
		RequestNeeded: chatInviteDO.RequestNeeded,
		Link:          chatInviteDO.Link,
		AdminId:       chatInviteDO.AdminId,
		Date:          int32(chatInviteDO.Date2),
		StartDate:     mtproto.MakeFlagsInt32(int32(chatInviteDO.StartDate)),
		ExpireDate:    mtproto.MakeFlagsInt32(int32(chatInviteDO.ExpireDate)),
		UsageLimit:    mtproto.MakeFlagsInt32(chatInviteDO.UsageLimit),
		Usage:         mtproto.MakeFlagsInt32(chatInviteDO.Usage2),
		Requested:     mtproto.MakeFlagsInt32(chatInviteDO.Requested),
		Title:         mtproto.MakeFlagsString(chatInviteDO.Title),
	}).To_ExportedChatInvite()

	// TODO: calc
	sz := c.svcCtx.Dao.CommonDAO.CalcSize(c.ctx, "chat_invite_participants", map[string]interface{}{
		"link": chatInviteDO.Link,
	})
	rValue.Usage = mtproto.MakeFlagsInt32(int32(sz))

	return rValue
}
