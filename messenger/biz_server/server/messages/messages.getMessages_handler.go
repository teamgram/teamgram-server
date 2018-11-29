// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package messages

import (
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/message"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

// messages.getMessages#63c66506 id:Vector<InputMessage> = messages.Messages;
func (s *MessagesServiceImpl) MessagesGetMessages(ctx context.Context, request *mtproto.TLMessagesGetMessages) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getMessages#63c66506 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var idList = make([]int32, len(request.Id))

	// TODO(@benqi): Read client source code.
	for _, id := range request.GetId() {
		switch id.GetConstructor() {
		case mtproto.TLConstructor_CRC32_inputMessageID:
			idList = append(idList, id.GetData2().GetId())
		case mtproto.TLConstructor_CRC32_inputMessageReplyTo:
			idList = append(idList, id.GetData2().GetId())
		case mtproto.TLConstructor_CRC32_inputMessagePinned:
			// idList = append(idList, id.GetData2().GetId())
		}
	}
	messages := s.MessageModel.GetUserMessagesByMessageIdList(md.UserId, idList)
	userIdList, chatIdList, _ := message.PickAllIDListByMessages(messages)
	userList := s.UserModel.GetUserListByIdList(md.UserId, userIdList)
	chatList := s.ChatModel.GetChatListBySelfAndIDList(md.UserId, chatIdList)

	messagesMessages := &mtproto.TLMessagesMessages{Data2: &mtproto.Messages_Messages_Data{
		Messages: messages,
		Users:    userList,
		Chats:    chatList,
	}}

	glog.Infof("messages.getMessages#63c66506 - reply: %s", logger.JsonDebugData(messagesMessages))
	return messagesMessages.To_Messages_Messages(), nil
}
