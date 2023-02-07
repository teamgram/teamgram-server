// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package message

import (
	"github.com/teamgram/proto/mtproto"
)

func (m *Vector_MessageBox) Visit(toUserId int64,
	cb1 func(messageList []*mtproto.Message),
	cb2 func(userIdList []int64),
	cb3 func(chatIdList []int64),
	cb4 func(channelIdList []int64)) {
	var (
		idHelper    = mtproto.NewIDListHelper(toUserId)
		messageList = make([]*mtproto.Message, 0, len(m.GetDatas()))
	)

	for _, box := range m.GetDatas() {
		message := box.ToMessage(toUserId)
		messageList = append(messageList, message)
		idHelper.PickByMessage(message)
	}

	if cb1 != nil {
		cb1(messageList)
	}

	idHelper.Visit(cb2, cb3, cb4)
}

func (m *Vector_MessageBox) Length() int32 {
	return int32(len(m.GetDatas()))
}

func (m *Vector_MessageBox) Walk(cb func(idx int, v *mtproto.MessageBox)) {
	if cb == nil {
		return
	}

	for i, v := range m.Datas {
		cb(i, v)
	}
}
