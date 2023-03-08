// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package chat

import (
	"github.com/teamgram/proto/mtproto"
)

func (m *Vector_MutableChat) Length() int32 {
	return int32(len(m.GetDatas()))
}

func (m *Vector_MutableChat) GetChatListByIdList(selfId int64, id ...int64) []*mtproto.Chat {
	if m == nil {
		return []*mtproto.Chat{}
	}

	chatList := make([]*mtproto.Chat, 0, len(m.Datas))
	for _, chat2 := range m.Datas {
		chatList = append(chatList, chat2.ToUnsafeChat(selfId))
	}

	return chatList
}

func (m *Vector_MutableChat) Walk(cb func(idx int, v *mtproto.MutableChat)) {
	if m == nil {
		return
	}

	if cb == nil {
		return
	}

	for i, v := range m.Datas {
		cb(i, v)
	}
}
