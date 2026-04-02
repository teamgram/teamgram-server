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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// ChannelsGetSendAs
// channels.getSendAs#dc770ee peer:InputPeer = channels.SendAsPeers;
func (c *MessagesCore) ChannelsGetSendAs(in *tg.TLChannelsGetSendAs) (*tg.ChannelsSendAsPeers, error) {
	peer, err := bffPeerFromInput(c, in.Peer)
	if err != nil {
		return nil, err
	}

	return tg.MakeTLChannelsSendAsPeers(&tg.TLChannelsSendAsPeers{
		Peers: []tg.SendAsPeerClazz{
			tg.MakeTLSendAsPeer(&tg.TLSendAsPeer{
				Peer: peer,
			}),
		},
		Chats: []tg.ChatClazz{},
		Users: []tg.UserClazz{},
	}).ToChannelsSendAsPeers(), nil
}
