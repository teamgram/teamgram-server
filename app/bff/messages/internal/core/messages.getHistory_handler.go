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
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"github.com/zeromicro/go-zero/core/mr"
)

// MessagesGetHistory
// messages.getHistory#4423e6c5 peer:InputPeer offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:long = messages.Messages;
func (c *MessagesCore) MessagesGetHistory(in *mtproto.TLMessagesGetHistory) (*mtproto.Messages_Messages, error) {
	// TODO(@benqi): 重复FromInputPeer2
	var (
		err  error
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.GetPeer())
		chat *mtproto.MutableChat
		//channel   *channelpb.MutableChannel
		//isChannel bool
		limit = in.Limit
	)

	if limit > 50 {
		limit = 50
	}

	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		if peer.PeerType == mtproto.PEER_CHAT {
			// 400	CHAT_ID_INVALID	The provided chat id is invalid
			if chat, err = c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
				ChatId: peer.PeerId,
			}); err != nil {
				err = mtproto.ErrPeerIdInvalid
				c.Logger.Errorf("messages.getHistory - error: %v", err)
				return nil, err
			}
			// TODO(@benqi): check migratedToId
			_ = chat
		}
	case mtproto.PEER_CHANNEL:
		//// 400	CHANNEL_INVALID	The provided channel is invalid
		//// 400	CHANNEL_PRIVATE	You haven't joined this channel/supergroup
		//channel, err = c.svcCtx.Dao.ChannelClient.ChannelGetMutableChannel(
		//	c.ctx,
		//	&channelpb.TLChannelGetMutableChannel{
		//		ChannelId: peer.PeerId,
		//		Id:        []int64{c.MD.UserId},
		//	})
		//if err != nil {
		//	err = mtproto.ErrChannelInvalid
		//	c.Logger.Errorf("messages.getHistory - error: %v", err)
		//	return nil, err
		//}
		//
		//me := channel.GetImmutableChannelParticipant(c.MD.UserId)
		//if me != nil {
		//	// TODO
		//	// minId = mathx.MaxInt(int(me.AvailableMinId), int(in.MinId))
		//}
		//
		//isChannel = true
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getHistory - error: %v", err)
		return nil, err
	}

	var (
		boxList *message.Vector_MessageBox
		count   *mtproto.Int32
	)

	err = mr.Finish(
		func() error {
			var err2 error
			boxList, err2 = c.svcCtx.Dao.MessageClient.MessageGetHistoryMessages(c.ctx, &message.TLMessageGetHistoryMessages{
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
			if err2 != nil {
				c.Logger.Errorf("messages.getHistory - error: %v", err2)
			}

			return err2
		},
		func() error {
			var err2 error
			count, err2 = c.svcCtx.Dao.MessageClient.MessageGetHistoryMessagesCount(
				c.ctx,
				&message.TLMessageGetHistoryMessagesCount{
					UserId:   c.MD.UserId,
					PeerType: peer.PeerType,
					PeerId:   peer.PeerId,
				})
			if err2 != nil {
				c.Logger.Errorf("messages.getHistory - error: %v", err2)
			}

			return err2
		})
	if err != nil {
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
			mChats, _ := c.svcCtx.Dao.ChatClient.Client().ChatGetChatListByIdList(c.ctx,
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
			Count:          count.V,
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
	//} else {
	//	rValues = mtproto.MakeTLMessagesChannelMessages(&mtproto.Messages_Messages{
	//		Inexact:        false, // TODO: ???
	//		Pts:            channel.Pts(),
	//		Count:          count.V,
	//		OffsetIdOffset: &wrapperspb.Int32Value{Value: channel.TopMessage() - in.OffsetId}, // TODO: ???
	//		Messages:       messages,
	//		Chats:          mtproto.ToSafeChats(chats),
	//		Users:          mtproto.ToSafeUsers(users),
	//	}).To_Messages_Messages()
	//}

	return rValues, nil
}
