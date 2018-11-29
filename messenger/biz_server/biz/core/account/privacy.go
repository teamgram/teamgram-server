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
	"encoding/json"
	"github.com/golang/glog"
	"github.com/nebula-chat/chatengine/messenger/biz_server/biz/dal/dataobject"
	"github.com/nebula-chat/chatengine/mtproto"
)

/*
 privacy:
	inputPrivacyKeyStatusTimestamp#4f96cb18 = InputPrivacyKey;
	inputPrivacyKeyChatInvite#bdfb0426 = InputPrivacyKey;
	inputPrivacyKeyPhoneCall#fabadc5f = InputPrivacyKey;

	privacyKeyStatusTimestamp#bc2eab30 = PrivacyKey;
	privacyKeyChatInvite#500e6dfa = PrivacyKey;
	privacyKeyPhoneCall#3d662b7b = PrivacyKey;

	inputPrivacyValueAllowContacts#d09e07b = InputPrivacyRule;
	inputPrivacyValueAllowAll#184b35ce = InputPrivacyRule;
	inputPrivacyValueAllowUsers#131cc67f users:Vector<InputUser> = InputPrivacyRule;
	inputPrivacyValueDisallowContacts#ba52007 = InputPrivacyRule;
	inputPrivacyValueDisallowAll#d66b66c9 = InputPrivacyRule;
	inputPrivacyValueDisallowUsers#90110467 users:Vector<InputUser> = InputPrivacyRule;

	privacyValueAllowContacts#fffe1bac = PrivacyRule;
	privacyValueAllowAll#65427b82 = PrivacyRule;
	privacyValueAllowUsers#4d5bbe0c users:Vector<int> = PrivacyRule;
	privacyValueDisallowContacts#f888fa1a = PrivacyRule;
	privacyValueDisallowAll#8b73e763 = PrivacyRule;
	privacyValueDisallowUsers#c7f49b7 users:Vector<int> = PrivacyRule;

	account.privacyRules#554abb6f rules:Vector<PrivacyRule> users:Vector<User> = account.PrivacyRules;

 update:
	updatePrivacy#ee3b272a key:PrivacyKey rules:Vector<PrivacyRule> = Update;

 api:
	account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
	account.setPrivacy#c9f81ce8 key:InputPrivacyKey rules:Vector<InputPrivacyRule> = account.PrivacyRules;
*/

/*
	int newType = currentType;
	if (position == nobodyRow) {
		newType = 1;
	} else if (position == everybodyRow) {
		newType = 0;
	} else if (position == myContactsRow) {
		newType = 2;
	}
	if (newType == currentType) {
		return;
	}
	enableAnimation = true;
	doneButton.setVisibility(View.VISIBLE);
	lastCheckedType = currentType;
	currentType = newType;
	updateRows();
*/
//
//type privacyKeyType int8
//
//// TLRpcErrorCodes_OTHER2 TLRpcErrorCodes = 502
//
//const (
//	kPrivacyKeyType_Invalid privacyKeyType = 0
//	kPrivacyKeyType_StatusTimestamp privacyKeyType = 1
//	kPrivacyKeyType_ChatInvite privacyKeyType = 2
//	kPrivacyKeyType_PhoneCall privacyKeyType = 3
//)

//const (
//	kPrivacyRuleType_Invalid = 0
//	kPrivacyRuleType_AllowContacts = 1
//	kPrivacyRuleType_AllowAll = 2
//	kPrivacyRuleType_AllowUsers = 3
//	kPrivacyRuleType_DisallowContacts = 4
//	kPrivacyRuleType_DisallowAll = 5
//	kPrivacyRuleType_DisallowUsers = 6
//)

//func getInputPrivacyKeyType(key *mtproto.InputPrivacyKey) int8 {
//}

func FromInputPrivacyKey(key *mtproto.InputPrivacyKey) PrivacyKeyType {
	var m PrivacyKeyType
	switch key.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputPrivacyKeyStatusTimestamp:
		m = PrivacyKeyType_STATUS_TIMESTAMP
	case mtproto.TLConstructor_CRC32_inputPrivacyKeyChatInvite:
		m = PrivacyKeyType_CHAT_INVITE
	case mtproto.TLConstructor_CRC32_inputPrivacyKeyPhoneCall:
		m = PrivacyKeyType_PHONE_CALL
	default:
		glog.Fatal("key is invalid")
	}
	return m
}

func (x PrivacyKeyType) ToPrivacyKey() (key *mtproto.PrivacyKey) {
	switch x {
	case PrivacyKeyType_STATUS_TIMESTAMP:
		key = mtproto.NewTLPrivacyKeyStatusTimestamp().To_PrivacyKey()
	case PrivacyKeyType_CHAT_INVITE:
		key = mtproto.NewTLPrivacyKeyChatInvite().To_PrivacyKey()
	case PrivacyKeyType_PHONE_CALL:
		key = mtproto.NewTLPrivacyKeyPhoneCall().To_PrivacyKey()
	default:
		glog.Fatal("type is invalid")
	}
	return
}

////////////////////////////////////////////////////////////////////////////////////////
type privacyLogic struct {
	selfUserId int32
	dao        *accountsDAO
}

func (m *AccountModel) MakePrivacyLogic(userId int32) *privacyLogic {
	return &privacyLogic{
		selfUserId: userId,
		dao:        m.dao,
	}
}

