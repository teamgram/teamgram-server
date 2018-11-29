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

package dialog

import (
	"encoding/json"
	"github.com/nebula-chat/chatengine/mtproto"
)

func (m *DialogModel) SaveDraftMessage(userId int32, peerType int32, peerId int32, message *mtproto.DraftMessage) {
	draft, _ := json.Marshal(message)
	m.dao.UserDialogsDAO.SaveDraft(string(draft), userId, int8(peerType), peerId)
}

func (m *DialogModel) ClearDraftMessage(userId int32, peerType int32, peerId int32) bool {
	// draft, _ := json.Marshal(message)
	// m.dao.UserDialogsDAO.SaveDraft(string(draft), userId, int8(peerType), peerId)
	affectedRows := m.dao.UserDialogsDAO.ClearDraft(userId, int8(peerType), peerId)
	return affectedRows > 0
}
