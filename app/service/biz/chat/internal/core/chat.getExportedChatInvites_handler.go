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

// ChatGetExportedChatInvites
// chat.getExportedChatInvites flags:# chat_id:long admin_id:long revoked:flags.3?true offset_date:flags.2?int offset_link:flags.2?string limit:int = Vector<ExportedChatInvite>;
func (c *ChatCore) ChatGetExportedChatInvites(in *chat.TLChatGetExportedChatInvites) (*chat.Vector_ExportedChatInvite, error) {
	var (
		rInvites []*mtproto.ExportedChatInvite
		limit    = in.Limit
	)

	if limit == 0 {
		limit = 50
	}

	c.svcCtx.Dao.ChatInvitesDAO.SelectListByAdminIdWithCB(
		c.ctx,
		in.ChatId,
		in.AdminId,
		func(sz, i int, v *dataobject.ChatInvitesDO) {
			if in.Revoked {
				if v.Revoked {
					rInvites = append(rInvites, c.svcCtx.Dao.MakeChatInviteExported(c.ctx, v))
				}
			} else {
				if !v.Revoked {
					rInvites = append(rInvites, c.svcCtx.Dao.MakeChatInviteExported(c.ctx, v))
				}
			}
		})

	if rInvites == nil {
		rInvites = []*mtproto.ExportedChatInvite{}
	}

	var (
		offset = -1
	)

	if in.OffsetLink.GetValue() != "" && in.OffsetDate.GetValue() != 0 {
		for i, v := range rInvites {
			if in.OffsetLink.GetValue() == v.Link && in.OffsetDate.GetValue() == v.Date {
				offset = i
				break
			}
		}
	} else {
		offset = 0
	}

	if offset == -1 {
		rInvites = rInvites[0:0]
	} else {
		if len(rInvites) > offset+int(limit) {
			rInvites = rInvites[offset : offset+int(limit)]
		} else {
			rInvites = rInvites[offset:]
		}
	}

	return &chat.Vector_ExportedChatInvite{
		Datas: rInvites,
	}, nil
}
