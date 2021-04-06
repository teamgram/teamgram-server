/*
 *  Copyright (c) 2018, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package channel

import (
	"math"
	"time"

	"github.com/nebula-chat/chatengine/mtproto"
)

//
// channelAdminRights#5d7ceba5 flags:#
// 	change_info:flags.0?true
// 	post_messages:flags.1?true
// 	edit_messages:flags.2?true
// 	delete_messages:flags.3?true
// 	ban_users:flags.4?true
// 	invite_users:flags.5?true
// 	invite_link:flags.6?true
// 	pin_messages:flags.7?true
// 	add_admins:flags.9?true
// 	manage_call:flags.10?true = ChannelAdminRights;
//
// type AdminRights int32

const (
	// OK is returned on success.
	CHANGE_INFO     int32 = 1 << 0
	POST_MESSAGES   int32 = 1 << 1
	EDIT_MESSAGES   int32 = 1 << 2
	DELETE_MESSAGES int32 = 1 << 3
	BAN_USERS       int32 = 1 << 4
	INVITE_USERS    int32 = 1 << 5
	INVITE_LINK     int32 = 1 << 6
	PIN_MESSAGES    int32 = 1 << 7
	ADD_ADMINS      int32 = 1 << 9
	MANAGE_CALL     int32 = 1 << 10
)

func FromChannelAdminRights(adminRights *mtproto.TLChannelAdminRights) int32 {
	var rights = int32(0)

	if adminRights.GetChangeInfo() {
		rights |= CHANGE_INFO
	}
	if adminRights.GetPostMessages() {
		rights |= POST_MESSAGES
	}
	if adminRights.GetEditMessages() {
		rights |= EDIT_MESSAGES
	}
	if adminRights.GetDeleteMessages() {
		rights |= DELETE_MESSAGES
	}
	if adminRights.GetBanUsers() {
		rights |= BAN_USERS
	}
	if adminRights.GetInviteUsers() {
		rights |= INVITE_USERS
	}
	if adminRights.GetInviteLink() {
		rights |= INVITE_LINK
	}
	if adminRights.GetPinMessages() {
		rights |= PIN_MESSAGES
	}
	if adminRights.GetAddAdmins() {
		rights |= ADD_ADMINS
	}
	if adminRights.GetManageCall() {
		rights |= MANAGE_CALL
	}

	return rights
}

func ToChannelAdminRights(rights int32) *mtproto.ChannelAdminRights {
	if rights == 0 {
		return nil
	}

	adminRights := mtproto.NewTLChannelAdminRights()

	if (rights & CHANGE_INFO) != 0 {
		adminRights.SetChangeInfo(true)
	}
	if (rights & POST_MESSAGES) != 0 {
		adminRights.SetPostMessages(true)
	}
	if (rights & EDIT_MESSAGES) != 0 {
		adminRights.SetEditMessages(true)
	}
	if (rights & DELETE_MESSAGES) != 0 {
		adminRights.SetDeleteMessages(true)
	}
	if (rights & BAN_USERS) != 0 {
		adminRights.SetBanUsers(true)
	}
	if (rights & INVITE_USERS) != 0 {
		adminRights.SetInviteUsers(true)
	}
	if (rights & INVITE_LINK) != 0 {
		adminRights.SetInviteLink(true)
	}
	if (rights & PIN_MESSAGES) != 0 {
		adminRights.SetPinMessages(true)
	}
	if (rights & ADD_ADMINS) != 0 {
		adminRights.SetAddAdmins(true)
	}
	if (rights & MANAGE_CALL) != 0 {
		adminRights.SetManageCall(true)
	}

	return adminRights.To_ChannelAdminRights()
}

// channelBannedRights#58cf4249 flags:#
// 	view_messages:flags.0?true
// 	send_messages:flags.1?true
// 	send_media:flags.2?true
// 	send_stickers:flags.3?true
// 	send_gifs:flags.4?true
// 	send_games:flags.5?true
// 	send_inline:flags.6?true
// 	embed_links:flags.7?true
// 	until_date:int = ChannelBannedRights;
//

// type BannedRights int32

const (
	// OK is returned on success.
	VIEW_MESSAGES int32 = 1 << 0
	SEND_MESSAGES int32 = 1 << 1
	SEND_MEDIA    int32 = 1 << 2
	SEND_STICKERS int32 = 1 << 3
	SEND_GIFS     int32 = 1 << 4
	SEND_GAMES    int32 = 1 << 5
	SEND_INLINE   int32 = 1 << 6
	EMBED_LINKS   int32 = 1 << 7
	// UNTIL_DATE 		BannedRights = 1 << 32
)

func FromChannelBannedRights(bannedRights *mtproto.TLChannelBannedRights) (int32, int32) {
	var (
		rights    int32 = 0
		untilDate int32 = 0
	)

	if bannedRights.GetViewMessages() {
		rights |= VIEW_MESSAGES
	}
	if bannedRights.GetSendMessages() {
		rights |= SEND_MESSAGES
	}
	if bannedRights.GetSendMedia() {
		rights |= SEND_MEDIA
	}
	if bannedRights.GetSendStickers() {
		rights |= SEND_STICKERS
	}
	if bannedRights.GetSendGifs() {
		rights |= SEND_GIFS
	}
	if bannedRights.GetSendGames() {
		rights |= SEND_GAMES
	}
	if bannedRights.GetSendInline() {
		rights |= SEND_INLINE
	}
	if bannedRights.GetEmbedLinks() {
		rights |= EMBED_LINKS
	}

	untilDate = bannedRights.GetUntilDate()

	return rights, untilDate
}

func ToChannelBannedRights(rights, untilDate int32) *mtproto.ChannelBannedRights {
	if rights == 0 {
		return nil
	}

	bannedRights := mtproto.NewTLChannelBannedRights()

	if (rights & VIEW_MESSAGES) != 0 {
		bannedRights.SetViewMessages(true)
	}
	if (rights & SEND_MESSAGES) != 0 {
		bannedRights.SetSendMessages(true)
	}
	if (rights & SEND_MEDIA) != 0 {
		bannedRights.SetSendMedia(true)
	}
	if (rights & SEND_STICKERS) != 0 {
		bannedRights.SetSendStickers(true)
	}
	if (rights & SEND_GIFS) != 0 {
		bannedRights.SetSendGifs(true)
	}
	if (rights & SEND_GAMES) != 0 {
		bannedRights.SetSendGames(true)
	}
	if (rights & SEND_INLINE) != 0 {
		bannedRights.SetSendInline(true)
	}
	if (rights & EMBED_LINKS) != 0 {
		bannedRights.SetEmbedLinks(true)
	}

	bannedRights.SetUntilDate(untilDate)

	return bannedRights.To_ChannelBannedRights()
}

func bannedRightsIsLeft(rights, untilDate int32) bool {
	_ = untilDate
	return rights&VIEW_MESSAGES == 0 && untilDate == math.MaxInt32
}

func BannedRightsIsForbidden(rights, untilDate int32) bool {
	_ = untilDate
	return rights != 0
}

func MakeUpdateChannel(channelId int32) *mtproto.Update {
	update := mtproto.NewTLUpdateChannel()
	update.SetChannelId(channelId)
	return update.To_Update()
}

func MakeChannelPeer(channelId int32) *mtproto.Peer {
	peer := &mtproto.TLPeerChannel{Data2: &mtproto.Peer_Data{
		ChannelId: channelId,
	}}
	return peer.To_Peer()
}

func MakeChannelMessageService(fromId, channelId int32, action *mtproto.MessageAction) *mtproto.Message {
	message := &mtproto.TLMessageService{Data2: &mtproto.Message_Data{
		Date:   int32(time.Now().Unix()),
		FromId: fromId,
		ToId:   MakeChannelPeer(channelId),
		Post:   true,
		Action: action,
	}}
	return message.To_Message()
}
