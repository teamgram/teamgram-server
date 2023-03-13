// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dialog

import (
	"context"
	"math"
	"sync"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/mr"
)

type (
	DialogExtList []*DialogExt
)

func (m DialogExtList) Len() int {
	return len(m)
}

func (m DialogExtList) Swap(i, j int) {
	m[j], m[i] = m[i], m[j]
}

func (m DialogExtList) Less(i, j int) bool {
	// TODO(@benqi): if date[i] == date[j]
	return m[i].Order < m[j].Order
}

func (m DialogExtList) ChannelIdList() (idList []int64) {
	for _, dlgExt := range m {
		peer := dlgExt.GetDialog().GetPeer()
		if mtproto.PeerIsChannel(peer) {
			idList = append(idList, peer.ChannelId)
		}
	}

	return
}

func (m DialogExtList) GetDialogsByOffsetLimit(offsetDate int32, offsetId int32, offsetPeer *mtproto.PeerUtil, limit int32) DialogExtList {
	var (
		dialogExtList2 DialogExtList
	)

	if offsetPeer.IsEmpty() {
		if (offsetId == 0 || offsetId == 2147483647) && (offsetDate == 0 || offsetDate == 2147483647) {
			if m.Len() >= int(limit) {
				dialogExtList2 = m[:limit]
			} else {
				dialogExtList2 = m
			}
		} else {
			idx := -1
			if offsetDate > 0 && offsetDate < math.MaxInt32 {
				for i, dialog := range m {
					if int32(dialog.Order) == offsetDate {
						idx = i
						// logx.Debugf("idx: %v", idx)
						break
					}
				}
			} else if offsetId > 0 && offsetId < math.MaxInt32 {
				for i, dialog := range m {
					if dialog.Dialog.TopMessage == offsetId {
						idx = i
						// logx.Debugf("idx: %v", idx)
						break
					}
				}
			}
			if idx > 0 {
				if idx+1+int(limit) > m.Len() {
					dialogExtList2 = m[idx+1:]
				} else {
					dialogExtList2 = m[idx+1 : idx+1+int(limit)]
				}
			} else {
				dialogExtList2 = m[:0]
			}
		}
	} else {
		idx := -1

		switch offsetPeer.PeerType {
		case mtproto.PEER_SELF, mtproto.PEER_USER:
			for i, dialog := range m {
				if dialog.Dialog.TopMessage == offsetId &&
					// int32(dialog.Order) == offsetDate &&
					dialog.Dialog.Peer.UserId == offsetPeer.PeerId {
					idx = i
					break
				}
			}
		case mtproto.PEER_CHAT:
			for i, dialog := range m {
				if dialog.Dialog.TopMessage == offsetId &&
					// int32(dialog.Order) == offsetDate &&
					dialog.Dialog.Peer.ChatId == offsetPeer.PeerId {
					idx = i
					break
				}
			}
		case mtproto.PEER_CHANNEL:
			for i, dialog := range m {
				if dialog.Dialog.TopMessage == offsetId &&
					// int32(dialog.Order) == offsetDate &&
					dialog.Dialog.Peer.ChannelId == offsetPeer.PeerId {
					idx = i
					break
				}
			}
		}
		if idx > 0 {
			if idx+int(limit) > m.Len()-2 {
				dialogExtList2 = m[idx+1:]
			} else {
				dialogExtList2 = m[idx+1 : idx+1+int(limit)]
			}
		} else {
			dialogExtList2 = m[:0]
		}
	}

	return dialogExtList2
}

func (m DialogExtList) DebugString() string {
	s, _ := jsonx.MarshalToString(m)
	return s
}

type TopMessageId struct {
	Peer       *mtproto.PeerUtil
	TopMessage int32
}

type PeerIdList struct {
	PeerType int
	IdList   []int64
}

func (m DialogExtList) DoGetMessagesDialogs(
	ctx context.Context,
	selfUserId int64,
	// cbNotifySettings func(ctx context.Context, selfUserId int64, peers ...mtproto.PeerUtil) []*mtproto.PeerNotifySettings,
	cbMsgF func(ctx context.Context, selfUserId int64, id ...TopMessageId) []*mtproto.Message,
	cbUserF func(ctx context.Context, selfUserId int64, id ...int64) []*mtproto.User,
	cbChatF func(ctx context.Context, selfUserId int64, id ...int64) []*mtproto.Chat,
	cbChannelF func(ctx context.Context, selfUserId int64, id ...int64) []*mtproto.Chat) *DialogsDataHelper {

	dialogsData := &DialogsDataHelper{
		Dialogs:  []*mtproto.Dialog{},
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}

	if len(m) == 0 {
		return dialogsData
	}

	var (
		idHelper = mtproto.NewIDListHelper(selfUserId)
		mu       sync.Mutex
	)

	mr.ForEach(
		func(source chan<- interface{}) {
			var (
				idList   []TopMessageId
				chIdList []TopMessageId
			)

			for _, dialogExt := range m {
				peer2 := mtproto.FromPeer(dialogExt.Dialog.Peer)
				if peer2.IsChannel() {
					chIdList = append(chIdList, TopMessageId{Peer: peer2, TopMessage: dialogExt.Dialog.TopMessage})
				} else {
					idList = append(idList, TopMessageId{Peer: peer2, TopMessage: dialogExt.Dialog.TopMessage})
				}
			}
			if len(idList) > 0 {
				source <- idList
			}
			if len(chIdList) > 0 {
				source <- chIdList
			}
		},
		func(item interface{}) {
			idList := item.([]TopMessageId)
			mList := cbMsgF(ctx, selfUserId, idList...)
			for _, msg := range mList {
				mu.Lock()
				idHelper.PickByMessage(msg)
				dialogsData.Messages = append(dialogsData.Messages, msg)
				mu.Unlock()
			}
		})

	for _, dialogExt := range m {
		dialogsData.Dialogs = append(dialogsData.Dialogs, dialogExt.Dialog)
	}

	mr.ForEach(
		func(source chan<- interface{}) {
			idHelper.Visit(
				func(userIdList []int64) {
					if cbUserF != nil {
						source <- PeerIdList{mtproto.PEER_USER, userIdList}
					}
				},
				func(chatIdList []int64) {
					if cbChatF != nil {
						source <- PeerIdList{mtproto.PEER_CHAT, chatIdList}
					}
				},
				func(channelIdList []int64) {
					if cbChatF != nil {
						source <- PeerIdList{mtproto.PEER_CHANNEL, channelIdList}
					}
				})

		},
		func(item interface{}) {
			idList := item.(PeerIdList)
			switch idList.PeerType {
			case mtproto.PEER_USER:
				users := cbUserF(ctx, selfUserId, idList.IdList...)
				if len(users) > 0 {
					mu.Lock()
					dialogsData.Users = append(dialogsData.Users, users...)
					mu.Unlock()
				}
			case mtproto.PEER_CHAT:
				chats := cbChatF(ctx, selfUserId, idList.IdList...)
				if len(chats) > 0 {
					mu.Lock()
					dialogsData.Chats = append(dialogsData.Chats, chats...)
					mu.Unlock()
				}
			case mtproto.PEER_CHANNEL:
				chats := cbChannelF(ctx, selfUserId, idList.IdList...)
				if len(chats) > 0 {
					mu.Lock()
					dialogsData.Chats = append(dialogsData.Chats, chats...)
					mu.Unlock()
				}
			}
		})

	return dialogsData
}

func (m *DialogExt) HasDialog() bool {
	name := m.GetDialog().GetDraft().GetPredicateName()
	if name != mtproto.Predicate_draftMessage {
		return false
	}
	return true
}
