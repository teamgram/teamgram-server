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
	"time"

	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
)

type channelParticipantData struct {
	*dataobject.ChannelParticipantsDO
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *channelParticipantData) CanChangeInfo() bool {
	return m.AdminRights&CHANGE_INFO != 0
}

func (m *channelParticipantData) CanPostMessages() bool {
	return m.AdminRights&POST_MESSAGES != 0
}

func (m *channelParticipantData) CanEditMessages() bool {
	return m.AdminRights&EDIT_MESSAGES != 0
}

func (m *channelParticipantData) CanDeleteMessages() bool {
	return m.AdminRights&DELETE_MESSAGES != 0
}

func (m *channelParticipantData) CanBanUsers() bool {
	return m.AdminRights&BAN_USERS != 0
}

func (m *channelParticipantData) CanInviteUsers() bool {
	return m.AdminRights&INVITE_USERS != 0
}

func (m *channelParticipantData) CanInviteLink() bool {
	return m.AdminRights&INVITE_LINK != 0
}

func (m *channelParticipantData) CanPinMessages() bool {
	return m.AdminRights&PIN_MESSAGES != 0
}

func (m *channelParticipantData) CanAddAdmins() bool {
	return m.AdminRights&ADD_ADMINS != 0
}

func (m *channelParticipantData) CanManageCall() bool {
	return m.AdminRights&MANAGE_CALL != 0
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *channelParticipantData) CanViewMessages() bool {
	return m.BannedRights&VIEW_MESSAGES != 0 && int32(time.Now().Unix()) <= m.BannedUntilDate
}

func (m *channelParticipantData) CanSendMessages() bool {
	return m.BannedRights&SEND_MESSAGES != 0 && int32(time.Now().Unix()) <= m.BannedUntilDate
}

func (m *channelParticipantData) CanSendMedia() bool {
	return m.BannedRights&SEND_MEDIA != 0 && int32(time.Now().Unix()) <= m.BannedUntilDate
}

func (m *channelParticipantData) CanSendStickers() bool {
	return m.BannedRights&SEND_STICKERS != 0 && int32(time.Now().Unix()) <= m.BannedUntilDate
}

func (m *channelParticipantData) CanSendGifs() bool {
	return m.BannedRights&SEND_GIFS != 0 && int32(time.Now().Unix()) <= m.BannedUntilDate
}

func (m *channelParticipantData) CanSendGames() bool {
	return m.BannedRights&SEND_GAMES != 0 && int32(time.Now().Unix()) <= m.BannedUntilDate
}

func (m *channelParticipantData) CanSendInline() bool {
	return m.BannedRights&SEND_INLINE != 0 && int32(time.Now().Unix()) <= m.BannedUntilDate
}

func (m *channelParticipantData) CanEmbedLinks() bool {
	return m.BannedRights&EMBED_LINKS != 0 && int32(time.Now().Unix()) <= m.BannedUntilDate
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *channelParticipantData) IsCreator() bool {
	return m.ChannelParticipantsDO.IsCreator == 1
}

func (m *channelParticipantData) IsAdmin() bool {
	return m.IsCreator() || m.AdminRights != 0
}

func (m *channelParticipantData) IsBanned() bool {
	// TODO(@benqi): check banned_until_date
	return m.IsKicked() && m.BannedRights == 255 && int32(time.Now().Unix()) <= m.BannedUntilDate
}

func (m *channelParticipantData) IsRestricted() bool {
	return m.IsBanned() && int32(time.Now().Unix()) <= m.BannedUntilDate
}

func (m *channelParticipantData) IsKicked() bool {
	return m.ChannelParticipantsDO.IsKicked != 0
}

func (m *channelParticipantData) IsLeft() bool {
	return m.ChannelParticipantsDO.IsLeft != 0
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *channelParticipantData) ToChannelAdminRights() *mtproto.ChannelAdminRights {
	return ToChannelAdminRights(m.AdminRights)
}

