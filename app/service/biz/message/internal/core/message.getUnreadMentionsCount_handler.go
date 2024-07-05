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
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessageGetUnreadMentionsCount
// message.getUnreadMentionsCount user_id:long peer_type:int peer_id:long = Int32;
func (c *MessageCore) MessageGetUnreadMentionsCount(in *message.TLMessageGetUnreadMentionsCount) (*mtproto.Int32, error) {
	var (
		sz int
	)

	switch in.PeerType {
	case mtproto.PEER_CHAT:
		sz = c.svcCtx.Dao.CommonDAO.CalcSize(
			c.ctx,
			c.svcCtx.Dao.MessagesDAO.CalcTableName(in.UserId),
			map[string]interface{}{
				"user_id":      in.UserId,
				"peer_type":    mtproto.PEER_CHAT,
				"peer_id":      in.PeerId,
				"mentioned":    1,
				"media_unread": 1,
				"deleted":      0,
			})
	case mtproto.PEER_CHANNEL:
		sz = c.svcCtx.Dao.CommonDAO.CalcSize(c.ctx, "channel_unread_mentions", map[string]interface{}{
			"user_id":    in.UserId,
			"channel_id": in.PeerId,
			"deleted":    0,
		})
	default:
		// TODO: log
	}

	return &mtproto.Int32{
		V: int32(sz),
	}, nil
}
