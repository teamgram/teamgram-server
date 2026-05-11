// Copyright (c) 2026-present, The Teamgram Authors (https://teamgram.net).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tg

const (
	RULE_TYPE_INVALID          = 0
	ALLOW_CONTACTS             = 1
	ALLOW_ALL                  = 2
	ALLOW_USERS                = 3
	DISALLOW_CONTACTS          = 4
	DISALLOW_ALL               = 5
	DISALLOW_USERS             = 6
	ALLOW_CHAT_PARTICIPANTS    = 7
	DISALLOW_CHAT_PARTICIPANTS = 8
	ALLOW_CLOSE_FRIENDS        = 9
	ALLOW_PREMIUM              = 10
	ALLOW_BOTS                 = 11
	DISALLOW_BOTS              = 12
)

const (
	KEY_TYPE_INVALID     = 0
	STATUS_TIMESTAMP     = 1 //
	CHAT_INVITE          = 2
	PHONE_CALL           = 3
	PHONE_P2P            = 4
	FORWARDS             = 5
	PROFILE_PHOTO        = 6
	PHONE_NUMBER         = 7
	ADDED_BY_PHONE       = 8
	VOICE_MESSAGES       = 9
	ABOUT                = 10
	BIRTHDAY             = 11
	STAR_GIFTS_AUTO_SAVE = 12
	NO_PAID_MESSAGES     = 13
	SAVED_MUSIC          = 14
)

func FromInputPrivacyKeyType(k InputPrivacyKeyClazz) int {
	switch k.(type) {
	case *TLInputPrivacyKeyStatusTimestamp:
		return STATUS_TIMESTAMP
	case *TLInputPrivacyKeyChatInvite:
		return CHAT_INVITE
	case *TLInputPrivacyKeyPhoneCall:
		return PHONE_CALL
	case *TLInputPrivacyKeyPhoneP2P:
		return PHONE_P2P
	case *TLInputPrivacyKeyForwards:
		return FORWARDS
	case *TLInputPrivacyKeyProfilePhoto:
		return PROFILE_PHOTO
	case *TLInputPrivacyKeyPhoneNumber:
		return PHONE_NUMBER
	case *TLInputPrivacyKeyAddedByPhone:
		return ADDED_BY_PHONE
	case *TLInputPrivacyKeyVoiceMessages:
		return VOICE_MESSAGES
	case *TLInputPrivacyKeyAbout:
		return ABOUT
	case *TLInputPrivacyKeyBirthday:
		return BIRTHDAY
	case *TLInputPrivacyKeyStarGiftsAutoSave:
		return STAR_GIFTS_AUTO_SAVE
	case *TLInputPrivacyKeyNoPaidMessages:
		return NO_PAID_MESSAGES
	case *TLInputPrivacyKeySavedMusic:
		return SAVED_MUSIC
	default:
		return KEY_TYPE_INVALID
	}
}

func ToPrivacyKey(keyType int) (key PrivacyKeyClazz) {
	switch keyType {
	case STATUS_TIMESTAMP:
		key = PrivacyKeyStatusTimestampClazz
	case CHAT_INVITE:
		key = PrivacyKeyChatInviteClazz
	case PHONE_CALL:
		key = PrivacyKeyPhoneCallClazz
	case PHONE_P2P:
		key = PrivacyKeyPhoneP2PClazz
	case FORWARDS:
		key = PrivacyKeyForwardsClazz
	case PROFILE_PHOTO:
		key = PrivacyKeyProfilePhotoClazz
	case PHONE_NUMBER:
		key = PrivacyKeyPhoneNumberClazz
	case ADDED_BY_PHONE:
		key = PrivacyKeyAddedByPhoneClazz
	case VOICE_MESSAGES:
		key = PrivacyKeyVoiceMessagesClazz
	case ABOUT:
		key = PrivacyKeyAboutClazz
	case BIRTHDAY:
		key = PrivacyKeyBirthdayClazz
	case STAR_GIFTS_AUTO_SAVE:
		key = PrivacyKeyStarGiftsAutoSaveClazz
	case NO_PAID_MESSAGES:
		key = PrivacyKeyNoPaidMessagesClazz
	case SAVED_MUSIC:
		key = PrivacyKeySavedMusicClazz
	default:
		panic("type is invalid")
	}
	return
}

