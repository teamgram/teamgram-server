// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
	"time"
)

// ChatImportChatInvite2
// chat.importChatInvite2 self_id:long hash:string = ChatInviteImported;
func (c *ChatCore) ChatImportChatInvite2(in *chat.TLChatImportChatInvite2) (*chat.ChatInviteImported, error) {
	chatInviteDO, err := c.svcCtx.Dao.ChatInvitesDAO.SelectByLink(c.ctx, in.Hash)
	if err != nil {
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	} else if chatInviteDO == nil {
		err = mtproto.ErrInviteHashInvalid
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	}

	// check expire
	if chatInviteDO.Revoked {
		c.Logger.Errorf("chat.checkChatInvite - error: invite hash %s expired", in.Hash)
		err = mtproto.ErrInviteHashExpired
		return nil, err
	}

	if chatInviteDO.ExpireDate != 0 && time.Now().Unix() > chatInviteDO.ExpireDate {
		err = mtproto.ErrInviteHashExpired
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	}
	if chatInviteDO.UsageLimit > 0 {
		sz := c.svcCtx.Dao.GetLinkInviteSize(c.ctx, chatInviteDO.Link)
		if sz >= chatInviteDO.UsageLimit {
			err = mtproto.ErrInviteHashExpired
			c.Logger.Errorf("chat.importChatInvite - error: %v", err)
			return nil, err
		}
	}

	if chatInviteDO.RequestNeeded {
		mChat, err2 := c.svcCtx.Dao.GetMutableChat(c.ctx, chatInviteDO.ChatId, chatInviteDO.AdminId, in.SelfId)
		if err2 != nil {
			err2 = mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("chat.importChatInvite - error: %v", err2)
			return nil, err2
		} else if mChat.Deactivated() && mChat.GetChat().GetMigratedTo() != nil {
			err2 = mtproto.ErrMigratedToChannel
			c.Logger.Errorf("chat.importChatInvite - error: %v", err2)
			return nil, err2
		}

		if mChat.ParticipantsCount() >= 200 {
			err2 = mtproto.ErrUsersTooMuch
			c.Logger.Errorf("chat.importChatInvite - error: %v", err2)
			return nil, err2
		}

		c.svcCtx.Dao.ChatInviteParticipantsDAO.Insert(c.ctx, &dataobject.ChatInviteParticipantsDO{
			ChatId:    chatInviteDO.ChatId,
			Link:      in.Hash,
			UserId:    in.SelfId,
			Requested: chatInviteDO.RequestNeeded,
			Date2:     time.Now().Unix(),
		})

		requesters := chat.MakeTLRecentChatInviteRequesters(&chat.RecentChatInviteRequesters{
			RequestsPending:  0,
			RecentRequesters: []int64{},
		}).To_RecentChatInviteRequesters()
		c.svcCtx.Dao.ChatInviteParticipantsDAO.SelectRecentRequestedListWithCB(
			c.ctx,
			mChat.Id(),
			func(sz, i int, v *dataobject.ChatInviteParticipantsDO) {
				requesters.RequestsPending += 1
				requesters.RecentRequesters = append(requesters.RecentRequesters, v.UserId)
			})

		return chat.MakeTLChatInviteImported(&chat.ChatInviteImported{
			Chat:       mChat,
			Requesters: requesters,
		}).To_ChatInviteImported(), nil
	} else {
		chat2, err := c.ChatAddChatUser(&chat.TLChatAddChatUser{
			ChatId:    chatInviteDO.ChatId,
			InviterId: chatInviteDO.AdminId,
			UserId:    in.SelfId,
		})
		if err != nil {
			c.Logger.Errorf("chat.importChatInvite - error: %v", err)
			return nil, err
		}

		c.svcCtx.Dao.ChatInviteParticipantsDAO.Insert(c.ctx, &dataobject.ChatInviteParticipantsDO{
			ChatId:    chatInviteDO.ChatId,
			Link:      in.Hash,
			UserId:    in.SelfId,
			Requested: false,
			Date2:     time.Now().Unix(),
		})

		return chat.MakeTLChatInviteImported(&chat.ChatInviteImported{
			Chat:       chat2,
			Requesters: nil,
		}).To_ChatInviteImported(), nil
	}
}
