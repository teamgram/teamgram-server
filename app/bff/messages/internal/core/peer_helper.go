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
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type resolvedMessagePeer struct {
	PeerType int32
	PeerID   int64
}

func resolveUserPeerID(peer tg.InputPeerClazz, selfUserID int64) (int64, bool) {
	p := tg.FromInputPeer2(selfUserID, peer)
	switch p.PeerType {
	case tg.PEER_SELF, tg.PEER_USER:
		return p.PeerId, p.PeerId > 0
	default:
		return 0, false
	}
}

func resolveMessagePeer(peer tg.InputPeerClazz, selfUserID int64) (resolvedMessagePeer, bool) {
	p := tg.FromInputPeer2(selfUserID, peer)
	switch p.PeerType {
	case tg.PEER_SELF, tg.PEER_USER:
		return resolvedMessagePeer{PeerType: payload.PeerTypeUser, PeerID: p.PeerId}, p.PeerId > 0
	case tg.PEER_CHAT:
		return resolvedMessagePeer{PeerType: payload.PeerTypeChat, PeerID: p.PeerId}, p.PeerId > 0
	default:
		return resolvedMessagePeer{}, false
	}
}

func messagePeerClazz(peerType int32, peerID int64) tg.PeerClazz {
	switch peerType {
	case 0, payload.PeerTypeUser:
		return tg.MakePeerUser(peerID)
	case payload.PeerTypeChat:
		return tg.MakePeerChat(peerID)
	default:
		return nil
	}
}
