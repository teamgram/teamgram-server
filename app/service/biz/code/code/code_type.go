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

package code

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

/**
  auth.codeTypeSms#72a3158c = auth.CodeType;
  auth.codeTypeCall#741cd3e3 = auth.CodeType;
  auth.codeTypeFlashCall#226ccefb = auth.CodeType;

  auth.sentCodeTypeApp#3dbb5986 length:int = auth.SentCodeType;
  auth.sentCodeTypeSms#c000bba2 length:int = auth.SentCodeType;
  auth.sentCodeTypeCall#5353e5a7 length:int = auth.SentCodeType;
  auth.sentCodeTypeFlashCall#ab03c6d9 pattern:string = auth.SentCodeType;
*/

const (
	CodeTypeNone      = 0
	CodeTypeApp       = 1
	CodeTypeSms       = 2
	CodeTypeCall      = 3
	CodeTypeFlashCall = 4
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

// MakeCodeType
// by params(phoneRegistered, allowFlashCall, currentNumber) ==> sentType and nextType
//
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

	sentCodeType := CodeTypeApp
	nextCodeType := CodeTypeNone
	return sentCodeType, nextCodeType
}

func makeAuthCodeType(codeType int32) *mtproto.Auth_CodeType {
	switch codeType {
	case CodeTypeSms:
		return mtproto.MakeTLAuthCodeTypeSms(nil).To_Auth_CodeType()
	case CodeTypeCall:
		return mtproto.MakeTLAuthCodeTypeCall(nil).To_Auth_CodeType()
	case CodeTypeFlashCall:
		return mtproto.MakeTLAuthCodeTypeFlashCall(nil).To_Auth_CodeType()
	default:
		return nil
	}
}

func makeAuthSentCodeType(codeType int32, codeLength int, pattern string) (authSentCodeType *mtproto.Auth_SentCodeType) {
	switch codeType {
	case CodeTypeApp:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeApp(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case CodeTypeSms:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeSms(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case CodeTypeCall:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeCall(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case CodeTypeFlashCall:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeFlashCall(&mtproto.Auth_SentCodeType{
			Length:  int32(codeLength),
			Pattern: pattern,
		}).To_Auth_SentCodeType()
	default:
		// code bug.
		err := fmt.Errorf("invalid sentCodeType: %d", codeType)
		logx.Errorf("makeAuthSentCodeType - %v", err)
		panic(err)
	}

	return
}
