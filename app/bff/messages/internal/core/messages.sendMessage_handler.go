// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesSendMessage
// messages.sendMessage#983f9745 flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true update_stickersets_order:flags.15?true invert_media:flags.16?true allow_paid_floodskip:flags.19?true peer:InputPeer reply_to:flags.0?InputReplyTo message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut effect:flags.18?long = Updates;
func (c *MessagesCore) MessagesSendMessage(in *tg.TLMessagesSendMessage) (*tg.Updates, error) {
	peer := tg.FromInputPeer2(0, in.Peer)
	if c.MD != nil {
		peer = tg.FromInputPeer2(c.MD.UserId, in.Peer)
	}

	switch peer.PeerType {
	case tg.PEER_SELF, tg.PEER_USER, tg.PEER_CHAT:
	case tg.PEER_CHANNEL:
		return nil, tg.ErrEnterpriseIsBlocked
	default:
		return nil, tg.ErrPeerIdInvalid
	}

	if in.Message == "" {
		return nil, tg.ErrMessageEmpty
	}

	return tg.MakeTLUpdateShortSentMessage(&tg.TLUpdateShortSentMessage{
		Out:      true,
		Id:       makePlaceholderMessageID(in.RandomId),
		Pts:      1,
		PtsCount: 1,
		Date:     int32(time.Now().Unix()),
		Entities: in.Entities,
	}).ToUpdates(), nil
}

func makePlaceholderMessageID(randomID int64) int32 {
	if randomID < 0 {
		randomID = -randomID
	}
	id := int32(randomID % 0x7fffffff)
	if id == 0 {
		id = 1
	}
	return id
}
