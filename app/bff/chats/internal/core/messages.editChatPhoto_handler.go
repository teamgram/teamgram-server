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

		photo, err = c.svcCtx.Dao.MediaClient.MediaUploadProfilePhotoFile(c.ctx, &mediapb.TLMediaUploadProfilePhotoFile{
			OwnerId:      c.MD.PermAuthKeyId,
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

		id := in.GetPhoto().GetId()
		if id.GetPredicateName() == mtproto.Predicate_inputPhotoEmpty {
			action = mtproto.MakeTLMessageActionChatDeletePhoto(nil).To_MessageAction()
		} else {
			photo, err = c.svcCtx.Dao.MediaClient.MediaGetPhoto(c.ctx, &mediapb.TLMediaGetPhoto{
				PhotoId: id.GetId(),
			})
			if err != nil {
				c.Logger.Errorf("messages.editChatPhoto - error: %v", err)
				return nil, err
			}

			action = mtproto.MakeMessageActionChatEditPhoto(photo)
		}
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

	replyUpdates, err := c.svcCtx.MsgClient.MsgSendMessageV2(
		c.ctx,
		&msgpb.TLMsgSendMessageV2{
			UserId:    c.MD.UserId,
			AuthKeyId: c.MD.PermAuthKeyId,
			PeerType:  mtproto.PEER_CHAT,
			PeerId:    in.ChatId,
			Message: []*msgpb.OutboxMessage{
				msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
					NoWebpage:    true,
					Background:   false,
					RandomId:     rand.Int63(),
					Message:      chat.MakeMessageService(c.MD.UserId, action),
					ScheduleDate: nil,
				}).To_OutboxMessage(),
			},
		})
	if err != nil {
		c.Logger.Errorf("messages.editChatPhoto - error: %v", err)
		return nil, err
	}

	return replyUpdates, nil
}
