// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package chat_client

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"

	"github.com/zeromicro/go-zero/zrpc"
)

type ChatClientHelper struct {
	cli ChatClient
}

func NewChatClientHelper(cli zrpc.Client) *ChatClientHelper {
	return &ChatClientHelper{
		cli: NewChatClient(cli),
	}
}

func (m *ChatClientHelper) Client() ChatClient {
	return m.cli
}

func (m *ChatClientHelper) GetChatListByIdList(ctx context.Context, selfId int64, id ...int64) []*mtproto.Chat {
	chatList, _ := m.cli.ChatGetChatListByIdList(ctx, &chat.TLChatGetChatListByIdList{
		SelfId: selfId,
		IdList: id,
	})
	if chatList == nil {
		return []*mtproto.Chat{}
	}

	return chatList.GetChatListByIdList(selfId, id...)
}

func (m *ChatClientHelper) CheckParticipantIsExist(ctx context.Context, userId int64, chatIdList []int64) bool {
	usersChatIdList, _ := m.cli.ChatGetUsersChatIdList(ctx, &chat.TLChatGetUsersChatIdList{
		Id: []int64{userId},
	})

	if len(usersChatIdList.GetDatas()) == 0 {
		return false
	}

	for _, v := range usersChatIdList.GetDatas() {
		if v.UserId == userId {
			for _, id := range chatIdList {
				for _, id2 := range v.ChatIdList {
					if id == id2 {
						return true
					}
				}
			}
		}
	}

	return false
}
