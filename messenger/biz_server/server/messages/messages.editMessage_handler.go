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
	"time"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/pkg/grpc_util"
	"github.com/nebula-chat/chatengine/pkg/logger"
	"github.com/nebula-chat/chatengine/mtproto"
	"golang.org/x/net/context"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	message2 "github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/message"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/sync/sync_client"
)

func (s *MessagesServiceImpl) makeUpdateEditMessageUpdates(selfUserId int32, message *mtproto.Message) *mtproto.TLUpdates {
	userIdList, _, _ := message2.PickAllIDListByMessages([]*mtproto.Message{message})
	userList := s.UserModel.GetUserListByIdList(selfUserId, userIdList)

	updateNew := &mtproto.TLUpdateEditMessage{Data2: &mtproto.Update_Data{
		Message_1: message,
		Pts:       int32(core.NextPtsId(selfUserId)),
		PtsCount:  1,
	}}
	return &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{updateNew.To_Update()},
		Users:   userList,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}}
}

func setEditMessageData(request *mtproto.TLMessagesEditMessage, messageBox *message2.MessageBox2) {
	data2 := messageBox.Message.GetData2()
	data2.Message = request.GetMessage()
	data2.ReplyMarkup = request.GetReplyMarkup()
	data2.Entities = request.GetEntities()
	data2.EditDate = int32(time.Now().Unix())

	messageBox.EditDate = data2.EditDate
	messageBox.EditMessage = request.GetMessage()
}

// messages.editMessage#5d1b8dd flags:# no_webpage:flags.1?true stop_geo_live:flags.12?true peer:InputPeer id:int message:flags.11?string reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> geo_point:flags.13?InputGeoPoint = Updates;
func (s *MessagesServiceImpl) MessagesEditMessage(ctx context.Context, request *mtproto.TLMessagesEditMessage) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.editMessage#5d1b8dd - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		peer               *base.PeerUtil
		err                error
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

	editOutbox, err := s.MessageModel.GetMessageBox2(int32(peer.PeerType), md.UserId, request.GetId())
	// TODO(@benqi): check invalid

	// Edit outbox
	setEditMessageData(request, editOutbox)
	syncUpdates := s.makeUpdateEditMessageUpdates(md.UserId, editOutbox.ToMessage(md.UserId))
	sync_client.GetSyncClient().SyncUpdatesNotMe(md.UserId, md.AuthId, syncUpdates.To_Updates())

	editOutbox.SaveMessageData()

	// TODO(@benqi):
	// push edit peer message
	peerEditMessages := s.MessageModel.GetPeerMessageListByMessageDataId(md.UserId, editOutbox.MessageDataId)
	for i := 0; i < len(peerEditMessages); i++ {
		editMessage := peerEditMessages[i]
		editMessage.MessageData = editOutbox.MessageData
		setEditMessageData(request, editMessage)
		editUpdates := s.makeUpdateEditMessageUpdates(editMessage.OwnerId, editMessage.ToMessage(editMessage.OwnerId))
		sync_client.GetSyncClient().PushUpdates(editMessage.OwnerId, editUpdates.To_Updates())
	}

	glog.Infof("messages.editMessage#5d1b8dd - reply: %s", logger.JsonDebugData(syncUpdates))
	return syncUpdates.To_Updates(), nil
}
