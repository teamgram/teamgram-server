// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
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

// Author: Benqi (wubenqi@gmail.com)

package message

import (
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
)

func (m *MessageModel) GetUnreadMentionCount(userId int32, peerType int8, peerId int32) int32 {
	params := map[string]interface{}{
		"user_id":   userId,
		"peer_type": peerType,
		"peer_id":   peerId,
		"deleted":   0,
	}
	return int32(m.dao.CommonDAO.CalcSize("mentions", params))
}

func (m *MessageModel) InsertUnreadMention(userId int32 , peerType int8, peerId, mentionedMessageId int32) {
	do := &dataobject.UnreadMentionsDO{
		UserId:             userId,
		PeerType:           peerType,
		PeerId:             peerId,
		MentionedMessageId: mentionedMessageId,
	}
	m.dao.UnreadMentionsDAO.InsertIgnore(do)
}

func (m *MessageModel) UpdateUnreadReadMention(userId int32, peerType int8, peerId, mentionedMessageId int32) {
	m.dao.UnreadMentionsDAO.Delete(userId, peerType, peerId, mentionedMessageId)
	m.dao.MessageBoxesDAO.UpdateMediaUnread(userId, mentionedMessageId)
	m.dao.UserDialogsDAO.UpdateUnreadMentionCount(userId, peerType, peerId)
}
