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
	"math"

	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// messages.getDialogs#191ba9c5 flags:# exclude_pinned:flags.0?true offset_date:int offset_id:int offset_peer:InputPeer limit:int = messages.Dialogs;
func (s *MessagesServiceImpl) MessagesGetDialogsLayer62(ctx context.Context, request *mtproto.TLMessagesGetDialogsLayer62) (*mtproto.Messages_Dialogs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getDialogs#191ba9c5 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	offsetId := request.GetOffsetId()
	if offsetId == 0 {
		offsetId = math.MaxInt32
	}

	dialogs := s.DialogModel.GetDialogsByOffsetId(md.UserId, false, offsetId, request.GetLimit())
	// glog.Infof("dialogs - {%v}", dialogs)

	// messageIdList, userIdList, chatIdList, channelIdList
	dialogItems := s.DialogModel.PickAllIDListByDialogs(dialogs)
	glog.Info(dialogItems)
	messages := s.MessageModel.GetUserMessagesByMessageIdList(md.UserId, dialogItems.MessageIdList)

	// TODO(@benqi): add channel message.
	for k, v := range dialogItems.ChannelMessageIdMap {
		m := s.MessageModel.GetChannelMessage(md.UserId, k, v)
		if m != nil {
			messages = append(messages, m)
		}
	}

	users := s.UserModel.GetUserListByIdList(md.UserId, dialogItems.UserIdList)
	chats := s.ChatModel.GetChatListBySelfAndIDList(md.UserId, dialogItems.ChatIdList)
	chats = append(chats, s.ChannelModel.GetChannelListBySelfAndIDList(md.UserId, dialogItems.ChannelIdList)...)

	messageDialogs := mtproto.TLMessagesDialogs{Data2: &mtproto.Messages_Dialogs_Data{
		Dialogs:  dialogs,
		Messages: messages,
		Users:    users,
		Chats:    chats,
	}}

	glog.Infof("messages.getDialogs#191ba9c5 - reply: %s", logger.JsonDebugData(messageDialogs))
	return messageDialogs.To_Messages_Dialogs(), nil
}
