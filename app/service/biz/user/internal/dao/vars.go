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

package dao

import (
	"github.com/teamgram/proto/v2/tg"
	"github.com/zeromicro/go-zero/core/jsonx"
)

const (
	versionField = "0"
)

var (
	defaultRules = []*tg.PrivacyRule{
		tg.MakePrivacyRule(&tg.TLPrivacyValueAllowAll{}),
	}
	phoneNumberRules = []*tg.PrivacyRule{
		tg.MakePrivacyRule(&tg.TLPrivacyValueDisallowAll{}),
	}

	defaultRulesData     string
	phoneNumberRulesData string
)

func init() {
	defaultRulesData, _ = jsonx.MarshalToString(defaultRules)
	phoneNumberRulesData, _ = jsonx.MarshalToString(phoneNumberRules)
}

type idxId struct {
	idx int
	id  int64
}

func removeAllNil(contacts []*tg.ContactData) []*tg.ContactData {
	for i := 0; i < len(contacts); {
		if contacts[i] != nil {
			i++
			continue
		}

		if i < len(contacts)-1 {
			copy(contacts[i:], contacts[i+1:])
		}

		contacts[len(contacts)-1] = nil
		contacts = contacts[:len(contacts)-1]
	}

	return contacts
}

func makeDefaultPrivacyRules(key int32) []*tg.PrivacyRule {
	if key == tg.PHONE_NUMBER {
		return phoneNumberRules
	} else {
		return defaultRules
	}
}
