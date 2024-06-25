// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/messenger/msg/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/jsonx"
)

func (d *Dao) DeleteByMessageIdList(ctx context.Context, userId int64, idList []int32) (rowsAffected int64, err error) {
	if len(idList) == 0 {
		return 0, nil
	}
	return d.MessagesDAO.DeleteMessagesByMessageIdList(ctx, userId, idList)
}

func (d *Dao) GetLastMessageAndIdListByDialog(ctx context.Context, userId int64, peer *mtproto.PeerUtil) (lastMessage *mtproto.Message, idList []int32) {
	dialogId := mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
	d.MessagesDAO.SelectDialogMessageIdListWithCB(
		ctx,
		userId,
		dialogId.A,
		dialogId.B,
		func(sz, i int, v *dataobject.MessagesDO) {
			if i == 0 {
				var (
					m = new(mtproto.Message)
				)
				err := jsonx.UnmarshalFromString(v.MessageData, m)
				if err != nil {
					logx.WithContext(ctx).Errorf("error: %v, do: %v", err, v)
				} else {
					lastMessage = m.FixData()
				}
			}

			idList = append(idList, v.UserMessageBoxId)
		})
	return
}

//func (d *Dao) GetPeerMessageId(ctx context.Context, userId, messageId, peerId int32) int32 {
//	//do, _ := d.MessagesDAO.SelectPeerMessageId(ctx, peerId, userId, messageId)
//	//if do == nil {
//	//	return 0
//	//} else {
//	//	return do.UserMessageBoxId
//	//}
//}
//
//func (d *Dao) GetPeerDialogMessageIdList(ctx context.Context, userId int64, idList []int32) map[int64][]int32 {
//	doList, _ := d.MessagesDAO.SelectPeerDialogMessageIdList(ctx, userId, idList)
//	peerMessageIdListMap := make(map[int64][]int32)
//
//	for _, do := range doList {
//		if messageIdList, ok := peerMessageIdListMap[do.UserId]; !ok {
//			peerMessageIdListMap[do.UserId] = []int32{do.UserMessageBoxId}
//		} else {
//			peerMessageIdListMap[do.UserId] = append(messageIdList, do.UserMessageBoxId)
//		}
//	}
//
//	return peerMessageIdListMap
//}

func (d *Dao) GetMessageIdListByDialog(ctx context.Context, userId int64, peer *mtproto.PeerUtil) []int32 {
	var (
		dialogId = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		idList   []int32
	)

	d.MessagesDAO.SelectDialogMessageIdListWithCB(
		ctx,
		userId,
		dialogId.A,
		dialogId.B,
		func(sz, i int, v *dataobject.MessagesDO) {
			idList = append(idList, v.UserMessageBoxId)
		})

	return idList
}
