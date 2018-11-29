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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"time"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
)


/*
 ## just_clear:
 - affectedHistory: pts_count is size(history_messages-1) + updateLastEditMessage(messageActionHistoryClear)
 - sync
	- updateDeleteMessages + updateEditMessage(messageActionHistoryClear)

 ##
 */
// messages.deleteHistory#1c015b09 flags:# just_clear:flags.0?true peer:InputPeer max_id:int = messages.AffectedHistory;
func (s *MessagesServiceImpl) MessagesDeleteHistory(ctx context.Context, request *mtproto.TLMessagesDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.deleteHistory#1c015b09 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// peer
	var (
		peer               *base.PeerUtil
		err                error
		pts, ptsCount      int32
	)

	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerEmpty {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("messages.sendMessage#fa88427a - invalid peer", err)
		return nil, err
	}

	// TODO(@benqi): check user or channels's access_hash
	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerSelf {
		peer = &base.PeerUtil{
			PeerType: base.PEER_USER,
			PeerId:   md.UserId,
		}
	} else {
		peer = base.FromInputPeer(request.GetPeer())
	}

	if request.GetJustClear() {
		lastMessageBox, deleteIds := s.MessageModel.GetClearHistoryMessages(md.UserId, peer)

		if lastMessageBox == nil || len(deleteIds) == 0 {
			pts = int32(core.CurrentPtsId(md.UserId))
			ptsCount = 0
		} else {
			pts = int32(core.NextNPtsId(md.UserId, len(deleteIds)))
			ptsCount = int32(len(deleteIds))

			updateDeleteMessages := &mtproto.TLUpdateDeleteMessages{Data2: &mtproto.Update_Data{
				Messages: deleteIds,
				Pts:      pts,
				PtsCount: ptsCount,
			}}

			clearHistoryMessage := &mtproto.TLMessageService{Data2: &mtproto.Message_Data{
				Out:    true,
				Id:     int32(core.NextMessageBoxId(md.UserId)),
				FromId: md.UserId,
				ToId:   peer.ToPeer(),
				Date:   lastMessageBox.Message.GetData2().GetDate(),
				Action: mtproto.NewTLMessageActionHistoryClear().To_MessageAction(),
			}}
			lastMessageBox.Message = clearHistoryMessage.To_Message()
			lastMessageBox.EditDate = clearHistoryMessage.GetDate()
			lastMessageBox.EditMessage = ""
			lastMessageBox.SaveMessageData()

			pts = int32(core.NextPtsId(md.UserId))
			ptsCount += 1
			updateEditMessage := &mtproto.TLUpdateEditMessage{Data2: &mtproto.Update_Data{
				Message_1: lastMessageBox.ToMessage(md.UserId),
				Pts:       pts,
				PtsCount:  1,
			}}

			syncUpdates := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
				Updates: []*mtproto.Update{updateDeleteMessages.To_Update(), updateEditMessage.To_Update()},
				Users:   s.UserModel.GetUserListByIdList(md.UserId, []int32{md.UserId, peer.PeerId}),
				Chats:   []*mtproto.Chat{},
				Date:    int32(time.Now().Unix()),
				Seq:     0,
			}}

			s.MessageModel.DeleteByMessageIdList(md.UserId, deleteIds)
			sync_client.GetSyncClient().SyncUpdatesNotMe(md.UserId, md.AuthId, syncUpdates.To_Updates())
		}
	} else {
		deleteIds := s.MessageModel.GetMessageIdListByDialog(md.UserId, peer)
		pts = int32(core.NextNPtsId(md.UserId, len(deleteIds)))
		ptsCount = int32(len(deleteIds))

		updateDeleteMessages := &mtproto.TLUpdateDeleteMessages{Data2: &mtproto.Update_Data{
			Messages: deleteIds,
			Pts:      pts,
			PtsCount: ptsCount,
		}}

		syncUpdats := &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
			Updates: []*mtproto.Update{updateDeleteMessages.To_Update()},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}}

		s.MessageModel.DeleteByMessageIdList(md.UserId, deleteIds)
		s.DialogModel.InsertOrUpdateDialog(md.UserId, peer.PeerType, peer.PeerId, 0, false, false)
		sync_client.GetSyncClient().SyncUpdatesNotMe(md.UserId, md.AuthId, syncUpdats.To_Updates())
	}

	affectedHistory := &mtproto.TLMessagesAffectedHistory{Data2: &mtproto.Messages_AffectedHistory_Data{
		Pts:      pts,
		PtsCount: ptsCount,
		Offset:   0,
	}}

	glog.Infof("messages.deleteHistory#1c015b09 - reply: %s", logger.JsonDebugData(affectedHistory))
	return affectedHistory.To_Messages_AffectedHistory(), nil
}
