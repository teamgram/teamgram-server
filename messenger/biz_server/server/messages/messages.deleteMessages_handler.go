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
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"golang.org/x/net/context"
	"time"
)

// messages.deleteMessages#e58e95d2 flags:# revoke:flags.0?true id:Vector<int> = messages.AffectedMessages;
func (s *MessagesServiceImpl) MessagesDeleteMessages(ctx context.Context, request *mtproto.TLMessagesDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.deleteMessages#e58e95d2 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		deleteIdList  = request.GetId()
		pts, ptsCount int32
	)

	type didInfo struct {
		idList  []int32
		didList []int64
	}

	// TODO(@benqi): 更新dialog的TopMessage
	var (
		deletedDialogsMap = map[int32]*didInfo{md.UserId: {idList: request.Id, didList: []int64{}}}
	)

	if request.GetRevoke() {
		//  消息撤回
		deleteIdListMap := s.MessageModel.GetPeerDialogMessageIdList(md.UserId, request.GetId())
		glog.Info("messages.deleteMessages#e58e95d2 - deleteIdListMap: ", deleteIdListMap)
		for k, v := range deleteIdListMap {
			pts = int32(core.NextNPtsId(k, len(v)))
			ptsCount = int32(len(v))
			var idList []int32
			var didList []int64

			for _, id2 := range v {
				idList = append(idList, id2.A)
				didList = append(didList, id2.B)
				//
				deletedDialogsMap[md.UserId].didList = append(deletedDialogsMap[md.UserId].didList, id2.B)
			}

			deletedDialogsMap[k] = &didInfo{idList: idList, didList: didList}
			deleteMessages := &mtproto.TLUpdateDeleteMessages{Data2: &mtproto.Update_Data{
				Messages: idList,
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
			s.MessageModel.DeleteByMessageIdList(k, idList)
		}

		// TODO(@benqi): 更新dialog的TopMessage
	} else {
		dialogs := s.MessageModel.GetDialogListMessageIdList(md.UserId, request.Id)
		for k, v := range dialogs {
			deletedDialogsMap[md.UserId].didList = append(deletedDialogsMap[md.UserId].didList, k)
			deletedDialogsMap[md.UserId].idList = v
		}
	}

	// s.MessageModel.DeleteByMessageIdList(md.UserId, deleteIdList)
	pts = int32(core.NextNPtsId(md.UserId, len(request.GetId())))
	ptsCount = int32(len(request.GetId()))

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

	s.MessageModel.DeleteByMessageIdList(md.UserId, deleteIdList)
	sync_client.GetSyncClient().SyncUpdatesNotMe(md.UserId, md.AuthId, syncDeleteMessagesUpdates.To_Updates())

	affectedMessages := &mtproto.TLMessagesAffectedMessages{Data2: &mtproto.Messages_AffectedMessages_Data{
		Pts:      pts,
		PtsCount: ptsCount,
	}}

	getPeerByDid := func(userId int32, did int64) (int32, int32) {
		if did < 0 {
			return base.PEER_CHAT, int32(-did)
		} else {
			id1 := int32(did & 0xffffffff)
			id2 := int32(did >> 32)
			if userId == id1 {
				return base.PEER_USER, id2
			} else {
				return base.PEER_USER, id1
			}
		}
	}

	glog.Info("messages.deleteMessages#e58e95d2 - deletedDialogsMap: ", deletedDialogsMap)
	for k, v := range deletedDialogsMap {
		for _, did := range v.didList {
			pType, pId := getPeerByDid(k, did)
			// currentTopMessage
			//if topMessage != 0 {
			topMessage := s.DialogModel.GetTopMessage(k, int8(pType), pId)
			glog.Info("topMessage - ", topMessage)
			for _, id := range v.idList {
				if id == topMessage {
					s.DialogModel.InsertOrUpdateDialog(k, pType, pId, s.MessageModel.GetLastPeerMessageId(k, did), false, false)
					break
				}
			}
			//}
		}
	}

	glog.Infof("messages.deleteMessages#e58e95d2 - reply: %s", logger.JsonDebugData(affectedMessages))
	return affectedMessages.To_Messages_AffectedMessages(), nil
}
