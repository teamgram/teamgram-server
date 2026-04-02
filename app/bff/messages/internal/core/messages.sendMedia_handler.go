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

// MessagesSendMedia
// messages.sendMedia#7852834e flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true noforwards:flags.14?true update_stickersets_order:flags.15?true invert_media:flags.16?true allow_paid_floodskip:flags.19?true peer:InputPeer reply_to:flags.0?InputReplyTo media:InputMedia message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> schedule_date:flags.10?int send_as:flags.13?InputPeer quick_reply_shortcut:flags.17?InputQuickReplyShortcut effect:flags.18?long = Updates;
func (c *MessagesCore) MessagesSendMedia(in *tg.TLMessagesSendMedia) (*tg.Updates, error) {
	if _, err := bffPeerFromInput(c, in.Peer); err != nil {
		return nil, err
	}
	if in.Media == nil {
		return nil, tg.ErrInputRequestInvalid
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
