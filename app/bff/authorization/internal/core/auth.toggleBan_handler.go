// Copyright (c) 2024 The Teamgooo Authors. All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AuthToggleBan
// auth.toggleBan flags:# phone:string predefined:flags.0?true expires:flags.1?int reason:flags.1?string = PredefinedUser;
func (c *AuthorizationCore) AuthToggleBan(in *tg.TLAuthToggleBan) (*tg.PredefinedUser, error) {
	if in.Phone == "" {
		return nil, tg.ErrInputMethodInvalid
	}

	code := "00000"

	return tg.MakeTLPredefinedUser(&tg.TLPredefinedUser{
		Phone:  in.Phone,
		Code:   code,
		Banned: true,
	}).ToPredefinedUser(), nil
}
