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
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/proto/mtproto/rpc/metadata"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/internal/svc"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type InboxCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *InboxCore {
	return &InboxCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}

func (c *InboxCore) makeUpdateNewMessageListUpdates(selfUserId int64, boxList ...*mtproto.MessageBox) *mtproto.Updates {
	var (
		messages      = make([]*mtproto.Message, 0, len(boxList))
		updateNewList = make([]*mtproto.Update, 0, len(boxList))
	)

	for _, box := range boxList {
		if box == nil {
			continue
		}
		if box.PeerType == mtproto.PEER_CHANNEL {
			m := box.ToMessage(selfUserId)
			messages = append(messages, m)
			updateNewMessage := mtproto.MakeTLUpdateNewChannelMessage(&mtproto.Update{
				Message_MESSAGE: box.ToMessage(selfUserId),
				Pts_INT32:       box.Pts,
				PtsCount:        box.PtsCount,
			})
			updateNewList = append(updateNewList, updateNewMessage.To_Update())
		} else {
			m := box.ToMessage(selfUserId)
			messages = append(messages, m)
			updateNewMessage := mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
				Message_MESSAGE: box.ToMessage(selfUserId),
				Pts_INT32:       box.Pts,
				PtsCount:        box.PtsCount,
			})
			updateNewList = append(updateNewList, updateNewMessage.To_Update())
		}
	}

	rUpdates := mtproto.MakeUpdatesByUpdates(updateNewList...)

	idHelper := mtproto.NewIDListHelper(selfUserId)
	idHelper.PickByMessages(messages...)
	idHelper.Visit(
		func(userIdList []int64) {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			rUpdates.PushUser(users.GetUserListByIdList(selfUserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			rUpdates.PushChat(chats.GetChatListByIdList(selfUserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			// TODO
		})

	return rUpdates
}
