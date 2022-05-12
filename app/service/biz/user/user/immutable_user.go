// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package user

import (
	"sort"

	"github.com/teamgram/proto/mtproto"

	"github.com/gogo/protobuf/types"
)

//func (m *ImmutableUser) ResetSetContacts(contacts []*ContactData) {
//	m.Contacts = contacts
//}
//
//func (m *ImmutableUser) ResetSetPrivacyRules(rules []*PrivacyKeyRules) {
//	m.KeysPrivacyRules = rules
//}

func (m *ImmutableUser) Id() int64 {
	return m.User.Id
}

func (m *ImmutableUser) AccessHash() int64 {
	return m.User.AccessHash
}

func (m *ImmutableUser) FirstName() string {
	return m.User.FirstName
}

func (m *ImmutableUser) LastName() string {
	return m.User.LastName
}

func (m *ImmutableUser) SetFirstName(v string) {
	m.User.FirstName = v
}

func (m *ImmutableUser) SetLastName(v string) {
	m.User.LastName = v
}

func (m *ImmutableUser) Username() string {
	return m.User.Username
}

func (m *ImmutableUser) SetUsername(v string) {
	m.User.Username = v
}

func (m *ImmutableUser) Phone() string {
	return m.User.Phone
}

func (m *ImmutableUser) Deleted() bool {
	return m.User.Deleted
}

func (m *ImmutableUser) Restricted() bool {
	return m.User.Restricted
}

func (m *ImmutableUser) ProfilePhoto() *mtproto.UserProfilePhoto {
	return mtproto.MakeUserProfilePhotoByPhoto(m.User.ProfilePhoto)
}

func (m *ImmutableUser) Photo() *mtproto.Photo {
	return m.User.ProfilePhoto
}

func (m *ImmutableUser) About() string {
	return m.User.About.GetValue()
}

func (m *ImmutableUser) SetAbout(v string) {
	m.User.About = mtproto.MakeFlagsString(v)
}

func (m *ImmutableUser) IsBot() bool {
	return m.User.Bot != nil
}

func (m *ImmutableUser) BotChatHistory() bool {
	return m.User.Bot.GetBotChatHistory()
}

func (m *ImmutableUser) BotNochats() bool {
	return m.User.Bot.GetBotNochats()
}

func (m *ImmutableUser) BotInlineGeo() bool {
	return m.User.Bot.GetBotInlineGeo()
}

func (m *ImmutableUser) RestrictionReason() []*mtproto.RestrictionReason {
	return m.User.RestrictionReason
}

func (m *ImmutableUser) BotInlinePlaceholder() *types.StringValue {
	return m.User.Bot.GetBotInlinePlaceholder()
}

func (m *ImmutableUser) Verified() bool {
	return m.User.Verified
}

func (m *ImmutableUser) SetVerified(v bool) {
	m.User.Verified = v
}

func (m *ImmutableUser) Support() bool {
	return m.User.Support
}

func (m *ImmutableUser) Scam() bool {
	return m.User.Scam
}

func (m *ImmutableUser) Fake() bool {
	return m.User.Fake
}

func (m *ImmutableUser) BotInfoVersion() int32 {
	return m.User.Bot.GetBotInfoVersion()
}

func (m *ImmutableUser) CheckContact(cId int64) (bool, bool) {
	i := sort.Search(len(m.Contacts), func(i int) bool {
		return m.Contacts[i].ContactUserId >= cId
	})
	if i < len(m.Contacts) && m.Contacts[i].ContactUserId == cId {
		return true, m.Contacts[i].MutualContact
	} else {
		return false, false
	}
}

func (m *ImmutableUser) GetContactData(cId int64) *ContactData {
	// logx.Info("GetContactData: %d ==> %s", cId, m.DebugString())
	i2 := sort.Search(len(m.Contacts), func(i int) bool {
		// logx.Info("GetContactData: %d ==> %v", cId, m.Contacts[i].ContactUserId)
		return m.Contacts[i].ContactUserId >= cId
	})
	if i2 < len(m.Contacts) && m.Contacts[i2].ContactUserId == cId {
		return m.Contacts[i2]
	} else {
		return nil
	}
}

func (m *ImmutableUser) CheckPrivacy(keyType int, id int64) bool {
	var (
		rules *PrivacyKeyRules
	)

	for _, v := range m.KeysPrivacyRules {
		if v.Key == int32(keyType) {
			rules = v
			break
		}
	}

	if rules == nil {
		return true
	}

	isContact, _ := m.CheckContact(id)
	allow := privacyIsAllow(rules.Rules, id, isContact)

	// logx.Infof("CheckPrivacy(%d, %s, %d): %v", m.Id(), rules.DebugString(), id, allow)
	return allow
}

