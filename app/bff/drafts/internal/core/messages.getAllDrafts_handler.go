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
	"github.com/teamgram/teamgram-server/v2/app/bff/drafts/internal/repository"
	chatprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/chatprojection"
	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetAllDrafts
// messages.getAllDrafts#6a3f8d65 = Updates;
func (c *DraftsCore) MessagesGetAllDrafts(in *tg.TLMessagesGetAllDrafts) (*tg.Updates, error) {
	_ = in

	dialogClient, err := c.dialogClient()
	if err != nil {
		return nil, err
	}

	drafts, err := dialogClient.DialogGetAllDrafts(c.ctx, &repository.DialogGetAllDrafts{
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

		switch peer := v.Peer.(type) {
		case *tg.TLPeerUser:
			userIdList = append(userIdList, peer.UserId)
		case *tg.TLPeerChat:
			chatIdList = append(chatIdList, peer.ChatId)
		case *tg.TLPeerChannel:
			// TODO: channel plugin required (enterprise feature)
		}
	}

	if len(userIdList) > 0 {
		users, err := userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, c.MD.UserId, userIdList, userprojection.MissingStoredReference)
		if err != nil {
			c.Logger.Errorf("messages.getAllDrafts - user.getUserProjectionBundle error: %v", err)
		}
		rUpdates.Users = append(rUpdates.Users, users...)
	}

	if len(chatIdList) > 0 {
		chats, err := chatprojection.ProjectChats(c.ctx, c.svcCtx.Repo.ChatClient, c.MD.UserId, chatIdList, chatprojection.MissingStoredReference)
		if err != nil {
			c.Logger.Errorf("messages.getAllDrafts - chat.getChatProjectionBundle error: %v", err)
		} else {
			rUpdates.Chats = append(rUpdates.Chats, chats...)
		}
	}

	return rUpdates.ToUpdates(), nil
}
