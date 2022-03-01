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
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// AccountGetNotifySettings
// account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
func (c *NotificationCore) AccountGetNotifySettings(in *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	var (
		err      error
		settings *mtproto.PeerNotifySettings
	)

	peer := mtproto.FromInputNotifyPeer(c.MD.UserId, in.Peer)
	switch peer.PeerType {
	case mtproto.PEER_USER:
		// TODO(@benqi): check peerUser Exists
	case mtproto.PEER_CHAT:
		// TODO(@benqi): check peerChat exists
	case mtproto.PEER_CHANNEL:
		// TODO(@benqi): check peerChannel exists
	case mtproto.PEER_USERS:
	case mtproto.PEER_CHATS:
	case mtproto.PEER_BROADCASTS:
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("account.updateNotifySettings - error: %v", err)
		return nil, err
	}

	settings, err = c.svcCtx.Dao.UserClient.UserGetNotifySettings(c.ctx, &userpb.TLUserGetNotifySettings{
		UserId:   c.MD.UserId,
		PeerType: peer.PeerType,
		PeerId:   peer.PeerId,
	})
	if err != nil {
		c.Logger.Errorf("getNotifySettings error - %v", err)
		return nil, err
	}

	return settings, err
}
