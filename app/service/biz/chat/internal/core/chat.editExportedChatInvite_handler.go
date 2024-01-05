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
	"time"
)

// ChatEditExportedChatInvite
// chat.editExportedChatInvite flags:# self_id:long chat_id:long revoked:flags.2?true link:string expire_date:flags.0?int usage_limit:flags.1?int request_needed:flags.3?Bool title:flags.4?string = ExportedChatInvite;
func (c *ChatCore) ChatEditExportedChatInvite(in *chat.TLChatEditExportedChatInvite) (*chat.Vector_ExportedChatInvite, error) {
	var (
		hash        = chat.GetInviteHashByLink(in.Link)
		chatInvites = make([]*mtproto.ExportedChatInvite, 0, 2)
	)

	chatInviteDO, err := c.svcCtx.Dao.ChatInvitesDAO.SelectByLink(c.ctx, hash)
	if err != nil {
		c.Logger.Errorf("chat.editExportedChatInvite - error: %v", err)
		return nil, err
	} else if chatInviteDO == nil {
		err = mtproto.ErrInternalServerError
		c.Logger.Errorf("chat.editExportedChatInvite - error: %v", err)
		return nil, err
	}

	if in.Revoked {
		c.svcCtx.Dao.ChatInvitesDAO.Update(
			c.ctx,
			map[string]interface{}{
				"revoked": in.Revoked,
			},
			in.ChatId,
			hash)
		chatInviteDO.Revoked = in.Revoked
		chatInvites = append(chatInvites, c.svcCtx.Dao.MakeChatInviteExported(c.ctx, chatInviteDO))

		// chatInvites
		if chatInviteDO.Permanent {
			chatInviteDO = &dataobject.ChatInvitesDO{
				ChatId:        in.ChatId,
				AdminId:       chatInviteDO.AdminId,
				Link:          chat.GenChatInviteHash(),
				Permanent:     chatInviteDO.Permanent,
				Revoked:       false,
				RequestNeeded: false,
				StartDate:     0,
				ExpireDate:    0,
				UsageLimit:    0,
				Usage2:        0,
				Requested:     0,
				Title:         "",
				Date2:         time.Now().Unix(),
			}
			c.svcCtx.Dao.ChatInvitesDAO.Insert(c.ctx, chatInviteDO)
			chatInvites = append(chatInvites, c.svcCtx.Dao.MakeChatInviteExported(c.ctx, chatInviteDO))
			c.svcCtx.Dao.ChatParticipantsDAO.UpdateLink(
				c.ctx,
				chatInviteDO.Link,
				in.ChatId,
				chatInviteDO.AdminId)
		}
	} else {
		cMap := map[string]interface{}{}

		if in.GetExpireDate() != nil {
			cMap["expire_date"] = in.GetExpireDate().GetValue()
			chatInviteDO.ExpireDate = int64(in.GetExpireDate().GetValue())
		}
		if in.GetUsageLimit() != nil {
			cMap["usage_limit"] = in.GetUsageLimit().GetValue()
			chatInviteDO.UsageLimit = in.GetUsageLimit().GetValue()
		}
		if in.GetRequestNeeded() != nil {
			cMap["request_needed"] = mtproto.FromBool(in.GetRequestNeeded())
			chatInviteDO.RequestNeeded = mtproto.FromBool(in.GetRequestNeeded())
		}
		if in.GetTitle() != nil {
			cMap["title"] = in.GetTitle().GetValue()
			chatInviteDO.Title = in.GetTitle().GetValue()
		}

		c.svcCtx.Dao.ChatInvitesDAO.Update(
			c.ctx,
			cMap,
			in.ChatId,
			hash)
		chatInvites = append(chatInvites, c.svcCtx.Dao.MakeChatInviteExported(c.ctx, chatInviteDO))
	}

	return &chat.Vector_ExportedChatInvite{
		Datas: chatInvites,
	}, nil
}