func ToPrivacyRuleByInput(userSelfId int64, inputRule InputPrivacyRuleClazz) PrivacyRuleClazz {
	switch r := inputRule.(type) {
	case *TLInputPrivacyValueAllowAll:
		return MakeTLPrivacyValueAllowAll(&TLPrivacyValueAllowAll{})
	case *TLInputPrivacyValueAllowContacts:
		return MakeTLPrivacyValueAllowContacts(&TLPrivacyValueAllowContacts{})
	case *TLInputPrivacyValueAllowUsers:
		return MakeTLPrivacyValueAllowUsers(&TLPrivacyValueAllowUsers{
			Users: ToUserIdListByInput(userSelfId, r.Users),
		})
	case *TLInputPrivacyValueDisallowAll:
		return MakeTLPrivacyValueDisallowAll(&TLPrivacyValueDisallowAll{})
	case *TLInputPrivacyValueDisallowContacts:
		return MakeTLPrivacyValueDisallowContacts(&TLPrivacyValueDisallowContacts{})
	case *TLInputPrivacyValueDisallowUsers:
		return MakeTLPrivacyValueDisallowUsers(&TLPrivacyValueDisallowUsers{
			Users: ToUserIdListByInput(userSelfId, r.Users),
		})
	case *TLInputPrivacyValueAllowChatParticipants:
		return MakeTLPrivacyValueAllowChatParticipants(&TLPrivacyValueAllowChatParticipants{
			Chats: r.Chats,
		})
	case *TLInputPrivacyValueDisallowChatParticipants:
		return MakeTLPrivacyValueDisallowChatParticipants(&TLPrivacyValueDisallowChatParticipants{
			Chats: r.Chats,
		})
	case *TLInputPrivacyValueAllowCloseFriends:
		return MakeTLPrivacyValueAllowCloseFriends(&TLPrivacyValueAllowCloseFriends{})
	case *TLInputPrivacyValueAllowPremium:
		return MakeTLPrivacyValueAllowPremium(&TLPrivacyValueAllowPremium{})
	case *TLInputPrivacyValueAllowBots:
		return MakeTLPrivacyValueAllowBots(&TLPrivacyValueAllowBots{})
	case *TLInputPrivacyValueDisallowBots:
		return MakeTLPrivacyValueDisallowBots(&TLPrivacyValueDisallowBots{})
	default:
		panic("type is invalid")
	}
}

func ToPrivacyRuleListByInput(userSelfId int64, inputRules []InputPrivacyRuleClazz) (rules []PrivacyRuleClazz) {
	rules = make([]PrivacyRuleClazz, 0, len(inputRules))
	for _, inputRule := range inputRules {
		rules = append(rules, ToPrivacyRuleByInput(userSelfId, inputRule))
	}
	return
}

