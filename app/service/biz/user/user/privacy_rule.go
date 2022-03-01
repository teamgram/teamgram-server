// Copyright 2022 Teamgram Authors
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

package user

import (
	"github.com/teamgram/proto/mtproto"
)

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
)

/*
```
inputPrivacyKeyStatusTimestamp#4f96cb18 = InputPrivacyKey;
inputPrivacyKeyChatInvite#bdfb0426 = InputPrivacyKey;
inputPrivacyKeyPhoneCall#fabadc5f = InputPrivacyKey;
inputPrivacyKeyPhoneP2P#db9e70d2 = InputPrivacyKey;
inputPrivacyKeyForwards#a4dd4c08 = InputPrivacyKey;
inputPrivacyKeyProfilePhoto#5719bacc = InputPrivacyKey;
inputPrivacyKeyPhoneNumber#352dafa = InputPrivacyKey;
inputPrivacyKeyAddedByPhone#d1219bdd = InputPrivacyKey;
```
*/
const (
	KEY_TYPE_INVALID = 0
	STATUS_TIMESTAMP = 1 //
	CHAT_INVITE      = 2
	PHONE_CALL       = 3
	PHONE_P2P        = 4
	FORWARDS         = 5
	PROFILE_PHOTO    = 6
	PHONE_NUMBER     = 7
	ADDED_BY_PHONE   = 8
	MAX_KEY_TYPE     = 8
)

func FromInputPrivacyKeyType(k *mtproto.InputPrivacyKey) int {
	switch k.PredicateName {
	case mtproto.Predicate_inputPrivacyKeyStatusTimestamp:
		return STATUS_TIMESTAMP
	case mtproto.Predicate_inputPrivacyKeyChatInvite:
		return CHAT_INVITE
	case mtproto.Predicate_inputPrivacyKeyPhoneCall:
		return PHONE_CALL
	case mtproto.Predicate_inputPrivacyKeyPhoneP2P:
		return PHONE_P2P
	case mtproto.Predicate_inputPrivacyKeyForwards:
		return FORWARDS
	case mtproto.Predicate_inputPrivacyKeyProfilePhoto:
		return PROFILE_PHOTO
	case mtproto.Predicate_inputPrivacyKeyPhoneNumber:
		return PHONE_NUMBER
	case mtproto.Predicate_inputPrivacyKeyAddedByPhone:
		return ADDED_BY_PHONE
	}
	return KEY_TYPE_INVALID
}

func ToPrivacyKey(keyType int) (key *mtproto.PrivacyKey) {
	switch keyType {
	case STATUS_TIMESTAMP:
		key = mtproto.MakeTLPrivacyKeyStatusTimestamp(nil).To_PrivacyKey()
	case CHAT_INVITE:
		key = mtproto.MakeTLPrivacyKeyChatInvite(nil).To_PrivacyKey()
	case PHONE_CALL:
		key = mtproto.MakeTLPrivacyKeyPhoneCall(nil).To_PrivacyKey()
	case PHONE_P2P:
		key = mtproto.MakeTLPrivacyKeyPhoneP2P(nil).To_PrivacyKey()
	case FORWARDS:
		key = mtproto.MakeTLPrivacyKeyForwards(nil).To_PrivacyKey()
	case PROFILE_PHOTO:
		key = mtproto.MakeTLPrivacyKeyProfilePhoto(nil).To_PrivacyKey()
	case PHONE_NUMBER:
		key = mtproto.MakeTLPrivacyKeyPhoneNumber(nil).To_PrivacyKey()
	case ADDED_BY_PHONE:
		key = mtproto.MakeTLPrivacyKeyAddedByPhone(nil).To_PrivacyKey()
	default:
		panic("type is invalid")
	}
	return
}

// ToPrivacyRuleByInput
/*
```
inputPrivacyValueAllowContacts#d09e07b = InputPrivacyRule;
inputPrivacyValueAllowAll#184b35ce = InputPrivacyRule;
inputPrivacyValueAllowUsers#131cc67f users:Vector<InputUser> = InputPrivacyRule;
inputPrivacyValueDisallowContacts#ba52007 = InputPrivacyRule;
inputPrivacyValueDisallowAll#d66b66c9 = InputPrivacyRule;
inputPrivacyValueDisallowUsers#90110467 users:Vector<InputUser> = InputPrivacyRule;
inputPrivacyValueAllowChatParticipants#4c81c1ba chats:Vector<int> = InputPrivacyRule;
inputPrivacyValueDisallowChatParticipants#d82363af chats:Vector<int> = InputPrivacyRule;
```
*/
func ToPrivacyRuleByInput(userSelfId int64, inputRule *mtproto.InputPrivacyRule) *mtproto.PrivacyRule {
	switch inputRule.PredicateName {
	case mtproto.Predicate_inputPrivacyValueAllowAll:
		return mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueAllowContacts:
		return mtproto.MakeTLPrivacyValueAllowContacts(nil).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueAllowUsers:
		return mtproto.MakeTLPrivacyValueAllowUsers(&mtproto.PrivacyRule{
			Users: mtproto.ToUserIdListByInput(userSelfId, inputRule.GetUsers()),
		}).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueDisallowAll:
		return mtproto.MakeTLPrivacyValueDisallowAll(nil).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueDisallowContacts:
		return mtproto.MakeTLPrivacyValueDisallowContacts(nil).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueDisallowUsers:
		return mtproto.MakeTLPrivacyValueDisallowUsers(&mtproto.PrivacyRule{
			Users: mtproto.ToUserIdListByInput(userSelfId, inputRule.GetUsers()),
		}).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueAllowChatParticipants:
		return mtproto.MakeTLPrivacyValueAllowChatParticipants(&mtproto.PrivacyRule{
			Chats: inputRule.GetChats(),
		}).To_PrivacyRule()
	case mtproto.Predicate_inputPrivacyValueDisallowChatParticipants:
		return mtproto.MakeTLPrivacyValueDisallowChatParticipants(&mtproto.PrivacyRule{
			Chats: inputRule.GetChats(),
		}).To_PrivacyRule()
	default:
		// log.Errorf("type is invalid")
	}
	return nil
}

