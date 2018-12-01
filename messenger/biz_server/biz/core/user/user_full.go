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

package user

import (
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/base"
)

func (m *UserModel) GetUserFull(selfId, userId int32) *mtproto.UserFull {
	uItem := m.makeUserItem(selfId, userId)
	user := uItem.ToUser()
	if user == nil {
		return nil
	}

	userFull := mtproto.NewTLUserFull()

	userFull.SetUser(user)

	// TODO(@benqi): blocked
	userFull.SetBlocked(false)

	// TODO(@benqi): phone_calls_available and phone_calls_private
	userFull.SetPhoneCallsAvailable(true)

	// TODO(@benqi): get privacy from user_privacy
	userFull.SetPhoneCallsPrivate(false)

	myLink, foreignLink := m.GetContactLink(selfId, userId)
	link := &mtproto.TLContactsLink{Data2: &mtproto.Contacts_Link_Data{
		User:        user,
		MyLink:      myLink,
		ForeignLink: foreignLink,
	}}
	userFull.SetLink(link.To_Contacts_Link())

	userFull.SetAbout(uItem.UsersDO.About)

	photo := m.photoCallback.GetPhoto(user.GetData2().GetPhoto().GetData2().GetPhotoId())
	userFull.SetProfilePhoto(photo)

	// notifySettings
	peer := &base.PeerUtil{
		PeerType: base.PEER_USER,
		PeerId:   userId,
	}
	userFull.SetNotifySettings(m.accountCallback.GetNotifySettings(selfId, peer))

	userFull.SetBotInfo(m.GetBotInfo(userId))
	userFull.SetCommonChatsCount(0)

	return userFull.To_UserFull()
}