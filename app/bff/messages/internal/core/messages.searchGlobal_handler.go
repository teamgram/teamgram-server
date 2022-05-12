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
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	messagepb "github.com/teamgram/teamgram-server/app/service/biz/message/message"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"math"
)

// MessagesSearchGlobal
// messages.searchGlobal#4bc6589a flags:# folder_id:flags.0?int q:string filter:MessagesFilter min_date:int max_date:int offset_rate:int offset_peer:InputPeer offset_id:int limit:int = messages.Messages;
func (c *MessagesCore) MessagesSearchGlobal(in *mtproto.TLMessagesSearchGlobal) (*mtproto.Messages_Messages, error) {
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	SEARCH_QUERY_EMPTY	The search query is empty
	if c.MD.IsBot {
		err := mtproto.ErrBotMethodInvalid
		c.Logger.Errorf("messages.searchGlobal - error: %v", err)
		return nil, err
	}

	if in.Q == "" {
		err := mtproto.ErrSearchQueryEmpty
		c.Logger.Errorf("messages.searchGlobal - error: %v", err)
		return nil, err
	}

	var (
		offsetId = in.OffsetId
		limit    = in.Limit
	)

	if offsetId == 0 {
		offsetId = math.MaxInt32
	}

	if limit > 50 {
		limit = 50
	}

	// TODO(@benqi): Impl MessagesSearchGlobal logic
	rValues := mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_Messages()

	boxList, err := c.svcCtx.Dao.MessageClient.MessageSearchGlobal(
		c.ctx,
		&messagepb.TLMessageSearchGlobal{
			UserId: c.MD.UserId,
			Q:      in.Q,
			Offset: offsetId,
			Limit:  limit,
		})
	if err != nil {
		c.Logger.Errorf("messages.searchGlobal - error: %v", err)
		return rValues, nil
	}

	boxList.Visit(c.MD.UserId,
		func(messageList []*mtproto.Message) {
			rValues.Messages = messageList
		},
		func(userIdList []int64) {
			mUsers, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			rValues.Users = append(rValues.Users, mUsers.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			mChats, _ := c.svcCtx.Dao.ChatClient.Client().ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			rValues.Chats = append(rValues.Chats, mChats.GetChatListByIdList(c.MD.UserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			//mChannels, _ := c.svcCtx.Dao.ChannelClient.ChannelGetChannelListByIdList(c.ctx,
			//	&channelpb.TLChannelGetChannelListByIdList{
			//		SelfUserId: c.MD.UserId,
			//		Id:         channelIdList,
			//	})
			//if len(mChannels.GetDatas()) > 0 {
			//	rValues.Chats = append(rValues.Chats, mChannels.GetDatas()...)
			//}
		})

	return rValues, nil
}
