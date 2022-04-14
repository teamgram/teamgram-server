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
)

// ChatGetExportedChatInvite
// chat.getExportedChatInvite chat_id:long link:string = ExportedChatInvite;
func (c *ChatCore) ChatGetExportedChatInvite(in *chat.TLChatGetExportedChatInvite) (*mtproto.ExportedChatInvite, error) {
	var (
		link = chat.GetInviteHashByLink(in.Link)
	)

	chatInviteDO, err := c.svcCtx.Dao.ChatInvitesDAO.SelectByLink(c.ctx, link)
	if err != nil {
		c.Logger.Errorf("chat.getExportedChatInvite - error: %v", err)
		return nil, err
	} else if chatInviteDO == nil {
		err = mtproto.ErrChatLinkExists
		c.Logger.Errorf("chat.getExportedChatInvite - error: %v", err)
		return nil, err
	}

	return c.svcCtx.Dao.MakeChatInviteExported(c.ctx, chatInviteDO), nil
}
