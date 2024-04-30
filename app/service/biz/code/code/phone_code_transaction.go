// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package code

import (
	"github.com/teamgram/proto/mtproto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ToAuthSentCode
// ///////////////////////////////////////////////////////////////////////////////////////////////////
// TODO(@benqi): 如果手机号已经注册，检查是否有其他设备在线，有则使用sentCodeTypeApp
//
//	否则使用sentCodeTypeSms
//
// TODO(@benqi): 有则使用sentCodeTypeFlashCall和entCodeTypeCall？？
func (m *PhoneCodeTransaction) ToAuthSentCode() *mtproto.Auth_SentCode {
	// TODO(@benqi): only use sms

	authSentCode := mtproto.MakeTLAuthSentCode(&mtproto.Auth_SentCode{
		Type:          makeAuthSentCodeType(m.SentCodeType, len(m.PhoneCode), m.FlashCallPattern),
		PhoneCodeHash: m.PhoneCodeHash,
		NextType:      makeAuthCodeType(m.NextCodeType),
		Timeout:       &wrapperspb.Int32Value{Value: 60}, // TODO(@benqi): 默认60s
	}).To_Auth_SentCode()
	if m.SentCodeType == CodeTypeApp {
		authSentCode.Timeout = nil
	}
	return authSentCode
}
