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
	"context"
	"fmt"

	"github.com/teamgram/marmota/pkg/threading2"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
)

// MsgReadHistory
// msg.readHistory user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (c *MsgCore) MsgReadHistory(in *msg.TLMsgReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	var (
		pts, ptsCount int32
		maxId               = in.MaxId
		did                 = mtproto.MakeDialogId(in.UserId, in.PeerType, in.PeerId)
		senderId      int64 = 0
		unreadCount   int32 = 0
		peerDialogId        = mtproto.MakePeerDialogId(in.PeerType, in.PeerId)
	)

	dlg, err := c.svcCtx.Dao.DialogsDAO.SelectDialog(c.ctx, in.UserId, in.PeerType, in.PeerId)
	if err != nil {
		c.Logger.Errorf("messages.readHistory - error: invalid peer %v", err)
		return nil, mtproto.ErrInternalServerError
	} else if dlg == nil {
		c.Logger.Errorf("messages.readHistory - error: not found dialog, request: %s", in)
		return nil, mtproto.ErrPeerIdInvalid
	}

	if maxId == 0 || maxId >= 1000000000 {
		maxId = dlg.TopMessage
	}

	// inbox readed
	if dlg.UnreadCount > 0 || maxId > dlg.ReadInboxMaxId {
		maxInboxMsg, err3 := c.svcCtx.Dao.MessagesDAO.SelectByMessageId(c.ctx, in.UserId, maxId)
		if err3 != nil {
			c.Logger.Errorf("messages.readHistory - error: not found dialog(%d,%d), error is %v", in.UserId, maxId, err3)
			return nil, mtproto.ErrInternalServerError
		} else if maxInboxMsg == nil {
			c.Logger.Errorf("messages.readHistory - error: not found dialog(%d,%d)", in.UserId, maxId)
			return nil, mtproto.ErrMsgIdInvalid
		}

		senderId = maxInboxMsg.SenderUserId
		if maxInboxMsg.SenderUserId == in.UserId {
			maxId = 0
		}
	}
	// inbox readed
	if dlg.ReadInboxMaxId >= maxId || dlg.UnreadCount == 0 {
		pts = c.svcCtx.Dao.IDGenClient2.CurrentPtsId(c.ctx, in.UserId)
		ptsCount = 0
		return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
			Pts:      pts,
			PtsCount: 0,
		}).To_Messages_AffectedMessages(), nil
	}

	if maxId > dlg.ReadInboxMaxId {
		readCount := c.svcCtx.Dao.CommonDAO.CalcSizeByWhere(
			c.ctx,
			c.svcCtx.Dao.MessagesDAO.CalcTableName(in.UserId),
			fmt.Sprintf("user_id = %d AND dialog_id1 = %d AND dialog_id2 = %d AND sender_user_id <> %d AND user_message_box_id > %d AND user_message_box_id <= %d AND deleted = 0",
				in.UserId, did.A, did.B, in.UserId, dlg.ReadInboxMaxId, maxId))
		unreadCount = dlg.UnreadCount - int32(readCount)
		if unreadCount < 0 {
			unreadCount = 0
		}
	}

	c.svcCtx.Dao.DialogsDAO.UpdateReadInboxMaxId(c.ctx, unreadCount, maxId, in.UserId, peerDialogId)

	//
	pts = c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, in.UserId)
	ptsCount = 1

	// result
	return threading2.WrapperGoFunc(
		c.ctx,
		mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
			Pts:      pts,
			PtsCount: ptsCount,
		}).To_Messages_AffectedMessages(),
		func(ctx context.Context) {
			// syncNotMe
			syncUpdates := mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
				Peer_PEER: mtproto.MakePeer(in.PeerType, in.PeerId),
				MaxId:     maxId,
				Pts_INT32: pts,
				PtsCount:  ptsCount,
			}).To_Update())

			c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(ctx, &sync.TLSyncUpdatesNotMe{
				UserId:        in.UserId,
				PermAuthKeyId: in.AuthKeyId,
				Updates:       syncUpdates,
			})

			if maxId > dlg.ReadOutboxMaxId {
				c.svcCtx.Dao.InboxClient.InboxUpdateHistoryReaded(ctx, &inbox.TLInboxUpdateHistoryReaded{
					FromId:   in.UserId,
					PeerType: in.PeerType,
					PeerId:   in.PeerId,
					MaxId:    maxId,
					Sender:   senderId,
				})
			}
		}).(*mtproto.Messages_AffectedMessages), nil
}
