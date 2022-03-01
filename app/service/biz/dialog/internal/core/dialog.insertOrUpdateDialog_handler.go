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
	"time"
)

// DialogInsertOrUpdateDialog
// dialog.insertOrUpdateDialog flags:# user_id:long peer_type:int peer_id:long top_message:flags.0?int read_outbox_id:flags.1?int read_inbox_max_id:flags.2?int unread_count:flags.3?int unread_mark:flags.4?true date:int = Bool;
func (c *DialogCore) DialogInsertOrUpdateDialog(in *dialog.TLDialogInsertOrUpdateDialog) (*mtproto.Bool, error) {
	var (
		cMap = make(map[string]interface{}, 0)
		date = time.Now().Unix()
	)

	if in.GetTopMessage() != nil {
		cMap["top_message"] = in.GetTopMessage().GetValue()
	}
	if in.GetReadOutboxMaxId() != nil {
		cMap["read_outbox_max_id"] = in.GetReadOutboxMaxId().GetValue()
	}
	if in.GetReadInboxMaxId() != nil {
		cMap["read_inbox_max_id"] = in.GetReadInboxMaxId().GetValue()
	}
	if in.GetUnreadCount() != nil {
		cMap["unread_count"] = in.GetUnreadCount().GetValue()
	}
	if in.GetUnreadMark() {
		cMap["unread_mark"] = 1
	}
	cMap["date2"] = date
	cMap["deleted"] = 0

	rowsAffected, err := c.svcCtx.Dao.DialogsDAO.UpdateCustomMap(
		c.ctx,
		cMap,
		in.UserId,
		in.PeerType,
		in.PeerId)
	if err != nil {
		c.Logger.Errorf("dialog.insertOrUpdateDialog - error: %v", err)
		return nil, err
	}

	if rowsAffected == 0 {
		dlgDO := &dataobject.DialogsDO{
			UserId:           in.UserId,
			PeerType:         in.PeerType,
			PeerId:           in.PeerId,
			PeerDialogId:     mtproto.MakePeerDialogId(in.PeerType, in.PeerId),
			DraftMessageData: "null",
			Date2:            date,
		}

		if in.GetTopMessage() != nil {
			dlgDO.TopMessage = in.GetTopMessage().GetValue()
		}
		if in.GetReadOutboxMaxId() != nil {
			dlgDO.ReadOutboxMaxId = in.GetReadOutboxMaxId().GetValue()
		}
		if in.GetReadInboxMaxId() != nil {
			dlgDO.ReadInboxMaxId = in.GetReadInboxMaxId().GetValue()
		}
		if in.GetUnreadCount() != nil {
			dlgDO.UnreadCount = in.GetUnreadCount().GetValue()
		}
		if in.GetUnreadMark() {
			dlgDO.UnreadMark = false
		}

		c.svcCtx.Dao.DialogsDAO.InsertIgnore(c.ctx, dlgDO)
	}

	return mtproto.BoolTrue, nil
}
