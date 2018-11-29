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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"golang.org/x/net/context"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"time"
)

// messages.readHistory#e306d3a peer:InputPeer max_id:int = messages.AffectedMessages;
func (s *MessagesServiceImpl) MessagesReadHistory(ctx context.Context, request *mtproto.TLMessagesReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.readHistory#e306d3a - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		updates *mtproto.TLUpdates
		pts, ptsCount int32
	)

	peer := base.FromInputPeer(request.GetPeer())
	if peer.PeerType == base.PEER_SELF {
		// TODO(@benqi): 太土！
		peer.PeerType = base.PEER_USER
		peer.PeerId = md.UserId
	}

	// 消息已读逻辑
	// 1. inbox，设置unread_count为0以及read_inbox_max_id
	s.DialogModel.UpdateUnreadByPeer(md.UserId, int8(peer.PeerType), peer.PeerId, request.GetMaxId())

	pts = int32(core.NextPtsId(md.UserId))
	ptsCount = 1

	updateReadHistoryInbox := &mtproto.TLUpdateReadHistoryInbox{Data2: &mtproto.Update_Data{
		Peer_39:  peer.ToPeer(),
		MaxId:    request.MaxId,
		Pts:      pts,
		PtsCount: ptsCount,
	}}
	updates = &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{updateReadHistoryInbox.To_Update()},
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

	// result
	affected := &mtproto.TLMessagesAffectedMessages{Data2: &mtproto.Messages_AffectedMessages_Data{
		Pts:      pts,
		PtsCount: ptsCount,
	}}

	// 2. outbox, 设置read_outbox_max_id
	if peer.PeerType == base.PEER_USER {
		outboxTopMessage := s.DialogModel.GetTopMessage(peer.PeerId, int8(peer.PeerType), md.UserId)
		s.DialogModel.UpdateReadOutboxMaxIdByPeer(peer.PeerId, int8(peer.PeerType), md.UserId, outboxTopMessage)

		outboxPeer := &mtproto.TLPeerUser{Data2: &mtproto.Peer_Data{
			UserId: md.UserId,
		}}
		updateReadHistoryOutbox := &mtproto.TLUpdateReadHistoryOutbox{Data2: &mtproto.Update_Data{
			Peer_39:  outboxPeer.To_Peer(),
			MaxId:    outboxTopMessage,
			Pts:      int32(core.NextPtsId(peer.PeerId)),
			PtsCount: 1,
		}}

		updates = &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
			Updates: []*mtproto.Update{updateReadHistoryOutbox.To_Update()},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}}
		sync_client.GetSyncClient().PushUpdates(peer.PeerId, updates.To_Updates())
	} else {
		chatLogic, _ := s.ChatModel.NewChatLogicById(peer.PeerId)
		chatParticipants := chatLogic.GetChatParticipantList()
		for _, participant := range chatParticipants {
			if participant.GetData2().GetUserId() == md.UserId {
				continue
			}
			outboxTopMessage := s.DialogModel.GetTopMessage(participant.GetData2().GetUserId(), int8(peer.PeerType), peer.PeerId)
			s.DialogModel.UpdateReadOutboxMaxIdByPeer(participant.GetData2().GetUserId(), int8(peer.PeerType), peer.PeerId, outboxTopMessage)

			outboxPeer := &mtproto.TLPeerChat{Data2: &mtproto.Peer_Data{
				ChatId: peer.PeerId,
			}}
			updateReadHistoryOutbox := &mtproto.TLUpdateReadHistoryOutbox{Data2: &mtproto.Update_Data{
				Peer_39:  outboxPeer.To_Peer(),
				MaxId:    outboxTopMessage,
				Pts:      int32(core.NextPtsId(participant.GetData2().GetUserId())),
				PtsCount: 1,
			}}

			updates = &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
				Updates: []*mtproto.Update{updateReadHistoryOutbox.To_Update()},
				Users:   []*mtproto.User{},
				Chats:   []*mtproto.Chat{},
				Date:    int32(time.Now().Unix()),
				Seq:     0,
			}}
			sync_client.GetSyncClient().PushUpdates(participant.GetData2().GetUserId(), updates.To_Updates())
		}
	}

	glog.Infof("messages.readHistory#e306d3a - reply: {%s}", logger.JsonDebugData(affected))
	return affected.To_Messages_AffectedMessages(), err
}
