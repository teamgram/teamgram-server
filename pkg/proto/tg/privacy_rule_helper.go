// Copyright 2024 Teamgram Authors
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

package tg

/*************************************************************
	inputPrivacyValueAllowContacts#d09e07b = InputPrivacyRule;
	inputPrivacyValueAllowAll#184b35ce = InputPrivacyRule;
	inputPrivacyValueAllowUsers#131cc67f users:Vector<InputUser> = InputPrivacyRule;
	inputPrivacyValueDisallowContacts#ba52007 = InputPrivacyRule;
	inputPrivacyValueDisallowAll#d66b66c9 = InputPrivacyRule;
	inputPrivacyValueDisallowUsers#90110467 users:Vector<InputUser> = InputPrivacyRule;
	inputPrivacyValueAllowChatParticipants#840649cf chats:Vector<long> = InputPrivacyRule;
	inputPrivacyValueDisallowChatParticipants#e94f0f86 chats:Vector<long> = InputPrivacyRule;
	inputPrivacyValueAllowCloseFriends#2f453e49 = InputPrivacyRule;
	inputPrivacyValueAllowPremium#77cdc9f1 = InputPrivacyRule;
	inputPrivacyValueAllowBots#5a4fcce5 = InputPrivacyRule;
	inputPrivacyValueDisallowBots#c4e57915 = InputPrivacyRule;

	privacyValueAllowContacts#fffe1bac = PrivacyRule;
	privacyValueAllowAll#65427b82 = PrivacyRule;
	privacyValueAllowUsers#b8905fb2 users:Vector<long> = PrivacyRule;
	privacyValueDisallowContacts#f888fa1a = PrivacyRule;
	privacyValueDisallowAll#8b73e763 = PrivacyRule;
	privacyValueDisallowUsers#e4621141 users:Vector<long> = PrivacyRule;
	privacyValueAllowChatParticipants#6b134e8e chats:Vector<long> = PrivacyRule;
	privacyValueDisallowChatParticipants#41c87565 chats:Vector<long> = PrivacyRule;
	privacyValueAllowCloseFriends#f7e8d89b = PrivacyRule;
	privacyValueAllowPremium#ece9814b = PrivacyRule;
	privacyValueAllowBots#21461b5d = PrivacyRule;
	privacyValueDisallowBots#f6a5f82f = PrivacyRule;
**/

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
	ALOW_BOTS                  = 11
	DISALLOW_BOTS              = 12
)

/*************************************************************
	inputPrivacyKeyStatusTimestamp#4f96cb18 = InputPrivacyKey;
	inputPrivacyKeyChatInvite#bdfb0426 = InputPrivacyKey;
	inputPrivacyKeyPhoneCall#fabadc5f = InputPrivacyKey;
	inputPrivacyKeyPhoneP2P#db9e70d2 = InputPrivacyKey;
	inputPrivacyKeyForwards#a4dd4c08 = InputPrivacyKey;
	inputPrivacyKeyProfilePhoto#5719bacc = InputPrivacyKey;
	inputPrivacyKeyPhoneNumber#352dafa = InputPrivacyKey;
	inputPrivacyKeyAddedByPhone#d1219bdd = InputPrivacyKey;
	inputPrivacyKeyVoiceMessages#aee69d68 = InputPrivacyKey;
	inputPrivacyKeyAbout#3823cc40 = InputPrivacyKey;
	inputPrivacyKeyBirthday#d65a11cc = InputPrivacyKey;
	inputPrivacyKeyStarGiftsAutoSave#e1732341 = InputPrivacyKey;
	inputPrivacyKeyNoPaidMessages#bdc597b4 = InputPrivacyKey;

	privacyKeyStatusTimestamp#bc2eab30 = PrivacyKey;
	privacyKeyChatInvite#500e6dfa = PrivacyKey;
	privacyKeyPhoneCall#3d662b7b = PrivacyKey;
	privacyKeyPhoneP2P#39491cc8 = PrivacyKey;
	privacyKeyForwards#69ec56a3 = PrivacyKey;
	privacyKeyProfilePhoto#96151fed = PrivacyKey;
	privacyKeyPhoneNumber#d19ae46d = PrivacyKey;
	privacyKeyAddedByPhone#42ffd42b = PrivacyKey;
	privacyKeyVoiceMessages#697f414 = PrivacyKey;
	privacyKeyAbout#a486b761 = PrivacyKey;
	privacyKeyBirthday#2000a518 = PrivacyKey;
	privacyKeyStarGiftsAutoSave#2ca4fdf8 = PrivacyKey;
	privacyKeyNoPaidMessages#17d348d2 = PrivacyKey;
**/

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
)

