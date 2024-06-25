/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"
)

// android client request source code
/*
	TLRPC.TL_messages_getDialogs req = new TLRPC.TL_messages_getDialogs();
	req.limit = count;
	req.exclude_pinned = true;
	if (UserConfig.dialogsLoadOffsetId != -1) {
		if (UserConfig.dialogsLoadOffsetId == Integer.MAX_VALUE) {
			dialogsEndReached = true;
			serverDialogsEndReached = true;
			loadingDialogs = false;
			NotificationCenter.getInstance().postNotificationName(NotificationCenter.dialogsNeedReload);
			return;
		}
		req.offset_id = UserConfig.dialogsLoadOffsetId;
		req.offset_date = UserConfig.dialogsLoadOffsetDate;
		if (req.offset_id == 0) {
			req.offset_peer = new TLRPC.TL_inputPeerEmpty();
		} else {
			if (UserConfig.dialogsLoadOffsetChannelId != 0) {
				req.offset_peer = new TLRPC.TL_inputPeerChannel();
				req.offset_peer.channel_id = UserConfig.dialogsLoadOffsetChannelId;
			} else if (UserConfig.dialogsLoadOffsetUserId != 0) {
				req.offset_peer = new TLRPC.TL_inputPeerUser();
				req.offset_peer.user_id = UserConfig.dialogsLoadOffsetUserId;
			} else {
				req.offset_peer = new TLRPC.TL_inputPeerChat();
				req.offset_peer.chat_id = UserConfig.dialogsLoadOffsetChatId;
			}
			req.offset_peer.access_hash = UserConfig.dialogsLoadOffsetAccess;
		}
	} else {
		boolean found = false;
		for (int a = dialogs.size() - 1; a >= 0; a--) {
			TLRPC.TL_dialog dialog = dialogs.get(a);
			if (dialog.pinned) {
				continue;
			}
			int lower_id = (int) dialog.id;
			int high_id = (int) (dialog.id >> 32);
			if (lower_id != 0 && high_id != 1 && dialog.top_message > 0) {
				MessageObject message = dialogMessage.get(dialog.id);
				if (message != null && message.getId() > 0) {
					req.offset_date = message.messageOwner.date;
					req.offset_id = message.messageOwner.id;
					int id;
					if (message.messageOwner.to_id.channel_id != 0) {
						id = -message.messageOwner.to_id.channel_id;
					} else if (message.messageOwner.to_id.chat_id != 0) {
						id = -message.messageOwner.to_id.chat_id;
					} else {
						id = message.messageOwner.to_id.user_id;
					}
					req.offset_peer = getInputPeer(id);
					found = true;
					break;
				}
			}
		}
		if (!found) {
			req.offset_peer = new TLRPC.TL_inputPeerEmpty();
		}
	}
*/

// DialogGetDialogs
// dialog.getDialogs user_id:long exclude_pinned:Bool folder_id:int = Vector<DialogExt>;
func (c *DialogCore) DialogGetDialogs(in *dialog.TLDialogGetDialogs) (*dialog.Vector_DialogExt, error) {
	var (
		dList         dialog.DialogExtList
		excludePinned = mtproto.FromBool(in.GetExcludePinned())
		folderId      = in.GetFolderId()
		meId          = in.GetUserId()
	)

	if excludePinned {
		if folderId == 0 {
			_ = dList
			c.svcCtx.Dao.DialogsDAO.SelectExcludePinnedDialogsWithCB(
				c.ctx,
				meId,
				func(sz, i int, v *dataobject.DialogsDO) {
					dList = append(dList, makeDialog(v))
				})
		} else {
			c.svcCtx.Dao.DialogsDAO.SelectExcludeFolderPinnedDialogsWithCB(
				c.ctx,
				meId,
				func(sz, i int, v *dataobject.DialogsDO) {
					dList = append(dList, makeDialog(v))
				})
		}
	} else {
		c.svcCtx.Dao.DialogsDAO.SelectDialogsWithCB(
			c.ctx,
			in.GetUserId(),
			in.GetFolderId(),
			func(sz, i int, v *dataobject.DialogsDO) {
				dList = append(dList, makeDialog(v))
			})
	}

	return &dialog.Vector_DialogExt{
		Datas: dList,
	}, nil
}