func ToPrivacyRuleListByInput(userSelfId int64, inputRules []*mtproto.InputPrivacyRule) (rules []*mtproto.PrivacyRule) {
	rules = make([]*mtproto.PrivacyRule, 0, len(inputRules))
	for _, inputRule := range inputRules {
		rules = append(rules, ToPrivacyRuleByInput(userSelfId, inputRule))
	}
	return
}

// PickAllIdListByRules
// TODO(@benqi): pick chat and channel
func PickAllIdListByRules(rules []*mtproto.PrivacyRule) (userIdList, chatIdList, channelIdList []int64) {
	userIdList = make([]int64, 0)
	chatIdList = make([]int64, 0)
	channelIdList = make([]int64, 0)
	for _, r := range rules {
		if len(r.Users) > 0 {
			userIdList = append(userIdList, r.Users...)
		}
		for _, id := range r.Chats {
			if id >= mtproto.MinNebulaChatChannelID {
				channelIdList = append(channelIdList, id)
			} else {
				chatIdList = append(chatIdList, id)
			}
		}
	}
	return
}

func CheckPrivacyIsAllow(selfId int64,
	rules []*mtproto.PrivacyRule,
	checkId int64,
	cbContact func(id, checkId int64) bool,
	cbChat func(checkId int64, idList []int64) bool) bool {
	ruleType := RULE_TYPE_INVALID

	for _, r := range rules {
		switch r.PredicateName {
		case mtproto.Predicate_privacyValueAllowAll:
			ruleType = ALLOW_ALL
		case mtproto.Predicate_privacyValueAllowContacts:
			ruleType = ALLOW_CONTACTS
		case mtproto.Predicate_privacyValueDisallowAll:
			ruleType = DISALLOW_ALL
		}
	}

	switch ruleType {
	case ALLOW_ALL:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return false
					}
				}
			case mtproto.Predicate_privacyValueDisallowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return false
				}
			}
		}
		return true
	case ALLOW_CONTACTS:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return true
					}
				}
			case mtproto.Predicate_privacyValueAllowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return true
				}
			case mtproto.Predicate_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return false
					}
				}
			case mtproto.Predicate_privacyValueDisallowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return false
				}
			}
		}
		return cbContact(selfId, checkId)
	case DISALLOW_ALL:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == checkId {
						return true
					}
				}
			case mtproto.Predicate_privacyValueAllowChatParticipants:
				if len(r.Chats) > 0 && cbChat(checkId, r.Chats) {
					return true
				}
			}
		}
		return false
	}

	return false
}

/*
	privacyValueAllowContacts#fffe1bac = PrivacyRule;
	privacyValueAllowAll#65427b82 = PrivacyRule;
	privacyValueAllowUsers#4d5bbe0c users:Vector<int> = PrivacyRule;
	privacyValueDisallowContacts#f888fa1a = PrivacyRule;
	privacyValueDisallowAll#8b73e763 = PrivacyRule;
	privacyValueDisallowUsers#c7f49b7 users:Vector<int> = PrivacyRule;
	privacyValueAllowChatParticipants#18be796b chats:Vector<int> = PrivacyRule;
	privacyValueDisallowChatParticipants#acae0690 chats:Vector<int> = PrivacyRule;
*/
func privacyIsAllow(rules []*mtproto.PrivacyRule, userId int64, isContact bool) bool {
	ruleType := RULE_TYPE_INVALID

	for _, r := range rules {
		switch r.PredicateName {
		case mtproto.Predicate_privacyValueAllowAll:
			ruleType = ALLOW_ALL
		case mtproto.Predicate_privacyValueAllowContacts:
			ruleType = ALLOW_CONTACTS
		case mtproto.Predicate_privacyValueDisallowAll:
			ruleType = DISALLOW_ALL
		}
	}

	switch ruleType {
	case ALLOW_ALL:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == userId {
						return false
					}
				}
			case mtproto.Predicate_privacyValueDisallowChatParticipants:
				return true
			}
		}
		return true
	case ALLOW_CONTACTS:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == userId {
						return true
					}
				}
			case mtproto.Predicate_privacyValueAllowChatParticipants:
				return true
			case mtproto.Predicate_privacyValueDisallowUsers:
				for _, id := range r.Users {
					if id == userId {
						return false
					}
				}
			case mtproto.Predicate_privacyValueDisallowChatParticipants:
				return true
			}
		}
		return isContact
	case DISALLOW_ALL:
		for _, r := range rules {
			switch r.PredicateName {
			case mtproto.Predicate_privacyValueAllowUsers:
				for _, id := range r.Users {
					if id == userId {
						return true
					}
				}
			case mtproto.Predicate_privacyValueAllowChatParticipants:
				return true
			}
		}
		return false
	}

	return false
}