func FromInputPrivacyKeyType(k *InputPrivacyKey) int {
	switch k.Clazz.InputPrivacyKeyClazzName() {
	case ClazzName_inputPrivacyKeyStatusTimestamp:
		return STATUS_TIMESTAMP
	case ClazzName_inputPrivacyKeyChatInvite:
		return CHAT_INVITE
	case ClazzName_inputPrivacyKeyPhoneCall:
		return PHONE_CALL
	case ClazzName_inputPrivacyKeyPhoneP2P:
		return PHONE_P2P
	case ClazzName_inputPrivacyKeyForwards:
		return FORWARDS
	case ClazzName_inputPrivacyKeyProfilePhoto:
		return PROFILE_PHOTO
	case ClazzName_inputPrivacyKeyPhoneNumber:
		return PHONE_NUMBER
	case ClazzName_inputPrivacyKeyAddedByPhone:
		return ADDED_BY_PHONE
	case ClazzName_inputPrivacyKeyVoiceMessages:
		return VOICE_MESSAGES
	case ClazzName_inputPrivacyKeyAbout:
		return ABOUT
	case ClazzName_inputPrivacyKeyBirthday:
		return BIRTHDAY
	case ClazzName_inputPrivacyKeyStarGiftsAutoSave:
		return STAR_GIFTS_AUTO_SAVE
	case ClazzName_inputPrivacyKeyNoPaidMessages:
		return NO_PAID_MESSAGES
	default:
		panic("type is invalid")
	}
	return KEY_TYPE_INVALID
}

var (
	cachePrivacyKeyStatusTimestamp   = MakeTLPrivacyKeyStatusTimestamp(&TLPrivacyKeyStatusTimestamp{})
	cachePrivacyKeyChatInvite        = MakeTLPrivacyKeyChatInvite(&TLPrivacyKeyChatInvite{})
	cachePrivacyKeyPhoneCall         = MakeTLPrivacyKeyPhoneCall(&TLPrivacyKeyPhoneCall{})
	cachePrivacyKeyPhoneP2P          = MakeTLPrivacyKeyPhoneP2P(&TLPrivacyKeyPhoneP2P{})
	cachePrivacyKeyForwards          = MakeTLPrivacyKeyForwards(&TLPrivacyKeyForwards{})
	cachePrivacyKeyProfilePhoto      = MakeTLPrivacyKeyProfilePhoto(&TLPrivacyKeyProfilePhoto{})
	cachePrivacyKeyPhoneNumber       = MakeTLPrivacyKeyPhoneNumber(&TLPrivacyKeyPhoneNumber{})
	cachePrivacyKeyAddedByPhone      = MakeTLPrivacyKeyAddedByPhone(&TLPrivacyKeyAddedByPhone{})
	cachePrivacyKeyVoiceMessages     = MakeTLPrivacyKeyVoiceMessages(&TLPrivacyKeyVoiceMessages{})
	cachePrivacyKeyAbout             = MakeTLPrivacyKeyAbout(&TLPrivacyKeyAbout{})
	cachePrivacyKeyBirthday          = MakeTLPrivacyKeyBirthday(&TLPrivacyKeyBirthday{})
	cachePrivacyKeyStarGiftsAutoSave = MakeTLPrivacyKeyStarGiftsAutoSave(&TLPrivacyKeyStarGiftsAutoSave{})
	cachePrivacyKeyNoPaidMessages    = MakeTLPrivacyKeyNoPaidMessages(&TLPrivacyKeyNoPaidMessages{})
)

func ToPrivacyKey(keyType int) (key PrivacyKeyClazz) {
	switch keyType {
	case STATUS_TIMESTAMP:
		key = cachePrivacyKeyStatusTimestamp
	case CHAT_INVITE:
		key = cachePrivacyKeyChatInvite
	case PHONE_CALL:
		key = cachePrivacyKeyPhoneCall
	case PHONE_P2P:
		key = cachePrivacyKeyPhoneP2P
	case FORWARDS:
		key = cachePrivacyKeyForwards
	case PROFILE_PHOTO:
		key = cachePrivacyKeyProfilePhoto
	case PHONE_NUMBER:
		key = cachePrivacyKeyPhoneNumber
	case ADDED_BY_PHONE:
		key = cachePrivacyKeyAddedByPhone
	case VOICE_MESSAGES:
		key = cachePrivacyKeyVoiceMessages
	case ABOUT:
		key = cachePrivacyKeyAbout
	case BIRTHDAY:
		key = cachePrivacyKeyBirthday
	case STAR_GIFTS_AUTO_SAVE:
		key = cachePrivacyKeyStarGiftsAutoSave
	case NO_PAID_MESSAGES:
		key = cachePrivacyKeyNoPaidMessages
	default:
		panic("type is invalid")
	}
	return
}

