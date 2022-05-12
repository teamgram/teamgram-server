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
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

/*
// android client source code.

```
if (!messageObject.scheduled) {
	TLRPC.TL_messages_getMessageEditData req = new TLRPC.TL_messages_getMessageEditData();
	req.peer = getMessagesController().getInputPeer((int) dialog_id);
	req.id = messageObject.getId();
	editingMessageObjectReqId = getConnectionsManager().sendRequest(req, (response, error) -> AndroidUtilities.runOnUIThread(() -> {
		editingMessageObjectReqId = 0;
		if (response == null) {
			if (getParentActivity() == null) {
				return;
			}
			AlertDialog.Builder builder = new AlertDialog.Builder(getParentActivity());
			builder.setTitle(LocaleController.getString("AppName", R.string.AppName));
			builder.setMessage(LocaleController.getString("EditMessageError", R.string.EditMessageError));
			builder.setPositiveButton(LocaleController.getString("OK", R.string.OK), null);
			showDialog(builder.create());

			if (chatActivityEnterView != null) {
				chatActivityEnterView.setEditingMessageObject(null, false);
				hideFieldPanel(true);
			}
		} else {
			if (chatActivityEnterView != null) {
				chatActivityEnterView.showEditDoneProgress(false, true);
			}
		}
	}));

```
*/

// MessagesGetMessageEditData
// messages.getMessageEditData#fda68d36 peer:InputPeer id:int = messages.MessageEditData;
func (c *MessagesCore) MessagesGetMessageEditData(in *mtproto.TLMessagesGetMessageEditData) (*mtproto.Messages_MessageEditData, error) {
	var (
		peer   = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		boxMsg *mtproto.MessageBox
		err    error
	)

	// 400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this.
	// 403	MESSAGE_AUTHOR_REQUIRED	Message author required.
	// 400	MESSAGE_ID_INVALID	The provided message id is invalid.
	// 400	PEER_ID_INVALID	The provided peer id is invalid.

	if c.MD.IsBot {
		err = mtproto.ErrBotMethodInvalid
		c.Logger.Errorf("messages.getMessageEditData - error: %v", err)
		return nil, err
	}

	switch peer.PeerType {
	case mtproto.PEER_SELF,
		mtproto.PEER_USER,
		mtproto.PEER_CHAT:
		boxMsg, err = c.svcCtx.Dao.MessageClient.MessageGetUserMessage(c.ctx, &message.TLMessageGetUserMessage{
			UserId: c.MD.UserId,
			Id:     in.Id,
		})
		if err != nil {
			c.Logger.Errorf("messages.getMessageEditData - error: %v", err)
			return nil, err
		}

		if boxMsg.PeerType != mtproto.PEER_USER && boxMsg.PeerId != peer.PeerId {
			err = mtproto.ErrMessageAuthorRequired
			c.Logger.Errorf("messages.getMessageEditData - error: %v", err)
			return nil, err
		}

		if peer.PeerType == mtproto.PEER_CHAT {
			//// TODO: need chat admin
			if c.MD.UserId != boxMsg.SenderUserId {
				mChat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetChatBySelfId(c.ctx, &chatpb.TLChatGetChatBySelfId{
					SelfId: c.MD.UserId,
					ChatId: peer.PeerId,
				})
				if err != nil {
					c.Logger.Errorf("messages.getMessageEditData - error: %v", err)
					return nil, mtproto.ErrChatAdminRequired
				}
				me, _ := mChat.GetImmutableChatParticipant(c.MD.UserId)
				if me == nil {
					c.Logger.Errorf("messages.getMessageEditData - error: %v", err)
					return nil, mtproto.ErrChatAdminRequired
				}
				if !me.CanAdminEditMessages() {
					err = mtproto.ErrChatAdminRequired
					c.Logger.Errorf("messages.getMessageEditData - error: %v", err)
					return nil, mtproto.ErrChatAdminRequired
				}
			}
		} else {
			if c.MD.UserId != boxMsg.SenderUserId {
				err = mtproto.ErrMessageAuthorRequired
				c.Logger.Errorf("messages.getMessageEditData - error: %v", err)
				return nil, err
			}
		}
	case mtproto.PEER_CHANNEL:
		// TODO: not impl
		c.Logger.Errorf("messages.getMessageEditData blocked, License key from https://teamgram.net required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getMessageEditData - error: %v", err)
		return nil, err
	}

	_ = boxMsg
	return mtproto.MakeTLMessagesMessageEditData(&mtproto.Messages_MessageEditData{
		Caption: false,
	}).To_Messages_MessageEditData(), nil
}
