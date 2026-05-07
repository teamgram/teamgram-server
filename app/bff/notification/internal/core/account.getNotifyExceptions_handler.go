// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"time"

	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/app/bff/notification/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountGetNotifyExceptions
// account.getNotifyExceptions#53577479 flags:# compare_sound:flags.1?true compare_stories:flags.2?true peer:flags.0?InputNotifyPeer = Updates;
func (c *NotificationCore) AccountGetNotifyExceptions(in *tg.TLAccountGetNotifyExceptions) (*tg.Updates, error) {
	settingsList, err := c.svcCtx.Repo.UserClient.UserGetNotifySettingsList(c.ctx, &repository.GetNotifySettingsList{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("account.getNotifyExceptions - error: %v", err)
		return nil, err
	}

	var (
		userIdList []int64
		chatIdList []int64
		updates    []tg.UpdateClazz
	)

	for _, settings := range settingsList.Datas {
		notifyPeer := makeNotifyPeer(settings.PeerType, settings.PeerId)
		update := tg.MakeTLUpdateNotifySettings(&tg.TLUpdateNotifySettings{
			Peer:           notifyPeer,
			NotifySettings: settings.Settings,
		})
		updates = append(updates, update)

		switch settings.PeerType {
		case tg.PEER_USER:
			userIdList = append(userIdList, settings.PeerId)
		case tg.PEER_CHAT:
			chatIdList = append(chatIdList, settings.PeerId)
		case tg.PEER_CHANNEL:
			// TODO: handle channels via plugin
		}
	}

	// Fetch users
	var users []tg.UserClazz
	if len(userIdList) > 0 {
		users, err = userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, c.MD.UserId, userIdList, userprojection.MissingStoredReference)
		if err != nil {
			c.Logger.Errorf("account.getNotifyExceptions - error fetching users: %v", err)
			return nil, err
		}
	}

	// Fetch chats
	var chats []tg.ChatClazz
	if len(chatIdList) > 0 {
		mChats, err := c.svcCtx.Repo.ChatClient.ChatGetChatListByIdList(c.ctx, &repository.GetChatListByIdList{
			SelfId: c.MD.UserId,
			IdList: chatIdList,
		})
		if err != nil {
			c.Logger.Errorf("account.getNotifyExceptions - error fetching chats: %v", err)
			return nil, err
		}
		for _, ch := range mChats.Datas {
			chat := projectMutableChat(ch, c.MD.UserId)
			if chat != nil {
				chats = append(chats, chat)
			}
		}
	}

	return tg.MakeTLUpdates(&tg.TLUpdates{
		Updates: updates,
		Users:   users,
		Chats:   chats,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).ToUpdates(), nil
}

func makeNotifyPeer(peerType int32, peerId int64) tg.NotifyPeerClazz {
	switch peerType {
	case tg.PEER_USER, tg.PEER_CHAT, tg.PEER_CHANNEL, tg.PEER_SELF:
		pu := tg.MakePeerUtilHelper(peerType, peerId).ToPeerUtil()
		return tg.MakeTLNotifyPeer(&tg.TLNotifyPeer{
			Peer: pu.ToPeer(),
		})
	case tg.PEER_USERS:
		return tg.MakeTLNotifyUsers(&tg.TLNotifyUsers{})
	case tg.PEER_CHATS:
		return tg.MakeTLNotifyChats(&tg.TLNotifyChats{})
	case tg.PEER_BROADCASTS:
		return tg.MakeTLNotifyBroadcasts(&tg.TLNotifyBroadcasts{})
	default:
		pu := tg.MakePeerUtilHelper(peerType, peerId).ToPeerUtil()
		return tg.MakeTLNotifyPeer(&tg.TLNotifyPeer{
			Peer: pu.ToPeer(),
		})
	}
}