func (m privacyLogic) GetPrivacy(key PrivacyKeyType) (rulesData *PrivacyRulesData) {
	do := m.dao.UserPrivacysDAO.SelectPrivacy(m.selfUserId, int8(key))
	if do == nil {
		return
	}
	rulesData2 := &PrivacyRulesData{}
	err := json.Unmarshal([]byte(do.Rules), rulesData2)
	if err != nil {
		glog.Errorf("getPrivacy - Unmarshal PrivacyRulesData(%d)error: %v", do.Id, err)
		// return nil, err
		return
	} else {
		rulesData = rulesData2
	}
	return
}

func (m privacyLogic) SetPrivacy(key PrivacyKeyType, rules []*mtproto.InputPrivacyRule) (rulesData *PrivacyRulesData) {
	rulesData = &PrivacyRulesData{
		Rules: make([]*PrivacyRuleData, 0, len(rules)),
	}
	for _, r := range rules {
		rulesData.Rules = append(rulesData.Rules, fromInputPrivacyRule(r))
	}

	// var err error
	rulesJson, _ := json.Marshal(rulesData)

	do := m.dao.UserPrivacysDAO.SelectPrivacy(m.selfUserId, int8(key))
	if do == nil {
		do := &dataobject.UserPrivacysDO{
			UserId:  m.selfUserId,
			KeyType: int8(key),
			Rules:   string(rulesJson),
		}
		m.dao.UserPrivacysDAO.Insert(do)
	} else {
		m.dao.UserPrivacysDAO.UpdatePrivacy(string(rulesJson), m.selfUserId, int8(key))
	}
	return
}

////////////////////////////////////////////////////////////////////////////////////////
func fromInputPrivacyRule(rule *mtproto.InputPrivacyRule) (ruleData *PrivacyRuleData) {
	switch rule.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputPrivacyValueAllowAll:
		ruleData = &PrivacyRuleData{
			Type: PrivacyRuleType_ALLOW_ALL,
		}
	case mtproto.TLConstructor_CRC32_inputPrivacyValueAllowContacts:
		ruleData = &PrivacyRuleData{
			Type: PrivacyRuleType_ALLOW_CONTACTS,
		}
	case mtproto.TLConstructor_CRC32_inputPrivacyValueAllowUsers:
		ruleData = &PrivacyRuleData{
			Type: PrivacyRuleType_ALLOW_USERS,
			// UserIdList: rule.GetData2().GetUsers(),
		}

		inputUsers := rule.GetData2().GetUsers()
		for _, u := range inputUsers {
			// TODO(@benqi): process inputUserSelf
			if u.GetConstructor() == mtproto.TLConstructor_CRC32_inputUser {
				ruleData.UserIdList = append(ruleData.UserIdList, u.GetData2().GetUserId())
			}
		}
	case mtproto.TLConstructor_CRC32_inputPrivacyValueDisallowAll:
		ruleData = &PrivacyRuleData{
			Type: PrivacyRuleType_DISALLOW_ALL,
		}
	case mtproto.TLConstructor_CRC32_inputPrivacyValueDisallowContacts:
		ruleData = &PrivacyRuleData{
			Type: PrivacyRuleType_DISALLOW_CONTACTS,
		}
	case mtproto.TLConstructor_CRC32_inputPrivacyValueDisallowUsers:
		ruleData = &PrivacyRuleData{
			Type: PrivacyRuleType_DISALLOW_USERS,
		}

		inputUsers := rule.GetData2().GetUsers()
		for _, u := range inputUsers {
			// TODO(@benqi): process inputUserSelf
			if u.GetConstructor() == mtproto.TLConstructor_CRC32_inputUser {
				ruleData.UserIdList = append(ruleData.UserIdList, u.GetData2().GetUserId())
			}
		}
	default:
		glog.Fatal("type is invalid")
	}
	return
}

func (m *PrivacyRuleData) ToPrivacyRule() (rule *mtproto.PrivacyRule) {
	switch m.GetType() {
	case PrivacyRuleType_ALLOW_ALL:
		rule = mtproto.NewTLPrivacyValueAllowAll().To_PrivacyRule()
	case PrivacyRuleType_ALLOW_CONTACTS:
		rule = mtproto.NewTLPrivacyValueAllowContacts().To_PrivacyRule()
	case PrivacyRuleType_ALLOW_USERS:
		rule = &mtproto.PrivacyRule{
			Constructor: mtproto.TLConstructor_CRC32_privacyValueAllowUsers,
			Data2: &mtproto.PrivacyRule_Data{
				Users: m.GetUserIdList(),
			},
		}
	case PrivacyRuleType_DISALLOW_ALL:
		rule = mtproto.NewTLPrivacyValueDisallowAll().To_PrivacyRule()
	case PrivacyRuleType_DISALLOW_CONTACTS:
		rule = mtproto.NewTLPrivacyValueDisallowContacts().To_PrivacyRule()
	case PrivacyRuleType_DISALLOW_USERS:
		rule = &mtproto.PrivacyRule{
			Constructor: mtproto.TLConstructor_CRC32_privacyValueAllowUsers,
			Data2: &mtproto.PrivacyRule_Data{
				Users: m.GetUserIdList(),
			},
		}
	default:
		glog.Fatal("rule type is invalid")
	}
	return
}

////////////////////////////////////////////////////////////////////////////////////////
func (m *PrivacyRulesData) ToPrivacyRuleList() (rules []*mtproto.PrivacyRule) {
	rules = make([]*mtproto.PrivacyRule, 0, len(m.GetRules()))
	for _, rule := range m.GetRules() {
		rules = append(rules, rule.ToPrivacyRule())
	}
	return
}

func (m *PrivacyRulesData) PickAllUserIdList() (idList []int32) {
	for _, rule := range m.Rules {
		idList = append(idList, rule.GetUserIdList()...)
	}
	return
}
