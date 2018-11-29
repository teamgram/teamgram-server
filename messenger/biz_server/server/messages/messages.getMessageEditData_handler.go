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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
)

// messages.getMessageEditData#fda68d36 peer:InputPeer id:int = messages.MessageEditData;
func (s *MessagesServiceImpl) MessagesGetMessageEditData(ctx context.Context, request *mtproto.TLMessagesGetMessageEditData) (*mtproto.Messages_MessageEditData, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getMessageEditData#fda68d36 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	peer := base.FromInputPeer2(md.UserId, request.GetPeer())
	editData := &mtproto.TLMessagesMessageEditData{Data2: &mtproto.Messages_MessageEditData_Data{}}
	edit, err := s.MessageModel.GetMessageBox2(peer.PeerType, md.UserId, request.GetId())
	if err == nil {
		// editData := &mtproto.TLMessagesMessageEditData{Data2: &mtproto.Messages_MessageEditData_Data{}}
		//editData.SetCaption(edit.Message.GetData2().e)
		_ = edit
	}
	glog.Infof("messages.getMessageEditData#fda68d36 - reply: %s", logger.JsonDebugData(editData))
	return editData.To_Messages_MessageEditData(), nil
}
