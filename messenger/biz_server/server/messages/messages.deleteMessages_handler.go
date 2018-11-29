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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"time"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
)

// messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true id:Vector<int> = messages.AffectedMessages;
func (s *MessagesServiceImpl) MessagesDeleteMessages(ctx context.Context, request *mtproto.TLMessagesDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.deleteMessages#e58e95d2 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		deleteIdList = request.GetId()
		pts, ptsCount int32
	)

	pts = int32(core.NextNPtsId(md.UserId, len(request.GetId())))
	ptsCount = int32(len(request.GetId()))

	s.MessageModel.DeleteByMessageIdList(md.UserId, deleteIdList)

	deleteMessages := &mtproto.TLUpdateDeleteMessages{Data2: &mtproto.Update_Data{
		Messages: deleteIdList,
		Pts:      pts,
		PtsCount: ptsCount,
	}}

	syncDeleteMessagesUpdates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{deleteMessages.To_Update()},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}}

	sync_client.GetSyncClient().SyncUpdatesNotMe(md.UserId, md.AuthId, syncDeleteMessagesUpdates.To_Updates())

	affectedMessages := &mtproto.TLMessagesAffectedMessages{Data2: &mtproto.Messages_AffectedMessages_Data{
		Pts:      pts,
		PtsCount: ptsCount,
	}}

	s.MessageModel.DeleteByMessageIdList(md.UserId, deleteIdList)

	// TODO(@benqi): 更新dialog的TopMessage

	// 1. delete messages
	// 2. updateTopMessage
	if request.GetRevoke() {
		//  消息撤回
		deleteIdListMap := s.MessageModel.GetPeerDialogMessageIdList(md.UserId, request.GetId())
		glog.Info("messages.deleteMessages#e58e95d2 - deleteIdListMap: ", deleteIdListMap)
		for k, v := range deleteIdListMap {
			pts = int32(core.NextNPtsId(k, len(v)))
			ptsCount = int32(len(v))

			deleteMessages := &mtproto.TLUpdateDeleteMessages{Data2: &mtproto.Update_Data{
				Messages: v,
				Pts:      pts,
				PtsCount: ptsCount,
			}}

			pushDeleteMessagesUpdates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
				Updates: []*mtproto.Update{deleteMessages.To_Update()},
				Users:   []*mtproto.User{},
				Chats:   []*mtproto.Chat{},
				Date:    int32(time.Now().Unix()),
				Seq:     0,
			}}

			sync_client.GetSyncClient().PushUpdates(k, pushDeleteMessagesUpdates.To_Updates())
			s.MessageModel.DeleteByMessageIdList(k, v)
		}

		// TODO(@benqi): 更新dialog的TopMessage
	}
	
	glog.Infof("messages.deleteMessages#e58e95d2 - reply: %s", logger.JsonDebugData(affectedMessages))
	return affectedMessages.To_Messages_AffectedMessages(), nil
}
