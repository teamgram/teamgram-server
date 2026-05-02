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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetPeerDialogs
// messages.getPeerDialogs#e470bcfd peers:Vector<InputDialogPeer> = messages.PeerDialogs;
func (c *DialogsCore) MessagesGetPeerDialogs(in *tg.TLMessagesGetPeerDialogs) (*tg.MessagesPeerDialogs, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}

	dialogs := make([]tg.DialogClazz, 0, len(in.Peers))
	messages := make([]tg.MessageClazz, 0, len(in.Peers))
	users := make([]tg.UserClazz, 0, len(in.Peers))
	for _, peer := range in.Peers {
		inputDialogPeer, ok := (&tg.InputDialogPeer{Clazz: peer}).ToInputDialogPeer()
		if !ok {
			if _, isFolder := (&tg.InputDialogPeer{Clazz: peer}).ToInputDialogPeerFolder(); isFolder {
				return nil, tg.ErrFolderIdInvalid
			}
			return nil, tg.ErrInputConstructorInvalid
		}

		peerUserID, ok := resolveDialogUserPeerID(inputDialogPeer.Peer, c.MD.UserId)
		if !ok {
			return nil, tg.Err400PeerIdInvalid
		}
		fallback, err := c.fetchCanonicalUserDialog("messages.getPeerDialogs", peerUserID)
		if err != nil {
			return nil, err
		}
		if fallback == nil {
			continue
		}
		dialogs = append(dialogs, fallback.Dialog)
		messages = append(messages, fallback.Messages...)
		users = append(users, fallback.Users...)
	}

	return tg.MakeTLMessagesPeerDialogs(&tg.TLMessagesPeerDialogs{
		Dialogs:  dialogs,
		Messages: messages,
		Chats:    []tg.ChatClazz{},
		Users:    users,
		State: tg.MakeTLUpdatesState(&tg.TLUpdatesState{
			Pts:         0,
			Qts:         0,
			Date:        0,
			Seq:         0,
			UnreadCount: 0,
		}),
	}).ToMessagesPeerDialogs(), nil
}

func resolveDialogUserPeerID(peer tg.InputPeerClazz, selfUserID int64) (int64, bool) {
	p := tg.FromInputPeer2(selfUserID, peer)
	switch p.PeerType {
	case tg.PEER_SELF:
		return selfUserID, selfUserID > 0
	case tg.PEER_USER:
		return p.PeerId, p.PeerId > 0
	default:
		return 0, false
	}
}
