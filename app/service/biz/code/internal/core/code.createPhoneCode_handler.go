// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"errors"

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
)

var _ *tg.Bool

// CodeCreatePhoneCode
// code.createPhoneCode flags:# auth_key_id:long session_id:long phone:string phone_number_registered:flags.0?true sent_code_type:int next_code_type:int state:int = PhoneCodeTransaction;
func (c *CodeCore) CodeCreatePhoneCode(in *code.TLCodeCreatePhoneCode) (*code.PhoneCodeTransaction, error) {
	// TODO: not impl
	// c.Logger.Errorf("code.createPhoneCode blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return nil, errors.New("code.createPhoneCode not implemented")
}
