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

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/pkg/phonenumber"
)

// AuthCancelCode
// auth.cancelCode#1f040578 phone_number:string phone_code_hash:string = Bool;
func (c *AuthorizationCore) AuthCancelCode(in *mtproto.TLAuthCancelCode) (*mtproto.Bool, error) {
	//// 1. check phone code
	//if request.PhoneCodeHash == "" {
	//	err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PHONE_CODE_EMPTY)
	//	return nil, err
	//}
	//
	// 2. check number
	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	_, err := phonenumber.CheckAndGetPhoneNumber(in.GetPhoneNumber())
	if err != nil {
		c.Logger.Errorf("check phone_number error - %v", err)
		err = mtproto.ErrPhoneNumberInvalid
		return nil, err
	}

	// code := logic.NewAuthSignLogic(s.AuthFacade)
	// canceled := mtproto.ToBool(code.DoAuthCancelCode(md.AuthId, phoneNumber, request.PhoneCodeHash))

	return mtproto.BoolTrue, nil
}
