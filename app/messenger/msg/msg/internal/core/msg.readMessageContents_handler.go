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
)

// MsgReadMessageContents
// msg.readMessageContents user_id:long auth_key_id:long peer_type:int peer_id:long id:Vector<ContentMessage> = messages.AffectedMessage;
func (c *MsgCore) MsgReadMessageContents(in *msg.TLMsgReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	var (
		pts, ptsCount int32
	)

	affected, _ := c.readMentionedMessageContents(in)
	ptsCount += affected
	affected, _ = c.readMediaUnreadMessageContents(in)
	ptsCount += affected
	affected, _ = c.readReactionUnreadMessageContents(in)
	ptsCount += affected

	if ptsCount > 0 {
		pts = c.svcCtx.Dao.IDGenClient2.NextNPtsId(c.ctx, in.UserId, int(ptsCount))
	} else {
		ptsCount = 0
		pts = c.svcCtx.Dao.IDGenClient2.CurrentPtsId(c.ctx, in.UserId)
	}

	return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	}).To_Messages_AffectedMessages(), nil
}

func (c *MsgCore) readMentionedMessageContents(in *msg.TLMsgReadMessageContents) (int32, error) {
	var (
		ptsCount int32 = 0
	)

	switch in.PeerType {
	case mtproto.PEER_USER:
		return 0, nil
	case mtproto.PEER_CHAT:
		for _, m := range in.Id {
			if m.Mentioned {
				ptsCount++
				c.svcCtx.Dao.MessagesDAO.UpdateMentionedAndMediaUnread(c.ctx, in.UserId, m.Id) //UpdateMentioned()
			}
		}

		return ptsCount, nil
	default:
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("DeleteMessages - error: %v", err)

		return 0, err
	}
}

func (c *MsgCore) readMediaUnreadMessageContents(in *msg.TLMsgReadMessageContents) (int32, error) {
	var (
		ptsCount int32 = 0
	)

	switch in.PeerType {
	case mtproto.PEER_USER:
		// id := make([]*inbox.InboxMessageId, 0, len(in.Id))
		for _, m := range in.Id {
			if m.MediaUnread {
				ptsCount++
				c.svcCtx.Dao.MessagesDAO.UpdateMediaUnread(c.ctx, in.UserId, m.Id)
				if in.UserId != in.PeerId {
					c.svcCtx.Dao.InboxClient.InboxReadMediaUnreadToInboxV2(
						c.ctx, &inbox.TLInboxReadMediaUnreadToInboxV2{
							UserId:          in.PeerId,
							PeerType:        mtproto.PEER_USER,
							PeerId:          in.UserId,
							DialogMessageId: m.DialogMessageId,
						})
				}
			}
		}

		return ptsCount, nil
	case mtproto.PEER_CHAT:
		// TODO: update sender
		for _, m := range in.Id {
			if m.MediaUnread {
				ptsCount++
				c.svcCtx.Dao.MessagesDAO.UpdateMediaUnread(c.ctx, in.UserId, m.Id)
				c.svcCtx.Dao.InboxClient.InboxReadMediaUnreadToInboxV2(
					c.ctx, &inbox.TLInboxReadMediaUnreadToInboxV2{
						UserId:          m.SendUserId,
						PeerType:        mtproto.PEER_CHAT,
						PeerId:          in.PeerId,
						DialogMessageId: m.DialogMessageId,
					})
			}
		}

		return ptsCount, nil
	default:
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("DeleteMessages - error: %v", err)

		return 0, err
	}
}

func (c *MsgCore) readReactionUnreadMessageContents(in *msg.TLMsgReadMessageContents) (int32, error) {
	for _, m := range in.Id {
		if m.Reaction {
			if c.svcCtx.MsgPlugin != nil {
				c.svcCtx.MsgPlugin.ReadReactionUnreadMessage(c.ctx, in.UserId, m.Id)
			}
		}
	}

	return 0, nil
}
