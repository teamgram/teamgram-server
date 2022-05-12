// Copyright 2022 Teamgram Authors
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
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"math/rand"

	"github.com/teamgram/proto/mtproto"
	msgpb "github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	mediapb "github.com/teamgram/teamgram-server/app/service/media/media"
)

// MessagesEditChatPhoto
// messages.editChatPhoto#35ddd674 chat_id:long photo:InputChatPhoto = Updates;
func (c *ChatsCore) MessagesEditChatPhoto(in *mtproto.TLMessagesEditChatPhoto) (*mtproto.Updates, error) {
	var (
		// photoId int64 = 0
		action *mtproto.MessageAction
		err    error
	)

	chatPhoto := in.GetPhoto()
	photo := mtproto.MakeTLPhotoEmpty(nil).To_Photo()
	switch chatPhoto.GetPredicateName() {
	case mtproto.Predicate_inputChatPhotoEmpty:
		// inputChatPhotoEmpty#1ca48f57 = InputChatPhoto;

		action = mtproto.MakeTLMessageActionChatDeletePhoto(nil).To_MessageAction()
	case mtproto.Predicate_inputChatUploadedPhoto:
		// inputChatUploadedPhoto#c642724e flags:# file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double = InputChatPhoto;

		// file := request.GetFile()
		photo, err = c.svcCtx.Dao.MediaClient.MediaUploadProfilePhotoFile(c.ctx, &mediapb.TLMediaUploadProfilePhotoFile{
			OwnerId:      c.MD.AuthId,
			File:         chatPhoto.GetFile(),
			Video:        chatPhoto.GetVideo(),
			VideoStartTs: chatPhoto.GetVideoStartTs(),
		})
		if err != nil {
			c.Logger.Errorf("messages.editChatPhoto - error: %v", err)
			return nil, err
		}
		action = mtproto.MakeMessageActionChatEditPhoto(photo)
	case mtproto.Predicate_inputChatPhoto:
		// inputChatPhoto#8953ad37 id:InputPhoto = InputChatPhoto;

		// TODO: android
		/*
			inputMessagesFilterChatPhotos - Return only chat photo changes

			body: { messages_search
			  flags: 0 [INT],
			  peer: { inputPeerChat
			    chat_id: 518812270 [INT],
			  },
			  q: "" [STRING],
			  from_id: [ SKIPPED BY BIT 0 IN FIELD flags ],
			  top_msg_id: [ SKIPPED BY BIT 1 IN FIELD flags ],
			  filter: { inputMessagesFilterChatPhotos },
			  min_date: 0 [INT],
			  max_date: 0 [INT],
			  offset_id: 57963 [INT],
			  add_offset: 0 [INT],
			  limit: 100 [INT],
			  max_id: 0 [INT],
			  min_id: 0 [INT],
			  hash: 0 [INT],
			},
		*/

		//switch chatPhoto.GetPredicateName() {
		//case mtproto.Predicate_inputPhotoEmpty:
		//	// inputPhotoEmpty#1cd7bf0d = InputPhoto;
		//	photos, err := s.UserFacade.GetCacheUserPhotos(ctx, md.UserId)
		//	if err != nil {
		//		log.Errorf("photos.updateProfilePhoto - error: %v", err)
		//		return nil, mtproto.ErrInternelServerError
		//	}
		//	photos.RemovePhotoId(photos.GetDefaultPhotoId(), func(id int64) *mtproto.Photo {
		//		photo := media_client.GetPhoto(id)
		//		return photo
		//	})
		//
		//	err = s.UserFacade.PutCacheUserPhotos(ctx, md.UserId, photos)
		//	if err != nil {
		//		log.Errorf("photos.updateProfilePhoto - error: %v", err)
		//	}
		//
		//	photos2 = mtproto.MakeTLPhotosPhoto(&mtproto.Photos_Photo{
		//		Photo: photos.Photo,
		//		Users: []*mtproto.User{},
		//	}).To_Photos_Photo()
		//case mtproto.Predicate_inputPhoto:
		//	// inputPhoto#3bb3b94a id:long access_hash:long file_reference:bytes = InputPhoto;
		//
		//	id := request.GetId().To_InputPhoto()
		//	// TODO(@benqi): check inputPhoto.access_hash
		//
		//	photo2 := media_client.GetPhoto(id.GetId())
		//	photos2 = mtproto.MakeTLPhotosPhoto(&mtproto.Photos_Photo{
		//		Photo: photo2,
		//		Users: []*mtproto.User{},
		//	}).To_Photos_Photo()
		//default:
		//}
	default:
		err = mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("messages.editChatPhoto - error: %v", err)
		return nil, err
	}

	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatEditChatPhoto(c.ctx, &chatpb.TLChatEditChatPhoto{
		ChatId:     in.ChatId,
		EditUserId: c.MD.UserId,
		ChatPhoto:  photo,
	})
	if err != nil {
		c.Logger.Errorf("messages.editChatPhoto - error: %v", err)
		return nil, err
	}

	replyUpdates, err := c.svcCtx.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  mtproto.PEER_CHAT,
		PeerId:    in.ChatId,
		Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      chat.MakeMessageService(c.MD.UserId, action),
			ScheduleDate: nil,
		}).To_OutboxMessage(),
	})
	if err != nil {
		c.Logger.Errorf("messages.editChatPhoto - error: %v", err)
		return nil, err
	}

	return replyUpdates, nil
}
