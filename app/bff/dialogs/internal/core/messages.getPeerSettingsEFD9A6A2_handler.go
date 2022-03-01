// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetPeerSettingsEFD9A6A2
// messages.getPeerSettings#efd9a6a2 peer:InputPeer = messages.PeerSettings;
func (c *DialogsCore) MessagesGetPeerSettingsEFD9A6A2(in *mtproto.TLMessagesGetPeerSettingsEFD9A6A2) (*mtproto.Messages_PeerSettings, error) {
	settings, err := c.MessagesGetPeerSettings3672E09C(&mtproto.TLMessagesGetPeerSettings3672E09C{
		Peer: in.Peer,
	})
	if err != nil {
		c.Logger.Errorf("messages.getPeerSettingsEFD9A6A2 - error: %v", err)
		return nil, err
	}

	return mtproto.MakeTLMessagesPeerSettings(&mtproto.Messages_PeerSettings{
		Settings: settings,
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_PeerSettings(), nil
}
