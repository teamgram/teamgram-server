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
	"github.com/teamgram/teamgram-server/pkg/env2"
	"strings"
)

// ChatGetChatInviteImporters
// chat.getChatInviteImporters flags:# self_id:long chat_id:long requested:flags.0?true link:flags.1?string q:flags.2?string offset_date:int offset_user:long limit:int = Vector<ChatInviteImporter>;
func (c *ChatCore) ChatGetChatInviteImporters(in *chat.TLChatGetChatInviteImporters) (*chat.Vector_ChatInviteImporter, error) {
	var (
		rInvites []*mtproto.ChatInviteImporter
		link     = in.GetLink().GetValue()
	)

	if strings.HasPrefix(link, "https://"+env2.TDotMe+"t.me/+") {
		link = link[len("https://"+env2.TDotMe+"/+"):]
	}
	c.Logger.Errorf("link: %s", link)

	c.svcCtx.Dao.ChatInviteParticipantsDAO.SelectListByLinkWithCB(
		c.ctx,
		link,
		func(i int, v *dataobject.ChatInviteParticipantsDO) {
			rInvites = append(rInvites, mtproto.MakeTLChatInviteImporter(&mtproto.ChatInviteImporter{
				Requested:  false,
				UserId:     v.UserId,
				Date:       int32(v.Date2),
				About:      nil,
				ApprovedBy: nil,
			}).To_ChatInviteImporter())
			c.Logger.Errorf("do: %v", v)
		})

	if rInvites == nil {
		rInvites = []*mtproto.ChatInviteImporter{}
	}

	var (
		offset = 0
	)

	for i, v := range rInvites {
		if in.OffsetUser == v.UserId && in.OffsetDate == v.Date {
			offset = i + 1
			break
		}
	}
	if len(rInvites) >= offset+int(in.Limit) {
		rInvites = rInvites[offset : offset+int(in.Limit)]
	} else {
		rInvites = rInvites[offset:]
	}

	c.Logger.Errorf("offset: %d", offset)

	return &chat.Vector_ChatInviteImporter{
		Datas: rInvites,
	}, nil
}
