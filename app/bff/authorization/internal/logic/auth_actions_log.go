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

package logic

import (
	"github.com/gogo/protobuf/proto"
	"github.com/teamgram/proto/mtproto"
)

const (
	opTypeUnknown    = 0
	opTypeSendCode   = 1
	opTypeSignUp     = 2
	opTypeSignIn     = 3
	opTypeLogout     = 4
	opTypeResendCode = 5
	opTypeCancelCode = 6
)

func GetActionType(request proto.Message) int {
	switch request.(type) {
	case *mtproto.TLAuthSendCode:
		return opTypeSendCode
	case *mtproto.TLAuthResendCode:
		return opTypeResendCode
	case *mtproto.TLAuthSignIn:
		return opTypeSignIn
	case *mtproto.TLAuthSignUp:
		return opTypeSignUp
	case *mtproto.TLAuthLogOut3E72BA19:
		return opTypeLogout
	case *mtproto.TLAuthLogOut5717DA40:
		return opTypeLogout
	case *mtproto.TLAuthCancelCode:
		return opTypeCancelCode
	}
	return opTypeUnknown
}

//// async
//func DoLogAuthAction(d *dao.Dao, md *metadata.RpcMetadata, phoneNumber string, actionType int, log string) {
//	go func(authKeyId, msgId int64, clientIp string, phoneNumber string, actionType int, log string) {
//		d.LogAuthAction(context.Background(), authKeyId, msgId, clientIp, phoneNumber, actionType, log)
//	}(md.AuthId, md.ClientMsgId, md.ClientAddr, phoneNumber, actionType, log)
//}