func (m *ImmutableUser) ToUnsafeUser(selfUser *ImmutableUser) *mtproto.User {
	if m.Deleted() {
		return m.ToDeletedUser()
	}

	if m.Id() == selfUser.Id() {
		return m.ToSelfUser()
	}

	user := mtproto.MakeTLUser(&mtproto.User{
		Self:                 false,
		Contact:              false,
		MutualContact:        false,
		Deleted:              false,
		Bot:                  m.IsBot(),
		BotChatHistory:       m.BotChatHistory(),
		BotNochats:           m.BotNochats(),
		Verified:             m.Verified(),
		Restricted:           m.Restricted(),
		Min:                  false,
		BotInlineGeo:         m.BotInlineGeo(),
		Support:              m.Support(),
		Scam:                 m.Scam(),
		ApplyMinPhoto:        false,
		Fake:                 m.Fake(),
		Id:                   m.Id(),
		AccessHash:           mtproto.MakeFlagsInt64(m.AccessHash()),
		FirstName:            mtproto.MakeFlagsString(m.FirstName()),
		LastName:             mtproto.MakeFlagsString(m.LastName()),
		Username:             mtproto.MakeFlagsString(m.Username()),
		Phone:                nil,
		Photo:                m.ProfilePhoto(),
		Status:               MakeUserStatus(m.LastSeenAt, true),
		BotInfoVersion:       mtproto.MakeFlagsInt32(m.BotInfoVersion()),
		RestrictionReason:    m.RestrictionReason(),
		BotInlinePlaceholder: m.BotInlinePlaceholder(),
		LangCode:             nil,
	}).To_User()

	contact := selfUser.GetContactData(m.Id())
	if contact != nil {
		user.Contact = true
		user.MutualContact = contact.MutualContact
		user.FirstName = contact.FirstName
		user.LastName = contact.LastName
	}

	// phone
	if m.CheckPrivacy(PHONE_NUMBER, selfUser.Id()) {
		user.Phone = mtproto.MakeFlagsString(m.Phone())
	}

	// photo
	if m.CheckPrivacy(PROFILE_PHOTO, selfUser.Id()) {
		user.Photo = m.ProfilePhoto()
	}

	// status
	allowTimestamp := m.CheckPrivacy(STATUS_TIMESTAMP, selfUser.Id())
	user.Status = MakeUserStatus(m.LastSeenAt, allowTimestamp)

	return user
}

func (m *ImmutableUser) ToSelfUser() *mtproto.User {
	return mtproto.MakeTLUser(&mtproto.User{
		Self:                 true,
		Contact:              true,
		MutualContact:        true,
		Deleted:              false,
		Bot:                  m.IsBot(),
		BotChatHistory:       m.BotChatHistory(),
		BotNochats:           m.BotNochats(),
		Verified:             m.Verified(),
		Restricted:           m.Restricted(),
		Min:                  false,
		BotInlineGeo:         m.BotInlineGeo(),
		Support:              m.Support(),
		Scam:                 m.Scam(),
		ApplyMinPhoto:        false,
		Fake:                 m.Fake(),
		Id:                   m.Id(),
		AccessHash:           mtproto.MakeFlagsInt64(m.AccessHash()),
		FirstName:            mtproto.MakeFlagsString(m.FirstName()),
		LastName:             mtproto.MakeFlagsString(m.LastName()),
		Username:             mtproto.MakeFlagsString(m.Username()),
		Phone:                mtproto.MakeFlagsString(m.Phone()),
		Photo:                m.ProfilePhoto(),
		Status:               MakeUserStatus(m.LastSeenAt, true),
		BotInfoVersion:       mtproto.MakeFlagsInt32(m.BotInfoVersion()),
		RestrictionReason:    m.RestrictionReason(),
		BotInlinePlaceholder: m.BotInlinePlaceholder(),
		LangCode:             nil,
	}).To_User()
}

func (m *ImmutableUser) ToDeletedUser() *mtproto.User {
	return mtproto.MakeTLUser(&mtproto.User{
		Id:                   m.Id(),
		Self:                 false,
		Contact:              false,
		MutualContact:        false,
		Deleted:              true,
		Bot:                  false,
		BotChatHistory:       false,
		BotNochats:           false,
		Verified:             false,
		Restricted:           false,
		Min:                  false,
		BotInlineGeo:         false,
		Support:              false,
		Scam:                 false,
		ApplyMinPhoto:        false,
		Fake:                 false,
		AccessHash:           mtproto.MakeFlagsInt64(m.AccessHash()),
		FirstName:            nil,
		LastName:             nil,
		Username:             nil,
		Phone:                nil,
		Photo:                nil,
		Status:               nil,
		BotInfoVersion:       nil,
		RestrictionReason:    nil,
		BotInlinePlaceholder: nil,
		LangCode:             nil,
	}).To_User()
}
