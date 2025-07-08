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
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/tg"
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

func GetActionType(request iface.TLObject) int {
	switch request.(type) {
	case *tg.TLAuthSendCode:
		return opTypeSendCode
	case *tg.TLAuthResendCode:
		return opTypeResendCode
	case *tg.TLAuthSignIn:
		return opTypeSignIn
	case *tg.TLAuthSignUp:
		return opTypeSignUp
	case *tg.TLAuthLogOut:
		return opTypeLogout
	case *tg.TLAuthCancelCode:
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
