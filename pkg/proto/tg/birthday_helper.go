// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

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

	birthday, _ := m.ToBirthday()
	if birthday.Year == nil {
		v = fmt.Sprintf("%02d-%02d", birthday.Month, birthday.Day)
	} else {
		v = fmt.Sprintf("%04d-%02d-%02d", GetFlagsInt32(birthday.Year), birthday.Month, birthday.Day)
	}

	return
}
