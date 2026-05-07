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
	userprojection "github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// MessagesGetPeerSettings
// messages.getPeerSettings#efd9a6a2 peer:InputPeer = messages.PeerSettings;
func (c *DialogsCore) MessagesGetPeerSettings(in *tg.TLMessagesGetPeerSettings) (*tg.MessagesPeerSettings, error) {
	if c.MD == nil || c.MD.UserId <= 0 {
		return nil, tg.ErrUserIdInvalid
	}
	if in == nil {
		return nil, tg.ErrInputRequestInvalid
	}
	peerUserID, ok := resolveDialogUserPeerID(in.Peer, c.MD.UserId)
	if !ok {
		return nil, tg.Err400PeerIdInvalid
	}
	users, err := userprojection.ProjectUsers(c.ctx, c.svcCtx.Repo.UserClient, c.MD.UserId, []int64{peerUserID}, userprojection.MissingExplicitInput)
	if err != nil {
		return nil, err
	}

	return tg.MakeTLMessagesPeerSettings(&tg.TLMessagesPeerSettings{
		Settings: tg.MakeTLPeerSettings(&tg.TLPeerSettings{}).ToPeerSettings(),
		Chats:    []tg.ChatClazz{},
		Users:    users,
	}).ToMessagesPeerSettings(), nil
}
