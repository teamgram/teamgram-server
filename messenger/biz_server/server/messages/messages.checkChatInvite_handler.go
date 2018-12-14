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
    "golang.org/x/net/context"
    "github.com/nebula-chat/chatengine/pkg/grpc_util"
    "github.com/nebula-chat/chatengine/pkg/logger"
    "github.com/nebula-chat/chatengine/mtproto"
)

// messages.checkChatInvite#3eadb1bb hash:string = ChatInvite;
func (s *MessagesServiceImpl) MessagesCheckChatInvite(ctx context.Context, request *mtproto.TLMessagesCheckChatInvite) (*mtproto.ChatInvite, error) {
    md := grpc_util.RpcMetadataFromIncoming(ctx)
    glog.Infof("messages.checkChatInvite#3eadb1bb - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

    var (
        chatInvite *mtproto.ChatInvite
    )

    chatLogic, err := s.ChatModel.NewChatLogicByLink(request.GetHash())
    if err != nil {
        glog.Error("messages.checkChatInvite#3eadb1bb - makeChatByLink error: ", err)
    }

    chatInvite = chatLogic.ToChatInvite(md.UserId, func(idList []int32) []*mtproto.User {
        return s.UserModel.GetUserListByIdList(md.UserId, idList)
    })
    glog.Infof("messages.checkChatInvite#3eadb1bb - reply: {%s}", logger.JsonDebugData(chatInvite))
    return chatInvite, nil
}
