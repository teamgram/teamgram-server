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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesEditMessage
// messages.editMessage#51e842e1 flags:# no_webpage:flags.1?true invert_media:flags.16?true peer:InputPeer id:int message:flags.11?string media:flags.14?InputMedia reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.15?int schedule_repeat_period:flags.18?int quick_reply_shortcut_id:flags.17?int = Updates;
func (c *MessagesCore) MessagesEditMessage(in *tg.TLMessagesEditMessage) (*tg.Updates, error) {
	md := c.MD
	if md == nil || md.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if in.Id <= 0 {
		return nil, tg.ErrMsgIdInvalid
	}
	if in.Media != nil {
		return nil, tg.ErrMediaInvalid
	}
	if in.InvertMedia ||
		in.ScheduleDate != nil ||
		in.ScheduleRepeatPeriod != nil ||
		in.QuickReplyShortcutId != nil {
		return nil, tg.ErrInputRequestInvalid
	}
	if in.Message == nil {
		return nil, tg.ErrMessageNotModified
	}
	if err := checkMessage(*in.Message); err != nil {
		return nil, err
	}

	peerUserID, ok := resolveUserPeerID(in.Peer, md.UserId)
	if !ok {
		return nil, tg.Err400PeerIdInvalid
	}

	now := int32(time.Now().Unix())
	newMessage := tg.MakeTLMessage(&tg.TLMessage{
		Out:         true,
		FromId:      tg.MakePeerUser(md.UserId),
		PeerId:      tg.MakePeerUser(peerUserID),
		Id:          in.Id,
		Date:        now,
		Message:     *in.Message,
		ReplyMarkup: in.ReplyMarkup,
		Entities:    in.Entities,
		EditDate:    &now,
		EditHide:    false,
	})
	dstMessage := tg.MakeTLMessageBox(&tg.TLMessageBox{
		UserId:       md.UserId,
		MessageId:    in.Id,
		SenderUserId: md.UserId,
		PeerType:     payload.PeerTypeUser,
		PeerId:       peerUserID,
		Message:      newMessage,
	})

	var editClient editMessageClient = c.svcCtx.Repo.MsgClient
	updates, err := editClient.MsgEditMessageV2(c.ctx, &msg.TLMsgEditMessageV2{
		UserId:    md.UserId,
		AuthKeyId: md.PermAuthKeyId,
		PeerType:  payload.PeerTypeUser,
		PeerId:    peerUserID,
		EditType:  0,
		NewMessage: msg.MakeTLOutboxMessage(&msg.TLOutboxMessage{
			NoWebpage:    in.NoWebpage,
			RandomId:     0,
			Message:      newMessage,
			ScheduleDate: in.ScheduleDate,
		}),
		DstMessage: dstMessage,
	})
	if err != nil {
		return nil, mapMsgEditError(err)
	}
	if err := userprojection.FillUpdatesUsers(c.ctx, c.svcCtx.Repo.UserClient, md.UserId, updates, userprojection.MissingStoredReference); err != nil {
		return nil, err
	}

	return updates, nil
}
