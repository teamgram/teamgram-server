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
	"time"

	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	idgenpb "github.com/teamgram/teamgram-server/v2/app/service/idgen/idgen"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const maxSendMultiMediaItems = 10

// MessagesSendMultiMedia
// messages.sendMultiMedia#1bf89d74 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true update_stickersets_order:flags.15?true invert_media:flags.16?true allow_paid_floodskip:flags.19?true peer:InputPeer reply_to:flags.0?InputReplyTo multi_media:Vector<InputSingleMedia> schedule_date:flags.10?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut effect:flags.18?long allow_paid_stars:flags.21?long = Updates;
func (c *MessagesCore) MessagesSendMultiMedia(in *tg.TLMessagesSendMultiMedia) (*tg.Updates, error) {
	md := c.MD
	if md == nil || md.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	selfUserID := md.UserId
	authKeyID := md.PermAuthKeyId

	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}

	peer, ok := resolveMessagePeer(in.Peer, selfUserID)
	if !ok {
		return nil, tg.Err400PeerIdInvalid
	}

	if err := checkSendMultiMediaUnsupportedFields(in); err != nil {
		return nil, err
	}
	replyHeader, err := makeMessageReplyHeader(in.ReplyTo)
	if err != nil {
		return nil, err
	}
	if err := checkSendMultiMediaItems(in.MultiMedia); err != nil {
		return nil, err
	}
	if peer.PeerType == payload.PeerTypeChat {
		if err := c.checkChatMessageAction(peer.PeerID, chatpb.MessageActionSendAlbum, "album"); err != nil {
			return nil, err
		}
	}

	groupedID, err := c.newMessageGroupedID("messages.sendMultiMedia")
	if err != nil {
		return nil, err
	}
	date := int32(time.Now().Unix())
	checkedMediaActions := make(map[string]struct{})
	outboxes := make([]msg.OutboxMessageClazz, 0, len(in.MultiMedia))
	for _, item := range in.MultiMedia {
		media, err := resolveMessageMedia(c.ctx, c.svcCtx.Repo.MediaClient, c.svcCtx.Repo.UserClient, authKeyID, item.Media)
		if err != nil {
			mappedErr := mapMediaResolveError(err)
			if mappedErr == tg.ErrInternalServerError {
				c.Logger.Errorf("messages.sendMultiMedia - media resolve failed: self_user_id: %d, peer_type: %d, peer_id: %d, random_id: %d, err: %v",
					selfUserID, peer.PeerType, peer.PeerID, item.RandomId, err)
			}
			return nil, mappedErr
		}
		if peer.PeerType == payload.PeerTypeChat {
			action, mediaKind := chatMessageActionForMedia(media)
			key := chatMessageActionKey(action, mediaKind)
			if _, ok := checkedMediaActions[key]; !ok {
				if err := c.checkChatMessageAction(peer.PeerID, action, mediaKind); err != nil {
					return nil, err
				}
				checkedMediaActions[key] = struct{}{}
			}
		}
		outboxes = append(outboxes, buildMessageMediaOutbox(mediaOutboxInput{
			RandomId:    item.RandomId,
			Background:  in.Background,
			Silent:      in.Silent,
			Noforwards:  in.Noforwards,
			InvertMedia: in.InvertMedia,
			FromId:      selfUserID,
			PeerType:    peer.PeerType,
			PeerId:      peer.PeerID,
			ReplyTo:     replyHeader,
			Date:        date,
			Message:     item.Message,
			Media:       media,
			Entities:    item.Entities,
			GroupedId:   &groupedID,
		}))
	}

	var sendClient sendMessageClient = c.svcCtx.Repo.MsgClient
	updates, err := sendClient.MsgSendMessage(c.ctx, &msg.TLMsgSendMessage{
		ClearDraft:           in.ClearDraft,
		UserId:               selfUserID,
		AuthKeyId:            authKeyID,
		SourcePermAuthKeyId:  &authKeyID,
		ClearDraftBeforeDate: &date,
		PeerType:             peer.PeerType,
		PeerId:               peer.PeerID,
		Message:              outboxes,
	})
	if err != nil {
		c.Logger.Errorf("messages.sendMultiMedia - msg error: self_user_id: %d, peer_type: %d, peer_id: %d, err: %v",
			selfUserID, peer.PeerType, peer.PeerID, err)
		return nil, mapMsgSendError(err)
	}
	if err := userprojection.FillUpdatesUsers(c.ctx, c.svcCtx.Repo.UserClient, selfUserID, updates, userprojection.MissingStoredReference); err != nil {
		return nil, err
	}

	return updates, nil
}

func checkSendMultiMediaUnsupportedFields(in *tg.TLMessagesSendMultiMedia) error {
	if in.ScheduleDate != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.UpdateStickersetsOrder {
		return tg.ErrInputRequestInvalid
	}
	if in.SendAs != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.QuickReplyShortcut != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.Effect != nil {
		return tg.ErrInputRequestInvalid
	}
	if in.AllowPaidFloodskip {
		return tg.ErrInputRequestInvalid
	}
	if in.AllowPaidStars != nil {
		return tg.ErrInputRequestInvalid
	}
	return nil
}

func checkSendMultiMediaItems(items []tg.InputSingleMediaClazz) error {
	if len(items) == 0 {
		return tg.ErrMediaEmpty
	}
	if len(items) > maxSendMultiMediaItems {
		return tg.ErrMultiMediaTooLong
	}
	seen := make(map[int64]struct{}, len(items))
	for _, item := range items {
		if item == nil {
			return tg.ErrInputRequestInvalid
		}
		if item.RandomId == 0 {
			return tg.ErrRandomIdEmpty
		}
		if err := checkCaption(item.Message); err != nil {
			return err
		}
		if _, ok := seen[item.RandomId]; ok {
			return tg.ErrRandomIdDuplicate
		}
		seen[item.RandomId] = struct{}{}
	}
	return nil
}

func (c *MessagesCore) newMessageGroupedID(logPrefix string) (int64, error) {
	var client idgenClient = c.svcCtx.Repo.IdgenClient
	if client == nil {
		c.Logger.Errorf("%s - idgen client is nil", logPrefix)
		return 0, tg.ErrInternalServerError
	}
	id, err := client.IdgenNextId(c.ctx, &idgenpb.TLIdgenNextId{})
	if err != nil {
		c.Logger.Errorf("%s - idgen next id failed: err: %v", logPrefix, err)
		return 0, tg.ErrInternalServerError
	}
	if id == nil || id.V <= 0 {
		c.Logger.Errorf("%s - idgen next id invalid: id: %#v", logPrefix, id)
		return 0, tg.ErrInternalServerError
	}
	return id.V, nil
}
