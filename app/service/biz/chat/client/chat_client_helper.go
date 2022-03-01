// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package chat_client

import (
	"context"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"

	"github.com/teamgram/proto/mtproto"
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

	return chatList.GetChatListByIdList(selfId, id...)
}
