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
	"github.com/teamgram/teamgram-server/v2/app/bff/notification/internal/repository"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AccountGetNotifySettings
// account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
func (c *NotificationCore) AccountGetNotifySettings(in *tg.TLAccountGetNotifySettings) (*tg.PeerNotifySettings, error) {
	peerUtil := fromInputNotifyPeer(c.MD.UserId, in.Peer)

	rValues, err := c.svcCtx.Repo.UserClient.UserGetNotifySettings(c.ctx, &repository.GetNotifySettings{
		UserId:   c.MD.UserId,
		PeerType: peerUtil.PeerType,
		PeerId:   peerUtil.PeerId,
	})
	if err != nil {
		c.Logger.Errorf("account.getNotifySettings - error: %v", err)
		return nil, err
	}

	return rValues, nil
}

func fromInputNotifyPeer(selfId int64, peer tg.InputNotifyPeerClazz) *tg.TLPeerUtil {
	p := &tg.TLPeerUtil{
		PeerType: tg.PEER_UNKNOWN,
	}

	if peer == nil {
		return p
	}

	switch c := peer.(type) {
	case *tg.TLInputNotifyPeer:
		p2 := tg.FromInputPeer2(selfId, c.Peer)
		return p2.ToPeerUtil()
	case *tg.TLInputNotifyUsers:
		p.PeerType = tg.PEER_USERS
		p.PeerId = 0
		p.AccessHash = 0
	case *tg.TLInputNotifyChats:
		p.PeerType = tg.PEER_CHATS
		p.PeerId = 0
		p.AccessHash = 0
	case *tg.TLInputNotifyBroadcasts:
		p.PeerType = tg.PEER_BROADCASTS
		p.PeerId = 0
		p.AccessHash = 0
	}

	return p
}
