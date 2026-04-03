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

// MessagesSendScreenshotNotification
// messages.sendScreenshotNotification#a1405817 peer:InputPeer reply_to:InputReplyTo random_id:long = Updates;
func (c *DialogsCore) MessagesSendScreenshotNotification(in *tg.TLMessagesSendScreenshotNotification) (*tg.Updates, error) {
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

	return tg.MakeTLUpdateShort(&tg.TLUpdateShort{
		Update: tg.MakeTLUpdateNewMessage(&tg.TLUpdateNewMessage{
			Message:  makePlaceholderDialogMessage(peer.PeerId, makePlaceholderDialogMessageID(in.RandomId)),
			Pts:      1,
			PtsCount: 1,
		}),
		Date: int32(time.Now().Unix()),
	}).ToUpdates(), nil
}
