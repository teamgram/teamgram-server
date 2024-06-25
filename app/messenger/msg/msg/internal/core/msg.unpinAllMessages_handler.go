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

// MsgUnpinAllMessages
// msg.unpinAllMessages user_id:long auth_key_id:long peer_type:int peer_id:long = messages.AffectedHistory;
func (c *MsgCore) MsgUnpinAllMessages(in *msg.TLMsgUnpinAllMessages) (*mtproto.Messages_AffectedHistory, error) {
	var (
		peer     = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
		dialogId = mtproto.MakeDialogId(in.UserId, peer.PeerType, peer.PeerId)
		idList   = make([]int32, 0)
		pts      int32
		ptsCount int32
	)

	switch peer.PeerType {
	case mtproto.PEER_SELF,
		mtproto.PEER_USER,
		mtproto.PEER_CHAT:
		boxMsgList, err := c.svcCtx.Dao.MessagesDAO.SelectPinnedListWithCB(
			c.ctx,
			in.UserId,
			dialogId.A,
			dialogId.B,
			func(sz, i int, v *dataobject.MessagesDO) {
				idList = append(idList, v.UserMessageBoxId)
			})
		if err != nil {
			c.Logger.Errorf("msg.updatePinnedMessage - error: %v", err)
			return nil, mtproto.ErrMsgIdInvalid
		}
		if len(boxMsgList) == 0 {
			c.Logger.Errorf("msg.updatePinnedMessage - error: %v", err)
			return nil, mtproto.ErrMsgIdInvalid
		}

		c.svcCtx.Dao.DialogsDAO.UpdatePinnedMsgId(c.ctx, 0, in.UserId, mtproto.MakePeerDialogId(peer.PeerType, peer.PeerId))
		c.svcCtx.Dao.MessagesDAO.UpdateUnPinnedByIdList(c.ctx, in.UserId, idList)

		// update
		pts = c.svcCtx.Dao.IDGenClient2.NextNPtsId(c.ctx, in.UserId, len(idList))
		ptsCount = int32(len(idList))
		c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
			c.ctx,
			&sync.TLSyncUpdatesNotMe{
				UserId:        in.UserId,
				PermAuthKeyId: in.AuthKeyId,
				Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdatePinnedMessages(&mtproto.Update{
					Pinned:    false,
					Peer_PEER: peer.ToPeer(),
					Messages:  idList,
					Pts_INT32: pts,
					PtsCount:  ptsCount,
				}).To_Update()),
			})

		// inbox
		c.svcCtx.Dao.InboxClient.InboxUnpinAllMessages(c.ctx, &inbox.TLInboxUnpinAllMessages{
			UserId:   in.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
		})

	case mtproto.PEER_CHANNEL:
	default:
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.updatePinnedMessage - error: %v", err)
		return nil, err
	}

	return mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
		Pts:      pts,
		PtsCount: ptsCount,
		Offset:   0,
	}).To_Messages_AffectedHistory(), nil
}
