// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: @benqi (wubenqi@gmail.com)

package model

import (
	"fmt"

	"github.com/teamgram/proto/v2/tg"

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
	CodeStateOk          = 1
	CodeStateSend        = 2
	CodeStateSent        = 3
	CodeStateReSent      = 6
	CodeStateSignIn      = 4
	CodeStateSignUp      = 5
	CodeStateChangePhone = 6
	CodeStateDeleted     = -1
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

func makeAuthCodeType(codeType int) tg.AuthCodeTypeClazz {
	switch codeType {
	case CodeTypeSms:
		return tg.MakeTLAuthCodeTypeSms(&tg.TLAuthCodeTypeSms{})
	case CodeTypeCall:
		return tg.MakeTLAuthCodeTypeCall(&tg.TLAuthCodeTypeCall{})
	case CodeTypeFlashCall:
		return tg.MakeTLAuthCodeTypeFlashCall(&tg.TLAuthCodeTypeFlashCall{})
	case CodeTypeFragmentSms:
		return tg.MakeTLAuthCodeTypeFragmentSms(&tg.TLAuthCodeTypeFragmentSms{})
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
func makeAuthSentCodeType(codeType, codeLength int, pattern string) (authSentCodeType tg.AuthSentCodeTypeClazz) {
	switch codeType {
	case SentCodeTypeApp:
		authSentCodeType = tg.MakeTLAuthSentCodeTypeApp(&tg.TLAuthSentCodeTypeApp{
			Length: int32(codeLength),
		})
	case SentCodeTypeSms:
		authSentCodeType = tg.MakeTLAuthSentCodeTypeSms(&tg.TLAuthSentCodeTypeSms{
			Length: int32(codeLength),
		})
	case SentCodeTypeCall:
		authSentCodeType = tg.MakeTLAuthSentCodeTypeCall(&tg.TLAuthSentCodeTypeCall{
			Length: int32(codeLength),
		})
	case SentCodeTypeFlashCall:
		authSentCodeType = tg.MakeTLAuthSentCodeTypeFlashCall(&tg.TLAuthSentCodeTypeFlashCall{
			Pattern: pattern,
		})
	case SentCodeTypeMissedCall:
		authSentCodeType = tg.MakeTLAuthSentCodeTypeMissedCall(&tg.TLAuthSentCodeTypeMissedCall{
			Prefix: "",
			Length: int32(codeLength),
		})
	case SentCodeTypeEmailCode:
		authSentCodeType = tg.MakeTLAuthSentCodeTypeEmailCode(&tg.TLAuthSentCodeTypeEmailCode{
			AppleSigninAllowed:   false,
			GoogleSigninAllowed:  false,
			EmailPattern:         "",
			Length:               int32(codeLength),
			ResetAvailablePeriod: nil,
			ResetPendingDate:     nil,
		})
	case SentCodeTypeSetUpEmailRequired:
		authSentCodeType = tg.MakeTLAuthSentCodeTypeSetUpEmailRequired(&tg.TLAuthSentCodeTypeSetUpEmailRequired{
			AppleSigninAllowed:  false,
			GoogleSigninAllowed: false,
		})
	case SentCodeTypeFirebaseSms:
		authSentCodeType = tg.MakeTLAuthSentCodeTypeFirebaseSms(&tg.TLAuthSentCodeTypeFirebaseSms{
			Nonce:                  nil,
			PlayIntegrityProjectId: nil,
			PlayIntegrityNonce:     nil,
			Receipt:                nil,
			PushTimeout:            nil,
			Length:                 0,
		})
	case SentCodeTypeSmsWord:
		authSentCodeType = tg.MakeTLAuthSentCodeTypeSmsWord(&tg.TLAuthSentCodeTypeSmsWord{
			Beginning: nil,
		})
	case SentCodeTypeSmsPhrase:
		authSentCodeType = tg.MakeTLAuthSentCodeTypeSmsPhrase(&tg.TLAuthSentCodeTypeSmsPhrase{
			Beginning: nil,
		})
	default:
		// code bug.
		err := fmt.Errorf("invalid sentCodeType: %d", codeType)
		logx.Errorf("makeAuthSentCodeType - %v", err)
		panic(err)
	}

	return
}
