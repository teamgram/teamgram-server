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
	"github.com/nebula-chat/chatengine/pkg/util"
)

type dialogLogic struct {
	selfUserId int32
	peerType   int32
	peerId     int32
	dao        *dialogsDAO
}

func (m *DialogModel) MakeDialogLogic(userId, peerType, peerId int32) *dialogLogic {
	return &dialogLogic{
		selfUserId: userId,
		peerType:   peerType,
		peerId:     peerId,
		dao:        m.dao,
	}
}

func (d *dialogLogic) ToggleDialogPin(pinned bool) error {
	d.dao.UserDialogsDAO.UpdatePinned(util.BoolToInt8(pinned), d.selfUserId, int8(d.peerType), d.peerId)
	return nil
}
