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

// ChannelsGetSendAs
// channels.getSendAs#dc770ee peer:InputPeer = channels.SendAsPeers;
func (c *MessagesCore) ChannelsGetSendAs(in *mtproto.TLChannelsGetSendAs) (*mtproto.Channels_SendAsPeers, error) {
	// TODO: not impl
	c.Logger.Errorf("contacts.getTopPeers blocked, License key from https://teamgram.net required to unlock enterprise features.")

	// Disable sendAsPeers
	return mtproto.MakeTLChannelsSendAsPeers(&mtproto.Channels_SendAsPeers{
		Peers: []*mtproto.Peer{},
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	}).To_Channels_SendAsPeers(), nil
}
