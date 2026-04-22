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

import (
	"fmt"
)

func FromBirthdayString(birthday string) (v BirthdayClazz) {
	if len(birthday) == 0 {
		return
	}

	// 1990-01-01
	var (
		year, month, day int32
	)

	_, err := fmt.Sscanf(birthday, "%d-%d-%d", &year, &month, &day)
	if err != nil {
		year = 0
		month = 0
		day = 0
	}

	v = MakeTLBirthday(&TLBirthday{
		Year:  MakeFlagsInt32(year),
		Month: month,
		Day:   day,
	})

	return
}

func (m *Birthday) ToBirthdayString() (v string) {
	if m == nil {
		return
	}

	birthday := m.ToBirthday()
	if birthday.Year == nil {
		v = fmt.Sprintf("%02d-%02d", birthday.Month, birthday.Day)
	} else {
		v = fmt.Sprintf("%04d-%02d-%02d", GetFlagsInt32(birthday.Year), birthday.Month, birthday.Day)
	}

	return
}
