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

package base

import (
	"fmt"
	"github.com/nebula-chat/chatengine/mtproto"
	"github.com/ttacon/libphonenumber"
)

type phoneNumberUtil struct {
	phoneNumber *libphonenumber.PhoneNumber
}

func MakePhoneNumberUtil(number, region string) (*phoneNumberUtil, error) {
	var (
		pnumber *libphonenumber.PhoneNumber
		err     error
	)

	if number == "" {
		err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), "phone number empty")
		return nil, err
	}

	// Android客户端手机号格式为: 8611111111111, Parse结果为invalid country code
	// 转换成+8611111111111，再进行Parse
	if region == "" && number[:1] != "+" {
		number = "+" + number
	}

	// fmt.Println(number)
	// check phone invalid
	pnumber, err = libphonenumber.Parse(number, region)
	// fmt.Println(pnumber)
	if err != nil {
		// fmt.Println(err)
		err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), fmt.Sprintf("invalid phone number: %v", err))
	} else {
		if !libphonenumber.IsValidNumber(pnumber) {
			err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), "invalid phone number")
		}
	}

	if err != nil {
		return nil, err
	} else {
		return &phoneNumberUtil{pnumber}, nil
	}
}

func (p *phoneNumberUtil) GetNormalizeDigits() string {
	// DB里存储归一化的phone
	return libphonenumber.NormalizeDigitsOnly(libphonenumber.Format(p.phoneNumber, libphonenumber.E164))
}

func (p *phoneNumberUtil) GetRegionCode() string {
	return libphonenumber.GetRegionCodeForNumber(p.phoneNumber)
}

// Check number
// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
func CheckAndGetPhoneNumber(number string) (phoneNumber string, err error) {
	var (
		pnumber *phoneNumberUtil
	)

	pnumber, err = MakePhoneNumberUtil(number, "")
	if err != nil {
		return
	}

	return pnumber.GetNormalizeDigits(), nil
}

//func CheckAndGetPhoneNumberByRegion(number, region string) (phoneNumber string, err error) {
//	if number == "" {
//		err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), "phone number empty")
//		return
//	}
//
//	// Android客户端手机号格式为: 8611111111111, Parse结果为invalid country code
//	// 转换成+8611111111111，再进行Parse
//	//if number[:1] != "+" {
//	//	number = "+" + number
//	//}
//
//	// fmt.Println(number)
//	// check phone invalid
//	var pnumber *libphonenumber.PhoneNumber
//	pnumber, err = libphonenumber.Parse(number, region)
//	// fmt.Println(pnumber)
//	if err != nil {
//		// fmt.Println(err)
//		err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), fmt.Sprintf("invalid phone number: %v", err))
//	} else {
//		if !libphonenumber.IsValidNumber(pnumber) {
//			err = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), "invalid phone number")
//		}
//	}
//
//	if err == nil {
//		// DB里存储归一化的phone
//		phoneNumber = libphonenumber.Format(pnumber, libphonenumber.E164)
//		// fmt.Println(phoneNumber)
//		// fmt.Println(libphonenumber.GetRegionCodeForNumber(pnumber))
//
//		phoneNumber = libphonenumber.NormalizeDigitsOnly(phoneNumber)
//
//		//phoneNumber = libphonenumber.Format(pnumber, libphonenumber.INTERNATIONAL)
//		//fmt.Println(phoneNumber)
//		//phoneNumber = libphonenumber.Format(pnumber, libphonenumber.NATIONAL)
//		//fmt.Println(phoneNumber)
//		//phoneNumber = libphonenumber.Format(pnumber, libphonenumber.RFC3966)
//		//fmt.Println(phoneNumber)
//		// phoneNumber = libphonenumber.NormalizeDigitsOnly(number)
//	}
//
//	return
//}
