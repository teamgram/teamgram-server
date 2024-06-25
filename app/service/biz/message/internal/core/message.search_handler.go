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
	"math"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessageSearch
// message.search user_id:long peer:PeerUtil q:string offset:int limit:int = Vector<MessageBox>;
func (c *MessageCore) MessageSearch(in *message.TLMessageSearch) (*mtproto.MessageBoxList, error) {
	var (
		offset  = in.Offset
		q       = in.Q
		boxList []*mtproto.MessageBox
		limit   = in.Limit
	)

	// TODO(@benqi): check q
	if offset == 0 {
		offset = math.MaxInt32
	}
	if limit > 50 {
		limit = 50
	}

	switch in.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		if q[0] == '#' {
			idList, _ := c.svcCtx.Dao.HashTagsDAO.SelectPeerHashTagList(
				c.ctx,
				in.UserId,
				in.PeerType,
				in.PeerId,
				q)

			if len(idList) > 0 {
				c.svcCtx.Dao.MessagesDAO.SelectByMessageIdListWithCB(
					c.ctx,
					in.UserId,
					idList,
					func(sz, i int, v *dataobject.MessagesDO) {
						boxList = append(boxList, c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, v))
					})
			}
		} else {
			dialogId := mtproto.MakeDialogId(in.UserId, in.PeerType, in.PeerId)
			c.svcCtx.Dao.MessagesDAO.SearchWithCB(
				c.ctx,
				in.UserId,
				dialogId.A,
				dialogId.B,
				in.Offset,
				"%"+q+"%",
				in.Limit,
				func(sz, i int, v *dataobject.MessagesDO) {
					boxList = append(boxList, c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, v))
				})
		}
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("message.search blocked, License key from https://teamgram.net required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	if boxList == nil {
		boxList = []*mtproto.MessageBox{}
	}

	return mtproto.MakeTLMessageBoxList(&mtproto.MessageBoxList{
		BoxList: boxList,
	}).To_MessageBoxList(), nil
}
