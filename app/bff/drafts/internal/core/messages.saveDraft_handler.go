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
	"time"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// MessagesSaveDraft
// messages.saveDraft#bc39e14b flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> = Bool;
func (c *DraftsCore) MessagesSaveDraft(in *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
	var (
		peer                = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		draft               *mtproto.DraftMessage
		isDraftMessageEmpty = true
		date                = int32(time.Now().Unix())
	)

	if in.NoWebpage == true {
		isDraftMessageEmpty = false
	} else if in.ReplyToMsgId != nil {
		isDraftMessageEmpty = false
	} else if in.Message != "" {
		isDraftMessageEmpty = false
	} else if in.Entities != nil {
		isDraftMessageEmpty = false
	}

	if isDraftMessageEmpty {
		draft = mtproto.MakeTLDraftMessageEmpty(&mtproto.DraftMessage{
			Date_FLAGINT32: mtproto.MakeFlagsInt32(date),
		}).To_DraftMessage()

		c.svcCtx.Dao.DialogClient.DialogClearDraftMessage(c.ctx, &dialog.TLDialogClearDraftMessage{
			UserId:   c.MD.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
		})
	} else {
		draft = mtproto.MakeTLDraftMessage(&mtproto.DraftMessage{
			NoWebpage:    in.GetNoWebpage(),
			ReplyToMsgId: in.GetReplyToMsgId(),
			Message:      in.GetMessage(),
			Entities:     in.GetEntities(),
			Date_INT32:   date,
		}).To_DraftMessage()

		c.svcCtx.Dao.DialogClient.DialogSaveDraftMessage(c.ctx, &dialog.TLDialogSaveDraftMessage{
			UserId:   c.MD.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
			Message:  draft,
		})
	}

	syncUpdates := mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDraftMessage(&mtproto.Update{
		Peer_PEER: peer.ToPeer(),
		Draft:     draft,
	}).To_Update())

	switch peer.PeerType {
	case mtproto.PEER_SELF:
		user, _ := c.svcCtx.Dao.UserClient.UserGetImmutableUser(c.ctx, &userpb.TLUserGetImmutableUser{
			Id: c.MD.UserId,
		})
		if user != nil {
			syncUpdates.PushUser(user.ToSelfUser())
		}
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
			c.Logger.Errorf("messages.saveDraft blocked, License key from https://teamgram.net required to unlock enterprise features.")
			return nil, mtproto.ErrEnterpriseIsBlocked
		}
	}

	// sync
	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:        c.MD.UserId,
		PermAuthKeyId: c.MD.PermAuthKeyId,
		Updates:       syncUpdates,
	})

	return mtproto.BoolTrue, nil
}
