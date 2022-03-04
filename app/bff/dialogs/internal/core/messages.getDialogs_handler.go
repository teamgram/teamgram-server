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
	"math"
	"sort"
)

// MessagesGetDialogs
// messages.getDialogs#a0f4cb4f flags:# exclude_pinned:flags.0?true folder_id:flags.1?int offset_date:int offset_id:int offset_peer:InputPeer limit:int hash:long = messages.Dialogs;
func (c *DialogsCore) MessagesGetDialogs(in *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	var (
		peer     = mtproto.FromInputPeer2(c.MD.UserId, in.OffsetPeer)
		folderId = in.GetFolderId().GetValue()
		limit    = in.Limit
	)

	if limit > 500 {
		limit = 500
	}

	dialogs, err := c.svcCtx.Dao.DialogClient.DialogGetDialogs(c.ctx, &dialog.TLDialogGetDialogs{
		UserId:        c.MD.UserId,
		ExcludePinned: mtproto.ToBool(in.ExcludePinned),
		FolderId:      folderId,
	})
	if err != nil {
		c.Logger.Errorf("messages.getDialogs - error: %v", err)
		return nil, err
	} else if len(dialogs.Datas) == 0 {
		return mtproto.MakeTLMessagesDialogsSlice(&mtproto.Messages_Dialogs{
			Dialogs:  []*mtproto.Dialog{},
			Messages: []*mtproto.Message{},
			Chats:    []*mtproto.Chat{},
			Users:    []*mtproto.User{},
			Count:    0,
		}).To_Messages_Dialogs(), nil
	}

	var (
		dialogCount                        = int32(len(dialogs.Datas))
		dialogExtList dialog.DialogExtList = dialogs.Datas
	)

	for _, dialogEx := range dialogExtList {
		peer2 := dialogEx.GetDialog().GetPeer()

		if peer2.GetPredicateName() == mtproto.Predicate_peerChannel {
			if c.svcCtx.Plugin != nil {
				dialog, _ := c.svcCtx.Plugin.GetChannelDialogById(c.ctx, c.MD.UserId, peer2.ChannelId)
				if dialog != nil {
					dialogEx.Dialog.TopMessage = dialog.Dialog.TopMessage
					dialogEx.Dialog.Pts = dialog.Dialog.Pts
					// dialogEx.Dialog.ReadOutboxMaxId = channel2.Megagroup.ReadInboxMaxId
					dialogEx.Date = dialog.Date
					dialogEx.Order = dialog.Order
					// TODO:
					// dialogEx.AvailableMinId = megagroup2.GetParticipants()[0].AvailableMinId
				}
			} else {
				c.Logger.Errorf("messages.getDialogs blocked, License key from https://teamgram.net required to unlock enterprise features.")
			}
		}
	}

	r2 := sort.Reverse(dialogExtList)
	sort.Sort(r2)

	//for _, dialog := range dialogs {
	switch peer.PeerType {
	case mtproto.PEER_EMPTY:
		if (in.OffsetId == 0 || in.OffsetId == 2147483647) && (in.OffsetDate == 0 || in.OffsetDate == 2147483647) {
			if len(dialogExtList) >= int(limit) {
				dialogExtList = dialogExtList[:limit]
			}
		} else {
			idx := -1
			if in.OffsetId > 0 && in.OffsetId < math.MaxInt32 {
				for i, dialog := range dialogExtList {
					if dialog.Dialog.TopMessage == in.OffsetId {
						idx = i
						break
					}
				}
			} else if in.OffsetDate > 0 && in.OffsetDate < math.MaxInt32 {
				for i, dialog := range dialogExtList {
					if int32(dialog.Order) == in.OffsetDate {
						idx = i
						break
					}
				}
			}
			if idx > 0 {
				if idx+int(limit) > len(dialogExtList)-2 {
					dialogExtList = dialogExtList[idx+1:]
				} else {
					dialogExtList = dialogExtList[idx+1 : idx+1+int(limit)]
				}
			} else {
				dialogExtList = dialogExtList[:0]
			}
		}
	default:
		idx := -1

		switch peer.PeerType {
		case mtproto.PEER_SELF, mtproto.PEER_USER:
			for i, dialog := range dialogExtList {
				if dialog.Dialog.TopMessage == in.OffsetId &&
					int32(dialog.Order) == in.OffsetDate &&
					dialog.Dialog.Peer.UserId == peer.PeerId {
					idx = i
					break
				}
			}
		case mtproto.PEER_CHAT:
			for i, dialog := range dialogExtList {
				if dialog.Dialog.TopMessage == in.OffsetId &&
					int32(dialog.Order) == in.OffsetDate &&
					dialog.Dialog.Peer.ChatId == peer.PeerId {
					idx = i
					break
				}
			}
		case mtproto.PEER_CHANNEL:
			for i, dialog := range dialogExtList {
				if dialog.Dialog.TopMessage == in.OffsetId &&
					int32(dialog.Order) == in.OffsetDate &&
					dialog.Dialog.Peer.ChannelId == peer.PeerId {
					idx = i
					break
				}
			}
		}
		if idx > 0 {
			if idx+int(limit) > len(dialogExtList)-2 {
				dialogExtList = dialogExtList[idx+1:]
			} else {
				dialogExtList = dialogExtList[idx+1 : idx+1+int(limit)]
			}
		} else {
			dialogExtList = dialogExtList[:0]
		}
	}

	return c.makeMessagesDialogs(dialogExtList).ToMessagesDialogs(dialogCount), nil
}
