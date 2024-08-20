// Copyright 2024 Teamgram Authors
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
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MessagesGetSavedHistory
// messages.getSavedHistory#3d9a414d peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (c *SavedMessageDialogsCore) MessagesGetSavedHistory(in *mtproto.TLMessagesGetSavedHistory) (*mtproto.Messages_Messages, error) {
	// TODO(@benqi): 重复FromInputPeer2
	var (
		err   error
		peer  = mtproto.FromInputPeer2(c.MD.UserId, in.GetPeer())
		limit = in.Limit
	)

	if limit > 50 {
		limit = 50
	}

	var (
		boxList *mtproto.MessageBoxList
	)

	boxList, err = c.svcCtx.Dao.MessageClient.MessageGetSavedHistoryMessages(c.ctx, &message.TLMessageGetSavedHistoryMessages{
		UserId:     c.MD.UserId,
		PeerType:   peer.PeerType,
		PeerId:     peer.PeerId,
		OffsetId:   in.OffsetId,
		OffsetDate: in.OffsetDate,
		AddOffset:  in.AddOffset,
		Limit:      limit,
		MaxId:      in.MaxId,
		MinId:      in.MinId,
		Hash:       in.Hash,
	})
	if err != nil {
		c.Logger.Errorf("messages.getHistory - error: %v", err)
		return nil, err
	}

	var (
		messages []*mtproto.Message
		users    []*mtproto.User
		chats    []*mtproto.Chat
	)
	boxList.Visit(c.MD.UserId,
		func(messageList []*mtproto.Message) {
			messages = messageList
		},
		func(userIdList []int64) {
			mUsers, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: userIdList,
				})
			users = append(users, mUsers.GetUserListByIdList(c.MD.UserId, userIdList...)...)
		},
		func(chatIdList []int64) {
			mChats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			chats = append(chats, mChats.GetChatListByIdList(c.MD.UserId, chatIdList...)...)
		},
		func(channelIdList []int64) {
			//// TODO: handler other...
			//if channel != nil {
			//	chats = append(chats, channel.ToUnsafeChat(c.MD.UserId))
			//}
		})

	var (
		rValues *mtproto.Messages_Messages
	)

	//if !isChannel {
	if boxList.Length() == limit {
		rValues = mtproto.MakeTLMessagesMessagesSlice(&mtproto.Messages_Messages{
			Inexact:        false, // TODO: ???
			Count:          boxList.Count,
			NextRate:       nil, // TODO: ???
			OffsetIdOffset: nil, // TODO: ???
			Messages:       messages,
			Users:          mtproto.ToSafeUsers(users),
			Chats:          mtproto.ToSafeChats(chats),
		}).To_Messages_Messages()
	} else {
		rValues = mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
			Messages: messages,
			Users:    mtproto.ToSafeUsers(users),
			Chats:    mtproto.ToSafeChats(chats),
		}).To_Messages_Messages()
	}

	return rValues, nil
}
