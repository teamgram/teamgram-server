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

package phonenumber

import (
	"errors"
	"fmt"
	"strings"

	"github.com/teamgram/proto/v2/tg"

	"github.com/nyaruka/phonenumbers"
)

type phoneNumberHelper struct {
	*phonenumbers.PhoneNumber
}

func MakePhoneNumberHelper(number, region string) (*phoneNumberHelper, error) {
	var (
		pNumber *phonenumbers.PhoneNumber
		err     error
	)

	if number == "" {
		return nil, errors.New("empty phone number")
	}

	// Android phone number format: 8611111111111, parse error: invalid country code
	// convert +8611111111111
	if region == "" && number[:1] != "+" {
		number = "+" + number
	}

	// check phone invalid
	pNumber, err = phonenumbers.Parse(number, region)
	if err != nil {
		err = fmt.Errorf("parse phone number %s err: %v", number, err)
	} else {
		if !phonenumbers.IsValidNumber(pNumber) {
			err = fmt.Errorf("invalid phone number: %s - %v", number, pNumber)
		}
	}

	if err != nil {
		return nil, err
	} else {
		return &phoneNumberHelper{pNumber}, nil
	}
}

func (p *phoneNumberHelper) GetNormalizeDigits() string {
	// DB store normalize phone number
	return phonenumbers.NormalizeDigitsOnly(phonenumbers.Format(p.PhoneNumber, phonenumbers.E164))
}

func (p *phoneNumberHelper) GetRegionCode() string {
	return phonenumbers.GetRegionCodeForNumber(p.PhoneNumber)
}

func (p *phoneNumberHelper) GetCountryCode() int32 {
	return p.PhoneNumber.GetCountryCode()
}

// CheckAndGetPhoneNumber
// Check number
// receive from client : "+86 111 1111 1111", need normalize
func CheckAndGetPhoneNumber(number string) (phoneNumber string, err error) {
	var (
		pNumber *phoneNumberHelper
	)

	pNumber, err = MakePhoneNumberHelper(number, "")
	if err != nil {
		return
	}

	return pNumber.GetNormalizeDigits(), nil
}

func CheckPhoneNumberInvalid(phone string) (string, string, error) {
	// 3. check number
	// 3.1. empty
	if phone == "" {
		// log.Errorf("check phone_number error - empty")
		return "", "", tg.Err406PhoneNumberInvalid
	}

	phone = strings.ReplaceAll(phone, " ", "")
	switch phone {
	case "+42400":
		return "", phone[1:], nil
	case "+42401":
		return "", phone[1:], nil
	case "+42777":
		return "", phone[1:], nil
	case "42400":
		return "", phone, nil
	case "42401":
		return "", phone, nil
	case "42777":
		return "", phone, nil
	}

	// fragment
	if strings.HasPrefix(phone, "+888") {
		if len(phone) == 12 {
			// +888 0888 0080
			return "", phone[1:], nil
		} else {
			return "", "", tg.Err406PhoneNumberInvalid
		}
	} else if strings.HasPrefix(phone, "888") {
		if len(phone) == 11 {
			// +888 0888 0080
			return "", phone, nil
		} else {
			return "", "", tg.Err406PhoneNumberInvalid
		}
	}

	// 3.2. check phone_number
	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	// We need getRegionCode from phone_number
	pNumber, err := MakePhoneNumberHelper(phone, "")
	if err != nil {
		// log.Errorf("check phone_number error - %v", err)
		// err = mtproto.ErrPhoneNumberInvalid
		return "", "", tg.Err406PhoneNumberInvalid
	}

	return pNumber.GetRegionCode(), pNumber.GetNormalizeDigits(), nil
}
