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

import "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"

// MessagesGetPinnedDialogs
// messages.getPinnedDialogs#d6b94df2 folder_id:int = messages.PeerDialogs;
func (c *DialogsCore) MessagesGetPinnedDialogs(in *tg.TLMessagesGetPinnedDialogs) (*tg.MessagesPeerDialogs, error) {
	return &tg.MessagesPeerDialogs{
		Dialogs: []tg.DialogClazz{
			tg.MakeTLDialog(&tg.TLDialog{
				Peer:            tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1}),
				TopMessage:      10,
				ReadInboxMaxId:  10,
				ReadOutboxMaxId: 10,
				UnreadCount:     0,
				NotifySettings:  tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}),
			}),
		},
		Messages: []tg.MessageClazz{
			tg.MakeTLMessageEmpty(&tg.TLMessageEmpty{Id: 10}),
		},
		Chats: []tg.ChatClazz{},
		Users: []tg.UserClazz{
			tg.MakeTLUserEmpty(&tg.TLUserEmpty{Id: 1}),
		},
		State: tg.MakeTLUpdatesState(&tg.TLUpdatesState{Pts: 1, Qts: 1, Date: 10, Seq: 1, UnreadCount: 0}),
	}, nil
}
