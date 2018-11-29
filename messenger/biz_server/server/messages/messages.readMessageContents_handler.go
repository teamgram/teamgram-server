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
	// "time"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"time"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
)

// messages.readMessageContents#36a73f77 id:Vector<int> = messages.AffectedMessages;
func (s *MessagesServiceImpl) MessagesReadMessageContents(ctx context.Context, request *mtproto.TLMessagesReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.readMessageContents#36a73f77 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		updates *mtproto.TLUpdates
		pts, ptsCount int32
	)

	messages := s.MessageModel.GetUserMessagesByMessageIdList(md.UserId, request.GetId())
	for _, m := range messages {
		if m.GetData2().GetMentioned() {
			s.MessageModel.UpdateUnreadReadMention(md.UserId,
				int8(base.PEER_CHAT),
				m.GetData2().GetToId().GetData2().GetChatId(),
				m.GetData2().GetId())
		}
	}

	ptsCount = int32(len(request.GetId()))
	pts = int32(core.NextNPtsId(md.UserId, int(ptsCount))) - ptsCount + 1

	updReadMessageContents := &mtproto.TLUpdateReadMessagesContents{Data2: &mtproto.Update_Data{
		Messages: request.GetId(),
		Pts:      pts,
		PtsCount: ptsCount,
	}}

	updates = &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{updReadMessageContents.To_Update()},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}}

	// sync
	_, err := sync_client.GetSyncClient().SyncUpdatesNotMe(md.UserId, md.AuthId, updates.To_Updates())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	// TODO(@benqi): voice and videos

	// result
	affected := &mtproto.TLMessagesAffectedMessages{Data2: &mtproto.Messages_AffectedMessages_Data{
		Pts:      int32(core.NextPtsId(md.UserId)),
		PtsCount: 1,
	}}

	glog.Infof("messages.readMessageContents#36a73f77 - reply: {%s}", logger.JsonDebugData(affected))
	return affected.To_Messages_AffectedMessages(), nil
}
