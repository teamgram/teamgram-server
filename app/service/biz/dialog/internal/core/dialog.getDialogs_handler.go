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
	"github.com/teamgram/teamgram-server/v2/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

var _ *tg.Bool

// DialogGetDialogs
// dialog.getDialogs user_id:long exclude_pinned:Bool folder_id:int = Vector<DialogExt>;
func (c *DialogCore) DialogGetDialogs(in *dialog.TLDialogGetDialogs) (*dialog.VectorDialogExt, error) {
	if in != nil && in.UserId != 0 {
		return &dialog.VectorDialogExt{
			Datas: []dialog.DialogExtClazz{
				makeDialogExtPlaceholder(in.UserId, tg.PEER_USER, in.UserId, 10),
			},
		}, nil
	}

	return &dialog.VectorDialogExt{
		Datas: []dialog.DialogExtClazz{},
	}, nil
}

func makeDialogExtPlaceholder(userID, peerType, peerID int64, topMessage int32) dialog.DialogExtClazz {
	return dialog.MakeTLDialogExt(&dialog.TLDialogExt{
		Order:          10,
		Dialog:         makeDialogPlaceholder(peerType, peerID, topMessage),
		AvailableMinId: 1,
		Date:           10,
		ThemeEmoticon:  "",
		TtlPeriod:      0,
		WallpaperId:    0,
	})
}

func makeDialogPlaceholder(peerType, peerID int64, topMessage int32) tg.DialogClazz {
	return tg.MakeTLDialog(&tg.TLDialog{
		Peer:            makeDialogPeer(peerType, peerID),
		TopMessage:      topMessage,
		ReadInboxMaxId:  topMessage,
		ReadOutboxMaxId: topMessage,
		UnreadCount:     0,
		NotifySettings:  tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
	})
}

func makeDialogPeer(peerType, peerID int64) tg.PeerClazz {
	switch peerType {
	case tg.PEER_SELF, tg.PEER_USER:
		return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peerID})
	case tg.PEER_CHAT:
		return tg.MakeTLPeerChat(&tg.TLPeerChat{ChatId: peerID})
	case tg.PEER_CHANNEL:
		return tg.MakeTLPeerChannel(&tg.TLPeerChannel{ChannelId: peerID})
	default:
		return tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: peerID})
	}
}
