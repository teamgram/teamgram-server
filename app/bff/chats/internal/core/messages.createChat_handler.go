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
	"math/rand"

	"github.com/teamgram/proto/mtproto"
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

func (c *ChatsCore) createChat(iUsers []*mtproto.InputUser, chatTitle string, ttlPeriod int32) (*mtproto.Updates, error) {
	var (
		chatUserIdList []int64
		userAddList    = make([]int64, 0)
		botAddList     = make([]int64, 0)
	)

	// check chat title
	if chatTitle == "" {
		err := mtproto.ErrChatTitleEmpty
		c.Logger.Errorf("messages.createChat - error: %v", err)
		return nil, err
	}

	if len(iUsers) == 0 {
		err := mtproto.ErrUsersTooFew
		c.Logger.Errorf("messages.createChat - error: %v", err)
		return nil, err
	}

	// check user too much
	if len(iUsers) > 200-1 {
		err := mtproto.ErrUsersTooMuch
		c.Logger.Errorf("messages.createChat - error: %v", err)
		return nil, err
	}

	// check len(users)
	chatUserIdList = []int64{c.MD.UserId}
	for _, u := range iUsers {
		if u.PredicateName != mtproto.Predicate_inputUser {
			err := mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("messages.createChat - error: %v", err)
			return nil, err
		} else {
			chatUserIdList = append(chatUserIdList, u.UserId)
		}
	}

	users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: chatUserIdList,
	})

	if me, _ := users.GetImmutableUser(c.MD.UserId); me.Restricted() {
		err := mtproto.ErrUserRestricted
		c.Logger.Errorf("messages.createChat - error: %v", err)
		return nil, err
	}

	for _, u := range iUsers {
		if addUser, ok := users.GetImmutableUser(u.UserId); !ok {
			err := mtproto.ErrInputUserDeactivated
			c.Logger.Errorf("messages.createChat - error: %v", err)
			return nil, err
		} else {
			if addUser.IsBot() {
				if !addUser.BotNochats() {
					c.Logger.Errorf("user is bot and nochats, ignore %d", u.UserId)
					continue
				} else {
					botAddList = append(botAddList, addUser.Id())
				}
			} else {
				rules, _ := c.svcCtx.Dao.UserClient.UserGetPrivacy(c.ctx, &userpb.TLUserGetPrivacy{
					UserId:  addUser.Id(),
					KeyType: mtproto.CHAT_INVITE,
				})
				if len(rules.Datas) > 0 {
					allowAddChat := mtproto.CheckPrivacyIsAllow(
						addUser.Id(),
						rules.Datas,
						c.MD.UserId,
						func(id, checkId int64) bool {
							contact, _ := c.svcCtx.Dao.UserClient.UserCheckContact(c.ctx, &userpb.TLUserCheckContact{
								UserId: id,
								Id:     checkId,
							})
							return mtproto.FromBool(contact)
						},
						func(checkId int64, idList []int64) bool {
							chatIdList, _ := mtproto.SplitChatAndChannelIdList(idList)
							return c.svcCtx.Dao.ChatClient.CheckParticipantIsExist(c.ctx, checkId, chatIdList)
						})
					if !allowAddChat {
						c.Logger.Errorf("chatInvite privacy, ignore %d", u.UserId)
						continue
					}
				}
				userAddList = append(userAddList, addUser.Id())
			}
		}
	}

	if len(userAddList) == 0 {
		err := mtproto.ErrUsersTooFew
		c.Logger.Errorf("messages.createChat - error: %v", err)
		return nil, err
	}

	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatCreateChat2(c.ctx, &chatpb.TLChatCreateChat2{
		CreatorId:  c.MD.UserId,
		UserIdList: userAddList,
		Title:      chatTitle,
		Bots:       botAddList,
	})
	if err != nil {
		c.Logger.Errorf("createChat duplicate: %v", err)
		return nil, err
	}

	// TODO: add attach_data (chat and chat_participants)
	rValue, err := c.svcCtx.Dao.MsgClient.MsgSendMessageV2(c.ctx, &msgpb.TLMsgSendMessageV2{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.PermAuthKeyId,
		PeerType:  mtproto.PEER_CHAT,
		PeerId:    chat.Chat.Id,
		Message: []*msgpb.OutboxMessage{
			msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
				NoWebpage:    true,
				Background:   false,
				RandomId:     rand.Int63(),
				Message:      chat.MakeMessageService(c.MD.UserId, mtproto.MakeMessageActionChatCreate(chatTitle, append(userAddList, botAddList...))),
				ScheduleDate: nil,
			}).To_OutboxMessage(),
		},
	})
	if err != nil {
		c.Logger.Errorf("messages.createChat - error: %v", err)
		return nil, err
	}

	return rValue, nil
}
