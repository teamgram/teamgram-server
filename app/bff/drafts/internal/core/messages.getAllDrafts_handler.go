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
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetAllDrafts
// messages.getAllDrafts#6a3f8d65 = Updates;
func (c *DraftsCore) MessagesGetAllDrafts(in *tg.TLMessagesGetAllDrafts) (*tg.Updates, error) {
	_ = in

	drafts, err := c.svcCtx.Repo.DialogClient.DialogGetAllDrafts(c.ctx, &repository.DialogGetAllDrafts{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("messages.getAllDrafts - error: %v", err)
		return nil, err
	}

	rUpdates := tg.MakeTLUpdates(&tg.TLUpdates{})

	var (
		userIdList []int64
		chatIdList []int64
	)

	for _, v := range drafts.Datas {
		rUpdates.Updates = append(rUpdates.Updates, tg.MakeTLUpdateDraftMessage(&tg.TLUpdateDraftMessage{
			Peer:  v.Peer,
			Draft: v.Draft,
		}))

		peerUtil := tg.FromPeer(v.Peer)
		switch peerUtil.PeerType {
		case tg.PEER_SELF, tg.PEER_USER:
			userIdList = append(userIdList, peerUtil.PeerId)
		case tg.PEER_CHAT:
			chatIdList = append(chatIdList, peerUtil.PeerId)
		case tg.PEER_CHANNEL:
			// TODO: channel plugin required (enterprise feature)
		}
	}

	if len(userIdList) > 0 {
		mutableUsers, _ := c.svcCtx.Repo.UserClient.UserGetMutableUsersV2(c.ctx,
			&userpb.TLUserGetMutableUsersV2{
				Id: userIdList,
			})
		if mutableUsers != nil {
			for _, u := range mutableUsers.Users {
				user := projectImmutableUser(u)
				if user != nil {
					rUpdates.Users = append(rUpdates.Users, user)
				}
			}
		}
	}

	if len(chatIdList) > 0 {
		chats, _ := c.svcCtx.Repo.ChatClient.ChatGetChatListByIdList(c.ctx,
			&chatpb.TLChatGetChatListByIdList{
				SelfId: c.MD.UserId,
				IdList: chatIdList,
			})
		if chats != nil {
			for _, ch := range chats.Datas {
				if chat := projectMutableChat(ch, c.MD.UserId); chat != nil {
					rUpdates.Chats = append(rUpdates.Chats, chat)
				}
			}
		}
	}

	return rUpdates.ToUpdates(), nil
}
