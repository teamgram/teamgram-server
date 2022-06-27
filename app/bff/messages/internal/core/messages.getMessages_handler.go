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
	"github.com/teamgram/marmota/pkg/container2/linkedmap"
	"github.com/teamgram/proto/mtproto"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MessagesGetMessages
// messages.getMessages#63c66506 id:Vector<InputMessage> = messages.Messages;
func (c *MessagesCore) MessagesGetMessages(in *mtproto.TLMessagesGetMessages) (*mtproto.Messages_Messages, error) {
	var (
		idList    []int32
		rMessages = linkedmap.New()
		rValues   = mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
			Messages: []*mtproto.Message{},
			Users:    []*mtproto.User{},
			Chats:    []*mtproto.Chat{},
		}).To_Messages_Messages()
	)

	for _, id := range in.GetId_VECTORINPUTMESSAGE() {
		switch id.PredicateName {
		case mtproto.Predicate_inputMessageID:
			idList = append(idList, id.Id)
		default:
			// client not use: inputMessageReplyTo, inputMessagePinned
			err := mtproto.ErrInputConstructorInvalid
			c.Logger.Errorf("messages.getMessages - error: %v", err)
			return nil, err
		}
	}
	if len(in.GetId_VECTORINT32()) > 0 {
		idList = append(idList, in.Id_VECTORINT32...)
	}

	if len(idList) == 0 {
		return rValues, nil
	} else {
		for _, id := range idList {
			rMessages.Add(id, mtproto.MakeTLMessageEmpty(&mtproto.Message{
				Id:     id,
				PeerId: nil,
			}).To_Message())
		}
	}

	boxList, _ := c.svcCtx.Dao.MessageClient.MessageGetUserMessageList(
		c.ctx,
		&message.TLMessageGetUserMessageList{
			UserId: c.MD.UserId,
			IdList: idList,
		})

	boxList.Visit(c.MD.UserId,
		func(messageList []*mtproto.Message) {
			for _, msg := range messageList {
				rMessages.Add(msg.Id, msg)
			}
			for i := rMessages.First(); i != nil; i = i.Next() {
				rValues.Messages = append(rValues.Messages, i.Value().(*mtproto.Message))
			}
		},
		func(userIdList []int64) {
			mUsers, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(
				c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			rValues.Users = append(rValues.Users, mUsers.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			mChats, _ := c.svcCtx.Dao.ChatClient.Client().ChatGetChatListByIdList(
				c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			rValues.Chats = append(rValues.Chats, mChats.GetChatListByIdList(c.MD.UserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			//mChannels, _ := c.svcCtx.Dao.ChannelClient.ChannelGetChannelListByIdList(
			//	c.ctx,
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