/*
// PickAllIdListByRules
// TODO(@benqi): pick chat and channel
func PickAllIdListByRules(rules []*PrivacyRule) (userIdList, chatIdList, channelIdList []int64) {
	userIdList = make([]int64, 0)
	chatIdList = make([]int64, 0)
	channelIdList = make([]int64, 0)

	for _, r := range rules {
		if r == nil {
			continue
		}

		switch c := r.Clazz.(type) {
		case *TLPrivacyValueAllowUsers:
			if len(c.Users) > 0 {
				userIdList = append(userIdList, c.Users...)
			}
		case *TLPrivacyValueDisallowUsers:
			if len(c.Users) > 0 {
				userIdList = append(userIdList, c.Users...)
			}
		case *TLPrivacyValueAllowChatParticipants:
			for _, id := range c.Chats {
				if id >= MinNebulaChatChannelID {
					channelIdList = append(channelIdList, id)
				} else {
					chatIdList = append(chatIdList, id)
				}
			}
		case *TLPrivacyValueDisallowChatParticipants:
			for _, id := range c.Chats {
				if id >= MinNebulaChatChannelID {
					channelIdList = append(channelIdList, id)
				} else {
					chatIdList = append(chatIdList, id)
				}
			}
		}
	}

	return
}

func CheckPrivacyIsAllow(selfId int64,
	rules []*PrivacyRule,
	checkId int64,
	cbContact func(id, checkId int64) bool,
	cbChat func(checkId int64, idList []int64) bool) bool {
	ruleType := RULE_TYPE_INVALID

	for _, r := range rules {
		switch r.PrivacyRuleClazzName() {
		case ClazzName_privacyValueAllowAll:
			ruleType = ALLOW_ALL
		case ClazzName_privacyValueAllowContacts:
			ruleType = ALLOW_CONTACTS
		case ClazzName_privacyValueDisallowAll:
			ruleType = DISALLOW_ALL
		}
	}

	switch ruleType {
	case ALLOW_ALL:
		for _, r := range rules {
			switch r.PrivacyRuleClazzName() {
			case ClazzName_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return false
					}
				}
			case ClazzName_privacyValueDisallowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return false
				}
			}
		}
		return true
	case ALLOW_CONTACTS:
		for _, r := range rules {
			switch r.PrivacyRuleClazzName() {
			case ClazzName_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return true
					}
				}
			case ClazzName_privacyValueAllowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return true
				}
			case ClazzName_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return false
					}
				}
			case ClazzName_privacyValueDisallowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return false
				}
			}
		}
		return cbContact(selfId, checkId)
	case DISALLOW_ALL:
		for _, r := range rules {
			switch r.PrivacyRuleClazzName() {
			case ClazzName_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return true
					}
				}
			case ClazzName_privacyValueAllowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return true
				}
			}
		}
		return false
	}

	return false
}

//// *
privacyValueAllowContacts#fffe1bac = PrivacyRule;
privacyValueAllowAll#65427b82 = PrivacyRule;
privacyValueAllowUsers#4d5bbe0c users:Vector<int> = PrivacyRule;
privacyValueDisallowContacts#f888fa1a = PrivacyRule;
privacyValueDisallowAll#8b73e763 = PrivacyRule;
privacyValueDisallowUsers#c7f49b7 users:Vector<int> = PrivacyRule;
privacyValueAllowChatParticipants#18be796b chats:Vector<int> = PrivacyRule;
privacyValueDisallowChatParticipants#acae0690 chats:Vector<int> = PrivacyRule;
/// * /
func privacyIsAllow(rules []*PrivacyRule, userId int64, isContact bool) bool {
	ruleType := RULE_TYPE_INVALID

	for _, r := range rules {
		switch r.PrivacyRuleClazzName() {
		case ClazzName_privacyValueAllowAll:
			ruleType = ALLOW_ALL
		case ClazzName_privacyValueAllowContacts:
			ruleType = ALLOW_CONTACTS
		case ClazzName_privacyValueDisallowAll:
			ruleType = DISALLOW_ALL
		}
	}

	switch ruleType {
	case ALLOW_ALL:
		for _, r := range rules {
			switch r.PrivacyRuleClazzName() {
			case ClazzName_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == userId {
						return false
					}
				}
			case ClazzName_privacyValueDisallowChatParticipants:
				return true
			}
		}
		return true
	case ALLOW_CONTACTS:
		for _, r := range rules {
			switch r.PrivacyRuleClazzName() {
			case ClazzName_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == userId {
						return true
					}
				}
			case ClazzName_privacyValueAllowChatParticipants:
				return true
			case ClazzName_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == userId {
						return false
					}
				}
			case ClazzName_privacyValueDisallowChatParticipants:
				return true
			}
		}
		return isContact
	case DISALLOW_ALL:
		for _, r := range rules {
			switch r.PrivacyRuleClazzName() {
			case ClazzName_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == userId {
						return true
					}
				}
			case ClazzName_privacyValueAllowChatParticipants:
				return true
			}
		}
		return false
	}

	return false
}
*/
