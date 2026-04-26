// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
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
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// AuthsessionSetInitConnection
// authsession.setInitConnection auth_key_id:long ip:string api_id:int device_model:string system_version:string app_version:string system_lang_code:string lang_pack:string lang_code:string proxy:string params:string = Bool;
func (c *AuthsessionCore) AuthsessionSetInitConnection(in *authsession.TLAuthsessionSetInitConnection) (*tg.Bool, error) {
	keyData, err := c.svcCtx.Repo.ResolvePermAuthKey(c.ctx, in.AuthKeyId)
	if err != nil {
		return nil, err
	}
	in.AuthKeyId = keyData.PermAuthKeyId
	if err := c.svcCtx.Repo.SetInitConnection(c.ctx, in); err != nil {
		return nil, err
	}
	return tg.BoolTrue, nil
}
