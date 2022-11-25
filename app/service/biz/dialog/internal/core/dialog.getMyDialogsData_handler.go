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
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/dialog"
	"github.com/teamgram/teamgram-server/app/service/biz/dialog/internal/dal/dataobject"
)

// DialogGetMyDialogsData
// dialog.getMyDialogsData flags:# user:flags.0?true chat:flags.1?true channel:flags.2?true = Vector<PeerUtil>;
func (c *DialogCore) DialogGetMyDialogsData(in *dialog.TLDialogGetMyDialogsData) (*dialog.DialogsData, error) {
	dialogsData := dialog.MakeTLSimpleDialogsData(&dialog.DialogsData{
		Users:    []int64{},
		Chats:    []int64{},
		Channels: []int64{},
	}).To_DialogsData()

	isGetAll := false
	if in.User == false && in.Chat == false && in.Channel == false {
		isGetAll = true
	}
	if in.User == true && in.Chat == true && in.Channel == true {
		isGetAll = true
	}

	if isGetAll {
		c.svcCtx.Dao.DialogsDAO.SelectAllDialogsWithCB(
			c.ctx,
			in.UserId,
			func(i int, v *dataobject.DialogsDO) {
				switch v.PeerType {
				case mtproto.PEER_USER:
					dialogsData.Users = append(dialogsData.Users, v.PeerId)
				case mtproto.PEER_CHAT:
					dialogsData.Chats = append(dialogsData.Chats, v.PeerId)
				case mtproto.PEER_CHANNEL:
					dialogsData.Channels = append(dialogsData.Channels, v.PeerId)
				}
			})
	} else {
		var (
			pList []int32
		)

		if in.User {
			pList = append(pList, mtproto.PEER_USER)
		}
		if in.Chat {
			pList = append(pList, mtproto.PEER_CHAT)
		}
		if in.Channel {
			pList = append(pList, mtproto.PEER_CHANNEL)
		}

		c.svcCtx.Dao.DialogsDAO.SelectDialogsByPeerTypeWithCB(
			c.ctx,
			in.UserId,
			pList,
			func(i int, v *dataobject.DialogsDO) {
				switch v.PeerType {
				case mtproto.PEER_USER:
					dialogsData.Users = append(dialogsData.Users, v.PeerId)
				case mtproto.PEER_CHAT:
					dialogsData.Chats = append(dialogsData.Chats, v.PeerId)
				case mtproto.PEER_CHANNEL:
					dialogsData.Channels = append(dialogsData.Channels, v.PeerId)
				}
			})
	}

	return dialogsData, nil
}
