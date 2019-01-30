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
	update2 "github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/update"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
)

// Ps：enable = true 的时候开关为关；enable = false 的时候标识为开；
// messages.toggleChatAdmins#ec8bd9e1 chat_id:int enabled:Bool = Updates;
func (s *MessagesServiceImpl) MessagesToggleChatAdmins(ctx context.Context, request *mtproto.TLMessagesToggleChatAdmins) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.toggleChatAdmins#ec8bd9e1 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	chatLogic, err := s.ChatModel.NewChatLogicById(request.ChatId)
	if err != nil {
		glog.Error("toggleChatAdmins#ec8bd9e1 - error: ", err)
		return nil, err
	}

	err = chatLogic.ToggleChatAdmins(md.UserId, mtproto.FromBool(request.GetEnabled()))
	if err != nil {
		glog.Error("toggleChatAdmins#ec8bd9e1 - error: ", err)
		return nil, err
	}

	syncUpdates := update2.NewUpdatesLogic(md.UserId)
	updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
		Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
	}}
	syncUpdates.AddUpdate(updateChatParticipants.To_Update())
	syncUpdates.AddChat(chatLogic.ToChat(md.UserId))

	replyUpdates := syncUpdates.ToUpdates()
	sync_client.GetSyncClient().SyncUpdatesMe(ctx, md.UserId, md.SessionId, 0, syncUpdates.ToUpdates())
	//updateChatAdmins := &mtproto.TLUpdateChatAdmins{Data2: &mtproto.Update_Data{
	//	ChatId:  chatLogic.GetChatId(),
	//	Enabled: request.GetEnabled(),
	//	Version: chatLogic.GetVersion(),
	//}}
	//
	//// sync_client.GetSyncClient().PushToUserNotMeUpdateShortData(md.AuthId, md.SessionId, md.UserId, updateChatAdmins.To_Update())
	//
	//idList := chatLogic.GetChatParticipantIdList()
	//for _, id := range idList {
	//	// sync_client.GetSyncClient().PushToUserUpdateShortData(id, updateChatAdmins.To_Update())
	//}

	glog.Infof("messages.toggleChatAdmins#ec8bd9e1 - reply: %s", logger.JsonDebugData(replyUpdates))
	return replyUpdates, nil
}
