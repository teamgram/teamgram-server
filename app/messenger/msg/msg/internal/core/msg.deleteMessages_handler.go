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
	"github.com/teamgram/teamgram-server/app/messenger/msg/msg/msg"
	"github.com/teamgram/teamgram-server/app/messenger/sync/sync"
)

// MsgDeleteMessages
// msg.deleteMessages user_id:long auth_key_id:long peer_type:int peer_id:long revoke:Bool id:Vector<int> = messages.AffectedMessages;
func (c *MsgCore) MsgDeleteMessages(in *msg.TLMsgDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	var (
		rValue *mtproto.Messages_AffectedMessages
		err    error
	)

	if in.UserId == 0 {
		err = mtproto.ErrInputRequestInvalid
		c.Logger.Errorf("msg.deleteMessages - error: %v", err)
		return nil, err
	}

	switch in.PeerType {
	case mtproto.PEER_EMPTY:
		if len(in.Id) == 0 {
			rValue = mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
				Pts:      c.svcCtx.Dao.IDGenClient2.CurrentPtsId(c.ctx, in.PeerId),
				PtsCount: 0,
			}).To_Messages_AffectedMessages()
		} else {
			rValue, err = c.deleteUserMessages(in)
			if err != nil {
				c.Logger.Errorf("msg.deleteMessages - error: method MsgDeleteMessages not impl")
				return nil, err
			}
		}
	case mtproto.PEER_CHANNEL:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("DeleteMessages - error: %v", err)
		return nil, err
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("DeleteMessages - error: %v", err)
		return nil, err
	}

	return rValue, nil
}

func (c *MsgCore) deleteUserMessages(in *msg.TLMsgDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	var (
		pts, ptsCount int32
		msgDataIdList []int64
	)

	peer, msgDataIdList, err := c.svcCtx.Dao.DeleteMessages(c.ctx, in.UserId, in.Id)
	if err != nil {
		c.Logger.Errorf("DeleteMessages - %v", err)
		return nil, err
	} else if len(msgDataIdList) == 0 {
		return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
			Pts:      c.svcCtx.Dao.IDGenClient2.CurrentPtsId(c.ctx, in.UserId),
			PtsCount: 0,
		}).To_Messages_AffectedMessages(), nil
	}

	pts = c.svcCtx.Dao.IDGenClient2.NextNPtsId(c.ctx, in.UserId, len(in.Id))
	ptsCount = int32(len(in.Id))

	// me
	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:        in.UserId,
		PermAuthKeyId: in.AuthKeyId,
		Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
			Messages:  in.Id,
			Pts_INT32: pts,
			PtsCount:  ptsCount,
		}).To_Update()),
	})

	if in.Revoke {
		c.svcCtx.Dao.InboxClient.InboxDeleteMessagesToInbox(
			c.ctx,
			&inbox.TLInboxDeleteMessagesToInbox{
				FromId:   in.UserId,
				PeerType: peer.PeerType,
				PeerId:   peer.PeerId,
				Id:       msgDataIdList,
			})
	}

	return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	}).To_Messages_AffectedMessages(), nil
}
