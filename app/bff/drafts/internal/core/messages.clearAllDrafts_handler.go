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
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MessagesClearAllDrafts
// messages.clearAllDrafts#7e58ee9c = Bool;
func (c *DraftsCore) MessagesClearAllDrafts(in *mtproto.TLMessagesClearAllDrafts) (*mtproto.Bool, error) {
	rValues, err := c.svcCtx.Dao.DialogClient.DialogClearAllDrafts(c.ctx, &dialog.TLDialogClearAllDrafts{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("messages.clearAllDrafts: %v", err)
		return nil, err
	}

	if len(rValues.Datas) == 0 {
		return mtproto.BoolTrue, nil
	}

	// sync
	for _, v := range rValues.Datas {
		syncUpdates := mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDraftMessage(&mtproto.Update{
			Peer_PEER: v.Peer,
			Draft:     v.Draft,
		}).To_Update())

		peer := mtproto.FromPeer(v.Peer)
		switch peer.PeerType {
		case mtproto.PEER_USER:
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
				Id: []int64{c.MD.UserId, peer.PeerId},
			})
			user, _ := users.GetUnsafeUser(c.MD.UserId, peer.PeerId)

			syncUpdates.AddSafeUser(user)
		case mtproto.PEER_CHAT:
			chat, _ := c.svcCtx.Dao.ChatClient.ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
				ChatId: peer.PeerId,
			})

			syncUpdates.AddSafeChat(chat.ToUnsafeChat(c.MD.UserId))
		case mtproto.PEER_CHANNEL:
			if c.svcCtx.Plugin != nil {
				chats := c.svcCtx.Plugin.GetChannelListByIdList(c.ctx, c.MD.UserId, peer.PeerId)
				syncUpdates.PushChat(chats...)
			} else {
				c.Logger.Errorf("messages.clearAllDrafts blocked, License key from https://teamgram.net required to unlock enterprise features.")
				return nil, mtproto.ErrEnterpriseIsBlocked
			}
		}

		c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
			UserId:        c.MD.UserId,
			PermAuthKeyId: c.MD.PermAuthKeyId,
			Updates:       syncUpdates,
		})
	}

	return mtproto.BoolTrue, nil
}
