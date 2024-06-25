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
	"github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/chat/internal/dal/dataobject"
)

// ChatGetUsersChatIdList
// chat.getUsersChatIdList id:Vector<long> = Vector<UserChatIdList>;
func (c *ChatCore) ChatGetUsersChatIdList(in *chat.TLChatGetUsersChatIdList) (*chat.Vector_UserChatIdList, error) {
	var (
		rValueList = make([]*chat.UserChatIdList, 0, len(in.Id))
	)

	c.svcCtx.Dao.ChatParticipantsDAO.SelectUsersChatIdListWithCB(
		c.ctx,
		in.Id,
		func(sz, i int, v *dataobject.ChatParticipantsDO) {
			found := false
			for _, ch := range rValueList {
				if ch.UserId == v.UserId {
					ch.ChatIdList = append(ch.ChatIdList, v.ChatId)
					found = true
					return
				}
			}
			if !found {
				rValueList = append(rValueList, &chat.UserChatIdList{
					UserId:     v.UserId,
					ChatIdList: []int64{v.ChatId},
				})
			}
		})

	return &chat.Vector_UserChatIdList{
		Datas: rValueList,
	}, nil
}
