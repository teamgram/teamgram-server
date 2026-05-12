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

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/msg"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	chatpb "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesSendMessage
// messages.sendMessage#545cd15a flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true update_stickersets_order:flags.15?true invert_media:flags.16?true allow_paid_floodskip:flags.19?true peer:InputPeer reply_to:flags.0?InputReplyTo message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int schedule_repeat_period:flags.24?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut effect:flags.18?long allow_paid_stars:flags.21?long suggested_post:flags.22?SuggestedPost = Updates;
func (c *MessagesCore) MessagesSendMessage(in *tg.TLMessagesSendMessage) (*tg.Updates, error) {
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

	if err := checkMessage(in.Message); err != nil {
		return nil, err
	}

	if in.RandomId == 0 {
		return nil, tg.ErrRandomIdEmpty
	}

	if err := checkUnsupportedFields(in); err != nil {
		return nil, err
	}
	if peer.PeerType == payload.PeerTypeChat {
		if err := c.checkChatMessageAction(peer.PeerID, chatpb.MessageActionSendText, ""); err != nil {
			return nil, err
		}
	}
	replyHeader, err := makeMessageReplyHeader(in.ReplyTo)
	if err != nil {
		return nil, err
	}
	peerID := tg.MakePeerUser(peer.PeerID)
	if peer.PeerType == payload.PeerTypeChat {
		peerID = tg.MakePeerChat(peer.PeerID)
	}

	outgoingMsg := tg.MakeTLMessage(&tg.TLMessage{
		Out:      true,
		FromId:   tg.MakePeerUser(selfUserID),
		PeerId:   peerID,
		ReplyTo:  replyHeader,
		Date:     int32(time.Now().Unix()),
		Message:  in.Message,
		Entities: remakeMessageTextEntities(in.Message, in.Entities, selfUserID, false),
	})

	outbox := msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
		RandomId: in.RandomId,
		Message:  outgoingMsg,
	})
	clearDraftBeforeDate := outgoingMsg.Date

	var sendClient sendMessageClient = c.svcCtx.Repo.MsgClient
	updates, err := sendClient.MsgSendMessageV2(c.ctx, &msg.TLMsgSendMessageV2{
		ClearDraft:           in.ClearDraft,
		UserId:               selfUserID,
		AuthKeyId:            authKeyID,
		SourcePermAuthKeyId:  &authKeyID,
		ClearDraftBeforeDate: &clearDraftBeforeDate,
		PeerType:             peer.PeerType,
		PeerId:               peer.PeerID,
		Message:              []msg.OutboxMessageClazz{outbox},
	})
	if err != nil {
		c.Logger.Errorf("messages.sendMessage - msg error: self_user_id: %d, peer_type: %d, peer_id: %d, random_id: %d, err: %v",
			selfUserID, peer.PeerType, peer.PeerID, in.RandomId, err)
		return nil, mapMsgSendError(err)
	}

	return updates, nil
}
