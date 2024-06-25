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
	"github.com/teamgram/teamgram-server/app/service/biz/message/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessageSearchByMediaType
// message.searchByMediaType user_id:long peer:PeerUtil media_type:int offset:int limit:int = Vector<MessageBox>;
func (c *MessageCore) MessageSearchByMediaType(in *message.TLMessageSearchByMediaType) (*mtproto.MessageBoxList, error) {
	var (
		boxList []*mtproto.MessageBox
	)

	if in.MediaType == mtproto.MEDIA_PHONE_CALL {
		boxList = c.searchByPhoneCall(in.UserId, in.Offset, in.Limit)
	} else {
		boxList = c.searchByMediaType(in.UserId, in.PeerType, in.PeerId, in.MediaType, in.Offset, in.Limit)
	}

	return mtproto.MakeTLMessageBoxList(&mtproto.MessageBoxList{
		BoxList: boxList,
	}).To_MessageBoxList(), nil
}

func (c *MessageCore) searchByMediaType(
	userId int64,
	peerType int32,
	peerId int64,
	mediaType int32,
	offset, limit int32) (boxList []*mtproto.MessageBox) {

	var (
		dialogId = mtproto.MakeDialogId(userId, peerType, peerId)
	)
	c.svcCtx.Dao.MessagesDAO.SelectByMediaTypeWithCB(
		c.ctx,
		userId,
		dialogId.A,
		dialogId.B,
		mediaType,
		offset,
		limit,
		func(sz, i int, v *dataobject.MessagesDO) {
			boxList = append(boxList, c.svcCtx.Dao.MakeMessageBox(c.ctx, userId, v))
		})

	if boxList == nil {
		boxList = []*mtproto.MessageBox{}
	}

	return
}

func (c *MessageCore) searchByPhoneCall(userId int64, offset, limit int32) (boxList []*mtproto.MessageBox) {
	c.svcCtx.Dao.MessagesDAO.SelectPhoneCallListWithCB(
		c.ctx,
		userId,
		mtproto.MEDIA_PHONE_CALL,
		offset,
		limit,
		func(sz, i int, v *dataobject.MessagesDO) {
			boxList = append(boxList, c.svcCtx.Dao.MakeMessageBox(c.ctx, userId, v))
		})

	if boxList == nil {
		boxList = []*mtproto.MessageBox{}
	}

	return
}
