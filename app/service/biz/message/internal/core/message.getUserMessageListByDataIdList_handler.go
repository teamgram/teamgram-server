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

// MessageGetUserMessageListByDataIdList
// message.getUserMessageListByDataIdList user_id:long id_list:Vector<long> = Vector<MessageBox>;
func (c *MessageCore) MessageGetUserMessageListByDataIdList(in *message.TLMessageGetUserMessageListByDataIdList) (*message.Vector_MessageBox, error) {
	rValueList := &message.Vector_MessageBox{
		Datas: make([]*mtproto.MessageBox, 0, len(in.IdList)),
	}

	c.svcCtx.Dao.MessagesDAO.SelectByMessageDataIdListWithCB(
		c.ctx,
		c.svcCtx.Dao.MessagesDAO.CalcTableName(in.UserId),
		in.IdList,
		func(sz, i int, v *dataobject.MessagesDO) {
			rValueList.Datas = append(rValueList.GetDatas(), c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, v))
		})

	return rValueList, nil
}
