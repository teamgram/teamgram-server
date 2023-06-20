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
	"container/list"

	"github.com/teamgram/proto/mtproto"
)

func equalPeer(p, o *mtproto.Peer) bool {
	if p.GetPredicateName() != o.GetPredicateName() {
		return false
	}
	switch p.GetPredicateName() {
	case mtproto.Predicate_peerUser:
		return p.GetUserId() == o.GetUserId()
	case mtproto.Predicate_peerChat:
		return p.GetChatId() == o.GetChatId()
	case mtproto.Predicate_peerChannel:
		return p.GetChannelId() == o.GetChannelId()
	default:
		return false
	}
}

type mergeUpdatesHelper struct {
	newMessages           []*mtproto.Message
	deleteMessages        *list.Element
	readHistoryInBoxList  []*list.Element
	readHistoryOutBoxList []*list.Element
	readMessagesContents  *list.Element
	folderPeers           *list.Element
	pinnedMessages        *list.Element
	phoneCallRequest      *list.Element
	otherUpdates          *list.List
}

func newMergeUpdatesHelper() *mergeUpdatesHelper {
	return &mergeUpdatesHelper{
		newMessages:  []*mtproto.Message{},
		otherUpdates: list.New(),
	}
}

func (m *mergeUpdatesHelper) merge(update *mtproto.Update, pts int32) {
	switch update.PredicateName {
	case mtproto.Predicate_updateNewMessage:
		// updateNewMessage#1f2b0afd message:Message pts:int pts_count:int = Update;
		m.newMessages = append(m.newMessages, update.Message_MESSAGE)
	case mtproto.Predicate_updateDeleteMessages:
		// updateDeleteMessages#a20db0e5 messages:Vector<int> pts:int pts_count:int = Update;
		if m.deleteMessages == nil {
			update.Pts_INT32 = pts
			update.PtsCount = 0
			m.deleteMessages = m.otherUpdates.PushBack(update)
		} else {
			tmp := m.deleteMessages.Value.(*mtproto.Update)
			m.otherUpdates.Remove(m.deleteMessages)
			tmp.Messages = append(tmp.Messages, update.Messages...)
			m.deleteMessages = m.otherUpdates.PushBack(tmp)
		}
	case mtproto.Predicate_updateReadHistoryInbox:
		// updateReadHistoryInbox#9c974fdf flags:# folder_id:flags.0?int peer:Peer max_id:int still_unread_count:int pts:int pts_count:int = Update;
		var (
			foundId = -1
			inbox   *mtproto.Update
		)
		for n, e := range m.readHistoryInBoxList {
			inbox = e.Value.(*mtproto.Update)
			if inbox.GetFolderId().GetValue() == update.GetFolderId().GetValue() &&
				equalPeer(inbox.GetPeer_PEER(), update.GetPeer_PEER()) {
				if update.GetMaxId() > inbox.GetMaxId() {
					inbox.MaxId = update.GetMaxId()
				}
				inbox.StillUnreadCount = update.StillUnreadCount

				m.otherUpdates.Remove(e)
				foundId = n
				break
			}
		}
		if foundId > 0 {
			m.readHistoryInBoxList[foundId] = m.otherUpdates.PushBack(inbox)
		} else {
			update.Pts_INT32 = pts
			update.PtsCount = 0
			e := m.otherUpdates.PushBack(update)
			m.readHistoryInBoxList = append(m.readHistoryInBoxList, e)
		}
	case mtproto.Predicate_updateReadHistoryOutbox:
		// updateReadHistoryOutbox#2f2f21bf peer:Peer max_id:int pts:int pts_count:int = Update;
		var (
			foundId = -1
			outbox  *mtproto.Update
		)
		for n, e := range m.readHistoryOutBoxList {
			outbox = e.Value.(*mtproto.Update)
			if equalPeer(outbox.GetPeer_PEER(), update.GetPeer_PEER()) {
				if update.GetMaxId() > outbox.GetMaxId() {
					outbox.MaxId = update.GetMaxId()
				}
				m.otherUpdates.Remove(e)
				foundId = n
				break
			}
		}
		if foundId > 0 {
			// found
			m.readHistoryOutBoxList[foundId] = m.otherUpdates.PushBack(outbox)
		} else {
			update.Pts_INT32 = pts
			update.PtsCount = 0
			e := m.otherUpdates.PushBack(update)
			m.readHistoryOutBoxList = append(m.readHistoryOutBoxList, e)
		}
	case mtproto.Predicate_updateWebPage:
		// updateWebPage#7f891213 webpage:WebPage pts:int pts_count:int = Update;
		update.Pts_INT32 = pts
		update.PtsCount = 0
		m.otherUpdates.PushBack(update)
	case mtproto.Predicate_updateReadMessagesContents:
		// updateReadMessagesContents#68c13933 messages:Vector<int> pts:int pts_count:int = Update;
		if m.readMessagesContents == nil {
			update.Pts_INT32 = pts
			update.PtsCount = 0
			m.readMessagesContents = m.otherUpdates.PushBack(update)
		} else {
			tmp := m.readMessagesContents.Value.(*mtproto.Update)
			m.otherUpdates.Remove(m.readMessagesContents)
			tmp.Messages = append(tmp.Messages, update.Messages...)
			m.readMessagesContents = m.otherUpdates.PushBack(tmp)
		}
	case mtproto.Predicate_updateNewChannelMessage:
		// updateNewChannelMessage#62ba04d9 message:Message pts:int pts_count:int = Update;
		// ignore
	case mtproto.Predicate_updateDeleteChannelMessages:
		// updateDeleteChannelMessages#c37521c9 channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
		// ignore
	case mtproto.Predicate_updateEditChannelMessage:
		// updateEditChannelMessage#1b3f4df7 message:Message pts:int pts_count:int = Update;
		// ignore
	case mtproto.Predicate_updateEditMessage:
		// updateEditMessage#e40370a3 message:Message pts:int pts_count:int = Update;
		update.Pts_INT32 = pts
		update.PtsCount = 0
		m.otherUpdates.PushBack(update)
	case mtproto.Predicate_updateChannelWebPage:
		// updateChannelWebPage#40771900 channel_id:int webpage:WebPage pts:int pts_count:int = Update;
		// ignore
	case mtproto.Predicate_updateFolderPeers:
		// updateFolderPeers#19360dc0 folder_peers:Vector<FolderPeer> pts:int pts_count:int = Update;
		if m.folderPeers != nil {
			m.otherUpdates.Remove(m.folderPeers)
		}
		update.Pts_INT32 = pts
		update.PtsCount = 0
		m.folderPeers = m.otherUpdates.PushBack(update)
	case mtproto.Predicate_updatePinnedMessages:
		// updatePinnedMessages#ed85eab5 flags:# pinned:flags.0?true peer:Peer messages:Vector<int> pts:int pts_count:int = Update;
		if m.pinnedMessages != nil {
			m.otherUpdates.Remove(m.pinnedMessages)
		}
		update.Pts_INT32 = pts
		update.PtsCount = 0
		m.pinnedMessages = m.otherUpdates.PushBack(update)
	case mtproto.Predicate_updatePinnedChannelMessages:
		// updatePinnedChannelMessages#8588878b flags:# pinned:flags.0?true channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
		// ignore
	case mtproto.Predicate_updatePhoneCall:
		switch update.GetPhoneCall().GetPredicateName() {
		case mtproto.Predicate_phoneCallRequested:
			if m.phoneCallRequest != nil {
				m.otherUpdates.Remove(m.phoneCallRequest)
			}
			m.phoneCallRequest = m.otherUpdates.PushBack(update)
		case mtproto.Predicate_phoneCallDiscarded:
			if m.phoneCallRequest != nil {
				m.otherUpdates.Remove(m.phoneCallRequest)
				m.phoneCallRequest = nil
			}
		}
	default:
		// TODO: merge
		m.otherUpdates.PushBack(update)
	}
}

func (m *mergeUpdatesHelper) toUpdates() []*mtproto.Update {
	updates := make([]*mtproto.Update, 0, m.otherUpdates.Len())
	for e := m.otherUpdates.Front(); e != nil; e = e.Next() {
		updates = append(updates, e.Value.(*mtproto.Update))
	}
	return updates
}
