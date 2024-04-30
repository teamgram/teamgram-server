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

package model

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

/*
	auth.codeTypeSms#72a3158c = auth.CodeType;
	auth.codeTypeCall#741cd3e3 = auth.CodeType;
	auth.codeTypeFlashCall#226ccefb = auth.CodeType;
	auth.codeTypeMissedCall#d61ad6ee = auth.CodeType;
	auth.codeTypeFragmentSms#6ed998c = auth.CodeType;

	auth.sentCodeTypeApp#3dbb5986 length:int = auth.SentCodeType;
	auth.sentCodeTypeSms#c000bba2 length:int = auth.SentCodeType;
	auth.sentCodeTypeCall#5353e5a7 length:int = auth.SentCodeType;
	auth.sentCodeTypeFlashCall#ab03c6d9 pattern:string = auth.SentCodeType;
	auth.sentCodeTypeMissedCall#82006484 prefix:string length:int = auth.SentCodeType;
	auth.sentCodeTypeEmailCode#f450f59b flags:# apple_signin_allowed:flags.0?true google_signin_allowed:flags.1?true email_pattern:string length:int reset_available_period:flags.3?int reset_pending_date:flags.4?int = auth.SentCodeType;
	auth.sentCodeTypeSetUpEmailRequired#a5491dea flags:# apple_signin_allowed:flags.0?true google_signin_allowed:flags.1?true = auth.SentCodeType;
	auth.sentCodeTypeFragmentSms#d9565c39 url:string length:int = auth.SentCodeType;
	auth.sentCodeTypeFirebaseSms#e57b1432 flags:# nonce:flags.0?bytes receipt:flags.1?string push_timeout:flags.1?int length:int = auth.SentCodeType;
	auth.sentCodeTypeSmsWord#a416ac81 flags:# beginning:flags.0?string = auth.SentCodeType;
	auth.sentCodeTypeSmsPhrase#b37794af flags:# beginning:flags.0?string = auth.SentCodeType;
*/

const (
	CodeTypeNone        = 0
	CodeTypeSms         = 1
	CodeTypeCall        = 2
	CodeTypeFlashCall   = 3
	CodeTypeMissedCall  = 4
	CodeTypeFragmentSms = 5
)

const (
	SentCodeTypeNone               = 0
	SentCodeTypeApp                = 1
	SentCodeTypeSms                = 2
	SentCodeTypeCall               = 3
	SentCodeTypeFlashCall          = 4
	SentCodeTypeMissedCall         = 5
	SentCodeTypeEmailCode          = 6
	SentCodeTypeSetUpEmailRequired = 7
	SentCodeTypeFragmentSms        = 8
	SentCodeTypeFirebaseSms        = 9
	SentCodeTypeSmsWord            = 10
	SentCodeTypeSmsPhrase          = 11
)

const (
	CodeStateOk      = 1
	CodeStateSend    = 2
	CodeStateSent    = 3
	CodeStateReSent  = 6
	CodeStateSignIn  = 4
	CodeStateSignUp  = 5
	CodeStateDeleted = -1
)

/////////////////////////////////////////////////////////////////////////////////////////////////////////
// by params(phoneRegistered, allowFlashCall, currentNumber) ==> sentType and nextType

// MakeCodeType
// FIXME(@benqi): ignore it.
func MakeCodeType(phoneRegistered, allowFlashCall, currentNumber bool) (int, int) {
	//if phoneRegistered {
	//	// TODO(@benqi): check other session online
	//	authSentCodeType := &mtproto.TLAuthSentCodeTypeApp{Data2: &mtproto.Auth_SentCodeType_Data{
	//		Length: code.GetPhoneCodeLength(),
	//	}}
	//	authSentCode.SetType(authSentCodeType.To_Auth_SentCodeType())
	//} else {
	//	// TODO(@benqi): sentCodeTypeFlashCall and sentCodeTypeCall, nextType
	//	// telegramd, we only use sms
	//	authSentCodeType := &mtproto.TLAuthSentCodeTypeSms{Data2: &mtproto.Auth_SentCodeType_Data{
	//		Length: code.GetPhoneCodeLength(),
	//	}}
	//	authSentCode.SetType(authSentCodeType.To_Auth_SentCodeType())
	//
	//	// TODO(@benqi): nextType
	//	// authSentCode.SetNextType()
	//}
	_ = phoneRegistered
	_ = allowFlashCall
	_ = currentNumber

	sentCodeType := SentCodeTypeApp
	nextCodeType := CodeTypeNone
	return sentCodeType, nextCodeType
}

func makeAuthCodeType(codeType int) *mtproto.Auth_CodeType {
	switch codeType {
	case CodeTypeSms:
		return mtproto.MakeTLAuthCodeTypeSms(nil).To_Auth_CodeType()
	case CodeTypeCall:
		return mtproto.MakeTLAuthCodeTypeCall(nil).To_Auth_CodeType()
	case CodeTypeFlashCall:
		return mtproto.MakeTLAuthCodeTypeFlashCall(nil).To_Auth_CodeType()
	case CodeTypeFragmentSms:
		return mtproto.MakeTLAuthCodeTypeFragmentSms(nil).To_Auth_CodeType()
	default:
		return nil
	}
}

// SentCodeTypeApp                = 1
// SentCodeTypeSms                = 2
// SentCodeTypeCall               = 3
// SentCodeTypeFlashCall          = 4
// SentCodeTypeMissedCall         = 5
// SentCodeTypeEmailCode          = 6
// SentCodeTypeSetUpEmailRequired = 7
// SentCodeTypeFragmentSms        = 8
// SentCodeTypeFirebaseSms        = 9
// SentCodeTypeSmsWord            = 10
// SentCodeTypeSmsPhrase          = 11
func makeAuthSentCodeType(codeType, codeLength int, pattern string) (authSentCodeType *mtproto.Auth_SentCodeType) {
	switch codeType {
	case SentCodeTypeApp:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeApp(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case SentCodeTypeSms:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeSms(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case SentCodeTypeCall:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeCall(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case SentCodeTypeFlashCall:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeFlashCall(&mtproto.Auth_SentCodeType{
			Length:  int32(codeLength),
			Pattern: pattern,
		}).To_Auth_SentCodeType()
	case SentCodeTypeMissedCall:
	case SentCodeTypeEmailCode:
	case SentCodeTypeSetUpEmailRequired:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeSetUpEmailRequired(&mtproto.Auth_SentCodeType{
			AppleSigninAllowed:  false,
			GoogleSigninAllowed: false,
		}).To_Auth_SentCodeType()
	case SentCodeTypeFirebaseSms:
	case SentCodeTypeSmsWord:
	case SentCodeTypeSmsPhrase:
	default:
		// code bug.
		err := fmt.Errorf("invalid sentCodeType: %d", codeType)
		logx.Errorf("makeAuthSentCodeType - %v", err)
		panic(err)
	}

	return
}
