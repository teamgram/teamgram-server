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
	"fmt"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
)

// messages.exportChatInvite#7d885289 chat_id:int = ExportedChatInvite;
func (s *MessagesServiceImpl) MessagesExportChatInvite(ctx context.Context, request *mtproto.TLMessagesExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.exportChatInvite#7d885289 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl MessagesExportChatInvite logic
	//chatLogic, err := s.ChatModel.NewChatLogicById(request.GetChatId())
	//if err != nil {
	//}

	//exportedChatInvite := &mtproto.TLChatInviteExported{Data2: &mtproto.ExportedChatInvite_Data{
	//	// Link: chatLogic.ExportedChatInvite(),
	//}}

	// glog.Infof("messages.exportChatInvite#7d885289 - reply: %s", logger.JsonDebugData(syncUpdates))
	return nil, fmt.Errorf("not impl MessagesExportChatInvite")
}
