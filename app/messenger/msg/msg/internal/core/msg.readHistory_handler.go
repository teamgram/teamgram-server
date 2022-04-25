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
	"github.com/teamgram/teamgram-server/app/messenger/msg/inbox/inbox"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
)

// MsgReadHistory
// msg.readHistory user_id:long auth_key_id:long peer_type:int peer_id:long max_id:int = messages.AffectedMessages;
func (c *MsgCore) MsgReadHistory(in *msg.TLMsgReadHistory) (*mtproto.Messages_AffectedMessages, error) {
	var (
		pts, ptsCount int32
		maxId               = in.MaxId
		peerDialogId        = mtproto.MakePeerDialogId(in.PeerType, in.PeerId)
		sendMe              = false
		did                 = mtproto.MakeDialogId(in.UserId, in.PeerType, in.PeerId)
		senderId      int64 = 0
	)

	// 消息已读逻辑
	// 1. inbox，设置unread_count为0以及read_inbox_max_id
	if maxId == 0 || maxId >= 1000000000 {
		c.svcCtx.Dao.SelectDialogLastMessageListWithCB(
			c.ctx,
			in.UserId,
			did.A,
			did.B,
			1,
			func(i int, v *dataobject.MessagesDO) {
				senderId = v.SenderUserId
			})
		// TODO: check error
	} else {
		v, _ := c.svcCtx.Dao.MessagesDAO.SelectByMessageId(c.ctx, in.UserId, maxId)
		if v != nil {
			senderId = v.SenderUserId
		}
		// TODO: check error
	}
	switch in.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER:
		sendMe = in.UserId == in.PeerId
		if maxId == 0 || maxId >= 1000000000 {
			// topMessage, err :=
			_, err := c.svcCtx.Dao.DialogsDAO.SelectPeerDialogListWithCB(
				c.ctx,
				in.UserId,
				[]int64{peerDialogId},
				func(i int, v *dataobject.DialogsDO) {
					maxId = v.TopMessage
				},
			)
			if err != nil {
				c.Logger.Errorf("msg.readHistory - error: %v", err)
				return nil, mtproto.ErrMsgIdInvalid
			} else if maxId == 0 {
				c.Logger.Errorf("msg.readHistory - error: not found peer_dialog_id")
				return nil, mtproto.ErrMsgIdInvalid
			}
		}
		c.svcCtx.Dao.DialogsDAO.UpdateReadInboxMaxId(c.ctx, in.MaxId, in.UserId, peerDialogId)
	case mtproto.PEER_CHAT:
		if maxId == 0 || maxId >= 1000000000 {
			_, err := c.svcCtx.Dao.DialogsDAO.SelectPeerDialogListWithCB(
				c.ctx,
				in.UserId,
				[]int64{peerDialogId},
				func(i int, v *dataobject.DialogsDO) {
					maxId = v.TopMessage
				},
			)
			if err != nil {
				c.Logger.Errorf("msg.readHistory - error: %v", err)
				return nil, mtproto.ErrMsgIdInvalid
			} else if maxId == 0 {
				c.Logger.Errorf("msg.readHistory - error: not found peer_dialog_id")
				return nil, mtproto.ErrMsgIdInvalid
			}
		}

		c.svcCtx.Dao.DialogsDAO.UpdateReadInboxMaxId(c.ctx, maxId, in.UserId, peerDialogId)
	default:
		c.Logger.Errorf("messages.readHistory - error: invalid peer %v", in)
		err := mtproto.ErrPeerIdInvalid
		return nil, err
	}

	//
	pts = c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, in.UserId)
	ptsCount = 1

	// syncNotMe
	syncUpdates := mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
		Peer_PEER: mtproto.MakePeer(in.PeerType, in.PeerId),
		MaxId:     maxId,
		Pts_INT32: pts,
		PtsCount:  ptsCount,
	}).To_Update())

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:    in.UserId,
		AuthKeyId: in.AuthKeyId,
		Updates:   syncUpdates,
	})

	if !sendMe {
		c.svcCtx.Dao.InboxClient.InboxUpdateHistoryReaded(c.ctx, &inbox.TLInboxUpdateHistoryReaded{
			FromId:   in.UserId,
			PeerType: in.PeerType,
			PeerId:   in.PeerId,
			MaxId:    maxId,
			Sender:   senderId,
		})
	}

	// result
	return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	}).To_Messages_AffectedMessages(), nil
}
