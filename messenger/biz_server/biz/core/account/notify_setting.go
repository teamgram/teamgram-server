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

package account

import (
	base2 "github.com/nebula-chat/chatengine/pkg/util"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
)

func (m *AccountModel) GetNotifySettings(userId int32, peer *base.PeerUtil) *mtproto.PeerNotifySettings {
	do := m.dao.UserNotifySettingsDAO.Select(userId, int8(peer.PeerType), peer.PeerId)

	// var mute_until int32 = 0
	if do == nil {
		settings := &mtproto.TLPeerNotifySettings{Data2: &mtproto.PeerNotifySettings_Data{
			ShowPreviews: mtproto.ToBool(true),
			Silent:       mtproto.ToBool(false),
			MuteUntil:    1,
			Sound:        "default",
		}}
		return settings.To_PeerNotifySettings()
	} else {
		settings := mtproto.NewTLPeerNotifySettings()
		if do.ShowPreviews == 1 {
			settings.SetShowPreviews(mtproto.ToBool(true))
		}
		if do.Silent == 1 {
			settings.SetSilent(mtproto.ToBool(true))
		}
		if do.MuteUntil == 0 {
			settings.SetMuteUntil(0)
		} else {
			settings.SetMuteUntil(do.MuteUntil)
		}
		if do.Sound == "" {
			settings.SetSound("default")
		} else {
			settings.SetSound(do.Sound)
		}
		return settings.To_PeerNotifySettings()
	}
}

func (m *AccountModel) SetNotifySettings(userId int32, peer *base.PeerUtil, settings *mtproto.TLInputPeerNotifySettings) {
	var (
		showPreviews = base2.BoolToInt8(mtproto.FromBool(settings.GetShowPreviews()))
		silent       = base2.BoolToInt8(mtproto.FromBool(settings.GetSilent()))
	)

	do := &dataobject.UserNotifySettingsDO{
		UserId:       userId,
		PeerType:     int8(peer.PeerType),
		PeerId:       peer.PeerId,
		ShowPreviews: showPreviews,
		Silent:       silent,
		MuteUntil:    settings.GetMuteUntil(),
		Sound:        settings.GetSound(),
	}

	m.dao.UserNotifySettingsDAO.InsertOrUpdate(do)
}

func (m *AccountModel) ResetNotifySettings(userId int32) {
	m.dao.UserNotifySettingsDAO.DeleteAll(userId)
}
