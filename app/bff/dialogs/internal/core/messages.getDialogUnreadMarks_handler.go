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

// MessagesGetDialogUnreadMarks
// messages.getDialogUnreadMarks#22e24e22 = Vector<DialogPeer>;
func (c *DialogsCore) MessagesGetDialogUnreadMarks(in *tg.TLMessagesGetDialogUnreadMarks) (*tg.VectorDialogPeer, error) {
	return &tg.VectorDialogPeer{
		Datas: []tg.DialogPeerClazz{
			tg.MakeTLDialogPeer(&tg.TLDialogPeer{
				Peer: tg.MakeTLPeerUser(&tg.TLPeerUser{UserId: 1}),
			}),
		},
	}, nil
}
