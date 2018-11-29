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
	"github.com/nebula-chat/chatengine/service/document/client"
	"golang.org/x/net/context"
	"time"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/update"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/core/message"
)

/*
	body: { messages_editChatPhoto
	  chat_id: 283362015 [INT],
	  photo: { inputChatUploadedPhoto
		file: { inputFile
		  id: 970161873153942262 [LONG],
		  parts: 9 [INT],
		  name: ".jpg" [STRING],
		  md5_checksum: "f1987132d8a3949420f0130d9e6afe08" [STRING],
		},
	  },
	},
*/
// messages.editChatPhoto#ca4c79d8 chat_id:int photo:InputChatPhoto = Updates;
func (s *MessagesServiceImpl) MessagesEditChatPhoto(ctx context.Context, request *mtproto.TLMessagesEditChatPhoto) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.editChatPhoto#ca4c79d8 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	chatLogic, err := s.ChatModel.NewChatLogicById(request.ChatId)
	if err != nil {
		glog.Error("messages.editChatTitle#dc452855 - error: ", err)
		return nil, err
	}

	peer := &base.PeerUtil{
		PeerType: base.PEER_CHAT,
		PeerId:   chatLogic.GetChatId(),
	}
	_ = peer
	var (
		photoId int64 = 0
		action  *mtproto.MessageAction
	)

	chatPhoto := request.GetPhoto()
	switch chatPhoto.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputChatPhotoEmpty:
		photoId = 0
		action = mtproto.NewTLMessageActionChatDeletePhoto().To_MessageAction()
		// chatLogic.EditChatPhoto(md.UserId, photoId)
		// chatLogic.MakeMessageService(md.UserId, action.To_MessageAction())
	case mtproto.TLConstructor_CRC32_inputChatUploadedPhoto:
		file := chatPhoto.GetData2().GetFile()
		// photoId = helper.NextSnowflakeId()
		result, err := document_client.UploadPhotoFile(md.AuthId, file) // photoId, file.GetData2().GetId(), file.GetData2().GetParts(), file.GetData2().GetName(), file.GetData2().GetMd5Checksum())
		if err != nil {
			glog.Errorf("UploadPhoto error: %v", err)
			return nil, err
		}
		photoId = result.PhotoId
		// user.SetUserPhotoID(md.UserId, uuid)
		// fileData := mediaData.GetFile().GetData2()
		photo := &mtproto.TLPhoto{Data2: &mtproto.Photo_Data{
			Id:          photoId,
			HasStickers: false,
			AccessHash:  result.AccessHash, // photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
			Date:        int32(time.Now().Unix()),
			Sizes:       result.SizeList,
		}}
		action2 := &mtproto.TLMessageActionChatEditPhoto{Data2: &mtproto.MessageAction_Data{
			Photo: photo.To_Photo(),
		}}
		action = action2.To_MessageAction()
	case mtproto.TLConstructor_CRC32_inputChatPhoto:
		// photo := chatPhoto.GetData2().GetId()
	}

	chatLogic.EditChatPhoto(md.UserId, photoId)
	editChatPhotoMessage := chatLogic.MakeMessageService(md.UserId, action)
	randomId := core.GetUUID()

	resultCB := func(pts, ptsCount int32, outBox *message.MessageBox2) (*mtproto.Updates, error) {
		syncUpdates := updates.NewUpdatesLogic(md.UserId)

		updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
			Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
		}}
		syncUpdates.AddUpdate(updateChatParticipants.To_Update())
		syncUpdates.AddUpdateNewMessage(pts, ptsCount, outBox.ToMessage(outBox.OwnerId))
		syncUpdates.AddUsers(s.UserModel.GetUserListByIdList(md.UserId, chatLogic.GetChatParticipantIdList()))
		syncUpdates.AddChat(chatLogic.ToChat(md.UserId))

		syncUpdates.AddUpdateMessageId(outBox.MessageId, outBox.RandomId)

		return syncUpdates.ToUpdates(), nil
	}

	syncNotMeCB := func(pts, ptsCount int32, outBox *message.MessageBox2) (int64, *mtproto.Updates, error) {
		syncUpdates := updates.NewUpdatesLogic(md.UserId)

		updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
			Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
		}}
		syncUpdates.AddUpdate(updateChatParticipants.To_Update())
		syncUpdates.AddUpdateNewMessage(pts, ptsCount, outBox.ToMessage(outBox.OwnerId))
		syncUpdates.AddUsers(s.UserModel.GetUserListByIdList(md.UserId, chatLogic.GetChatParticipantIdList()))
		syncUpdates.AddChat(chatLogic.ToChat(md.UserId))

		return md.AuthId, syncUpdates.ToUpdates(), nil
	}

	pushCB := func(pts, ptsCount int32, inBox *message.MessageBox2) (*mtproto.Updates, error) {
		pushUpdates := updates.NewUpdatesLogic(md.UserId)

		updateChatParticipants := &mtproto.TLUpdateChatParticipants{Data2: &mtproto.Update_Data{
			Participants: chatLogic.GetChatParticipants().To_ChatParticipants(),
		}}
		pushUpdates.AddUpdate(updateChatParticipants.To_Update())
		pushUpdates.AddUpdateNewMessage(pts, ptsCount, inBox.ToMessage(inBox.OwnerId))
		pushUpdates.AddUsers(s.UserModel.GetUserListByIdList(inBox.OwnerId, chatLogic.GetChatParticipantIdList()))
		pushUpdates.AddChat(chatLogic.ToChat(inBox.OwnerId))

		return pushUpdates.ToUpdates(), nil
	}

	replyUpdates, _ := s.MessageModel.SendMessage(
		md.UserId,
		peer,
		randomId,
		editChatPhotoMessage,
		resultCB,
		syncNotMeCB,
		pushCB)

	glog.Infof("messages.editChatPhoto#ca4c79d8 - reply: {%v}", replyUpdates)
	return replyUpdates, nil
}
