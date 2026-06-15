// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	msgpb "github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesEditChatPhoto
// messages.editChatPhoto#35ddd674 chat_id:long photo:InputChatPhoto = Updates;
func (c *ChatsCore) MessagesEditChatPhoto(in *tg.TLMessagesEditChatPhoto) (*tg.Updates, error) {
	selfID := selfID(c.MD)
	photo, action, err := c.resolveEditChatPhoto(in.ChatId, in.Photo)
	if err != nil {
		return nil, err
	}

	mutableChat, err := c.svcCtx.Repo.ChatClient.ChatEditChatPhoto(c.ctx, &chatpb.TLChatEditChatPhoto{
		ChatId:     in.ChatId,
		EditUserId: selfID,
		ChatPhoto:  photo,
	})
	if err != nil {
		return nil, mapChatError(err)
	}

	updates, err := c.svcCtx.Repo.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    selfID,
		AuthKeyId: c.MD.PermAuthKeyId,
		PeerType:  payload.PeerTypeChat,
		PeerId:    mutableChat.Chat.Id,
		Message: []msgpb.OutboxMessageClazz{
			msgpb.MakeTLOutboxMessage(&msgpb.TLOutboxMessage{
				NoWebpage: true,
				RandomId:  normalizeCreateChatServiceMessageRandomID(rand.Int63()),
				Message: tg.MakeTLMessageService(&tg.TLMessageService{
					Out:    true,
					FromId: tg.MakePeerUser(selfID),
					PeerId: tg.MakePeerChat(mutableChat.Chat.Id),
					Date:   int32(time.Now().Unix()),
					Action: action,
				}),
			}),
		},
	})
	if err != nil {
		c.Logger.Errorf("messages.editChatPhoto - send edit photo service message failed: self_user_id=%d chat_id=%d err=%v", selfID, mutableChat.Chat.Id, err)
		return nil, tg.ErrInternalServerError
	}
	if updates == nil {
		c.Logger.Errorf("messages.editChatPhoto - send edit photo service message returned nil updates: self_user_id=%d chat_id=%d", selfID, mutableChat.Chat.Id)
		return nil, tg.ErrInternalServerError
	}

	return updates, nil
}

func (c *ChatsCore) resolveEditChatPhoto(chatID int64, input tg.InputChatPhotoClazz) (tg.PhotoClazz, tg.MessageActionClazz, error) {
	if input == nil {
		return nil, nil, tg.ErrInputRequestInvalid
	}

	switch p := input.(type) {
	case *tg.TLInputChatPhotoEmpty:
		return editChatPhotoDeleteAction()
	case *tg.TLInputChatUploadedPhoto:
		if c.svcCtx.Repo.MediaClient == nil {
			return nil, nil, tg.ErrInternalServerError
		}
		photo, err := c.svcCtx.Repo.MediaClient.MediaUploadProfilePhotoFile(c.ctx, &mediapb.TLMediaUploadProfilePhotoFile{
			OwnerId:          c.MD.PermAuthKeyId,
			File:             p.File,
			Video:            p.Video,
			VideoStartTs:     p.VideoStartTs,
			VideoEmojiMarkup: p.VideoEmojiMarkup,
		})
		if err != nil {
			return nil, nil, c.mapEditChatPhotoMediaError(chatID, err)
		}
		if photo == nil || photo.Clazz == nil {
			return nil, nil, tg.ErrPhotoInvalid
		}
		return photo.Clazz, tg.MakeTLMessageActionChatEditPhoto(&tg.TLMessageActionChatEditPhoto{
			Photo: photo.Clazz,
		}), nil
	case *tg.TLInputChatPhoto:
		switch inputPhoto := p.Id.(type) {
		case *tg.TLInputPhotoEmpty:
			return editChatPhotoDeleteAction()
		case *tg.TLInputPhoto:
			if c.svcCtx.Repo.MediaClient == nil {
				return nil, nil, tg.ErrInternalServerError
			}
			photo, err := c.svcCtx.Repo.MediaClient.MediaGetPhoto(c.ctx, &mediapb.TLMediaGetPhoto{
				PhotoId: inputPhoto.Id,
			})
			if err != nil {
				return nil, nil, c.mapEditChatPhotoMediaError(chatID, err)
			}
			if photo == nil || photo.Clazz == nil {
				return nil, nil, tg.ErrPhotoInvalid
			}
			return photo.Clazz, tg.MakeTLMessageActionChatEditPhoto(&tg.TLMessageActionChatEditPhoto{
				Photo: photo.Clazz,
			}), nil
		default:
			return nil, nil, tg.ErrInputRequestInvalid
		}
	default:
		return nil, nil, tg.ErrInputRequestInvalid
	}
}

func editChatPhotoDeleteAction() (tg.PhotoClazz, tg.MessageActionClazz, error) {
	return tg.MakeTLPhotoEmpty(&tg.TLPhotoEmpty{}), tg.MakeTLMessageActionChatDeletePhoto(&tg.TLMessageActionChatDeletePhoto{}), nil
}

func (c *ChatsCore) mapEditChatPhotoMediaError(chatID int64, err error) error {
	mappedErr := mapEditChatPhotoMediaError(err)
	if mappedErr == tg.ErrInternalServerError {
		c.Logger.Errorf("messages.editChatPhoto - media resolve failed: self_user_id=%d chat_id=%d err=%v", selfID(c.MD), chatID, err)
	}
	return mappedErr
}

func mapEditChatPhotoMediaError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, mediapb.ErrPhotoNotFound), isRemoteEditChatPhotoMediaError(err, mediapb.ErrPhotoNotFound):
		return tg.ErrPhotoIdInvalid
	case errors.Is(err, mediapb.ErrMediaStorage),
		isRemoteEditChatPhotoMediaError(err, mediapb.ErrMediaStorage),
		errors.Is(err, mediapb.ErrMediaDownstream),
		isRemoteEditChatPhotoMediaError(err, mediapb.ErrMediaDownstream):
		return tg.ErrInternalServerError
	case errors.Is(err, mediapb.ErrMediaInvalidArgument),
		isRemoteEditChatPhotoMediaError(err, mediapb.ErrMediaInvalidArgument),
		errors.Is(err, mediapb.ErrMediaInvalidUploadedFile),
		isRemoteEditChatPhotoMediaError(err, mediapb.ErrMediaInvalidUploadedFile),
		errors.Is(err, mediapb.ErrMediaChecksumInvalid),
		isRemoteEditChatPhotoMediaError(err, mediapb.ErrMediaChecksumInvalid),
		errors.Is(err, mediapb.ErrMediaBlocked),
		isRemoteEditChatPhotoMediaError(err, mediapb.ErrMediaBlocked):
		return tg.ErrPhotoInvalid
	default:
		return tg.ErrInternalServerError
	}
}

func isRemoteEditChatPhotoMediaError(err error, target error) bool {
	if err == nil || target == nil {
		return false
	}
	return strings.Contains(err.Error(), "biz error: "+target.Error())
}