/*
func ToPrivacyRuleByInput(userSelfId int64, inputRule *InputPrivacyRule) *PrivacyRule {
	switch inputRule.InputPrivacyRuleClazzName() {
	case ClazzName_inputPrivacyValueAllowAll:
		return MakePrivacyRule(&TLPrivacyValueAllowAll{})
	case ClazzName_inputPrivacyValueAllowContacts:
		return MakePrivacyRule(&TLPrivacyValueAllowContacts{})
	case ClazzName_inputPrivacyValueAllowUsers:
		return MakePrivacyRule(&TLPrivacyValueAllowUsers{
			Users: ToUserIdListByInput(userSelfId, inputRule.GetUsers()),
		})
	case ClazzName_inputPrivacyValueDisallowAll:
		return MakePrivacyRule(&TLPrivacyValueDisallowAll{})
	case ClazzName_inputPrivacyValueDisallowContacts:
		return MakePrivacyRule(&TLPrivacyValueDisallowContacts{})
	case ClazzName_inputPrivacyValueDisallowUsers:
		return MakePrivacyRule(&TLPrivacyValueDisallowUsers{
			Users: ToUserIdListByInput(userSelfId, inputRule.GetUsers()),
		})
	case ClazzName_inputPrivacyValueAllowChatParticipants:
		return MakePrivacyRule(&TLPrivacyValueAllowChatParticipants{
			Chats: inputRule.GetChats(),
		})
	case ClazzName_inputPrivacyValueDisallowChatParticipants:
		return MakePrivacyRule(&TLPrivacyValueDisallowChatParticipants{
			Chats: inputRule.GetChats(),
		})
	case ClazzName_inputPrivacyValueAllowCloseFriends:
		return MakePrivacyRule(&TLPrivacyValueAllowCloseFriends{})
	case ClazzName_inputPrivacyValueAllowPremium:
		return MakePrivacyRule(&TLPrivacyValueAllowPremium{})
	case ClazzName_inputPrivacyValueAllowBots:
		return MakePrivacyRule(&TLPrivacyValueAllowBots{})
	case ClazzName_inputPrivacyValueDisallowBots:
		return MakePrivacyRule(&TLPrivacyValueDisallowBots{})
	default:
		panic("type is invalid")
	}
	return nil
}

func ToPrivacyRuleListByInput(userSelfId int64, inputRules []*InputPrivacyRule) (rules []*PrivacyRule) {
	rules = make([]*PrivacyRule, 0, len(inputRules))
	for _, inputRule := range inputRules {
		rules = append(rules, ToPrivacyRuleByInput(userSelfId, inputRule))
	}
	return
}

// PickAllIdListByRules
// TODO(@benqi): pick chat and channel
func PickAllIdListByRules(rules []*PrivacyRule) (userIdList, chatIdList, channelIdList []int64) {
	userIdList = make([]int64, 0)
	chatIdList = make([]int64, 0)
	channelIdList = make([]int64, 0)

	for _, r := range rules {
		r.Match(
			func(c *TLPrivacyValueAllowUsers) interface{} {
				if len(c.Users) > 0 {
					userIdList = append(userIdList, c.Users...)
				}

				return nil
			},
			func(c *TLPrivacyValueDisallowUsers) interface{} {
				if len(c.Users) > 0 {
					userIdList = append(userIdList, c.Users...)
				}

				return nil
			},
			func(c *TLPrivacyValueAllowChatParticipants) interface{} {
				for _, id := range c.Chats {
					if id >= MinNebulaChatChannelID {
						channelIdList = append(channelIdList, id)
					} else {
						chatIdList = append(chatIdList, id)
					}
				}

				return nil
			},
			func(c *TLPrivacyValueDisallowChatParticipants) interface{} {
				for _, id := range c.Chats {
					if id >= MinNebulaChatChannelID {
						channelIdList = append(channelIdList, id)
					} else {
						chatIdList = append(chatIdList, id)
					}
				}

				return nil
			})
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
