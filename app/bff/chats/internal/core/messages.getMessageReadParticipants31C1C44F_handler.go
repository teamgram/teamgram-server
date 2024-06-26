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

package core

import (
	"github.com/teamgram/proto/mtproto"
	chatpb "github.com/teamgram/teamgram-server/app/service/biz/chat/chat"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/message/message"
)

// MessagesGetMessageReadParticipants31C1C44F
// messages.getMessageReadParticipants#31c1c44f peer:InputPeer msg_id:int = Vector<ReadParticipantDate>;
func (c *ChatsCore) MessagesGetMessageReadParticipants31C1C44F(in *mtproto.TLMessagesGetMessageReadParticipants31C1C44F) (*mtproto.Vector_ReadParticipantDate, error) {
	var (
		peer                 = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		readParticipantDates = make([]*mtproto.ReadParticipantDate, 0)
	)

	switch peer.PeerType {
	case mtproto.PEER_CHAT:
		msgBox, err := c.svcCtx.Dao.MessageClient.MessageGetUserMessage(c.ctx, &message.TLMessageGetUserMessage{
			UserId: c.MD.UserId,
			Id:     in.MsgId,
		})
		if err != nil {
			c.Logger.Errorf("messages.getMessageReadParticipants - error: %v", err)
			return nil, err
		}

		pIdList, err := c.svcCtx.Dao.ChatClient.Client().ChatGetChatParticipantIdList(c.ctx, &chatpb.TLChatGetChatParticipantIdList{
			ChatId: peer.PeerId,
		})
		if err != nil {
			c.Logger.Errorf("messages.getMessageReadParticipants - error: %v", err)
			return nil, err
		}

		boxList, err := c.svcCtx.Dao.MessageClient.MessageGetUserMessageListByDataIdUserIdList(c.ctx, &message.TLMessageGetUserMessageListByDataIdUserIdList{
			Id:         msgBox.DialogMessageId,
			UserIdList: pIdList.GetDatas(),
		})
		if err != nil {
			c.Logger.Errorf("messages.getMessageReadParticipants - error: %v", err)
			return nil, err
		}

		// TODO: 性能优化
		boxList.Walk(func(idx int, v *mtproto.MessageBox) {
			if v.UserId == c.MD.UserId {
				return
			}

			dialogList, _ := c.svcCtx.Dao.DialogClient.DialogGetDialogsByIdList(c.ctx, &dialog.TLDialogGetDialogsByIdList{
				UserId: v.UserId,
				IdList: []int64{mtproto.MakePeerDialogId(peer.PeerType, peer.PeerId)},
			})
			for _, d := range dialogList.GetDatas() {
				// c.Logger.Infof("messages.getMessageReadParticipants - dialog: %s", d)
				if d.GetDialog().GetReadInboxMaxId() >= v.MessageId {
					readParticipantDates = append(readParticipantDates, mtproto.MakeTLReadParticipantDate(&mtproto.ReadParticipantDate{
						UserId: v.UserId,
						Date:   0,
					}).To_ReadParticipantDate())
				}
			}
		})
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("messages.getMessageReadParticipants blocked, License key from https://teamgram.net required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getMessageReadParticipants - error: %v", err)
		return nil, err
	}

	// TODO: add readParticipantDates
	return &mtproto.Vector_ReadParticipantDate{
		Datas: readParticipantDates,
	}, nil
}
