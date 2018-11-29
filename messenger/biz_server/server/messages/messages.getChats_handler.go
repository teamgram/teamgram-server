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
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

// messages.getChats#3c6aa187 id:Vector<int> = messages.Chats;
func (s *MessagesServiceImpl) MessagesGetChats(ctx context.Context, request *mtproto.TLMessagesGetChats) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getChats#3c6aa187 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): messages.chatsSlice
	chats := &mtproto.TLMessagesChats{Data2: &mtproto.Messages_Chats_Data{
		Chats: s.ChatModel.GetChatListBySelfAndIDList(md.UserId, request.GetId()),
	}}

	glog.Infof("messages.getChats#3c6aa187 - reply: %s", chats)
	return chats.To_Messages_Chats(), nil
}