func (m *channelParticipantData) ToChannelBannedRights() *mtproto.ChannelBannedRights {
	return ToChannelBannedRights(m.BannedRights, m.BannedUntilDate)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *channelParticipantData) ToChannelParticipant() (channelParticipant *mtproto.ChannelParticipant) {
	if m.IsLeft() {
		return
	}

	if m.IsCreator() {
		participant := &mtproto.TLChannelParticipantCreator{Data2: &mtproto.ChannelParticipant_Data{
			UserId: m.UserId,
		}}
		channelParticipant = participant.To_ChannelParticipant()
		return
	}

	if m.IsAdmin() {
		participant := &mtproto.TLChannelParticipantAdmin{Data2: &mtproto.ChannelParticipant_Data{
			UserId:      m.UserId,
			CanEdit:     true,
			InviterId:   m.InviterUserId,
			Date:        m.JoinedAt,
			PromotedBy:  m.PromotedBy,
			AdminRights: ToChannelAdminRights(m.AdminRights),
		}}
		channelParticipant = participant.To_ChannelParticipant()
		return
	}

	if m.IsBanned() {
		participant := &mtproto.TLChannelParticipantBanned{Data2: &mtproto.ChannelParticipant_Data{
			UserId:       m.UserId,
			Left:         bannedRightsIsLeft(m.BannedRights, m.BannedUntilDate),
			KickedBy:     m.KickedBy,
			Date:         m.BannedAt,
			BannedRights: ToChannelBannedRights(m.BannedRights, m.BannedUntilDate),
		}}
		channelParticipant = participant.To_ChannelParticipant()
		return
	}

	participant := &mtproto.TLChannelParticipant{Data2: &mtproto.ChannelParticipant_Data{
		UserId: m.UserId,
		Date:   m.JoinedAt,
	}}
	channelParticipant = participant.To_ChannelParticipant()
	return
}

func (m *channelParticipantData) TryToChannelParticipantSelf(selfUserId int32) (channelParticipant *mtproto.ChannelParticipant) {
	if m.IsLeft() {
		return
	}

	if m.IsCreator() {
		participant := &mtproto.TLChannelParticipantCreator{Data2: &mtproto.ChannelParticipant_Data{
			UserId: m.UserId,
		}}
		channelParticipant = participant.To_ChannelParticipant()
		return
	}

	if m.IsAdmin() {
		participant := &mtproto.TLChannelParticipantAdmin{Data2: &mtproto.ChannelParticipant_Data{
			UserId:      m.UserId,
			CanEdit:     true,
			InviterId:   m.InviterUserId,
			Date:        m.JoinedAt,
			PromotedBy:  m.PromotedBy,
			AdminRights: ToChannelAdminRights(m.AdminRights),
		}}
		channelParticipant = participant.To_ChannelParticipant()
		return
	}

	if m.IsBanned() {
		participant := &mtproto.TLChannelParticipantBanned{Data2: &mtproto.ChannelParticipant_Data{
			UserId:       m.UserId,
			Left:         bannedRightsIsLeft(m.BannedRights, m.BannedUntilDate),
			KickedBy:     m.KickedBy,
			Date:         m.BannedAt,
			BannedRights: ToChannelBannedRights(m.BannedRights, m.BannedUntilDate),
		}}
		channelParticipant = participant.To_ChannelParticipant()
		return
	}

	if m.UserId == selfUserId {
		participant := &mtproto.TLChannelParticipantSelf{Data2: &mtproto.ChannelParticipant_Data{
			UserId:    m.UserId,
			InviterId: m.InviterUserId,
			Date:      m.JoinedAt,
		}}
		channelParticipant = participant.To_ChannelParticipant()
	} else {
		participant := &mtproto.TLChannelParticipant{Data2: &mtproto.ChannelParticipant_Data{
			UserId: m.UserId,
			Date:   m.JoinedAt,
		}}
		channelParticipant = participant.To_ChannelParticipant()
	}

	return
}
